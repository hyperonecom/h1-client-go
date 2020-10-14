# \BillingProjectReservationApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingProjectReservationAssign**](BillingProjectReservationApi.md#BillingProjectReservationAssign) | **Post** /billing/project/{projectId}/reservation/{reservationId}/actions/assign | Assign billing/reservation
[**BillingProjectReservationCreate**](BillingProjectReservationApi.md#BillingProjectReservationCreate) | **Post** /billing/project/{projectId}/reservation | Create billing/reservation
[**BillingProjectReservationDelete**](BillingProjectReservationApi.md#BillingProjectReservationDelete) | **Delete** /billing/project/{projectId}/reservation/{reservationId} | Delete billing/reservation
[**BillingProjectReservationEventGet**](BillingProjectReservationApi.md#BillingProjectReservationEventGet) | **Get** /billing/project/{projectId}/reservation/{reservationId}/event/{eventId} | Get billing/reservation.event
[**BillingProjectReservationEventList**](BillingProjectReservationApi.md#BillingProjectReservationEventList) | **Get** /billing/project/{projectId}/reservation/{reservationId}/event | List billing/reservation.event
[**BillingProjectReservationExtend**](BillingProjectReservationApi.md#BillingProjectReservationExtend) | **Post** /billing/project/{projectId}/reservation/{reservationId}/actions/extend | Extend billing/reservation
[**BillingProjectReservationGet**](BillingProjectReservationApi.md#BillingProjectReservationGet) | **Get** /billing/project/{projectId}/reservation/{reservationId} | Get billing/reservation
[**BillingProjectReservationList**](BillingProjectReservationApi.md#BillingProjectReservationList) | **Get** /billing/project/{projectId}/reservation | List billing/reservation
[**BillingProjectReservationServiceGet**](BillingProjectReservationApi.md#BillingProjectReservationServiceGet) | **Get** /billing/project/{projectId}/reservation/{reservationId}/service/{serviceId} | Get billing/reservation.service
[**BillingProjectReservationServiceList**](BillingProjectReservationApi.md#BillingProjectReservationServiceList) | **Get** /billing/project/{projectId}/reservation/{reservationId}/service | List billing/reservation.service
[**BillingProjectReservationTagCreate**](BillingProjectReservationApi.md#BillingProjectReservationTagCreate) | **Post** /billing/project/{projectId}/reservation/{reservationId}/tag | Create billing/reservation.tag
[**BillingProjectReservationTagDelete**](BillingProjectReservationApi.md#BillingProjectReservationTagDelete) | **Delete** /billing/project/{projectId}/reservation/{reservationId}/tag/{tagId} | Delete billing/reservation.tag
[**BillingProjectReservationTagGet**](BillingProjectReservationApi.md#BillingProjectReservationTagGet) | **Get** /billing/project/{projectId}/reservation/{reservationId}/tag/{tagId} | Get billing/reservation.tag
[**BillingProjectReservationTagList**](BillingProjectReservationApi.md#BillingProjectReservationTagList) | **Get** /billing/project/{projectId}/reservation/{reservationId}/tag | List billing/reservation.tag
[**BillingProjectReservationTagPut**](BillingProjectReservationApi.md#BillingProjectReservationTagPut) | **Put** /billing/project/{projectId}/reservation/{reservationId}/tag | Replace billing/reservation.tag
[**BillingProjectReservationUpdate**](BillingProjectReservationApi.md#BillingProjectReservationUpdate) | **Patch** /billing/project/{projectId}/reservation/{reservationId} | Update billing/reservation



## BillingProjectReservationAssign

> Reservation BillingProjectReservationAssign(ctx, projectId, reservationId, billingProjectReservationAssign, optional)

Assign billing/reservation

action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
**billingProjectReservationAssign** | [**BillingProjectReservationAssign**](BillingProjectReservationAssign.md)|  | 
 **optional** | ***BillingProjectReservationAssignOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingProjectReservationAssignOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationCreate

> Reservation BillingProjectReservationCreate(ctx, projectId, billingProjectReservationCreate, optional)

Create billing/reservation

Create reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**billingProjectReservationCreate** | [**BillingProjectReservationCreate**](BillingProjectReservationCreate.md)|  | 
 **optional** | ***BillingProjectReservationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingProjectReservationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationDelete

> BillingProjectReservationDelete(ctx, projectId, reservationId)

Delete billing/reservation

Delete reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 

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


## BillingProjectReservationEventGet

> Event BillingProjectReservationEventGet(ctx, projectId, reservationId, eventId)

Get billing/reservation.event

Get billing/reservation.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
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


## BillingProjectReservationEventList

> []Event BillingProjectReservationEventList(ctx, projectId, reservationId, optional)

List billing/reservation.event

List billing/reservation.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
 **optional** | ***BillingProjectReservationEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingProjectReservationEventListOpts struct


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


## BillingProjectReservationExtend

> Reservation BillingProjectReservationExtend(ctx, projectId, reservationId, optional)

Extend billing/reservation

action extend

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
 **optional** | ***BillingProjectReservationExtendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingProjectReservationExtendOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationGet

> Reservation BillingProjectReservationGet(ctx, projectId, reservationId)

Get billing/reservation

Returns a single reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationList

> []Reservation BillingProjectReservationList(ctx, projectId, optional)

List billing/reservation

List reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***BillingProjectReservationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingProjectReservationListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationServiceGet

> ResourceService BillingProjectReservationServiceGet(ctx, projectId, reservationId, serviceId)

Get billing/reservation.service

Get billing/reservation.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
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


## BillingProjectReservationServiceList

> []ResourceService BillingProjectReservationServiceList(ctx, projectId, reservationId)

List billing/reservation.service

List billing/reservation.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 

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


## BillingProjectReservationTagCreate

> Tag BillingProjectReservationTagCreate(ctx, projectId, reservationId, tag)

Create billing/reservation.tag

Create billing/reservation.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
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


## BillingProjectReservationTagDelete

> BillingProjectReservationTagDelete(ctx, projectId, reservationId, tagId)

Delete billing/reservation.tag

Delete billing/reservation.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
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


## BillingProjectReservationTagGet

> Tag BillingProjectReservationTagGet(ctx, projectId, reservationId, tagId)

Get billing/reservation.tag

Get billing/reservation.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
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


## BillingProjectReservationTagList

> []Tag BillingProjectReservationTagList(ctx, projectId, reservationId)

List billing/reservation.tag

List billing/reservation.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 

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


## BillingProjectReservationTagPut

> []Tag BillingProjectReservationTagPut(ctx, projectId, reservationId, tag)

Replace billing/reservation.tag

Replace billing/reservation.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
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


## BillingProjectReservationUpdate

> Reservation BillingProjectReservationUpdate(ctx, projectId, reservationId, billingProjectReservationUpdate)

Update billing/reservation

Returns modified reservation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**reservationId** | **string**| Reservation Id | 
**billingProjectReservationUpdate** | [**BillingProjectReservationUpdate**](BillingProjectReservationUpdate.md)|  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

