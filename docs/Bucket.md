# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Flavour** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organisation** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**SizeUsed** | Pointer to **float32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewBucket

`func NewBucket() *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *Bucket) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Bucket) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Bucket) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Bucket) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Bucket) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Bucket) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Bucket) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Bucket) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetFlavour

`func (o *Bucket) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Bucket) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Bucket) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.

### HasFlavour

`func (o *Bucket) HasFlavour() bool`

HasFlavour returns a boolean if a field has been set.

### GetId

`func (o *Bucket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bucket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bucket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Bucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Bucket) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Bucket) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Bucket) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Bucket) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Bucket) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Bucket) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Bucket) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Bucket) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganisation

`func (o *Bucket) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *Bucket) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *Bucket) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *Bucket) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### GetProject

`func (o *Bucket) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Bucket) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Bucket) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Bucket) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSizeUsed

`func (o *Bucket) GetSizeUsed() float32`

GetSizeUsed returns the SizeUsed field if non-nil, zero value otherwise.

### GetSizeUsedOk

`func (o *Bucket) GetSizeUsedOk() (*float32, bool)`

GetSizeUsedOk returns a tuple with the SizeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeUsed

`func (o *Bucket) SetSizeUsed(v float32)`

SetSizeUsed sets SizeUsed field to given value.

### HasSizeUsed

`func (o *Bucket) HasSizeUsed() bool`

HasSizeUsed returns a boolean if a field has been set.

### GetState

`func (o *Bucket) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Bucket) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Bucket) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Bucket) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUri

`func (o *Bucket) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Bucket) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Bucket) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Bucket) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


