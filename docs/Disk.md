# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Flavour** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Image** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Project** | Pointer to **string** |  | [optional] 
**Replica** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Vm** | Pointer to **string** |  | [optional] 

## Methods

### NewDisk

`func NewDisk(id string, name string, ) *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *Disk) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Disk) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Disk) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Disk) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Disk) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Disk) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Disk) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Disk) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetFlavour

`func (o *Disk) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Disk) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Disk) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.

### HasFlavour

`func (o *Disk) HasFlavour() bool`

HasFlavour returns a boolean if a field has been set.

### GetId

`func (o *Disk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *Disk) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Disk) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Disk) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Disk) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Disk) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Disk) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Disk) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Disk) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Disk) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Disk) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Disk) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Disk) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *Disk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Disk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Disk) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *Disk) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Disk) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Disk) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Disk) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReplica

`func (o *Disk) GetReplica() string`

GetReplica returns the Replica field if non-nil, zero value otherwise.

### GetReplicaOk

`func (o *Disk) GetReplicaOk() (*string, bool)`

GetReplicaOk returns a tuple with the Replica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplica

`func (o *Disk) SetReplica(v string)`

SetReplica sets Replica field to given value.

### HasReplica

`func (o *Disk) HasReplica() bool`

HasReplica returns a boolean if a field has been set.

### GetSize

`func (o *Disk) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Disk) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Disk) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Disk) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *Disk) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Disk) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Disk) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Disk) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *Disk) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Disk) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Disk) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Disk) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUri

`func (o *Disk) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Disk) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Disk) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Disk) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVm

`func (o *Disk) GetVm() string`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *Disk) GetVmOk() (*string, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *Disk) SetVm(v string)`

SetVm sets Vm field to given value.

### HasVm

`func (o *Disk) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


