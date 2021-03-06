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

// Reservation struct for Reservation
type Reservation struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	State *string `json:"state,omitempty"`
	Project *string `json:"project,omitempty"`
	Uri *string `json:"uri,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Assigned *string `json:"assigned,omitempty"`
	SourceService *string `json:"sourceService,omitempty"`
	Tag *[]Tag `json:"tag,omitempty"`
}

// NewReservation instantiates a new Reservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservation() *Reservation {
	this := Reservation{}
	return &this
}

// NewReservationWithDefaults instantiates a new Reservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationWithDefaults() *Reservation {
	this := Reservation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Reservation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Reservation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Reservation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Reservation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Reservation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Reservation) SetName(v string) {
	o.Name = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *Reservation) GetFlavour() string {
	if o == nil || o.Flavour == nil {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetFlavourOk() (*string, bool) {
	if o == nil || o.Flavour == nil {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *Reservation) HasFlavour() bool {
	if o != nil && o.Flavour != nil {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *Reservation) SetFlavour(v string) {
	o.Flavour = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Reservation) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Reservation) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Reservation) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *Reservation) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *Reservation) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *Reservation) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Reservation) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Reservation) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Reservation) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Reservation) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Reservation) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Reservation) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Reservation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Reservation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Reservation) SetState(v string) {
	o.State = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Reservation) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Reservation) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Reservation) SetProject(v string) {
	o.Project = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Reservation) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Reservation) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Reservation) SetUri(v string) {
	o.Uri = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Reservation) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Reservation) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Reservation) SetKind(v string) {
	o.Kind = &v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *Reservation) GetAssigned() string {
	if o == nil || o.Assigned == nil {
		var ret string
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetAssignedOk() (*string, bool) {
	if o == nil || o.Assigned == nil {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *Reservation) HasAssigned() bool {
	if o != nil && o.Assigned != nil {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given string and assigns it to the Assigned field.
func (o *Reservation) SetAssigned(v string) {
	o.Assigned = &v
}

// GetSourceService returns the SourceService field value if set, zero value otherwise.
func (o *Reservation) GetSourceService() string {
	if o == nil || o.SourceService == nil {
		var ret string
		return ret
	}
	return *o.SourceService
}

// GetSourceServiceOk returns a tuple with the SourceService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetSourceServiceOk() (*string, bool) {
	if o == nil || o.SourceService == nil {
		return nil, false
	}
	return o.SourceService, true
}

// HasSourceService returns a boolean if a field has been set.
func (o *Reservation) HasSourceService() bool {
	if o != nil && o.SourceService != nil {
		return true
	}

	return false
}

// SetSourceService gets a reference to the given string and assigns it to the SourceService field.
func (o *Reservation) SetSourceService(v string) {
	o.SourceService = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Reservation) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetTagOk() (*[]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Reservation) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *Reservation) SetTag(v []Tag) {
	o.Tag = &v
}

func (o Reservation) MarshalJSON() ([]byte, error) {
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
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Assigned != nil {
		toSerialize["assigned"] = o.Assigned
	}
	if o.SourceService != nil {
		toSerialize["sourceService"] = o.SourceService
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableReservation struct {
	value *Reservation
	isSet bool
}

func (v NullableReservation) Get() *Reservation {
	return v.value
}

func (v *NullableReservation) Set(val *Reservation) {
	v.value = val
	v.isSet = true
}

func (v NullableReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservation(val *Reservation) *NullableReservation {
	return &NullableReservation{value: val, isSet: true}
}

func (v NullableReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


