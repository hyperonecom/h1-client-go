# \SnapshotApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotApi.md#CreateSnapshot) | **Post** /snapshot | Create
[**DeleteSnapshot**](SnapshotApi.md#DeleteSnapshot) | **Delete** /snapshot/{snapshotId} | Delete
[**ListSnapshot**](SnapshotApi.md#ListSnapshot) | **Get** /snapshot | List
[**OperationSnapshotDeleteaccessrightsIdentity**](SnapshotApi.md#OperationSnapshotDeleteaccessrightsIdentity) | **Delete** /snapshot/{snapshotId}/accessrights/{identity} | /accessrights/:identity
[**OperationSnapshotDeletetagKey**](SnapshotApi.md#OperationSnapshotDeletetagKey) | **Delete** /snapshot/{snapshotId}/tag/{key} | /tag/:key
[**OperationSnapshotGetservicesServiceId**](SnapshotApi.md#OperationSnapshotGetservicesServiceId) | **Get** /snapshot/{snapshotId}/services/{serviceId} | /services/:serviceId
[**OperationSnapshotGettag**](SnapshotApi.md#OperationSnapshotGettag) | **Get** /snapshot/{snapshotId}/tag/ | /tag/
[**OperationSnapshotListaccessrights**](SnapshotApi.md#OperationSnapshotListaccessrights) | **Get** /snapshot/{snapshotId}/accessrights/ | /accessrights/
[**OperationSnapshotListqueue**](SnapshotApi.md#OperationSnapshotListqueue) | **Get** /snapshot/{snapshotId}/queue/ | /queue/
[**OperationSnapshotListservices**](SnapshotApi.md#OperationSnapshotListservices) | **Get** /snapshot/{snapshotId}/services/ | /services/
[**OperationSnapshotPatchtag**](SnapshotApi.md#OperationSnapshotPatchtag) | **Patch** /snapshot/{snapshotId}/tag/ | /tag/
[**OperationSnapshotPostaccessrights**](SnapshotApi.md#OperationSnapshotPostaccessrights) | **Post** /snapshot/{snapshotId}/accessrights/ | /accessrights/
[**ShowSnapshot**](SnapshotApi.md#ShowSnapshot) | **Get** /snapshot/{snapshotId} | Get
[**UpdateSnapshot**](SnapshotApi.md#UpdateSnapshot) | **Patch** /snapshot/{snapshotId} | Update


# **CreateSnapshot**
> Snapshot CreateSnapshot(ctx, snapshotCreate)
Create

Create snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotCreate** | [**SnapshotCreate**](SnapshotCreate.md)|  | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshot**
> DeleteSnapshot(ctx, snapshotId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshot**
> []Snapshot ListSnapshot(ctx, optional)
List

List snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListSnapshotOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **vault** | **optional.String**| Filter by vault | 

### Return type

[**[]Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotDeleteaccessrightsIdentity**
> Snapshot OperationSnapshotDeleteaccessrightsIdentity(ctx, snapshotId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 
  **identity** | **string**| identity | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotDeletetagKey**
> map[string]string OperationSnapshotDeletetagKey(ctx, snapshotId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotGetservicesServiceId**
> SnapshotServices OperationSnapshotGetservicesServiceId(ctx, snapshotId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 
  **serviceId** | **string**| serviceId | 

### Return type

[**SnapshotServices**](snapshot.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotGettag**
> map[string]string OperationSnapshotGettag(ctx, snapshotId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotListaccessrights**
> []string OperationSnapshotListaccessrights(ctx, snapshotId)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotListqueue**
> []Event OperationSnapshotListqueue(ctx, snapshotId)
/queue/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotListservices**
> []SnapshotServices OperationSnapshotListservices(ctx, snapshotId)
/services/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 

### Return type

[**[]SnapshotServices**](snapshot.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotPatchtag**
> map[string]string OperationSnapshotPatchtag(ctx, snapshotId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationSnapshotPostaccessrights**
> string OperationSnapshotPostaccessrights(ctx, snapshotId, resourceAccessRight)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 
  **resourceAccessRight** | [**ResourceAccessRight**](ResourceAccessRight.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSnapshot**
> Snapshot ShowSnapshot(ctx, snapshotId)
Get

Returns a single snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnapshot**
> Snapshot UpdateSnapshot(ctx, snapshotId, optional)
Update

Returns modified snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| ID of snapshot | 
 **optional** | ***UpdateSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateSnapshotOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject29** | [**optional.Interface of InlineObject29**](InlineObject29.md)|  | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

