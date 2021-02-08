# BillingProjectReservationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Service** | **string** |  | 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewBillingProjectReservationCreate

`func NewBillingProjectReservationCreate(name string, service string, ) *BillingProjectReservationCreate`

NewBillingProjectReservationCreate instantiates a new BillingProjectReservationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProjectReservationCreateWithDefaults

`func NewBillingProjectReservationCreateWithDefaults() *BillingProjectReservationCreate`

NewBillingProjectReservationCreateWithDefaults instantiates a new BillingProjectReservationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingProjectReservationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingProjectReservationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingProjectReservationCreate) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *BillingProjectReservationCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingProjectReservationCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingProjectReservationCreate) SetService(v string)`

SetService sets Service field to given value.


### GetTag

`func (o *BillingProjectReservationCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BillingProjectReservationCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BillingProjectReservationCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BillingProjectReservationCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


