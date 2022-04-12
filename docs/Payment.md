# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Credits** | Pointer to **float32** |  | [optional] 
**CreditsFree** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Payment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Payment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannel

`func (o *Payment) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Payment) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Payment) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Payment) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Payment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Payment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Payment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Payment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Payment) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Payment) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Payment) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Payment) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCredits

`func (o *Payment) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Payment) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Payment) SetCredits(v float32)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *Payment) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetCreditsFree

`func (o *Payment) GetCreditsFree() float32`

GetCreditsFree returns the CreditsFree field if non-nil, zero value otherwise.

### GetCreditsFreeOk

`func (o *Payment) GetCreditsFreeOk() (*float32, bool)`

GetCreditsFreeOk returns a tuple with the CreditsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsFree

`func (o *Payment) SetCreditsFree(v float32)`

SetCreditsFree sets CreditsFree field to given value.

### HasCreditsFree

`func (o *Payment) HasCreditsFree() bool`

HasCreditsFree returns a boolean if a field has been set.

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Payment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Payment) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Payment) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Payment) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Payment) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Payment) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Payment) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Payment) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Payment) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *Payment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Payment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Payment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Payment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *Payment) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Payment) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Payment) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Payment) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *Payment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Payment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Payment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Payment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUri

`func (o *Payment) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Payment) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Payment) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Payment) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


