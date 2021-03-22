# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AdditionalInfo** | Pointer to **string** |  | [optional] 
**Corrections** | Pointer to **[]string** |  | [optional] 
**InvoiceInfo** | Pointer to **string** |  | [optional] 
**InvoiceNo** | Pointer to **string** |  | [optional] 
**IssueDate** | Pointer to **time.Time** |  | [optional] 
**Payments** | Pointer to **[]string** |  | [optional] 
**Seller** | Pointer to [**InvoiceSeller**](InvoiceSeller.md) |  | [optional] 
**Items** | Pointer to [**[]InvoiceItems**](InvoiceItems.md) |  | [optional] 
**Buyer** | Pointer to [**InvoiceBuyer**](InvoiceBuyer.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Duplicate** | Pointer to [**InvoiceDuplicate**](InvoiceDuplicate.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Array** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *Invoice) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *Invoice) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *Invoice) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *Invoice) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetCorrections

`func (o *Invoice) GetCorrections() []string`

GetCorrections returns the Corrections field if non-nil, zero value otherwise.

### GetCorrectionsOk

`func (o *Invoice) GetCorrectionsOk() (*[]string, bool)`

GetCorrectionsOk returns a tuple with the Corrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrections

`func (o *Invoice) SetCorrections(v []string)`

SetCorrections sets Corrections field to given value.

### HasCorrections

`func (o *Invoice) HasCorrections() bool`

HasCorrections returns a boolean if a field has been set.

### GetInvoiceInfo

`func (o *Invoice) GetInvoiceInfo() string`

GetInvoiceInfo returns the InvoiceInfo field if non-nil, zero value otherwise.

### GetInvoiceInfoOk

`func (o *Invoice) GetInvoiceInfoOk() (*string, bool)`

GetInvoiceInfoOk returns a tuple with the InvoiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceInfo

`func (o *Invoice) SetInvoiceInfo(v string)`

SetInvoiceInfo sets InvoiceInfo field to given value.

### HasInvoiceInfo

`func (o *Invoice) HasInvoiceInfo() bool`

HasInvoiceInfo returns a boolean if a field has been set.

### GetInvoiceNo

`func (o *Invoice) GetInvoiceNo() string`

GetInvoiceNo returns the InvoiceNo field if non-nil, zero value otherwise.

### GetInvoiceNoOk

`func (o *Invoice) GetInvoiceNoOk() (*string, bool)`

GetInvoiceNoOk returns a tuple with the InvoiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNo

`func (o *Invoice) SetInvoiceNo(v string)`

SetInvoiceNo sets InvoiceNo field to given value.

### HasInvoiceNo

`func (o *Invoice) HasInvoiceNo() bool`

HasInvoiceNo returns a boolean if a field has been set.

### GetIssueDate

`func (o *Invoice) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Invoice) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Invoice) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *Invoice) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPayments

`func (o *Invoice) GetPayments() []string`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Invoice) GetPaymentsOk() (*[]string, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Invoice) SetPayments(v []string)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *Invoice) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetSeller

`func (o *Invoice) GetSeller() InvoiceSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *Invoice) GetSellerOk() (*InvoiceSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *Invoice) SetSeller(v InvoiceSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *Invoice) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetItems

`func (o *Invoice) GetItems() []InvoiceItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Invoice) GetItemsOk() (*[]InvoiceItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Invoice) SetItems(v []InvoiceItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *Invoice) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBuyer

`func (o *Invoice) GetBuyer() InvoiceBuyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Invoice) GetBuyerOk() (*InvoiceBuyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Invoice) SetBuyer(v InvoiceBuyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *Invoice) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetType

`func (o *Invoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Invoice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSummary

`func (o *Invoice) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Invoice) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Invoice) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Invoice) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetProject

`func (o *Invoice) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Invoice) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Invoice) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Invoice) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetDuplicate

`func (o *Invoice) GetDuplicate() InvoiceDuplicate`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *Invoice) GetDuplicateOk() (*InvoiceDuplicate, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *Invoice) SetDuplicate(v InvoiceDuplicate)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *Invoice) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetUri

`func (o *Invoice) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Invoice) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Invoice) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Invoice) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetArray

`func (o *Invoice) GetArray() string`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *Invoice) GetArrayOk() (*string, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *Invoice) SetArray(v string)`

SetArray sets Array field to given value.

### HasArray

`func (o *Invoice) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


