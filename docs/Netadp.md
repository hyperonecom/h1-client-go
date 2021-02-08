# Netadp

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
**Macaddress** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **float32** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Firewall** | Pointer to **string** |  | [optional] 
**Assigned** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewNetadp

`func NewNetadp() *Netadp`

NewNetadp instantiates a new Netadp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetadpWithDefaults

`func NewNetadpWithDefaults() *Netadp`

NewNetadpWithDefaults instantiates a new Netadp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Netadp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Netadp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Netadp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Netadp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Netadp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Netadp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Netadp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Netadp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlavour

`func (o *Netadp) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Netadp) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Netadp) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.

### HasFlavour

`func (o *Netadp) HasFlavour() bool`

HasFlavour returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Netadp) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Netadp) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Netadp) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Netadp) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Netadp) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Netadp) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Netadp) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Netadp) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Netadp) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Netadp) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Netadp) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Netadp) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Netadp) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Netadp) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Netadp) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Netadp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetState

`func (o *Netadp) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Netadp) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Netadp) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Netadp) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProject

`func (o *Netadp) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Netadp) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Netadp) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Netadp) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetUri

`func (o *Netadp) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Netadp) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Netadp) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Netadp) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetMacaddress

`func (o *Netadp) GetMacaddress() string`

GetMacaddress returns the Macaddress field if non-nil, zero value otherwise.

### GetMacaddressOk

`func (o *Netadp) GetMacaddressOk() (*string, bool)`

GetMacaddressOk returns a tuple with the Macaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacaddress

`func (o *Netadp) SetMacaddress(v string)`

SetMacaddress sets Macaddress field to given value.

### HasMacaddress

`func (o *Netadp) HasMacaddress() bool`

HasMacaddress returns a boolean if a field has been set.

### GetSpeed

`func (o *Netadp) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Netadp) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Netadp) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Netadp) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetFqdn

`func (o *Netadp) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Netadp) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Netadp) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Netadp) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetNetwork

`func (o *Netadp) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Netadp) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Netadp) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Netadp) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFirewall

`func (o *Netadp) GetFirewall() string`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *Netadp) GetFirewallOk() (*string, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *Netadp) SetFirewall(v string)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *Netadp) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetAssigned

`func (o *Netadp) GetAssigned() string`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *Netadp) GetAssignedOk() (*string, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *Netadp) SetAssigned(v string)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *Netadp) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetTag

`func (o *Netadp) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Netadp) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Netadp) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Netadp) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


