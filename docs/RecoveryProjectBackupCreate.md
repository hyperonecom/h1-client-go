# RecoveryProjectBackupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Source** | [**OneOfAnyTypeAnyType**](oneOf&lt;AnyType,AnyType&gt;.md) |  | 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewRecoveryProjectBackupCreate

`func NewRecoveryProjectBackupCreate(name string, source OneOfAnyTypeAnyType, ) *RecoveryProjectBackupCreate`

NewRecoveryProjectBackupCreate instantiates a new RecoveryProjectBackupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryProjectBackupCreateWithDefaults

`func NewRecoveryProjectBackupCreateWithDefaults() *RecoveryProjectBackupCreate`

NewRecoveryProjectBackupCreateWithDefaults instantiates a new RecoveryProjectBackupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoveryProjectBackupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryProjectBackupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryProjectBackupCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *RecoveryProjectBackupCreate) GetSource() OneOfAnyTypeAnyType`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoveryProjectBackupCreate) GetSourceOk() (*OneOfAnyTypeAnyType, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoveryProjectBackupCreate) SetSource(v OneOfAnyTypeAnyType)`

SetSource sets Source field to given value.


### GetTag

`func (o *RecoveryProjectBackupCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RecoveryProjectBackupCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RecoveryProjectBackupCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RecoveryProjectBackupCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


