/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package h1

import (
	"encoding/json"
)

// Metric struct for Metric
type Metric struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
	Dimension *[]MetricDimension `json:"dimension,omitempty"`
}

// NewMetric instantiates a new Metric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetric(id string, name string, unit string) *Metric {
	this := Metric{}
	this.Id = id
	this.Name = name
	this.Unit = unit
	return &this
}

// NewMetricWithDefaults instantiates a new Metric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricWithDefaults() *Metric {
	this := Metric{}
	return &this
}

// GetId returns the Id field value
func (o *Metric) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Metric) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Metric) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Metric) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metric) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Metric) SetName(v string) {
	o.Name = v
}

// GetUnit returns the Unit field value
func (o *Metric) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Metric) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Metric) SetUnit(v string) {
	o.Unit = v
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *Metric) GetDimension() []MetricDimension {
	if o == nil || o.Dimension == nil {
		var ret []MetricDimension
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetDimensionOk() (*[]MetricDimension, bool) {
	if o == nil || o.Dimension == nil {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *Metric) HasDimension() bool {
	if o != nil && o.Dimension != nil {
		return true
	}

	return false
}

// SetDimension gets a reference to the given []MetricDimension and assigns it to the Dimension field.
func (o *Metric) SetDimension(v []MetricDimension) {
	o.Dimension = &v
}

func (o Metric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if o.Dimension != nil {
		toSerialize["dimension"] = o.Dimension
	}
	return json.Marshal(toSerialize)
}

type NullableMetric struct {
	value *Metric
	isSet bool
}

func (v NullableMetric) Get() *Metric {
	return v.value
}

func (v *NullableMetric) Set(val *Metric) {
	v.value = val
	v.isSet = true
}

func (v NullableMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetric(val *Metric) *NullableMetric {
	return &NullableMetric{value: val, isSet: true}
}

func (v NullableMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


