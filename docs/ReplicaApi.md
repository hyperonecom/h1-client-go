# \ReplicaApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplicaActionImage**](ReplicaApi.md#ReplicaActionImage) | **Post** /replica/{replicaId}/actions/image | /actions/image
[**ReplicaCreate**](ReplicaApi.md#ReplicaCreate) | **Post** /replica | Create
[**ReplicaDelete**](ReplicaApi.md#ReplicaDelete) | **Delete** /replica/{replicaId} | Delete
[**ReplicaDeleteAccessrightsIdentity**](ReplicaApi.md#ReplicaDeleteAccessrightsIdentity) | **Delete** /replica/{replicaId}/accessrights/{identity} | /accessrights/:identity
[**ReplicaDeleteTagKey**](ReplicaApi.md#ReplicaDeleteTagKey) | **Delete** /replica/{replicaId}/tag/{key} | /tag/:key
[**ReplicaGetServicesServiceId**](ReplicaApi.md#ReplicaGetServicesServiceId) | **Get** /replica/{replicaId}/services/{serviceId} | /services/:serviceId
[**ReplicaGetTag**](ReplicaApi.md#ReplicaGetTag) | **Get** /replica/{replicaId}/tag | /tag
[**ReplicaList**](ReplicaApi.md#ReplicaList) | **Get** /replica | List
[**ReplicaListAccessrights**](ReplicaApi.md#ReplicaListAccessrights) | **Get** /replica/{replicaId}/accessrights | /accessrights
[**ReplicaListQueue**](ReplicaApi.md#ReplicaListQueue) | **Get** /replica/{replicaId}/queue | /queue
[**ReplicaListServices**](ReplicaApi.md#ReplicaListServices) | **Get** /replica/{replicaId}/services | /services
[**ReplicaPatchTag**](ReplicaApi.md#ReplicaPatchTag) | **Patch** /replica/{replicaId}/tag | /tag
[**ReplicaPostAccessrights**](ReplicaApi.md#ReplicaPostAccessrights) | **Post** /replica/{replicaId}/accessrights | /accessrights
[**ReplicaShow**](ReplicaApi.md#ReplicaShow) | **Get** /replica/{replicaId} | Get



## ReplicaActionImage

> Replica ReplicaActionImage(ctx, replicaId, replicaActionImage)
/actions/image

Action image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 
**replicaActionImage** | [**ReplicaActionImage**](ReplicaActionImage.md)|  | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaCreate

> Replica ReplicaCreate(ctx, replicaCreate)
Create

Create replica

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaCreate** | [**ReplicaCreate**](ReplicaCreate.md)|  | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaDelete

> ReplicaDelete(ctx, replicaId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaDeleteAccessrightsIdentity

> Replica ReplicaDeleteAccessrightsIdentity(ctx, replicaId, identity)
/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 
**identity** | **string**| identity | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaDeleteTagKey

> map[string]interface{} ReplicaDeleteTagKey(ctx, replicaId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 
**key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaGetServicesServiceId

> ReplicaServices ReplicaGetServicesServiceId(ctx, replicaId, serviceId)
/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 
**serviceId** | **string**| serviceId | 

### Return type

[**ReplicaServices**](replica.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaGetTag

> map[string]interface{} ReplicaGetTag(ctx, replicaId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaList

> []Replica ReplicaList(ctx, optional)
List

List replica

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicaListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplicaListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaListAccessrights

> []string ReplicaListAccessrights(ctx, replicaId)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 

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


## ReplicaListQueue

> []Event ReplicaListQueue(ctx, replicaId)
/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 

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


## ReplicaListServices

> []ReplicaServices ReplicaListServices(ctx, replicaId)
/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 

### Return type

[**[]ReplicaServices**](replica.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaPatchTag

> map[string]interface{} ReplicaPatchTag(ctx, replicaId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaPostAccessrights

> Replica ReplicaPostAccessrights(ctx, replicaId, replicaPostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 
**replicaPostAccessrights** | [**ReplicaPostAccessrights**](ReplicaPostAccessrights.md)|  | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicaShow

> Replica ReplicaShow(ctx, replicaId)
Get

Returns a single replica

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaId** | **string**| ID of replica | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

