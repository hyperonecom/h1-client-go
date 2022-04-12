# StorageProjectDiskCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Size** | **float32** |  | 
**Source** | Pointer to [**OneOfAnyTypeAnyType**](oneOf&lt;AnyType,AnyType&gt;.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Vm** | Pointer to **string** |  | [optional] 

## Methods

### NewStorageProjectDiskCreate

`func NewStorageProjectDiskCreate(name string, service string, size float32, ) *StorageProjectDiskCreate`

NewStorageProjectDiskCreate instantiates a new StorageProjectDiskCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProjectDiskCreateWithDefaults

`func NewStorageProjectDiskCreateWithDefaults() *StorageProjectDiskCreate`

NewStorageProjectDiskCreateWithDefaults instantiates a new StorageProjectDiskCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageProjectDiskCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageProjectDiskCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageProjectDiskCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *StorageProjectDiskCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *StorageProjectDiskCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *StorageProjectDiskCreate) SetService(v string)`

SetService sets Service field to given value.


### GetSize

`func (o *StorageProjectDiskCreate) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageProjectDiskCreate) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageProjectDiskCreate) SetSize(v float32)`

SetSize sets Size field to given value.


### GetSource

`func (o *StorageProjectDiskCreate) GetSource() OneOfAnyTypeAnyType`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageProjectDiskCreate) GetSourceOk() (*OneOfAnyTypeAnyType, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageProjectDiskCreate) SetSource(v OneOfAnyTypeAnyType)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageProjectDiskCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *StorageProjectDiskCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *StorageProjectDiskCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *StorageProjectDiskCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *StorageProjectDiskCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVm

`func (o *StorageProjectDiskCreate) GetVm() string`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *StorageProjectDiskCreate) GetVmOk() (*string, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *StorageProjectDiskCreate) SetVm(v string)`

SetVm sets Vm field to given value.

### HasVm

`func (o *StorageProjectDiskCreate) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


