// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package strimziadminclient

import (
	"context"
	"net/http"
	"sync"
)

// Ensure, that DefaultApiMock does implement DefaultApi.
// If this is not the case, regenerate this file with moq.
var _ DefaultApi = &DefaultApiMock{}

// DefaultApiMock is a mock implementation of DefaultApi.
//
//     func TestSomethingThatUsesDefaultApi(t *testing.T) {
//
//         // make and configure a mocked DefaultApi
//         mockedDefaultApi := &DefaultApiMock{
//             CreateTopicFunc: func(ctx context.Context) ApiCreateTopicRequest {
// 	               panic("mock out the CreateTopic method")
//             },
//             CreateTopicExecuteFunc: func(r ApiCreateTopicRequest) (Topic, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the CreateTopicExecute method")
//             },
//             DeleteTopicFunc: func(ctx context.Context, topicName string) ApiDeleteTopicRequest {
// 	               panic("mock out the DeleteTopic method")
//             },
//             DeleteTopicExecuteFunc: func(r ApiDeleteTopicRequest) (*http.Response, GenericOpenAPIError) {
// 	               panic("mock out the DeleteTopicExecute method")
//             },
//             GetTopicFunc: func(ctx context.Context, topicName string) ApiGetTopicRequest {
// 	               panic("mock out the GetTopic method")
//             },
//             GetTopicExecuteFunc: func(r ApiGetTopicRequest) (Topic, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the GetTopicExecute method")
//             },
//             GetTopicsListFunc: func(ctx context.Context) ApiGetTopicsListRequest {
// 	               panic("mock out the GetTopicsList method")
//             },
//             GetTopicsListExecuteFunc: func(r ApiGetTopicsListRequest) (TopicsList, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the GetTopicsListExecute method")
//             },
//             MetricsFunc: func(ctx context.Context) ApiMetricsRequest {
// 	               panic("mock out the Metrics method")
//             },
//             MetricsExecuteFunc: func(r ApiMetricsRequest) (*http.Response, GenericOpenAPIError) {
// 	               panic("mock out the MetricsExecute method")
//             },
//             UpdateTopicFunc: func(ctx context.Context, topicName string) ApiUpdateTopicRequest {
// 	               panic("mock out the UpdateTopic method")
//             },
//             UpdateTopicExecuteFunc: func(r ApiUpdateTopicRequest) (Topic, *http.Response, GenericOpenAPIError) {
// 	               panic("mock out the UpdateTopicExecute method")
//             },
//         }
//
//         // use mockedDefaultApi in code that requires DefaultApi
//         // and then make assertions.
//
//     }
type DefaultApiMock struct {
	// CreateTopicFunc mocks the CreateTopic method.
	CreateTopicFunc func(ctx context.Context) ApiCreateTopicRequest

	// CreateTopicExecuteFunc mocks the CreateTopicExecute method.
	CreateTopicExecuteFunc func(r ApiCreateTopicRequest) (Topic, *http.Response, GenericOpenAPIError)

	// DeleteTopicFunc mocks the DeleteTopic method.
	DeleteTopicFunc func(ctx context.Context, topicName string) ApiDeleteTopicRequest

	// DeleteTopicExecuteFunc mocks the DeleteTopicExecute method.
	DeleteTopicExecuteFunc func(r ApiDeleteTopicRequest) (*http.Response, GenericOpenAPIError)

	// GetTopicFunc mocks the GetTopic method.
	GetTopicFunc func(ctx context.Context, topicName string) ApiGetTopicRequest

	// GetTopicExecuteFunc mocks the GetTopicExecute method.
	GetTopicExecuteFunc func(r ApiGetTopicRequest) (Topic, *http.Response, GenericOpenAPIError)

	// GetTopicsListFunc mocks the GetTopicsList method.
	GetTopicsListFunc func(ctx context.Context) ApiGetTopicsListRequest

	// GetTopicsListExecuteFunc mocks the GetTopicsListExecute method.
	GetTopicsListExecuteFunc func(r ApiGetTopicsListRequest) (TopicsList, *http.Response, GenericOpenAPIError)

	// MetricsFunc mocks the Metrics method.
	MetricsFunc func(ctx context.Context) ApiMetricsRequest

	// MetricsExecuteFunc mocks the MetricsExecute method.
	MetricsExecuteFunc func(r ApiMetricsRequest) (*http.Response, GenericOpenAPIError)

	// UpdateTopicFunc mocks the UpdateTopic method.
	UpdateTopicFunc func(ctx context.Context, topicName string) ApiUpdateTopicRequest

	// UpdateTopicExecuteFunc mocks the UpdateTopicExecute method.
	UpdateTopicExecuteFunc func(r ApiUpdateTopicRequest) (Topic, *http.Response, GenericOpenAPIError)

	// calls tracks calls to the methods.
	calls struct {
		// CreateTopic holds details about calls to the CreateTopic method.
		CreateTopic []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// CreateTopicExecute holds details about calls to the CreateTopicExecute method.
		CreateTopicExecute []struct {
			// R is the r argument value.
			R ApiCreateTopicRequest
		}
		// DeleteTopic holds details about calls to the DeleteTopic method.
		DeleteTopic []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TopicName is the topicName argument value.
			TopicName string
		}
		// DeleteTopicExecute holds details about calls to the DeleteTopicExecute method.
		DeleteTopicExecute []struct {
			// R is the r argument value.
			R ApiDeleteTopicRequest
		}
		// GetTopic holds details about calls to the GetTopic method.
		GetTopic []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TopicName is the topicName argument value.
			TopicName string
		}
		// GetTopicExecute holds details about calls to the GetTopicExecute method.
		GetTopicExecute []struct {
			// R is the r argument value.
			R ApiGetTopicRequest
		}
		// GetTopicsList holds details about calls to the GetTopicsList method.
		GetTopicsList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetTopicsListExecute holds details about calls to the GetTopicsListExecute method.
		GetTopicsListExecute []struct {
			// R is the r argument value.
			R ApiGetTopicsListRequest
		}
		// Metrics holds details about calls to the Metrics method.
		Metrics []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// MetricsExecute holds details about calls to the MetricsExecute method.
		MetricsExecute []struct {
			// R is the r argument value.
			R ApiMetricsRequest
		}
		// UpdateTopic holds details about calls to the UpdateTopic method.
		UpdateTopic []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TopicName is the topicName argument value.
			TopicName string
		}
		// UpdateTopicExecute holds details about calls to the UpdateTopicExecute method.
		UpdateTopicExecute []struct {
			// R is the r argument value.
			R ApiUpdateTopicRequest
		}
	}
	lockCreateTopic          sync.RWMutex
	lockCreateTopicExecute   sync.RWMutex
	lockDeleteTopic          sync.RWMutex
	lockDeleteTopicExecute   sync.RWMutex
	lockGetTopic             sync.RWMutex
	lockGetTopicExecute      sync.RWMutex
	lockGetTopicsList        sync.RWMutex
	lockGetTopicsListExecute sync.RWMutex
	lockMetrics              sync.RWMutex
	lockMetricsExecute       sync.RWMutex
	lockUpdateTopic          sync.RWMutex
	lockUpdateTopicExecute   sync.RWMutex
}

// CreateTopic calls CreateTopicFunc.
func (mock *DefaultApiMock) CreateTopic(ctx context.Context) ApiCreateTopicRequest {
	if mock.CreateTopicFunc == nil {
		panic("DefaultApiMock.CreateTopicFunc: method is nil but DefaultApi.CreateTopic was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCreateTopic.Lock()
	mock.calls.CreateTopic = append(mock.calls.CreateTopic, callInfo)
	mock.lockCreateTopic.Unlock()
	return mock.CreateTopicFunc(ctx)
}

// CreateTopicCalls gets all the calls that were made to CreateTopic.
// Check the length with:
//     len(mockedDefaultApi.CreateTopicCalls())
func (mock *DefaultApiMock) CreateTopicCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCreateTopic.RLock()
	calls = mock.calls.CreateTopic
	mock.lockCreateTopic.RUnlock()
	return calls
}

// CreateTopicExecute calls CreateTopicExecuteFunc.
func (mock *DefaultApiMock) CreateTopicExecute(r ApiCreateTopicRequest) (Topic, *http.Response, GenericOpenAPIError) {
	if mock.CreateTopicExecuteFunc == nil {
		panic("DefaultApiMock.CreateTopicExecuteFunc: method is nil but DefaultApi.CreateTopicExecute was just called")
	}
	callInfo := struct {
		R ApiCreateTopicRequest
	}{
		R: r,
	}
	mock.lockCreateTopicExecute.Lock()
	mock.calls.CreateTopicExecute = append(mock.calls.CreateTopicExecute, callInfo)
	mock.lockCreateTopicExecute.Unlock()
	return mock.CreateTopicExecuteFunc(r)
}

// CreateTopicExecuteCalls gets all the calls that were made to CreateTopicExecute.
// Check the length with:
//     len(mockedDefaultApi.CreateTopicExecuteCalls())
func (mock *DefaultApiMock) CreateTopicExecuteCalls() []struct {
	R ApiCreateTopicRequest
} {
	var calls []struct {
		R ApiCreateTopicRequest
	}
	mock.lockCreateTopicExecute.RLock()
	calls = mock.calls.CreateTopicExecute
	mock.lockCreateTopicExecute.RUnlock()
	return calls
}

// DeleteTopic calls DeleteTopicFunc.
func (mock *DefaultApiMock) DeleteTopic(ctx context.Context, topicName string) ApiDeleteTopicRequest {
	if mock.DeleteTopicFunc == nil {
		panic("DefaultApiMock.DeleteTopicFunc: method is nil but DefaultApi.DeleteTopic was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		TopicName string
	}{
		Ctx:       ctx,
		TopicName: topicName,
	}
	mock.lockDeleteTopic.Lock()
	mock.calls.DeleteTopic = append(mock.calls.DeleteTopic, callInfo)
	mock.lockDeleteTopic.Unlock()
	return mock.DeleteTopicFunc(ctx, topicName)
}

// DeleteTopicCalls gets all the calls that were made to DeleteTopic.
// Check the length with:
//     len(mockedDefaultApi.DeleteTopicCalls())
func (mock *DefaultApiMock) DeleteTopicCalls() []struct {
	Ctx       context.Context
	TopicName string
} {
	var calls []struct {
		Ctx       context.Context
		TopicName string
	}
	mock.lockDeleteTopic.RLock()
	calls = mock.calls.DeleteTopic
	mock.lockDeleteTopic.RUnlock()
	return calls
}

// DeleteTopicExecute calls DeleteTopicExecuteFunc.
func (mock *DefaultApiMock) DeleteTopicExecute(r ApiDeleteTopicRequest) (*http.Response, GenericOpenAPIError) {
	if mock.DeleteTopicExecuteFunc == nil {
		panic("DefaultApiMock.DeleteTopicExecuteFunc: method is nil but DefaultApi.DeleteTopicExecute was just called")
	}
	callInfo := struct {
		R ApiDeleteTopicRequest
	}{
		R: r,
	}
	mock.lockDeleteTopicExecute.Lock()
	mock.calls.DeleteTopicExecute = append(mock.calls.DeleteTopicExecute, callInfo)
	mock.lockDeleteTopicExecute.Unlock()
	return mock.DeleteTopicExecuteFunc(r)
}

// DeleteTopicExecuteCalls gets all the calls that were made to DeleteTopicExecute.
// Check the length with:
//     len(mockedDefaultApi.DeleteTopicExecuteCalls())
func (mock *DefaultApiMock) DeleteTopicExecuteCalls() []struct {
	R ApiDeleteTopicRequest
} {
	var calls []struct {
		R ApiDeleteTopicRequest
	}
	mock.lockDeleteTopicExecute.RLock()
	calls = mock.calls.DeleteTopicExecute
	mock.lockDeleteTopicExecute.RUnlock()
	return calls
}

// GetTopic calls GetTopicFunc.
func (mock *DefaultApiMock) GetTopic(ctx context.Context, topicName string) ApiGetTopicRequest {
	if mock.GetTopicFunc == nil {
		panic("DefaultApiMock.GetTopicFunc: method is nil but DefaultApi.GetTopic was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		TopicName string
	}{
		Ctx:       ctx,
		TopicName: topicName,
	}
	mock.lockGetTopic.Lock()
	mock.calls.GetTopic = append(mock.calls.GetTopic, callInfo)
	mock.lockGetTopic.Unlock()
	return mock.GetTopicFunc(ctx, topicName)
}

// GetTopicCalls gets all the calls that were made to GetTopic.
// Check the length with:
//     len(mockedDefaultApi.GetTopicCalls())
func (mock *DefaultApiMock) GetTopicCalls() []struct {
	Ctx       context.Context
	TopicName string
} {
	var calls []struct {
		Ctx       context.Context
		TopicName string
	}
	mock.lockGetTopic.RLock()
	calls = mock.calls.GetTopic
	mock.lockGetTopic.RUnlock()
	return calls
}

// GetTopicExecute calls GetTopicExecuteFunc.
func (mock *DefaultApiMock) GetTopicExecute(r ApiGetTopicRequest) (Topic, *http.Response, GenericOpenAPIError) {
	if mock.GetTopicExecuteFunc == nil {
		panic("DefaultApiMock.GetTopicExecuteFunc: method is nil but DefaultApi.GetTopicExecute was just called")
	}
	callInfo := struct {
		R ApiGetTopicRequest
	}{
		R: r,
	}
	mock.lockGetTopicExecute.Lock()
	mock.calls.GetTopicExecute = append(mock.calls.GetTopicExecute, callInfo)
	mock.lockGetTopicExecute.Unlock()
	return mock.GetTopicExecuteFunc(r)
}

// GetTopicExecuteCalls gets all the calls that were made to GetTopicExecute.
// Check the length with:
//     len(mockedDefaultApi.GetTopicExecuteCalls())
func (mock *DefaultApiMock) GetTopicExecuteCalls() []struct {
	R ApiGetTopicRequest
} {
	var calls []struct {
		R ApiGetTopicRequest
	}
	mock.lockGetTopicExecute.RLock()
	calls = mock.calls.GetTopicExecute
	mock.lockGetTopicExecute.RUnlock()
	return calls
}

// GetTopicsList calls GetTopicsListFunc.
func (mock *DefaultApiMock) GetTopicsList(ctx context.Context) ApiGetTopicsListRequest {
	if mock.GetTopicsListFunc == nil {
		panic("DefaultApiMock.GetTopicsListFunc: method is nil but DefaultApi.GetTopicsList was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetTopicsList.Lock()
	mock.calls.GetTopicsList = append(mock.calls.GetTopicsList, callInfo)
	mock.lockGetTopicsList.Unlock()
	return mock.GetTopicsListFunc(ctx)
}

// GetTopicsListCalls gets all the calls that were made to GetTopicsList.
// Check the length with:
//     len(mockedDefaultApi.GetTopicsListCalls())
func (mock *DefaultApiMock) GetTopicsListCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetTopicsList.RLock()
	calls = mock.calls.GetTopicsList
	mock.lockGetTopicsList.RUnlock()
	return calls
}

// GetTopicsListExecute calls GetTopicsListExecuteFunc.
func (mock *DefaultApiMock) GetTopicsListExecute(r ApiGetTopicsListRequest) (TopicsList, *http.Response, GenericOpenAPIError) {
	if mock.GetTopicsListExecuteFunc == nil {
		panic("DefaultApiMock.GetTopicsListExecuteFunc: method is nil but DefaultApi.GetTopicsListExecute was just called")
	}
	callInfo := struct {
		R ApiGetTopicsListRequest
	}{
		R: r,
	}
	mock.lockGetTopicsListExecute.Lock()
	mock.calls.GetTopicsListExecute = append(mock.calls.GetTopicsListExecute, callInfo)
	mock.lockGetTopicsListExecute.Unlock()
	return mock.GetTopicsListExecuteFunc(r)
}

// GetTopicsListExecuteCalls gets all the calls that were made to GetTopicsListExecute.
// Check the length with:
//     len(mockedDefaultApi.GetTopicsListExecuteCalls())
func (mock *DefaultApiMock) GetTopicsListExecuteCalls() []struct {
	R ApiGetTopicsListRequest
} {
	var calls []struct {
		R ApiGetTopicsListRequest
	}
	mock.lockGetTopicsListExecute.RLock()
	calls = mock.calls.GetTopicsListExecute
	mock.lockGetTopicsListExecute.RUnlock()
	return calls
}

// Metrics calls MetricsFunc.
func (mock *DefaultApiMock) Metrics(ctx context.Context) ApiMetricsRequest {
	if mock.MetricsFunc == nil {
		panic("DefaultApiMock.MetricsFunc: method is nil but DefaultApi.Metrics was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockMetrics.Lock()
	mock.calls.Metrics = append(mock.calls.Metrics, callInfo)
	mock.lockMetrics.Unlock()
	return mock.MetricsFunc(ctx)
}

// MetricsCalls gets all the calls that were made to Metrics.
// Check the length with:
//     len(mockedDefaultApi.MetricsCalls())
func (mock *DefaultApiMock) MetricsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockMetrics.RLock()
	calls = mock.calls.Metrics
	mock.lockMetrics.RUnlock()
	return calls
}

// MetricsExecute calls MetricsExecuteFunc.
func (mock *DefaultApiMock) MetricsExecute(r ApiMetricsRequest) (*http.Response, GenericOpenAPIError) {
	if mock.MetricsExecuteFunc == nil {
		panic("DefaultApiMock.MetricsExecuteFunc: method is nil but DefaultApi.MetricsExecute was just called")
	}
	callInfo := struct {
		R ApiMetricsRequest
	}{
		R: r,
	}
	mock.lockMetricsExecute.Lock()
	mock.calls.MetricsExecute = append(mock.calls.MetricsExecute, callInfo)
	mock.lockMetricsExecute.Unlock()
	return mock.MetricsExecuteFunc(r)
}

// MetricsExecuteCalls gets all the calls that were made to MetricsExecute.
// Check the length with:
//     len(mockedDefaultApi.MetricsExecuteCalls())
func (mock *DefaultApiMock) MetricsExecuteCalls() []struct {
	R ApiMetricsRequest
} {
	var calls []struct {
		R ApiMetricsRequest
	}
	mock.lockMetricsExecute.RLock()
	calls = mock.calls.MetricsExecute
	mock.lockMetricsExecute.RUnlock()
	return calls
}

// UpdateTopic calls UpdateTopicFunc.
func (mock *DefaultApiMock) UpdateTopic(ctx context.Context, topicName string) ApiUpdateTopicRequest {
	if mock.UpdateTopicFunc == nil {
		panic("DefaultApiMock.UpdateTopicFunc: method is nil but DefaultApi.UpdateTopic was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		TopicName string
	}{
		Ctx:       ctx,
		TopicName: topicName,
	}
	mock.lockUpdateTopic.Lock()
	mock.calls.UpdateTopic = append(mock.calls.UpdateTopic, callInfo)
	mock.lockUpdateTopic.Unlock()
	return mock.UpdateTopicFunc(ctx, topicName)
}

// UpdateTopicCalls gets all the calls that were made to UpdateTopic.
// Check the length with:
//     len(mockedDefaultApi.UpdateTopicCalls())
func (mock *DefaultApiMock) UpdateTopicCalls() []struct {
	Ctx       context.Context
	TopicName string
} {
	var calls []struct {
		Ctx       context.Context
		TopicName string
	}
	mock.lockUpdateTopic.RLock()
	calls = mock.calls.UpdateTopic
	mock.lockUpdateTopic.RUnlock()
	return calls
}

// UpdateTopicExecute calls UpdateTopicExecuteFunc.
func (mock *DefaultApiMock) UpdateTopicExecute(r ApiUpdateTopicRequest) (Topic, *http.Response, GenericOpenAPIError) {
	if mock.UpdateTopicExecuteFunc == nil {
		panic("DefaultApiMock.UpdateTopicExecuteFunc: method is nil but DefaultApi.UpdateTopicExecute was just called")
	}
	callInfo := struct {
		R ApiUpdateTopicRequest
	}{
		R: r,
	}
	mock.lockUpdateTopicExecute.Lock()
	mock.calls.UpdateTopicExecute = append(mock.calls.UpdateTopicExecute, callInfo)
	mock.lockUpdateTopicExecute.Unlock()
	return mock.UpdateTopicExecuteFunc(r)
}

// UpdateTopicExecuteCalls gets all the calls that were made to UpdateTopicExecute.
// Check the length with:
//     len(mockedDefaultApi.UpdateTopicExecuteCalls())
func (mock *DefaultApiMock) UpdateTopicExecuteCalls() []struct {
	R ApiUpdateTopicRequest
} {
	var calls []struct {
		R ApiUpdateTopicRequest
	}
	mock.lockUpdateTopicExecute.RLock()
	calls = mock.calls.UpdateTopicExecute
	mock.lockUpdateTopicExecute.RUnlock()
	return calls
}
