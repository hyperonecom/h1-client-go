# \ReservationApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReservationActionAssign**](ReservationApi.md#ReservationActionAssign) | **Post** /reservation/{reservationId}/actions/assign | /actions/assign
[**ReservationActionExtend**](ReservationApi.md#ReservationActionExtend) | **Post** /reservation/{reservationId}/actions/extend | /actions/extend
[**ReservationCreate**](ReservationApi.md#ReservationCreate) | **Post** /reservation | Create
[**ReservationDelete**](ReservationApi.md#ReservationDelete) | **Delete** /reservation/{reservationId} | Delete
[**ReservationDeleteAccessrightsIdentity**](ReservationApi.md#ReservationDeleteAccessrightsIdentity) | **Delete** /reservation/{reservationId}/accessrights/{identity} | /accessrights/:identity
[**ReservationDeleteTagKey**](ReservationApi.md#ReservationDeleteTagKey) | **Delete** /reservation/{reservationId}/tag/{key} | /tag/:key
[**ReservationGetServicesServiceId**](ReservationApi.md#ReservationGetServicesServiceId) | **Get** /reservation/{reservationId}/services/{serviceId} | /services/:serviceId
[**ReservationGetTag**](ReservationApi.md#ReservationGetTag) | **Get** /reservation/{reservationId}/tag | /tag
[**ReservationList**](ReservationApi.md#ReservationList) | **Get** /reservation | List
[**ReservationListAccessrights**](ReservationApi.md#ReservationListAccessrights) | **Get** /reservation/{reservationId}/accessrights | /accessrights
[**ReservationListQueue**](ReservationApi.md#ReservationListQueue) | **Get** /reservation/{reservationId}/queue | /queue
[**ReservationListServices**](ReservationApi.md#ReservationListServices) | **Get** /reservation/{reservationId}/services | /services
[**ReservationPatchTag**](ReservationApi.md#ReservationPatchTag) | **Patch** /reservation/{reservationId}/tag | /tag
[**ReservationPostAccessrights**](ReservationApi.md#ReservationPostAccessrights) | **Post** /reservation/{reservationId}/accessrights | /accessrights
[**ReservationPutTag**](ReservationApi.md#ReservationPutTag) | **Put** /reservation/{reservationId}/tag | /tag
[**ReservationShow**](ReservationApi.md#ReservationShow) | **Get** /reservation/{reservationId} | Get
[**ReservationUpdate**](ReservationApi.md#ReservationUpdate) | **Patch** /reservation/{reservationId} | Update



## ReservationActionAssign

> Reservation ReservationActionAssign(ctx, reservationId, reservationActionAssign)

/actions/assign

Action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
**reservationActionAssign** | [**ReservationActionAssign**](ReservationActionAssign.md)|  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationActionExtend

> Reservation ReservationActionExtend(ctx, reservationId)

/actions/extend

Action extend

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationCreate

> Reservation ReservationCreate(ctx, reservationCreate)

Create

Create reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationCreate** | [**ReservationCreate**](ReservationCreate.md)|  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationDelete

> ReservationDelete(ctx, reservationId)

Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 

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


## ReservationDeleteAccessrightsIdentity

> Reservation ReservationDeleteAccessrightsIdentity(ctx, reservationId, identity)

/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
**identity** | **string**| identity | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationDeleteTagKey

> map[string]string ReservationDeleteTagKey(ctx, reservationId, key)

/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
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


## ReservationGetServicesServiceId

> ReservationServices ReservationGetServicesServiceId(ctx, reservationId, serviceId)

/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
**serviceId** | **string**| serviceId | 

### Return type

[**ReservationServices**](reservation.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationGetTag

> map[string]string ReservationGetTag(ctx, reservationId)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 

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


## ReservationList

> []Reservation ReservationList(ctx, optional)

List

List reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReservationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReservationListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationListAccessrights

> []string ReservationListAccessrights(ctx, reservationId)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 

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


## ReservationListQueue

> []Event ReservationListQueue(ctx, reservationId, optional)

/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
 **optional** | ***ReservationListQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReservationListQueueOpts struct


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


## ReservationListServices

> []ReservationServices ReservationListServices(ctx, reservationId)

/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 

### Return type

[**[]ReservationServices**](reservation.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationPatchTag

> map[string]string ReservationPatchTag(ctx, reservationId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
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


## ReservationPostAccessrights

> Reservation ReservationPostAccessrights(ctx, reservationId, reservationPostAccessrights)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
**reservationPostAccessrights** | [**ReservationPostAccessrights**](ReservationPostAccessrights.md)|  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationPutTag

> map[string]string ReservationPutTag(ctx, reservationId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
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


## ReservationShow

> Reservation ReservationShow(ctx, reservationId)

Get

Returns a single reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationUpdate

> Reservation ReservationUpdate(ctx, reservationId, reservationUpdate)

Update

Returns modified reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string**| ID of reservation | 
**reservationUpdate** | [**ReservationUpdate**](ReservationUpdate.md)|  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

