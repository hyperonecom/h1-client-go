# \StorageProjectImageApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectImageCreate**](StorageProjectImageApi.md#StorageProjectImageCreate) | **Post** /storage/{locationId}/project/{projectId}/image | Create storage/image
[**StorageProjectImageDelete**](StorageProjectImageApi.md#StorageProjectImageDelete) | **Delete** /storage/{locationId}/project/{projectId}/image/{imageId} | Delete storage/image
[**StorageProjectImageDiskList**](StorageProjectImageApi.md#StorageProjectImageDiskList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/disk | List storage/image.disk
[**StorageProjectImageEventGet**](StorageProjectImageApi.md#StorageProjectImageEventGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/event/{eventId} | Get storage/image.event
[**StorageProjectImageEventList**](StorageProjectImageApi.md#StorageProjectImageEventList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/event | List storage/image.event
[**StorageProjectImageGet**](StorageProjectImageApi.md#StorageProjectImageGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId} | Get storage/image
[**StorageProjectImageList**](StorageProjectImageApi.md#StorageProjectImageList) | **Get** /storage/{locationId}/project/{projectId}/image | List storage/image
[**StorageProjectImageServiceGet**](StorageProjectImageApi.md#StorageProjectImageServiceGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/service/{serviceId} | Get storage/image.service
[**StorageProjectImageServiceList**](StorageProjectImageApi.md#StorageProjectImageServiceList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/service | List storage/image.service
[**StorageProjectImageTagCreate**](StorageProjectImageApi.md#StorageProjectImageTagCreate) | **Post** /storage/{locationId}/project/{projectId}/image/{imageId}/tag | Create storage/image.tag
[**StorageProjectImageTagDelete**](StorageProjectImageApi.md#StorageProjectImageTagDelete) | **Delete** /storage/{locationId}/project/{projectId}/image/{imageId}/tag/{tagId} | Delete storage/image.tag
[**StorageProjectImageTagGet**](StorageProjectImageApi.md#StorageProjectImageTagGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/tag/{tagId} | Get storage/image.tag
[**StorageProjectImageTagList**](StorageProjectImageApi.md#StorageProjectImageTagList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/tag | List storage/image.tag
[**StorageProjectImageTagPut**](StorageProjectImageApi.md#StorageProjectImageTagPut) | **Put** /storage/{locationId}/project/{projectId}/image/{imageId}/tag | Replace storage/image.tag
[**StorageProjectImageTransfer**](StorageProjectImageApi.md#StorageProjectImageTransfer) | **Post** /storage/{locationId}/project/{projectId}/image/{imageId}/actions/transfer | Transfer storage/image
[**StorageProjectImageUpdate**](StorageProjectImageApi.md#StorageProjectImageUpdate) | **Patch** /storage/{locationId}/project/{projectId}/image/{imageId} | Update storage/image



## StorageProjectImageCreate

> Image StorageProjectImageCreate(ctx, projectId, locationId, storageProjectImageCreate, optional)

Create storage/image

Create image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**storageProjectImageCreate** | [**StorageProjectImageCreate**](StorageProjectImageCreate.md)|  | 
 **optional** | ***StorageProjectImageCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectImageCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Image**](image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageDelete

> StorageProjectImageDelete(ctx, projectId, locationId, imageId)

Delete storage/image

Delete image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 

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


## StorageProjectImageDiskList

> []Disk StorageProjectImageDiskList(ctx, projectId, locationId, imageId)

List storage/image.disk

List storage/image.disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 

### Return type

[**[]Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageEventGet

> Event StorageProjectImageEventGet(ctx, projectId, locationId, imageId, eventId)

Get storage/image.event

Get storage/image.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**eventId** | **string**| eventId | 

### Return type

[**Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageEventList

> []Event StorageProjectImageEventList(ctx, projectId, locationId, imageId, optional)

List storage/image.event

List storage/image.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
 **optional** | ***StorageProjectImageEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectImageEventListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Float32**| $limit | [default to 100]
 **skip** | **optional.Float32**| $skip | 

### Return type

[**[]Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageGet

> Image StorageProjectImageGet(ctx, projectId, locationId, imageId)

Get storage/image

Returns a single image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 

### Return type

[**Image**](image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageList

> []Image StorageProjectImageList(ctx, projectId, locationId, optional)

List storage/image

List image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***StorageProjectImageListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectImageListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Image**](image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageServiceGet

> ResourceService StorageProjectImageServiceGet(ctx, projectId, locationId, imageId, serviceId)

Get storage/image.service

Get storage/image.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**serviceId** | **string**| serviceId | 

### Return type

[**ResourceService**](resourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageServiceList

> []ResourceService StorageProjectImageServiceList(ctx, projectId, locationId, imageId)

List storage/image.service

List storage/image.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 

### Return type

[**[]ResourceService**](resourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageTagCreate

> Tag StorageProjectImageTagCreate(ctx, projectId, locationId, imageId, tag)

Create storage/image.tag

Create storage/image.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**tag** | [**Tag**](Tag.md)|  | 

### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageTagDelete

> StorageProjectImageTagDelete(ctx, projectId, locationId, imageId, tagId)

Delete storage/image.tag

Delete storage/image.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**tagId** | **string**| tagId | 

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


## StorageProjectImageTagGet

> Tag StorageProjectImageTagGet(ctx, projectId, locationId, imageId, tagId)

Get storage/image.tag

Get storage/image.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**tagId** | **string**| tagId | 

### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageTagList

> []Tag StorageProjectImageTagList(ctx, projectId, locationId, imageId)

List storage/image.tag

List storage/image.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 

### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageTagPut

> []Tag StorageProjectImageTagPut(ctx, projectId, locationId, imageId, tag)

Replace storage/image.tag

Replace storage/image.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**tag** | [**[]Tag**](tag.md)|  | 

### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageTransfer

> Image StorageProjectImageTransfer(ctx, projectId, locationId, imageId, storageProjectImageTransfer, optional)

Transfer storage/image

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**storageProjectImageTransfer** | [**StorageProjectImageTransfer**](StorageProjectImageTransfer.md)|  | 
 **optional** | ***StorageProjectImageTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectImageTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Image**](image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageUpdate

> Image StorageProjectImageUpdate(ctx, projectId, locationId, imageId, storageProjectImageUpdate)

Update storage/image

Returns modified image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**imageId** | **string**| Image Id | 
**storageProjectImageUpdate** | [**StorageProjectImageUpdate**](StorageProjectImageUpdate.md)|  | 

### Return type

[**Image**](image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

