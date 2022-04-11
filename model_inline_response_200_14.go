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

// InlineResponse20014 struct for InlineResponse20014
type InlineResponse20014 struct {
	Properties OneOfobjectobject `json:"properties,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Project *string `json:"project,omitempty"`
	State *string `json:"state,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	Tag []WebsiteLocationIdProjectProjectIdInstanceTag `json:"tag,omitempty"`
	Profile string `json:"profile"`
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewInlineResponse20014 instantiates a new InlineResponse20014 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014(profile string, name string) *InlineResponse20014 {
	this := InlineResponse20014{}
	this.Profile = profile
	this.Name = name
	return &this
}

// NewInlineResponse20014WithDefaults instantiates a new InlineResponse20014 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014WithDefaults() *InlineResponse20014 {
	this := InlineResponse20014{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20014) GetProperties() OneOfobjectobject {
	if o == nil  {
		var ret OneOfobjectobject
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20014) GetPropertiesOk() (*OneOfobjectobject, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *InlineResponse20014) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given OneOfobjectobject and assigns it to the Properties field.
func (o *InlineResponse20014) SetProperties(v OneOfobjectobject) {
	o.Properties = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *InlineResponse20014) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *InlineResponse20014) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *InlineResponse20014) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *InlineResponse20014) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *InlineResponse20014) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *InlineResponse20014) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *InlineResponse20014) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *InlineResponse20014) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *InlineResponse20014) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *InlineResponse20014) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *InlineResponse20014) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *InlineResponse20014) SetProject(v string) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *InlineResponse20014) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *InlineResponse20014) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *InlineResponse20014) SetState(v string) {
	o.State = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *InlineResponse20014) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *InlineResponse20014) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *InlineResponse20014) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *InlineResponse20014) GetTag() []WebsiteLocationIdProjectProjectIdInstanceTag {
	if o == nil || o.Tag == nil {
		var ret []WebsiteLocationIdProjectProjectIdInstanceTag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetTagOk() ([]WebsiteLocationIdProjectProjectIdInstanceTag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *InlineResponse20014) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []WebsiteLocationIdProjectProjectIdInstanceTag and assigns it to the Tag field.
func (o *InlineResponse20014) SetTag(v []WebsiteLocationIdProjectProjectIdInstanceTag) {
	o.Tag = v
}

// GetProfile returns the Profile field value
func (o *InlineResponse20014) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetProfileOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *InlineResponse20014) SetProfile(v string) {
	o.Profile = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20014) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20014) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20014) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *InlineResponse20014) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse20014) SetName(v string) {
	o.Name = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *InlineResponse20014) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *InlineResponse20014) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *InlineResponse20014) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineResponse20014) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineResponse20014) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineResponse20014) SetUri(v string) {
	o.Uri = &v
}

func (o InlineResponse20014) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["profile"] = o.Profile
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014 struct {
	value *InlineResponse20014
	isSet bool
}

func (v NullableInlineResponse20014) Get() *InlineResponse20014 {
	return v.value
}

func (v *NullableInlineResponse20014) Set(val *InlineResponse20014) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014(val *InlineResponse20014) *NullableInlineResponse20014 {
	return &NullableInlineResponse20014{value: val, isSet: true}
}

func (v NullableInlineResponse20014) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


