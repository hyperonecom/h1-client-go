# \NetworkingProjectNetworkApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectNetworkCreate**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkCreate) | **Post** /networking/{locationId}/project/{projectId}/network | Create networking/network
[**NetworkingProjectNetworkDelete**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkDelete) | **Delete** /networking/{locationId}/project/{projectId}/network/{networkId} | Delete networking/network
[**NetworkingProjectNetworkEventGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkEventGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/event/{eventId} | Get networking/network.event
[**NetworkingProjectNetworkEventList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkEventList) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/event | List networking/network.event
[**NetworkingProjectNetworkGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId} | Get networking/network
[**NetworkingProjectNetworkList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkList) | **Get** /networking/{locationId}/project/{projectId}/network | List networking/network
[**NetworkingProjectNetworkServiceGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkServiceGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/service/{serviceId} | Get networking/network.service
[**NetworkingProjectNetworkServiceList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkServiceList) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/service | List networking/network.service
[**NetworkingProjectNetworkTagCreate**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagCreate) | **Post** /networking/{locationId}/project/{projectId}/network/{networkId}/tag | Create networking/network.tag
[**NetworkingProjectNetworkTagDelete**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/network/{networkId}/tag/{tagId} | Delete networking/network.tag
[**NetworkingProjectNetworkTagGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/tag/{tagId} | Get networking/network.tag
[**NetworkingProjectNetworkTagList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagList) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/tag | List networking/network.tag
[**NetworkingProjectNetworkTagPut**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagPut) | **Put** /networking/{locationId}/project/{projectId}/network/{networkId}/tag | Replace networking/network.tag
[**NetworkingProjectNetworkUpdate**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkUpdate) | **Patch** /networking/{locationId}/project/{projectId}/network/{networkId} | Update networking/network



## NetworkingProjectNetworkCreate

> Network NetworkingProjectNetworkCreate(ctx, projectId, locationId, networkingProjectNetworkCreate, optional)

Create networking/network

Create network

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkingProjectNetworkCreate** | [**NetworkingProjectNetworkCreate**](NetworkingProjectNetworkCreate.md)|  | 
 **optional** | ***NetworkingProjectNetworkCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetworkCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Network**](network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetworkDelete

> NetworkingProjectNetworkDelete(ctx, projectId, locationId, networkId)

Delete networking/network

Delete network

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 

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


## NetworkingProjectNetworkEventGet

> Event NetworkingProjectNetworkEventGet(ctx, projectId, locationId, networkId, eventId)

Get networking/network.event

Get networking/network.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
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


## NetworkingProjectNetworkEventList

> []Event NetworkingProjectNetworkEventList(ctx, projectId, locationId, networkId, optional)

List networking/network.event

List networking/network.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
 **optional** | ***NetworkingProjectNetworkEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetworkEventListOpts struct


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


## NetworkingProjectNetworkGet

> Network NetworkingProjectNetworkGet(ctx, projectId, locationId, networkId)

Get networking/network

Returns a single network

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 

### Return type

[**Network**](network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetworkList

> []Network NetworkingProjectNetworkList(ctx, projectId, locationId, optional)

List networking/network

List network

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***NetworkingProjectNetworkListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetworkListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Network**](network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetworkServiceGet

> ResourceService NetworkingProjectNetworkServiceGet(ctx, projectId, locationId, networkId, serviceId)

Get networking/network.service

Get networking/network.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
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


## NetworkingProjectNetworkServiceList

> []ResourceService NetworkingProjectNetworkServiceList(ctx, projectId, locationId, networkId)

List networking/network.service

List networking/network.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 

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


## NetworkingProjectNetworkTagCreate

> Tag NetworkingProjectNetworkTagCreate(ctx, projectId, locationId, networkId, tag)

Create networking/network.tag

Create networking/network.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
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


## NetworkingProjectNetworkTagDelete

> NetworkingProjectNetworkTagDelete(ctx, projectId, locationId, networkId, tagId)

Delete networking/network.tag

Delete networking/network.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
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


## NetworkingProjectNetworkTagGet

> Tag NetworkingProjectNetworkTagGet(ctx, projectId, locationId, networkId, tagId)

Get networking/network.tag

Get networking/network.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
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


## NetworkingProjectNetworkTagList

> []Tag NetworkingProjectNetworkTagList(ctx, projectId, locationId, networkId)

List networking/network.tag

List networking/network.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 

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


## NetworkingProjectNetworkTagPut

> []Tag NetworkingProjectNetworkTagPut(ctx, projectId, locationId, networkId, tag)

Replace networking/network.tag

Replace networking/network.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
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


## NetworkingProjectNetworkUpdate

> Network NetworkingProjectNetworkUpdate(ctx, projectId, locationId, networkId, networkingProjectNetworkUpdate)

Update networking/network

Returns modified network

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkId** | **string**| Network Id | 
**networkingProjectNetworkUpdate** | [**NetworkingProjectNetworkUpdate**](NetworkingProjectNetworkUpdate.md)|  | 

### Return type

[**Network**](network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

