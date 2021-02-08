# Attempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **float32** |  | [optional] 
**Error** | Pointer to **bool** |  | [optional] [default to false]
**Stdout** | Pointer to **map[string]string** |  | [optional] 
**Stderr** | Pointer to **map[string]string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 

## Methods

### NewAttempt

`func NewAttempt() *Attempt`

NewAttempt instantiates a new Attempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptWithDefaults

`func NewAttemptWithDefaults() *Attempt`

NewAttemptWithDefaults instantiates a new Attempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attempt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attempt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attempt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attempt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Attempt) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attempt) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Attempt) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Attempt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *Attempt) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Attempt) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Attempt) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *Attempt) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStdout

`func (o *Attempt) GetStdout() map[string]string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *Attempt) GetStdoutOk() (*map[string]string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *Attempt) SetStdout(v map[string]string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *Attempt) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderr

`func (o *Attempt) GetStderr() map[string]string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *Attempt) GetStderrOk() (*map[string]string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *Attempt) SetStderr(v map[string]string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *Attempt) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Attempt) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Attempt) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Attempt) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Attempt) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetAgent

`func (o *Attempt) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *Attempt) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *Attempt) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *Attempt) HasAgent() bool`

HasAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


