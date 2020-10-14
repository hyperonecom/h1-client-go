# \NetworkingProjectIpApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectIpAssociate**](NetworkingProjectIpApi.md#NetworkingProjectIpAssociate) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/associate | Associate networking/ip
[**NetworkingProjectIpCreate**](NetworkingProjectIpApi.md#NetworkingProjectIpCreate) | **Post** /networking/{locationId}/project/{projectId}/ip | Create networking/ip
[**NetworkingProjectIpDelete**](NetworkingProjectIpApi.md#NetworkingProjectIpDelete) | **Delete** /networking/{locationId}/project/{projectId}/ip/{ipId} | Delete networking/ip
[**NetworkingProjectIpDisassociate**](NetworkingProjectIpApi.md#NetworkingProjectIpDisassociate) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/disassociate | Disassociate networking/ip
[**NetworkingProjectIpEventGet**](NetworkingProjectIpApi.md#NetworkingProjectIpEventGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/event/{eventId} | Get networking/ip.event
[**NetworkingProjectIpEventList**](NetworkingProjectIpApi.md#NetworkingProjectIpEventList) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/event | List networking/ip.event
[**NetworkingProjectIpGet**](NetworkingProjectIpApi.md#NetworkingProjectIpGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId} | Get networking/ip
[**NetworkingProjectIpList**](NetworkingProjectIpApi.md#NetworkingProjectIpList) | **Get** /networking/{locationId}/project/{projectId}/ip | List networking/ip
[**NetworkingProjectIpPersist**](NetworkingProjectIpApi.md#NetworkingProjectIpPersist) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/persist | Persist networking/ip
[**NetworkingProjectIpServiceGet**](NetworkingProjectIpApi.md#NetworkingProjectIpServiceGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/service/{serviceId} | Get networking/ip.service
[**NetworkingProjectIpServiceList**](NetworkingProjectIpApi.md#NetworkingProjectIpServiceList) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/service | List networking/ip.service
[**NetworkingProjectIpTagCreate**](NetworkingProjectIpApi.md#NetworkingProjectIpTagCreate) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag | Create networking/ip.tag
[**NetworkingProjectIpTagDelete**](NetworkingProjectIpApi.md#NetworkingProjectIpTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag/{tagId} | Delete networking/ip.tag
[**NetworkingProjectIpTagGet**](NetworkingProjectIpApi.md#NetworkingProjectIpTagGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag/{tagId} | Get networking/ip.tag
[**NetworkingProjectIpTagList**](NetworkingProjectIpApi.md#NetworkingProjectIpTagList) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag | List networking/ip.tag
[**NetworkingProjectIpTagPut**](NetworkingProjectIpApi.md#NetworkingProjectIpTagPut) | **Put** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag | Replace networking/ip.tag
[**NetworkingProjectIpTransfer**](NetworkingProjectIpApi.md#NetworkingProjectIpTransfer) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/transfer | Transfer networking/ip
[**NetworkingProjectIpUpdate**](NetworkingProjectIpApi.md#NetworkingProjectIpUpdate) | **Patch** /networking/{locationId}/project/{projectId}/ip/{ipId} | Update networking/ip



## NetworkingProjectIpAssociate

> Ip NetworkingProjectIpAssociate(ctx, projectId, locationId, ipId, networkingProjectIpAssociate, optional)

Associate networking/ip

action associate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
**networkingProjectIpAssociate** | [**NetworkingProjectIpAssociate**](NetworkingProjectIpAssociate.md)|  | 
 **optional** | ***NetworkingProjectIpAssociateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectIpAssociateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpCreate

> Ip NetworkingProjectIpCreate(ctx, projectId, locationId, networkingProjectIpCreate, optional)

Create networking/ip

Create ip

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkingProjectIpCreate** | [**NetworkingProjectIpCreate**](NetworkingProjectIpCreate.md)|  | 
 **optional** | ***NetworkingProjectIpCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectIpCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpDelete

> NetworkingProjectIpDelete(ctx, projectId, locationId, ipId)

Delete networking/ip

Delete ip

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 

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


## NetworkingProjectIpDisassociate

> Ip NetworkingProjectIpDisassociate(ctx, projectId, locationId, ipId, optional)

Disassociate networking/ip

action disassociate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
 **optional** | ***NetworkingProjectIpDisassociateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectIpDisassociateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpEventGet

> Event NetworkingProjectIpEventGet(ctx, projectId, locationId, ipId, eventId)

Get networking/ip.event

Get networking/ip.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
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


## NetworkingProjectIpEventList

> []Event NetworkingProjectIpEventList(ctx, projectId, locationId, ipId, optional)

List networking/ip.event

List networking/ip.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
 **optional** | ***NetworkingProjectIpEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectIpEventListOpts struct


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


## NetworkingProjectIpGet

> Ip NetworkingProjectIpGet(ctx, projectId, locationId, ipId)

Get networking/ip

Returns a single ip

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 

### Return type

[**Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpList

> []Ip NetworkingProjectIpList(ctx, projectId, locationId, optional)

List networking/ip

List ip

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***NetworkingProjectIpListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectIpListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **network** | **optional.String**| Filter by network | 
 **associatedNetadp** | **optional.String**| Filter by associated.netadp | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpPersist

> Ip NetworkingProjectIpPersist(ctx, projectId, locationId, ipId, optional)

Persist networking/ip

action persist

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
 **optional** | ***NetworkingProjectIpPersistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectIpPersistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpServiceGet

> ResourceService NetworkingProjectIpServiceGet(ctx, projectId, locationId, ipId, serviceId)

Get networking/ip.service

Get networking/ip.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
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


## NetworkingProjectIpServiceList

> []ResourceService NetworkingProjectIpServiceList(ctx, projectId, locationId, ipId)

List networking/ip.service

List networking/ip.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 

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


## NetworkingProjectIpTagCreate

> Tag NetworkingProjectIpTagCreate(ctx, projectId, locationId, ipId, tag)

Create networking/ip.tag

Create networking/ip.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
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


## NetworkingProjectIpTagDelete

> NetworkingProjectIpTagDelete(ctx, projectId, locationId, ipId, tagId)

Delete networking/ip.tag

Delete networking/ip.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
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


## NetworkingProjectIpTagGet

> Tag NetworkingProjectIpTagGet(ctx, projectId, locationId, ipId, tagId)

Get networking/ip.tag

Get networking/ip.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
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


## NetworkingProjectIpTagList

> []Tag NetworkingProjectIpTagList(ctx, projectId, locationId, ipId)

List networking/ip.tag

List networking/ip.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 

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


## NetworkingProjectIpTagPut

> []Tag NetworkingProjectIpTagPut(ctx, projectId, locationId, ipId, tag)

Replace networking/ip.tag

Replace networking/ip.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
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


## NetworkingProjectIpTransfer

> Ip NetworkingProjectIpTransfer(ctx, projectId, locationId, ipId, networkingProjectIpTransfer, optional)

Transfer networking/ip

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
**networkingProjectIpTransfer** | [**NetworkingProjectIpTransfer**](NetworkingProjectIpTransfer.md)|  | 
 **optional** | ***NetworkingProjectIpTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectIpTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpUpdate

> Ip NetworkingProjectIpUpdate(ctx, projectId, locationId, ipId, networkingProjectIpUpdate)

Update networking/ip

Returns modified ip

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**ipId** | **string**| Ip Id | 
**networkingProjectIpUpdate** | [**NetworkingProjectIpUpdate**](NetworkingProjectIpUpdate.md)|  | 

### Return type

[**Ip**](ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

