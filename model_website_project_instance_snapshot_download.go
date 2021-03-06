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

// WebsiteProjectInstanceSnapshotDownload struct for WebsiteProjectInstanceSnapshotDownload
type WebsiteProjectInstanceSnapshotDownload struct {
	Incremental *string `json:"incremental,omitempty"`
}

// NewWebsiteProjectInstanceSnapshotDownload instantiates a new WebsiteProjectInstanceSnapshotDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteProjectInstanceSnapshotDownload() *WebsiteProjectInstanceSnapshotDownload {
	this := WebsiteProjectInstanceSnapshotDownload{}
	return &this
}

// NewWebsiteProjectInstanceSnapshotDownloadWithDefaults instantiates a new WebsiteProjectInstanceSnapshotDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteProjectInstanceSnapshotDownloadWithDefaults() *WebsiteProjectInstanceSnapshotDownload {
	this := WebsiteProjectInstanceSnapshotDownload{}
	return &this
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *WebsiteProjectInstanceSnapshotDownload) GetIncremental() string {
	if o == nil || o.Incremental == nil {
		var ret string
		return ret
	}
	return *o.Incremental
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteProjectInstanceSnapshotDownload) GetIncrementalOk() (*string, bool) {
	if o == nil || o.Incremental == nil {
		return nil, false
	}
	return o.Incremental, true
}

// HasIncremental returns a boolean if a field has been set.
func (o *WebsiteProjectInstanceSnapshotDownload) HasIncremental() bool {
	if o != nil && o.Incremental != nil {
		return true
	}

	return false
}

// SetIncremental gets a reference to the given string and assigns it to the Incremental field.
func (o *WebsiteProjectInstanceSnapshotDownload) SetIncremental(v string) {
	o.Incremental = &v
}

func (o WebsiteProjectInstanceSnapshotDownload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Incremental != nil {
		toSerialize["incremental"] = o.Incremental
	}
	return json.Marshal(toSerialize)
}

type NullableWebsiteProjectInstanceSnapshotDownload struct {
	value *WebsiteProjectInstanceSnapshotDownload
	isSet bool
}

func (v NullableWebsiteProjectInstanceSnapshotDownload) Get() *WebsiteProjectInstanceSnapshotDownload {
	return v.value
}

func (v *NullableWebsiteProjectInstanceSnapshotDownload) Set(val *WebsiteProjectInstanceSnapshotDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteProjectInstanceSnapshotDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteProjectInstanceSnapshotDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteProjectInstanceSnapshotDownload(val *WebsiteProjectInstanceSnapshotDownload) *NullableWebsiteProjectInstanceSnapshotDownload {
	return &NullableWebsiteProjectInstanceSnapshotDownload{value: val, isSet: true}
}

func (v NullableWebsiteProjectInstanceSnapshotDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteProjectInstanceSnapshotDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


