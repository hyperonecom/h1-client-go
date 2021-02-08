# InvoiceBuyer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Address** | Pointer to [**InvoiceBuyerAddress**](invoice_buyer_address.md) |  | [optional] 
**Nip** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceBuyer

`func NewInvoiceBuyer(company string, ) *InvoiceBuyer`

NewInvoiceBuyer instantiates a new InvoiceBuyer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceBuyerWithDefaults

`func NewInvoiceBuyerWithDefaults() *InvoiceBuyer`

NewInvoiceBuyerWithDefaults instantiates a new InvoiceBuyer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *InvoiceBuyer) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InvoiceBuyer) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InvoiceBuyer) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetAddress

`func (o *InvoiceBuyer) GetAddress() InvoiceBuyerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InvoiceBuyer) GetAddressOk() (*InvoiceBuyerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InvoiceBuyer) SetAddress(v InvoiceBuyerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InvoiceBuyer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNip

`func (o *InvoiceBuyer) GetNip() string`

GetNip returns the Nip field if non-nil, zero value otherwise.

### GetNipOk

`func (o *InvoiceBuyer) GetNipOk() (*string, bool)`

GetNipOk returns a tuple with the Nip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNip

`func (o *InvoiceBuyer) SetNip(v string)`

SetNip sets Nip field to given value.

### HasNip

`func (o *InvoiceBuyer) HasNip() bool`

HasNip returns a boolean if a field has been set.

### GetEmail

`func (o *InvoiceBuyer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvoiceBuyer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvoiceBuyer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InvoiceBuyer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


