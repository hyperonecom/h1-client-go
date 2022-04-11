# IpProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **float32** |  | [optional] 

## Methods

### NewIpProperties

`func NewIpProperties() *IpProperties`

NewIpProperties instantiates a new IpProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpPropertiesWithDefaults

`func NewIpPropertiesWithDefaults() *IpProperties`

NewIpPropertiesWithDefaults instantiates a new IpProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *IpProperties) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *IpProperties) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *IpProperties) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *IpProperties) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetVersion

`func (o *IpProperties) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IpProperties) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IpProperties) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IpProperties) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


