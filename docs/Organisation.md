# Organisation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Services** | **string** |  | [optional] 
**Flavour** | **string** |  | [optional] 
**ModifiedOn** | [**time.Time**](time.Time.md) |  | [optional] 
**ModifiedBy** | **string** |  | [optional] 
**CreatedBy** | **string** |  | [optional] 
**CreatedOn** | [**time.Time**](time.Time.md) |  | [optional] 
**AccessRights** | [**[]OrganisationAccessRights**](organisation_accessRights.md) |  | [optional] 
**Processing** | **bool** |  | [optional] 
**Created** | **bool** |  | [optional] 
**Queue** | [**[]Event**](event.md) |  | [optional] 
**State** | **string** |  | [optional] 
**Tag** | [**map[string]interface{}**](.md) |  | [optional] 
**Project** | **string** |  | [optional] 
**Billing** | [**OrganisationBilling**](organisation_billing.md) |  | [optional] 
**Verified** | **float32** |  | [optional] 
**Limit** | [**OrganisationLimit**](organisation_limit.md) |  | [optional] 
**Roles** | [**[]ProjectRoles**](project_roles.md) |  | [optional] 
**Transfer** | [**OrganisationTransfer**](organisation_transfer.md) |  | [optional] 
**BankAccount** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


