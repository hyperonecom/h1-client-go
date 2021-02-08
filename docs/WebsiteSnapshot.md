# WebsiteSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Creation** | Pointer to **time.Time** |  | [optional] [readonly] 
**Used** | Pointer to **float32** |  | [optional] [readonly] 

## Methods

### NewWebsiteSnapshot

`func NewWebsiteSnapshot(name string, ) *WebsiteSnapshot`

NewWebsiteSnapshot instantiates a new WebsiteSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteSnapshotWithDefaults

`func NewWebsiteSnapshotWithDefaults() *WebsiteSnapshot`

NewWebsiteSnapshotWithDefaults instantiates a new WebsiteSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebsiteSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebsiteSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebsiteSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebsiteSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WebsiteSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsiteSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsiteSnapshot) SetName(v string)`

SetName sets Name field to given value.


### GetCreation

`func (o *WebsiteSnapshot) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *WebsiteSnapshot) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *WebsiteSnapshot) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.

### HasCreation

`func (o *WebsiteSnapshot) HasCreation() bool`

HasCreation returns a boolean if a field has been set.

### GetUsed

`func (o *WebsiteSnapshot) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *WebsiteSnapshot) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *WebsiteSnapshot) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *WebsiteSnapshot) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


