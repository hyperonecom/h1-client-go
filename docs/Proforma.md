# Proforma

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Array** | Pointer to [**ProformaArray**](ProformaArray.md) |  | [optional] 
**Buyer** | Pointer to [**InvoiceBuyer**](InvoiceBuyer.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceNo** | Pointer to **string** |  | [optional] 
**IssueDate** | Pointer to **time.Time** |  | [optional] 
**Items** | Pointer to [**[]InvoiceItems**](InvoiceItems.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Seller** | Pointer to [**ProformaSeller**](ProformaSeller.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewProforma

`func NewProforma() *Proforma`

NewProforma instantiates a new Proforma object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProformaWithDefaults

`func NewProformaWithDefaults() *Proforma`

NewProformaWithDefaults instantiates a new Proforma object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArray

`func (o *Proforma) GetArray() ProformaArray`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *Proforma) GetArrayOk() (*ProformaArray, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *Proforma) SetArray(v ProformaArray)`

SetArray sets Array field to given value.

### HasArray

`func (o *Proforma) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetBuyer

`func (o *Proforma) GetBuyer() InvoiceBuyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Proforma) GetBuyerOk() (*InvoiceBuyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Proforma) SetBuyer(v InvoiceBuyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *Proforma) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetId

`func (o *Proforma) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Proforma) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Proforma) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Proforma) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceNo

`func (o *Proforma) GetInvoiceNo() string`

GetInvoiceNo returns the InvoiceNo field if non-nil, zero value otherwise.

### GetInvoiceNoOk

`func (o *Proforma) GetInvoiceNoOk() (*string, bool)`

GetInvoiceNoOk returns a tuple with the InvoiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNo

`func (o *Proforma) SetInvoiceNo(v string)`

SetInvoiceNo sets InvoiceNo field to given value.

### HasInvoiceNo

`func (o *Proforma) HasInvoiceNo() bool`

HasInvoiceNo returns a boolean if a field has been set.

### GetIssueDate

`func (o *Proforma) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Proforma) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Proforma) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *Proforma) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetItems

`func (o *Proforma) GetItems() []InvoiceItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Proforma) GetItemsOk() (*[]InvoiceItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Proforma) SetItems(v []InvoiceItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *Proforma) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetProject

`func (o *Proforma) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Proforma) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Proforma) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Proforma) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSeller

`func (o *Proforma) GetSeller() ProformaSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *Proforma) GetSellerOk() (*ProformaSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *Proforma) SetSeller(v ProformaSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *Proforma) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetSummary

`func (o *Proforma) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Proforma) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Proforma) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Proforma) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUri

`func (o *Proforma) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Proforma) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Proforma) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Proforma) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


