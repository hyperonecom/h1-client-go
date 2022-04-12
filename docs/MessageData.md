# MessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Mime** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMessageData

`func NewMessageData(mime string, ) *MessageData`

NewMessageData instantiates a new MessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataWithDefaults

`func NewMessageDataWithDefaults() *MessageData`

NewMessageDataWithDefaults instantiates a new MessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *MessageData) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MessageData) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MessageData) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MessageData) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMime

`func (o *MessageData) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *MessageData) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *MessageData) SetMime(v string)`

SetMime sets Mime field to given value.


### GetUrl

`func (o *MessageData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessageData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessageData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessageData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


