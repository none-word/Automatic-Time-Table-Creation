// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package api

import (
	"github.com/yaattc/automatic-time-table-creation/backend/app/store"
	"sync"
)

// Ensure, that privStoreMock does implement privStore.
// If this is not the case, regenerate this file with moq.
var _ privStore = &privStoreMock{}

// privStoreMock is a mock implementation of privStore.
//
//     func TestSomethingThatUsesprivStore(t *testing.T) {
//
//         // make and configure a mocked privStore
//         mockedprivStore := &privStoreMock{
//             AddTeacherFunc: func(teacher store.Teacher) (string, error) {
// 	               panic("mock out the AddTeacher method")
//             },
//             DeleteTeacherFunc: func(teacherID string) error {
// 	               panic("mock out the DeleteTeacher method")
//             },
//             GetTeacherFunc: func(teacherID string) (store.Teacher, error) {
// 	               panic("mock out the GetTeacher method")
//             },
//             ListTeachersFunc: func() ([]store.Teacher, error) {
// 	               panic("mock out the ListTeachers method")
//             },
//         }
//
//         // use mockedprivStore in code that requires privStore
//         // and then make assertions.
//
//     }
type privStoreMock struct {
	// AddTeacherFunc mocks the AddTeacher method.
	AddTeacherFunc func(teacher store.Teacher) (string, error)

	// DeleteTeacherFunc mocks the DeleteTeacher method.
	DeleteTeacherFunc func(teacherID string) error

	// GetTeacherFunc mocks the GetTeacher method.
	GetTeacherFunc func(teacherID string) (store.Teacher, error)

	// ListTeachersFunc mocks the ListTeachers method.
	ListTeachersFunc func() ([]store.Teacher, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddTeacher holds details about calls to the AddTeacher method.
		AddTeacher []struct {
			// Teacher is the teacher argument value.
			Teacher store.Teacher
		}
		// DeleteTeacher holds details about calls to the DeleteTeacher method.
		DeleteTeacher []struct {
			// TeacherID is the teacherID argument value.
			TeacherID string
		}
		// GetTeacher holds details about calls to the GetTeacher method.
		GetTeacher []struct {
			// TeacherID is the teacherID argument value.
			TeacherID string
		}
		// ListTeachers holds details about calls to the ListTeachers method.
		ListTeachers []struct {
		}
	}
	lockAddTeacher    sync.RWMutex
	lockDeleteTeacher sync.RWMutex
	lockGetTeacher    sync.RWMutex
	lockListTeachers  sync.RWMutex
}

// AddTeacher calls AddTeacherFunc.
func (mock *privStoreMock) AddTeacher(teacher store.Teacher) (string, error) {
	if mock.AddTeacherFunc == nil {
		panic("privStoreMock.AddTeacherFunc: method is nil but privStore.AddTeacher was just called")
	}
	callInfo := struct {
		Teacher store.Teacher
	}{
		Teacher: teacher,
	}
	mock.lockAddTeacher.Lock()
	mock.calls.AddTeacher = append(mock.calls.AddTeacher, callInfo)
	mock.lockAddTeacher.Unlock()
	return mock.AddTeacherFunc(teacher)
}

// AddTeacherCalls gets all the calls that were made to AddTeacher.
// Check the length with:
//     len(mockedprivStore.AddTeacherCalls())
func (mock *privStoreMock) AddTeacherCalls() []struct {
	Teacher store.Teacher
} {
	var calls []struct {
		Teacher store.Teacher
	}
	mock.lockAddTeacher.RLock()
	calls = mock.calls.AddTeacher
	mock.lockAddTeacher.RUnlock()
	return calls
}

// DeleteTeacher calls DeleteTeacherFunc.
func (mock *privStoreMock) DeleteTeacher(teacherID string) error {
	if mock.DeleteTeacherFunc == nil {
		panic("privStoreMock.DeleteTeacherFunc: method is nil but privStore.DeleteTeacher was just called")
	}
	callInfo := struct {
		TeacherID string
	}{
		TeacherID: teacherID,
	}
	mock.lockDeleteTeacher.Lock()
	mock.calls.DeleteTeacher = append(mock.calls.DeleteTeacher, callInfo)
	mock.lockDeleteTeacher.Unlock()
	return mock.DeleteTeacherFunc(teacherID)
}

// DeleteTeacherCalls gets all the calls that were made to DeleteTeacher.
// Check the length with:
//     len(mockedprivStore.DeleteTeacherCalls())
func (mock *privStoreMock) DeleteTeacherCalls() []struct {
	TeacherID string
} {
	var calls []struct {
		TeacherID string
	}
	mock.lockDeleteTeacher.RLock()
	calls = mock.calls.DeleteTeacher
	mock.lockDeleteTeacher.RUnlock()
	return calls
}

// GetTeacher calls GetTeacherFunc.
func (mock *privStoreMock) GetTeacher(teacherID string) (store.Teacher, error) {
	if mock.GetTeacherFunc == nil {
		panic("privStoreMock.GetTeacherFunc: method is nil but privStore.GetTeacher was just called")
	}
	callInfo := struct {
		TeacherID string
	}{
		TeacherID: teacherID,
	}
	mock.lockGetTeacher.Lock()
	mock.calls.GetTeacher = append(mock.calls.GetTeacher, callInfo)
	mock.lockGetTeacher.Unlock()
	return mock.GetTeacherFunc(teacherID)
}

// GetTeacherCalls gets all the calls that were made to GetTeacher.
// Check the length with:
//     len(mockedprivStore.GetTeacherCalls())
func (mock *privStoreMock) GetTeacherCalls() []struct {
	TeacherID string
} {
	var calls []struct {
		TeacherID string
	}
	mock.lockGetTeacher.RLock()
	calls = mock.calls.GetTeacher
	mock.lockGetTeacher.RUnlock()
	return calls
}

// ListTeachers calls ListTeachersFunc.
func (mock *privStoreMock) ListTeachers() ([]store.Teacher, error) {
	if mock.ListTeachersFunc == nil {
		panic("privStoreMock.ListTeachersFunc: method is nil but privStore.ListTeachers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockListTeachers.Lock()
	mock.calls.ListTeachers = append(mock.calls.ListTeachers, callInfo)
	mock.lockListTeachers.Unlock()
	return mock.ListTeachersFunc()
}

// ListTeachersCalls gets all the calls that were made to ListTeachers.
// Check the length with:
//     len(mockedprivStore.ListTeachersCalls())
func (mock *privStoreMock) ListTeachersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockListTeachers.RLock()
	calls = mock.calls.ListTeachers
	mock.lockListTeachers.RUnlock()
	return calls
}