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

// DnsRecordset struct for DnsRecordset
type DnsRecordset struct {
	Id *string `json:"id,omitempty"`
	// use '@' to reference the zone origin
	Name *string `json:"name,omitempty"`
	Type string `json:"type"`
	Ttl *float32 `json:"ttl,omitempty"`
	Record *[]DnsRecord `json:"record,omitempty"`
}

// NewDnsRecordset instantiates a new DnsRecordset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRecordset(type_ string) *DnsRecordset {
	this := DnsRecordset{}
	var name string = "@"
	this.Name = &name
	this.Type = type_
	var ttl float32 = 3600
	this.Ttl = &ttl
	return &this
}

// NewDnsRecordsetWithDefaults instantiates a new DnsRecordset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRecordsetWithDefaults() *DnsRecordset {
	this := DnsRecordset{}
	var name string = "@"
	this.Name = &name
	var ttl float32 = 3600
	this.Ttl = &ttl
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DnsRecordset) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordset) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DnsRecordset) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DnsRecordset) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnsRecordset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnsRecordset) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnsRecordset) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *DnsRecordset) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DnsRecordset) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DnsRecordset) SetType(v string) {
	o.Type = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *DnsRecordset) GetTtl() float32 {
	if o == nil || o.Ttl == nil {
		var ret float32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordset) GetTtlOk() (*float32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *DnsRecordset) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given float32 and assigns it to the Ttl field.
func (o *DnsRecordset) SetTtl(v float32) {
	o.Ttl = &v
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *DnsRecordset) GetRecord() []DnsRecord {
	if o == nil || o.Record == nil {
		var ret []DnsRecord
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordset) GetRecordOk() (*[]DnsRecord, bool) {
	if o == nil || o.Record == nil {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *DnsRecordset) HasRecord() bool {
	if o != nil && o.Record != nil {
		return true
	}

	return false
}

// SetRecord gets a reference to the given []DnsRecord and assigns it to the Record field.
func (o *DnsRecordset) SetRecord(v []DnsRecord) {
	o.Record = &v
}

func (o DnsRecordset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.Record != nil {
		toSerialize["record"] = o.Record
	}
	return json.Marshal(toSerialize)
}

type NullableDnsRecordset struct {
	value *DnsRecordset
	isSet bool
}

func (v NullableDnsRecordset) Get() *DnsRecordset {
	return v.value
}

func (v *NullableDnsRecordset) Set(val *DnsRecordset) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecordset) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecordset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecordset(val *DnsRecordset) *NullableDnsRecordset {
	return &NullableDnsRecordset{value: val, isSet: true}
}

func (v NullableDnsRecordset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecordset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


