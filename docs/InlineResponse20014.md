# InlineResponse20014

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**OneOfobjectobject**](oneOf&lt;object,object&gt;.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedBy** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Project** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Fqdn** | Pointer to **string** |  | [optional] [readonly] 
**Tag** | Pointer to [**[]WebsiteLocationIdProjectProjectIdInstanceTag**](WebsiteLocationIdProjectProjectIdInstanceTag.md) |  | [optional] 
**Profile** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Uri** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewInlineResponse20014

`func NewInlineResponse20014(profile string, name string, ) *InlineResponse20014`

NewInlineResponse20014 instantiates a new InlineResponse20014 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20014WithDefaults

`func NewInlineResponse20014WithDefaults() *InlineResponse20014`

NewInlineResponse20014WithDefaults instantiates a new InlineResponse20014 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *InlineResponse20014) GetProperties() OneOfobjectobject`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InlineResponse20014) GetPropertiesOk() (*OneOfobjectobject, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InlineResponse20014) SetProperties(v OneOfobjectobject)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InlineResponse20014) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *InlineResponse20014) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *InlineResponse20014) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetCreatedBy

`func (o *InlineResponse20014) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse20014) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse20014) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse20014) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *InlineResponse20014) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *InlineResponse20014) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *InlineResponse20014) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *InlineResponse20014) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *InlineResponse20014) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *InlineResponse20014) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *InlineResponse20014) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *InlineResponse20014) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetProject

`func (o *InlineResponse20014) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InlineResponse20014) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InlineResponse20014) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *InlineResponse20014) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *InlineResponse20014) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse20014) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse20014) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InlineResponse20014) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFqdn

`func (o *InlineResponse20014) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *InlineResponse20014) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *InlineResponse20014) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *InlineResponse20014) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetTag

`func (o *InlineResponse20014) GetTag() []WebsiteLocationIdProjectProjectIdInstanceTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *InlineResponse20014) GetTagOk() (*[]WebsiteLocationIdProjectProjectIdInstanceTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *InlineResponse20014) SetTag(v []WebsiteLocationIdProjectProjectIdInstanceTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *InlineResponse20014) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse20014) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse20014) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse20014) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetId

`func (o *InlineResponse20014) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20014) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20014) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20014) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20014) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20014) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20014) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedOn

`func (o *InlineResponse20014) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *InlineResponse20014) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *InlineResponse20014) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *InlineResponse20014) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUri

`func (o *InlineResponse20014) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineResponse20014) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineResponse20014) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineResponse20014) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


