# IamProjectSaCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | Pointer to **string** |  | [optional] [default to "5e5fc76ff1fb3efe1842336a"]
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewIamProjectSaCreate

`func NewIamProjectSaCreate(name string, ) *IamProjectSaCreate`

NewIamProjectSaCreate instantiates a new IamProjectSaCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectSaCreateWithDefaults

`func NewIamProjectSaCreateWithDefaults() *IamProjectSaCreate`

NewIamProjectSaCreateWithDefaults instantiates a new IamProjectSaCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamProjectSaCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProjectSaCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProjectSaCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *IamProjectSaCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IamProjectSaCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IamProjectSaCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *IamProjectSaCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTag

`func (o *IamProjectSaCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamProjectSaCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamProjectSaCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamProjectSaCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


