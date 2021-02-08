# UserLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organisation** | Pointer to [**UserLimitOrganisation**](user_limit_organisation.md) |  | [optional] 

## Methods

### NewUserLimit

`func NewUserLimit() *UserLimit`

NewUserLimit instantiates a new UserLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLimitWithDefaults

`func NewUserLimitWithDefaults() *UserLimit`

NewUserLimitWithDefaults instantiates a new UserLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganisation

`func (o *UserLimit) GetOrganisation() UserLimitOrganisation`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *UserLimit) GetOrganisationOk() (*UserLimitOrganisation, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *UserLimit) SetOrganisation(v UserLimitOrganisation)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *UserLimit) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


