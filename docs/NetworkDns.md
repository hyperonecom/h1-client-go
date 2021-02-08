# NetworkDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nameservers** | Pointer to **[]string** |  | [optional] [default to ["9.9.9.9","8.8.8.8"]]

## Methods

### NewNetworkDns

`func NewNetworkDns() *NetworkDns`

NewNetworkDns instantiates a new NetworkDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDnsWithDefaults

`func NewNetworkDnsWithDefaults() *NetworkDns`

NewNetworkDnsWithDefaults instantiates a new NetworkDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameservers

`func (o *NetworkDns) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *NetworkDns) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *NetworkDns) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *NetworkDns) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


