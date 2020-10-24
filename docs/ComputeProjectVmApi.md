# \ComputeProjectVmApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeProjectVmConnectGet**](ComputeProjectVmApi.md#ComputeProjectVmConnectGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/connect/{connectId} | Get compute/vm.connect
[**ComputeProjectVmConnectList**](ComputeProjectVmApi.md#ComputeProjectVmConnectList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/connect | List compute/vm.connect
[**ComputeProjectVmConnectOpen**](ComputeProjectVmApi.md#ComputeProjectVmConnectOpen) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/connect/{connectId}/actions/open | Open compute/vm.connect
[**ComputeProjectVmCreate**](ComputeProjectVmApi.md#ComputeProjectVmCreate) | **Post** /compute/{locationId}/project/{projectId}/vm | Create compute/vm
[**ComputeProjectVmDelete**](ComputeProjectVmApi.md#ComputeProjectVmDelete) | **Delete** /compute/{locationId}/project/{projectId}/vm/{vmId} | Delete compute/vm
[**ComputeProjectVmDiskCreate**](ComputeProjectVmApi.md#ComputeProjectVmDiskCreate) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/disk | Create compute/vm.disk
[**ComputeProjectVmDiskList**](ComputeProjectVmApi.md#ComputeProjectVmDiskList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/disk | List compute/vm.disk
[**ComputeProjectVmEventGet**](ComputeProjectVmApi.md#ComputeProjectVmEventGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/event/{eventId} | Get compute/vm.event
[**ComputeProjectVmEventList**](ComputeProjectVmApi.md#ComputeProjectVmEventList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/event | List compute/vm.event
[**ComputeProjectVmFlavour**](ComputeProjectVmApi.md#ComputeProjectVmFlavour) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/flavour | Flavour compute/vm
[**ComputeProjectVmGet**](ComputeProjectVmApi.md#ComputeProjectVmGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId} | Get compute/vm
[**ComputeProjectVmIsoCreate**](ComputeProjectVmApi.md#ComputeProjectVmIsoCreate) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/iso | Create compute/vm.iso
[**ComputeProjectVmIsoList**](ComputeProjectVmApi.md#ComputeProjectVmIsoList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/iso | List compute/vm.iso
[**ComputeProjectVmList**](ComputeProjectVmApi.md#ComputeProjectVmList) | **Get** /compute/{locationId}/project/{projectId}/vm | List compute/vm
[**ComputeProjectVmMetricGet**](ComputeProjectVmApi.md#ComputeProjectVmMetricGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/metric/{metricId} | Get compute/vm.metric
[**ComputeProjectVmMetricList**](ComputeProjectVmApi.md#ComputeProjectVmMetricList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/metric | List compute/vm.metric
[**ComputeProjectVmMetricPointList**](ComputeProjectVmApi.md#ComputeProjectVmMetricPointList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/metric/{metricId}/point | List compute/vm.point
[**ComputeProjectVmPasswordReset**](ComputeProjectVmApi.md#ComputeProjectVmPasswordReset) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/password_reset | Password reset compute/vm
[**ComputeProjectVmRestart**](ComputeProjectVmApi.md#ComputeProjectVmRestart) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/restart | Restart compute/vm
[**ComputeProjectVmSerialport**](ComputeProjectVmApi.md#ComputeProjectVmSerialport) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/serialport | Serialport compute/vm
[**ComputeProjectVmServiceGet**](ComputeProjectVmApi.md#ComputeProjectVmServiceGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/service/{serviceId} | Get compute/vm.service
[**ComputeProjectVmServiceList**](ComputeProjectVmApi.md#ComputeProjectVmServiceList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/service | List compute/vm.service
[**ComputeProjectVmStart**](ComputeProjectVmApi.md#ComputeProjectVmStart) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/start | Start compute/vm
[**ComputeProjectVmStop**](ComputeProjectVmApi.md#ComputeProjectVmStop) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/stop | Stop compute/vm
[**ComputeProjectVmTagCreate**](ComputeProjectVmApi.md#ComputeProjectVmTagCreate) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag | Create compute/vm.tag
[**ComputeProjectVmTagDelete**](ComputeProjectVmApi.md#ComputeProjectVmTagDelete) | **Delete** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag/{tagId} | Delete compute/vm.tag
[**ComputeProjectVmTagGet**](ComputeProjectVmApi.md#ComputeProjectVmTagGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag/{tagId} | Get compute/vm.tag
[**ComputeProjectVmTagList**](ComputeProjectVmApi.md#ComputeProjectVmTagList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag | List compute/vm.tag
[**ComputeProjectVmTagPut**](ComputeProjectVmApi.md#ComputeProjectVmTagPut) | **Put** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag | Replace compute/vm.tag
[**ComputeProjectVmTurnoff**](ComputeProjectVmApi.md#ComputeProjectVmTurnoff) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/turnoff | Turnoff compute/vm
[**ComputeProjectVmUpdate**](ComputeProjectVmApi.md#ComputeProjectVmUpdate) | **Patch** /compute/{locationId}/project/{projectId}/vm/{vmId} | Update compute/vm



## ComputeProjectVmConnectGet

> Connect ComputeProjectVmConnectGet(ctx, projectId, locationId, vmId, connectId)

Get compute/vm.connect

Get compute/vm.connect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**connectId** | **string**| connectId | 

### Return type

[**Connect**](connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmConnectList

> []Connect ComputeProjectVmConnectList(ctx, projectId, locationId, vmId)

List compute/vm.connect

List compute/vm.connect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

### Return type

[**[]Connect**](connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmConnectOpen

> ComputeProjectVmConnectOpen(ctx, projectId, locationId, vmId, connectId, computeProjectVmConnectOpen)

Open compute/vm.connect

action open

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**connectId** | **string**| connectId | 
**computeProjectVmConnectOpen** | [**ComputeProjectVmConnectOpen**](ComputeProjectVmConnectOpen.md)|  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmCreate

> Vm ComputeProjectVmCreate(ctx, projectId, locationId, computeProjectVmCreate, optional)

Create compute/vm

Create vm

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**computeProjectVmCreate** | [**ComputeProjectVmCreate**](ComputeProjectVmCreate.md)|  | 
 **optional** | ***ComputeProjectVmCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmDelete

> ComputeProjectVmDelete(ctx, projectId, locationId, vmId)

Delete compute/vm

Delete vm

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

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


## ComputeProjectVmDiskCreate

> Disk ComputeProjectVmDiskCreate(ctx, projectId, locationId, vmId, computeProjectVmDiskCreate)

Create compute/vm.disk

Create compute/vm.disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**computeProjectVmDiskCreate** | [**ComputeProjectVmDiskCreate**](ComputeProjectVmDiskCreate.md)|  | 

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


## ComputeProjectVmDiskList

> []Disk ComputeProjectVmDiskList(ctx, projectId, locationId, vmId)

List compute/vm.disk

List compute/vm.disk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

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


## ComputeProjectVmEventGet

> Event ComputeProjectVmEventGet(ctx, projectId, locationId, vmId, eventId)

Get compute/vm.event

Get compute/vm.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
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


## ComputeProjectVmEventList

> []Event ComputeProjectVmEventList(ctx, projectId, locationId, vmId, optional)

List compute/vm.event

List compute/vm.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
 **optional** | ***ComputeProjectVmEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmEventListOpts struct


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


## ComputeProjectVmFlavour

> Vm ComputeProjectVmFlavour(ctx, projectId, locationId, vmId, computeProjectVmFlavour, optional)

Flavour compute/vm

action flavour

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**computeProjectVmFlavour** | [**ComputeProjectVmFlavour**](ComputeProjectVmFlavour.md)|  | 
 **optional** | ***ComputeProjectVmFlavourOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmFlavourOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmGet

> Vm ComputeProjectVmGet(ctx, projectId, locationId, vmId)

Get compute/vm

Returns a single vm

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmIsoCreate

> Iso ComputeProjectVmIsoCreate(ctx, projectId, locationId, vmId, computeProjectVmIsoCreate)

Create compute/vm.iso

Create compute/vm.iso

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**computeProjectVmIsoCreate** | [**ComputeProjectVmIsoCreate**](ComputeProjectVmIsoCreate.md)|  | 

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


## ComputeProjectVmIsoList

> []Iso ComputeProjectVmIsoList(ctx, projectId, locationId, vmId)

List compute/vm.iso

List compute/vm.iso

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

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


## ComputeProjectVmList

> []Vm ComputeProjectVmList(ctx, projectId, locationId, optional)

List compute/vm

List vm

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***ComputeProjectVmListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmMetricGet

> Metric ComputeProjectVmMetricGet(ctx, projectId, locationId, vmId, metricId)

Get compute/vm.metric

Get compute/vm.metric

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
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


## ComputeProjectVmMetricList

> []Metric ComputeProjectVmMetricList(ctx, projectId, locationId, vmId)

List compute/vm.metric

List compute/vm.metric

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

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


## ComputeProjectVmMetricPointList

> []Point ComputeProjectVmMetricPointList(ctx, projectId, locationId, vmId, metricId, optional)

List compute/vm.point

List compute/vm.point

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**metricId** | **string**| metricId | 
 **optional** | ***ComputeProjectVmMetricPointListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmMetricPointListOpts struct


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


## ComputeProjectVmPasswordReset

> Vm ComputeProjectVmPasswordReset(ctx, projectId, locationId, vmId, computeProjectVmPasswordReset, optional)

Password reset compute/vm

action password_reset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**computeProjectVmPasswordReset** | [**ComputeProjectVmPasswordReset**](ComputeProjectVmPasswordReset.md)|  | 
 **optional** | ***ComputeProjectVmPasswordResetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmPasswordResetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmRestart

> Vm ComputeProjectVmRestart(ctx, projectId, locationId, vmId, optional)

Restart compute/vm

action restart

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
 **optional** | ***ComputeProjectVmRestartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmRestartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmSerialport

> *os.File ComputeProjectVmSerialport(ctx, projectId, locationId, vmId, computeProjectVmSerialport, optional)

Serialport compute/vm

action serialport

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**computeProjectVmSerialport** | [**ComputeProjectVmSerialport**](ComputeProjectVmSerialport.md)|  | 
 **optional** | ***ComputeProjectVmSerialportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmSerialportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmServiceGet

> ResourceService ComputeProjectVmServiceGet(ctx, projectId, locationId, vmId, serviceId)

Get compute/vm.service

Get compute/vm.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
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


## ComputeProjectVmServiceList

> []ResourceService ComputeProjectVmServiceList(ctx, projectId, locationId, vmId)

List compute/vm.service

List compute/vm.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

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


## ComputeProjectVmStart

> Vm ComputeProjectVmStart(ctx, projectId, locationId, vmId, optional)

Start compute/vm

action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
 **optional** | ***ComputeProjectVmStartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmStartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmStop

> Vm ComputeProjectVmStop(ctx, projectId, locationId, vmId, optional)

Stop compute/vm

action stop

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
 **optional** | ***ComputeProjectVmStopOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmStopOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmTagCreate

> Tag ComputeProjectVmTagCreate(ctx, projectId, locationId, vmId, tag)

Create compute/vm.tag

Create compute/vm.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
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


## ComputeProjectVmTagDelete

> ComputeProjectVmTagDelete(ctx, projectId, locationId, vmId, tagId)

Delete compute/vm.tag

Delete compute/vm.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
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


## ComputeProjectVmTagGet

> Tag ComputeProjectVmTagGet(ctx, projectId, locationId, vmId, tagId)

Get compute/vm.tag

Get compute/vm.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
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


## ComputeProjectVmTagList

> []Tag ComputeProjectVmTagList(ctx, projectId, locationId, vmId)

List compute/vm.tag

List compute/vm.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 

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


## ComputeProjectVmTagPut

> []Tag ComputeProjectVmTagPut(ctx, projectId, locationId, vmId, tag)

Replace compute/vm.tag

Replace compute/vm.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
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


## ComputeProjectVmTurnoff

> Vm ComputeProjectVmTurnoff(ctx, projectId, locationId, vmId, optional)

Turnoff compute/vm

action turnoff

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
 **optional** | ***ComputeProjectVmTurnoffOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ComputeProjectVmTurnoffOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmUpdate

> Vm ComputeProjectVmUpdate(ctx, projectId, locationId, vmId, computeProjectVmUpdate)

Update compute/vm

Returns modified vm

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vmId** | **string**| Vm Id | 
**computeProjectVmUpdate** | [**ComputeProjectVmUpdate**](ComputeProjectVmUpdate.md)|  | 

### Return type

[**Vm**](vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

