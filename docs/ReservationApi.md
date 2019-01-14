# \ReservationApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionReservationAssign**](ReservationApi.md#ActionReservationAssign) | **Post** /reservation/{reservationId}/actions/assign | /actions/assign
[**ActionReservationExtend**](ReservationApi.md#ActionReservationExtend) | **Post** /reservation/{reservationId}/actions/extend | /actions/extend
[**ActionReservationUnassign**](ReservationApi.md#ActionReservationUnassign) | **Post** /reservation/{reservationId}/actions/unassign | /actions/unassign
[**CreateReservation**](ReservationApi.md#CreateReservation) | **Post** /reservation | Create
[**DeleteReservation**](ReservationApi.md#DeleteReservation) | **Delete** /reservation/{reservationId} | Delete
[**ListReservation**](ReservationApi.md#ListReservation) | **Get** /reservation | List
[**OperationReservationDeleteaccessrightsIdentity**](ReservationApi.md#OperationReservationDeleteaccessrightsIdentity) | **Delete** /reservation/{reservationId}/accessrights/{identity} | /accessrights/:identity
[**OperationReservationDeletetagKey**](ReservationApi.md#OperationReservationDeletetagKey) | **Delete** /reservation/{reservationId}/tag/{key} | /tag/:key
[**OperationReservationGetservicesServiceId**](ReservationApi.md#OperationReservationGetservicesServiceId) | **Get** /reservation/{reservationId}/services/{serviceId} | /services/:serviceId
[**OperationReservationGettag**](ReservationApi.md#OperationReservationGettag) | **Get** /reservation/{reservationId}/tag/ | /tag/
[**OperationReservationListaccessrights**](ReservationApi.md#OperationReservationListaccessrights) | **Get** /reservation/{reservationId}/accessrights/ | /accessrights/
[**OperationReservationListqueue**](ReservationApi.md#OperationReservationListqueue) | **Get** /reservation/{reservationId}/queue/ | /queue/
[**OperationReservationListservices**](ReservationApi.md#OperationReservationListservices) | **Get** /reservation/{reservationId}/services/ | /services/
[**OperationReservationPatchtag**](ReservationApi.md#OperationReservationPatchtag) | **Patch** /reservation/{reservationId}/tag/ | /tag/
[**OperationReservationPostaccessrights**](ReservationApi.md#OperationReservationPostaccessrights) | **Post** /reservation/{reservationId}/accessrights/ | /accessrights/
[**ShowReservation**](ReservationApi.md#ShowReservation) | **Get** /reservation/{reservationId} | Get
[**UpdateReservation**](ReservationApi.md#UpdateReservation) | **Patch** /reservation/{reservationId} | Update


# **ActionReservationAssign**
> Reservation ActionReservationAssign(ctx, reservationId)
/actions/assign

Action assign

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionReservationExtend**
> Reservation ActionReservationExtend(ctx, reservationId)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionReservationUnassign**
> Reservation ActionReservationUnassign(ctx, reservationId)
/actions/unassign

Action unassign

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReservation**
> Reservation CreateReservation(ctx, optional)
Create

Create reservation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateReservationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject57** | [**optional.Interface of InlineObject57**](InlineObject57.md)|  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReservation**
> DeleteReservation(ctx, reservationId)
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
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReservation**
> []Reservation ListReservation(ctx, optional)
List

List reservation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListReservationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationDeleteaccessrightsIdentity**
> Reservation OperationReservationDeleteaccessrightsIdentity(ctx, reservationId, identity)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationDeletetagKey**
> map[string]string OperationReservationDeletetagKey(ctx, reservationId, key)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationGetservicesServiceId**
> ReservationServices OperationReservationGetservicesServiceId(ctx, reservationId, serviceId)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationGettag**
> map[string]string OperationReservationGettag(ctx, reservationId)
/tag/

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationListaccessrights**
> []string OperationReservationListaccessrights(ctx, reservationId)
/accessrights/

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationListqueue**
> []Event OperationReservationListqueue(ctx, reservationId)
/queue/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reservationId** | **string**| ID of reservation | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationListservices**
> []ReservationServices OperationReservationListservices(ctx, reservationId)
/services/

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationPatchtag**
> map[string]string OperationReservationPatchtag(ctx, reservationId, requestBody)
/tag/

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReservationPostaccessrights**
> string OperationReservationPostaccessrights(ctx, reservationId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reservationId** | **string**| ID of reservation | 
 **optional** | ***OperationReservationPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationReservationPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject59** | [**optional.Interface of InlineObject59**](InlineObject59.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowReservation**
> Reservation ShowReservation(ctx, reservationId)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReservation**
> Reservation UpdateReservation(ctx, reservationId, optional)
Update

Returns modified reservation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reservationId** | **string**| ID of reservation | 
 **optional** | ***UpdateReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateReservationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject58** | [**optional.Interface of InlineObject58**](InlineObject58.md)|  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

