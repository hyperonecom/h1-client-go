# DnsProjectZoneCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | **string** |  | 
**Name** | **string** |  | 
**Service** | **string** |  | 
**Source** | Pointer to [**ZoneSource**](ZoneSource.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewDnsProjectZoneCreate

`func NewDnsProjectZoneCreate(dnsName string, name string, service string, ) *DnsProjectZoneCreate`

NewDnsProjectZoneCreate instantiates a new DnsProjectZoneCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsProjectZoneCreateWithDefaults

`func NewDnsProjectZoneCreateWithDefaults() *DnsProjectZoneCreate`

NewDnsProjectZoneCreateWithDefaults instantiates a new DnsProjectZoneCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *DnsProjectZoneCreate) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *DnsProjectZoneCreate) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *DnsProjectZoneCreate) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetName

`func (o *DnsProjectZoneCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsProjectZoneCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsProjectZoneCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *DnsProjectZoneCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DnsProjectZoneCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DnsProjectZoneCreate) SetService(v string)`

SetService sets Service field to given value.


### GetSource

`func (o *DnsProjectZoneCreate) GetSource() ZoneSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DnsProjectZoneCreate) GetSourceOk() (*ZoneSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DnsProjectZoneCreate) SetSource(v ZoneSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DnsProjectZoneCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *DnsProjectZoneCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DnsProjectZoneCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DnsProjectZoneCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DnsProjectZoneCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


