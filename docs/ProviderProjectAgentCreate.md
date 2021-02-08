# ProviderProjectAgentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewProviderProjectAgentCreate

`func NewProviderProjectAgentCreate(name string, service string, ) *ProviderProjectAgentCreate`

NewProviderProjectAgentCreate instantiates a new ProviderProjectAgentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderProjectAgentCreateWithDefaults

`func NewProviderProjectAgentCreateWithDefaults() *ProviderProjectAgentCreate`

NewProviderProjectAgentCreateWithDefaults instantiates a new ProviderProjectAgentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProviderProjectAgentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderProjectAgentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderProjectAgentCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *ProviderProjectAgentCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProviderProjectAgentCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProviderProjectAgentCreate) SetService(v string)`

SetService sets Service field to given value.


### GetTag

`func (o *ProviderProjectAgentCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ProviderProjectAgentCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ProviderProjectAgentCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ProviderProjectAgentCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


