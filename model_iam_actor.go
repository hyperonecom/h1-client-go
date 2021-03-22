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

// IamActor struct for IamActor
type IamActor struct {
	Id *string `json:"id,omitempty"`
	// one of: * E-mail * User URI * Service Account URI * 'me' -§ requestor
	Value string `json:"value"`
}

// NewIamActor instantiates a new IamActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamActor(value string) *IamActor {
	this := IamActor{}
	this.Value = value
	return &this
}

// NewIamActorWithDefaults instantiates a new IamActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamActorWithDefaults() *IamActor {
	this := IamActor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamActor) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamActor) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamActor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamActor) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value
func (o *IamActor) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IamActor) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IamActor) SetValue(v string) {
	o.Value = v
}

func (o IamActor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableIamActor struct {
	value *IamActor
	isSet bool
}

func (v NullableIamActor) Get() *IamActor {
	return v.value
}

func (v *NullableIamActor) Set(val *IamActor) {
	v.value = val
	v.isSet = true
}

func (v NullableIamActor) IsSet() bool {
	return v.isSet
}

func (v *NullableIamActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamActor(val *IamActor) *NullableIamActor {
	return &NullableIamActor{value: val, isSet: true}
}

func (v NullableIamActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


