# \ContainerApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionContainerRestart**](ContainerApi.md#ActionContainerRestart) | **Post** /container/{containerId}/actions/restart | /actions/restart
[**ActionContainerStart**](ContainerApi.md#ActionContainerStart) | **Post** /container/{containerId}/actions/start | /actions/start
[**ActionContainerStop**](ContainerApi.md#ActionContainerStop) | **Post** /container/{containerId}/actions/stop | /actions/stop
[**CreateContainer**](ContainerApi.md#CreateContainer) | **Post** /container | Create
[**DeleteContainer**](ContainerApi.md#DeleteContainer) | **Delete** /container/{containerId} | Delete
[**ListContainer**](ContainerApi.md#ListContainer) | **Get** /container | List
[**OperationContainerDeleteaccessrightsIdentity**](ContainerApi.md#OperationContainerDeleteaccessrightsIdentity) | **Delete** /container/{containerId}/accessrights/{identity} | /accessrights/:identity
[**OperationContainerDeletetagKey**](ContainerApi.md#OperationContainerDeletetagKey) | **Delete** /container/{containerId}/tag/{key} | /tag/:key
[**OperationContainerGetservicesServiceId**](ContainerApi.md#OperationContainerGetservicesServiceId) | **Get** /container/{containerId}/services/{serviceId} | /services/:serviceId
[**OperationContainerGettag**](ContainerApi.md#OperationContainerGettag) | **Get** /container/{containerId}/tag/ | /tag/
[**OperationContainerListaccessrights**](ContainerApi.md#OperationContainerListaccessrights) | **Get** /container/{containerId}/accessrights/ | /accessrights/
[**OperationContainerListqueue**](ContainerApi.md#OperationContainerListqueue) | **Get** /container/{containerId}/queue/ | /queue/
[**OperationContainerListservices**](ContainerApi.md#OperationContainerListservices) | **Get** /container/{containerId}/services/ | /services/
[**OperationContainerPatchtag**](ContainerApi.md#OperationContainerPatchtag) | **Patch** /container/{containerId}/tag/ | /tag/
[**OperationContainerPostaccessrights**](ContainerApi.md#OperationContainerPostaccessrights) | **Post** /container/{containerId}/accessrights/ | /accessrights/
[**ShowContainer**](ContainerApi.md#ShowContainer) | **Get** /container/{containerId} | Get
[**UpdateContainer**](ContainerApi.md#UpdateContainer) | **Patch** /container/{containerId} | Update


# **ActionContainerRestart**
> Container ActionContainerRestart(ctx, containerId)
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

# **ActionContainerStart**
> Container ActionContainerStart(ctx, containerId)
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

# **ActionContainerStop**
> Container ActionContainerStop(ctx, containerId)
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

# **CreateContainer**
> Container CreateContainer(ctx, optional)
Create

Create container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateContainerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateContainerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject7** | [**optional.Interface of InlineObject7**](InlineObject7.md)|  | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContainer**
> DeleteContainer(ctx, containerId)
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

# **ListContainer**
> []Container ListContainer(ctx, optional)
List

List container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListContainerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListContainerOpts struct

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

# **OperationContainerDeleteaccessrightsIdentity**
> Container OperationContainerDeleteaccessrightsIdentity(ctx, containerId, identity)
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

# **OperationContainerDeletetagKey**
> map[string]string OperationContainerDeletetagKey(ctx, containerId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationContainerGetservicesServiceId**
> ContainerServices OperationContainerGetservicesServiceId(ctx, containerId, serviceId)
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

# **OperationContainerGettag**
> map[string]string OperationContainerGettag(ctx, containerId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationContainerListaccessrights**
> []string OperationContainerListaccessrights(ctx, containerId)
/accessrights/

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

# **OperationContainerListqueue**
> []Event OperationContainerListqueue(ctx, containerId)
/queue/

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

# **OperationContainerListservices**
> []ContainerServices OperationContainerListservices(ctx, containerId)
/services/

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

# **OperationContainerPatchtag**
> map[string]string OperationContainerPatchtag(ctx, containerId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationContainerPostaccessrights**
> string OperationContainerPostaccessrights(ctx, containerId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
 **optional** | ***OperationContainerPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationContainerPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject9** | [**optional.Interface of InlineObject9**](InlineObject9.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowContainer**
> Container ShowContainer(ctx, containerId)
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

# **UpdateContainer**
> Container UpdateContainer(ctx, containerId, optional)
Update

Returns modified container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerId** | **string**| ID of container | 
 **optional** | ***UpdateContainerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateContainerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject8** | [**optional.Interface of InlineObject8**](InlineObject8.md)|  | 

### Return type

[**Container**](container.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

