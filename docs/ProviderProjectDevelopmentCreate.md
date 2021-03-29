# ProviderProjectDevelopmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Env** | Pointer to [**[]ProviderEnv**](ProviderEnv.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewProviderProjectDevelopmentCreate

`func NewProviderProjectDevelopmentCreate(name string, service string, ) *ProviderProjectDevelopmentCreate`

NewProviderProjectDevelopmentCreate instantiates a new ProviderProjectDevelopmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderProjectDevelopmentCreateWithDefaults

`func NewProviderProjectDevelopmentCreateWithDefaults() *ProviderProjectDevelopmentCreate`

NewProviderProjectDevelopmentCreateWithDefaults instantiates a new ProviderProjectDevelopmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProviderProjectDevelopmentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderProjectDevelopmentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderProjectDevelopmentCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *ProviderProjectDevelopmentCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProviderProjectDevelopmentCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProviderProjectDevelopmentCreate) SetService(v string)`

SetService sets Service field to given value.


### GetEnv

`func (o *ProviderProjectDevelopmentCreate) GetEnv() []ProviderEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ProviderProjectDevelopmentCreate) GetEnvOk() (*[]ProviderEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ProviderProjectDevelopmentCreate) SetEnv(v []ProviderEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ProviderProjectDevelopmentCreate) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetTag

`func (o *ProviderProjectDevelopmentCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ProviderProjectDevelopmentCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ProviderProjectDevelopmentCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ProviderProjectDevelopmentCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


