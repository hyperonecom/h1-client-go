# Container

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
**AccessRights** | **[]string** |  | [optional] 
**Processing** | **bool** |  | [optional] 
**Created** | **bool** |  | [optional] 
**Queue** | [**[]Event**](event.md) |  | [optional] 
**State** | **string** |  | [optional] 
**Tag** | **map[string]string** |  | [optional] 
**Project** | **string** |  | [optional] 
**Image** | **string** |  | [optional] 
**Command** | **string** |  | [optional] 
**Volumes** | [**[]ContainerVolumes**](container_volumes.md) |  | [optional] 
**Expose** | [**[]ContainerExpose**](container_expose.md) |  | [optional] 
**Env** | **[]string** |  | [optional] 
**Fqdn** | **string** |  | [optional] 
**Process** | [**ContainerProcess**](container_process.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


