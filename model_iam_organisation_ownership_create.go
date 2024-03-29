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

// IamOrganisationOwnershipCreate struct for IamOrganisationOwnershipCreate
type IamOrganisationOwnershipCreate struct {
	Email string `json:"email"`
}

// NewIamOrganisationOwnershipCreate instantiates a new IamOrganisationOwnershipCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamOrganisationOwnershipCreate(email string) *IamOrganisationOwnershipCreate {
	this := IamOrganisationOwnershipCreate{}
	this.Email = email
	return &this
}

// NewIamOrganisationOwnershipCreateWithDefaults instantiates a new IamOrganisationOwnershipCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamOrganisationOwnershipCreateWithDefaults() *IamOrganisationOwnershipCreate {
	this := IamOrganisationOwnershipCreate{}
	return &this
}

// GetEmail returns the Email field value
func (o *IamOrganisationOwnershipCreate) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IamOrganisationOwnershipCreate) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IamOrganisationOwnershipCreate) SetEmail(v string) {
	o.Email = v
}

func (o IamOrganisationOwnershipCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableIamOrganisationOwnershipCreate struct {
	value *IamOrganisationOwnershipCreate
	isSet bool
}

func (v NullableIamOrganisationOwnershipCreate) Get() *IamOrganisationOwnershipCreate {
	return v.value
}

func (v *NullableIamOrganisationOwnershipCreate) Set(val *IamOrganisationOwnershipCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableIamOrganisationOwnershipCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableIamOrganisationOwnershipCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamOrganisationOwnershipCreate(val *IamOrganisationOwnershipCreate) *NullableIamOrganisationOwnershipCreate {
	return &NullableIamOrganisationOwnershipCreate{value: val, isSet: true}
}

func (v NullableIamOrganisationOwnershipCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamOrganisationOwnershipCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


