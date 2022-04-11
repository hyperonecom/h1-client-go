/*
HyperOne

HyperOne API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InsightProjectJournalUpdate struct for InsightProjectJournalUpdate
type InsightProjectJournalUpdate struct {
	Name *string `json:"name,omitempty"`
	Retention *float32 `json:"retention,omitempty"`
}

// NewInsightProjectJournalUpdate instantiates a new InsightProjectJournalUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightProjectJournalUpdate() *InsightProjectJournalUpdate {
	this := InsightProjectJournalUpdate{}
	return &this
}

// NewInsightProjectJournalUpdateWithDefaults instantiates a new InsightProjectJournalUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightProjectJournalUpdateWithDefaults() *InsightProjectJournalUpdate {
	this := InsightProjectJournalUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InsightProjectJournalUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightProjectJournalUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InsightProjectJournalUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InsightProjectJournalUpdate) SetName(v string) {
	o.Name = &v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *InsightProjectJournalUpdate) GetRetention() float32 {
	if o == nil || o.Retention == nil {
		var ret float32
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightProjectJournalUpdate) GetRetentionOk() (*float32, bool) {
	if o == nil || o.Retention == nil {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *InsightProjectJournalUpdate) HasRetention() bool {
	if o != nil && o.Retention != nil {
		return true
	}

	return false
}

// SetRetention gets a reference to the given float32 and assigns it to the Retention field.
func (o *InsightProjectJournalUpdate) SetRetention(v float32) {
	o.Retention = &v
}

func (o InsightProjectJournalUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Retention != nil {
		toSerialize["retention"] = o.Retention
	}
	return json.Marshal(toSerialize)
}

type NullableInsightProjectJournalUpdate struct {
	value *InsightProjectJournalUpdate
	isSet bool
}

func (v NullableInsightProjectJournalUpdate) Get() *InsightProjectJournalUpdate {
	return v.value
}

func (v *NullableInsightProjectJournalUpdate) Set(val *InsightProjectJournalUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightProjectJournalUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightProjectJournalUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightProjectJournalUpdate(val *InsightProjectJournalUpdate) *NullableInsightProjectJournalUpdate {
	return &NullableInsightProjectJournalUpdate{value: val, isSet: true}
}

func (v NullableInsightProjectJournalUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightProjectJournalUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


