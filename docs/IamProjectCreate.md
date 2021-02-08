# IamProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Organisation** | **string** |  | 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewIamProjectCreate

`func NewIamProjectCreate(name string, organisation string, ) *IamProjectCreate`

NewIamProjectCreate instantiates a new IamProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectCreateWithDefaults

`func NewIamProjectCreateWithDefaults() *IamProjectCreate`

NewIamProjectCreateWithDefaults instantiates a new IamProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamProjectCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProjectCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProjectCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOrganisation

`func (o *IamProjectCreate) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *IamProjectCreate) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *IamProjectCreate) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetTag

`func (o *IamProjectCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamProjectCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamProjectCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamProjectCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


