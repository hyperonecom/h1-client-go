# \NetgwApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionNetgwAttach**](NetgwApi.md#ActionNetgwAttach) | **Post** /netgw/{netgwId}/actions/attach | /actions/attach
[**ActionNetgwDetach**](NetgwApi.md#ActionNetgwDetach) | **Post** /netgw/{netgwId}/actions/detach | /actions/detach
[**CreateNetgw**](NetgwApi.md#CreateNetgw) | **Post** /netgw | Create
[**DeleteNetgw**](NetgwApi.md#DeleteNetgw) | **Delete** /netgw/{netgwId} | Delete
[**ListNetgw**](NetgwApi.md#ListNetgw) | **Get** /netgw | List
[**OperationNetgwDeleteaccessrightsIdentity**](NetgwApi.md#OperationNetgwDeleteaccessrightsIdentity) | **Delete** /netgw/{netgwId}/accessrights/{identity} | /accessrights/:identity
[**OperationNetgwDeletetagKey**](NetgwApi.md#OperationNetgwDeletetagKey) | **Delete** /netgw/{netgwId}/tag/{key} | /tag/:key
[**OperationNetgwGetservicesServiceId**](NetgwApi.md#OperationNetgwGetservicesServiceId) | **Get** /netgw/{netgwId}/services/{serviceId} | /services/:serviceId
[**OperationNetgwGettag**](NetgwApi.md#OperationNetgwGettag) | **Get** /netgw/{netgwId}/tag/ | /tag/
[**OperationNetgwListaccessrights**](NetgwApi.md#OperationNetgwListaccessrights) | **Get** /netgw/{netgwId}/accessrights/ | /accessrights/
[**OperationNetgwListqueue**](NetgwApi.md#OperationNetgwListqueue) | **Get** /netgw/{netgwId}/queue/ | /queue/
[**OperationNetgwListservices**](NetgwApi.md#OperationNetgwListservices) | **Get** /netgw/{netgwId}/services/ | /services/
[**OperationNetgwPatchtag**](NetgwApi.md#OperationNetgwPatchtag) | **Patch** /netgw/{netgwId}/tag/ | /tag/
[**OperationNetgwPostaccessrights**](NetgwApi.md#OperationNetgwPostaccessrights) | **Post** /netgw/{netgwId}/accessrights/ | /accessrights/
[**ShowNetgw**](NetgwApi.md#ShowNetgw) | **Get** /netgw/{netgwId} | Get
[**UpdateNetgw**](NetgwApi.md#UpdateNetgw) | **Patch** /netgw/{netgwId} | Update


# **ActionNetgwAttach**
> Netgw ActionNetgwAttach(ctx, netgwId)
/actions/attach

Action attach

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionNetgwDetach**
> Netgw ActionNetgwDetach(ctx, netgwId)
/actions/detach

Action detach

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNetgw**
> Netgw CreateNetgw(ctx, netgwCreate)
Create

Create netgw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwCreate** | [**NetgwCreate**](NetgwCreate.md)|  | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetgw**
> DeleteNetgw(ctx, netgwId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetgw**
> []Netgw ListNetgw(ctx, optional)
List

List netgw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListNetgwOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNetgwOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwDeleteaccessrightsIdentity**
> Netgw OperationNetgwDeleteaccessrightsIdentity(ctx, netgwId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **identity** | **string**| identity | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwDeletetagKey**
> map[string]string OperationNetgwDeletetagKey(ctx, netgwId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwGetservicesServiceId**
> NetgwServices OperationNetgwGetservicesServiceId(ctx, netgwId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **serviceId** | **string**| serviceId | 

### Return type

[**NetgwServices**](netgw.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwGettag**
> map[string]string OperationNetgwGettag(ctx, netgwId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwListaccessrights**
> []string OperationNetgwListaccessrights(ctx, netgwId)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwListqueue**
> []Event OperationNetgwListqueue(ctx, netgwId)
/queue/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwListservices**
> []NetgwServices OperationNetgwListservices(ctx, netgwId)
/services/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

[**[]NetgwServices**](netgw.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwPatchtag**
> map[string]string OperationNetgwPatchtag(ctx, netgwId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetgwPostaccessrights**
> string OperationNetgwPostaccessrights(ctx, netgwId, resourceAccessRight)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **resourceAccessRight** | [**ResourceAccessRight**](ResourceAccessRight.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowNetgw**
> Netgw ShowNetgw(ctx, netgwId)
Get

Returns a single netgw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetgw**
> Netgw UpdateNetgw(ctx, netgwId, optional)
Update

Returns modified netgw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
 **optional** | ***UpdateNetgwOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateNetgwOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject10** | [**optional.Interface of InlineObject10**](InlineObject10.md)|  | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

