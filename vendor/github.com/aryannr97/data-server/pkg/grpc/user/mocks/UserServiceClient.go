// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	user "github.com/aryannr97/data-server/pkg/grpc/user"
)

// MockUserServiceClient is an autogenerated mock type for the UserServiceClient type
type MockUserServiceClient struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) AddUser(ctx context.Context, in *user.UserAddRequest, opts ...grpc.CallOption) (*user.UserAddResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *user.UserAddResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.UserAddRequest, ...grpc.CallOption) *user.UserAddResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UserAddResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.UserAddRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) DeleteUser(ctx context.Context, in *user.UserDeleteRequest, opts ...grpc.CallOption) (*user.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *user.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *user.UserDeleteRequest, ...grpc.CallOption) *user.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.UserDeleteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) GetUser(ctx context.Context, in *user.UserGetReuqest, opts ...grpc.CallOption) (*user.UserGetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *user.UserGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.UserGetReuqest, ...grpc.CallOption) *user.UserGetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UserGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.UserGetReuqest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) UpdateUser(ctx context.Context, in *user.UserUpdateRequest, opts ...grpc.CallOption) (*user.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *user.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *user.UserUpdateRequest, ...grpc.CallOption) *user.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.UserUpdateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}