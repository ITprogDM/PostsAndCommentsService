package mock

import (
	"OdinVOdin/internal/models"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockPosts struct {
	ctrl     *gomock.Controller
	recorder *MockPostsMockRecorder
}

// MockPostsMockRecorder is the mock recorder for MockPosts.
type MockPostsMockRecorder struct {
	mock *MockPosts
}

// NewMockPosts creates a new mock instance.
func NewMockPosts(ctrl *gomock.Controller) *MockPosts {
	mock := &MockPosts{ctrl: ctrl}
	mock.recorder = &MockPostsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPosts) EXPECT() *MockPostsMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPosts) CreatePost(post models.Post) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", post)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostsMockRecorder) CreatePost(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPosts)(nil).CreatePost), post)
}

// GetAllPosts mocks base method.
func (m *MockPosts) GetAllPosts(limit, offset int) ([]models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPosts", limit, offset)
	ret0, _ := ret[0].([]models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPosts indicates an expected call of GetAllPosts.
func (mr *MockPostsMockRecorder) GetAllPosts(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPosts", reflect.TypeOf((*MockPosts)(nil).GetAllPosts), limit, offset)
}

// GetPostById mocks base method.
func (m *MockPosts) GetPostById(id int) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostById", id)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostById indicates an expected call of GetPostById.
func (mr *MockPostsMockRecorder) GetPostById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostById", reflect.TypeOf((*MockPosts)(nil).GetPostById), id)
}

// MockComments is a mock of Comments interface.
type MockComments struct {
	ctrl     *gomock.Controller
	recorder *MockCommentsMockRecorder
}

// MockCommentsMockRecorder is the mock recorder for MockComments.
type MockCommentsMockRecorder struct {
	mock *MockComments
}

// NewMockComments creates a new mock instance.
func NewMockComments(ctrl *gomock.Controller) *MockComments {
	mock := &MockComments{ctrl: ctrl}
	mock.recorder = &MockCommentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComments) EXPECT() *MockCommentsMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockComments) CreateComment(comment models.Comment) (models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", comment)
	ret0, _ := ret[0].(models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentsMockRecorder) CreateComment(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockComments)(nil).CreateComment), comment)
}

// GetCommentsByPost mocks base method.
func (m *MockComments) GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsByPost", postId, limit, offset)
	ret0, _ := ret[0].([]*models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByPost indicates an expected call of GetCommentsByPost.
func (mr *MockCommentsMockRecorder) GetCommentsByPost(postId, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByPost", reflect.TypeOf((*MockComments)(nil).GetCommentsByPost), postId, limit, offset)
}

// GetRepliesOfComment mocks base method.
func (m *MockComments) GetRepliesOfComment(commentId int) ([]*models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepliesOfComment", commentId)
	ret0, _ := ret[0].([]*models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepliesOfComment indicates an expected call of GetRepliesOfComment.
func (mr *MockCommentsMockRecorder) GetRepliesOfComment(commentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepliesOfComment", reflect.TypeOf((*MockComments)(nil).GetRepliesOfComment), commentId)
}
