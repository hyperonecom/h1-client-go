# \NetworkingProjectNetgwApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectNetgwAttach**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwAttach) | **Post** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/actions/attach | Attach networking/netgw
[**NetworkingProjectNetgwCreate**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwCreate) | **Post** /networking/{locationId}/project/{projectId}/netgw | Create networking/netgw
[**NetworkingProjectNetgwDelete**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwDelete) | **Delete** /networking/{locationId}/project/{projectId}/netgw/{netgwId} | Delete networking/netgw
[**NetworkingProjectNetgwDetach**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwDetach) | **Post** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/actions/detach | Detach networking/netgw
[**NetworkingProjectNetgwEventGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwEventGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/event/{eventId} | Get networking/netgw.event
[**NetworkingProjectNetgwEventList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwEventList) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/event | List networking/netgw.event
[**NetworkingProjectNetgwGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId} | Get networking/netgw
[**NetworkingProjectNetgwList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwList) | **Get** /networking/{locationId}/project/{projectId}/netgw | List networking/netgw
[**NetworkingProjectNetgwServiceGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwServiceGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/service/{serviceId} | Get networking/netgw.service
[**NetworkingProjectNetgwServiceList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwServiceList) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/service | List networking/netgw.service
[**NetworkingProjectNetgwTagCreate**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagCreate) | **Post** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag | Create networking/netgw.tag
[**NetworkingProjectNetgwTagDelete**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag/{tagId} | Delete networking/netgw.tag
[**NetworkingProjectNetgwTagGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag/{tagId} | Get networking/netgw.tag
[**NetworkingProjectNetgwTagList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagList) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag | List networking/netgw.tag
[**NetworkingProjectNetgwTagPut**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagPut) | **Put** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag | Replace networking/netgw.tag
[**NetworkingProjectNetgwUpdate**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwUpdate) | **Patch** /networking/{locationId}/project/{projectId}/netgw/{netgwId} | Update networking/netgw



## NetworkingProjectNetgwAttach

> Netgw NetworkingProjectNetgwAttach(ctx, projectId, locationId, netgwId, networkingProjectNetgwAttach, optional)

Attach networking/netgw

action attach

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
**networkingProjectNetgwAttach** | [**NetworkingProjectNetgwAttach**](NetworkingProjectNetgwAttach.md)|  | 
 **optional** | ***NetworkingProjectNetgwAttachOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetgwAttachOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwCreate

> Netgw NetworkingProjectNetgwCreate(ctx, projectId, locationId, networkingProjectNetgwCreate, optional)

Create networking/netgw

Create netgw

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkingProjectNetgwCreate** | [**NetworkingProjectNetgwCreate**](NetworkingProjectNetgwCreate.md)|  | 
 **optional** | ***NetworkingProjectNetgwCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetgwCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwDelete

> NetworkingProjectNetgwDelete(ctx, projectId, locationId, netgwId)

Delete networking/netgw

Delete netgw

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 

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


## NetworkingProjectNetgwDetach

> Netgw NetworkingProjectNetgwDetach(ctx, projectId, locationId, netgwId, optional)

Detach networking/netgw

action detach

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
 **optional** | ***NetworkingProjectNetgwDetachOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetgwDetachOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwEventGet

> Event NetworkingProjectNetgwEventGet(ctx, projectId, locationId, netgwId, eventId)

Get networking/netgw.event

Get networking/netgw.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
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


## NetworkingProjectNetgwEventList

> []Event NetworkingProjectNetgwEventList(ctx, projectId, locationId, netgwId, optional)

List networking/netgw.event

List networking/netgw.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
 **optional** | ***NetworkingProjectNetgwEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetgwEventListOpts struct


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


## NetworkingProjectNetgwGet

> Netgw NetworkingProjectNetgwGet(ctx, projectId, locationId, netgwId)

Get networking/netgw

Returns a single netgw

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwList

> []Netgw NetworkingProjectNetgwList(ctx, projectId, locationId, optional)

List networking/netgw

List netgw

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***NetworkingProjectNetgwListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetgwListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Netgw**](netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwServiceGet

> ResourceService NetworkingProjectNetgwServiceGet(ctx, projectId, locationId, netgwId, serviceId)

Get networking/netgw.service

Get networking/netgw.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
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


## NetworkingProjectNetgwServiceList

> []ResourceService NetworkingProjectNetgwServiceList(ctx, projectId, locationId, netgwId)

List networking/netgw.service

List networking/netgw.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 

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


## NetworkingProjectNetgwTagCreate

> Tag NetworkingProjectNetgwTagCreate(ctx, projectId, locationId, netgwId, tag)

Create networking/netgw.tag

Create networking/netgw.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
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


## NetworkingProjectNetgwTagDelete

> NetworkingProjectNetgwTagDelete(ctx, projectId, locationId, netgwId, tagId)

Delete networking/netgw.tag

Delete networking/netgw.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
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


## NetworkingProjectNetgwTagGet

> Tag NetworkingProjectNetgwTagGet(ctx, projectId, locationId, netgwId, tagId)

Get networking/netgw.tag

Get networking/netgw.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
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


## NetworkingProjectNetgwTagList

> []Tag NetworkingProjectNetgwTagList(ctx, projectId, locationId, netgwId)

List networking/netgw.tag

List networking/netgw.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 

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


## NetworkingProjectNetgwTagPut

> []Tag NetworkingProjectNetgwTagPut(ctx, projectId, locationId, netgwId, tag)

Replace networking/netgw.tag

Replace networking/netgw.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
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


## NetworkingProjectNetgwUpdate

> Netgw NetworkingProjectNetgwUpdate(ctx, projectId, locationId, netgwId, networkingProjectNetgwUpdate)

Update networking/netgw

Returns modified netgw

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netgwId** | **string**| Netgw Id | 
**networkingProjectNetgwUpdate** | [**NetworkingProjectNetgwUpdate**](NetworkingProjectNetgwUpdate.md)|  | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

