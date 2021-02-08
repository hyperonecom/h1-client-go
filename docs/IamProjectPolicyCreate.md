# IamProjectPolicyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Role** | **string** |  | 
**Resource** | **string** |  | 
**Actor** | Pointer to [**[]IamProjectPolicyCreateActor**](IamProjectPolicyCreateActor.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewIamProjectPolicyCreate

`func NewIamProjectPolicyCreate(name string, role string, resource string, ) *IamProjectPolicyCreate`

NewIamProjectPolicyCreate instantiates a new IamProjectPolicyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectPolicyCreateWithDefaults

`func NewIamProjectPolicyCreateWithDefaults() *IamProjectPolicyCreate`

NewIamProjectPolicyCreateWithDefaults instantiates a new IamProjectPolicyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamProjectPolicyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProjectPolicyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProjectPolicyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *IamProjectPolicyCreate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IamProjectPolicyCreate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IamProjectPolicyCreate) SetRole(v string)`

SetRole sets Role field to given value.


### GetResource

`func (o *IamProjectPolicyCreate) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamProjectPolicyCreate) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamProjectPolicyCreate) SetResource(v string)`

SetResource sets Resource field to given value.


### GetActor

`func (o *IamProjectPolicyCreate) GetActor() []IamProjectPolicyCreateActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *IamProjectPolicyCreate) GetActorOk() (*[]IamProjectPolicyCreateActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *IamProjectPolicyCreate) SetActor(v []IamProjectPolicyCreateActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *IamProjectPolicyCreate) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTag

`func (o *IamProjectPolicyCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamProjectPolicyCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamProjectPolicyCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamProjectPolicyCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


