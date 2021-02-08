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

// PlanRetention struct for PlanRetention
type PlanRetention struct {
	Id *string `json:"id,omitempty"`
	Interval *string `json:"interval,omitempty"`
	Count *float32 `json:"count,omitempty"`
}

// NewPlanRetention instantiates a new PlanRetention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanRetention() *PlanRetention {
	this := PlanRetention{}
	return &this
}

// NewPlanRetentionWithDefaults instantiates a new PlanRetention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanRetentionWithDefaults() *PlanRetention {
	this := PlanRetention{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlanRetention) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRetention) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlanRetention) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlanRetention) SetId(v string) {
	o.Id = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *PlanRetention) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRetention) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *PlanRetention) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *PlanRetention) SetInterval(v string) {
	o.Interval = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PlanRetention) GetCount() float32 {
	if o == nil || o.Count == nil {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRetention) GetCountOk() (*float32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PlanRetention) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *PlanRetention) SetCount(v float32) {
	o.Count = &v
}

func (o PlanRetention) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullablePlanRetention struct {
	value *PlanRetention
	isSet bool
}

func (v NullablePlanRetention) Get() *PlanRetention {
	return v.value
}

func (v *NullablePlanRetention) Set(val *PlanRetention) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanRetention) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanRetention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanRetention(val *PlanRetention) *NullablePlanRetention {
	return &NullablePlanRetention{value: val, isSet: true}
}

func (v NullablePlanRetention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanRetention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


