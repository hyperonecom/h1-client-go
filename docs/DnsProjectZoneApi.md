# \DnsProjectZoneApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsProjectZoneCreate**](DnsProjectZoneApi.md#DnsProjectZoneCreate) | **Post** /dns/{locationId}/project/{projectId}/zone | Create dns/zone
[**DnsProjectZoneDelete**](DnsProjectZoneApi.md#DnsProjectZoneDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId} | Delete dns/zone
[**DnsProjectZoneEventGet**](DnsProjectZoneApi.md#DnsProjectZoneEventGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/event/{eventId} | Get dns/zone.event
[**DnsProjectZoneEventList**](DnsProjectZoneApi.md#DnsProjectZoneEventList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/event | List dns/zone.event
[**DnsProjectZoneGet**](DnsProjectZoneApi.md#DnsProjectZoneGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId} | Get dns/zone
[**DnsProjectZoneList**](DnsProjectZoneApi.md#DnsProjectZoneList) | **Get** /dns/{locationId}/project/{projectId}/zone | List dns/zone
[**DnsProjectZoneRecordsetCreate**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetCreate) | **Post** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset | Create dns/zone.recordset
[**DnsProjectZoneRecordsetDelete**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId} | Delete dns/zone.recordset
[**DnsProjectZoneRecordsetGet**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId} | Get dns/zone.recordset
[**DnsProjectZoneRecordsetList**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset | List dns/zone.recordset
[**DnsProjectZoneRecordsetPatch**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetPatch) | **Patch** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId} | Update dns/zone.recordset
[**DnsProjectZoneRecordsetRecordCreate**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordCreate) | **Post** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record | Create dns/zone.record
[**DnsProjectZoneRecordsetRecordDelete**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record/{recordId} | Delete dns/zone.record
[**DnsProjectZoneRecordsetRecordGet**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record/{recordId} | Get dns/zone.record
[**DnsProjectZoneRecordsetRecordList**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record | List dns/zone.record
[**DnsProjectZoneRecordsetRecordPut**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordPut) | **Put** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record | Replace dns/zone.record
[**DnsProjectZoneServiceGet**](DnsProjectZoneApi.md#DnsProjectZoneServiceGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/service/{serviceId} | Get dns/zone.service
[**DnsProjectZoneServiceList**](DnsProjectZoneApi.md#DnsProjectZoneServiceList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/service | List dns/zone.service
[**DnsProjectZoneTagCreate**](DnsProjectZoneApi.md#DnsProjectZoneTagCreate) | **Post** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag | Create dns/zone.tag
[**DnsProjectZoneTagDelete**](DnsProjectZoneApi.md#DnsProjectZoneTagDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag/{tagId} | Delete dns/zone.tag
[**DnsProjectZoneTagGet**](DnsProjectZoneApi.md#DnsProjectZoneTagGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag/{tagId} | Get dns/zone.tag
[**DnsProjectZoneTagList**](DnsProjectZoneApi.md#DnsProjectZoneTagList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag | List dns/zone.tag
[**DnsProjectZoneTagPut**](DnsProjectZoneApi.md#DnsProjectZoneTagPut) | **Put** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag | Replace dns/zone.tag
[**DnsProjectZoneUpdate**](DnsProjectZoneApi.md#DnsProjectZoneUpdate) | **Patch** /dns/{locationId}/project/{projectId}/zone/{zoneId} | Update dns/zone



## DnsProjectZoneCreate

> Zone DnsProjectZoneCreate(ctx, projectId, locationId, dnsProjectZoneCreate, optional)

Create dns/zone

Create zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**dnsProjectZoneCreate** | [**DnsProjectZoneCreate**](DnsProjectZoneCreate.md)|  | 
 **optional** | ***DnsProjectZoneCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DnsProjectZoneCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Zone**](zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneDelete

> DnsProjectZoneDelete(ctx, projectId, locationId, zoneId)

Delete dns/zone

Delete zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 

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


## DnsProjectZoneEventGet

> Event DnsProjectZoneEventGet(ctx, projectId, locationId, zoneId, eventId)

Get dns/zone.event

Get dns/zone.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
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


## DnsProjectZoneEventList

> []Event DnsProjectZoneEventList(ctx, projectId, locationId, zoneId, optional)

List dns/zone.event

List dns/zone.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
 **optional** | ***DnsProjectZoneEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DnsProjectZoneEventListOpts struct


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


## DnsProjectZoneGet

> Zone DnsProjectZoneGet(ctx, projectId, locationId, zoneId)

Get dns/zone

Returns a single zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 

### Return type

[**Zone**](zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneList

> []Zone DnsProjectZoneList(ctx, projectId, locationId, optional)

List dns/zone

List zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***DnsProjectZoneListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DnsProjectZoneListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Zone**](zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetCreate

> DnsRecordset DnsProjectZoneRecordsetCreate(ctx, projectId, locationId, zoneId, dnsRecordset)

Create dns/zone.recordset

Create dns/zone.recordset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**dnsRecordset** | [**DnsRecordset**](DnsRecordset.md)|  | 

### Return type

[**DnsRecordset**](dns.recordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetDelete

> Zone DnsProjectZoneRecordsetDelete(ctx, projectId, locationId, zoneId, recordsetId)

Delete dns/zone.recordset

Delete dns/zone.recordset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 

### Return type

[**Zone**](zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetGet

> DnsRecordset DnsProjectZoneRecordsetGet(ctx, projectId, locationId, zoneId, recordsetId)

Get dns/zone.recordset

Get dns/zone.recordset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 

### Return type

[**DnsRecordset**](dns.recordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetList

> []DnsRecordset DnsProjectZoneRecordsetList(ctx, projectId, locationId, zoneId)

List dns/zone.recordset

List dns/zone.recordset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 

### Return type

[**[]DnsRecordset**](dns.recordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetPatch

> DnsRecordset DnsProjectZoneRecordsetPatch(ctx, projectId, locationId, zoneId, recordsetId, dnsProjectZoneRecordsetPatch)

Update dns/zone.recordset

Update dns/zone.recordset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 
**dnsProjectZoneRecordsetPatch** | [**DnsProjectZoneRecordsetPatch**](DnsProjectZoneRecordsetPatch.md)|  | 

### Return type

[**DnsRecordset**](dns.recordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordCreate

> DnsRecord DnsProjectZoneRecordsetRecordCreate(ctx, projectId, locationId, zoneId, recordsetId, dnsRecord)

Create dns/zone.record

Create dns/zone.record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 
**dnsRecord** | [**DnsRecord**](DnsRecord.md)|  | 

### Return type

[**DnsRecord**](dns.record.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordDelete

> DnsProjectZoneRecordsetRecordDelete(ctx, projectId, locationId, zoneId, recordsetId, recordId)

Delete dns/zone.record

Delete dns/zone.record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 
**recordId** | **string**| recordId | 

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


## DnsProjectZoneRecordsetRecordGet

> DnsRecord DnsProjectZoneRecordsetRecordGet(ctx, projectId, locationId, zoneId, recordsetId, recordId)

Get dns/zone.record

Get dns/zone.record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 
**recordId** | **string**| recordId | 

### Return type

[**DnsRecord**](dns.record.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordList

> []DnsRecord DnsProjectZoneRecordsetRecordList(ctx, projectId, locationId, zoneId, recordsetId)

List dns/zone.record

List dns/zone.record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 

### Return type

[**[]DnsRecord**](dns.record.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordPut

> []DnsRecord DnsProjectZoneRecordsetRecordPut(ctx, projectId, locationId, zoneId, recordsetId, dnsRecord)

Replace dns/zone.record

Replace dns/zone.record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**recordsetId** | **string**| recordsetId | 
**dnsRecord** | [**[]DnsRecord**](dns.record.md)|  | 

### Return type

[**[]DnsRecord**](dns.record.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneServiceGet

> ResourceService DnsProjectZoneServiceGet(ctx, projectId, locationId, zoneId, serviceId)

Get dns/zone.service

Get dns/zone.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
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


## DnsProjectZoneServiceList

> []ResourceService DnsProjectZoneServiceList(ctx, projectId, locationId, zoneId)

List dns/zone.service

List dns/zone.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 

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


## DnsProjectZoneTagCreate

> Tag DnsProjectZoneTagCreate(ctx, projectId, locationId, zoneId, tag)

Create dns/zone.tag

Create dns/zone.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
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


## DnsProjectZoneTagDelete

> DnsProjectZoneTagDelete(ctx, projectId, locationId, zoneId, tagId)

Delete dns/zone.tag

Delete dns/zone.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
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


## DnsProjectZoneTagGet

> Tag DnsProjectZoneTagGet(ctx, projectId, locationId, zoneId, tagId)

Get dns/zone.tag

Get dns/zone.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
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


## DnsProjectZoneTagList

> []Tag DnsProjectZoneTagList(ctx, projectId, locationId, zoneId)

List dns/zone.tag

List dns/zone.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 

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


## DnsProjectZoneTagPut

> []Tag DnsProjectZoneTagPut(ctx, projectId, locationId, zoneId, tag)

Replace dns/zone.tag

Replace dns/zone.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
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


## DnsProjectZoneUpdate

> Zone DnsProjectZoneUpdate(ctx, projectId, locationId, zoneId, dnsProjectZoneUpdate)

Update dns/zone

Returns modified zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**zoneId** | **string**| Zone Id | 
**dnsProjectZoneUpdate** | [**DnsProjectZoneUpdate**](DnsProjectZoneUpdate.md)|  | 

### Return type

[**Zone**](zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

