# VmCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] 
**Image** | **string** |  | [optional] 
**Iso** | **string** |  | [optional] 
**Service** | **string** |  | 
**Username** | **string** |  | [optional] 
**Password** | [**VmCreatePassword**](vm_create_password.md) |  | [optional] 
**SshKeys** | **[]string** | - user credential (by id or name) - project credential (by id or name) - raw openssh public key (starting with &#x60;ssh-rsa &#x60;) | [optional] 
**UserMetadata** | **string** |  | [optional] 
**Disk** | [**[]VmCreateDisk**](vm_create_disk.md) |  | [optional] 
**Netadp** | [**[]VmCreateNetadp**](vm_create_netadp.md) |  | [optional] 
**Boot** | **bool** |  | [optional] 
**Cloud** | **string** |  | [optional] 
**Tag** | [**map[string]interface{}**](.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


