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

// Netgw struct for Netgw
type Netgw struct {
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	Id string `json:"id"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Name string `json:"name"`
	Network *string `json:"network,omitempty"`
	PrimaryIP *string `json:"primaryIP,omitempty"`
	Project *string `json:"project,omitempty"`
	State *string `json:"state,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewNetgw instantiates a new Netgw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetgw(id string, name string) *Netgw {
	this := Netgw{}
	this.Id = id
	this.Name = name
	return &this
}

// NewNetgwWithDefaults instantiates a new Netgw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetgwWithDefaults() *Netgw {
	this := Netgw{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Netgw) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Netgw) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Netgw) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Netgw) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Netgw) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Netgw) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *Netgw) GetFlavour() string {
	if o == nil || o.Flavour == nil {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetFlavourOk() (*string, bool) {
	if o == nil || o.Flavour == nil {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *Netgw) HasFlavour() bool {
	if o != nil && o.Flavour != nil {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *Netgw) SetFlavour(v string) {
	o.Flavour = &v
}

// GetId returns the Id field value
func (o *Netgw) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Netgw) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Netgw) SetId(v string) {
	o.Id = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *Netgw) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *Netgw) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *Netgw) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Netgw) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Netgw) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Netgw) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value
func (o *Netgw) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Netgw) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Netgw) SetName(v string) {
	o.Name = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Netgw) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Netgw) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *Netgw) SetNetwork(v string) {
	o.Network = &v
}

// GetPrimaryIP returns the PrimaryIP field value if set, zero value otherwise.
func (o *Netgw) GetPrimaryIP() string {
	if o == nil || o.PrimaryIP == nil {
		var ret string
		return ret
	}
	return *o.PrimaryIP
}

// GetPrimaryIPOk returns a tuple with the PrimaryIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetPrimaryIPOk() (*string, bool) {
	if o == nil || o.PrimaryIP == nil {
		return nil, false
	}
	return o.PrimaryIP, true
}

// HasPrimaryIP returns a boolean if a field has been set.
func (o *Netgw) HasPrimaryIP() bool {
	if o != nil && o.PrimaryIP != nil {
		return true
	}

	return false
}

// SetPrimaryIP gets a reference to the given string and assigns it to the PrimaryIP field.
func (o *Netgw) SetPrimaryIP(v string) {
	o.PrimaryIP = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Netgw) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Netgw) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Netgw) SetProject(v string) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Netgw) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Netgw) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Netgw) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Netgw) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Netgw) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *Netgw) SetTag(v []Tag) {
	o.Tag = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Netgw) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Netgw) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Netgw) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Netgw) SetUri(v string) {
	o.Uri = &v
}

func (o Netgw) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Flavour != nil {
		toSerialize["flavour"] = o.Flavour
	}
	if true {
		toSerialize["id"] = o.Id
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
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.PrimaryIP != nil {
		toSerialize["primaryIP"] = o.PrimaryIP
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
	return json.Marshal(toSerialize)
}

type NullableNetgw struct {
	value *Netgw
	isSet bool
}

func (v NullableNetgw) Get() *Netgw {
	return v.value
}

func (v *NullableNetgw) Set(val *Netgw) {
	v.value = val
	v.isSet = true
}

func (v NullableNetgw) IsSet() bool {
	return v.isSet
}

func (v *NullableNetgw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetgw(val *Netgw) *NullableNetgw {
	return &NullableNetgw{value: val, isSet: true}
}

func (v NullableNetgw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetgw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


