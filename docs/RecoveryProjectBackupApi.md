# \RecoveryProjectBackupApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecoveryProjectBackupCreate**](RecoveryProjectBackupApi.md#RecoveryProjectBackupCreate) | **Post** /recovery/{locationId}/project/{projectId}/backup | Create recovery/backup
[**RecoveryProjectBackupDelete**](RecoveryProjectBackupApi.md#RecoveryProjectBackupDelete) | **Delete** /recovery/{locationId}/project/{projectId}/backup/{backupId} | Delete recovery/backup
[**RecoveryProjectBackupEventGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupEventGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/event/{eventId} | Get recovery/backup.event
[**RecoveryProjectBackupEventList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupEventList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/event | List recovery/backup.event
[**RecoveryProjectBackupGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId} | Get recovery/backup
[**RecoveryProjectBackupList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupList) | **Get** /recovery/{locationId}/project/{projectId}/backup | List recovery/backup
[**RecoveryProjectBackupServiceGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupServiceGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/service/{serviceId} | Get recovery/backup.service
[**RecoveryProjectBackupServiceList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupServiceList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/service | List recovery/backup.service
[**RecoveryProjectBackupTagCreate**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagCreate) | **Post** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag | Create recovery/backup.tag
[**RecoveryProjectBackupTagDelete**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagDelete) | **Delete** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag/{tagId} | Delete recovery/backup.tag
[**RecoveryProjectBackupTagGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag/{tagId} | Get recovery/backup.tag
[**RecoveryProjectBackupTagList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag | List recovery/backup.tag
[**RecoveryProjectBackupTagPut**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagPut) | **Put** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag | Replace recovery/backup.tag
[**RecoveryProjectBackupUpdate**](RecoveryProjectBackupApi.md#RecoveryProjectBackupUpdate) | **Patch** /recovery/{locationId}/project/{projectId}/backup/{backupId} | Update recovery/backup



## RecoveryProjectBackupCreate

> Backup RecoveryProjectBackupCreate(ctx, projectId, locationId, recoveryProjectBackupCreate, optional)

Create recovery/backup

Create backup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**recoveryProjectBackupCreate** | [**RecoveryProjectBackupCreate**](RecoveryProjectBackupCreate.md)|  | 
 **optional** | ***RecoveryProjectBackupCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RecoveryProjectBackupCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupDelete

> RecoveryProjectBackupDelete(ctx, projectId, locationId, backupId)

Delete recovery/backup

Delete backup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 

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


## RecoveryProjectBackupEventGet

> Event RecoveryProjectBackupEventGet(ctx, projectId, locationId, backupId, eventId)

Get recovery/backup.event

Get recovery/backup.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
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


## RecoveryProjectBackupEventList

> []Event RecoveryProjectBackupEventList(ctx, projectId, locationId, backupId, optional)

List recovery/backup.event

List recovery/backup.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
 **optional** | ***RecoveryProjectBackupEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RecoveryProjectBackupEventListOpts struct


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


## RecoveryProjectBackupGet

> Backup RecoveryProjectBackupGet(ctx, projectId, locationId, backupId)

Get recovery/backup

Returns a single backup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 

### Return type

[**Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupList

> []Backup RecoveryProjectBackupList(ctx, projectId, locationId, optional)

List recovery/backup

List backup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***RecoveryProjectBackupListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RecoveryProjectBackupListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **source** | **optional.String**| Filter by source | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupServiceGet

> ResourceService RecoveryProjectBackupServiceGet(ctx, projectId, locationId, backupId, serviceId)

Get recovery/backup.service

Get recovery/backup.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
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


## RecoveryProjectBackupServiceList

> []ResourceService RecoveryProjectBackupServiceList(ctx, projectId, locationId, backupId)

List recovery/backup.service

List recovery/backup.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 

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


## RecoveryProjectBackupTagCreate

> Tag RecoveryProjectBackupTagCreate(ctx, projectId, locationId, backupId, tag)

Create recovery/backup.tag

Create recovery/backup.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
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


## RecoveryProjectBackupTagDelete

> RecoveryProjectBackupTagDelete(ctx, projectId, locationId, backupId, tagId)

Delete recovery/backup.tag

Delete recovery/backup.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
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


## RecoveryProjectBackupTagGet

> Tag RecoveryProjectBackupTagGet(ctx, projectId, locationId, backupId, tagId)

Get recovery/backup.tag

Get recovery/backup.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
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


## RecoveryProjectBackupTagList

> []Tag RecoveryProjectBackupTagList(ctx, projectId, locationId, backupId)

List recovery/backup.tag

List recovery/backup.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 

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


## RecoveryProjectBackupTagPut

> []Tag RecoveryProjectBackupTagPut(ctx, projectId, locationId, backupId, tag)

Replace recovery/backup.tag

Replace recovery/backup.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
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


## RecoveryProjectBackupUpdate

> Backup RecoveryProjectBackupUpdate(ctx, projectId, locationId, backupId, recoveryProjectBackupUpdate)

Update recovery/backup

Returns modified backup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**backupId** | **string**| Backup Id | 
**recoveryProjectBackupUpdate** | [**RecoveryProjectBackupUpdate**](RecoveryProjectBackupUpdate.md)|  | 

### Return type

[**Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

