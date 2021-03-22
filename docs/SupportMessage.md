# SupportMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**User** | Pointer to **string** |  | [optional] [readonly] 
**Data** | Pointer to [**MessageData**](MessageData.md) |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] [readonly] 
**Date** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewSupportMessage

`func NewSupportMessage(type_ string, ) *SupportMessage`

NewSupportMessage instantiates a new SupportMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportMessageWithDefaults

`func NewSupportMessageWithDefaults() *SupportMessage`

NewSupportMessageWithDefaults instantiates a new SupportMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SupportMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupportMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupportMessage) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *SupportMessage) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SupportMessage) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SupportMessage) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SupportMessage) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetData

`func (o *SupportMessage) GetData() MessageData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SupportMessage) GetDataOk() (*MessageData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SupportMessage) SetData(v MessageData)`

SetData sets Data field to given value.

### HasData

`func (o *SupportMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOrigin

`func (o *SupportMessage) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SupportMessage) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SupportMessage) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *SupportMessage) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetDate

`func (o *SupportMessage) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SupportMessage) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SupportMessage) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *SupportMessage) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


