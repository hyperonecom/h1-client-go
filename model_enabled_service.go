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

// EnabledService struct for EnabledService
type EnabledService struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Service *string `json:"service,omitempty"`
}

// NewEnabledService instantiates a new EnabledService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnabledService(id string, name string) *EnabledService {
	this := EnabledService{}
	this.Id = id
	this.Name = name
	return &this
}

// NewEnabledServiceWithDefaults instantiates a new EnabledService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnabledServiceWithDefaults() *EnabledService {
	this := EnabledService{}
	return &this
}

// GetId returns the Id field value
func (o *EnabledService) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnabledService) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnabledService) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EnabledService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnabledService) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnabledService) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *EnabledService) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledService) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *EnabledService) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *EnabledService) SetService(v string) {
	o.Service = &v
}

func (o EnabledService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableEnabledService struct {
	value *EnabledService
	isSet bool
}

func (v NullableEnabledService) Get() *EnabledService {
	return v.value
}

func (v *NullableEnabledService) Set(val *EnabledService) {
	v.value = val
	v.isSet = true
}

func (v NullableEnabledService) IsSet() bool {
	return v.isSet
}

func (v *NullableEnabledService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnabledService(val *EnabledService) *NullableEnabledService {
	return &NullableEnabledService{value: val, isSet: true}
}

func (v NullableEnabledService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnabledService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


