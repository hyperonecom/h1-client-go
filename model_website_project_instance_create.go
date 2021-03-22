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

// WebsiteProjectInstanceCreate struct for WebsiteProjectInstanceCreate
type WebsiteProjectInstanceCreate struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Image string `json:"image"`
	Source *interface{} `json:"source,omitempty"`
	Env *[]WebsiteEnv `json:"env,omitempty"`
	Tag *[]Tag `json:"tag,omitempty"`
}

// NewWebsiteProjectInstanceCreate instantiates a new WebsiteProjectInstanceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteProjectInstanceCreate(name string, service string, image string) *WebsiteProjectInstanceCreate {
	this := WebsiteProjectInstanceCreate{}
	this.Name = name
	this.Service = service
	this.Image = image
	return &this
}

// NewWebsiteProjectInstanceCreateWithDefaults instantiates a new WebsiteProjectInstanceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteProjectInstanceCreateWithDefaults() *WebsiteProjectInstanceCreate {
	this := WebsiteProjectInstanceCreate{}
	return &this
}

// GetName returns the Name field value
func (o *WebsiteProjectInstanceCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebsiteProjectInstanceCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebsiteProjectInstanceCreate) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value
func (o *WebsiteProjectInstanceCreate) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *WebsiteProjectInstanceCreate) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *WebsiteProjectInstanceCreate) SetService(v string) {
	o.Service = v
}

// GetImage returns the Image field value
func (o *WebsiteProjectInstanceCreate) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *WebsiteProjectInstanceCreate) GetImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *WebsiteProjectInstanceCreate) SetImage(v string) {
	o.Image = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *WebsiteProjectInstanceCreate) GetSource() interface{} {
	if o == nil || o.Source == nil {
		var ret interface{}
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteProjectInstanceCreate) GetSourceOk() (*interface{}, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *WebsiteProjectInstanceCreate) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given interface{} and assigns it to the Source field.
func (o *WebsiteProjectInstanceCreate) SetSource(v interface{}) {
	o.Source = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *WebsiteProjectInstanceCreate) GetEnv() []WebsiteEnv {
	if o == nil || o.Env == nil {
		var ret []WebsiteEnv
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteProjectInstanceCreate) GetEnvOk() (*[]WebsiteEnv, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *WebsiteProjectInstanceCreate) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []WebsiteEnv and assigns it to the Env field.
func (o *WebsiteProjectInstanceCreate) SetEnv(v []WebsiteEnv) {
	o.Env = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *WebsiteProjectInstanceCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteProjectInstanceCreate) GetTagOk() (*[]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *WebsiteProjectInstanceCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *WebsiteProjectInstanceCreate) SetTag(v []Tag) {
	o.Tag = &v
}

func (o WebsiteProjectInstanceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["image"] = o.Image
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableWebsiteProjectInstanceCreate struct {
	value *WebsiteProjectInstanceCreate
	isSet bool
}

func (v NullableWebsiteProjectInstanceCreate) Get() *WebsiteProjectInstanceCreate {
	return v.value
}

func (v *NullableWebsiteProjectInstanceCreate) Set(val *WebsiteProjectInstanceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteProjectInstanceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteProjectInstanceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteProjectInstanceCreate(val *WebsiteProjectInstanceCreate) *NullableWebsiteProjectInstanceCreate {
	return &NullableWebsiteProjectInstanceCreate{value: val, isSet: true}
}

func (v NullableWebsiteProjectInstanceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteProjectInstanceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


