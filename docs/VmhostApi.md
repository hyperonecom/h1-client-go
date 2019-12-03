# \VmhostApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
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
[**VmhostPutTag**](VmhostApi.md#VmhostPutTag) | **Put** /vmhost/{vmhostId}/tag | /tag
[**VmhostShow**](VmhostApi.md#VmhostShow) | **Get** /vmhost/{vmhostId} | Get



## VmhostDeleteAccessrightsIdentity

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostDeleteTagKey

> map[string]string VmhostDeleteTagKey(ctx, vmhostId, key)

/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmhostId** | **string**| ID of vmhost | 
**key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostGetServicesServiceId

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostGetTag

> map[string]string VmhostGetTag(ctx, vmhostId)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmhostId** | **string**| ID of vmhost | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostList

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostListAccessrights

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostListQueue

> []Event VmhostListQueue(ctx, vmhostId, optional)

/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmhostId** | **string**| ID of vmhost | 
 **optional** | ***VmhostListQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VmhostListQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Float32**| $limit | 
 **skip** | **optional.Float32**| $skip | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostListServices

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostPatchTag

> map[string]string VmhostPatchTag(ctx, vmhostId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmhostId** | **string**| ID of vmhost | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostPostAccessrights

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostPutTag

> map[string]string VmhostPutTag(ctx, vmhostId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmhostId** | **string**| ID of vmhost | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostShow

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

