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
	"time"
)

// DevelopmentCredential struct for DevelopmentCredential
type DevelopmentCredential struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Type string `json:"type"`
	Value string `json:"value"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewDevelopmentCredential instantiates a new DevelopmentCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevelopmentCredential(name string, type_ string, value string) *DevelopmentCredential {
	this := DevelopmentCredential{}
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewDevelopmentCredentialWithDefaults instantiates a new DevelopmentCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevelopmentCredentialWithDefaults() *DevelopmentCredential {
	this := DevelopmentCredential{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevelopmentCredential) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevelopmentCredential) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevelopmentCredential) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DevelopmentCredential) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DevelopmentCredential) SetName(v string) {
	o.Name = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DevelopmentCredential) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DevelopmentCredential) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DevelopmentCredential) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *DevelopmentCredential) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *DevelopmentCredential) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *DevelopmentCredential) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetType returns the Type field value
func (o *DevelopmentCredential) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DevelopmentCredential) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *DevelopmentCredential) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DevelopmentCredential) SetValue(v string) {
	o.Value = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *DevelopmentCredential) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *DevelopmentCredential) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *DevelopmentCredential) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DevelopmentCredential) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentCredential) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DevelopmentCredential) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DevelopmentCredential) SetToken(v string) {
	o.Token = &v
}

func (o DevelopmentCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableDevelopmentCredential struct {
	value *DevelopmentCredential
	isSet bool
}

func (v NullableDevelopmentCredential) Get() *DevelopmentCredential {
	return v.value
}

func (v *NullableDevelopmentCredential) Set(val *DevelopmentCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableDevelopmentCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableDevelopmentCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevelopmentCredential(val *DevelopmentCredential) *NullableDevelopmentCredential {
	return &NullableDevelopmentCredential{value: val, isSet: true}
}

func (v NullableDevelopmentCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevelopmentCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

