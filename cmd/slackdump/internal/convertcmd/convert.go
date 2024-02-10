package convertcmd

import (
	"context"
	"errors"
	"time"

	"github.com/rusq/fsadapter"
	"github.com/rusq/slackdump/v3/cmd/slackdump/internal/cfg"
	"github.com/rusq/slackdump/v3/cmd/slackdump/internal/golang/base"
	"github.com/rusq/slackdump/v3/internal/chunk"
	"github.com/rusq/slackdump/v3/internal/chunk/transform/fileproc"
	"github.com/rusq/slackdump/v3/internal/convert"
	"github.com/rusq/slackdump/v3/logger"
)

var CmdConvert = &base.Command{
	Run:       runConvert,
	UsageLine: "slackdump convert [flags] <source>",
	Short:     "convert slackdump chunks to various formats",
	Long: `
# Convert Command

Convert slackdump Chunks (output of "record") to various formats.

By default it converts a directory with chunks to an archive or directory
in Slack Export format.
`,
	CustomFlags: false,
	FlagMask:    cfg.OmitAll & ^cfg.OmitDownloadFlag &^ cfg.OmitOutputFlag,
	PrintFlags:  true,
}

type tparams struct {
	storageType fileproc.StorageType
	inputfmt    datafmt
	outputfmt   datafmt
}

var params = tparams{
	storageType: fileproc.STmattermost,
	inputfmt:    Fchunk,
	outputfmt:   Fexport,
}

func init() {
	CmdConvert.Flag.Var(&params.storageType, "storage", "storage type")
	CmdConvert.Flag.Var(&params.inputfmt, "input", "input format")
	CmdConvert.Flag.Var(&params.outputfmt, "output", "output format")
}

func runConvert(ctx context.Context, cmd *base.Command, args []string) error {
	if len(args) < 1 {
		base.SetExitStatus(base.SInvalidParameters)
		return errors.New("source and destination are required")
	}
	fn, exist := converter(params.inputfmt, params.outputfmt)
	if !exist {
		base.SetExitStatus(base.SInvalidParameters)
		return errors.New("unsupported conversion type")
	}

	lg := logger.FromContext(ctx)
	lg.Printf("converting (%s) %q to (%s) %q", params.inputfmt, args[0], params.outputfmt, cfg.Output)

	cflg := convertflags{
		withFiles: cfg.DownloadFiles,
		stt:       params.storageType,
	}
	start := time.Now()
	if err := fn(ctx, args[0], cfg.Output, cflg); err != nil {
		base.SetExitStatus(base.SApplicationError)
		return err
	}

	lg.Printf("completed in %s", time.Since(start))
	return nil
}

func converter(input datafmt, output datafmt) (convertFunc, bool) {
	if _, ok := converters[input]; !ok {
		return nil, false
	}
	if cvt, ok := converters[input][output]; ok {
		return cvt, true
	}
	return nil, false
}

type convertFunc func(ctx context.Context, input, output string, cflg convertflags) error

// ..................input.......output..............
var converters = map[datafmt]map[datafmt]convertFunc{
	Fchunk: {
		Fexport: chunk2export,
	},
}

type convertflags struct {
	withFiles bool
	stt       fileproc.StorageType
}

func chunk2export(ctx context.Context, src, trg string, cflg convertflags) error {
	cd, err := chunk.OpenDir(src)
	if err != nil {
		return err
	}
	defer cd.Close()
	fsa, err := fsadapter.New(trg)
	if err != nil {
		return err
	}
	defer fsa.Close()

	sttFn, ok := fileproc.StorageTypeFuncs[cflg.stt]
	if !ok {
		return errors.New("unknown storage type")
	}

	cvt := convert.NewChunkToExport(
		cd,
		fsa,
		convert.WithIncludeFiles(cflg.withFiles),
		convert.WithTrgFileLoc(sttFn),
		convert.WithLogger(logger.FromContext(ctx)),
	)
	if err := cvt.Convert(ctx); err != nil {
		return err
	}

	return nil
}
