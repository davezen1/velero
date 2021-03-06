/*
Copyright 2018 the Heptio Ark contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import restore "github.com/heptio/velero/pkg/restore"

// ItemAction is an autogenerated mock type for the ItemAction type
type ItemAction struct {
	mock.Mock
}

// AppliesTo provides a mock function with given fields:
func (_m *ItemAction) AppliesTo() (restore.ResourceSelector, error) {
	ret := _m.Called()

	var r0 restore.ResourceSelector
	if rf, ok := ret.Get(0).(func() restore.ResourceSelector); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(restore.ResourceSelector)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Execute provides a mock function with given fields: input
func (_m *ItemAction) Execute(input *restore.RestoreItemActionExecuteInput) (*restore.RestoreItemActionExecuteOutput, error) {
	ret := _m.Called(input)

	var r0 *restore.RestoreItemActionExecuteOutput
	if rf, ok := ret.Get(0).(func(*restore.RestoreItemActionExecuteInput) *restore.RestoreItemActionExecuteOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*restore.RestoreItemActionExecuteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*restore.RestoreItemActionExecuteInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
