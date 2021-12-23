// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	markets "github.com/go-numb/go-ftx/rest/public/markets"
	mock "github.com/stretchr/testify/mock"
)

// FXTClient is an autogenerated mock type for the FXTClient type
type FXTClient struct {
	mock.Mock
}

// Markets provides a mock function with given fields: _a0
func (_m *FXTClient) Markets(_a0 *markets.RequestForMarkets) (*markets.ResponseForMarkets, error) {
	ret := _m.Called(_a0)

	var r0 *markets.ResponseForMarkets
	if rf, ok := ret.Get(0).(func(*markets.RequestForMarkets) *markets.ResponseForMarkets); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*markets.ResponseForMarkets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*markets.RequestForMarkets) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
