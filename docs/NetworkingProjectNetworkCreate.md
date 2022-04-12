# NetworkingProjectNetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewNetworkingProjectNetworkCreate

`func NewNetworkingProjectNetworkCreate(name string, ) *NetworkingProjectNetworkCreate`

NewNetworkingProjectNetworkCreate instantiates a new NetworkingProjectNetworkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingProjectNetworkCreateWithDefaults

`func NewNetworkingProjectNetworkCreateWithDefaults() *NetworkingProjectNetworkCreate`

NewNetworkingProjectNetworkCreateWithDefaults instantiates a new NetworkingProjectNetworkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NetworkingProjectNetworkCreate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworkingProjectNetworkCreate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworkingProjectNetworkCreate) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NetworkingProjectNetworkCreate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *NetworkingProjectNetworkCreate) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkingProjectNetworkCreate) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkingProjectNetworkCreate) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkingProjectNetworkCreate) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetName

`func (o *NetworkingProjectNetworkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkingProjectNetworkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkingProjectNetworkCreate) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *NetworkingProjectNetworkCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NetworkingProjectNetworkCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NetworkingProjectNetworkCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *NetworkingProjectNetworkCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


