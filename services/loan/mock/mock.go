// Code generated by MockGen. DO NOT EDIT.
// Source: dep.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "loan.com/models"
)

// MockloanRepo is a mock of loanRepo interface.
type MockloanRepo struct {
	ctrl     *gomock.Controller
	recorder *MockloanRepoMockRecorder
}

// MockloanRepoMockRecorder is the mock recorder for MockloanRepo.
type MockloanRepoMockRecorder struct {
	mock *MockloanRepo
}

// NewMockloanRepo creates a new mock instance.
func NewMockloanRepo(ctrl *gomock.Controller) *MockloanRepo {
	mock := &MockloanRepo{ctrl: ctrl}
	mock.recorder = &MockloanRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockloanRepo) EXPECT() *MockloanRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockloanRepo) Get(ctx context.Context, loanID int32) (models.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, loanID)
	ret0, _ := ret[0].(models.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockloanRepoMockRecorder) Get(ctx, loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockloanRepo)(nil).Get), ctx, loanID)
}

// MockpaymentRepo is a mock of paymentRepo interface.
type MockpaymentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockpaymentRepoMockRecorder
}

// MockpaymentRepoMockRecorder is the mock recorder for MockpaymentRepo.
type MockpaymentRepoMockRecorder struct {
	mock *MockpaymentRepo
}

// NewMockpaymentRepo creates a new mock instance.
func NewMockpaymentRepo(ctrl *gomock.Controller) *MockpaymentRepo {
	mock := &MockpaymentRepo{ctrl: ctrl}
	mock.recorder = &MockpaymentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpaymentRepo) EXPECT() *MockpaymentRepoMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockpaymentRepo) Count(ctx context.Context, loanID int32) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, loanID)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockpaymentRepoMockRecorder) Count(ctx, loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockpaymentRepo)(nil).Count), ctx, loanID)
}

// List mocks base method.
func (m *MockpaymentRepo) List(ctx context.Context, loanID int32) ([]models.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, loanID)
	ret0, _ := ret[0].([]models.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockpaymentRepoMockRecorder) List(ctx, loanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockpaymentRepo)(nil).List), ctx, loanID)
}
