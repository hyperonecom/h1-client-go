# WebsiteProjectInstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Image** | **string** |  | 
**Source** | Pointer to [**OneOfAnyTypeAnyType**](oneOf&lt;AnyType,AnyType&gt;.md) |  | [optional] 
**Env** | Pointer to [**[]WebsiteEnv**](WebsiteEnv.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewWebsiteProjectInstanceCreate

`func NewWebsiteProjectInstanceCreate(name string, service string, image string, ) *WebsiteProjectInstanceCreate`

NewWebsiteProjectInstanceCreate instantiates a new WebsiteProjectInstanceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteProjectInstanceCreateWithDefaults

`func NewWebsiteProjectInstanceCreateWithDefaults() *WebsiteProjectInstanceCreate`

NewWebsiteProjectInstanceCreateWithDefaults instantiates a new WebsiteProjectInstanceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebsiteProjectInstanceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsiteProjectInstanceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsiteProjectInstanceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *WebsiteProjectInstanceCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WebsiteProjectInstanceCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WebsiteProjectInstanceCreate) SetService(v string)`

SetService sets Service field to given value.


### GetImage

`func (o *WebsiteProjectInstanceCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WebsiteProjectInstanceCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WebsiteProjectInstanceCreate) SetImage(v string)`

SetImage sets Image field to given value.


### GetSource

`func (o *WebsiteProjectInstanceCreate) GetSource() OneOfAnyTypeAnyType`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebsiteProjectInstanceCreate) GetSourceOk() (*OneOfAnyTypeAnyType, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebsiteProjectInstanceCreate) SetSource(v OneOfAnyTypeAnyType)`

SetSource sets Source field to given value.

### HasSource

`func (o *WebsiteProjectInstanceCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEnv

`func (o *WebsiteProjectInstanceCreate) GetEnv() []WebsiteEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *WebsiteProjectInstanceCreate) GetEnvOk() (*[]WebsiteEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *WebsiteProjectInstanceCreate) SetEnv(v []WebsiteEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *WebsiteProjectInstanceCreate) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetTag

`func (o *WebsiteProjectInstanceCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *WebsiteProjectInstanceCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *WebsiteProjectInstanceCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *WebsiteProjectInstanceCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


