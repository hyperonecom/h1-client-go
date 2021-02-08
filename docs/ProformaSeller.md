# ProformaSeller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Address** | Pointer to [**InvoiceSellerAddress**](invoice_seller_address.md) |  | [optional] 
**Nip** | **string** |  | 
**Iban** | **string** |  | 

## Methods

### NewProformaSeller

`func NewProformaSeller(company string, nip string, iban string, ) *ProformaSeller`

NewProformaSeller instantiates a new ProformaSeller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProformaSellerWithDefaults

`func NewProformaSellerWithDefaults() *ProformaSeller`

NewProformaSellerWithDefaults instantiates a new ProformaSeller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *ProformaSeller) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ProformaSeller) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ProformaSeller) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetAddress

`func (o *ProformaSeller) GetAddress() InvoiceSellerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ProformaSeller) GetAddressOk() (*InvoiceSellerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ProformaSeller) SetAddress(v InvoiceSellerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ProformaSeller) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNip

`func (o *ProformaSeller) GetNip() string`

GetNip returns the Nip field if non-nil, zero value otherwise.

### GetNipOk

`func (o *ProformaSeller) GetNipOk() (*string, bool)`

GetNipOk returns a tuple with the Nip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNip

`func (o *ProformaSeller) SetNip(v string)`

SetNip sets Nip field to given value.


### GetIban

`func (o *ProformaSeller) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *ProformaSeller) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *ProformaSeller) SetIban(v string)`

SetIban sets Iban field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


