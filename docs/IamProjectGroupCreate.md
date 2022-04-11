# IamProjectGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Actor** | Pointer to [**[]IamActor**](IamActor.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewIamProjectGroupCreate

`func NewIamProjectGroupCreate(name string, ) *IamProjectGroupCreate`

NewIamProjectGroupCreate instantiates a new IamProjectGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectGroupCreateWithDefaults

`func NewIamProjectGroupCreateWithDefaults() *IamProjectGroupCreate`

NewIamProjectGroupCreateWithDefaults instantiates a new IamProjectGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamProjectGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProjectGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProjectGroupCreate) SetName(v string)`

SetName sets Name field to given value.


### GetActor

`func (o *IamProjectGroupCreate) GetActor() []IamActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *IamProjectGroupCreate) GetActorOk() (*[]IamActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *IamProjectGroupCreate) SetActor(v []IamActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *IamProjectGroupCreate) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTag

`func (o *IamProjectGroupCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IamProjectGroupCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IamProjectGroupCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IamProjectGroupCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


