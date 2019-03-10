# \NetgwApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetgwActionAttach**](NetgwApi.md#NetgwActionAttach) | **Post** /netgw/{netgwId}/actions/attach | /actions/attach
[**NetgwActionDetach**](NetgwApi.md#NetgwActionDetach) | **Post** /netgw/{netgwId}/actions/detach | /actions/detach
[**NetgwCreate**](NetgwApi.md#NetgwCreate) | **Post** /netgw | Create
[**NetgwDelete**](NetgwApi.md#NetgwDelete) | **Delete** /netgw/{netgwId} | Delete
[**NetgwDeleteAccessrightsIdentity**](NetgwApi.md#NetgwDeleteAccessrightsIdentity) | **Delete** /netgw/{netgwId}/accessrights/{identity} | /accessrights/:identity
[**NetgwDeleteTagKey**](NetgwApi.md#NetgwDeleteTagKey) | **Delete** /netgw/{netgwId}/tag/{key} | /tag/:key
[**NetgwGetServicesServiceId**](NetgwApi.md#NetgwGetServicesServiceId) | **Get** /netgw/{netgwId}/services/{serviceId} | /services/:serviceId
[**NetgwGetTag**](NetgwApi.md#NetgwGetTag) | **Get** /netgw/{netgwId}/tag | /tag
[**NetgwList**](NetgwApi.md#NetgwList) | **Get** /netgw | List
[**NetgwListAccessrights**](NetgwApi.md#NetgwListAccessrights) | **Get** /netgw/{netgwId}/accessrights | /accessrights
[**NetgwListQueue**](NetgwApi.md#NetgwListQueue) | **Get** /netgw/{netgwId}/queue | /queue
[**NetgwListServices**](NetgwApi.md#NetgwListServices) | **Get** /netgw/{netgwId}/services | /services
[**NetgwPatchTag**](NetgwApi.md#NetgwPatchTag) | **Patch** /netgw/{netgwId}/tag | /tag
[**NetgwPostAccessrights**](NetgwApi.md#NetgwPostAccessrights) | **Post** /netgw/{netgwId}/accessrights | /accessrights
[**NetgwShow**](NetgwApi.md#NetgwShow) | **Get** /netgw/{netgwId} | Get
[**NetgwUpdate**](NetgwApi.md#NetgwUpdate) | **Patch** /netgw/{netgwId} | Update


# **NetgwActionAttach**
> Netgw NetgwActionAttach(ctx, netgwId, netgwActionAttach)
/actions/attach

Action attach

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **netgwActionAttach** | [**NetgwActionAttach**](NetgwActionAttach.md)|  | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetgwActionDetach**
> Netgw NetgwActionDetach(ctx, netgwId)
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

# **NetgwCreate**
> Netgw NetgwCreate(ctx, netgwCreate)
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

# **NetgwDelete**
> NetgwDelete(ctx, netgwId)
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

# **NetgwDeleteAccessrightsIdentity**
> Netgw NetgwDeleteAccessrightsIdentity(ctx, netgwId, identity)
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

# **NetgwDeleteTagKey**
> map[string]interface{} NetgwDeleteTagKey(ctx, netgwId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetgwGetServicesServiceId**
> NetgwServices NetgwGetServicesServiceId(ctx, netgwId, serviceId)
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

# **NetgwGetTag**
> map[string]interface{} NetgwGetTag(ctx, netgwId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetgwList**
> []Netgw NetgwList(ctx, optional)
List

List netgw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetgwListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetgwListOpts struct

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

# **NetgwListAccessrights**
> []string NetgwListAccessrights(ctx, netgwId)
/accessrights

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

# **NetgwListQueue**
> []Event NetgwListQueue(ctx, netgwId)
/queue

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

# **NetgwListServices**
> []NetgwServices NetgwListServices(ctx, netgwId)
/services

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

# **NetgwPatchTag**
> map[string]interface{} NetgwPatchTag(ctx, netgwId, requestBody)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetgwPostAccessrights**
> Netgw NetgwPostAccessrights(ctx, netgwId, netgwPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **netgwPostAccessrights** | [**NetgwPostAccessrights**](NetgwPostAccessrights.md)|  | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetgwShow**
> Netgw NetgwShow(ctx, netgwId)
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

# **NetgwUpdate**
> Netgw NetgwUpdate(ctx, netgwId, netgwUpdate)
Update

Returns modified netgw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netgwId** | **string**| ID of netgw | 
  **netgwUpdate** | [**NetgwUpdate**](NetgwUpdate.md)|  | 

### Return type

[**Netgw**](netgw.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

