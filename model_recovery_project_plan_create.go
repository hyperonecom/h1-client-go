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

// RecoveryProjectPlanCreate struct for RecoveryProjectPlanCreate
type RecoveryProjectPlanCreate struct {
	Name string `json:"name"`
	Window []RecoveryProjectPlanCreateWindow `json:"window,omitempty"`
	Retention []RecoveryProjectPlanCreateRetention `json:"retention,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}

// NewRecoveryProjectPlanCreate instantiates a new RecoveryProjectPlanCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryProjectPlanCreate(name string) *RecoveryProjectPlanCreate {
	this := RecoveryProjectPlanCreate{}
	this.Name = name
	return &this
}

// NewRecoveryProjectPlanCreateWithDefaults instantiates a new RecoveryProjectPlanCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryProjectPlanCreateWithDefaults() *RecoveryProjectPlanCreate {
	this := RecoveryProjectPlanCreate{}
	return &this
}

// GetName returns the Name field value
func (o *RecoveryProjectPlanCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RecoveryProjectPlanCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RecoveryProjectPlanCreate) SetName(v string) {
	o.Name = v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *RecoveryProjectPlanCreate) GetWindow() []RecoveryProjectPlanCreateWindow {
	if o == nil || o.Window == nil {
		var ret []RecoveryProjectPlanCreateWindow
		return ret
	}
	return o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryProjectPlanCreate) GetWindowOk() ([]RecoveryProjectPlanCreateWindow, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *RecoveryProjectPlanCreate) HasWindow() bool {
	if o != nil && o.Window != nil {
		return true
	}

	return false
}

// SetWindow gets a reference to the given []RecoveryProjectPlanCreateWindow and assigns it to the Window field.
func (o *RecoveryProjectPlanCreate) SetWindow(v []RecoveryProjectPlanCreateWindow) {
	o.Window = v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *RecoveryProjectPlanCreate) GetRetention() []RecoveryProjectPlanCreateRetention {
	if o == nil || o.Retention == nil {
		var ret []RecoveryProjectPlanCreateRetention
		return ret
	}
	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryProjectPlanCreate) GetRetentionOk() ([]RecoveryProjectPlanCreateRetention, bool) {
	if o == nil || o.Retention == nil {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *RecoveryProjectPlanCreate) HasRetention() bool {
	if o != nil && o.Retention != nil {
		return true
	}

	return false
}

// SetRetention gets a reference to the given []RecoveryProjectPlanCreateRetention and assigns it to the Retention field.
func (o *RecoveryProjectPlanCreate) SetRetention(v []RecoveryProjectPlanCreateRetention) {
	o.Retention = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *RecoveryProjectPlanCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryProjectPlanCreate) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *RecoveryProjectPlanCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *RecoveryProjectPlanCreate) SetTag(v []Tag) {
	o.Tag = v
}

func (o RecoveryProjectPlanCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}
	if o.Retention != nil {
		toSerialize["retention"] = o.Retention
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryProjectPlanCreate struct {
	value *RecoveryProjectPlanCreate
	isSet bool
}

func (v NullableRecoveryProjectPlanCreate) Get() *RecoveryProjectPlanCreate {
	return v.value
}

func (v *NullableRecoveryProjectPlanCreate) Set(val *RecoveryProjectPlanCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryProjectPlanCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryProjectPlanCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryProjectPlanCreate(val *RecoveryProjectPlanCreate) *NullableRecoveryProjectPlanCreate {
	return &NullableRecoveryProjectPlanCreate{value: val, isSet: true}
}

func (v NullableRecoveryProjectPlanCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryProjectPlanCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


