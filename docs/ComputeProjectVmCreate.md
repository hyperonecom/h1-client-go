# ComputeProjectVmCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**[]ComputeProjectVmCreateCredential**](ComputeProjectVmCreateCredential.md) |  | [optional] 
**Disk** | Pointer to [**[]ComputeProjectVmCreateDisk**](ComputeProjectVmCreateDisk.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Iso** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Netadp** | Pointer to [**[]ComputeProjectVmCreateNetadp**](ComputeProjectVmCreateNetadp.md) |  | [optional] 
**Service** | **string** |  | 
**Start** | Pointer to **bool** |  | [optional] [default to true]
**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**UserMetadata** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewComputeProjectVmCreate

`func NewComputeProjectVmCreate(name string, service string, ) *ComputeProjectVmCreate`

NewComputeProjectVmCreate instantiates a new ComputeProjectVmCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeProjectVmCreateWithDefaults

`func NewComputeProjectVmCreateWithDefaults() *ComputeProjectVmCreate`

NewComputeProjectVmCreateWithDefaults instantiates a new ComputeProjectVmCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *ComputeProjectVmCreate) GetCredential() []ComputeProjectVmCreateCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ComputeProjectVmCreate) GetCredentialOk() (*[]ComputeProjectVmCreateCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ComputeProjectVmCreate) SetCredential(v []ComputeProjectVmCreateCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ComputeProjectVmCreate) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDisk

`func (o *ComputeProjectVmCreate) GetDisk() []ComputeProjectVmCreateDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ComputeProjectVmCreate) GetDiskOk() (*[]ComputeProjectVmCreateDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ComputeProjectVmCreate) SetDisk(v []ComputeProjectVmCreateDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ComputeProjectVmCreate) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetImage

`func (o *ComputeProjectVmCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ComputeProjectVmCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ComputeProjectVmCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ComputeProjectVmCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIso

`func (o *ComputeProjectVmCreate) GetIso() string`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *ComputeProjectVmCreate) GetIsoOk() (*string, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *ComputeProjectVmCreate) SetIso(v string)`

SetIso sets Iso field to given value.

### HasIso

`func (o *ComputeProjectVmCreate) HasIso() bool`

HasIso returns a boolean if a field has been set.

### GetName

`func (o *ComputeProjectVmCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeProjectVmCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeProjectVmCreate) SetName(v string)`

SetName sets Name field to given value.


### GetNetadp

`func (o *ComputeProjectVmCreate) GetNetadp() []ComputeProjectVmCreateNetadp`

GetNetadp returns the Netadp field if non-nil, zero value otherwise.

### GetNetadpOk

`func (o *ComputeProjectVmCreate) GetNetadpOk() (*[]ComputeProjectVmCreateNetadp, bool)`

GetNetadpOk returns a tuple with the Netadp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetadp

`func (o *ComputeProjectVmCreate) SetNetadp(v []ComputeProjectVmCreateNetadp)`

SetNetadp sets Netadp field to given value.

### HasNetadp

`func (o *ComputeProjectVmCreate) HasNetadp() bool`

HasNetadp returns a boolean if a field has been set.

### GetService

`func (o *ComputeProjectVmCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ComputeProjectVmCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ComputeProjectVmCreate) SetService(v string)`

SetService sets Service field to given value.


### GetStart

`func (o *ComputeProjectVmCreate) GetStart() bool`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ComputeProjectVmCreate) GetStartOk() (*bool, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ComputeProjectVmCreate) SetStart(v bool)`

SetStart sets Start field to given value.

### HasStart

`func (o *ComputeProjectVmCreate) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTag

`func (o *ComputeProjectVmCreate) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ComputeProjectVmCreate) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ComputeProjectVmCreate) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ComputeProjectVmCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUserMetadata

`func (o *ComputeProjectVmCreate) GetUserMetadata() string`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *ComputeProjectVmCreate) GetUserMetadataOk() (*string, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *ComputeProjectVmCreate) SetUserMetadata(v string)`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *ComputeProjectVmCreate) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.

### GetUsername

`func (o *ComputeProjectVmCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ComputeProjectVmCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ComputeProjectVmCreate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ComputeProjectVmCreate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


