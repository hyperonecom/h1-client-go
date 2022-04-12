# InsightProjectJournalCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Retention** | Pointer to **float32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] [default to "5c9cc2d0255c16c3e899a4ea"]
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewInsightProjectJournalCreate

`func NewInsightProjectJournalCreate(name string, ) *InsightProjectJournalCreate`

NewInsightProjectJournalCreate instantiates a new InsightProjectJournalCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightProjectJournalCreateWithDefaults

`func NewInsightProjectJournalCreateWithDefaults() *InsightProjectJournalCreate`

NewInsightProjectJournalCreateWithDefaults instantiates a new InsightProjectJournalCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InsightProjectJournalCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightProjectJournalCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightProjectJournalCreate) SetName(v string)`

SetName sets Name field to given value.


### GetRetention

`func (o *InsightProjectJournalCreate) GetRetention() float32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *InsightProjectJournalCreate) GetRetentionOk() (*float32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *InsightProjectJournalCreate) SetRetention(v float32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *InsightProjectJournalCreate) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetService

`func (o *InsightProjectJournalCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *InsightProjectJournalCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *InsightProjectJournalCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *InsightProjectJournalCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTag

`func (o *InsightProjectJournalCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *InsightProjectJournalCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *InsightProjectJournalCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *InsightProjectJournalCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


