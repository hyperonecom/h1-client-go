# InvoiceSeller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Address** | Pointer to [**InvoiceSellerAddress**](InvoiceSellerAddress.md) |  | [optional] 
**Nip** | **string** |  | 

## Methods

### NewInvoiceSeller

`func NewInvoiceSeller(company string, nip string, ) *InvoiceSeller`

NewInvoiceSeller instantiates a new InvoiceSeller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSellerWithDefaults

`func NewInvoiceSellerWithDefaults() *InvoiceSeller`

NewInvoiceSellerWithDefaults instantiates a new InvoiceSeller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *InvoiceSeller) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InvoiceSeller) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InvoiceSeller) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetAddress

`func (o *InvoiceSeller) GetAddress() InvoiceSellerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InvoiceSeller) GetAddressOk() (*InvoiceSellerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InvoiceSeller) SetAddress(v InvoiceSellerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InvoiceSeller) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNip

`func (o *InvoiceSeller) GetNip() string`

GetNip returns the Nip field if non-nil, zero value otherwise.

### GetNipOk

`func (o *InvoiceSeller) GetNipOk() (*string, bool)`

GetNipOk returns a tuple with the Nip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNip

`func (o *InvoiceSeller) SetNip(v string)`

SetNip sets Nip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


