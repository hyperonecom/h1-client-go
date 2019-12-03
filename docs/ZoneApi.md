# \ZoneApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZoneCreate**](ZoneApi.md#ZoneCreate) | **Post** /zone | Create
[**ZoneDelete**](ZoneApi.md#ZoneDelete) | **Delete** /zone/{zoneId} | Delete
[**ZoneDeleteAccessrightsIdentity**](ZoneApi.md#ZoneDeleteAccessrightsIdentity) | **Delete** /zone/{zoneId}/accessrights/{identity} | /accessrights/:identity
[**ZoneDeleteRecordsetRecordsetId**](ZoneApi.md#ZoneDeleteRecordsetRecordsetId) | **Delete** /zone/{zoneId}/recordset/{recordsetId} | /recordset/:recordsetId
[**ZoneDeleteTagKey**](ZoneApi.md#ZoneDeleteTagKey) | **Delete** /zone/{zoneId}/tag/{key} | /tag/:key
[**ZoneGetRecordsetRecordsetId**](ZoneApi.md#ZoneGetRecordsetRecordsetId) | **Get** /zone/{zoneId}/recordset/{recordsetId} | /recordset/:recordsetId
[**ZoneGetServicesServiceId**](ZoneApi.md#ZoneGetServicesServiceId) | **Get** /zone/{zoneId}/services/{serviceId} | /services/:serviceId
[**ZoneGetTag**](ZoneApi.md#ZoneGetTag) | **Get** /zone/{zoneId}/tag | /tag
[**ZoneList**](ZoneApi.md#ZoneList) | **Get** /zone | List
[**ZoneListAccessrights**](ZoneApi.md#ZoneListAccessrights) | **Get** /zone/{zoneId}/accessrights | /accessrights
[**ZoneListQueue**](ZoneApi.md#ZoneListQueue) | **Get** /zone/{zoneId}/queue | /queue
[**ZoneListRecordset**](ZoneApi.md#ZoneListRecordset) | **Get** /zone/{zoneId}/recordset | /recordset
[**ZoneListRecordsetRecordsetIdrecord**](ZoneApi.md#ZoneListRecordsetRecordsetIdrecord) | **Get** /zone/{zoneId}/recordset/{recordsetId}/record | /recordset/:recordsetId/record
[**ZoneListServices**](ZoneApi.md#ZoneListServices) | **Get** /zone/{zoneId}/services | /services
[**ZonePatchRecordsetRecordsetId**](ZoneApi.md#ZonePatchRecordsetRecordsetId) | **Patch** /zone/{zoneId}/recordset/{recordsetId} | /recordset/:recordsetId
[**ZonePatchTag**](ZoneApi.md#ZonePatchTag) | **Patch** /zone/{zoneId}/tag | /tag
[**ZonePostAccessrights**](ZoneApi.md#ZonePostAccessrights) | **Post** /zone/{zoneId}/accessrights | /accessrights
[**ZonePostRecordset**](ZoneApi.md#ZonePostRecordset) | **Post** /zone/{zoneId}/recordset | /recordset
[**ZonePutRecordsetRecordsetIdrecord**](ZoneApi.md#ZonePutRecordsetRecordsetIdrecord) | **Put** /zone/{zoneId}/recordset/{recordsetId}/record | /recordset/:recordsetId/record
[**ZonePutTag**](ZoneApi.md#ZonePutTag) | **Put** /zone/{zoneId}/tag | /tag
[**ZoneShow**](ZoneApi.md#ZoneShow) | **Get** /zone/{zoneId} | Get
[**ZoneUpdate**](ZoneApi.md#ZoneUpdate) | **Patch** /zone/{zoneId} | Update



## ZoneCreate

> Zone ZoneCreate(ctx, zoneCreate)

Create

Create zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneCreate** | [**ZoneCreate**](ZoneCreate.md)|  | 

### Return type

[**Zone**](zone.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneDelete

> ZoneDelete(ctx, zoneId)

Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneDeleteAccessrightsIdentity

> Zone ZoneDeleteAccessrightsIdentity(ctx, zoneId, identity)

/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**identity** | **string**| identity | 

### Return type

[**Zone**](zone.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneDeleteRecordsetRecordsetId

> Zone ZoneDeleteRecordsetRecordsetId(ctx, zoneId, recordsetId)

/recordset/:recordsetId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**recordsetId** | **string**| recordsetId | 

### Return type

[**Zone**](zone.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneDeleteTagKey

> map[string]string ZoneDeleteTagKey(ctx, zoneId, key)

/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneGetRecordsetRecordsetId

> ZoneRecordset ZoneGetRecordsetRecordsetId(ctx, zoneId, recordsetId)

/recordset/:recordsetId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**recordsetId** | **string**| recordsetId | 

### Return type

[**ZoneRecordset**](zone.recordset.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneGetServicesServiceId

> ZoneServices ZoneGetServicesServiceId(ctx, zoneId, serviceId)

/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**serviceId** | **string**| serviceId | 

### Return type

[**ZoneServices**](zone.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneGetTag

> map[string]string ZoneGetTag(ctx, zoneId)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneList

> []Zone ZoneList(ctx, optional)

List

List zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ZoneListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZoneListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Zone**](zone.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneListAccessrights

> []string ZoneListAccessrights(ctx, zoneId)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneListQueue

> []Event ZoneListQueue(ctx, zoneId, optional)

/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
 **optional** | ***ZoneListQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ZoneListQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Float32**| $limit | 
 **skip** | **optional.Float32**| $skip | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneListRecordset

> []ZoneRecordset ZoneListRecordset(ctx, zoneId)

/recordset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 

### Return type

[**[]ZoneRecordset**](zone.recordset.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneListRecordsetRecordsetIdrecord

> []ZoneRecordsetRecord ZoneListRecordsetRecordsetIdrecord(ctx, zoneId, recordsetId)

/recordset/:recordsetId/record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**recordsetId** | **string**| recordsetId | 

### Return type

[**[]ZoneRecordsetRecord**](zone.recordset.record.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneListServices

> []ZoneServices ZoneListServices(ctx, zoneId)

/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 

### Return type

[**[]ZoneServices**](zone.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZonePatchRecordsetRecordsetId

> ZoneRecordset ZonePatchRecordsetRecordsetId(ctx, zoneId, recordsetId, zonePatchRecordsetRecordsetId)

/recordset/:recordsetId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**recordsetId** | **string**| recordsetId | 
**zonePatchRecordsetRecordsetId** | [**ZonePatchRecordsetRecordsetId**](ZonePatchRecordsetRecordsetId.md)|  | 

### Return type

[**ZoneRecordset**](zone.recordset.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZonePatchTag

> map[string]string ZonePatchTag(ctx, zoneId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZonePostAccessrights

> Zone ZonePostAccessrights(ctx, zoneId, zonePostAccessrights)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**zonePostAccessrights** | [**ZonePostAccessrights**](ZonePostAccessrights.md)|  | 

### Return type

[**Zone**](zone.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZonePostRecordset

> ZoneRecordset ZonePostRecordset(ctx, zoneId, zonePostRecordset)

/recordset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**zonePostRecordset** | [**ZonePostRecordset**](ZonePostRecordset.md)|  | 

### Return type

[**ZoneRecordset**](zone.recordset.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZonePutRecordsetRecordsetIdrecord

> ZoneRecordsetRecord ZonePutRecordsetRecordsetIdrecord(ctx, zoneId, recordsetId, mapStringinterface)

/recordset/:recordsetId/record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**recordsetId** | **string**| recordsetId | 
**mapStringinterface** | [**[]map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**ZoneRecordsetRecord**](zone.recordset.record.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZonePutTag

> map[string]string ZonePutTag(ctx, zoneId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneShow

> Zone ZoneShow(ctx, zoneId)

Get

Returns a single zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 

### Return type

[**Zone**](zone.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneUpdate

> Zone ZoneUpdate(ctx, zoneId, zoneUpdate)

Update

Returns modified zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| ID of zone | 
**zoneUpdate** | [**ZoneUpdate**](ZoneUpdate.md)|  | 

### Return type

[**Zone**](zone.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

