# ReconciliationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **float32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Kind** | Pointer to **string** |  | [optional] [readonly] 
**Measurement** | Pointer to **string** |  | [optional] [readonly] 
**Profile** | Pointer to **string** |  | [optional] [readonly] 
**Resource** | Pointer to [**[]ReconciliationTargetResource**](ReconciliationTargetResource.md) |  | [optional] 
**Usage** | Pointer to **float32** |  | [optional] [readonly] 

## Methods

### NewReconciliationTarget

`func NewReconciliationTarget() *ReconciliationTarget`

NewReconciliationTarget instantiates a new ReconciliationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationTargetWithDefaults

`func NewReconciliationTargetWithDefaults() *ReconciliationTarget`

NewReconciliationTargetWithDefaults instantiates a new ReconciliationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *ReconciliationTarget) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ReconciliationTarget) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ReconciliationTarget) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ReconciliationTarget) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetId

`func (o *ReconciliationTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReconciliationTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReconciliationTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReconciliationTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ReconciliationTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReconciliationTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReconciliationTarget) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ReconciliationTarget) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeasurement

`func (o *ReconciliationTarget) GetMeasurement() string`

GetMeasurement returns the Measurement field if non-nil, zero value otherwise.

### GetMeasurementOk

`func (o *ReconciliationTarget) GetMeasurementOk() (*string, bool)`

GetMeasurementOk returns a tuple with the Measurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurement

`func (o *ReconciliationTarget) SetMeasurement(v string)`

SetMeasurement sets Measurement field to given value.

### HasMeasurement

`func (o *ReconciliationTarget) HasMeasurement() bool`

HasMeasurement returns a boolean if a field has been set.

### GetProfile

`func (o *ReconciliationTarget) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ReconciliationTarget) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ReconciliationTarget) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ReconciliationTarget) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResource

`func (o *ReconciliationTarget) GetResource() []ReconciliationTargetResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ReconciliationTarget) GetResourceOk() (*[]ReconciliationTargetResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ReconciliationTarget) SetResource(v []ReconciliationTargetResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ReconciliationTarget) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetUsage

`func (o *ReconciliationTarget) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ReconciliationTarget) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ReconciliationTarget) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ReconciliationTarget) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


