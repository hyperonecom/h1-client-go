# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to [**QuotaMetric**](QuotaMetric.md) |  | [optional] 
**Usage** | Pointer to **float32** |  | [optional] 
**Limit** | Pointer to [**QuotaLimit**](QuotaLimit.md) |  | [optional] 

## Methods

### NewQuota

`func NewQuota() *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Quota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Quota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Quota) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Quota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Quota) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Quota) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Quota) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Quota) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetric

`func (o *Quota) GetMetric() QuotaMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Quota) GetMetricOk() (*QuotaMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Quota) SetMetric(v QuotaMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Quota) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetUsage

`func (o *Quota) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Quota) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Quota) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Quota) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLimit

`func (o *Quota) GetLimit() QuotaLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Quota) GetLimitOk() (*QuotaLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Quota) SetLimit(v QuotaLimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Quota) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


