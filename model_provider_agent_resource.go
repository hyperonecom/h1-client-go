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

// ProviderAgentResource struct for ProviderAgentResource
type ProviderAgentResource struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Project *string `json:"project,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

// NewProviderAgentResource instantiates a new ProviderAgentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderAgentResource() *ProviderAgentResource {
	this := ProviderAgentResource{}
	return &this
}

// NewProviderAgentResourceWithDefaults instantiates a new ProviderAgentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderAgentResourceWithDefaults() *ProviderAgentResource {
	this := ProviderAgentResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProviderAgentResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderAgentResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProviderAgentResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProviderAgentResource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProviderAgentResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderAgentResource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProviderAgentResource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProviderAgentResource) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProviderAgentResource) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderAgentResource) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProviderAgentResource) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ProviderAgentResource) SetState(v string) {
	o.State = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ProviderAgentResource) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderAgentResource) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ProviderAgentResource) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *ProviderAgentResource) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProviderAgentResource) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderAgentResource) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProviderAgentResource) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *ProviderAgentResource) SetProject(v string) {
	o.Project = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ProviderAgentResource) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderAgentResource) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ProviderAgentResource) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *ProviderAgentResource) SetResource(v string) {
	o.Resource = &v
}

func (o ProviderAgentResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableProviderAgentResource struct {
	value *ProviderAgentResource
	isSet bool
}

func (v NullableProviderAgentResource) Get() *ProviderAgentResource {
	return v.value
}

func (v *NullableProviderAgentResource) Set(val *ProviderAgentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderAgentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderAgentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderAgentResource(val *ProviderAgentResource) *NullableProviderAgentResource {
	return &NullableProviderAgentResource{value: val, isSet: true}
}

func (v NullableProviderAgentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderAgentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


