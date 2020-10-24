# \StorageProjectDiskApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectDiskAttach**](StorageProjectDiskApi.md#StorageProjectDiskAttach) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/attach | Attach storage/disk
[**StorageProjectDiskCreate**](StorageProjectDiskApi.md#StorageProjectDiskCreate) | **Post** /storage/{locationId}/project/{projectId}/disk | Create storage/disk
[**StorageProjectDiskDelete**](StorageProjectDiskApi.md#StorageProjectDiskDelete) | **Delete** /storage/{locationId}/project/{projectId}/disk/{diskId} | Delete storage/disk
[**StorageProjectDiskDetach**](StorageProjectDiskApi.md#StorageProjectDiskDetach) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/detach | Detach storage/disk
[**StorageProjectDiskDownload**](StorageProjectDiskApi.md#StorageProjectDiskDownload) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/download | Download storage/disk
[**StorageProjectDiskEventGet**](StorageProjectDiskApi.md#StorageProjectDiskEventGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/event/{eventId} | Get storage/disk.event
[**StorageProjectDiskEventList**](StorageProjectDiskApi.md#StorageProjectDiskEventList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/event | List storage/disk.event
[**StorageProjectDiskGet**](StorageProjectDiskApi.md#StorageProjectDiskGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId} | Get storage/disk
[**StorageProjectDiskList**](StorageProjectDiskApi.md#StorageProjectDiskList) | **Get** /storage/{locationId}/project/{projectId}/disk | List storage/disk
[**StorageProjectDiskMetricGet**](StorageProjectDiskApi.md#StorageProjectDiskMetricGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/metric/{metricId} | Get storage/disk.metric
[**StorageProjectDiskMetricList**](StorageProjectDiskApi.md#StorageProjectDiskMetricList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/metric | List storage/disk.metric
[**StorageProjectDiskMetricPointList**](StorageProjectDiskApi.md#StorageProjectDiskMetricPointList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/metric/{metricId}/point | List storage/disk.point
[**StorageProjectDiskResize**](StorageProjectDiskApi.md#StorageProjectDiskResize) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/resize | Resize storage/disk
[**StorageProjectDiskServiceGet**](StorageProjectDiskApi.md#StorageProjectDiskServiceGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/service/{serviceId} | Get storage/disk.service
[**StorageProjectDiskServiceList**](StorageProjectDiskApi.md#StorageProjectDiskServiceList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/service | List storage/disk.service
[**StorageProjectDiskTagCreate**](StorageProjectDiskApi.md#StorageProjectDiskTagCreate) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag | Create storage/disk.tag
[**StorageProjectDiskTagDelete**](StorageProjectDiskApi.md#StorageProjectDiskTagDelete) | **Delete** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag/{tagId} | Delete storage/disk.tag
[**StorageProjectDiskTagGet**](StorageProjectDiskApi.md#StorageProjectDiskTagGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag/{tagId} | Get storage/disk.tag
[**StorageProjectDiskTagList**](StorageProjectDiskApi.md#StorageProjectDiskTagList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag | List storage/disk.tag
[**StorageProjectDiskTagPut**](StorageProjectDiskApi.md#StorageProjectDiskTagPut) | **Put** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag | Replace storage/disk.tag
[**StorageProjectDiskTransfer**](StorageProjectDiskApi.md#StorageProjectDiskTransfer) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/transfer | Transfer storage/disk
[**StorageProjectDiskUpdate**](StorageProjectDiskApi.md#StorageProjectDiskUpdate) | **Patch** /storage/{locationId}/project/{projectId}/disk/{diskId} | Update storage/disk



## StorageProjectDiskAttach

> Disk StorageProjectDiskAttach(ctx, projectId, locationId, diskId, storageProjectDiskAttach, optional)

Attach storage/disk

action attach

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
**storageProjectDiskAttach** | [**StorageProjectDiskAttach**](StorageProjectDiskAttach.md)|  | 
 **optional** | ***StorageProjectDiskAttachOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskAttachOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskCreate

> Disk StorageProjectDiskCreate(ctx, projectId, locationId, storageProjectDiskCreate, optional)

Create storage/disk

Create disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**storageProjectDiskCreate** | [**StorageProjectDiskCreate**](StorageProjectDiskCreate.md)|  | 
 **optional** | ***StorageProjectDiskCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskDelete

> StorageProjectDiskDelete(ctx, projectId, locationId, diskId)

Delete storage/disk

Delete disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 

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


## StorageProjectDiskDetach

> Disk StorageProjectDiskDetach(ctx, projectId, locationId, diskId, optional)

Detach storage/disk

action detach

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
 **optional** | ***StorageProjectDiskDetachOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskDetachOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskDownload

> StorageProjectDiskDownload(ctx, projectId, locationId, diskId, optional)

Download storage/disk

action download

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
 **optional** | ***StorageProjectDiskDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskDownloadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

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


## StorageProjectDiskEventGet

> Event StorageProjectDiskEventGet(ctx, projectId, locationId, diskId, eventId)

Get storage/disk.event

Get storage/disk.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
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


## StorageProjectDiskEventList

> []Event StorageProjectDiskEventList(ctx, projectId, locationId, diskId, optional)

List storage/disk.event

List storage/disk.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
 **optional** | ***StorageProjectDiskEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskEventListOpts struct


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


## StorageProjectDiskGet

> Disk StorageProjectDiskGet(ctx, projectId, locationId, diskId)

Get storage/disk

Returns a single disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 

### Return type

[**Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskList

> []Disk StorageProjectDiskList(ctx, projectId, locationId, optional)

List storage/disk

List disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***StorageProjectDiskListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **vm** | **optional.String**| Filter by vm | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

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


## StorageProjectDiskMetricGet

> Metric StorageProjectDiskMetricGet(ctx, projectId, locationId, diskId, metricId)

Get storage/disk.metric

Get storage/disk.metric

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
**metricId** | **string**| metricId | 

### Return type

[**Metric**](metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskMetricList

> []Metric StorageProjectDiskMetricList(ctx, projectId, locationId, diskId)

List storage/disk.metric

List storage/disk.metric

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 

### Return type

[**[]Metric**](metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskMetricPointList

> []Point StorageProjectDiskMetricPointList(ctx, projectId, locationId, diskId, metricId, optional)

List storage/disk.point

List storage/disk.point

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
**metricId** | **string**| metricId | 
 **optional** | ***StorageProjectDiskMetricPointListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskMetricPointListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **interval** | **optional.String**| interval | 
 **timespan** | **optional.String**| timespan | 

### Return type

[**[]Point**](point.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskResize

> Disk StorageProjectDiskResize(ctx, projectId, locationId, diskId, storageProjectDiskResize, optional)

Resize storage/disk

action resize

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
**storageProjectDiskResize** | [**StorageProjectDiskResize**](StorageProjectDiskResize.md)|  | 
 **optional** | ***StorageProjectDiskResizeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskResizeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskServiceGet

> ResourceService StorageProjectDiskServiceGet(ctx, projectId, locationId, diskId, serviceId)

Get storage/disk.service

Get storage/disk.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
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


## StorageProjectDiskServiceList

> []ResourceService StorageProjectDiskServiceList(ctx, projectId, locationId, diskId)

List storage/disk.service

List storage/disk.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 

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


## StorageProjectDiskTagCreate

> Tag StorageProjectDiskTagCreate(ctx, projectId, locationId, diskId, tag)

Create storage/disk.tag

Create storage/disk.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
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


## StorageProjectDiskTagDelete

> StorageProjectDiskTagDelete(ctx, projectId, locationId, diskId, tagId)

Delete storage/disk.tag

Delete storage/disk.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
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


## StorageProjectDiskTagGet

> Tag StorageProjectDiskTagGet(ctx, projectId, locationId, diskId, tagId)

Get storage/disk.tag

Get storage/disk.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
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


## StorageProjectDiskTagList

> []Tag StorageProjectDiskTagList(ctx, projectId, locationId, diskId)

List storage/disk.tag

List storage/disk.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 

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


## StorageProjectDiskTagPut

> []Tag StorageProjectDiskTagPut(ctx, projectId, locationId, diskId, tag)

Replace storage/disk.tag

Replace storage/disk.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
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


## StorageProjectDiskTransfer

> Disk StorageProjectDiskTransfer(ctx, projectId, locationId, diskId, storageProjectDiskTransfer, optional)

Transfer storage/disk

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
**storageProjectDiskTransfer** | [**StorageProjectDiskTransfer**](StorageProjectDiskTransfer.md)|  | 
 **optional** | ***StorageProjectDiskTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectDiskTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskUpdate

> Disk StorageProjectDiskUpdate(ctx, projectId, locationId, diskId, storageProjectDiskUpdate)

Update storage/disk

Returns modified disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**diskId** | **string**| Disk Id | 
**storageProjectDiskUpdate** | [**StorageProjectDiskUpdate**](StorageProjectDiskUpdate.md)|  | 

### Return type

[**Disk**](disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

