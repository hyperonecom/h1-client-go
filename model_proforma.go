/*
HyperOne

HyperOne API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Proforma struct for Proforma
type Proforma struct {
	Id *string `json:"id,omitempty"`
	InvoiceNo *string `json:"invoiceNo,omitempty"`
	Seller *ProformaSeller `json:"seller,omitempty"`
	Buyer *InvoiceBuyer `json:"buyer,omitempty"`
	IssueDate *time.Time `json:"issueDate,omitempty"`
	Items []InvoiceItems `json:"items,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Project *string `json:"project,omitempty"`
	Uri *string `json:"uri,omitempty"`
	Array *ProformaArray `json:"__array__,omitempty"`
}

// NewProforma instantiates a new Proforma object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProforma() *Proforma {
	this := Proforma{}
	return &this
}

// NewProformaWithDefaults instantiates a new Proforma object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProformaWithDefaults() *Proforma {
	this := Proforma{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Proforma) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Proforma) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Proforma) SetId(v string) {
	o.Id = &v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *Proforma) GetInvoiceNo() string {
	if o == nil || o.InvoiceNo == nil {
		var ret string
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetInvoiceNoOk() (*string, bool) {
	if o == nil || o.InvoiceNo == nil {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *Proforma) HasInvoiceNo() bool {
	if o != nil && o.InvoiceNo != nil {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given string and assigns it to the InvoiceNo field.
func (o *Proforma) SetInvoiceNo(v string) {
	o.InvoiceNo = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *Proforma) GetSeller() ProformaSeller {
	if o == nil || o.Seller == nil {
		var ret ProformaSeller
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetSellerOk() (*ProformaSeller, bool) {
	if o == nil || o.Seller == nil {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *Proforma) HasSeller() bool {
	if o != nil && o.Seller != nil {
		return true
	}

	return false
}

// SetSeller gets a reference to the given ProformaSeller and assigns it to the Seller field.
func (o *Proforma) SetSeller(v ProformaSeller) {
	o.Seller = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *Proforma) GetBuyer() InvoiceBuyer {
	if o == nil || o.Buyer == nil {
		var ret InvoiceBuyer
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetBuyerOk() (*InvoiceBuyer, bool) {
	if o == nil || o.Buyer == nil {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *Proforma) HasBuyer() bool {
	if o != nil && o.Buyer != nil {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given InvoiceBuyer and assigns it to the Buyer field.
func (o *Proforma) SetBuyer(v InvoiceBuyer) {
	o.Buyer = &v
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise.
func (o *Proforma) GetIssueDate() time.Time {
	if o == nil || o.IssueDate == nil {
		var ret time.Time
		return ret
	}
	return *o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetIssueDateOk() (*time.Time, bool) {
	if o == nil || o.IssueDate == nil {
		return nil, false
	}
	return o.IssueDate, true
}

// HasIssueDate returns a boolean if a field has been set.
func (o *Proforma) HasIssueDate() bool {
	if o != nil && o.IssueDate != nil {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given time.Time and assigns it to the IssueDate field.
func (o *Proforma) SetIssueDate(v time.Time) {
	o.IssueDate = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Proforma) GetItems() []InvoiceItems {
	if o == nil || o.Items == nil {
		var ret []InvoiceItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetItemsOk() ([]InvoiceItems, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Proforma) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InvoiceItems and assigns it to the Items field.
func (o *Proforma) SetItems(v []InvoiceItems) {
	o.Items = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Proforma) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Proforma) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Proforma) SetSummary(v string) {
	o.Summary = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Proforma) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Proforma) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Proforma) SetProject(v string) {
	o.Project = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Proforma) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Proforma) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Proforma) SetUri(v string) {
	o.Uri = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *Proforma) GetArray() ProformaArray {
	if o == nil || o.Array == nil {
		var ret ProformaArray
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proforma) GetArrayOk() (*ProformaArray, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *Proforma) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given ProformaArray and assigns it to the Array field.
func (o *Proforma) SetArray(v ProformaArray) {
	o.Array = &v
}

func (o Proforma) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InvoiceNo != nil {
		toSerialize["invoiceNo"] = o.InvoiceNo
	}
	if o.Seller != nil {
		toSerialize["seller"] = o.Seller
	}
	if o.Buyer != nil {
		toSerialize["buyer"] = o.Buyer
	}
	if o.IssueDate != nil {
		toSerialize["issueDate"] = o.IssueDate
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Array != nil {
		toSerialize["__array__"] = o.Array
	}
	return json.Marshal(toSerialize)
}

type NullableProforma struct {
	value *Proforma
	isSet bool
}

func (v NullableProforma) Get() *Proforma {
	return v.value
}

func (v *NullableProforma) Set(val *Proforma) {
	v.value = val
	v.isSet = true
}

func (v NullableProforma) IsSet() bool {
	return v.isSet
}

func (v *NullableProforma) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProforma(val *Proforma) *NullableProforma {
	return &NullableProforma{value: val, isSet: true}
}

func (v NullableProforma) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProforma) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


