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

// IamProjectPolicyCreate struct for IamProjectPolicyCreate
type IamProjectPolicyCreate struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Resource string `json:"resource"`
	Actor []IamProjectPolicyCreateActor `json:"actor,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}

// NewIamProjectPolicyCreate instantiates a new IamProjectPolicyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamProjectPolicyCreate(name string, role string, resource string) *IamProjectPolicyCreate {
	this := IamProjectPolicyCreate{}
	this.Name = name
	this.Role = role
	this.Resource = resource
	return &this
}

// NewIamProjectPolicyCreateWithDefaults instantiates a new IamProjectPolicyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamProjectPolicyCreateWithDefaults() *IamProjectPolicyCreate {
	this := IamProjectPolicyCreate{}
	return &this
}

// GetName returns the Name field value
func (o *IamProjectPolicyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IamProjectPolicyCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IamProjectPolicyCreate) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *IamProjectPolicyCreate) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *IamProjectPolicyCreate) GetRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *IamProjectPolicyCreate) SetRole(v string) {
	o.Role = v
}

// GetResource returns the Resource field value
func (o *IamProjectPolicyCreate) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *IamProjectPolicyCreate) GetResourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *IamProjectPolicyCreate) SetResource(v string) {
	o.Resource = v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *IamProjectPolicyCreate) GetActor() []IamProjectPolicyCreateActor {
	if o == nil || o.Actor == nil {
		var ret []IamProjectPolicyCreateActor
		return ret
	}
	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamProjectPolicyCreate) GetActorOk() ([]IamProjectPolicyCreateActor, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *IamProjectPolicyCreate) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given []IamProjectPolicyCreateActor and assigns it to the Actor field.
func (o *IamProjectPolicyCreate) SetActor(v []IamProjectPolicyCreateActor) {
	o.Actor = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *IamProjectPolicyCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamProjectPolicyCreate) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *IamProjectPolicyCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *IamProjectPolicyCreate) SetTag(v []Tag) {
	o.Tag = v
}

func (o IamProjectPolicyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableIamProjectPolicyCreate struct {
	value *IamProjectPolicyCreate
	isSet bool
}

func (v NullableIamProjectPolicyCreate) Get() *IamProjectPolicyCreate {
	return v.value
}

func (v *NullableIamProjectPolicyCreate) Set(val *IamProjectPolicyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableIamProjectPolicyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableIamProjectPolicyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamProjectPolicyCreate(val *IamProjectPolicyCreate) *NullableIamProjectPolicyCreate {
	return &NullableIamProjectPolicyCreate{value: val, isSet: true}
}

func (v NullableIamProjectPolicyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamProjectPolicyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


