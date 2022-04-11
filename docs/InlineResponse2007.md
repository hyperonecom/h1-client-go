# InlineResponse2007

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**RecoveryLocationIdProjectProjectIdPlanProperties**](RecoveryLocationIdProjectProjectIdPlanProperties.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedBy** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Project** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Fqdn** | Pointer to **string** |  | [optional] [readonly] 
**Tag** | Pointer to [**[]DatabaseLocationIdProjectProjectIdInstanceTag**](DatabaseLocationIdProjectProjectIdInstanceTag.md) |  | [optional] 
**Profile** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Uri** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewInlineResponse2007

`func NewInlineResponse2007(profile string, name string, ) *InlineResponse2007`

NewInlineResponse2007 instantiates a new InlineResponse2007 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2007WithDefaults

`func NewInlineResponse2007WithDefaults() *InlineResponse2007`

NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *InlineResponse2007) GetProperties() RecoveryLocationIdProjectProjectIdPlanProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InlineResponse2007) GetPropertiesOk() (*RecoveryLocationIdProjectProjectIdPlanProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InlineResponse2007) SetProperties(v RecoveryLocationIdProjectProjectIdPlanProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InlineResponse2007) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InlineResponse2007) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse2007) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse2007) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse2007) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *InlineResponse2007) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *InlineResponse2007) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *InlineResponse2007) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *InlineResponse2007) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *InlineResponse2007) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *InlineResponse2007) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *InlineResponse2007) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *InlineResponse2007) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetProject

`func (o *InlineResponse2007) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InlineResponse2007) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InlineResponse2007) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *InlineResponse2007) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *InlineResponse2007) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse2007) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse2007) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InlineResponse2007) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFqdn

`func (o *InlineResponse2007) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *InlineResponse2007) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *InlineResponse2007) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *InlineResponse2007) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetTag

`func (o *InlineResponse2007) GetTag() []DatabaseLocationIdProjectProjectIdInstanceTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *InlineResponse2007) GetTagOk() (*[]DatabaseLocationIdProjectProjectIdInstanceTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *InlineResponse2007) SetTag(v []DatabaseLocationIdProjectProjectIdInstanceTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *InlineResponse2007) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse2007) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse2007) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse2007) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetId

`func (o *InlineResponse2007) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2007) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2007) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2007) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2007) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2007) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2007) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedOn

`func (o *InlineResponse2007) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *InlineResponse2007) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *InlineResponse2007) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *InlineResponse2007) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUri

`func (o *InlineResponse2007) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineResponse2007) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineResponse2007) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineResponse2007) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


