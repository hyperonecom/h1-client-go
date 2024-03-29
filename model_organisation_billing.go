/*
HyperOne

HyperOne API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OrganisationBilling struct for OrganisationBilling
type OrganisationBilling struct {
	Address *InvoiceSellerAddress `json:"address,omitempty"`
	Company *string `json:"company,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Email *string `json:"email,omitempty"`
	Nip *string `json:"nip,omitempty"`
}

// NewOrganisationBilling instantiates a new OrganisationBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationBilling() *OrganisationBilling {
	this := OrganisationBilling{}
	return &this
}

// NewOrganisationBillingWithDefaults instantiates a new OrganisationBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationBillingWithDefaults() *OrganisationBilling {
	this := OrganisationBilling{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganisationBilling) GetAddress() InvoiceSellerAddress {
	if o == nil || o.Address == nil {
		var ret InvoiceSellerAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationBilling) GetAddressOk() (*InvoiceSellerAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganisationBilling) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given InvoiceSellerAddress and assigns it to the Address field.
func (o *OrganisationBilling) SetAddress(v InvoiceSellerAddress) {
	o.Address = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *OrganisationBilling) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationBilling) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *OrganisationBilling) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *OrganisationBilling) SetCompany(v string) {
	o.Company = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OrganisationBilling) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationBilling) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OrganisationBilling) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *OrganisationBilling) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganisationBilling) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationBilling) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganisationBilling) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganisationBilling) SetEmail(v string) {
	o.Email = &v
}

// GetNip returns the Nip field value if set, zero value otherwise.
func (o *OrganisationBilling) GetNip() string {
	if o == nil || o.Nip == nil {
		var ret string
		return ret
	}
	return *o.Nip
}

// GetNipOk returns a tuple with the Nip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationBilling) GetNipOk() (*string, bool) {
	if o == nil || o.Nip == nil {
		return nil, false
	}
	return o.Nip, true
}

// HasNip returns a boolean if a field has been set.
func (o *OrganisationBilling) HasNip() bool {
	if o != nil && o.Nip != nil {
		return true
	}

	return false
}

// SetNip gets a reference to the given string and assigns it to the Nip field.
func (o *OrganisationBilling) SetNip(v string) {
	o.Nip = &v
}

func (o OrganisationBilling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Nip != nil {
		toSerialize["nip"] = o.Nip
	}
	return json.Marshal(toSerialize)
}

type NullableOrganisationBilling struct {
	value *OrganisationBilling
	isSet bool
}

func (v NullableOrganisationBilling) Get() *OrganisationBilling {
	return v.value
}

func (v *NullableOrganisationBilling) Set(val *OrganisationBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationBilling(val *OrganisationBilling) *NullableOrganisationBilling {
	return &NullableOrganisationBilling{value: val, isSet: true}
}

func (v NullableOrganisationBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


