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

// InvoiceItems struct for InvoiceItems
type InvoiceItems struct {
	Brutto string `json:"brutto"`
	Name string `json:"name"`
	Netto string `json:"netto"`
	Price string `json:"price"`
	Quantity float32 `json:"quantity"`
	VatAmount string `json:"vatAmount"`
	VatRate string `json:"vatRate"`
}

// NewInvoiceItems instantiates a new InvoiceItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceItems(brutto string, name string, netto string, price string, quantity float32, vatAmount string, vatRate string) *InvoiceItems {
	this := InvoiceItems{}
	this.Brutto = brutto
	this.Name = name
	this.Netto = netto
	this.Price = price
	this.Quantity = quantity
	this.VatAmount = vatAmount
	this.VatRate = vatRate
	return &this
}

// NewInvoiceItemsWithDefaults instantiates a new InvoiceItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceItemsWithDefaults() *InvoiceItems {
	this := InvoiceItems{}
	return &this
}

// GetBrutto returns the Brutto field value
func (o *InvoiceItems) GetBrutto() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brutto
}

// GetBruttoOk returns a tuple with the Brutto field value
// and a boolean to check if the value has been set.
func (o *InvoiceItems) GetBruttoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Brutto, true
}

// SetBrutto sets field value
func (o *InvoiceItems) SetBrutto(v string) {
	o.Brutto = v
}

// GetName returns the Name field value
func (o *InvoiceItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvoiceItems) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvoiceItems) SetName(v string) {
	o.Name = v
}

// GetNetto returns the Netto field value
func (o *InvoiceItems) GetNetto() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Netto
}

// GetNettoOk returns a tuple with the Netto field value
// and a boolean to check if the value has been set.
func (o *InvoiceItems) GetNettoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Netto, true
}

// SetNetto sets field value
func (o *InvoiceItems) SetNetto(v string) {
	o.Netto = v
}

// GetPrice returns the Price field value
func (o *InvoiceItems) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InvoiceItems) GetPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InvoiceItems) SetPrice(v string) {
	o.Price = v
}

// GetQuantity returns the Quantity field value
func (o *InvoiceItems) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InvoiceItems) GetQuantityOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InvoiceItems) SetQuantity(v float32) {
	o.Quantity = v
}

// GetVatAmount returns the VatAmount field value
func (o *InvoiceItems) GetVatAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VatAmount
}

// GetVatAmountOk returns a tuple with the VatAmount field value
// and a boolean to check if the value has been set.
func (o *InvoiceItems) GetVatAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VatAmount, true
}

// SetVatAmount sets field value
func (o *InvoiceItems) SetVatAmount(v string) {
	o.VatAmount = v
}

// GetVatRate returns the VatRate field value
func (o *InvoiceItems) GetVatRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VatRate
}

// GetVatRateOk returns a tuple with the VatRate field value
// and a boolean to check if the value has been set.
func (o *InvoiceItems) GetVatRateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VatRate, true
}

// SetVatRate sets field value
func (o *InvoiceItems) SetVatRate(v string) {
	o.VatRate = v
}

func (o InvoiceItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brutto"] = o.Brutto
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["netto"] = o.Netto
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["vatAmount"] = o.VatAmount
	}
	if true {
		toSerialize["vatRate"] = o.VatRate
	}
	return json.Marshal(toSerialize)
}

type NullableInvoiceItems struct {
	value *InvoiceItems
	isSet bool
}

func (v NullableInvoiceItems) Get() *InvoiceItems {
	return v.value
}

func (v *NullableInvoiceItems) Set(val *InvoiceItems) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceItems) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceItems(val *InvoiceItems) *NullableInvoiceItems {
	return &NullableInvoiceItems{value: val, isSet: true}
}

func (v NullableInvoiceItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


