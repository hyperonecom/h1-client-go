# \DiskApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiskActionResize**](DiskApi.md#DiskActionResize) | **Post** /disk/{diskId}/actions/resize | /actions/resize
[**DiskActionTransfer**](DiskApi.md#DiskActionTransfer) | **Post** /disk/{diskId}/actions/transfer | /actions/transfer
[**DiskCreate**](DiskApi.md#DiskCreate) | **Post** /disk | Create
[**DiskDelete**](DiskApi.md#DiskDelete) | **Delete** /disk/{diskId} | Delete
[**DiskDeleteAccessrightsIdentity**](DiskApi.md#DiskDeleteAccessrightsIdentity) | **Delete** /disk/{diskId}/accessrights/{identity} | /accessrights/:identity
[**DiskDeleteTagKey**](DiskApi.md#DiskDeleteTagKey) | **Delete** /disk/{diskId}/tag/{key} | /tag/:key
[**DiskGetServicesServiceId**](DiskApi.md#DiskGetServicesServiceId) | **Get** /disk/{diskId}/services/{serviceId} | /services/:serviceId
[**DiskGetTag**](DiskApi.md#DiskGetTag) | **Get** /disk/{diskId}/tag | /tag
[**DiskList**](DiskApi.md#DiskList) | **Get** /disk | List
[**DiskListAccessrights**](DiskApi.md#DiskListAccessrights) | **Get** /disk/{diskId}/accessrights | /accessrights
[**DiskListQueue**](DiskApi.md#DiskListQueue) | **Get** /disk/{diskId}/queue | /queue
[**DiskListServices**](DiskApi.md#DiskListServices) | **Get** /disk/{diskId}/services | /services
[**DiskPatchTag**](DiskApi.md#DiskPatchTag) | **Patch** /disk/{diskId}/tag | /tag
[**DiskPostAccessrights**](DiskApi.md#DiskPostAccessrights) | **Post** /disk/{diskId}/accessrights | /accessrights
[**DiskShow**](DiskApi.md#DiskShow) | **Get** /disk/{diskId} | Get
[**DiskUpdate**](DiskApi.md#DiskUpdate) | **Patch** /disk/{diskId} | Update


# **DiskActionResize**
> Disk DiskActionResize(ctx, diskId, diskActionResize)
/actions/resize

Action resize

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **diskActionResize** | [**DiskActionResize**](DiskActionResize.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskActionTransfer**
> Disk DiskActionTransfer(ctx, diskId, diskActionTransfer)
/actions/transfer

Action transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **diskActionTransfer** | [**DiskActionTransfer**](DiskActionTransfer.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskCreate**
> Disk DiskCreate(ctx, diskCreate)
Create

Create disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskCreate** | [**DiskCreate**](DiskCreate.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskDelete**
> DiskDelete(ctx, diskId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskDeleteAccessrightsIdentity**
> Disk DiskDeleteAccessrightsIdentity(ctx, diskId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **identity** | **string**| identity | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskDeleteTagKey**
> map[string]interface{} DiskDeleteTagKey(ctx, diskId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskGetServicesServiceId**
> DiskServices DiskGetServicesServiceId(ctx, diskId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **serviceId** | **string**| serviceId | 

### Return type

[**DiskServices**](disk.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskGetTag**
> map[string]interface{} DiskGetTag(ctx, diskId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskList**
> []Disk DiskList(ctx, optional)
List

List disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiskListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiskListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskListAccessrights**
> []string DiskListAccessrights(ctx, diskId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskListQueue**
> []Event DiskListQueue(ctx, diskId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskListServices**
> []DiskServices DiskListServices(ctx, diskId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 

### Return type

[**[]DiskServices**](disk.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskPatchTag**
> map[string]interface{} DiskPatchTag(ctx, diskId, body)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **body** | **map[string]interface{}**|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskPostAccessrights**
> string DiskPostAccessrights(ctx, diskId, diskPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **diskPostAccessrights** | [**DiskPostAccessrights**](DiskPostAccessrights.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskShow**
> Disk DiskShow(ctx, diskId)
Get

Returns a single disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiskUpdate**
> Disk DiskUpdate(ctx, diskId, diskUpdate)
Update

Returns modified disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **diskUpdate** | [**DiskUpdate**](DiskUpdate.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

