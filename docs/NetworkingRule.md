# NetworkingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**External** | Pointer to **[]string** |  | [optional] [default to ["0.0.0.0/0"]]
**Filter** | **[]string** |  | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Internal** | Pointer to **[]string** |  | [optional] [default to ["*"]]
**Name** | **string** |  | 
**Priority** | **float32** |  | 

## Methods

### NewNetworkingRule

`func NewNetworkingRule(action string, filter []string, name string, priority float32, ) *NetworkingRule`

NewNetworkingRule instantiates a new NetworkingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingRuleWithDefaults

`func NewNetworkingRuleWithDefaults() *NetworkingRule`

NewNetworkingRuleWithDefaults instantiates a new NetworkingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *NetworkingRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NetworkingRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NetworkingRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetExternal

`func (o *NetworkingRule) GetExternal() []string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *NetworkingRule) GetExternalOk() (*[]string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *NetworkingRule) SetExternal(v []string)`

SetExternal sets External field to given value.

### HasExternal

`func (o *NetworkingRule) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetFilter

`func (o *NetworkingRule) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *NetworkingRule) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *NetworkingRule) SetFilter(v []string)`

SetFilter sets Filter field to given value.


### GetId

`func (o *NetworkingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternal

`func (o *NetworkingRule) GetInternal() []string`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *NetworkingRule) GetInternalOk() (*[]string, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *NetworkingRule) SetInternal(v []string)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *NetworkingRule) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetName

`func (o *NetworkingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkingRule) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *NetworkingRule) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkingRule) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkingRule) SetPriority(v float32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


