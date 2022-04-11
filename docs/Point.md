# Point

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **map[string]string** |  | [optional] 
**Value** | Pointer to [**[]PointValue**](PointValue.md) |  | [optional] 

## Methods

### NewPoint

`func NewPoint() *Point`

NewPoint instantiates a new Point object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointWithDefaults

`func NewPointWithDefaults() *Point`

NewPointWithDefaults instantiates a new Point object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *Point) GetDimension() map[string]string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *Point) GetDimensionOk() (*map[string]string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *Point) SetDimension(v map[string]string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *Point) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetValue

`func (o *Point) GetValue() []PointValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Point) GetValueOk() (*[]PointValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Point) SetValue(v []PointValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *Point) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


