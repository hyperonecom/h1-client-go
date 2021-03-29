# DevelopmentCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Type** | **string** |  | 
**Value** | **string** |  | 
**Fingerprint** | Pointer to **string** |  | [optional] [readonly] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewDevelopmentCredential

`func NewDevelopmentCredential(name string, type_ string, value string, ) *DevelopmentCredential`

NewDevelopmentCredential instantiates a new DevelopmentCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevelopmentCredentialWithDefaults

`func NewDevelopmentCredentialWithDefaults() *DevelopmentCredential`

NewDevelopmentCredentialWithDefaults instantiates a new DevelopmentCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DevelopmentCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevelopmentCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevelopmentCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevelopmentCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DevelopmentCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevelopmentCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevelopmentCredential) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedBy

`func (o *DevelopmentCredential) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DevelopmentCredential) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DevelopmentCredential) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DevelopmentCredential) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *DevelopmentCredential) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *DevelopmentCredential) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *DevelopmentCredential) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *DevelopmentCredential) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetType

`func (o *DevelopmentCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevelopmentCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevelopmentCredential) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *DevelopmentCredential) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DevelopmentCredential) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DevelopmentCredential) SetValue(v string)`

SetValue sets Value field to given value.


### GetFingerprint

`func (o *DevelopmentCredential) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *DevelopmentCredential) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *DevelopmentCredential) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *DevelopmentCredential) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetToken

`func (o *DevelopmentCredential) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DevelopmentCredential) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DevelopmentCredential) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DevelopmentCredential) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


