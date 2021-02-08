# Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**OneTime** | Pointer to **bool** |  | [optional] 
**Service** | Pointer to [**BillingService**](billing_service.md) |  | [optional] 
**Resource** | Pointer to [**BillingResource**](billing_resource.md) |  | [optional] 
**Charges** | Pointer to [**[]BillingCharges**](BillingCharges.md) |  | [optional] 

## Methods

### NewBilling

`func NewBilling() *Billing`

NewBilling instantiates a new Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWithDefaults

`func NewBillingWithDefaults() *Billing`

NewBillingWithDefaults instantiates a new Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Billing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Billing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Billing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Billing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPeriod

`func (o *Billing) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Billing) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Billing) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Billing) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrice

`func (o *Billing) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Billing) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Billing) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Billing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *Billing) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Billing) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Billing) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Billing) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetProject

`func (o *Billing) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Billing) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Billing) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Billing) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOneTime

`func (o *Billing) GetOneTime() bool`

GetOneTime returns the OneTime field if non-nil, zero value otherwise.

### GetOneTimeOk

`func (o *Billing) GetOneTimeOk() (*bool, bool)`

GetOneTimeOk returns a tuple with the OneTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTime

`func (o *Billing) SetOneTime(v bool)`

SetOneTime sets OneTime field to given value.

### HasOneTime

`func (o *Billing) HasOneTime() bool`

HasOneTime returns a boolean if a field has been set.

### GetService

`func (o *Billing) GetService() BillingService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Billing) GetServiceOk() (*BillingService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Billing) SetService(v BillingService)`

SetService sets Service field to given value.

### HasService

`func (o *Billing) HasService() bool`

HasService returns a boolean if a field has been set.

### GetResource

`func (o *Billing) GetResource() BillingResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Billing) GetResourceOk() (*BillingResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Billing) SetResource(v BillingResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Billing) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetCharges

`func (o *Billing) GetCharges() []BillingCharges`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *Billing) GetChargesOk() (*[]BillingCharges, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *Billing) SetCharges(v []BillingCharges)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *Billing) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


