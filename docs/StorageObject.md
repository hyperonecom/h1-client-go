# StorageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStorageObject

`func NewStorageObject() *StorageObject`

NewStorageObject instantiates a new StorageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageObjectWithDefaults

`func NewStorageObjectWithDefaults() *StorageObject`

NewStorageObjectWithDefaults instantiates a new StorageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StorageObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StorageObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageObject) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageObject) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageObject) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageObject) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreatedOn

`func (o *StorageObject) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *StorageObject) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *StorageObject) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *StorageObject) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


