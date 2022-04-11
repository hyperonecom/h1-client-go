# IamProjectApplicationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Redirect** | Pointer to [**[]IamRedirect**](IamRedirect.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewIamProjectApplicationCreate

`func NewIamProjectApplicationCreate(name string, service string, ) *IamProjectApplicationCreate`

NewIamProjectApplicationCreate instantiates a new IamProjectApplicationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectApplicationCreateWithDefaults

`func NewIamProjectApplicationCreateWithDefaults() *IamProjectApplicationCreate`

NewIamProjectApplicationCreateWithDefaults instantiates a new IamProjectApplicationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamProjectApplicationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProjectApplicationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProjectApplicationCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *IamProjectApplicationCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IamProjectApplicationCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IamProjectApplicationCreate) SetService(v string)`

SetService sets Service field to given value.


### GetRedirect

`func (o *IamProjectApplicationCreate) GetRedirect() []IamRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *IamProjectApplicationCreate) GetRedirectOk() (*[]IamRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *IamProjectApplicationCreate) SetRedirect(v []IamRedirect)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *IamProjectApplicationCreate) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetTag

`func (o *IamProjectApplicationCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamProjectApplicationCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamProjectApplicationCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamProjectApplicationCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


