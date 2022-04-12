# Iso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Flavour** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**IsoMetadata**](IsoMetadata.md) |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewIso

`func NewIso() *Iso`

NewIso instantiates a new Iso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsoWithDefaults

`func NewIsoWithDefaults() *Iso`

NewIsoWithDefaults instantiates a new Iso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *Iso) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Iso) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Iso) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Iso) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Iso) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Iso) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Iso) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Iso) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetFlavour

`func (o *Iso) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Iso) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Iso) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.

### HasFlavour

`func (o *Iso) HasFlavour() bool`

HasFlavour returns a boolean if a field has been set.

### GetId

`func (o *Iso) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Iso) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Iso) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Iso) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *Iso) GetMetadata() IsoMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Iso) GetMetadataOk() (*IsoMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Iso) SetMetadata(v IsoMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Iso) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Iso) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Iso) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Iso) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Iso) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Iso) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Iso) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Iso) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Iso) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *Iso) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Iso) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Iso) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Iso) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *Iso) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Iso) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Iso) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Iso) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSize

`func (o *Iso) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Iso) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Iso) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Iso) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *Iso) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Iso) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Iso) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Iso) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *Iso) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Iso) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Iso) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Iso) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUri

`func (o *Iso) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Iso) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Iso) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Iso) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


