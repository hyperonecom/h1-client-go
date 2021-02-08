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
)

// ServiceBilling struct for ServiceBilling
type ServiceBilling struct {
	Price *ServiceBillingPrice `json:"price,omitempty"`
	Period *string `json:"period,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
	OneTime *bool `json:"oneTime,omitempty"`
	Reservations *ServiceBillingReservations `json:"reservations,omitempty"`
}

// NewServiceBilling instantiates a new ServiceBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBilling() *ServiceBilling {
	this := ServiceBilling{}
	return &this
}

// NewServiceBillingWithDefaults instantiates a new ServiceBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBillingWithDefaults() *ServiceBilling {
	this := ServiceBilling{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ServiceBilling) GetPrice() ServiceBillingPrice {
	if o == nil || o.Price == nil {
		var ret ServiceBillingPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBilling) GetPriceOk() (*ServiceBillingPrice, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ServiceBilling) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ServiceBillingPrice and assigns it to the Price field.
func (o *ServiceBilling) SetPrice(v ServiceBillingPrice) {
	o.Price = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ServiceBilling) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBilling) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ServiceBilling) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *ServiceBilling) SetPeriod(v string) {
	o.Period = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ServiceBilling) GetQuantity() float32 {
	if o == nil || o.Quantity == nil {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBilling) GetQuantityOk() (*float32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ServiceBilling) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *ServiceBilling) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetOneTime returns the OneTime field value if set, zero value otherwise.
func (o *ServiceBilling) GetOneTime() bool {
	if o == nil || o.OneTime == nil {
		var ret bool
		return ret
	}
	return *o.OneTime
}

// GetOneTimeOk returns a tuple with the OneTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBilling) GetOneTimeOk() (*bool, bool) {
	if o == nil || o.OneTime == nil {
		return nil, false
	}
	return o.OneTime, true
}

// HasOneTime returns a boolean if a field has been set.
func (o *ServiceBilling) HasOneTime() bool {
	if o != nil && o.OneTime != nil {
		return true
	}

	return false
}

// SetOneTime gets a reference to the given bool and assigns it to the OneTime field.
func (o *ServiceBilling) SetOneTime(v bool) {
	o.OneTime = &v
}

// GetReservations returns the Reservations field value if set, zero value otherwise.
func (o *ServiceBilling) GetReservations() ServiceBillingReservations {
	if o == nil || o.Reservations == nil {
		var ret ServiceBillingReservations
		return ret
	}
	return *o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBilling) GetReservationsOk() (*ServiceBillingReservations, bool) {
	if o == nil || o.Reservations == nil {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *ServiceBilling) HasReservations() bool {
	if o != nil && o.Reservations != nil {
		return true
	}

	return false
}

// SetReservations gets a reference to the given ServiceBillingReservations and assigns it to the Reservations field.
func (o *ServiceBilling) SetReservations(v ServiceBillingReservations) {
	o.Reservations = &v
}

func (o ServiceBilling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.OneTime != nil {
		toSerialize["oneTime"] = o.OneTime
	}
	if o.Reservations != nil {
		toSerialize["reservations"] = o.Reservations
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBilling struct {
	value *ServiceBilling
	isSet bool
}

func (v NullableServiceBilling) Get() *ServiceBilling {
	return v.value
}

func (v *NullableServiceBilling) Set(val *ServiceBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBilling(val *ServiceBilling) *NullableServiceBilling {
	return &NullableServiceBilling{value: val, isSet: true}
}

func (v NullableServiceBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


