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

// Zone struct for Zone
type Zone struct {
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	DnsName *string `json:"dnsName,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Name *string `json:"name,omitempty"`
	Nameserver []string `json:"nameserver,omitempty"`
	Project *string `json:"project,omitempty"`
	State *string `json:"state,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewZone instantiates a new Zone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZone() *Zone {
	this := Zone{}
	return &this
}

// NewZoneWithDefaults instantiates a new Zone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneWithDefaults() *Zone {
	this := Zone{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Zone) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Zone) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Zone) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Zone) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Zone) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Zone) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *Zone) GetDnsName() string {
	if o == nil || o.DnsName == nil {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetDnsNameOk() (*string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *Zone) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *Zone) SetDnsName(v string) {
	o.DnsName = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *Zone) GetFlavour() string {
	if o == nil || o.Flavour == nil {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetFlavourOk() (*string, bool) {
	if o == nil || o.Flavour == nil {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *Zone) HasFlavour() bool {
	if o != nil && o.Flavour != nil {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *Zone) SetFlavour(v string) {
	o.Flavour = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *Zone) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *Zone) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *Zone) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Zone) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Zone) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Zone) SetId(v string) {
	o.Id = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *Zone) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *Zone) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *Zone) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Zone) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Zone) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Zone) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Zone) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Zone) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Zone) SetName(v string) {
	o.Name = &v
}

// GetNameserver returns the Nameserver field value if set, zero value otherwise.
func (o *Zone) GetNameserver() []string {
	if o == nil || o.Nameserver == nil {
		var ret []string
		return ret
	}
	return o.Nameserver
}

// GetNameserverOk returns a tuple with the Nameserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNameserverOk() ([]string, bool) {
	if o == nil || o.Nameserver == nil {
		return nil, false
	}
	return o.Nameserver, true
}

// HasNameserver returns a boolean if a field has been set.
func (o *Zone) HasNameserver() bool {
	if o != nil && o.Nameserver != nil {
		return true
	}

	return false
}

// SetNameserver gets a reference to the given []string and assigns it to the Nameserver field.
func (o *Zone) SetNameserver(v []string) {
	o.Nameserver = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Zone) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Zone) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Zone) SetProject(v string) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Zone) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Zone) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Zone) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Zone) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Zone) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *Zone) SetTag(v []Tag) {
	o.Tag = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Zone) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Zone) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Zone) SetUri(v string) {
	o.Uri = &v
}

func (o Zone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.DnsName != nil {
		toSerialize["dnsName"] = o.DnsName
	}
	if o.Flavour != nil {
		toSerialize["flavour"] = o.Flavour
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Nameserver != nil {
		toSerialize["nameserver"] = o.Nameserver
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

type NullableZone struct {
	value *Zone
	isSet bool
}

func (v NullableZone) Get() *Zone {
	return v.value
}

func (v *NullableZone) Set(val *Zone) {
	v.value = val
	v.isSet = true
}

func (v NullableZone) IsSet() bool {
	return v.isSet
}

func (v *NullableZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZone(val *Zone) *NullableZone {
	return &NullableZone{value: val, isSet: true}
}

func (v NullableZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


