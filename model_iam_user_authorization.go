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

// IamUserAuthorization struct for IamUserAuthorization
type IamUserAuthorization struct {
	Application string `json:"application"`
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Uri *string `json:"uri,omitempty"`
}

// NewIamUserAuthorization instantiates a new IamUserAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserAuthorization(application string, name string) *IamUserAuthorization {
	this := IamUserAuthorization{}
	this.Application = application
	this.Name = name
	return &this
}

// NewIamUserAuthorizationWithDefaults instantiates a new IamUserAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserAuthorizationWithDefaults() *IamUserAuthorization {
	this := IamUserAuthorization{}
	return &this
}

// GetApplication returns the Application field value
func (o *IamUserAuthorization) GetApplication() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *IamUserAuthorization) GetApplicationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *IamUserAuthorization) SetApplication(v string) {
	o.Application = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamUserAuthorization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAuthorization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamUserAuthorization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamUserAuthorization) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *IamUserAuthorization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IamUserAuthorization) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IamUserAuthorization) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *IamUserAuthorization) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserAuthorization) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *IamUserAuthorization) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *IamUserAuthorization) SetUri(v string) {
	o.Uri = &v
}

func (o IamUserAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application"] = o.Application
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableIamUserAuthorization struct {
	value *IamUserAuthorization
	isSet bool
}

func (v NullableIamUserAuthorization) Get() *IamUserAuthorization {
	return v.value
}

func (v *NullableIamUserAuthorization) Set(val *IamUserAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserAuthorization(val *IamUserAuthorization) *NullableIamUserAuthorization {
	return &NullableIamUserAuthorization{value: val, isSet: true}
}

func (v NullableIamUserAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


