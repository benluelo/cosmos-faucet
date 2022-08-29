// Code generated by mockery v2.14.0. DO NOT EDIT.

package mock

import (
	actor "github.com/asynkron/protoactor-go/actor"
	ctxext "github.com/asynkron/protoactor-go/ctxext"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ActorContext is an autogenerated mock type for the Context type
type ActorContext struct {
	mock.Mock
}

// Actor provides a mock function with given fields:
func (_m *ActorContext) Actor() actor.Actor {
	ret := _m.Called()

	var r0 actor.Actor
	if rf, ok := ret.Get(0).(func() actor.Actor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actor.Actor)
		}
	}

	return r0
}

// ActorSystem provides a mock function with given fields:
func (_m *ActorContext) ActorSystem() *actor.ActorSystem {
	ret := _m.Called()

	var r0 *actor.ActorSystem
	if rf, ok := ret.Get(0).(func() *actor.ActorSystem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.ActorSystem)
		}
	}

	return r0
}

// CancelReceiveTimeout provides a mock function with given fields:
func (_m *ActorContext) CancelReceiveTimeout() {
	_m.Called()
}

// Children provides a mock function with given fields:
func (_m *ActorContext) Children() []*actor.PID {
	ret := _m.Called()

	var r0 []*actor.PID
	if rf, ok := ret.Get(0).(func() []*actor.PID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*actor.PID)
		}
	}

	return r0
}

// Forward provides a mock function with given fields: pid
func (_m *ActorContext) Forward(pid *actor.PID) {
	_m.Called(pid)
}

// Get provides a mock function with given fields: id
func (_m *ActorContext) Get(id ctxext.ContextExtensionID) ctxext.ContextExtension {
	ret := _m.Called(id)

	var r0 ctxext.ContextExtension
	if rf, ok := ret.Get(0).(func(ctxext.ContextExtensionID) ctxext.ContextExtension); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ctxext.ContextExtension)
		}
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *ActorContext) Message() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MessageHeader provides a mock function with given fields:
func (_m *ActorContext) MessageHeader() actor.ReadonlyMessageHeader {
	ret := _m.Called()

	var r0 actor.ReadonlyMessageHeader
	if rf, ok := ret.Get(0).(func() actor.ReadonlyMessageHeader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actor.ReadonlyMessageHeader)
		}
	}

	return r0
}

// Parent provides a mock function with given fields:
func (_m *ActorContext) Parent() *actor.PID {
	ret := _m.Called()

	var r0 *actor.PID
	if rf, ok := ret.Get(0).(func() *actor.PID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.PID)
		}
	}

	return r0
}

// Poison provides a mock function with given fields: pid
func (_m *ActorContext) Poison(pid *actor.PID) {
	_m.Called(pid)
}

// PoisonFuture provides a mock function with given fields: pid
func (_m *ActorContext) PoisonFuture(pid *actor.PID) *actor.Future {
	ret := _m.Called(pid)

	var r0 *actor.Future
	if rf, ok := ret.Get(0).(func(*actor.PID) *actor.Future); ok {
		r0 = rf(pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.Future)
		}
	}

	return r0
}

// Receive provides a mock function with given fields: envelope
func (_m *ActorContext) Receive(envelope *actor.MessageEnvelope) {
	_m.Called(envelope)
}

// ReceiveTimeout provides a mock function with given fields:
func (_m *ActorContext) ReceiveTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// ReenterAfter provides a mock function with given fields: f, continuation
func (_m *ActorContext) ReenterAfter(f *actor.Future, continuation func(interface{}, error)) {
	_m.Called(f, continuation)
}

// Request provides a mock function with given fields: pid, message
func (_m *ActorContext) Request(pid *actor.PID, message interface{}) {
	_m.Called(pid, message)
}

// RequestFuture provides a mock function with given fields: pid, message, timeout
func (_m *ActorContext) RequestFuture(pid *actor.PID, message interface{}, timeout time.Duration) *actor.Future {
	ret := _m.Called(pid, message, timeout)

	var r0 *actor.Future
	if rf, ok := ret.Get(0).(func(*actor.PID, interface{}, time.Duration) *actor.Future); ok {
		r0 = rf(pid, message, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.Future)
		}
	}

	return r0
}

// RequestWithCustomSender provides a mock function with given fields: pid, message, sender
func (_m *ActorContext) RequestWithCustomSender(pid *actor.PID, message interface{}, sender *actor.PID) {
	_m.Called(pid, message, sender)
}

// Respond provides a mock function with given fields: response
func (_m *ActorContext) Respond(response interface{}) {
	_m.Called(response)
}

// Self provides a mock function with given fields:
func (_m *ActorContext) Self() *actor.PID {
	ret := _m.Called()

	var r0 *actor.PID
	if rf, ok := ret.Get(0).(func() *actor.PID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.PID)
		}
	}

	return r0
}

// Send provides a mock function with given fields: pid, message
func (_m *ActorContext) Send(pid *actor.PID, message interface{}) {
	_m.Called(pid, message)
}

// Sender provides a mock function with given fields:
func (_m *ActorContext) Sender() *actor.PID {
	ret := _m.Called()

	var r0 *actor.PID
	if rf, ok := ret.Get(0).(func() *actor.PID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.PID)
		}
	}

	return r0
}

// Set provides a mock function with given fields: ext
func (_m *ActorContext) Set(ext ctxext.ContextExtension) {
	_m.Called(ext)
}

// SetReceiveTimeout provides a mock function with given fields: d
func (_m *ActorContext) SetReceiveTimeout(d time.Duration) {
	_m.Called(d)
}

// Spawn provides a mock function with given fields: props
func (_m *ActorContext) Spawn(props *actor.Props) *actor.PID {
	ret := _m.Called(props)

	var r0 *actor.PID
	if rf, ok := ret.Get(0).(func(*actor.Props) *actor.PID); ok {
		r0 = rf(props)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.PID)
		}
	}

	return r0
}

// SpawnNamed provides a mock function with given fields: props, id
func (_m *ActorContext) SpawnNamed(props *actor.Props, id string) (*actor.PID, error) {
	ret := _m.Called(props, id)

	var r0 *actor.PID
	if rf, ok := ret.Get(0).(func(*actor.Props, string) *actor.PID); ok {
		r0 = rf(props, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.PID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*actor.Props, string) error); ok {
		r1 = rf(props, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SpawnPrefix provides a mock function with given fields: props, prefix
func (_m *ActorContext) SpawnPrefix(props *actor.Props, prefix string) *actor.PID {
	ret := _m.Called(props, prefix)

	var r0 *actor.PID
	if rf, ok := ret.Get(0).(func(*actor.Props, string) *actor.PID); ok {
		r0 = rf(props, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.PID)
		}
	}

	return r0
}

// Stash provides a mock function with given fields:
func (_m *ActorContext) Stash() {
	_m.Called()
}

// Stop provides a mock function with given fields: pid
func (_m *ActorContext) Stop(pid *actor.PID) {
	_m.Called(pid)
}

// StopFuture provides a mock function with given fields: pid
func (_m *ActorContext) StopFuture(pid *actor.PID) *actor.Future {
	ret := _m.Called(pid)

	var r0 *actor.Future
	if rf, ok := ret.Get(0).(func(*actor.PID) *actor.Future); ok {
		r0 = rf(pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.Future)
		}
	}

	return r0
}

// Unwatch provides a mock function with given fields: pid
func (_m *ActorContext) Unwatch(pid *actor.PID) {
	_m.Called(pid)
}

// Watch provides a mock function with given fields: pid
func (_m *ActorContext) Watch(pid *actor.PID) {
	_m.Called(pid)
}

type mockConstructorTestingTNewMockedContext interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockedContext creates a new instance of ActorContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockedContext(t mockConstructorTestingTNewMockedContext) *ActorContext {
	mock := &ActorContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
