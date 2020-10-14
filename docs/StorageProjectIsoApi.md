# \StorageProjectIsoApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectIsoCreate**](StorageProjectIsoApi.md#StorageProjectIsoCreate) | **Post** /storage/{locationId}/project/{projectId}/iso | Create storage/iso
[**StorageProjectIsoDelete**](StorageProjectIsoApi.md#StorageProjectIsoDelete) | **Delete** /storage/{locationId}/project/{projectId}/iso/{isoId} | Delete storage/iso
[**StorageProjectIsoDetach**](StorageProjectIsoApi.md#StorageProjectIsoDetach) | **Post** /storage/{locationId}/project/{projectId}/iso/{isoId}/actions/detach | Detach storage/iso
[**StorageProjectIsoEventGet**](StorageProjectIsoApi.md#StorageProjectIsoEventGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/event/{eventId} | Get storage/iso.event
[**StorageProjectIsoEventList**](StorageProjectIsoApi.md#StorageProjectIsoEventList) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/event | List storage/iso.event
[**StorageProjectIsoGet**](StorageProjectIsoApi.md#StorageProjectIsoGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId} | Get storage/iso
[**StorageProjectIsoList**](StorageProjectIsoApi.md#StorageProjectIsoList) | **Get** /storage/{locationId}/project/{projectId}/iso | List storage/iso
[**StorageProjectIsoServiceGet**](StorageProjectIsoApi.md#StorageProjectIsoServiceGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/service/{serviceId} | Get storage/iso.service
[**StorageProjectIsoServiceList**](StorageProjectIsoApi.md#StorageProjectIsoServiceList) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/service | List storage/iso.service
[**StorageProjectIsoTagCreate**](StorageProjectIsoApi.md#StorageProjectIsoTagCreate) | **Post** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag | Create storage/iso.tag
[**StorageProjectIsoTagDelete**](StorageProjectIsoApi.md#StorageProjectIsoTagDelete) | **Delete** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag/{tagId} | Delete storage/iso.tag
[**StorageProjectIsoTagGet**](StorageProjectIsoApi.md#StorageProjectIsoTagGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag/{tagId} | Get storage/iso.tag
[**StorageProjectIsoTagList**](StorageProjectIsoApi.md#StorageProjectIsoTagList) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag | List storage/iso.tag
[**StorageProjectIsoTagPut**](StorageProjectIsoApi.md#StorageProjectIsoTagPut) | **Put** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag | Replace storage/iso.tag
[**StorageProjectIsoTransfer**](StorageProjectIsoApi.md#StorageProjectIsoTransfer) | **Post** /storage/{locationId}/project/{projectId}/iso/{isoId}/actions/transfer | Transfer storage/iso
[**StorageProjectIsoUpdate**](StorageProjectIsoApi.md#StorageProjectIsoUpdate) | **Patch** /storage/{locationId}/project/{projectId}/iso/{isoId} | Update storage/iso



## StorageProjectIsoCreate

> Iso StorageProjectIsoCreate(ctx, projectId, locationId, storageProjectIsoCreate, optional)

Create storage/iso

Create iso

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**storageProjectIsoCreate** | [**StorageProjectIsoCreate**](StorageProjectIsoCreate.md)|  | 
 **optional** | ***StorageProjectIsoCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectIsoCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Iso**](iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoDelete

> StorageProjectIsoDelete(ctx, projectId, locationId, isoId)

Delete storage/iso

Delete iso

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 

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


## StorageProjectIsoDetach

> Iso StorageProjectIsoDetach(ctx, projectId, locationId, isoId, storageProjectIsoDetach, optional)

Detach storage/iso

action detach

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
**storageProjectIsoDetach** | [**StorageProjectIsoDetach**](StorageProjectIsoDetach.md)|  | 
 **optional** | ***StorageProjectIsoDetachOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectIsoDetachOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Iso**](iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoEventGet

> Event StorageProjectIsoEventGet(ctx, projectId, locationId, isoId, eventId)

Get storage/iso.event

Get storage/iso.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
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


## StorageProjectIsoEventList

> []Event StorageProjectIsoEventList(ctx, projectId, locationId, isoId, optional)

List storage/iso.event

List storage/iso.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
 **optional** | ***StorageProjectIsoEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectIsoEventListOpts struct


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


## StorageProjectIsoGet

> Iso StorageProjectIsoGet(ctx, projectId, locationId, isoId)

Get storage/iso

Returns a single iso

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 

### Return type

[**Iso**](iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoList

> []Iso StorageProjectIsoList(ctx, projectId, locationId, optional)

List storage/iso

List iso

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***StorageProjectIsoListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectIsoListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Iso**](iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoServiceGet

> ResourceService StorageProjectIsoServiceGet(ctx, projectId, locationId, isoId, serviceId)

Get storage/iso.service

Get storage/iso.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
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


## StorageProjectIsoServiceList

> []ResourceService StorageProjectIsoServiceList(ctx, projectId, locationId, isoId)

List storage/iso.service

List storage/iso.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 

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


## StorageProjectIsoTagCreate

> Tag StorageProjectIsoTagCreate(ctx, projectId, locationId, isoId, tag)

Create storage/iso.tag

Create storage/iso.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
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


## StorageProjectIsoTagDelete

> StorageProjectIsoTagDelete(ctx, projectId, locationId, isoId, tagId)

Delete storage/iso.tag

Delete storage/iso.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
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


## StorageProjectIsoTagGet

> Tag StorageProjectIsoTagGet(ctx, projectId, locationId, isoId, tagId)

Get storage/iso.tag

Get storage/iso.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
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


## StorageProjectIsoTagList

> []Tag StorageProjectIsoTagList(ctx, projectId, locationId, isoId)

List storage/iso.tag

List storage/iso.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 

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


## StorageProjectIsoTagPut

> []Tag StorageProjectIsoTagPut(ctx, projectId, locationId, isoId, tag)

Replace storage/iso.tag

Replace storage/iso.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
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


## StorageProjectIsoTransfer

> Iso StorageProjectIsoTransfer(ctx, projectId, locationId, isoId, storageProjectIsoTransfer, optional)

Transfer storage/iso

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
**storageProjectIsoTransfer** | [**StorageProjectIsoTransfer**](StorageProjectIsoTransfer.md)|  | 
 **optional** | ***StorageProjectIsoTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectIsoTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Iso**](iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoUpdate

> Iso StorageProjectIsoUpdate(ctx, projectId, locationId, isoId, storageProjectIsoUpdate)

Update storage/iso

Returns modified iso

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**isoId** | **string**| Iso Id | 
**storageProjectIsoUpdate** | [**StorageProjectIsoUpdate**](StorageProjectIsoUpdate.md)|  | 

### Return type

[**Iso**](iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

