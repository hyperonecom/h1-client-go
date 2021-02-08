# NetworkingProjectFirewallCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | Pointer to **string** |  | [optional] [default to "5bacaf7202deee0c100eda3b"]
**Ingress** | Pointer to [**[]NetworkingRule**](NetworkingRule.md) |  | [optional] 
**Egress** | Pointer to [**[]NetworkingRule**](NetworkingRule.md) |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewNetworkingProjectFirewallCreate

`func NewNetworkingProjectFirewallCreate(name string, ) *NetworkingProjectFirewallCreate`

NewNetworkingProjectFirewallCreate instantiates a new NetworkingProjectFirewallCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingProjectFirewallCreateWithDefaults

`func NewNetworkingProjectFirewallCreateWithDefaults() *NetworkingProjectFirewallCreate`

NewNetworkingProjectFirewallCreateWithDefaults instantiates a new NetworkingProjectFirewallCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkingProjectFirewallCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkingProjectFirewallCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkingProjectFirewallCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *NetworkingProjectFirewallCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *NetworkingProjectFirewallCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *NetworkingProjectFirewallCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *NetworkingProjectFirewallCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetIngress

`func (o *NetworkingProjectFirewallCreate) GetIngress() []NetworkingRule`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *NetworkingProjectFirewallCreate) GetIngressOk() (*[]NetworkingRule, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *NetworkingProjectFirewallCreate) SetIngress(v []NetworkingRule)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *NetworkingProjectFirewallCreate) HasIngress() bool`

HasIngress returns a boolean if a field has been set.

### GetEgress

`func (o *NetworkingProjectFirewallCreate) GetEgress() []NetworkingRule`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *NetworkingProjectFirewallCreate) GetEgressOk() (*[]NetworkingRule, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *NetworkingProjectFirewallCreate) SetEgress(v []NetworkingRule)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *NetworkingProjectFirewallCreate) HasEgress() bool`

HasEgress returns a boolean if a field has been set.

### GetTag

`func (o *NetworkingProjectFirewallCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NetworkingProjectFirewallCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NetworkingProjectFirewallCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *NetworkingProjectFirewallCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


