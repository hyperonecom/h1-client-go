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

// NetworkingRule struct for NetworkingRule
type NetworkingRule struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Action string `json:"action"`
	Priority float32 `json:"priority"`
	Filter []string `json:"filter"`
	External *[]string `json:"external,omitempty"`
	Internal *[]string `json:"internal,omitempty"`
}

// NewNetworkingRule instantiates a new NetworkingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingRule(name string, action string, priority float32, filter []string, ) *NetworkingRule {
	this := NetworkingRule{}
	this.Name = name
	this.Action = action
	this.Priority = priority
	this.Filter = filter
	return &this
}

// NewNetworkingRuleWithDefaults instantiates a new NetworkingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingRuleWithDefaults() *NetworkingRule {
	this := NetworkingRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkingRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkingRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkingRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *NetworkingRule) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkingRule) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkingRule) SetName(v string) {
	o.Name = v
}

// GetAction returns the Action field value
func (o *NetworkingRule) GetAction() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *NetworkingRule) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *NetworkingRule) SetAction(v string) {
	o.Action = v
}

// GetPriority returns the Priority field value
func (o *NetworkingRule) GetPriority() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *NetworkingRule) GetPriorityOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *NetworkingRule) SetPriority(v float32) {
	o.Priority = v
}

// GetFilter returns the Filter field value
func (o *NetworkingRule) GetFilter() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *NetworkingRule) GetFilterOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *NetworkingRule) SetFilter(v []string) {
	o.Filter = v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *NetworkingRule) GetExternal() []string {
	if o == nil || o.External == nil {
		var ret []string
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingRule) GetExternalOk() (*[]string, bool) {
	if o == nil || o.External == nil {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *NetworkingRule) HasExternal() bool {
	if o != nil && o.External != nil {
		return true
	}

	return false
}

// SetExternal gets a reference to the given []string and assigns it to the External field.
func (o *NetworkingRule) SetExternal(v []string) {
	o.External = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *NetworkingRule) GetInternal() []string {
	if o == nil || o.Internal == nil {
		var ret []string
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingRule) GetInternalOk() (*[]string, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *NetworkingRule) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given []string and assigns it to the Internal field.
func (o *NetworkingRule) SetInternal(v []string) {
	o.Internal = &v
}

func (o NetworkingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if o.External != nil {
		toSerialize["external"] = o.External
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingRule struct {
	value *NetworkingRule
	isSet bool
}

func (v NullableNetworkingRule) Get() *NetworkingRule {
	return v.value
}

func (v *NullableNetworkingRule) Set(val *NetworkingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingRule(val *NetworkingRule) *NullableNetworkingRule {
	return &NullableNetworkingRule{value: val, isSet: true}
}

func (v NullableNetworkingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


