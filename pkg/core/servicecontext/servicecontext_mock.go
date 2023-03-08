package servicecontext

import "sync"

var _ IContext = &IContextMock{}

type IContextMock struct {
	// LoadFunc mocks the Load method.
	LoadFunc func() (*Context, error)

	// LocationFunc mocks the Location method.
	LocationFunc func() (string, error)

	// RemoveFunc mocks the Remove method.
	RemoveFunc func() error

	// SaveFunc mocks the Save method.
	SaveFunc func(context *Context) error

	// calls tracks calls to the methods.
	calls struct {
		// Load holds details about calls to the Load method.
		Load []struct {
		}
		// Location holds details about calls to the Location method.
		Location []struct {
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Context is the context argument value.
			Context *Context
		}
	}
	lockLoad     sync.RWMutex
	lockLocation sync.RWMutex
	lockRemove   sync.RWMutex
	lockSave     sync.RWMutex
}

// Load calls LoadFunc.
func (mock *IContextMock) Load() (*Context, error) {
	if mock.LoadFunc == nil {
		panic("IContextMock.LoadFunc: method is nil but IContext.Load was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLoad.Lock()
	mock.calls.Load = append(mock.calls.Load, callInfo)
	mock.lockLoad.Unlock()
	return mock.LoadFunc()
}

// Location calls LocationFunc.
func (mock *IContextMock) Location() (string, error) {
	if mock.LocationFunc == nil {
		panic("IContextMock.LocationFunc: method is nil but IContext.Location was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLocation.Lock()
	mock.calls.Location = append(mock.calls.Location, callInfo)
	mock.lockLocation.Unlock()
	return mock.LocationFunc()
}

// Remove calls RemoveFunc.
func (mock *IContextMock) Remove() error {
	if mock.RemoveFunc == nil {
		panic("IContextMock.RemoveFunc: method is nil but IContext.Remove was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc()
}

// Save calls SaveFunc.
func (mock *IContextMock) Save(context *Context) error {
	if mock.SaveFunc == nil {
		panic("IContextMock.SaveFunc: method is nil but IContext.Save was just called")
	}
	callInfo := struct {
		Context *Context
	}{
		Context: context,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(context)
}
