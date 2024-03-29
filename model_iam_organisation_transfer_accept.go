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

// IamOrganisationTransferAccept struct for IamOrganisationTransferAccept
type IamOrganisationTransferAccept struct {
	Payment string `json:"payment"`
}

// NewIamOrganisationTransferAccept instantiates a new IamOrganisationTransferAccept object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamOrganisationTransferAccept(payment string) *IamOrganisationTransferAccept {
	this := IamOrganisationTransferAccept{}
	this.Payment = payment
	return &this
}

// NewIamOrganisationTransferAcceptWithDefaults instantiates a new IamOrganisationTransferAccept object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamOrganisationTransferAcceptWithDefaults() *IamOrganisationTransferAccept {
	this := IamOrganisationTransferAccept{}
	return &this
}

// GetPayment returns the Payment field value
func (o *IamOrganisationTransferAccept) GetPayment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value
// and a boolean to check if the value has been set.
func (o *IamOrganisationTransferAccept) GetPaymentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payment, true
}

// SetPayment sets field value
func (o *IamOrganisationTransferAccept) SetPayment(v string) {
	o.Payment = v
}

func (o IamOrganisationTransferAccept) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment"] = o.Payment
	}
	return json.Marshal(toSerialize)
}

type NullableIamOrganisationTransferAccept struct {
	value *IamOrganisationTransferAccept
	isSet bool
}

func (v NullableIamOrganisationTransferAccept) Get() *IamOrganisationTransferAccept {
	return v.value
}

func (v *NullableIamOrganisationTransferAccept) Set(val *IamOrganisationTransferAccept) {
	v.value = val
	v.isSet = true
}

func (v NullableIamOrganisationTransferAccept) IsSet() bool {
	return v.isSet
}

func (v *NullableIamOrganisationTransferAccept) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamOrganisationTransferAccept(val *IamOrganisationTransferAccept) *NullableIamOrganisationTransferAccept {
	return &NullableIamOrganisationTransferAccept{value: val, isSet: true}
}

func (v NullableIamOrganisationTransferAccept) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamOrganisationTransferAccept) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


