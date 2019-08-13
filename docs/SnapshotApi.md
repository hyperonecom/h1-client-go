# \SnapshotApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnapshotCreate**](SnapshotApi.md#SnapshotCreate) | **Post** /snapshot | Create
[**SnapshotDelete**](SnapshotApi.md#SnapshotDelete) | **Delete** /snapshot/{snapshotId} | Delete
[**SnapshotDeleteAccessrightsIdentity**](SnapshotApi.md#SnapshotDeleteAccessrightsIdentity) | **Delete** /snapshot/{snapshotId}/accessrights/{identity} | /accessrights/:identity
[**SnapshotDeleteTagKey**](SnapshotApi.md#SnapshotDeleteTagKey) | **Delete** /snapshot/{snapshotId}/tag/{key} | /tag/:key
[**SnapshotGetServicesServiceId**](SnapshotApi.md#SnapshotGetServicesServiceId) | **Get** /snapshot/{snapshotId}/services/{serviceId} | /services/:serviceId
[**SnapshotGetTag**](SnapshotApi.md#SnapshotGetTag) | **Get** /snapshot/{snapshotId}/tag | /tag
[**SnapshotList**](SnapshotApi.md#SnapshotList) | **Get** /snapshot | List
[**SnapshotListAccessrights**](SnapshotApi.md#SnapshotListAccessrights) | **Get** /snapshot/{snapshotId}/accessrights | /accessrights
[**SnapshotListQueue**](SnapshotApi.md#SnapshotListQueue) | **Get** /snapshot/{snapshotId}/queue | /queue
[**SnapshotListServices**](SnapshotApi.md#SnapshotListServices) | **Get** /snapshot/{snapshotId}/services | /services
[**SnapshotPatchTag**](SnapshotApi.md#SnapshotPatchTag) | **Patch** /snapshot/{snapshotId}/tag | /tag
[**SnapshotPostAccessrights**](SnapshotApi.md#SnapshotPostAccessrights) | **Post** /snapshot/{snapshotId}/accessrights | /accessrights
[**SnapshotShow**](SnapshotApi.md#SnapshotShow) | **Get** /snapshot/{snapshotId} | Get
[**SnapshotUpdate**](SnapshotApi.md#SnapshotUpdate) | **Patch** /snapshot/{snapshotId} | Update



## SnapshotCreate

> Snapshot SnapshotCreate(ctx, snapshotCreate)
Create

Create snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotCreate** | [**SnapshotCreate**](SnapshotCreate.md)|  | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotDelete

> SnapshotDelete(ctx, snapshotId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 

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


## SnapshotDeleteAccessrightsIdentity

> Snapshot SnapshotDeleteAccessrightsIdentity(ctx, snapshotId, identity)
/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 
**identity** | **string**| identity | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotDeleteTagKey

> map[string]string SnapshotDeleteTagKey(ctx, snapshotId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 
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


## SnapshotGetServicesServiceId

> SnapshotServices SnapshotGetServicesServiceId(ctx, snapshotId, serviceId)
/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 
**serviceId** | **string**| serviceId | 

### Return type

[**SnapshotServices**](snapshot.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotGetTag

> map[string]string SnapshotGetTag(ctx, snapshotId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 

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


## SnapshotList

> []Snapshot SnapshotList(ctx, optional)
List

List snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **vault** | **optional.String**| Filter by vault | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotListAccessrights

> []string SnapshotListAccessrights(ctx, snapshotId)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 

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


## SnapshotListQueue

> []Event SnapshotListQueue(ctx, snapshotId)
/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 

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


## SnapshotListServices

> []SnapshotServices SnapshotListServices(ctx, snapshotId)
/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 

### Return type

[**[]SnapshotServices**](snapshot.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotPatchTag

> map[string]string SnapshotPatchTag(ctx, snapshotId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 
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


## SnapshotPostAccessrights

> Snapshot SnapshotPostAccessrights(ctx, snapshotId, snapshotPostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 
**snapshotPostAccessrights** | [**SnapshotPostAccessrights**](SnapshotPostAccessrights.md)|  | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotShow

> Snapshot SnapshotShow(ctx, snapshotId)
Get

Returns a single snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotUpdate

> Snapshot SnapshotUpdate(ctx, snapshotId, snapshotUpdate)
Update

Returns modified snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string**| ID of snapshot | 
**snapshotUpdate** | [**SnapshotUpdate**](SnapshotUpdate.md)|  | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

