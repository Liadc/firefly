// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package batchmocks

import (
	batch "github.com/kaleido-io/firefly/internal/batch"
	fftypes "github.com/kaleido-io/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// BatchManager is an autogenerated mock type for the BatchManager type
type BatchManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *BatchManager) Close() {
	_m.Called()
}

// NewMessages provides a mock function with given fields:
func (_m *BatchManager) NewMessages() chan<- *fftypes.UUID {
	ret := _m.Called()

	var r0 chan<- *fftypes.UUID
	if rf, ok := ret.Get(0).(func() chan<- *fftypes.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- *fftypes.UUID)
		}
	}

	return r0
}

// RegisterDispatcher provides a mock function with given fields: batchType, handler, batchOptions
func (_m *BatchManager) RegisterDispatcher(batchType fftypes.MessageType, handler batch.DispatchHandler, batchOptions batch.BatchOptions) {
	_m.Called(batchType, handler, batchOptions)
}

// Start provides a mock function with given fields:
func (_m *BatchManager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitStop provides a mock function with given fields:
func (_m *BatchManager) WaitStop() {
	_m.Called()
}
