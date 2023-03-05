package dump

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime/trace"
	"strings"
	"text/template"
	"time"

	"github.com/rusq/dlog"
	"github.com/rusq/fsadapter"
	"github.com/rusq/slackdump/v2"
	"github.com/rusq/slackdump/v2/auth"
	"github.com/rusq/slackdump/v2/cmd/slackdump/internal/cfg"
	"github.com/rusq/slackdump/v2/cmd/slackdump/internal/fetch"
	"github.com/rusq/slackdump/v2/cmd/slackdump/internal/golang/base"
	"github.com/rusq/slackdump/v2/internal/app/config"
	"github.com/rusq/slackdump/v2/internal/chunk/state"
	"github.com/rusq/slackdump/v2/internal/nametmpl"
	"github.com/rusq/slackdump/v2/internal/structures"
	"github.com/rusq/slackdump/v2/internal/transform"
	"github.com/rusq/slackdump/v2/types"
	"golang.org/x/sync/errgroup"
)

//go:embed assets/list_conversation.md
var dumpMd string

// CmdDump is the dump command.
var CmdDump = &base.Command{
	UsageLine:   "slackdump dump [flags] <IDs or URLs>",
	Short:       "dump individual conversations or threads",
	Long:        dumpMd,
	RequireAuth: true,
	PrintFlags:  true,
}

func init() {
	CmdDump.Run = RunDump
	CmdDump.Wizard = WizDump
	CmdDump.Long = HelpDump(CmdDump)
}

// ErrNothingToDo is returned if there's no links to dump.
var ErrNothingToDo = errors.New("no conversations to dump, run \"slackdump help dump\"")

type options struct {
	Oldest       time.Time // Oldest is the timestamp of the oldest message to fetch.
	Latest       time.Time // Latest is the timestamp of the newest message to fetch.
	NameTemplate string    // NameTemplate is the template for the output file name.
	JSONL        bool      // JSONL should be true if the output should be JSONL instead of JSON.
}

var opts options

// ptr returns a pointer to the given value.
func ptr[T any](a T) *T { return &a }

// InitDumpFlagset initializes the flagset for the dump command.
func InitDumpFlagset(fs *flag.FlagSet) {
	fs.Var(ptr(config.TimeValue(opts.Oldest)), "from", "timestamp of the oldest message to fetch")
	fs.Var(ptr(config.TimeValue(opts.Latest)), "to", "timestamp of the newest message to fetch")
	fs.StringVar(&opts.NameTemplate, "ft", nametmpl.Default, "output file naming template.\n")
	fs.BoolVar(&opts.JSONL, "jsonl", false, "output JSONL instead of JSON")
}

// RunDump is the main entry point for the dump command.
func RunDump(ctx context.Context, cmd *base.Command, args []string) error {
	if len(args) == 0 {
		base.SetExitStatus(base.SInvalidParameters)
		return ErrNothingToDo
	}
	// Retrieve the Authentication provider.
	prov, err := auth.FromContext(ctx)
	if err != nil {
		base.SetExitStatus(base.SApplicationError)
		return err
	}

	// initialize the list of entities to dump.
	list, err := structures.NewEntityList(args)
	if err != nil {
		base.SetExitStatus(base.SInvalidParameters)
		return err
	} else if list.IsEmpty() {
		base.SetExitStatus(base.SInvalidParameters)
		return ErrNothingToDo
	}

	// initialize the file naming template.
	if opts.NameTemplate == "" {
		opts.NameTemplate = nametmpl.Default
	}
	tmpl, err := nametmpl.New(opts.NameTemplate + ".json")
	if err != nil {
		base.SetExitStatus(base.SUserError)
		return fmt.Errorf("file template error: %w", err)
	}

	sess, err := slackdump.New(ctx, prov, cfg.SlackConfig, slackdump.WithLogger(dlog.FromContext(ctx)))
	if err != nil {
		base.SetExitStatus(base.SApplicationError)
		return err
	}
	defer sess.Close()

	tmpdir, err := os.MkdirTemp("", "slackdump-*")
	if err != nil {
		base.SetExitStatus(base.SGenericError)
		return err
	}

	// Dump conversations.
	p := &fetch.Parameters{
		Oldest:    opts.Oldest,
		Latest:    opts.Latest,
		List:      list,
		DumpFiles: cfg.SlackConfig.DumpFiles,
	}

	if err := dumpv3(ctx, sess, tmpdir, tmpl, p); err != nil {
		base.SetExitStatus(base.SApplicationError)
		return err
	}
	return nil
}

func dumpv3(ctx context.Context, sess *slackdump.Session, dir string, tmpl *nametmpl.Template, p *fetch.Parameters) error {
	ctx, task := trace.NewTask(ctx, "dumpv3")
	defer task.End()

	if p == nil {
		trace.Log(ctx, "p", "nil")
		return errors.New("nil parameters")
	}
	lg := dlog.FromContext(ctx)

	lg.Printf("using temporary directory:  %s ", dir)

	tf := transform.NewStandard(sess.Filesystem(), transform.WithNameFn(tmpl.Execute))
	var eg errgroup.Group
	for _, link := range p.List.Include {
		lg.Printf("fetching %q", link)

		cr := trace.StartRegion(ctx, "fetch.Conversation")
		statefile, err := fetch.Conversation(ctx, sess, dir, link, p)
		cr.End()
		if err != nil {
			return err
		}
		eg.Go(func() error {
			return convertChunks(ctx, tf, statefile, dir)
		})
	}
	lg.Printf("waiting for all conversations to finish conversion...")
	if err := eg.Wait(); err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

func convertChunks(ctx context.Context, tf transform.Interface, statefile string, dir string) error {
	ctx, task := trace.NewTask(ctx, "convert")
	defer task.End()

	lg := dlog.FromContext(ctx)

	lg.Printf("converting %q", statefile)
	st, err := state.Load(statefile)
	if err != nil {
		return err
	}
	if err := tf.Transform(ctx, dir, st); err != nil {
		return err
	}
	return nil
}

func dumpv2(ctx context.Context, sess *slackdump.Session, list *structures.EntityList, t *nametmpl.Template) error {
	for _, link := range list.Include {
		conv, err := sess.Dump(ctx, link, opts.Oldest, opts.Latest)
		if err != nil {
			return err
		}
		name, err := t.Execute(conv)
		if err != nil {
			return err
		}
		if err := save(ctx, sess.Filesystem(), name, conv); err != nil {
			return err
		}
	}
	return nil
}

func save(ctx context.Context, fs fsadapter.FS, filename string, conv *types.Conversation) error {
	_, task := trace.NewTask(ctx, "saveData")
	defer task.End()

	f, err := fs.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(conv)
}

var helpTmpl = template.Must(template.New("dumphelp").Parse(string(dumpMd)))

// HelpDump returns the help message for the dump command.
func HelpDump(cmd *base.Command) string {
	var buf strings.Builder
	if err := helpTmpl.Execute(&buf, cmd); err != nil {
		panic(err)
	}
	return buf.String()
}
