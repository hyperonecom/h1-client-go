# StorageProjectImageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | Pointer to **string** |  | [optional] [default to "564639bc052c084e2f2e3266"]
**Vm** | Pointer to **string** |  | [optional] 
**Replica** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewStorageProjectImageCreate

`func NewStorageProjectImageCreate(name string, ) *StorageProjectImageCreate`

NewStorageProjectImageCreate instantiates a new StorageProjectImageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProjectImageCreateWithDefaults

`func NewStorageProjectImageCreateWithDefaults() *StorageProjectImageCreate`

NewStorageProjectImageCreateWithDefaults instantiates a new StorageProjectImageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageProjectImageCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageProjectImageCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageProjectImageCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *StorageProjectImageCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *StorageProjectImageCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *StorageProjectImageCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *StorageProjectImageCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetVm

`func (o *StorageProjectImageCreate) GetVm() string`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *StorageProjectImageCreate) GetVmOk() (*string, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *StorageProjectImageCreate) SetVm(v string)`

SetVm sets Vm field to given value.

### HasVm

`func (o *StorageProjectImageCreate) HasVm() bool`

HasVm returns a boolean if a field has been set.

### GetReplica

`func (o *StorageProjectImageCreate) GetReplica() string`

GetReplica returns the Replica field if non-nil, zero value otherwise.

### GetReplicaOk

`func (o *StorageProjectImageCreate) GetReplicaOk() (*string, bool)`

GetReplicaOk returns a tuple with the Replica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplica

`func (o *StorageProjectImageCreate) SetReplica(v string)`

SetReplica sets Replica field to given value.

### HasReplica

`func (o *StorageProjectImageCreate) HasReplica() bool`

HasReplica returns a boolean if a field has been set.

### GetDescription

`func (o *StorageProjectImageCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageProjectImageCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageProjectImageCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageProjectImageCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTag

`func (o *StorageProjectImageCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *StorageProjectImageCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *StorageProjectImageCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *StorageProjectImageCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


