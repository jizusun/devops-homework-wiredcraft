// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DependenciesInterface is an autogenerated mock type for the DependenciesInterface type
type DependenciesInterface struct {
	mock.Mock
}

// AppendToFile provides a mock function with given fields: filePath, content
func (_m *DependenciesInterface) AppendToFile(filePath string, content string) error {
	ret := _m.Called(filePath, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(filePath, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DirExists provides a mock function with given fields: path
func (_m *DependenciesInterface) DirExists(path string) (bool, error) {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecHugo provides a mock function with given fields: argString, workingDir
func (_m *DependenciesInterface) ExecHugo(argString string, workingDir string) (string, error) {
	ret := _m.Called(argString, workingDir)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(argString, workingDir)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(argString, workingDir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFortune provides a mock function with given fields:
func (_m *DependenciesInterface) GetFortune() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHugoWorkingDir provides a mock function with given fields:
func (_m *DependenciesInterface) GetHugoWorkingDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNow provides a mock function with given fields:
func (_m *DependenciesInterface) GetNow() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetWorkingDir provides a mock function with given fields:
func (_m *DependenciesInterface) GetWorkingDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// JoinPath provides a mock function with given fields: elem
func (_m *DependenciesInterface) JoinPath(elem ...string) string {
	_va := make([]interface{}, len(elem))
	for _i := range elem {
		_va[_i] = elem[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(elem...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Println provides a mock function with given fields: a
func (_m *DependenciesInterface) Println(a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// ReadFileContent provides a mock function with given fields: filename
func (_m *DependenciesInterface) ReadFileContent(filename string) ([]byte, error) {
	ret := _m.Called(filename)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
