# ComputeProjectVmCreateCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **string** | - user credential (by id or name) - project credential (by id or name) - raw openssh public key (starting with &#x60;ssh-rsa/ssh-ed25519&#x60;) - defaults to user keys | 

## Methods

### NewComputeProjectVmCreateCredential

`func NewComputeProjectVmCreateCredential(type_ string, value string, ) *ComputeProjectVmCreateCredential`

NewComputeProjectVmCreateCredential instantiates a new ComputeProjectVmCreateCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeProjectVmCreateCredentialWithDefaults

`func NewComputeProjectVmCreateCredentialWithDefaults() *ComputeProjectVmCreateCredential`

NewComputeProjectVmCreateCredentialWithDefaults instantiates a new ComputeProjectVmCreateCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComputeProjectVmCreateCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputeProjectVmCreateCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputeProjectVmCreateCredential) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ComputeProjectVmCreateCredential) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComputeProjectVmCreateCredential) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComputeProjectVmCreateCredential) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


