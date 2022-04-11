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

// IamOrganisationCreate struct for IamOrganisationCreate
type IamOrganisationCreate struct {
	Name string `json:"name"`
	Billing *OrganisationBilling `json:"billing,omitempty"`
}

// NewIamOrganisationCreate instantiates a new IamOrganisationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamOrganisationCreate(name string) *IamOrganisationCreate {
	this := IamOrganisationCreate{}
	this.Name = name
	return &this
}

// NewIamOrganisationCreateWithDefaults instantiates a new IamOrganisationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamOrganisationCreateWithDefaults() *IamOrganisationCreate {
	this := IamOrganisationCreate{}
	return &this
}

// GetName returns the Name field value
func (o *IamOrganisationCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IamOrganisationCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IamOrganisationCreate) SetName(v string) {
	o.Name = v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *IamOrganisationCreate) GetBilling() OrganisationBilling {
	if o == nil || o.Billing == nil {
		var ret OrganisationBilling
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOrganisationCreate) GetBillingOk() (*OrganisationBilling, bool) {
	if o == nil || o.Billing == nil {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *IamOrganisationCreate) HasBilling() bool {
	if o != nil && o.Billing != nil {
		return true
	}

	return false
}

// SetBilling gets a reference to the given OrganisationBilling and assigns it to the Billing field.
func (o *IamOrganisationCreate) SetBilling(v OrganisationBilling) {
	o.Billing = &v
}

func (o IamOrganisationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Billing != nil {
		toSerialize["billing"] = o.Billing
	}
	return json.Marshal(toSerialize)
}

type NullableIamOrganisationCreate struct {
	value *IamOrganisationCreate
	isSet bool
}

func (v NullableIamOrganisationCreate) Get() *IamOrganisationCreate {
	return v.value
}

func (v *NullableIamOrganisationCreate) Set(val *IamOrganisationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableIamOrganisationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableIamOrganisationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamOrganisationCreate(val *IamOrganisationCreate) *NullableIamOrganisationCreate {
	return &NullableIamOrganisationCreate{value: val, isSet: true}
}

func (v NullableIamOrganisationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamOrganisationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


