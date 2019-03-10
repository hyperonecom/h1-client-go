# \VmhostApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmhostActionMoveDisk**](VmhostApi.md#VmhostActionMoveDisk) | **Post** /vmhost/{vmhostId}/actions/moveDisk | /actions/moveDisk
[**VmhostActionMoveISO**](VmhostApi.md#VmhostActionMoveISO) | **Post** /vmhost/{vmhostId}/actions/moveISO | /actions/moveISO
[**VmhostActionMoveImage**](VmhostApi.md#VmhostActionMoveImage) | **Post** /vmhost/{vmhostId}/actions/moveImage | /actions/moveImage
[**VmhostActionMoveVM**](VmhostApi.md#VmhostActionMoveVM) | **Post** /vmhost/{vmhostId}/actions/moveVM | /actions/moveVM
[**VmhostDelete**](VmhostApi.md#VmhostDelete) | **Delete** /vmhost/{vmhostId} | Delete
[**VmhostDeleteAccessrightsIdentity**](VmhostApi.md#VmhostDeleteAccessrightsIdentity) | **Delete** /vmhost/{vmhostId}/accessrights/{identity} | /accessrights/:identity
[**VmhostDeleteTagKey**](VmhostApi.md#VmhostDeleteTagKey) | **Delete** /vmhost/{vmhostId}/tag/{key} | /tag/:key
[**VmhostGetServicesServiceId**](VmhostApi.md#VmhostGetServicesServiceId) | **Get** /vmhost/{vmhostId}/services/{serviceId} | /services/:serviceId
[**VmhostGetTag**](VmhostApi.md#VmhostGetTag) | **Get** /vmhost/{vmhostId}/tag | /tag
[**VmhostList**](VmhostApi.md#VmhostList) | **Get** /vmhost | List
[**VmhostListAccessrights**](VmhostApi.md#VmhostListAccessrights) | **Get** /vmhost/{vmhostId}/accessrights | /accessrights
[**VmhostListQueue**](VmhostApi.md#VmhostListQueue) | **Get** /vmhost/{vmhostId}/queue | /queue
[**VmhostListServices**](VmhostApi.md#VmhostListServices) | **Get** /vmhost/{vmhostId}/services | /services
[**VmhostPatchTag**](VmhostApi.md#VmhostPatchTag) | **Patch** /vmhost/{vmhostId}/tag | /tag
[**VmhostPostAccessrights**](VmhostApi.md#VmhostPostAccessrights) | **Post** /vmhost/{vmhostId}/accessrights | /accessrights
[**VmhostShow**](VmhostApi.md#VmhostShow) | **Get** /vmhost/{vmhostId} | Get
[**VmhostUpdate**](VmhostApi.md#VmhostUpdate) | **Patch** /vmhost/{vmhostId} | Update


# **VmhostActionMoveDisk**
> Vmhost VmhostActionMoveDisk(ctx, vmhostId, vmhostActionMoveDisk)
/actions/moveDisk

Action moveDisk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **vmhostActionMoveDisk** | [**VmhostActionMoveDisk**](VmhostActionMoveDisk.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostActionMoveISO**
> Vmhost VmhostActionMoveISO(ctx, vmhostId, vmhostActionMoveIso)
/actions/moveISO

Action moveISO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **vmhostActionMoveIso** | [**VmhostActionMoveIso**](VmhostActionMoveIso.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostActionMoveImage**
> Vmhost VmhostActionMoveImage(ctx, vmhostId, vmhostActionMoveImage)
/actions/moveImage

Action moveImage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **vmhostActionMoveImage** | [**VmhostActionMoveImage**](VmhostActionMoveImage.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostActionMoveVM**
> Vmhost VmhostActionMoveVM(ctx, vmhostId, vmhostActionMoveVm)
/actions/moveVM

Action moveVM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **vmhostActionMoveVm** | [**VmhostActionMoveVm**](VmhostActionMoveVm.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostDelete**
> VmhostDelete(ctx, vmhostId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostDeleteAccessrightsIdentity**
> Vmhost VmhostDeleteAccessrightsIdentity(ctx, vmhostId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **identity** | **string**| identity | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostDeleteTagKey**
> map[string]interface{} VmhostDeleteTagKey(ctx, vmhostId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostGetServicesServiceId**
> VmhostServices VmhostGetServicesServiceId(ctx, vmhostId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **serviceId** | **string**| serviceId | 

### Return type

[**VmhostServices**](vmhost.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostGetTag**
> map[string]interface{} VmhostGetTag(ctx, vmhostId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostList**
> []Vmhost VmhostList(ctx, optional)
List

List vmhost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VmhostListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmhostListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabledServices** | **optional.String**| Filter by enabledServices | 

### Return type

[**[]Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostListAccessrights**
> []string VmhostListAccessrights(ctx, vmhostId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostListQueue**
> []Event VmhostListQueue(ctx, vmhostId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostListServices**
> []VmhostServices VmhostListServices(ctx, vmhostId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

[**[]VmhostServices**](vmhost.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostPatchTag**
> map[string]interface{} VmhostPatchTag(ctx, vmhostId, requestBody)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostPostAccessrights**
> Vmhost VmhostPostAccessrights(ctx, vmhostId, vmhostPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **vmhostPostAccessrights** | [**VmhostPostAccessrights**](VmhostPostAccessrights.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostShow**
> Vmhost VmhostShow(ctx, vmhostId)
Get

Returns a single vmhost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmhostUpdate**
> Vmhost VmhostUpdate(ctx, vmhostId, vmhostUpdate)
Update

Returns modified vmhost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **vmhostUpdate** | [**VmhostUpdate**](VmhostUpdate.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

