# RecoveryProjectPlanCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Window** | Pointer to [**[]RecoveryProjectPlanCreateWindow**](RecoveryProjectPlanCreateWindow.md) |  | [optional] 
**Retention** | Pointer to [**[]RecoveryProjectPlanCreateRetention**](RecoveryProjectPlanCreateRetention.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewRecoveryProjectPlanCreate

`func NewRecoveryProjectPlanCreate(name string, ) *RecoveryProjectPlanCreate`

NewRecoveryProjectPlanCreate instantiates a new RecoveryProjectPlanCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryProjectPlanCreateWithDefaults

`func NewRecoveryProjectPlanCreateWithDefaults() *RecoveryProjectPlanCreate`

NewRecoveryProjectPlanCreateWithDefaults instantiates a new RecoveryProjectPlanCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoveryProjectPlanCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryProjectPlanCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryProjectPlanCreate) SetName(v string)`

SetName sets Name field to given value.


### GetWindow

`func (o *RecoveryProjectPlanCreate) GetWindow() []RecoveryProjectPlanCreateWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *RecoveryProjectPlanCreate) GetWindowOk() (*[]RecoveryProjectPlanCreateWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *RecoveryProjectPlanCreate) SetWindow(v []RecoveryProjectPlanCreateWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *RecoveryProjectPlanCreate) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetRetention

`func (o *RecoveryProjectPlanCreate) GetRetention() []RecoveryProjectPlanCreateRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RecoveryProjectPlanCreate) GetRetentionOk() (*[]RecoveryProjectPlanCreateRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RecoveryProjectPlanCreate) SetRetention(v []RecoveryProjectPlanCreateRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *RecoveryProjectPlanCreate) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetTag

`func (o *RecoveryProjectPlanCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RecoveryProjectPlanCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RecoveryProjectPlanCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RecoveryProjectPlanCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


