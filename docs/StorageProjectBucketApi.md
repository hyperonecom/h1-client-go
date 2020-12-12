# \StorageProjectBucketApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectBucketGet**](StorageProjectBucketApi.md#StorageProjectBucketGet) | **Get** /storage/{locationId}/project/{projectId}/bucket/{bucketId} | Get storage/bucket
[**StorageProjectBucketList**](StorageProjectBucketApi.md#StorageProjectBucketList) | **Get** /storage/{locationId}/project/{projectId}/bucket | List storage/bucket
[**StorageProjectBucketObjectDelete**](StorageProjectBucketApi.md#StorageProjectBucketObjectDelete) | **Delete** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object/{objectId} | Delete storage/bucket.object
[**StorageProjectBucketObjectDownload**](StorageProjectBucketApi.md#StorageProjectBucketObjectDownload) | **Post** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object/{objectId}/actions/download | Download storage/bucket.object
[**StorageProjectBucketObjectGet**](StorageProjectBucketApi.md#StorageProjectBucketObjectGet) | **Get** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object/{objectId} | Get storage/bucket.object
[**StorageProjectBucketObjectList**](StorageProjectBucketApi.md#StorageProjectBucketObjectList) | **Get** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object | List storage/bucket.object
[**StorageProjectBucketUpload**](StorageProjectBucketApi.md#StorageProjectBucketUpload) | **Post** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/actions/upload | Upload storage/bucket



## StorageProjectBucketGet

> Bucket StorageProjectBucketGet(ctx, projectId, locationId, bucketId)

Get storage/bucket

Returns a single bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**bucketId** | **string**| Bucket Id | 

### Return type

[**Bucket**](bucket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketList

> []Bucket StorageProjectBucketList(ctx, projectId, locationId)

List storage/bucket

List bucket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 

### Return type

[**[]Bucket**](bucket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketObjectDelete

> StorageProjectBucketObjectDelete(ctx, projectId, locationId, bucketId, objectId)

Delete storage/bucket.object

Delete storage/bucket.object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**bucketId** | **string**| Bucket Id | 
**objectId** | **string**| objectId | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketObjectDownload

> StorageProjectBucketObjectDownload(ctx, projectId, locationId, bucketId, objectId)

Download storage/bucket.object

action download

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**bucketId** | **string**| Bucket Id | 
**objectId** | **string**| objectId | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketObjectGet

> StorageObject StorageProjectBucketObjectGet(ctx, projectId, locationId, bucketId, objectId)

Get storage/bucket.object

Get storage/bucket.object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**bucketId** | **string**| Bucket Id | 
**objectId** | **string**| objectId | 

### Return type

[**StorageObject**](storage.object.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketObjectList

> []StorageObject StorageProjectBucketObjectList(ctx, projectId, locationId, bucketId)

List storage/bucket.object

List storage/bucket.object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**bucketId** | **string**| Bucket Id | 

### Return type

[**[]StorageObject**](storage.object.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketUpload

> Bucket StorageProjectBucketUpload(ctx, projectId, locationId, bucketId, storageProjectBucketUpload, optional)

Upload storage/bucket

action upload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**bucketId** | **string**| Bucket Id | 
**storageProjectBucketUpload** | [**StorageProjectBucketUpload**](StorageProjectBucketUpload.md)|  | 
 **optional** | ***StorageProjectBucketUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectBucketUploadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Bucket**](bucket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

