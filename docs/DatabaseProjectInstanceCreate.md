# DatabaseProjectInstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewDatabaseProjectInstanceCreate

`func NewDatabaseProjectInstanceCreate(name string, service string, ) *DatabaseProjectInstanceCreate`

NewDatabaseProjectInstanceCreate instantiates a new DatabaseProjectInstanceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseProjectInstanceCreateWithDefaults

`func NewDatabaseProjectInstanceCreateWithDefaults() *DatabaseProjectInstanceCreate`

NewDatabaseProjectInstanceCreateWithDefaults instantiates a new DatabaseProjectInstanceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseProjectInstanceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseProjectInstanceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseProjectInstanceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *DatabaseProjectInstanceCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DatabaseProjectInstanceCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DatabaseProjectInstanceCreate) SetService(v string)`

SetService sets Service field to given value.


### GetSource

`func (o *DatabaseProjectInstanceCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DatabaseProjectInstanceCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DatabaseProjectInstanceCreate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DatabaseProjectInstanceCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *DatabaseProjectInstanceCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DatabaseProjectInstanceCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DatabaseProjectInstanceCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DatabaseProjectInstanceCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


