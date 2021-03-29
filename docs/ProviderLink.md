# ProviderLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Actor** | **string** |  | 
**Purpose** | **string** |  | 
**Resource** | **string** |  | 

## Methods

### NewProviderLink

`func NewProviderLink(actor string, purpose string, resource string, ) *ProviderLink`

NewProviderLink instantiates a new ProviderLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderLinkWithDefaults

`func NewProviderLinkWithDefaults() *ProviderLink`

NewProviderLinkWithDefaults instantiates a new ProviderLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActor

`func (o *ProviderLink) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ProviderLink) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ProviderLink) SetActor(v string)`

SetActor sets Actor field to given value.


### GetPurpose

`func (o *ProviderLink) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ProviderLink) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ProviderLink) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetResource

`func (o *ProviderLink) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ProviderLink) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ProviderLink) SetResource(v string)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


