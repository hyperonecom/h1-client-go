# QuotaLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **float32** |  | [optional] 
**User** | Pointer to **float32** |  | [optional] 
**Effective** | Pointer to **float32** |  | [optional] [readonly] 

## Methods

### NewQuotaLimit

`func NewQuotaLimit() *QuotaLimit`

NewQuotaLimit instantiates a new QuotaLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaLimitWithDefaults

`func NewQuotaLimitWithDefaults() *QuotaLimit`

NewQuotaLimitWithDefaults instantiates a new QuotaLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *QuotaLimit) GetPlatform() float32`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *QuotaLimit) GetPlatformOk() (*float32, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *QuotaLimit) SetPlatform(v float32)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *QuotaLimit) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUser

`func (o *QuotaLimit) GetUser() float32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *QuotaLimit) GetUserOk() (*float32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *QuotaLimit) SetUser(v float32)`

SetUser sets User field to given value.

### HasUser

`func (o *QuotaLimit) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEffective

`func (o *QuotaLimit) GetEffective() float32`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *QuotaLimit) GetEffectiveOk() (*float32, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *QuotaLimit) SetEffective(v float32)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *QuotaLimit) HasEffective() bool`

HasEffective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


