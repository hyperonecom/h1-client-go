# \IsoApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IsoActionTransfer**](IsoApi.md#IsoActionTransfer) | **Post** /iso/{isoId}/actions/transfer | /actions/transfer
[**IsoCreate**](IsoApi.md#IsoCreate) | **Post** /iso | Create
[**IsoDelete**](IsoApi.md#IsoDelete) | **Delete** /iso/{isoId} | Delete
[**IsoDeleteAccessrightsIdentity**](IsoApi.md#IsoDeleteAccessrightsIdentity) | **Delete** /iso/{isoId}/accessrights/{identity} | /accessrights/:identity
[**IsoDeleteTagKey**](IsoApi.md#IsoDeleteTagKey) | **Delete** /iso/{isoId}/tag/{key} | /tag/:key
[**IsoGetServicesServiceId**](IsoApi.md#IsoGetServicesServiceId) | **Get** /iso/{isoId}/services/{serviceId} | /services/:serviceId
[**IsoGetTag**](IsoApi.md#IsoGetTag) | **Get** /iso/{isoId}/tag | /tag
[**IsoList**](IsoApi.md#IsoList) | **Get** /iso | List
[**IsoListAccessrights**](IsoApi.md#IsoListAccessrights) | **Get** /iso/{isoId}/accessrights | /accessrights
[**IsoListQueue**](IsoApi.md#IsoListQueue) | **Get** /iso/{isoId}/queue | /queue
[**IsoListServices**](IsoApi.md#IsoListServices) | **Get** /iso/{isoId}/services | /services
[**IsoPatchTag**](IsoApi.md#IsoPatchTag) | **Patch** /iso/{isoId}/tag | /tag
[**IsoPostAccessrights**](IsoApi.md#IsoPostAccessrights) | **Post** /iso/{isoId}/accessrights | /accessrights
[**IsoShow**](IsoApi.md#IsoShow) | **Get** /iso/{isoId} | Get
[**IsoUpdate**](IsoApi.md#IsoUpdate) | **Patch** /iso/{isoId} | Update


# **IsoActionTransfer**
> Iso IsoActionTransfer(ctx, isoId, isoActionTransfer)
/actions/transfer

Action transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 
  **isoActionTransfer** | [**IsoActionTransfer**](IsoActionTransfer.md)|  | 

### Return type

[**Iso**](iso.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoCreate**
> Iso IsoCreate(ctx, isoCreate)
Create

Create iso

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoCreate** | [**IsoCreate**](IsoCreate.md)|  | 

### Return type

[**Iso**](iso.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoDelete**
> IsoDelete(ctx, isoId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoDeleteAccessrightsIdentity**
> Iso IsoDeleteAccessrightsIdentity(ctx, isoId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 
  **identity** | **string**| identity | 

### Return type

[**Iso**](iso.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoDeleteTagKey**
> map[string]interface{} IsoDeleteTagKey(ctx, isoId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoGetServicesServiceId**
> IsoServices IsoGetServicesServiceId(ctx, isoId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 
  **serviceId** | **string**| serviceId | 

### Return type

[**IsoServices**](iso.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoGetTag**
> map[string]interface{} IsoGetTag(ctx, isoId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoList**
> []Iso IsoList(ctx, optional)
List

List iso

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IsoListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IsoListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Iso**](iso.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoListAccessrights**
> []string IsoListAccessrights(ctx, isoId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoListQueue**
> []Event IsoListQueue(ctx, isoId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoListServices**
> []IsoServices IsoListServices(ctx, isoId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 

### Return type

[**[]IsoServices**](iso.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoPatchTag**
> map[string]interface{} IsoPatchTag(ctx, isoId, body)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 
  **body** | **map[string]interface{}**|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoPostAccessrights**
> Iso IsoPostAccessrights(ctx, isoId, isoPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 
  **isoPostAccessrights** | [**IsoPostAccessrights**](IsoPostAccessrights.md)|  | 

### Return type

[**Iso**](iso.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoShow**
> Iso IsoShow(ctx, isoId)
Get

Returns a single iso

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 

### Return type

[**Iso**](iso.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsoUpdate**
> Iso IsoUpdate(ctx, isoId, isoUpdate)
Update

Returns modified iso

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**| ID of iso | 
  **isoUpdate** | [**IsoUpdate**](IsoUpdate.md)|  | 

### Return type

[**Iso**](iso.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

