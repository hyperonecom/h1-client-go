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

// StorageS3credential struct for StorageS3credential
type StorageS3credential struct {
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	SessionToken *string `json:"sessionToken,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Region *string `json:"region,omitempty"`
	Location *string `json:"location,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewStorageS3credential instantiates a new StorageS3credential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageS3credential() *StorageS3credential {
	this := StorageS3credential{}
	return &this
}

// NewStorageS3credentialWithDefaults instantiates a new StorageS3credential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageS3credentialWithDefaults() *StorageS3credential {
	this := StorageS3credential{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *StorageS3credential) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *StorageS3credential) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *StorageS3credential) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *StorageS3credential) GetSessionToken() string {
	if o == nil || o.SessionToken == nil {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetSessionTokenOk() (*string, bool) {
	if o == nil || o.SessionToken == nil {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *StorageS3credential) HasSessionToken() bool {
	if o != nil && o.SessionToken != nil {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *StorageS3credential) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *StorageS3credential) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *StorageS3credential) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *StorageS3credential) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *StorageS3credential) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *StorageS3credential) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *StorageS3credential) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *StorageS3credential) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StorageS3credential) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StorageS3credential) SetRegion(v string) {
	o.Region = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *StorageS3credential) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *StorageS3credential) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *StorageS3credential) SetLocation(v string) {
	o.Location = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *StorageS3credential) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *StorageS3credential) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *StorageS3credential) SetBucket(v string) {
	o.Bucket = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageS3credential) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageS3credential) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageS3credential) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageS3credential) SetKey(v string) {
	o.Key = &v
}

func (o StorageS3credential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeyId != nil {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if o.SessionToken != nil {
		toSerialize["sessionToken"] = o.SessionToken
	}
	if o.SecretAccessKey != nil {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableStorageS3credential struct {
	value *StorageS3credential
	isSet bool
}

func (v NullableStorageS3credential) Get() *StorageS3credential {
	return v.value
}

func (v *NullableStorageS3credential) Set(val *StorageS3credential) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageS3credential) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageS3credential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageS3credential(val *StorageS3credential) *NullableStorageS3credential {
	return &NullableStorageS3credential{value: val, isSet: true}
}

func (v NullableStorageS3credential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageS3credential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

