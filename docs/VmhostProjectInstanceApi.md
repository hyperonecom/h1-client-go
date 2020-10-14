# \VmhostProjectInstanceApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmhostProjectInstanceEventGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceEventGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/event/{eventId} | Get vmhost/instance.event
[**VmhostProjectInstanceEventList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceEventList) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/event | List vmhost/instance.event
[**VmhostProjectInstanceGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId} | Get vmhost/instance
[**VmhostProjectInstanceList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceList) | **Get** /vmhost/{locationId}/project/{projectId}/instance | List vmhost/instance
[**VmhostProjectInstanceServiceGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceServiceGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/service/{serviceId} | Get vmhost/instance.service
[**VmhostProjectInstanceServiceList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceServiceList) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/service | List vmhost/instance.service
[**VmhostProjectInstanceTagCreate**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagCreate) | **Post** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag | Create vmhost/instance.tag
[**VmhostProjectInstanceTagDelete**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagDelete) | **Delete** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Delete vmhost/instance.tag
[**VmhostProjectInstanceTagGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Get vmhost/instance.tag
[**VmhostProjectInstanceTagList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagList) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag | List vmhost/instance.tag
[**VmhostProjectInstanceTagPut**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagPut) | **Put** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag | Replace vmhost/instance.tag



## VmhostProjectInstanceEventGet

> Event VmhostProjectInstanceEventGet(ctx, projectId, locationId, instanceId, eventId)

Get vmhost/instance.event

Get vmhost/instance.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## VmhostProjectInstanceEventList

> []Event VmhostProjectInstanceEventList(ctx, projectId, locationId, instanceId, optional)

List vmhost/instance.event

List vmhost/instance.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
 **optional** | ***VmhostProjectInstanceEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VmhostProjectInstanceEventListOpts struct


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


## VmhostProjectInstanceGet

> Vmhost VmhostProjectInstanceGet(ctx, projectId, locationId, instanceId)

Get vmhost/instance

Returns a single instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostProjectInstanceList

> []Vmhost VmhostProjectInstanceList(ctx, projectId, locationId, optional)

List vmhost/instance

List instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***VmhostProjectInstanceListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VmhostProjectInstanceListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabledServices** | **optional.String**| Filter by enabledServices | 

### Return type

[**[]Vmhost**](vmhost.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostProjectInstanceServiceGet

> ResourceService VmhostProjectInstanceServiceGet(ctx, projectId, locationId, instanceId, serviceId)

Get vmhost/instance.service

Get vmhost/instance.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## VmhostProjectInstanceServiceList

> []ResourceService VmhostProjectInstanceServiceList(ctx, projectId, locationId, instanceId)

List vmhost/instance.service

List vmhost/instance.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

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


## VmhostProjectInstanceTagCreate

> Tag VmhostProjectInstanceTagCreate(ctx, projectId, locationId, instanceId, tag)

Create vmhost/instance.tag

Create vmhost/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## VmhostProjectInstanceTagDelete

> VmhostProjectInstanceTagDelete(ctx, projectId, locationId, instanceId, tagId)

Delete vmhost/instance.tag

Delete vmhost/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## VmhostProjectInstanceTagGet

> Tag VmhostProjectInstanceTagGet(ctx, projectId, locationId, instanceId, tagId)

Get vmhost/instance.tag

Get vmhost/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## VmhostProjectInstanceTagList

> []Tag VmhostProjectInstanceTagList(ctx, projectId, locationId, instanceId)

List vmhost/instance.tag

List vmhost/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

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


## VmhostProjectInstanceTagPut

> []Tag VmhostProjectInstanceTagPut(ctx, projectId, locationId, instanceId, tag)

Replace vmhost/instance.tag

Replace vmhost/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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

