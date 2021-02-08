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

// ProjectBilling struct for ProjectBilling
type ProjectBilling struct {
	Credits *float32 `json:"credits,omitempty"`
	CreditsBonus *float32 `json:"creditsBonus,omitempty"`
	CreditLimit *float32 `json:"creditLimit,omitempty"`
}

// NewProjectBilling instantiates a new ProjectBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectBilling() *ProjectBilling {
	this := ProjectBilling{}
	return &this
}

// NewProjectBillingWithDefaults instantiates a new ProjectBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectBillingWithDefaults() *ProjectBilling {
	this := ProjectBilling{}
	return &this
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *ProjectBilling) GetCredits() float32 {
	if o == nil || o.Credits == nil {
		var ret float32
		return ret
	}
	return *o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBilling) GetCreditsOk() (*float32, bool) {
	if o == nil || o.Credits == nil {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *ProjectBilling) HasCredits() bool {
	if o != nil && o.Credits != nil {
		return true
	}

	return false
}

// SetCredits gets a reference to the given float32 and assigns it to the Credits field.
func (o *ProjectBilling) SetCredits(v float32) {
	o.Credits = &v
}

// GetCreditsBonus returns the CreditsBonus field value if set, zero value otherwise.
func (o *ProjectBilling) GetCreditsBonus() float32 {
	if o == nil || o.CreditsBonus == nil {
		var ret float32
		return ret
	}
	return *o.CreditsBonus
}

// GetCreditsBonusOk returns a tuple with the CreditsBonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBilling) GetCreditsBonusOk() (*float32, bool) {
	if o == nil || o.CreditsBonus == nil {
		return nil, false
	}
	return o.CreditsBonus, true
}

// HasCreditsBonus returns a boolean if a field has been set.
func (o *ProjectBilling) HasCreditsBonus() bool {
	if o != nil && o.CreditsBonus != nil {
		return true
	}

	return false
}

// SetCreditsBonus gets a reference to the given float32 and assigns it to the CreditsBonus field.
func (o *ProjectBilling) SetCreditsBonus(v float32) {
	o.CreditsBonus = &v
}

// GetCreditLimit returns the CreditLimit field value if set, zero value otherwise.
func (o *ProjectBilling) GetCreditLimit() float32 {
	if o == nil || o.CreditLimit == nil {
		var ret float32
		return ret
	}
	return *o.CreditLimit
}

// GetCreditLimitOk returns a tuple with the CreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBilling) GetCreditLimitOk() (*float32, bool) {
	if o == nil || o.CreditLimit == nil {
		return nil, false
	}
	return o.CreditLimit, true
}

// HasCreditLimit returns a boolean if a field has been set.
func (o *ProjectBilling) HasCreditLimit() bool {
	if o != nil && o.CreditLimit != nil {
		return true
	}

	return false
}

// SetCreditLimit gets a reference to the given float32 and assigns it to the CreditLimit field.
func (o *ProjectBilling) SetCreditLimit(v float32) {
	o.CreditLimit = &v
}

func (o ProjectBilling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credits != nil {
		toSerialize["credits"] = o.Credits
	}
	if o.CreditsBonus != nil {
		toSerialize["creditsBonus"] = o.CreditsBonus
	}
	if o.CreditLimit != nil {
		toSerialize["creditLimit"] = o.CreditLimit
	}
	return json.Marshal(toSerialize)
}

type NullableProjectBilling struct {
	value *ProjectBilling
	isSet bool
}

func (v NullableProjectBilling) Get() *ProjectBilling {
	return v.value
}

func (v *NullableProjectBilling) Set(val *ProjectBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectBilling(val *ProjectBilling) *NullableProjectBilling {
	return &NullableProjectBilling{value: val, isSet: true}
}

func (v NullableProjectBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


