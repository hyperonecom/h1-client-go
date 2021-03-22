# OrganisationBilling1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**BillingAddress1**](BillingAddress1.md) |  | [optional] 

## Methods

### NewOrganisationBilling1

`func NewOrganisationBilling1() *OrganisationBilling1`

NewOrganisationBilling1 instantiates a new OrganisationBilling1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationBilling1WithDefaults

`func NewOrganisationBilling1WithDefaults() *OrganisationBilling1`

NewOrganisationBilling1WithDefaults instantiates a new OrganisationBilling1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganisationBilling1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganisationBilling1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganisationBilling1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganisationBilling1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *OrganisationBilling1) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *OrganisationBilling1) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *OrganisationBilling1) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *OrganisationBilling1) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAddress

`func (o *OrganisationBilling1) GetAddress() BillingAddress1`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganisationBilling1) GetAddressOk() (*BillingAddress1, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganisationBilling1) SetAddress(v BillingAddress1)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganisationBilling1) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


