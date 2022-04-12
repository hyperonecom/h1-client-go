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

// IamRedirect struct for IamRedirect
type IamRedirect struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewIamRedirect instantiates a new IamRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRedirect(name string, value string) *IamRedirect {
	this := IamRedirect{}
	this.Name = name
	this.Value = value
	return &this
}

// NewIamRedirectWithDefaults instantiates a new IamRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRedirectWithDefaults() *IamRedirect {
	this := IamRedirect{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamRedirect) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRedirect) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamRedirect) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamRedirect) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *IamRedirect) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IamRedirect) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IamRedirect) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *IamRedirect) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IamRedirect) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IamRedirect) SetValue(v string) {
	o.Value = v
}

func (o IamRedirect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableIamRedirect struct {
	value *IamRedirect
	isSet bool
}

func (v NullableIamRedirect) Get() *IamRedirect {
	return v.value
}

func (v *NullableIamRedirect) Set(val *IamRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullableIamRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableIamRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamRedirect(val *IamRedirect) *NullableIamRedirect {
	return &NullableIamRedirect{value: val, isSet: true}
}

func (v NullableIamRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

