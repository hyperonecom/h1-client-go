/*
HyperOne

HyperOne API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Vm struct for Vm
type Vm struct {
	Cpu *float32 `json:"cpu,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	Id string `json:"id"`
	Memory *float32 `json:"memory,omitempty"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Name string `json:"name"`
	Project *string `json:"project,omitempty"`
	State *string `json:"state,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
	Uri *string `json:"uri,omitempty"`
	UserMetadata *string `json:"userMetadata,omitempty"`
}

// NewVm instantiates a new Vm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVm(id string, name string) *Vm {
	this := Vm{}
	this.Id = id
	this.Name = name
	return &this
}

// NewVmWithDefaults instantiates a new Vm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmWithDefaults() *Vm {
	this := Vm{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Vm) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Vm) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *Vm) SetCpu(v float32) {
	o.Cpu = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Vm) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Vm) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Vm) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Vm) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Vm) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Vm) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *Vm) GetFlavour() string {
	if o == nil || o.Flavour == nil {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetFlavourOk() (*string, bool) {
	if o == nil || o.Flavour == nil {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *Vm) HasFlavour() bool {
	if o != nil && o.Flavour != nil {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *Vm) SetFlavour(v string) {
	o.Flavour = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *Vm) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *Vm) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *Vm) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetId returns the Id field value
func (o *Vm) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Vm) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Vm) SetId(v string) {
	o.Id = v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Vm) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Vm) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *Vm) SetMemory(v float32) {
	o.Memory = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *Vm) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *Vm) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *Vm) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Vm) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Vm) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Vm) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value
func (o *Vm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Vm) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Vm) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Vm) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Vm) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Vm) SetProject(v string) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Vm) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Vm) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Vm) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Vm) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Vm) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *Vm) SetTag(v []Tag) {
	o.Tag = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Vm) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Vm) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Vm) SetUri(v string) {
	o.Uri = &v
}

// GetUserMetadata returns the UserMetadata field value if set, zero value otherwise.
func (o *Vm) GetUserMetadata() string {
	if o == nil || o.UserMetadata == nil {
		var ret string
		return ret
	}
	return *o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vm) GetUserMetadataOk() (*string, bool) {
	if o == nil || o.UserMetadata == nil {
		return nil, false
	}
	return o.UserMetadata, true
}

// HasUserMetadata returns a boolean if a field has been set.
func (o *Vm) HasUserMetadata() bool {
	if o != nil && o.UserMetadata != nil {
		return true
	}

	return false
}

// SetUserMetadata gets a reference to the given string and assigns it to the UserMetadata field.
func (o *Vm) SetUserMetadata(v string) {
	o.UserMetadata = &v
}

func (o Vm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Flavour != nil {
		toSerialize["flavour"] = o.Flavour
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.UserMetadata != nil {
		toSerialize["userMetadata"] = o.UserMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableVm struct {
	value *Vm
	isSet bool
}

func (v NullableVm) Get() *Vm {
	return v.value
}

func (v *NullableVm) Set(val *Vm) {
	v.value = val
	v.isSet = true
}

func (v NullableVm) IsSet() bool {
	return v.isSet
}

func (v *NullableVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVm(val *Vm) *NullableVm {
	return &NullableVm{value: val, isSet: true}
}

func (v NullableVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


