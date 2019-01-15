# \NetadpApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetadp**](NetadpApi.md#ListNetadp) | **Get** /netadp | List
[**OperationNetadpDeleteaccessrightsIdentity**](NetadpApi.md#OperationNetadpDeleteaccessrightsIdentity) | **Delete** /netadp/{netadpId}/accessrights/{identity} | /accessrights/:identity
[**OperationNetadpDeletetagKey**](NetadpApi.md#OperationNetadpDeletetagKey) | **Delete** /netadp/{netadpId}/tag/{key} | /tag/:key
[**OperationNetadpGetservicesServiceId**](NetadpApi.md#OperationNetadpGetservicesServiceId) | **Get** /netadp/{netadpId}/services/{serviceId} | /services/:serviceId
[**OperationNetadpGettag**](NetadpApi.md#OperationNetadpGettag) | **Get** /netadp/{netadpId}/tag/ | /tag/
[**OperationNetadpListaccessrights**](NetadpApi.md#OperationNetadpListaccessrights) | **Get** /netadp/{netadpId}/accessrights/ | /accessrights/
[**OperationNetadpListqueue**](NetadpApi.md#OperationNetadpListqueue) | **Get** /netadp/{netadpId}/queue/ | /queue/
[**OperationNetadpListservices**](NetadpApi.md#OperationNetadpListservices) | **Get** /netadp/{netadpId}/services/ | /services/
[**OperationNetadpPatchtag**](NetadpApi.md#OperationNetadpPatchtag) | **Patch** /netadp/{netadpId}/tag/ | /tag/
[**OperationNetadpPostaccessrights**](NetadpApi.md#OperationNetadpPostaccessrights) | **Post** /netadp/{netadpId}/accessrights/ | /accessrights/
[**ShowNetadp**](NetadpApi.md#ShowNetadp) | **Get** /netadp/{netadpId} | Get


# **ListNetadp**
> []Netadp ListNetadp(ctx, optional)
List

List netadp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListNetadpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNetadpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignedResource** | **optional.String**| Filter by assigned.resource | 
 **assignedId** | **optional.String**| Filter by assigned.id | 

### Return type

[**[]Netadp**](netadp.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpDeleteaccessrightsIdentity**
> Netadp OperationNetadpDeleteaccessrightsIdentity(ctx, netadpId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 
  **identity** | **string**| identity | 

### Return type

[**Netadp**](netadp.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpDeletetagKey**
> map[string]string OperationNetadpDeletetagKey(ctx, netadpId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpGetservicesServiceId**
> NetadpServices OperationNetadpGetservicesServiceId(ctx, netadpId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 
  **serviceId** | **string**| serviceId | 

### Return type

[**NetadpServices**](netadp.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpGettag**
> map[string]string OperationNetadpGettag(ctx, netadpId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpListaccessrights**
> []string OperationNetadpListaccessrights(ctx, netadpId)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpListqueue**
> []Event OperationNetadpListqueue(ctx, netadpId)
/queue/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpListservices**
> []NetadpServices OperationNetadpListservices(ctx, netadpId)
/services/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 

### Return type

[**[]NetadpServices**](netadp.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpPatchtag**
> map[string]string OperationNetadpPatchtag(ctx, netadpId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationNetadpPostaccessrights**
> string OperationNetadpPostaccessrights(ctx, netadpId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 
 **optional** | ***OperationNetadpPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationNetadpPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject35** | [**optional.Interface of InlineObject35**](InlineObject35.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowNetadp**
> Netadp ShowNetadp(ctx, netadpId)
Get

Returns a single netadp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netadpId** | **string**| ID of netadp | 

### Return type

[**Netadp**](netadp.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

