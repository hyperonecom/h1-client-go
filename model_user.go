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

// User struct for User
type User struct {
	Id *string `json:"id,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	FamilyName *string `json:"familyName,omitempty"`
	Name *string `json:"name,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Lang *string `json:"lang,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Limit *UserLimit `json:"limit,omitempty"`
	State *string `json:"state,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *User) GetFamilyName() string {
	if o == nil || o.FamilyName == nil {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFamilyNameOk() (*string, bool) {
	if o == nil || o.FamilyName == nil {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *User) HasFamilyName() bool {
	if o != nil && o.FamilyName != nil {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *User) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *User) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *User) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *User) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *User) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *User) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *User) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *User) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *User) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *User) SetLang(v string) {
	o.Lang = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *User) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *User) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *User) SetPhone(v string) {
	o.Phone = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *User) GetLimit() UserLimit {
	if o == nil || o.Limit == nil {
		var ret UserLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLimitOk() (*UserLimit, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *User) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given UserLimit and assigns it to the Limit field.
func (o *User) SetLimit(v UserLimit) {
	o.Limit = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *User) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *User) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *User) SetState(v string) {
	o.State = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *User) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *User) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *User) SetUri(v string) {
	o.Uri = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.FamilyName != nil {
		toSerialize["familyName"] = o.FamilyName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


