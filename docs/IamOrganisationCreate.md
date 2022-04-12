# IamOrganisationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | Pointer to [**OrganisationBilling**](OrganisationBilling.md) |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewIamOrganisationCreate

`func NewIamOrganisationCreate(name string, ) *IamOrganisationCreate`

NewIamOrganisationCreate instantiates a new IamOrganisationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamOrganisationCreateWithDefaults

`func NewIamOrganisationCreateWithDefaults() *IamOrganisationCreate`

NewIamOrganisationCreateWithDefaults instantiates a new IamOrganisationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *IamOrganisationCreate) GetBilling() OrganisationBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *IamOrganisationCreate) GetBillingOk() (*OrganisationBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *IamOrganisationCreate) SetBilling(v OrganisationBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *IamOrganisationCreate) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetName

`func (o *IamOrganisationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamOrganisationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamOrganisationCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


