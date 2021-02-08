# StorageProjectIsoCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Source** | **string** |  | 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewStorageProjectIsoCreate

`func NewStorageProjectIsoCreate(name string, source string, ) *StorageProjectIsoCreate`

NewStorageProjectIsoCreate instantiates a new StorageProjectIsoCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProjectIsoCreateWithDefaults

`func NewStorageProjectIsoCreateWithDefaults() *StorageProjectIsoCreate`

NewStorageProjectIsoCreateWithDefaults instantiates a new StorageProjectIsoCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageProjectIsoCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageProjectIsoCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageProjectIsoCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *StorageProjectIsoCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageProjectIsoCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageProjectIsoCreate) SetSource(v string)`

SetSource sets Source field to given value.


### GetTag

`func (o *StorageProjectIsoCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *StorageProjectIsoCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *StorageProjectIsoCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *StorageProjectIsoCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


