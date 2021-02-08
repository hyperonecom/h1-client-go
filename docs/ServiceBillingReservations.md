# ServiceBillingReservations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**ServiceBillingPrice**](service_billing_price.md) |  | [optional] 
**ResourcePrice** | Pointer to [**ServiceBillingPrice**](service_billing_price.md) |  | [optional] 

## Methods

### NewServiceBillingReservations

`func NewServiceBillingReservations() *ServiceBillingReservations`

NewServiceBillingReservations instantiates a new ServiceBillingReservations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBillingReservationsWithDefaults

`func NewServiceBillingReservationsWithDefaults() *ServiceBillingReservations`

NewServiceBillingReservationsWithDefaults instantiates a new ServiceBillingReservations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceBillingReservations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceBillingReservations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceBillingReservations) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceBillingReservations) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPeriod

`func (o *ServiceBillingReservations) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ServiceBillingReservations) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ServiceBillingReservations) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ServiceBillingReservations) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrice

`func (o *ServiceBillingReservations) GetPrice() ServiceBillingPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ServiceBillingReservations) GetPriceOk() (*ServiceBillingPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ServiceBillingReservations) SetPrice(v ServiceBillingPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ServiceBillingReservations) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetResourcePrice

`func (o *ServiceBillingReservations) GetResourcePrice() ServiceBillingPrice`

GetResourcePrice returns the ResourcePrice field if non-nil, zero value otherwise.

### GetResourcePriceOk

`func (o *ServiceBillingReservations) GetResourcePriceOk() (*ServiceBillingPrice, bool)`

GetResourcePriceOk returns a tuple with the ResourcePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePrice

`func (o *ServiceBillingReservations) SetResourcePrice(v ServiceBillingPrice)`

SetResourcePrice sets ResourcePrice field to given value.

### HasResourcePrice

`func (o *ServiceBillingReservations) HasResourcePrice() bool`

HasResourcePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


