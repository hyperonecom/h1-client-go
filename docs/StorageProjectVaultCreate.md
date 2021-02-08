# StorageProjectVaultCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | Pointer to **string** |  | [optional] [default to "5a0332c4eb8f4ed95c206a12"]
**Size** | **float32** |  | 
**Source** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewStorageProjectVaultCreate

`func NewStorageProjectVaultCreate(name string, size float32, ) *StorageProjectVaultCreate`

NewStorageProjectVaultCreate instantiates a new StorageProjectVaultCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProjectVaultCreateWithDefaults

`func NewStorageProjectVaultCreateWithDefaults() *StorageProjectVaultCreate`

NewStorageProjectVaultCreateWithDefaults instantiates a new StorageProjectVaultCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageProjectVaultCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageProjectVaultCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageProjectVaultCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *StorageProjectVaultCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *StorageProjectVaultCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *StorageProjectVaultCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *StorageProjectVaultCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSize

`func (o *StorageProjectVaultCreate) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageProjectVaultCreate) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageProjectVaultCreate) SetSize(v float32)`

SetSize sets Size field to given value.


### GetSource

`func (o *StorageProjectVaultCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageProjectVaultCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageProjectVaultCreate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageProjectVaultCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *StorageProjectVaultCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *StorageProjectVaultCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *StorageProjectVaultCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *StorageProjectVaultCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


