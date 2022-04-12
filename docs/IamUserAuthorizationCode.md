# IamUserAuthorizationCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **string** |  | [optional] 
**Redirect** | **string** |  | 

## Methods

### NewIamUserAuthorizationCode

`func NewIamUserAuthorizationCode(redirect string, ) *IamUserAuthorizationCode`

NewIamUserAuthorizationCode instantiates a new IamUserAuthorizationCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserAuthorizationCodeWithDefaults

`func NewIamUserAuthorizationCodeWithDefaults() *IamUserAuthorizationCode`

NewIamUserAuthorizationCodeWithDefaults instantiates a new IamUserAuthorizationCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *IamUserAuthorizationCode) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *IamUserAuthorizationCode) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *IamUserAuthorizationCode) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *IamUserAuthorizationCode) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetRedirect

`func (o *IamUserAuthorizationCode) GetRedirect() string`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *IamUserAuthorizationCode) GetRedirectOk() (*string, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *IamUserAuthorizationCode) SetRedirect(v string)`

SetRedirect sets Redirect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


