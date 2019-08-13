# \NetadpApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetadpActionFirewallAdd**](NetadpApi.md#NetadpActionFirewallAdd) | **Post** /netadp/{netadpId}/actions/firewall_add | /actions/firewall_add
[**NetadpActionFirewallRemove**](NetadpApi.md#NetadpActionFirewallRemove) | **Post** /netadp/{netadpId}/actions/firewall_remove | /actions/firewall_remove
[**NetadpDeleteAccessrightsIdentity**](NetadpApi.md#NetadpDeleteAccessrightsIdentity) | **Delete** /netadp/{netadpId}/accessrights/{identity} | /accessrights/:identity
[**NetadpDeleteTagKey**](NetadpApi.md#NetadpDeleteTagKey) | **Delete** /netadp/{netadpId}/tag/{key} | /tag/:key
[**NetadpGetServicesServiceId**](NetadpApi.md#NetadpGetServicesServiceId) | **Get** /netadp/{netadpId}/services/{serviceId} | /services/:serviceId
[**NetadpGetTag**](NetadpApi.md#NetadpGetTag) | **Get** /netadp/{netadpId}/tag | /tag
[**NetadpList**](NetadpApi.md#NetadpList) | **Get** /netadp | List
[**NetadpListAccessrights**](NetadpApi.md#NetadpListAccessrights) | **Get** /netadp/{netadpId}/accessrights | /accessrights
[**NetadpListQueue**](NetadpApi.md#NetadpListQueue) | **Get** /netadp/{netadpId}/queue | /queue
[**NetadpListServices**](NetadpApi.md#NetadpListServices) | **Get** /netadp/{netadpId}/services | /services
[**NetadpPatchTag**](NetadpApi.md#NetadpPatchTag) | **Patch** /netadp/{netadpId}/tag | /tag
[**NetadpPostAccessrights**](NetadpApi.md#NetadpPostAccessrights) | **Post** /netadp/{netadpId}/accessrights | /accessrights
[**NetadpShow**](NetadpApi.md#NetadpShow) | **Get** /netadp/{netadpId} | Get



## NetadpActionFirewallAdd

> Netadp NetadpActionFirewallAdd(ctx, netadpId, netadpActionFirewallAdd)
/actions/firewall_add

Action firewall_add

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**netadpId** | **string**| ID of netadp | 
**netadpActionFirewallAdd** | [**NetadpActionFirewallAdd**](NetadpActionFirewallAdd.md)|  | 

### Return type

[**Netadp**](netadp.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpActionFirewallRemove

> Netadp NetadpActionFirewallRemove(ctx, netadpId)
/actions/firewall_remove

Action firewall_remove

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpDeleteAccessrightsIdentity

> Netadp NetadpDeleteAccessrightsIdentity(ctx, netadpId, identity)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpDeleteTagKey

> map[string]string NetadpDeleteTagKey(ctx, netadpId, key)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpGetServicesServiceId

> NetadpServices NetadpGetServicesServiceId(ctx, netadpId, serviceId)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpGetTag

> map[string]string NetadpGetTag(ctx, netadpId)
/tag

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpList

> []Netadp NetadpList(ctx, optional)
List

List netadp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetadpListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetadpListOpts struct


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpListAccessrights

> []string NetadpListAccessrights(ctx, netadpId)
/accessrights

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpListQueue

> []Event NetadpListQueue(ctx, netadpId)
/queue

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpListServices

> []NetadpServices NetadpListServices(ctx, netadpId)
/services

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpPatchTag

> map[string]string NetadpPatchTag(ctx, netadpId, requestBody)
/tag

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpPostAccessrights

> Netadp NetadpPostAccessrights(ctx, netadpId, netadpPostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**netadpId** | **string**| ID of netadp | 
**netadpPostAccessrights** | [**NetadpPostAccessrights**](NetadpPostAccessrights.md)|  | 

### Return type

[**Netadp**](netadp.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetadpShow

> Netadp NetadpShow(ctx, netadpId)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

