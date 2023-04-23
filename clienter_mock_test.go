// Code generated by MockGen. DO NOT EDIT.
// Source: slackdump.go

// Package slackdump is a generated GoMock package.
package slackdump

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	slack "github.com/slack-go/slack"
)

// Mockstreamer is a mock of streamer interface.
type Mockstreamer struct {
	ctrl     *gomock.Controller
	recorder *MockstreamerMockRecorder
}

// MockstreamerMockRecorder is the mock recorder for Mockstreamer.
type MockstreamerMockRecorder struct {
	mock *Mockstreamer
}

// NewMockstreamer creates a new mock instance.
func NewMockstreamer(ctrl *gomock.Controller) *Mockstreamer {
	mock := &Mockstreamer{ctrl: ctrl}
	mock.recorder = &MockstreamerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockstreamer) EXPECT() *MockstreamerMockRecorder {
	return m.recorder
}

// AuthTestContext mocks base method.
func (m *Mockstreamer) AuthTestContext(arg0 context.Context) (*slack.AuthTestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthTestContext", arg0)
	ret0, _ := ret[0].(*slack.AuthTestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthTestContext indicates an expected call of AuthTestContext.
func (mr *MockstreamerMockRecorder) AuthTestContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthTestContext", reflect.TypeOf((*Mockstreamer)(nil).AuthTestContext), arg0)
}

// GetConversationHistoryContext mocks base method.
func (m *Mockstreamer) GetConversationHistoryContext(ctx context.Context, params *slack.GetConversationHistoryParameters) (*slack.GetConversationHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationHistoryContext", ctx, params)
	ret0, _ := ret[0].(*slack.GetConversationHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConversationHistoryContext indicates an expected call of GetConversationHistoryContext.
func (mr *MockstreamerMockRecorder) GetConversationHistoryContext(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationHistoryContext", reflect.TypeOf((*Mockstreamer)(nil).GetConversationHistoryContext), ctx, params)
}

// GetConversationInfoContext mocks base method.
func (m *Mockstreamer) GetConversationInfoContext(ctx context.Context, input *slack.GetConversationInfoInput) (*slack.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationInfoContext", ctx, input)
	ret0, _ := ret[0].(*slack.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConversationInfoContext indicates an expected call of GetConversationInfoContext.
func (mr *MockstreamerMockRecorder) GetConversationInfoContext(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationInfoContext", reflect.TypeOf((*Mockstreamer)(nil).GetConversationInfoContext), ctx, input)
}

// GetConversationRepliesContext mocks base method.
func (m *Mockstreamer) GetConversationRepliesContext(ctx context.Context, params *slack.GetConversationRepliesParameters) ([]slack.Message, bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationRepliesContext", ctx, params)
	ret0, _ := ret[0].([]slack.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetConversationRepliesContext indicates an expected call of GetConversationRepliesContext.
func (mr *MockstreamerMockRecorder) GetConversationRepliesContext(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationRepliesContext", reflect.TypeOf((*Mockstreamer)(nil).GetConversationRepliesContext), ctx, params)
}

// GetConversationsContext mocks base method.
func (m *Mockstreamer) GetConversationsContext(ctx context.Context, params *slack.GetConversationsParameters) ([]slack.Channel, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationsContext", ctx, params)
	ret0, _ := ret[0].([]slack.Channel)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetConversationsContext indicates an expected call of GetConversationsContext.
func (mr *MockstreamerMockRecorder) GetConversationsContext(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationsContext", reflect.TypeOf((*Mockstreamer)(nil).GetConversationsContext), ctx, params)
}

// GetUsersPaginated mocks base method.
func (m *Mockstreamer) GetUsersPaginated(options ...slack.GetUsersOption) slack.UserPagination {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersPaginated", varargs...)
	ret0, _ := ret[0].(slack.UserPagination)
	return ret0
}

// GetUsersPaginated indicates an expected call of GetUsersPaginated.
func (mr *MockstreamerMockRecorder) GetUsersPaginated(options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersPaginated", reflect.TypeOf((*Mockstreamer)(nil).GetUsersPaginated), options...)
}

// mockClienter is a mock of clienter interface.
type mockClienter struct {
	ctrl     *gomock.Controller
	recorder *mockClienterMockRecorder
}

// mockClienterMockRecorder is the mock recorder for mockClienter.
type mockClienterMockRecorder struct {
	mock *mockClienter
}

// NewmockClienter creates a new mock instance.
func NewmockClienter(ctrl *gomock.Controller) *mockClienter {
	mock := &mockClienter{ctrl: ctrl}
	mock.recorder = &mockClienterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *mockClienter) EXPECT() *mockClienterMockRecorder {
	return m.recorder
}

// AuthTestContext mocks base method.
func (m *mockClienter) AuthTestContext(arg0 context.Context) (*slack.AuthTestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthTestContext", arg0)
	ret0, _ := ret[0].(*slack.AuthTestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthTestContext indicates an expected call of AuthTestContext.
func (mr *mockClienterMockRecorder) AuthTestContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthTestContext", reflect.TypeOf((*mockClienter)(nil).AuthTestContext), arg0)
}

// GetConversationHistoryContext mocks base method.
func (m *mockClienter) GetConversationHistoryContext(ctx context.Context, params *slack.GetConversationHistoryParameters) (*slack.GetConversationHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationHistoryContext", ctx, params)
	ret0, _ := ret[0].(*slack.GetConversationHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConversationHistoryContext indicates an expected call of GetConversationHistoryContext.
func (mr *mockClienterMockRecorder) GetConversationHistoryContext(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationHistoryContext", reflect.TypeOf((*mockClienter)(nil).GetConversationHistoryContext), ctx, params)
}

// GetConversationInfoContext mocks base method.
func (m *mockClienter) GetConversationInfoContext(ctx context.Context, input *slack.GetConversationInfoInput) (*slack.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationInfoContext", ctx, input)
	ret0, _ := ret[0].(*slack.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConversationInfoContext indicates an expected call of GetConversationInfoContext.
func (mr *mockClienterMockRecorder) GetConversationInfoContext(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationInfoContext", reflect.TypeOf((*mockClienter)(nil).GetConversationInfoContext), ctx, input)
}

// GetConversationRepliesContext mocks base method.
func (m *mockClienter) GetConversationRepliesContext(ctx context.Context, params *slack.GetConversationRepliesParameters) ([]slack.Message, bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationRepliesContext", ctx, params)
	ret0, _ := ret[0].([]slack.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetConversationRepliesContext indicates an expected call of GetConversationRepliesContext.
func (mr *mockClienterMockRecorder) GetConversationRepliesContext(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationRepliesContext", reflect.TypeOf((*mockClienter)(nil).GetConversationRepliesContext), ctx, params)
}

// GetConversationsContext mocks base method.
func (m *mockClienter) GetConversationsContext(ctx context.Context, params *slack.GetConversationsParameters) ([]slack.Channel, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversationsContext", ctx, params)
	ret0, _ := ret[0].([]slack.Channel)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetConversationsContext indicates an expected call of GetConversationsContext.
func (mr *mockClienterMockRecorder) GetConversationsContext(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversationsContext", reflect.TypeOf((*mockClienter)(nil).GetConversationsContext), ctx, params)
}

// GetEmojiContext mocks base method.
func (m *mockClienter) GetEmojiContext(ctx context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmojiContext", ctx)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmojiContext indicates an expected call of GetEmojiContext.
func (mr *mockClienterMockRecorder) GetEmojiContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmojiContext", reflect.TypeOf((*mockClienter)(nil).GetEmojiContext), ctx)
}

// GetFile mocks base method.
func (m *mockClienter) GetFile(downloadURL string, writer io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", downloadURL, writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetFile indicates an expected call of GetFile.
func (mr *mockClienterMockRecorder) GetFile(downloadURL, writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*mockClienter)(nil).GetFile), downloadURL, writer)
}

// GetUsersContext mocks base method.
func (m *mockClienter) GetUsersContext(ctx context.Context, options ...slack.GetUsersOption) ([]slack.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersContext", varargs...)
	ret0, _ := ret[0].([]slack.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersContext indicates an expected call of GetUsersContext.
func (mr *mockClienterMockRecorder) GetUsersContext(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersContext", reflect.TypeOf((*mockClienter)(nil).GetUsersContext), varargs...)
}

// GetUsersPaginated mocks base method.
func (m *mockClienter) GetUsersPaginated(options ...slack.GetUsersOption) slack.UserPagination {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersPaginated", varargs...)
	ret0, _ := ret[0].(slack.UserPagination)
	return ret0
}

// GetUsersPaginated indicates an expected call of GetUsersPaginated.
func (mr *mockClienterMockRecorder) GetUsersPaginated(options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersPaginated", reflect.TypeOf((*mockClienter)(nil).GetUsersPaginated), options...)
}
