// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	confluent_kafka_gokafka "github.com/confluentinc/confluent-kafka-go/kafka"

	mock "github.com/stretchr/testify/mock"
)

// producer is an autogenerated mock type for the producer type
type producer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *producer) Close() {
	_m.Called()
}

// Events provides a mock function with given fields:
func (_m *producer) Events() chan confluent_kafka_gokafka.Event {
	ret := _m.Called()

	var r0 chan confluent_kafka_gokafka.Event
	if rf, ok := ret.Get(0).(func() chan confluent_kafka_gokafka.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan confluent_kafka_gokafka.Event)
		}
	}

	return r0
}

// Produce provides a mock function with given fields: _a0, _a1
func (_m *producer) Produce(_a0 *confluent_kafka_gokafka.Message, _a1 chan confluent_kafka_gokafka.Event) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*confluent_kafka_gokafka.Message, chan confluent_kafka_gokafka.Event) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewProducer interface {
	mock.TestingT
	Cleanup(func())
}

// newProducer creates a new instance of producer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newProducer(t mockConstructorTestingTnewProducer) *producer {
	mock := &producer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}