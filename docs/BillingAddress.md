# BillingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** |  | 
**Country** | Pointer to **string** |  | [optional] [default to "PL"]
**Street** | **string** |  | 
**Zipcode** | **string** |  | 

## Methods

### NewBillingAddress

`func NewBillingAddress(city string, street string, zipcode string, ) *BillingAddress`

NewBillingAddress instantiates a new BillingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddressWithDefaults

`func NewBillingAddressWithDefaults() *BillingAddress`

NewBillingAddressWithDefaults instantiates a new BillingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *BillingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *BillingAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BillingAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStreet

`func (o *BillingAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *BillingAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *BillingAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipcode

`func (o *BillingAddress) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *BillingAddress) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *BillingAddress) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


