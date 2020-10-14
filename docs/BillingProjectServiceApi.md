# \BillingProjectServiceApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingProjectServiceGet**](BillingProjectServiceApi.md#BillingProjectServiceGet) | **Get** /billing/project/{projectId}/service/{serviceId} | Get billing/service
[**BillingProjectServiceList**](BillingProjectServiceApi.md#BillingProjectServiceList) | **Get** /billing/project/{projectId}/service | List billing/service



## BillingProjectServiceGet

> Service BillingProjectServiceGet(ctx, projectId, serviceId)

Get billing/service

Returns a single service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**serviceId** | **string**| Service Id | 

### Return type

[**Service**](service.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectServiceList

> []Service BillingProjectServiceList(ctx, projectId, optional)

List billing/service

List service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***BillingProjectServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingProjectServiceListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kind** | **optional.String**| Filter by kind | 
 **name** | **optional.String**| Filter by name | 
 **type_** | **optional.String**| Filter by type | 

### Return type

[**[]Service**](service.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

