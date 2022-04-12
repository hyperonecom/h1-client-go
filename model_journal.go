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

// Journal struct for Journal
type Journal struct {
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	Id string `json:"id"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Name *string `json:"name,omitempty"`
	Project *string `json:"project,omitempty"`
	Retention *float32 `json:"retention,omitempty"`
	SizeUsed *float32 `json:"sizeUsed,omitempty"`
	State *string `json:"state,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewJournal instantiates a new Journal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournal(id string) *Journal {
	this := Journal{}
	this.Id = id
	return &this
}

// NewJournalWithDefaults instantiates a new Journal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalWithDefaults() *Journal {
	this := Journal{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Journal) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Journal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Journal) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Journal) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Journal) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Journal) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *Journal) GetFlavour() string {
	if o == nil || o.Flavour == nil {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetFlavourOk() (*string, bool) {
	if o == nil || o.Flavour == nil {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *Journal) HasFlavour() bool {
	if o != nil && o.Flavour != nil {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *Journal) SetFlavour(v string) {
	o.Flavour = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *Journal) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *Journal) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *Journal) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetId returns the Id field value
func (o *Journal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Journal) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Journal) SetId(v string) {
	o.Id = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *Journal) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *Journal) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *Journal) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Journal) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Journal) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Journal) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Journal) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Journal) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Journal) SetName(v string) {
	o.Name = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Journal) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Journal) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Journal) SetProject(v string) {
	o.Project = &v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *Journal) GetRetention() float32 {
	if o == nil || o.Retention == nil {
		var ret float32
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetRetentionOk() (*float32, bool) {
	if o == nil || o.Retention == nil {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *Journal) HasRetention() bool {
	if o != nil && o.Retention != nil {
		return true
	}

	return false
}

// SetRetention gets a reference to the given float32 and assigns it to the Retention field.
func (o *Journal) SetRetention(v float32) {
	o.Retention = &v
}

// GetSizeUsed returns the SizeUsed field value if set, zero value otherwise.
func (o *Journal) GetSizeUsed() float32 {
	if o == nil || o.SizeUsed == nil {
		var ret float32
		return ret
	}
	return *o.SizeUsed
}

// GetSizeUsedOk returns a tuple with the SizeUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetSizeUsedOk() (*float32, bool) {
	if o == nil || o.SizeUsed == nil {
		return nil, false
	}
	return o.SizeUsed, true
}

// HasSizeUsed returns a boolean if a field has been set.
func (o *Journal) HasSizeUsed() bool {
	if o != nil && o.SizeUsed != nil {
		return true
	}

	return false
}

// SetSizeUsed gets a reference to the given float32 and assigns it to the SizeUsed field.
func (o *Journal) SetSizeUsed(v float32) {
	o.SizeUsed = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Journal) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Journal) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Journal) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Journal) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Journal) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *Journal) SetTag(v []Tag) {
	o.Tag = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Journal) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Journal) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Journal) SetUri(v string) {
	o.Uri = &v
}

func (o Journal) MarshalJSON() ([]byte, error) {
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
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Retention != nil {
		toSerialize["retention"] = o.Retention
	}
	if o.SizeUsed != nil {
		toSerialize["sizeUsed"] = o.SizeUsed
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

type NullableJournal struct {
	value *Journal
	isSet bool
}

func (v NullableJournal) Get() *Journal {
	return v.value
}

func (v *NullableJournal) Set(val *Journal) {
	v.value = val
	v.isSet = true
}

func (v NullableJournal) IsSet() bool {
	return v.isSet
}

func (v *NullableJournal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournal(val *Journal) *NullableJournal {
	return &NullableJournal{value: val, isSet: true}
}

func (v NullableJournal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


