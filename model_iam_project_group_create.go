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

// IamProjectGroupCreate struct for IamProjectGroupCreate
type IamProjectGroupCreate struct {
	Actor []IamActor `json:"actor,omitempty"`
	Name string `json:"name"`
	Tag []Tag `json:"tag,omitempty"`
}

// NewIamProjectGroupCreate instantiates a new IamProjectGroupCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamProjectGroupCreate(name string) *IamProjectGroupCreate {
	this := IamProjectGroupCreate{}
	this.Name = name
	return &this
}

// NewIamProjectGroupCreateWithDefaults instantiates a new IamProjectGroupCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamProjectGroupCreateWithDefaults() *IamProjectGroupCreate {
	this := IamProjectGroupCreate{}
	return &this
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *IamProjectGroupCreate) GetActor() []IamActor {
	if o == nil || o.Actor == nil {
		var ret []IamActor
		return ret
	}
	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamProjectGroupCreate) GetActorOk() ([]IamActor, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *IamProjectGroupCreate) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given []IamActor and assigns it to the Actor field.
func (o *IamProjectGroupCreate) SetActor(v []IamActor) {
	o.Actor = v
}

// GetName returns the Name field value
func (o *IamProjectGroupCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IamProjectGroupCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IamProjectGroupCreate) SetName(v string) {
	o.Name = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *IamProjectGroupCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamProjectGroupCreate) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *IamProjectGroupCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *IamProjectGroupCreate) SetTag(v []Tag) {
	o.Tag = v
}

func (o IamProjectGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableIamProjectGroupCreate struct {
	value *IamProjectGroupCreate
	isSet bool
}

func (v NullableIamProjectGroupCreate) Get() *IamProjectGroupCreate {
	return v.value
}

func (v *NullableIamProjectGroupCreate) Set(val *IamProjectGroupCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableIamProjectGroupCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableIamProjectGroupCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamProjectGroupCreate(val *IamProjectGroupCreate) *NullableIamProjectGroupCreate {
	return &NullableIamProjectGroupCreate{value: val, isSet: true}
}

func (v NullableIamProjectGroupCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamProjectGroupCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


