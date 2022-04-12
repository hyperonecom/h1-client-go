# Ip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**IpAssociated**](IpAssociated.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Flavour** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Persistent** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**IpProperties**](IpProperties.md) |  | [optional] 
**PtrRecord** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewIp

`func NewIp() *Ip`

NewIp instantiates a new Ip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpWithDefaults

`func NewIpWithDefaults() *Ip`

NewIpWithDefaults instantiates a new Ip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Ip) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Ip) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Ip) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Ip) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAssociated

`func (o *Ip) GetAssociated() IpAssociated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *Ip) GetAssociatedOk() (*IpAssociated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *Ip) SetAssociated(v IpAssociated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *Ip) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Ip) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Ip) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Ip) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Ip) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Ip) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Ip) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Ip) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Ip) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetFlavour

`func (o *Ip) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Ip) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Ip) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.

### HasFlavour

`func (o *Ip) HasFlavour() bool`

HasFlavour returns a boolean if a field has been set.

### GetFqdn

`func (o *Ip) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Ip) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Ip) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Ip) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetId

`func (o *Ip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ip) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ip) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Ip) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Ip) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Ip) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Ip) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Ip) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Ip) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Ip) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Ip) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *Ip) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ip) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ip) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ip) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *Ip) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Ip) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Ip) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Ip) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPersistent

`func (o *Ip) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *Ip) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *Ip) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *Ip) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetProject

`func (o *Ip) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Ip) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Ip) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Ip) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProperties

`func (o *Ip) GetProperties() IpProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Ip) GetPropertiesOk() (*IpProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Ip) SetProperties(v IpProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Ip) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPtrRecord

`func (o *Ip) GetPtrRecord() string`

GetPtrRecord returns the PtrRecord field if non-nil, zero value otherwise.

### GetPtrRecordOk

`func (o *Ip) GetPtrRecordOk() (*string, bool)`

GetPtrRecordOk returns a tuple with the PtrRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrRecord

`func (o *Ip) SetPtrRecord(v string)`

SetPtrRecord sets PtrRecord field to given value.

### HasPtrRecord

`func (o *Ip) HasPtrRecord() bool`

HasPtrRecord returns a boolean if a field has been set.

### GetState

`func (o *Ip) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Ip) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Ip) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Ip) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *Ip) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Ip) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Ip) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Ip) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUri

`func (o *Ip) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Ip) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Ip) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Ip) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


