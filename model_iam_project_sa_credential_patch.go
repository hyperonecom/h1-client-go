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

// IamProjectSaCredentialPatch struct for IamProjectSaCredentialPatch
type IamProjectSaCredentialPatch struct {
	Name string `json:"name"`
}

// NewIamProjectSaCredentialPatch instantiates a new IamProjectSaCredentialPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamProjectSaCredentialPatch(name string) *IamProjectSaCredentialPatch {
	this := IamProjectSaCredentialPatch{}
	this.Name = name
	return &this
}

// NewIamProjectSaCredentialPatchWithDefaults instantiates a new IamProjectSaCredentialPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamProjectSaCredentialPatchWithDefaults() *IamProjectSaCredentialPatch {
	this := IamProjectSaCredentialPatch{}
	return &this
}

// GetName returns the Name field value
func (o *IamProjectSaCredentialPatch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IamProjectSaCredentialPatch) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IamProjectSaCredentialPatch) SetName(v string) {
	o.Name = v
}

func (o IamProjectSaCredentialPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIamProjectSaCredentialPatch struct {
	value *IamProjectSaCredentialPatch
	isSet bool
}

func (v NullableIamProjectSaCredentialPatch) Get() *IamProjectSaCredentialPatch {
	return v.value
}

func (v *NullableIamProjectSaCredentialPatch) Set(val *IamProjectSaCredentialPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableIamProjectSaCredentialPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableIamProjectSaCredentialPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamProjectSaCredentialPatch(val *IamProjectSaCredentialPatch) *NullableIamProjectSaCredentialPatch {
	return &NullableIamProjectSaCredentialPatch{value: val, isSet: true}
}

func (v NullableIamProjectSaCredentialPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamProjectSaCredentialPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


