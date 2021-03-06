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

// StorageSnapshot struct for StorageSnapshot
type StorageSnapshot struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	State *string `json:"state,omitempty"`
	Project *string `json:"project,omitempty"`
	Organisation *string `json:"organisation,omitempty"`
	Uri *string `json:"uri,omitempty"`
	SizeUsed *float32 `json:"sizeUsed,omitempty"`
}

// NewStorageSnapshot instantiates a new StorageSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSnapshot() *StorageSnapshot {
	this := StorageSnapshot{}
	return &this
}

// NewStorageSnapshotWithDefaults instantiates a new StorageSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSnapshotWithDefaults() *StorageSnapshot {
	this := StorageSnapshot{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StorageSnapshot) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StorageSnapshot) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StorageSnapshot) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageSnapshot) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageSnapshot) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageSnapshot) SetName(v string) {
	o.Name = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *StorageSnapshot) GetFlavour() string {
	if o == nil || o.Flavour == nil {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetFlavourOk() (*string, bool) {
	if o == nil || o.Flavour == nil {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *StorageSnapshot) HasFlavour() bool {
	if o != nil && o.Flavour != nil {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *StorageSnapshot) SetFlavour(v string) {
	o.Flavour = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *StorageSnapshot) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *StorageSnapshot) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *StorageSnapshot) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *StorageSnapshot) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *StorageSnapshot) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *StorageSnapshot) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *StorageSnapshot) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *StorageSnapshot) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *StorageSnapshot) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *StorageSnapshot) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *StorageSnapshot) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *StorageSnapshot) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageSnapshot) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageSnapshot) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageSnapshot) SetState(v string) {
	o.State = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *StorageSnapshot) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *StorageSnapshot) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *StorageSnapshot) SetProject(v string) {
	o.Project = &v
}

// GetOrganisation returns the Organisation field value if set, zero value otherwise.
func (o *StorageSnapshot) GetOrganisation() string {
	if o == nil || o.Organisation == nil {
		var ret string
		return ret
	}
	return *o.Organisation
}

// GetOrganisationOk returns a tuple with the Organisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetOrganisationOk() (*string, bool) {
	if o == nil || o.Organisation == nil {
		return nil, false
	}
	return o.Organisation, true
}

// HasOrganisation returns a boolean if a field has been set.
func (o *StorageSnapshot) HasOrganisation() bool {
	if o != nil && o.Organisation != nil {
		return true
	}

	return false
}

// SetOrganisation gets a reference to the given string and assigns it to the Organisation field.
func (o *StorageSnapshot) SetOrganisation(v string) {
	o.Organisation = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *StorageSnapshot) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *StorageSnapshot) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *StorageSnapshot) SetUri(v string) {
	o.Uri = &v
}

// GetSizeUsed returns the SizeUsed field value if set, zero value otherwise.
func (o *StorageSnapshot) GetSizeUsed() float32 {
	if o == nil || o.SizeUsed == nil {
		var ret float32
		return ret
	}
	return *o.SizeUsed
}

// GetSizeUsedOk returns a tuple with the SizeUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetSizeUsedOk() (*float32, bool) {
	if o == nil || o.SizeUsed == nil {
		return nil, false
	}
	return o.SizeUsed, true
}

// HasSizeUsed returns a boolean if a field has been set.
func (o *StorageSnapshot) HasSizeUsed() bool {
	if o != nil && o.SizeUsed != nil {
		return true
	}

	return false
}

// SetSizeUsed gets a reference to the given float32 and assigns it to the SizeUsed field.
func (o *StorageSnapshot) SetSizeUsed(v float32) {
	o.SizeUsed = &v
}

func (o StorageSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Flavour != nil {
		toSerialize["flavour"] = o.Flavour
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Organisation != nil {
		toSerialize["organisation"] = o.Organisation
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.SizeUsed != nil {
		toSerialize["sizeUsed"] = o.SizeUsed
	}
	return json.Marshal(toSerialize)
}

type NullableStorageSnapshot struct {
	value *StorageSnapshot
	isSet bool
}

func (v NullableStorageSnapshot) Get() *StorageSnapshot {
	return v.value
}

func (v *NullableStorageSnapshot) Set(val *StorageSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSnapshot(val *StorageSnapshot) *NullableStorageSnapshot {
	return &NullableStorageSnapshot{value: val, isSet: true}
}

func (v NullableStorageSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


