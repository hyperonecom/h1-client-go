# EnabledService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewEnabledService

`func NewEnabledService(id string, name string, ) *EnabledService`

NewEnabledService instantiates a new EnabledService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledServiceWithDefaults

`func NewEnabledServiceWithDefaults() *EnabledService`

NewEnabledServiceWithDefaults instantiates a new EnabledService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnabledService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnabledService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnabledService) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EnabledService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnabledService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnabledService) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *EnabledService) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *EnabledService) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *EnabledService) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *EnabledService) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


