# NetworkingProjectNetadpCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vm** | **string** |  | 
**Network** | **string** |  | 
**Firewall** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **[]string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewNetworkingProjectNetadpCreate

`func NewNetworkingProjectNetadpCreate(vm string, network string, ) *NetworkingProjectNetadpCreate`

NewNetworkingProjectNetadpCreate instantiates a new NetworkingProjectNetadpCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingProjectNetadpCreateWithDefaults

`func NewNetworkingProjectNetadpCreateWithDefaults() *NetworkingProjectNetadpCreate`

NewNetworkingProjectNetadpCreateWithDefaults instantiates a new NetworkingProjectNetadpCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVm

`func (o *NetworkingProjectNetadpCreate) GetVm() string`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *NetworkingProjectNetadpCreate) GetVmOk() (*string, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *NetworkingProjectNetadpCreate) SetVm(v string)`

SetVm sets Vm field to given value.


### GetNetwork

`func (o *NetworkingProjectNetadpCreate) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingProjectNetadpCreate) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingProjectNetadpCreate) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFirewall

`func (o *NetworkingProjectNetadpCreate) GetFirewall() string`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *NetworkingProjectNetadpCreate) GetFirewallOk() (*string, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *NetworkingProjectNetadpCreate) SetFirewall(v string)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *NetworkingProjectNetadpCreate) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetIp

`func (o *NetworkingProjectNetadpCreate) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkingProjectNetadpCreate) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkingProjectNetadpCreate) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworkingProjectNetadpCreate) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetTag

`func (o *NetworkingProjectNetadpCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NetworkingProjectNetadpCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NetworkingProjectNetadpCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *NetworkingProjectNetadpCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


