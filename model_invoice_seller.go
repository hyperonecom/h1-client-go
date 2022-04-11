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

// InvoiceSeller struct for InvoiceSeller
type InvoiceSeller struct {
	Company string `json:"company"`
	Address *InvoiceSellerAddress `json:"address,omitempty"`
	Nip string `json:"nip"`
}

// NewInvoiceSeller instantiates a new InvoiceSeller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceSeller(company string, nip string) *InvoiceSeller {
	this := InvoiceSeller{}
	this.Company = company
	this.Nip = nip
	return &this
}

// NewInvoiceSellerWithDefaults instantiates a new InvoiceSeller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceSellerWithDefaults() *InvoiceSeller {
	this := InvoiceSeller{}
	return &this
}

// GetCompany returns the Company field value
func (o *InvoiceSeller) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *InvoiceSeller) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *InvoiceSeller) SetCompany(v string) {
	o.Company = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InvoiceSeller) GetAddress() InvoiceSellerAddress {
	if o == nil || o.Address == nil {
		var ret InvoiceSellerAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSeller) GetAddressOk() (*InvoiceSellerAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InvoiceSeller) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given InvoiceSellerAddress and assigns it to the Address field.
func (o *InvoiceSeller) SetAddress(v InvoiceSellerAddress) {
	o.Address = &v
}

// GetNip returns the Nip field value
func (o *InvoiceSeller) GetNip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nip
}

// GetNipOk returns a tuple with the Nip field value
// and a boolean to check if the value has been set.
func (o *InvoiceSeller) GetNipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nip, true
}

// SetNip sets field value
func (o *InvoiceSeller) SetNip(v string) {
	o.Nip = v
}

func (o InvoiceSeller) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["company"] = o.Company
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["nip"] = o.Nip
	}
	return json.Marshal(toSerialize)
}

type NullableInvoiceSeller struct {
	value *InvoiceSeller
	isSet bool
}

func (v NullableInvoiceSeller) Get() *InvoiceSeller {
	return v.value
}

func (v *NullableInvoiceSeller) Set(val *InvoiceSeller) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceSeller) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceSeller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceSeller(val *InvoiceSeller) *NullableInvoiceSeller {
	return &NullableInvoiceSeller{value: val, isSet: true}
}

func (v NullableInvoiceSeller) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceSeller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


