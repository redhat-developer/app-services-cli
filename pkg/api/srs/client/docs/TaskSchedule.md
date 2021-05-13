# TaskSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstExecuteAt** | **string** | ISO 8601 UTC timestamp. | 
**Priority** | Pointer to **int32** | Higher number means higher priority. Default priority is 5. | [optional] 
**IntervalSec** | Pointer to **int32** |  | [optional] 

## Methods

### NewTaskSchedule

`func NewTaskSchedule(firstExecuteAt string, ) *TaskSchedule`

NewTaskSchedule instantiates a new TaskSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskScheduleWithDefaults

`func NewTaskScheduleWithDefaults() *TaskSchedule`

NewTaskScheduleWithDefaults instantiates a new TaskSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstExecuteAt

`func (o *TaskSchedule) GetFirstExecuteAt() string`

GetFirstExecuteAt returns the FirstExecuteAt field if non-nil, zero value otherwise.

### GetFirstExecuteAtOk

`func (o *TaskSchedule) GetFirstExecuteAtOk() (*string, bool)`

GetFirstExecuteAtOk returns a tuple with the FirstExecuteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstExecuteAt

`func (o *TaskSchedule) SetFirstExecuteAt(v string)`

SetFirstExecuteAt sets FirstExecuteAt field to given value.


### GetPriority

`func (o *TaskSchedule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskSchedule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskSchedule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskSchedule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIntervalSec

`func (o *TaskSchedule) GetIntervalSec() int32`

GetIntervalSec returns the IntervalSec field if non-nil, zero value otherwise.

### GetIntervalSecOk

`func (o *TaskSchedule) GetIntervalSecOk() (*int32, bool)`

GetIntervalSecOk returns a tuple with the IntervalSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSec

`func (o *TaskSchedule) SetIntervalSec(v int32)`

SetIntervalSec sets IntervalSec field to given value.

### HasIntervalSec

`func (o *TaskSchedule) HasIntervalSec() bool`

HasIntervalSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


