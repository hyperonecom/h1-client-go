# IamProjectRoleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Permission** | Pointer to [**[]IamPermission**](IamPermission.md) |  | [optional] 
**Service** | Pointer to **string** |  | [optional] [default to "5e679c282b39c4353cd86f34"]
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewIamProjectRoleCreate

`func NewIamProjectRoleCreate(name string, ) *IamProjectRoleCreate`

NewIamProjectRoleCreate instantiates a new IamProjectRoleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectRoleCreateWithDefaults

`func NewIamProjectRoleCreateWithDefaults() *IamProjectRoleCreate`

NewIamProjectRoleCreateWithDefaults instantiates a new IamProjectRoleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IamProjectRoleCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamProjectRoleCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamProjectRoleCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamProjectRoleCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamProjectRoleCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProjectRoleCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProjectRoleCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPermission

`func (o *IamProjectRoleCreate) GetPermission() []IamPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamProjectRoleCreate) GetPermissionOk() (*[]IamPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamProjectRoleCreate) SetPermission(v []IamPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamProjectRoleCreate) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetService

`func (o *IamProjectRoleCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IamProjectRoleCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IamProjectRoleCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *IamProjectRoleCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTag

`func (o *IamProjectRoleCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamProjectRoleCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamProjectRoleCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamProjectRoleCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


