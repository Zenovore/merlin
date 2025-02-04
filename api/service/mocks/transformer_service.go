// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/caraml-dev/merlin/models"
	mock "github.com/stretchr/testify/mock"

	types "github.com/caraml-dev/merlin/pkg/transformer/types"
)

// TransformerService is an autogenerated mock type for the TransformerService type
type TransformerService struct {
	mock.Mock
}

// SimulateTransformer provides a mock function with given fields: ctx, simulationPayload
func (_m *TransformerService) SimulateTransformer(ctx context.Context, simulationPayload *models.TransformerSimulation) (*types.PredictResponse, error) {
	ret := _m.Called(ctx, simulationPayload)

	var r0 *types.PredictResponse
	if rf, ok := ret.Get(0).(func(context.Context, *models.TransformerSimulation) *types.PredictResponse); ok {
		r0 = rf(ctx, simulationPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.PredictResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.TransformerSimulation) error); ok {
		r1 = rf(ctx, simulationPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransformerService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransformerService creates a new instance of TransformerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransformerService(t mockConstructorTestingTNewTransformerService) *TransformerService {
	mock := &TransformerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
