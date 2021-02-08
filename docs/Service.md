# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Billing** | Pointer to [**ServiceBilling**](service_billing.md) |  | [optional] 
**Display** | Pointer to [**ServiceDisplay**](service_display.md) |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AvailableServices** | Pointer to **[]string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResource

`func (o *Service) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Service) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Service) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Service) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBilling

`func (o *Service) GetBilling() ServiceBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Service) GetBillingOk() (*ServiceBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Service) SetBilling(v ServiceBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Service) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetDisplay

`func (o *Service) GetDisplay() ServiceDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Service) GetDisplayOk() (*ServiceDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Service) SetDisplay(v ServiceDisplay)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Service) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetData

`func (o *Service) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Service) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Service) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *Service) HasData() bool`

HasData returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvailableServices

`func (o *Service) GetAvailableServices() []string`

GetAvailableServices returns the AvailableServices field if non-nil, zero value otherwise.

### GetAvailableServicesOk

`func (o *Service) GetAvailableServicesOk() (*[]string, bool)`

GetAvailableServicesOk returns a tuple with the AvailableServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableServices

`func (o *Service) SetAvailableServices(v []string)`

SetAvailableServices sets AvailableServices field to given value.

### HasAvailableServices

`func (o *Service) HasAvailableServices() bool`

HasAvailableServices returns a boolean if a field has been set.

### GetUri

`func (o *Service) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Service) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Service) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Service) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


