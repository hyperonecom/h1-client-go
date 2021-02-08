# Reservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Flavour** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Assigned** | Pointer to **string** |  | [optional] 
**SourceService** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewReservation

`func NewReservation() *Reservation`

NewReservation instantiates a new Reservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationWithDefaults

`func NewReservationWithDefaults() *Reservation`

NewReservationWithDefaults instantiates a new Reservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Reservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Reservation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reservation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reservation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Reservation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlavour

`func (o *Reservation) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Reservation) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Reservation) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.

### HasFlavour

`func (o *Reservation) HasFlavour() bool`

HasFlavour returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Reservation) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Reservation) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Reservation) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Reservation) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Reservation) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Reservation) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Reservation) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Reservation) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Reservation) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Reservation) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Reservation) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Reservation) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Reservation) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Reservation) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Reservation) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Reservation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetState

`func (o *Reservation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Reservation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Reservation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Reservation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProject

`func (o *Reservation) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Reservation) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Reservation) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Reservation) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetUri

`func (o *Reservation) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Reservation) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Reservation) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Reservation) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetKind

`func (o *Reservation) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Reservation) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Reservation) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Reservation) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAssigned

`func (o *Reservation) GetAssigned() string`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *Reservation) GetAssignedOk() (*string, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *Reservation) SetAssigned(v string)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *Reservation) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetSourceService

`func (o *Reservation) GetSourceService() string`

GetSourceService returns the SourceService field if non-nil, zero value otherwise.

### GetSourceServiceOk

`func (o *Reservation) GetSourceServiceOk() (*string, bool)`

GetSourceServiceOk returns a tuple with the SourceService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceService

`func (o *Reservation) SetSourceService(v string)`

SetSourceService sets SourceService field to given value.

### HasSourceService

`func (o *Reservation) HasSourceService() bool`

HasSourceService returns a boolean if a field has been set.

### GetTag

`func (o *Reservation) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Reservation) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Reservation) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Reservation) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


