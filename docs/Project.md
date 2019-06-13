# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Services** | [**[]ProjectServices**](project_services.md) |  | [optional] 
**Flavour** | **string** |  | [optional] 
**ModifiedOn** | [**time.Time**](time.Time.md) |  | [optional] 
**ModifiedBy** | **string** |  | [optional] [default to ]
**CreatedBy** | **string** |  | [optional] [default to ]
**CreatedOn** | [**time.Time**](time.Time.md) |  | [optional] 
**AccessRights** | [**[]ProjectAccessRights**](project_accessRights.md) |  | [optional] 
**Processing** | **bool** |  | [optional] 
**Created** | **bool** |  | [optional] 
**Queue** | [**[]Event**](event.md) |  | [optional] 
**State** | **string** |  | [optional] 
**Tag** | [**map[string]interface{}**](.md) |  | [optional] [default to {}]
**Project** | **string** |  | [optional] 
**BankAccount** | **string** |  | [optional] 
**Organisation** | **string** |  | [optional] 
**Billing** | [**ProjectBilling**](project_billing.md) |  | [optional] 
**Invoices** | [**[]ProjectInvoices**](project_invoices.md) |  | [optional] 
**Payments** | [**[]Payment**](payment.md) |  | [optional] 
**Verified** | **string** |  | [optional] 
**Active** | **bool** |  | [optional] [default to true]
**Limit** | [**ProjectLimit**](project_limit.md) |  | [optional] 
**Threshold** | [**ProjectThreshold**](project_threshold.md) |  | [optional] 
**Roles** | [**[]ProjectRoles**](project_roles.md) |  | [optional] 
**NetworkAcl** | [**[]UserNetworkAcl**](user_networkAcl.md) |  | [optional] 
**Compliance** | [**ProjectCompliance**](project_compliance.md) |  | [optional] 
**Transfer** | [**ProjectTransfer**](project_transfer.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


