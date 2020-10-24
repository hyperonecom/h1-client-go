# \NetworkingProjectNetadpApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectNetadpCreate**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpCreate) | **Post** /networking/{locationId}/project/{projectId}/netadp | Create networking/netadp
[**NetworkingProjectNetadpDelete**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpDelete) | **Delete** /networking/{locationId}/project/{projectId}/netadp/{netadpId} | Delete networking/netadp
[**NetworkingProjectNetadpEventGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpEventGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/event/{eventId} | Get networking/netadp.event
[**NetworkingProjectNetadpEventList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpEventList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/event | List networking/netadp.event
[**NetworkingProjectNetadpGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId} | Get networking/netadp
[**NetworkingProjectNetadpList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpList) | **Get** /networking/{locationId}/project/{projectId}/netadp | List networking/netadp
[**NetworkingProjectNetadpMetricGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpMetricGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/metric/{metricId} | Get networking/netadp.metric
[**NetworkingProjectNetadpMetricList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpMetricList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/metric | List networking/netadp.metric
[**NetworkingProjectNetadpMetricPointList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpMetricPointList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/metric/{metricId}/point | List networking/netadp.point
[**NetworkingProjectNetadpServiceGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpServiceGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/service/{serviceId} | Get networking/netadp.service
[**NetworkingProjectNetadpServiceList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpServiceList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/service | List networking/netadp.service
[**NetworkingProjectNetadpTagCreate**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagCreate) | **Post** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag | Create networking/netadp.tag
[**NetworkingProjectNetadpTagDelete**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag/{tagId} | Delete networking/netadp.tag
[**NetworkingProjectNetadpTagGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag/{tagId} | Get networking/netadp.tag
[**NetworkingProjectNetadpTagList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag | List networking/netadp.tag
[**NetworkingProjectNetadpTagPut**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagPut) | **Put** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag | Replace networking/netadp.tag
[**NetworkingProjectNetadpUpdate**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpUpdate) | **Patch** /networking/{locationId}/project/{projectId}/netadp/{netadpId} | Update networking/netadp



## NetworkingProjectNetadpCreate

> Netadp NetworkingProjectNetadpCreate(ctx, projectId, locationId, networkingProjectNetadpCreate, optional)

Create networking/netadp

Create netadp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkingProjectNetadpCreate** | [**NetworkingProjectNetadpCreate**](NetworkingProjectNetadpCreate.md)|  | 
 **optional** | ***NetworkingProjectNetadpCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetadpCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Netadp**](netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetadpDelete

> NetworkingProjectNetadpDelete(ctx, projectId, locationId, netadpId)

Delete networking/netadp

Delete netadp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 

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


## NetworkingProjectNetadpEventGet

> Event NetworkingProjectNetadpEventGet(ctx, projectId, locationId, netadpId, eventId)

Get networking/netadp.event

Get networking/netadp.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
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


## NetworkingProjectNetadpEventList

> []Event NetworkingProjectNetadpEventList(ctx, projectId, locationId, netadpId, optional)

List networking/netadp.event

List networking/netadp.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
 **optional** | ***NetworkingProjectNetadpEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetadpEventListOpts struct


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


## NetworkingProjectNetadpGet

> Netadp NetworkingProjectNetadpGet(ctx, projectId, locationId, netadpId)

Get networking/netadp

Returns a single netadp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 

### Return type

[**Netadp**](netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetadpList

> []Netadp NetworkingProjectNetadpList(ctx, projectId, locationId, optional)

List networking/netadp

List netadp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***NetworkingProjectNetadpListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetadpListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignedResource** | **optional.String**| Filter by assigned.resource | 
 **assignedId** | **optional.String**| Filter by assigned.id | 
 **network** | **optional.String**| Filter by network | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Netadp**](netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetadpMetricGet

> Metric NetworkingProjectNetadpMetricGet(ctx, projectId, locationId, netadpId, metricId)

Get networking/netadp.metric

Get networking/netadp.metric

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
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


## NetworkingProjectNetadpMetricList

> []Metric NetworkingProjectNetadpMetricList(ctx, projectId, locationId, netadpId)

List networking/netadp.metric

List networking/netadp.metric

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 

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


## NetworkingProjectNetadpMetricPointList

> []Point NetworkingProjectNetadpMetricPointList(ctx, projectId, locationId, netadpId, metricId, optional)

List networking/netadp.point

List networking/netadp.point

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
**metricId** | **string**| metricId | 
 **optional** | ***NetworkingProjectNetadpMetricPointListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectNetadpMetricPointListOpts struct


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


## NetworkingProjectNetadpServiceGet

> ResourceService NetworkingProjectNetadpServiceGet(ctx, projectId, locationId, netadpId, serviceId)

Get networking/netadp.service

Get networking/netadp.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
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


## NetworkingProjectNetadpServiceList

> []ResourceService NetworkingProjectNetadpServiceList(ctx, projectId, locationId, netadpId)

List networking/netadp.service

List networking/netadp.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 

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


## NetworkingProjectNetadpTagCreate

> Tag NetworkingProjectNetadpTagCreate(ctx, projectId, locationId, netadpId, tag)

Create networking/netadp.tag

Create networking/netadp.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
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


## NetworkingProjectNetadpTagDelete

> NetworkingProjectNetadpTagDelete(ctx, projectId, locationId, netadpId, tagId)

Delete networking/netadp.tag

Delete networking/netadp.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
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


## NetworkingProjectNetadpTagGet

> Tag NetworkingProjectNetadpTagGet(ctx, projectId, locationId, netadpId, tagId)

Get networking/netadp.tag

Get networking/netadp.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
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


## NetworkingProjectNetadpTagList

> []Tag NetworkingProjectNetadpTagList(ctx, projectId, locationId, netadpId)

List networking/netadp.tag

List networking/netadp.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 

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


## NetworkingProjectNetadpTagPut

> []Tag NetworkingProjectNetadpTagPut(ctx, projectId, locationId, netadpId, tag)

Replace networking/netadp.tag

Replace networking/netadp.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
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


## NetworkingProjectNetadpUpdate

> Netadp NetworkingProjectNetadpUpdate(ctx, projectId, locationId, netadpId, networkingProjectNetadpUpdate)

Update networking/netadp

Returns modified netadp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**netadpId** | **string**| Netadp Id | 
**networkingProjectNetadpUpdate** | [**NetworkingProjectNetadpUpdate**](NetworkingProjectNetadpUpdate.md)|  | 

### Return type

[**Netadp**](netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

