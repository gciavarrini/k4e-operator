// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-flotta/flotta-operator/internal/common/repository/edgedevice (interfaces: Repository)

// Package edgedevice is a generated GoMock package.
package edgedevice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/project-flotta/flotta-operator/api/v1alpha1"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 context.Context, arg1 *v1alpha1.EdgeDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1)
}

// ListForEdgeConfig mocks base method.
func (m *MockRepository) ListForEdgeConfig(arg0 context.Context, arg1, arg2 string) ([]v1alpha1.EdgeDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForEdgeConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].([]v1alpha1.EdgeDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForEdgeConfig indicates an expected call of ListForEdgeConfig.
func (mr *MockRepositoryMockRecorder) ListForEdgeConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForEdgeConfig", reflect.TypeOf((*MockRepository)(nil).ListForEdgeConfig), arg0, arg1, arg2)
}

// ListForSelector mocks base method.
func (m *MockRepository) ListForSelector(arg0 context.Context, arg1 *v1.LabelSelector, arg2 string) ([]v1alpha1.EdgeDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForSelector", arg0, arg1, arg2)
	ret0, _ := ret[0].([]v1alpha1.EdgeDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForSelector indicates an expected call of ListForSelector.
func (mr *MockRepositoryMockRecorder) ListForSelector(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForSelector", reflect.TypeOf((*MockRepository)(nil).ListForSelector), arg0, arg1, arg2)
}

// ListForWorkload mocks base method.
func (m *MockRepository) ListForWorkload(arg0 context.Context, arg1, arg2 string) ([]v1alpha1.EdgeDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForWorkload", arg0, arg1, arg2)
	ret0, _ := ret[0].([]v1alpha1.EdgeDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForWorkload indicates an expected call of ListForWorkload.
func (mr *MockRepositoryMockRecorder) ListForWorkload(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForWorkload", reflect.TypeOf((*MockRepository)(nil).ListForWorkload), arg0, arg1, arg2)
}

// Patch mocks base method.
func (m *MockRepository) Patch(arg0 context.Context, arg1, arg2 *v1alpha1.EdgeDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockRepositoryMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRepository)(nil).Patch), arg0, arg1, arg2)
}

// PatchStatus mocks base method.
func (m *MockRepository) PatchStatus(arg0 context.Context, arg1 *v1alpha1.EdgeDevice, arg2 *client.Patch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchStatus indicates an expected call of PatchStatus.
func (mr *MockRepositoryMockRecorder) PatchStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchStatus", reflect.TypeOf((*MockRepository)(nil).PatchStatus), arg0, arg1, arg2)
}

// Read mocks base method.
func (m *MockRepository) Read(arg0 context.Context, arg1, arg2 string) (*v1alpha1.EdgeDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.EdgeDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRepositoryMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRepository)(nil).Read), arg0, arg1, arg2)
}

// ReadForPlaybookExecution mocks base method.
func (m *MockRepository) ReadForPlaybookExecution(arg0 context.Context, arg1, arg2 string) (*v1alpha1.EdgeDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadForPlaybookExecution", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.EdgeDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadForPlaybookExecution indicates an expected call of ReadForPlaybookExecution.
func (mr *MockRepositoryMockRecorder) ReadForPlaybookExecution(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadForPlaybookExecution", reflect.TypeOf((*MockRepository)(nil).ReadForPlaybookExecution), arg0, arg1, arg2)
}

// RemoveFinalizer mocks base method.
func (m *MockRepository) RemoveFinalizer(arg0 context.Context, arg1 *v1alpha1.EdgeDevice, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFinalizer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFinalizer indicates an expected call of RemoveFinalizer.
func (mr *MockRepositoryMockRecorder) RemoveFinalizer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFinalizer", reflect.TypeOf((*MockRepository)(nil).RemoveFinalizer), arg0, arg1, arg2)
}

// UpdateLabels mocks base method.
func (m *MockRepository) UpdateLabels(arg0 context.Context, arg1 *v1alpha1.EdgeDevice, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLabels", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLabels indicates an expected call of UpdateLabels.
func (mr *MockRepositoryMockRecorder) UpdateLabels(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabels", reflect.TypeOf((*MockRepository)(nil).UpdateLabels), arg0, arg1, arg2)
}
