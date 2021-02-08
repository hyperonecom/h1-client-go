# ResourceService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourceService** | Pointer to **string** |  | [optional] 
**Billing** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewResourceService

`func NewResourceService() *ResourceService`

NewResourceService instantiates a new ResourceService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceServiceWithDefaults

`func NewResourceServiceWithDefaults() *ResourceService`

NewResourceServiceWithDefaults instantiates a new ResourceService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ResourceService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceService

`func (o *ResourceService) GetSourceService() string`

GetSourceService returns the SourceService field if non-nil, zero value otherwise.

### GetSourceServiceOk

`func (o *ResourceService) GetSourceServiceOk() (*string, bool)`

GetSourceServiceOk returns a tuple with the SourceService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceService

`func (o *ResourceService) SetSourceService(v string)`

SetSourceService sets SourceService field to given value.

### HasSourceService

`func (o *ResourceService) HasSourceService() bool`

HasSourceService returns a boolean if a field has been set.

### GetBilling

`func (o *ResourceService) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *ResourceService) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *ResourceService) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *ResourceService) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetData

`func (o *ResourceService) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceService) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceService) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *ResourceService) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


