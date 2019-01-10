# \ReplicaApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionReplicaImage**](ReplicaApi.md#ActionReplicaImage) | **Post** /replica/{replicaId}/actions/image | /actions/image
[**CreateReplica**](ReplicaApi.md#CreateReplica) | **Post** /replica | Create
[**DeleteReplica**](ReplicaApi.md#DeleteReplica) | **Delete** /replica/{replicaId} | Delete
[**ListReplica**](ReplicaApi.md#ListReplica) | **Get** /replica | List
[**OperationReplicaDeleteaccessrightsIdentity**](ReplicaApi.md#OperationReplicaDeleteaccessrightsIdentity) | **Delete** /replica/{replicaId}/accessrights/{identity} | /accessrights/:identity
[**OperationReplicaDeletetagKey**](ReplicaApi.md#OperationReplicaDeletetagKey) | **Delete** /replica/{replicaId}/tag/{key} | /tag/:key
[**OperationReplicaGetservicesServiceId**](ReplicaApi.md#OperationReplicaGetservicesServiceId) | **Get** /replica/{replicaId}/services/{serviceId} | /services/:serviceId
[**OperationReplicaGettag**](ReplicaApi.md#OperationReplicaGettag) | **Get** /replica/{replicaId}/tag/ | /tag/
[**OperationReplicaListaccessrights**](ReplicaApi.md#OperationReplicaListaccessrights) | **Get** /replica/{replicaId}/accessrights/ | /accessrights/
[**OperationReplicaListqueue**](ReplicaApi.md#OperationReplicaListqueue) | **Get** /replica/{replicaId}/queue/ | /queue/
[**OperationReplicaListservices**](ReplicaApi.md#OperationReplicaListservices) | **Get** /replica/{replicaId}/services/ | /services/
[**OperationReplicaPatchtag**](ReplicaApi.md#OperationReplicaPatchtag) | **Patch** /replica/{replicaId}/tag/ | /tag/
[**OperationReplicaPostaccessrights**](ReplicaApi.md#OperationReplicaPostaccessrights) | **Post** /replica/{replicaId}/accessrights/ | /accessrights/
[**ShowReplica**](ReplicaApi.md#ShowReplica) | **Get** /replica/{replicaId} | Get


# **ActionReplicaImage**
> Replica ActionReplicaImage(ctx, replicaId, optional)
/actions/image

Action image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 
 **optional** | ***ActionReplicaImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionReplicaImageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject72** | [**optional.Interface of InlineObject72**](InlineObject72.md)|  | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReplica**
> Replica CreateReplica(ctx, optional)
Create

Create replica

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateReplicaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateReplicaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject71** | [**optional.Interface of InlineObject71**](InlineObject71.md)|  | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReplica**
> DeleteReplica(ctx, replicaId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReplica**
> []Replica ListReplica(ctx, optional)
List

List replica

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListReplicaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListReplicaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaDeleteaccessrightsIdentity**
> Replica OperationReplicaDeleteaccessrightsIdentity(ctx, replicaId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 
  **identity** | **string**| identity | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaDeletetagKey**
> map[string]interface{} OperationReplicaDeletetagKey(ctx, replicaId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaGetservicesServiceId**
> ReplicaServices OperationReplicaGetservicesServiceId(ctx, replicaId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 
  **serviceId** | **string**| serviceId | 

### Return type

[**ReplicaServices**](replica.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaGettag**
> map[string]interface{} OperationReplicaGettag(ctx, replicaId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaListaccessrights**
> []string OperationReplicaListaccessrights(ctx, replicaId)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaListqueue**
> []Event OperationReplicaListqueue(ctx, replicaId)
/queue/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaListservices**
> []ReplicaServices OperationReplicaListservices(ctx, replicaId)
/services/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 

### Return type

[**[]ReplicaServices**](replica.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaPatchtag**
> map[string]interface{} OperationReplicaPatchtag(ctx, replicaId, optional)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 
 **optional** | ***OperationReplicaPatchtagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationReplicaPatchtagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.Map[string]interface{}**|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationReplicaPostaccessrights**
> string OperationReplicaPostaccessrights(ctx, replicaId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 
 **optional** | ***OperationReplicaPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationReplicaPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject73** | [**optional.Interface of InlineObject73**](InlineObject73.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowReplica**
> Replica ShowReplica(ctx, replicaId)
Get

Returns a single replica

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **replicaId** | **string**| ID of replica | 

### Return type

[**Replica**](replica.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

