package mock_service

import (
	"OdinVOdin/internal/models"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockPostGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPostGetterMockRecorder
}

// MockPostGetterMockRecorder is the mock recorder for MockPostGetter.
type MockPostGetterMockRecorder struct {
	mock *MockPostGetter
}

// NewMockPostGetter creates a new mock instance.
func NewMockPostGetter(ctrl *gomock.Controller) *MockPostGetter {
	mock := &MockPostGetter{ctrl: ctrl}
	mock.recorder = &MockPostGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostGetter) EXPECT() *MockPostGetterMockRecorder {
	return m.recorder
}

// GetPostById mocks base method.
func (m *MockPostGetter) GetPostById(id int) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostById", id)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostById indicates an expected call of GetPostById.
func (mr *MockPostGetterMockRecorder) GetPostById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostById", reflect.TypeOf((*MockPostGetter)(nil).GetPostById), id)
}
