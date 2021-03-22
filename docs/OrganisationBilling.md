# OrganisationBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**InvoiceSellerAddress**](InvoiceSellerAddress.md) |  | [optional] 
**Nip** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganisationBilling

`func NewOrganisationBilling() *OrganisationBilling`

NewOrganisationBilling instantiates a new OrganisationBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationBillingWithDefaults

`func NewOrganisationBillingWithDefaults() *OrganisationBilling`

NewOrganisationBillingWithDefaults instantiates a new OrganisationBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *OrganisationBilling) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrganisationBilling) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrganisationBilling) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrganisationBilling) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCompany

`func (o *OrganisationBilling) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *OrganisationBilling) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *OrganisationBilling) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *OrganisationBilling) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *OrganisationBilling) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganisationBilling) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganisationBilling) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganisationBilling) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *OrganisationBilling) GetAddress() InvoiceSellerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganisationBilling) GetAddressOk() (*InvoiceSellerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganisationBilling) SetAddress(v InvoiceSellerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganisationBilling) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNip

`func (o *OrganisationBilling) GetNip() string`

GetNip returns the Nip field if non-nil, zero value otherwise.

### GetNipOk

`func (o *OrganisationBilling) GetNipOk() (*string, bool)`

GetNipOk returns a tuple with the Nip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNip

`func (o *OrganisationBilling) SetNip(v string)`

SetNip sets Nip field to given value.

### HasNip

`func (o *OrganisationBilling) HasNip() bool`

HasNip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


