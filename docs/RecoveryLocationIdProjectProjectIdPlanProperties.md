# RecoveryLocationIdProjectProjectIdPlanProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] [readonly] 
**Window** | [**[]RecoveryLocationIdProjectProjectIdPlanPropertiesWindow**](RecoveryLocationIdProjectProjectIdPlanPropertiesWindow.md) |  | 
**Retention** | [**[]RecoveryLocationIdProjectProjectIdPlanPropertiesRetention**](RecoveryLocationIdProjectProjectIdPlanPropertiesRetention.md) |  | 

## Methods

### NewRecoveryLocationIdProjectProjectIdPlanProperties

`func NewRecoveryLocationIdProjectProjectIdPlanProperties(window []RecoveryLocationIdProjectProjectIdPlanPropertiesWindow, retention []RecoveryLocationIdProjectProjectIdPlanPropertiesRetention, ) *RecoveryLocationIdProjectProjectIdPlanProperties`

NewRecoveryLocationIdProjectProjectIdPlanProperties instantiates a new RecoveryLocationIdProjectProjectIdPlanProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryLocationIdProjectProjectIdPlanPropertiesWithDefaults

`func NewRecoveryLocationIdProjectProjectIdPlanPropertiesWithDefaults() *RecoveryLocationIdProjectProjectIdPlanProperties`

NewRecoveryLocationIdProjectProjectIdPlanPropertiesWithDefaults instantiates a new RecoveryLocationIdProjectProjectIdPlanProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWindow

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) GetWindow() []RecoveryLocationIdProjectProjectIdPlanPropertiesWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) GetWindowOk() (*[]RecoveryLocationIdProjectProjectIdPlanPropertiesWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) SetWindow(v []RecoveryLocationIdProjectProjectIdPlanPropertiesWindow)`

SetWindow sets Window field to given value.


### GetRetention

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) GetRetention() []RecoveryLocationIdProjectProjectIdPlanPropertiesRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) GetRetentionOk() (*[]RecoveryLocationIdProjectProjectIdPlanPropertiesRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RecoveryLocationIdProjectProjectIdPlanProperties) SetRetention(v []RecoveryLocationIdProjectProjectIdPlanPropertiesRetention)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


