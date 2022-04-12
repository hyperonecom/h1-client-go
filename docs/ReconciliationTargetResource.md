# ReconciliationTargetResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Plan** | Pointer to [**[]ReconciliationTargetPlan**](ReconciliationTargetPlan.md) |  | [optional] 
**Usage** | Pointer to **float32** |  | [optional] [readonly] 

## Methods

### NewReconciliationTargetResource

`func NewReconciliationTargetResource() *ReconciliationTargetResource`

NewReconciliationTargetResource instantiates a new ReconciliationTargetResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationTargetResourceWithDefaults

`func NewReconciliationTargetResourceWithDefaults() *ReconciliationTargetResource`

NewReconciliationTargetResourceWithDefaults instantiates a new ReconciliationTargetResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReconciliationTargetResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReconciliationTargetResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReconciliationTargetResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReconciliationTargetResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlan

`func (o *ReconciliationTargetResource) GetPlan() []ReconciliationTargetPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ReconciliationTargetResource) GetPlanOk() (*[]ReconciliationTargetPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ReconciliationTargetResource) SetPlan(v []ReconciliationTargetPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ReconciliationTargetResource) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUsage

`func (o *ReconciliationTargetResource) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ReconciliationTargetResource) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ReconciliationTargetResource) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ReconciliationTargetResource) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


