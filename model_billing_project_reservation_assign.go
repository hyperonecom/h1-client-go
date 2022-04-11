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

// BillingProjectReservationAssign struct for BillingProjectReservationAssign
type BillingProjectReservationAssign struct {
	Resource *string `json:"resource,omitempty"`
}

// NewBillingProjectReservationAssign instantiates a new BillingProjectReservationAssign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingProjectReservationAssign() *BillingProjectReservationAssign {
	this := BillingProjectReservationAssign{}
	return &this
}

// NewBillingProjectReservationAssignWithDefaults instantiates a new BillingProjectReservationAssign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingProjectReservationAssignWithDefaults() *BillingProjectReservationAssign {
	this := BillingProjectReservationAssign{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *BillingProjectReservationAssign) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingProjectReservationAssign) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *BillingProjectReservationAssign) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *BillingProjectReservationAssign) SetResource(v string) {
	o.Resource = &v
}

func (o BillingProjectReservationAssign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableBillingProjectReservationAssign struct {
	value *BillingProjectReservationAssign
	isSet bool
}

func (v NullableBillingProjectReservationAssign) Get() *BillingProjectReservationAssign {
	return v.value
}

func (v *NullableBillingProjectReservationAssign) Set(val *BillingProjectReservationAssign) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingProjectReservationAssign) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingProjectReservationAssign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingProjectReservationAssign(val *BillingProjectReservationAssign) *NullableBillingProjectReservationAssign {
	return &NullableBillingProjectReservationAssign{value: val, isSet: true}
}

func (v NullableBillingProjectReservationAssign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingProjectReservationAssign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


