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

// MessageData struct for MessageData
type MessageData struct {
	Mime string `json:"mime"`
	Url *string `json:"url,omitempty"`
	Body *string `json:"body,omitempty"`
}

// NewMessageData instantiates a new MessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageData(mime string) *MessageData {
	this := MessageData{}
	this.Mime = mime
	return &this
}

// NewMessageDataWithDefaults instantiates a new MessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDataWithDefaults() *MessageData {
	this := MessageData{}
	return &this
}

// GetMime returns the Mime field value
func (o *MessageData) GetMime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mime
}

// GetMimeOk returns a tuple with the Mime field value
// and a boolean to check if the value has been set.
func (o *MessageData) GetMimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mime, true
}

// SetMime sets field value
func (o *MessageData) SetMime(v string) {
	o.Mime = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MessageData) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageData) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageData) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MessageData) SetUrl(v string) {
	o.Url = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *MessageData) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageData) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *MessageData) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *MessageData) SetBody(v string) {
	o.Body = &v
}

func (o MessageData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mime"] = o.Mime
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableMessageData struct {
	value *MessageData
	isSet bool
}

func (v NullableMessageData) Get() *MessageData {
	return v.value
}

func (v *NullableMessageData) Set(val *MessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageData(val *MessageData) *NullableMessageData {
	return &NullableMessageData{value: val, isSet: true}
}

func (v NullableMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


