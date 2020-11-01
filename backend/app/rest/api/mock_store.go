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
//             AddGroupFunc: func(name string, studyYearID string) (string, error) {
// 	               panic("mock out the AddGroup method")
//             },
//             AddTeacherFunc: func(teacher store.Teacher) (string, error) {
// 	               panic("mock out the AddTeacher method")
//             },
//             DeleteGroupFunc: func(id string) error {
// 	               panic("mock out the DeleteGroup method")
//             },
//             DeleteTeacherFunc: func(teacherID string) error {
// 	               panic("mock out the DeleteTeacher method")
//             },
//             GetTeacherFullFunc: func(teacherID string) (store.Teacher, error) {
// 	               panic("mock out the GetTeacherFull method")
//             },
//             ListGroupsFunc: func() ([]store.Group, error) {
// 	               panic("mock out the ListGroups method")
//             },
//             ListTeachersFunc: func() ([]store.TeacherDetails, error) {
// 	               panic("mock out the ListTeachers method")
//             },
//             SetTeacherPreferencesFunc: func(teacherID string, pref store.TeacherPreferences) error {
// 	               panic("mock out the SetTeacherPreferences method")
//             },
//         }
//
//         // use mockedprivStore in code that requires privStore
//         // and then make assertions.
//
//     }
type privStoreMock struct {
	// AddGroupFunc mocks the AddGroup method.
	AddGroupFunc func(name string, studyYearID string) (string, error)

	// AddTeacherFunc mocks the AddTeacher method.
	AddTeacherFunc func(teacher store.Teacher) (string, error)

	// DeleteGroupFunc mocks the DeleteGroup method.
	DeleteGroupFunc func(id string) error

	// DeleteTeacherFunc mocks the DeleteTeacher method.
	DeleteTeacherFunc func(teacherID string) error

	// GetTeacherFullFunc mocks the GetTeacherFull method.
	GetTeacherFullFunc func(teacherID string) (store.Teacher, error)

	// ListGroupsFunc mocks the ListGroups method.
	ListGroupsFunc func() ([]store.Group, error)

	// ListTeachersFunc mocks the ListTeachers method.
	ListTeachersFunc func() ([]store.TeacherDetails, error)

	// SetTeacherPreferencesFunc mocks the SetTeacherPreferences method.
	SetTeacherPreferencesFunc func(teacherID string, pref store.TeacherPreferences) error

	// calls tracks calls to the methods.
	calls struct {
		// AddGroup holds details about calls to the AddGroup method.
		AddGroup []struct {
			// Name is the name argument value.
			Name string
			// StudyYearID is the studyYearID argument value.
			StudyYearID string
		}
		// AddTeacher holds details about calls to the AddTeacher method.
		AddTeacher []struct {
			// Teacher is the teacher argument value.
			Teacher store.Teacher
		}
		// DeleteGroup holds details about calls to the DeleteGroup method.
		DeleteGroup []struct {
			// ID is the id argument value.
			ID string
		}
		// DeleteTeacher holds details about calls to the DeleteTeacher method.
		DeleteTeacher []struct {
			// TeacherID is the teacherID argument value.
			TeacherID string
		}
		// GetTeacherFull holds details about calls to the GetTeacherFull method.
		GetTeacherFull []struct {
			// TeacherID is the teacherID argument value.
			TeacherID string
		}
		// ListGroups holds details about calls to the ListGroups method.
		ListGroups []struct {
		}
		// ListTeachers holds details about calls to the ListTeachers method.
		ListTeachers []struct {
		}
		// SetTeacherPreferences holds details about calls to the SetTeacherPreferences method.
		SetTeacherPreferences []struct {
			// TeacherID is the teacherID argument value.
			TeacherID string
			// Pref is the pref argument value.
			Pref store.TeacherPreferences
		}
	}
	lockAddGroup              sync.RWMutex
	lockAddTeacher            sync.RWMutex
	lockDeleteGroup           sync.RWMutex
	lockDeleteTeacher         sync.RWMutex
	lockGetTeacherFull        sync.RWMutex
	lockListGroups            sync.RWMutex
	lockListTeachers          sync.RWMutex
	lockSetTeacherPreferences sync.RWMutex
}

// AddGroup calls AddGroupFunc.
func (mock *privStoreMock) AddGroup(name string, studyYearID string) (string, error) {
	if mock.AddGroupFunc == nil {
		panic("privStoreMock.AddGroupFunc: method is nil but privStore.AddGroup was just called")
	}
	callInfo := struct {
		Name        string
		StudyYearID string
	}{
		Name:        name,
		StudyYearID: studyYearID,
	}
	mock.lockAddGroup.Lock()
	mock.calls.AddGroup = append(mock.calls.AddGroup, callInfo)
	mock.lockAddGroup.Unlock()
	return mock.AddGroupFunc(name, studyYearID)
}

// AddGroupCalls gets all the calls that were made to AddGroup.
// Check the length with:
//     len(mockedprivStore.AddGroupCalls())
func (mock *privStoreMock) AddGroupCalls() []struct {
	Name        string
	StudyYearID string
} {
	var calls []struct {
		Name        string
		StudyYearID string
	}
	mock.lockAddGroup.RLock()
	calls = mock.calls.AddGroup
	mock.lockAddGroup.RUnlock()
	return calls
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

// DeleteGroup calls DeleteGroupFunc.
func (mock *privStoreMock) DeleteGroup(id string) error {
	if mock.DeleteGroupFunc == nil {
		panic("privStoreMock.DeleteGroupFunc: method is nil but privStore.DeleteGroup was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockDeleteGroup.Lock()
	mock.calls.DeleteGroup = append(mock.calls.DeleteGroup, callInfo)
	mock.lockDeleteGroup.Unlock()
	return mock.DeleteGroupFunc(id)
}

// DeleteGroupCalls gets all the calls that were made to DeleteGroup.
// Check the length with:
//     len(mockedprivStore.DeleteGroupCalls())
func (mock *privStoreMock) DeleteGroupCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockDeleteGroup.RLock()
	calls = mock.calls.DeleteGroup
	mock.lockDeleteGroup.RUnlock()
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

// GetTeacherFull calls GetTeacherFullFunc.
func (mock *privStoreMock) GetTeacherFull(teacherID string) (store.Teacher, error) {
	if mock.GetTeacherFullFunc == nil {
		panic("privStoreMock.GetTeacherFullFunc: method is nil but privStore.GetTeacherFull was just called")
	}
	callInfo := struct {
		TeacherID string
	}{
		TeacherID: teacherID,
	}
	mock.lockGetTeacherFull.Lock()
	mock.calls.GetTeacherFull = append(mock.calls.GetTeacherFull, callInfo)
	mock.lockGetTeacherFull.Unlock()
	return mock.GetTeacherFullFunc(teacherID)
}

// GetTeacherFullCalls gets all the calls that were made to GetTeacherFull.
// Check the length with:
//     len(mockedprivStore.GetTeacherFullCalls())
func (mock *privStoreMock) GetTeacherFullCalls() []struct {
	TeacherID string
} {
	var calls []struct {
		TeacherID string
	}
	mock.lockGetTeacherFull.RLock()
	calls = mock.calls.GetTeacherFull
	mock.lockGetTeacherFull.RUnlock()
	return calls
}

// ListGroups calls ListGroupsFunc.
func (mock *privStoreMock) ListGroups() ([]store.Group, error) {
	if mock.ListGroupsFunc == nil {
		panic("privStoreMock.ListGroupsFunc: method is nil but privStore.ListGroups was just called")
	}
	callInfo := struct {
	}{}
	mock.lockListGroups.Lock()
	mock.calls.ListGroups = append(mock.calls.ListGroups, callInfo)
	mock.lockListGroups.Unlock()
	return mock.ListGroupsFunc()
}

// ListGroupsCalls gets all the calls that were made to ListGroups.
// Check the length with:
//     len(mockedprivStore.ListGroupsCalls())
func (mock *privStoreMock) ListGroupsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockListGroups.RLock()
	calls = mock.calls.ListGroups
	mock.lockListGroups.RUnlock()
	return calls
}

// ListTeachers calls ListTeachersFunc.
func (mock *privStoreMock) ListTeachers() ([]store.TeacherDetails, error) {
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

// SetTeacherPreferences calls SetTeacherPreferencesFunc.
func (mock *privStoreMock) SetTeacherPreferences(teacherID string, pref store.TeacherPreferences) error {
	if mock.SetTeacherPreferencesFunc == nil {
		panic("privStoreMock.SetTeacherPreferencesFunc: method is nil but privStore.SetTeacherPreferences was just called")
	}
	callInfo := struct {
		TeacherID string
		Pref      store.TeacherPreferences
	}{
		TeacherID: teacherID,
		Pref:      pref,
	}
	mock.lockSetTeacherPreferences.Lock()
	mock.calls.SetTeacherPreferences = append(mock.calls.SetTeacherPreferences, callInfo)
	mock.lockSetTeacherPreferences.Unlock()
	return mock.SetTeacherPreferencesFunc(teacherID, pref)
}

// SetTeacherPreferencesCalls gets all the calls that were made to SetTeacherPreferences.
// Check the length with:
//     len(mockedprivStore.SetTeacherPreferencesCalls())
func (mock *privStoreMock) SetTeacherPreferencesCalls() []struct {
	TeacherID string
	Pref      store.TeacherPreferences
} {
	var calls []struct {
		TeacherID string
		Pref      store.TeacherPreferences
	}
	mock.lockSetTeacherPreferences.RLock()
	calls = mock.calls.SetTeacherPreferences
	mock.lockSetTeacherPreferences.RUnlock()
	return calls
}