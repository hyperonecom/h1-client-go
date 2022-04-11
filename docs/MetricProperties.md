# MetricProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** |  | 
**Dimension** | Pointer to [**[]MetricPropertiesDimension**](MetricPropertiesDimension.md) |  | [optional] 

## Methods

### NewMetricProperties

`func NewMetricProperties(unit string, ) *MetricProperties`

NewMetricProperties instantiates a new MetricProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricPropertiesWithDefaults

`func NewMetricPropertiesWithDefaults() *MetricProperties`

NewMetricPropertiesWithDefaults instantiates a new MetricProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *MetricProperties) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricProperties) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricProperties) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDimension

`func (o *MetricProperties) GetDimension() []MetricPropertiesDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *MetricProperties) GetDimensionOk() (*[]MetricPropertiesDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *MetricProperties) SetDimension(v []MetricPropertiesDimension)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *MetricProperties) HasDimension() bool`

HasDimension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


