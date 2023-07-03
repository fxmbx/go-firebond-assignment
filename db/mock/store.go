// Code generated by MockGen. DO NOT EDIT.
// Source: go-firebond-assignment/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	sql "database/sql"
	db "go-firebond-assignment/db/sqlc"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetCurrentCryptoExchangeRate mocks base method.
func (m *MockStore) GetCurrentCryptoExchangeRate(arg0 context.Context, arg1 sql.NullString) ([]db.GetCurrentCryptoExchangeRateRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentCryptoExchangeRate", arg0, arg1)
	ret0, _ := ret[0].([]db.GetCurrentCryptoExchangeRateRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentCryptoExchangeRate indicates an expected call of GetCurrentCryptoExchangeRate.
func (mr *MockStoreMockRecorder) GetCurrentCryptoExchangeRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentCryptoExchangeRate", reflect.TypeOf((*MockStore)(nil).GetCurrentCryptoExchangeRate), arg0, arg1)
}

// GetCurrentCryptoFiatRate mocks base method.
func (m *MockStore) GetCurrentCryptoFiatRate(arg0 context.Context, arg1 db.GetCurrentCryptoFiatRateParams) (db.GetCurrentCryptoFiatRateRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentCryptoFiatRate", arg0, arg1)
	ret0, _ := ret[0].(db.GetCurrentCryptoFiatRateRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentCryptoFiatRate indicates an expected call of GetCurrentCryptoFiatRate.
func (mr *MockStoreMockRecorder) GetCurrentCryptoFiatRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentCryptoFiatRate", reflect.TypeOf((*MockStore)(nil).GetCurrentCryptoFiatRate), arg0, arg1)
}

// GetCurrentExchangeRates mocks base method.
func (m *MockStore) GetCurrentExchangeRates(arg0 context.Context) ([]db.GetCurrentExchangeRatesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentExchangeRates", arg0)
	ret0, _ := ret[0].([]db.GetCurrentExchangeRatesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentExchangeRates indicates an expected call of GetCurrentExchangeRates.
func (mr *MockStoreMockRecorder) GetCurrentExchangeRates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentExchangeRates", reflect.TypeOf((*MockStore)(nil).GetCurrentExchangeRates), arg0)
}

// GetExchangeRateHistory mocks base method.
func (m *MockStore) GetExchangeRateHistory(arg0 context.Context, arg1 db.GetExchangeRateHistoryParams) ([]db.GetExchangeRateHistoryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRateHistory", arg0, arg1)
	ret0, _ := ret[0].([]db.GetExchangeRateHistoryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRateHistory indicates an expected call of GetExchangeRateHistory.
func (mr *MockStoreMockRecorder) GetExchangeRateHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRateHistory", reflect.TypeOf((*MockStore)(nil).GetExchangeRateHistory), arg0, arg1)
}

// InsertExchangeRate mocks base method.
func (m *MockStore) InsertExchangeRate(arg0 context.Context, arg1 db.InsertExchangeRateParams) (db.ExchangeRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertExchangeRate", arg0, arg1)
	ret0, _ := ret[0].(db.ExchangeRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertExchangeRate indicates an expected call of InsertExchangeRate.
func (mr *MockStoreMockRecorder) InsertExchangeRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertExchangeRate", reflect.TypeOf((*MockStore)(nil).InsertExchangeRate), arg0, arg1)
}