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

// ServiceBillingReservations struct for ServiceBillingReservations
type ServiceBillingReservations struct {
	Id *string `json:"id,omitempty"`
	Period *string `json:"period,omitempty"`
	Price *ServiceBillingPrice `json:"price,omitempty"`
	ResourcePrice *ServiceBillingPrice `json:"resourcePrice,omitempty"`
}

// NewServiceBillingReservations instantiates a new ServiceBillingReservations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBillingReservations() *ServiceBillingReservations {
	this := ServiceBillingReservations{}
	return &this
}

// NewServiceBillingReservationsWithDefaults instantiates a new ServiceBillingReservations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBillingReservationsWithDefaults() *ServiceBillingReservations {
	this := ServiceBillingReservations{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceBillingReservations) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBillingReservations) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceBillingReservations) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceBillingReservations) SetId(v string) {
	o.Id = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ServiceBillingReservations) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBillingReservations) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ServiceBillingReservations) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *ServiceBillingReservations) SetPeriod(v string) {
	o.Period = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ServiceBillingReservations) GetPrice() ServiceBillingPrice {
	if o == nil || o.Price == nil {
		var ret ServiceBillingPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBillingReservations) GetPriceOk() (*ServiceBillingPrice, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ServiceBillingReservations) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ServiceBillingPrice and assigns it to the Price field.
func (o *ServiceBillingReservations) SetPrice(v ServiceBillingPrice) {
	o.Price = &v
}

// GetResourcePrice returns the ResourcePrice field value if set, zero value otherwise.
func (o *ServiceBillingReservations) GetResourcePrice() ServiceBillingPrice {
	if o == nil || o.ResourcePrice == nil {
		var ret ServiceBillingPrice
		return ret
	}
	return *o.ResourcePrice
}

// GetResourcePriceOk returns a tuple with the ResourcePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBillingReservations) GetResourcePriceOk() (*ServiceBillingPrice, bool) {
	if o == nil || o.ResourcePrice == nil {
		return nil, false
	}
	return o.ResourcePrice, true
}

// HasResourcePrice returns a boolean if a field has been set.
func (o *ServiceBillingReservations) HasResourcePrice() bool {
	if o != nil && o.ResourcePrice != nil {
		return true
	}

	return false
}

// SetResourcePrice gets a reference to the given ServiceBillingPrice and assigns it to the ResourcePrice field.
func (o *ServiceBillingReservations) SetResourcePrice(v ServiceBillingPrice) {
	o.ResourcePrice = &v
}

func (o ServiceBillingReservations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.ResourcePrice != nil {
		toSerialize["resourcePrice"] = o.ResourcePrice
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBillingReservations struct {
	value *ServiceBillingReservations
	isSet bool
}

func (v NullableServiceBillingReservations) Get() *ServiceBillingReservations {
	return v.value
}

func (v *NullableServiceBillingReservations) Set(val *ServiceBillingReservations) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBillingReservations) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBillingReservations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBillingReservations(val *ServiceBillingReservations) *NullableServiceBillingReservations {
	return &NullableServiceBillingReservations{value: val, isSet: true}
}

func (v NullableServiceBillingReservations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBillingReservations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


