# ServiceBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**ServiceBillingPrice**](ServiceBillingPrice.md) |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**OneTime** | Pointer to **bool** |  | [optional] 
**Reservations** | Pointer to [**ServiceBillingReservations**](ServiceBillingReservations.md) |  | [optional] 

## Methods

### NewServiceBilling

`func NewServiceBilling() *ServiceBilling`

NewServiceBilling instantiates a new ServiceBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBillingWithDefaults

`func NewServiceBillingWithDefaults() *ServiceBilling`

NewServiceBillingWithDefaults instantiates a new ServiceBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *ServiceBilling) GetPrice() ServiceBillingPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ServiceBilling) GetPriceOk() (*ServiceBillingPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ServiceBilling) SetPrice(v ServiceBillingPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ServiceBilling) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPeriod

`func (o *ServiceBilling) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ServiceBilling) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ServiceBilling) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ServiceBilling) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetQuantity

`func (o *ServiceBilling) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ServiceBilling) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ServiceBilling) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ServiceBilling) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOneTime

`func (o *ServiceBilling) GetOneTime() bool`

GetOneTime returns the OneTime field if non-nil, zero value otherwise.

### GetOneTimeOk

`func (o *ServiceBilling) GetOneTimeOk() (*bool, bool)`

GetOneTimeOk returns a tuple with the OneTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTime

`func (o *ServiceBilling) SetOneTime(v bool)`

SetOneTime sets OneTime field to given value.

### HasOneTime

`func (o *ServiceBilling) HasOneTime() bool`

HasOneTime returns a boolean if a field has been set.

### GetReservations

`func (o *ServiceBilling) GetReservations() ServiceBillingReservations`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *ServiceBilling) GetReservationsOk() (*ServiceBillingReservations, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *ServiceBilling) SetReservations(v ServiceBillingReservations)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *ServiceBilling) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


