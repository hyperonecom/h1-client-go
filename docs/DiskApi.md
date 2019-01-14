# \DiskApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionDiskResize**](DiskApi.md#ActionDiskResize) | **Post** /disk/{diskId}/actions/resize | /actions/resize
[**ActionDiskTransfer**](DiskApi.md#ActionDiskTransfer) | **Post** /disk/{diskId}/actions/transfer | /actions/transfer
[**CreateDisk**](DiskApi.md#CreateDisk) | **Post** /disk | Create
[**DeleteDisk**](DiskApi.md#DeleteDisk) | **Delete** /disk/{diskId} | Delete
[**ListDisk**](DiskApi.md#ListDisk) | **Get** /disk | List
[**OperationDiskDeleteaccessrightsIdentity**](DiskApi.md#OperationDiskDeleteaccessrightsIdentity) | **Delete** /disk/{diskId}/accessrights/{identity} | /accessrights/:identity
[**OperationDiskDeletetagKey**](DiskApi.md#OperationDiskDeletetagKey) | **Delete** /disk/{diskId}/tag/{key} | /tag/:key
[**OperationDiskGetservicesServiceId**](DiskApi.md#OperationDiskGetservicesServiceId) | **Get** /disk/{diskId}/services/{serviceId} | /services/:serviceId
[**OperationDiskGettag**](DiskApi.md#OperationDiskGettag) | **Get** /disk/{diskId}/tag/ | /tag/
[**OperationDiskListaccessrights**](DiskApi.md#OperationDiskListaccessrights) | **Get** /disk/{diskId}/accessrights/ | /accessrights/
[**OperationDiskListqueue**](DiskApi.md#OperationDiskListqueue) | **Get** /disk/{diskId}/queue/ | /queue/
[**OperationDiskListservices**](DiskApi.md#OperationDiskListservices) | **Get** /disk/{diskId}/services/ | /services/
[**OperationDiskPatchtag**](DiskApi.md#OperationDiskPatchtag) | **Patch** /disk/{diskId}/tag/ | /tag/
[**OperationDiskPostaccessrights**](DiskApi.md#OperationDiskPostaccessrights) | **Post** /disk/{diskId}/accessrights/ | /accessrights/
[**ShowDisk**](DiskApi.md#ShowDisk) | **Get** /disk/{diskId} | Get
[**UpdateDisk**](DiskApi.md#UpdateDisk) | **Patch** /disk/{diskId} | Update


# **ActionDiskResize**
> Disk ActionDiskResize(ctx, diskId)
/actions/resize

Action resize

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

# **ActionDiskTransfer**
> Disk ActionDiskTransfer(ctx, diskId, optional)
/actions/transfer

Action transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
 **optional** | ***ActionDiskTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionDiskTransferOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject29** | [**optional.Interface of InlineObject29**](InlineObject29.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDisk**
> Disk CreateDisk(ctx, optional)
Create

Create disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDiskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateDiskOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject27** | [**optional.Interface of InlineObject27**](InlineObject27.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDisk**
> DeleteDisk(ctx, diskId)
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

# **ListDisk**
> []Disk ListDisk(ctx, optional)
List

List disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDiskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDiskOpts struct

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

# **OperationDiskDeleteaccessrightsIdentity**
> Disk OperationDiskDeleteaccessrightsIdentity(ctx, diskId, identity)
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

# **OperationDiskDeletetagKey**
> map[string]string OperationDiskDeletetagKey(ctx, diskId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationDiskGetservicesServiceId**
> DiskServices OperationDiskGetservicesServiceId(ctx, diskId, serviceId)
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

# **OperationDiskGettag**
> map[string]string OperationDiskGettag(ctx, diskId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationDiskListaccessrights**
> []string OperationDiskListaccessrights(ctx, diskId)
/accessrights/

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

# **OperationDiskListqueue**
> []Event OperationDiskListqueue(ctx, diskId)
/queue/

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

# **OperationDiskListservices**
> []DiskServices OperationDiskListservices(ctx, diskId)
/services/

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

# **OperationDiskPatchtag**
> map[string]string OperationDiskPatchtag(ctx, diskId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationDiskPostaccessrights**
> string OperationDiskPostaccessrights(ctx, diskId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
 **optional** | ***OperationDiskPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationDiskPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject30** | [**optional.Interface of InlineObject30**](InlineObject30.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowDisk**
> Disk ShowDisk(ctx, diskId)
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

# **UpdateDisk**
> Disk UpdateDisk(ctx, diskId, optional)
Update

Returns modified disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diskId** | **string**| ID of disk | 
 **optional** | ***UpdateDiskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDiskOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject28** | [**optional.Interface of InlineObject28**](InlineObject28.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

