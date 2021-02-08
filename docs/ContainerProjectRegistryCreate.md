# ContainerProjectRegistryCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewContainerProjectRegistryCreate

`func NewContainerProjectRegistryCreate(name string, service string, ) *ContainerProjectRegistryCreate`

NewContainerProjectRegistryCreate instantiates a new ContainerProjectRegistryCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerProjectRegistryCreateWithDefaults

`func NewContainerProjectRegistryCreateWithDefaults() *ContainerProjectRegistryCreate`

NewContainerProjectRegistryCreateWithDefaults instantiates a new ContainerProjectRegistryCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerProjectRegistryCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerProjectRegistryCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerProjectRegistryCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *ContainerProjectRegistryCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ContainerProjectRegistryCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ContainerProjectRegistryCreate) SetService(v string)`

SetService sets Service field to given value.


### GetTag

`func (o *ContainerProjectRegistryCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ContainerProjectRegistryCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ContainerProjectRegistryCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ContainerProjectRegistryCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


