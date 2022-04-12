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

// IamUserUpdate struct for IamUserUpdate
type IamUserUpdate struct {
	FamilyName *string `json:"familyName,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Lang *string `json:"lang,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// NewIamUserUpdate instantiates a new IamUserUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserUpdate() *IamUserUpdate {
	this := IamUserUpdate{}
	return &this
}

// NewIamUserUpdateWithDefaults instantiates a new IamUserUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserUpdateWithDefaults() *IamUserUpdate {
	this := IamUserUpdate{}
	return &this
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *IamUserUpdate) GetFamilyName() string {
	if o == nil || o.FamilyName == nil {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserUpdate) GetFamilyNameOk() (*string, bool) {
	if o == nil || o.FamilyName == nil {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *IamUserUpdate) HasFamilyName() bool {
	if o != nil && o.FamilyName != nil {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *IamUserUpdate) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *IamUserUpdate) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserUpdate) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *IamUserUpdate) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *IamUserUpdate) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *IamUserUpdate) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserUpdate) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *IamUserUpdate) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *IamUserUpdate) SetLang(v string) {
	o.Lang = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *IamUserUpdate) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserUpdate) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *IamUserUpdate) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *IamUserUpdate) SetPhone(v string) {
	o.Phone = &v
}

func (o IamUserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FamilyName != nil {
		toSerialize["familyName"] = o.FamilyName
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableIamUserUpdate struct {
	value *IamUserUpdate
	isSet bool
}

func (v NullableIamUserUpdate) Get() *IamUserUpdate {
	return v.value
}

func (v *NullableIamUserUpdate) Set(val *IamUserUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserUpdate(val *IamUserUpdate) *NullableIamUserUpdate {
	return &NullableIamUserUpdate{value: val, isSet: true}
}

func (v NullableIamUserUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


