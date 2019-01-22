# \ContainerApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContainerActionRestart**](ContainerApi.md#ContainerActionRestart) | **Post** /container/{containerId}/actions/restart | /actions/restart
[**ContainerActionStart**](ContainerApi.md#ContainerActionStart) | **Post** /container/{containerId}/actions/start | /actions/start
[**ContainerActionStop**](ContainerApi.md#ContainerActionStop) | **Post** /container/{containerId}/actions/stop | /actions/stop
[**ContainerCreate**](ContainerApi.md#ContainerCreate) | **Post** /container | Create
[**ContainerDelete**](ContainerApi.md#ContainerDelete) | **Delete** /container/{containerId} | Delete
[**ContainerDeleteAccessrightsIdentity**](ContainerApi.md#ContainerDeleteAccessrightsIdentity) | **Delete** /container/{containerId}/accessrights/{identity} | /accessrights/:identity
[**ContainerDeleteTagKey**](ContainerApi.md#ContainerDeleteTagKey) | **Delete** /container/{containerId}/tag/{key} | /tag/:key
[**ContainerGetServicesServiceId**](ContainerApi.md#ContainerGetServicesServiceId) | **Get** /container/{containerId}/services/{serviceId} | /services/:serviceId
[**ContainerGetTag**](ContainerApi.md#ContainerGetTag) | **Get** /container/{containerId}/tag | /tag
[**ContainerList**](ContainerApi.md#ContainerList) | **Get** /container | List
[**ContainerListAccessrights**](ContainerApi.md#ContainerListAccessrights) | **Get** /container/{containerId}/accessrights | /accessrights
[**ContainerListQueue**](ContainerApi.md#ContainerListQueue) | **Get** /container/{containerId}/queue | /queue
[**ContainerListServices**](ContainerApi.md#ContainerListServices) | **Get** /container/{containerId}/services | /services
[**ContainerPatchTag**](ContainerApi.md#ContainerPatchTag) | **Patch** /container/{containerId}/tag | /tag
[**ContainerPostAccessrights**](ContainerApi.md#ContainerPostAccessrights) | **Post** /container/{containerId}/accessrights | /accessrights
[**ContainerShow**](ContainerApi.md#ContainerShow) | **Get** /container/{containerId} | Get
[**ContainerUpdate**](ContainerApi.md#ContainerUpdate) | **Patch** /container/{containerId} | Update


# **ContainerActionRestart**
> Container ContainerActionRestart(ctx, containerId)
/actions/restart

Action restart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerActionStart**
> Container ContainerActionStart(ctx, containerId)
/actions/start

Action start

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerActionStop**
> Container ContainerActionStop(ctx, containerId)
/actions/stop

Action stop

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerCreate**
> Container ContainerCreate(ctx, containerCreate)
Create

Create container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerCreate** | [**ContainerCreate**](ContainerCreate.md)|  | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerDelete**
> ContainerDelete(ctx, containerId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerDeleteAccessrightsIdentity**
> Container ContainerDeleteAccessrightsIdentity(ctx, containerId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **identity** | **string**| identity | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerDeleteTagKey**
> map[string]interface{} ContainerDeleteTagKey(ctx, containerId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerGetServicesServiceId**
> ContainerServices ContainerGetServicesServiceId(ctx, containerId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **serviceId** | **string**| serviceId | 

### Return type

[**ContainerServices**](container.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerGetTag**
> map[string]interface{} ContainerGetTag(ctx, containerId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerList**
> []Container ContainerList(ctx, optional)
List

List container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainerListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainerListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerListAccessrights**
> []string ContainerListAccessrights(ctx, containerId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerListQueue**
> []Event ContainerListQueue(ctx, containerId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerListServices**
> []ContainerServices ContainerListServices(ctx, containerId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

[**[]ContainerServices**](container.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerPatchTag**
> map[string]interface{} ContainerPatchTag(ctx, containerId, body)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **body** | **map[string]interface{}**|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerPostAccessrights**
> Container ContainerPostAccessrights(ctx, containerId, containerPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **containerPostAccessrights** | [**ContainerPostAccessrights**](ContainerPostAccessrights.md)|  | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerShow**
> Container ContainerShow(ctx, containerId)
Get

Returns a single container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerUpdate**
> Container ContainerUpdate(ctx, containerId, containerUpdate)
Update

Returns modified container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **containerUpdate** | [**ContainerUpdate**](ContainerUpdate.md)|  | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

