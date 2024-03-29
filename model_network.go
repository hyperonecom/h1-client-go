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

// Network struct for Network
type Network struct {
	Address *string `json:"address,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Dns *NetworkDns `json:"dns,omitempty"`
	Firewall *string `json:"firewall,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Id string `json:"id"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Name string `json:"name"`
	Netgw *string `json:"netgw,omitempty"`
	Project *string `json:"project,omitempty"`
	State *string `json:"state,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
	Type *string `json:"type,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork(id string, name string) *Network {
	this := Network{}
	this.Id = id
	this.Name = name
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Network) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Network) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Network) SetAddress(v string) {
	o.Address = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Network) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Network) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Network) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Network) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Network) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Network) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *Network) GetDns() NetworkDns {
	if o == nil || o.Dns == nil {
		var ret NetworkDns
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetDnsOk() (*NetworkDns, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *Network) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given NetworkDns and assigns it to the Dns field.
func (o *Network) SetDns(v NetworkDns) {
	o.Dns = &v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *Network) GetFirewall() string {
	if o == nil || o.Firewall == nil {
		var ret string
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetFirewallOk() (*string, bool) {
	if o == nil || o.Firewall == nil {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *Network) HasFirewall() bool {
	if o != nil && o.Firewall != nil {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given string and assigns it to the Firewall field.
func (o *Network) SetFirewall(v string) {
	o.Firewall = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *Network) GetFlavour() string {
	if o == nil || o.Flavour == nil {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetFlavourOk() (*string, bool) {
	if o == nil || o.Flavour == nil {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *Network) HasFlavour() bool {
	if o != nil && o.Flavour != nil {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *Network) SetFlavour(v string) {
	o.Flavour = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *Network) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *Network) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *Network) SetGateway(v string) {
	o.Gateway = &v
}

// GetId returns the Id field value
func (o *Network) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Network) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Network) SetId(v string) {
	o.Id = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *Network) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *Network) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *Network) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Network) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Network) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Network) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value
func (o *Network) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Network) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Network) SetName(v string) {
	o.Name = v
}

// GetNetgw returns the Netgw field value if set, zero value otherwise.
func (o *Network) GetNetgw() string {
	if o == nil || o.Netgw == nil {
		var ret string
		return ret
	}
	return *o.Netgw
}

// GetNetgwOk returns a tuple with the Netgw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetNetgwOk() (*string, bool) {
	if o == nil || o.Netgw == nil {
		return nil, false
	}
	return o.Netgw, true
}

// HasNetgw returns a boolean if a field has been set.
func (o *Network) HasNetgw() bool {
	if o != nil && o.Netgw != nil {
		return true
	}

	return false
}

// SetNetgw gets a reference to the given string and assigns it to the Netgw field.
func (o *Network) SetNetgw(v string) {
	o.Netgw = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Network) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Network) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Network) SetProject(v string) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Network) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Network) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Network) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Network) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Network) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *Network) SetTag(v []Tag) {
	o.Tag = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Network) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Network) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Network) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Network) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Network) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Network) SetUri(v string) {
	o.Uri = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.Firewall != nil {
		toSerialize["firewall"] = o.Firewall
	}
	if o.Flavour != nil {
		toSerialize["flavour"] = o.Flavour
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
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
	if o.Netgw != nil {
		toSerialize["netgw"] = o.Netgw
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


