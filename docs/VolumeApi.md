# \VolumeApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumeCreate**](VolumeApi.md#VolumeCreate) | **Post** /volume | Create
[**VolumeDelete**](VolumeApi.md#VolumeDelete) | **Delete** /volume/{volumeId} | Delete
[**VolumeDeleteAccessrightsIdentity**](VolumeApi.md#VolumeDeleteAccessrightsIdentity) | **Delete** /volume/{volumeId}/accessrights/{identity} | /accessrights/:identity
[**VolumeDeleteTagKey**](VolumeApi.md#VolumeDeleteTagKey) | **Delete** /volume/{volumeId}/tag/{key} | /tag/:key
[**VolumeGetServicesServiceId**](VolumeApi.md#VolumeGetServicesServiceId) | **Get** /volume/{volumeId}/services/{serviceId} | /services/:serviceId
[**VolumeGetTag**](VolumeApi.md#VolumeGetTag) | **Get** /volume/{volumeId}/tag | /tag
[**VolumeList**](VolumeApi.md#VolumeList) | **Get** /volume | List
[**VolumeListAccessrights**](VolumeApi.md#VolumeListAccessrights) | **Get** /volume/{volumeId}/accessrights | /accessrights
[**VolumeListQueue**](VolumeApi.md#VolumeListQueue) | **Get** /volume/{volumeId}/queue | /queue
[**VolumeListServices**](VolumeApi.md#VolumeListServices) | **Get** /volume/{volumeId}/services | /services
[**VolumePatchTag**](VolumeApi.md#VolumePatchTag) | **Patch** /volume/{volumeId}/tag | /tag
[**VolumePostAccessrights**](VolumeApi.md#VolumePostAccessrights) | **Post** /volume/{volumeId}/accessrights | /accessrights
[**VolumeShow**](VolumeApi.md#VolumeShow) | **Get** /volume/{volumeId} | Get
[**VolumeUpdate**](VolumeApi.md#VolumeUpdate) | **Patch** /volume/{volumeId} | Update


# **VolumeCreate**
> Volume VolumeCreate(ctx, volumeCreate)
Create

Create volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeCreate** | [**VolumeCreate**](VolumeCreate.md)|  | 

### Return type

[**Volume**](volume.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeDelete**
> VolumeDelete(ctx, volumeId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeDeleteAccessrightsIdentity**
> Volume VolumeDeleteAccessrightsIdentity(ctx, volumeId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 
  **identity** | **string**| identity | 

### Return type

[**Volume**](volume.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeDeleteTagKey**
> map[string]interface{} VolumeDeleteTagKey(ctx, volumeId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeGetServicesServiceId**
> VolumeServices VolumeGetServicesServiceId(ctx, volumeId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 
  **serviceId** | **string**| serviceId | 

### Return type

[**VolumeServices**](volume.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeGetTag**
> map[string]interface{} VolumeGetTag(ctx, volumeId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeList**
> []Volume VolumeList(ctx, optional)
List

List volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Volume**](volume.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeListAccessrights**
> []string VolumeListAccessrights(ctx, volumeId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeListQueue**
> []Event VolumeListQueue(ctx, volumeId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeListServices**
> []VolumeServices VolumeListServices(ctx, volumeId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 

### Return type

[**[]VolumeServices**](volume.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumePatchTag**
> map[string]interface{} VolumePatchTag(ctx, volumeId, requestBody)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumePostAccessrights**
> Volume VolumePostAccessrights(ctx, volumeId, volumePostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 
  **volumePostAccessrights** | [**VolumePostAccessrights**](VolumePostAccessrights.md)|  | 

### Return type

[**Volume**](volume.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeShow**
> Volume VolumeShow(ctx, volumeId)
Get

Returns a single volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 

### Return type

[**Volume**](volume.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeUpdate**
> Volume VolumeUpdate(ctx, volumeId, volumeUpdate)
Update

Returns modified volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| ID of volume | 
  **volumeUpdate** | [**VolumeUpdate**](VolumeUpdate.md)|  | 

### Return type

[**Volume**](volume.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

