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

// ProviderProjectAgentTransfer struct for ProviderProjectAgentTransfer
type ProviderProjectAgentTransfer struct {
	Project string `json:"project"`
}

// NewProviderProjectAgentTransfer instantiates a new ProviderProjectAgentTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderProjectAgentTransfer(project string, ) *ProviderProjectAgentTransfer {
	this := ProviderProjectAgentTransfer{}
	this.Project = project
	return &this
}

// NewProviderProjectAgentTransferWithDefaults instantiates a new ProviderProjectAgentTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderProjectAgentTransferWithDefaults() *ProviderProjectAgentTransfer {
	this := ProviderProjectAgentTransfer{}
	return &this
}

// GetProject returns the Project field value
func (o *ProviderProjectAgentTransfer) GetProject() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProviderProjectAgentTransfer) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ProviderProjectAgentTransfer) SetProject(v string) {
	o.Project = v
}

func (o ProviderProjectAgentTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableProviderProjectAgentTransfer struct {
	value *ProviderProjectAgentTransfer
	isSet bool
}

func (v NullableProviderProjectAgentTransfer) Get() *ProviderProjectAgentTransfer {
	return v.value
}

func (v *NullableProviderProjectAgentTransfer) Set(val *ProviderProjectAgentTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderProjectAgentTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderProjectAgentTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderProjectAgentTransfer(val *ProviderProjectAgentTransfer) *NullableProviderProjectAgentTransfer {
	return &NullableProviderProjectAgentTransfer{value: val, isSet: true}
}

func (v NullableProviderProjectAgentTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderProjectAgentTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


