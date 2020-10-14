# \ComputeProjectReplicaApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeProjectReplicaCreate**](ComputeProjectReplicaApi.md#ComputeProjectReplicaCreate) | **Post** /compute/{locationId}/project/{projectId}/replica | Create compute/replica
[**ComputeProjectReplicaDelete**](ComputeProjectReplicaApi.md#ComputeProjectReplicaDelete) | **Delete** /compute/{locationId}/project/{projectId}/replica/{replicaId} | Delete compute/replica
[**ComputeProjectReplicaEventGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaEventGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/event/{eventId} | Get compute/replica.event
[**ComputeProjectReplicaEventList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaEventList) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/event | List compute/replica.event
[**ComputeProjectReplicaGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId} | Get compute/replica
[**ComputeProjectReplicaList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaList) | **Get** /compute/{locationId}/project/{projectId}/replica | List compute/replica
[**ComputeProjectReplicaServiceGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaServiceGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/service/{serviceId} | Get compute/replica.service
[**ComputeProjectReplicaServiceList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaServiceList) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/service | List compute/replica.service
[**ComputeProjectReplicaTagCreate**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagCreate) | **Post** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag | Create compute/replica.tag
[**ComputeProjectReplicaTagDelete**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagDelete) | **Delete** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag/{tagId} | Delete compute/replica.tag
[**ComputeProjectReplicaTagGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag/{tagId} | Get compute/replica.tag
[**ComputeProjectReplicaTagList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagList) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag | List compute/replica.tag
[**ComputeProjectReplicaTagPut**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagPut) | **Put** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag | Replace compute/replica.tag



## ComputeProjectReplicaCreate

> Replica ComputeProjectReplicaCreate(ctx, projectId, locationId, computeProjectReplicaCreate, optional)

Create compute/replica

Create replica

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**computeProjectReplicaCreate** | [**ComputeProjectReplicaCreate**](ComputeProjectReplicaCreate.md)|  | 
 **optional** | ***ComputeProjectReplicaCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectReplicaCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Replica**](replica.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectReplicaDelete

> ComputeProjectReplicaDelete(ctx, projectId, locationId, replicaId)

Delete compute/replica

Delete replica

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 

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


## ComputeProjectReplicaEventGet

> Event ComputeProjectReplicaEventGet(ctx, projectId, locationId, replicaId, eventId)

Get compute/replica.event

Get compute/replica.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 
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


## ComputeProjectReplicaEventList

> []Event ComputeProjectReplicaEventList(ctx, projectId, locationId, replicaId, optional)

List compute/replica.event

List compute/replica.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 
 **optional** | ***ComputeProjectReplicaEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectReplicaEventListOpts struct


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


## ComputeProjectReplicaGet

> Replica ComputeProjectReplicaGet(ctx, projectId, locationId, replicaId)

Get compute/replica

Returns a single replica

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 

### Return type

[**Replica**](replica.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectReplicaList

> []Replica ComputeProjectReplicaList(ctx, projectId, locationId, optional)

List compute/replica

List replica

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***ComputeProjectReplicaListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectReplicaListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Replica**](replica.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectReplicaServiceGet

> ResourceService ComputeProjectReplicaServiceGet(ctx, projectId, locationId, replicaId, serviceId)

Get compute/replica.service

Get compute/replica.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 
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


## ComputeProjectReplicaServiceList

> []ResourceService ComputeProjectReplicaServiceList(ctx, projectId, locationId, replicaId)

List compute/replica.service

List compute/replica.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 

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


## ComputeProjectReplicaTagCreate

> Tag ComputeProjectReplicaTagCreate(ctx, projectId, locationId, replicaId, tag)

Create compute/replica.tag

Create compute/replica.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 
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


## ComputeProjectReplicaTagDelete

> ComputeProjectReplicaTagDelete(ctx, projectId, locationId, replicaId, tagId)

Delete compute/replica.tag

Delete compute/replica.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 
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


## ComputeProjectReplicaTagGet

> Tag ComputeProjectReplicaTagGet(ctx, projectId, locationId, replicaId, tagId)

Get compute/replica.tag

Get compute/replica.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 
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


## ComputeProjectReplicaTagList

> []Tag ComputeProjectReplicaTagList(ctx, projectId, locationId, replicaId)

List compute/replica.tag

List compute/replica.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 

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


## ComputeProjectReplicaTagPut

> []Tag ComputeProjectReplicaTagPut(ctx, projectId, locationId, replicaId, tag)

Replace compute/replica.tag

Replace compute/replica.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**replicaId** | **string**| Replica Id | 
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

