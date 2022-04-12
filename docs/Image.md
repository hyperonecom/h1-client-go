# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Flavour** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**License** | Pointer to **[]string** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *Image) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Image) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Image) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Image) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Image) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Image) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Image) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Image) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDescription

`func (o *Image) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Image) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Image) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Image) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlavour

`func (o *Image) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Image) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Image) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.

### HasFlavour

`func (o *Image) HasFlavour() bool`

HasFlavour returns a boolean if a field has been set.

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Image) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicense

`func (o *Image) GetLicense() []string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Image) GetLicenseOk() (*[]string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Image) SetLicense(v []string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Image) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Image) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Image) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Image) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Image) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Image) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Image) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Image) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Image) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Image) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *Image) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Image) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Image) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Image) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSize

`func (o *Image) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Image) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Image) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Image) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *Image) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Image) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Image) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Image) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *Image) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Image) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Image) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Image) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUri

`func (o *Image) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Image) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Image) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Image) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


