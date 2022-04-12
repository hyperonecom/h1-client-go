# DnsRecordset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | use &#39;@&#39; to reference the zone origin | [optional] [default to "@"]
**Record** | Pointer to [**[]DnsRecord**](DnsRecord.md) |  | [optional] 
**Ttl** | Pointer to **float32** |  | [optional] [default to 3600]
**Type** | **string** |  | 

## Methods

### NewDnsRecordset

`func NewDnsRecordset(type_ string, ) *DnsRecordset`

NewDnsRecordset instantiates a new DnsRecordset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordsetWithDefaults

`func NewDnsRecordsetWithDefaults() *DnsRecordset`

NewDnsRecordsetWithDefaults instantiates a new DnsRecordset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsRecordset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecordset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecordset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnsRecordset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DnsRecordset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecordset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecordset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsRecordset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecord

`func (o *DnsRecordset) GetRecord() []DnsRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DnsRecordset) GetRecordOk() (*[]DnsRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DnsRecordset) SetRecord(v []DnsRecord)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *DnsRecordset) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTtl

`func (o *DnsRecordset) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecordset) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecordset) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecordset) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *DnsRecordset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecordset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecordset) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


