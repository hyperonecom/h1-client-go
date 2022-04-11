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

// RecoveryProjectBackupExport struct for RecoveryProjectBackupExport
type RecoveryProjectBackupExport struct {
	Bucket string `json:"bucket"`
}

// NewRecoveryProjectBackupExport instantiates a new RecoveryProjectBackupExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryProjectBackupExport(bucket string) *RecoveryProjectBackupExport {
	this := RecoveryProjectBackupExport{}
	this.Bucket = bucket
	return &this
}

// NewRecoveryProjectBackupExportWithDefaults instantiates a new RecoveryProjectBackupExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryProjectBackupExportWithDefaults() *RecoveryProjectBackupExport {
	this := RecoveryProjectBackupExport{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *RecoveryProjectBackupExport) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *RecoveryProjectBackupExport) GetBucketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *RecoveryProjectBackupExport) SetBucket(v string) {
	o.Bucket = v
}

func (o RecoveryProjectBackupExport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryProjectBackupExport struct {
	value *RecoveryProjectBackupExport
	isSet bool
}

func (v NullableRecoveryProjectBackupExport) Get() *RecoveryProjectBackupExport {
	return v.value
}

func (v *NullableRecoveryProjectBackupExport) Set(val *RecoveryProjectBackupExport) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryProjectBackupExport) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryProjectBackupExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryProjectBackupExport(val *RecoveryProjectBackupExport) *NullableRecoveryProjectBackupExport {
	return &NullableRecoveryProjectBackupExport{value: val, isSet: true}
}

func (v NullableRecoveryProjectBackupExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryProjectBackupExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


