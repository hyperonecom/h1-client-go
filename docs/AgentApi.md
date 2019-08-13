# \AgentApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentActionStart**](AgentApi.md#AgentActionStart) | **Post** /agent/{agentId}/actions/start | /actions/start
[**AgentActionSuspend**](AgentApi.md#AgentActionSuspend) | **Post** /agent/{agentId}/actions/suspend | /actions/suspend
[**AgentActionTransfer**](AgentApi.md#AgentActionTransfer) | **Post** /agent/{agentId}/actions/transfer | /actions/transfer
[**AgentActionUpdateEnabledServices**](AgentApi.md#AgentActionUpdateEnabledServices) | **Post** /agent/{agentId}/actions/update_enabledServices | /actions/update_enabledServices
[**AgentCreate**](AgentApi.md#AgentCreate) | **Post** /agent | Create
[**AgentDelete**](AgentApi.md#AgentDelete) | **Delete** /agent/{agentId} | Delete
[**AgentDeleteAccessrightsIdentity**](AgentApi.md#AgentDeleteAccessrightsIdentity) | **Delete** /agent/{agentId}/accessrights/{identity} | /accessrights/:identity
[**AgentDeleteCredentialId**](AgentApi.md#AgentDeleteCredentialId) | **Delete** /agent/{agentId}/credential/{id} | /credential/:id
[**AgentDeleteCredentialcertificateId**](AgentApi.md#AgentDeleteCredentialcertificateId) | **Delete** /agent/{agentId}/credential/certificate/{id} | /credential/certificate/:id
[**AgentDeleteTagKey**](AgentApi.md#AgentDeleteTagKey) | **Delete** /agent/{agentId}/tag/{key} | /tag/:key
[**AgentGetCredentialId**](AgentApi.md#AgentGetCredentialId) | **Get** /agent/{agentId}/credential/{id} | /credential/:id
[**AgentGetCredentialcertificateId**](AgentApi.md#AgentGetCredentialcertificateId) | **Get** /agent/{agentId}/credential/certificate/{id} | /credential/certificate/:id
[**AgentGetInspect**](AgentApi.md#AgentGetInspect) | **Get** /agent/{agentId}/inspect | /inspect
[**AgentGetResourceResourceIdinspect**](AgentApi.md#AgentGetResourceResourceIdinspect) | **Get** /agent/{agentId}/resource/{resourceId}/inspect | /resource/:resourceId/inspect
[**AgentGetServicesServiceId**](AgentApi.md#AgentGetServicesServiceId) | **Get** /agent/{agentId}/services/{serviceId} | /services/:serviceId
[**AgentGetTag**](AgentApi.md#AgentGetTag) | **Get** /agent/{agentId}/tag | /tag
[**AgentList**](AgentApi.md#AgentList) | **Get** /agent | List
[**AgentListAccessrights**](AgentApi.md#AgentListAccessrights) | **Get** /agent/{agentId}/accessrights | /accessrights
[**AgentListCredential**](AgentApi.md#AgentListCredential) | **Get** /agent/{agentId}/credential | /credential
[**AgentListCredentialcertificate**](AgentApi.md#AgentListCredentialcertificate) | **Get** /agent/{agentId}/credential/certificate | /credential/certificate
[**AgentListEnabledServices**](AgentApi.md#AgentListEnabledServices) | **Get** /agent/{agentId}/enabledServices | /enabledServices
[**AgentListQueue**](AgentApi.md#AgentListQueue) | **Get** /agent/{agentId}/queue | /queue
[**AgentListResource**](AgentApi.md#AgentListResource) | **Get** /agent/{agentId}/resource | /resource
[**AgentListResourceResourceIdqueue**](AgentApi.md#AgentListResourceResourceIdqueue) | **Get** /agent/{agentId}/resource/{resourceId}/queue | /resource/:resourceId/queue
[**AgentListServices**](AgentApi.md#AgentListServices) | **Get** /agent/{agentId}/services | /services
[**AgentPatchCredentialId**](AgentApi.md#AgentPatchCredentialId) | **Patch** /agent/{agentId}/credential/{id} | /credential/:id
[**AgentPatchCredentialcertificateId**](AgentApi.md#AgentPatchCredentialcertificateId) | **Patch** /agent/{agentId}/credential/certificate/{id} | /credential/certificate/:id
[**AgentPatchTag**](AgentApi.md#AgentPatchTag) | **Patch** /agent/{agentId}/tag | /tag
[**AgentPostAccessrights**](AgentApi.md#AgentPostAccessrights) | **Post** /agent/{agentId}/accessrights | /accessrights
[**AgentPostCredential**](AgentApi.md#AgentPostCredential) | **Post** /agent/{agentId}/credential | /credential
[**AgentPostCredentialcertificate**](AgentApi.md#AgentPostCredentialcertificate) | **Post** /agent/{agentId}/credential/certificate | /credential/certificate
[**AgentPostResourceResourceIdactionsrecreate**](AgentApi.md#AgentPostResourceResourceIdactionsrecreate) | **Post** /agent/{agentId}/resource/{resourceId}/actions/recreate | /resource/:resourceId/actions/recreate
[**AgentPutEnabledServices**](AgentApi.md#AgentPutEnabledServices) | **Put** /agent/{agentId}/enabledServices | /enabledServices
[**AgentShow**](AgentApi.md#AgentShow) | **Get** /agent/{agentId} | Get
[**AgentUpdate**](AgentApi.md#AgentUpdate) | **Patch** /agent/{agentId} | Update



## AgentActionStart

> Agent AgentActionStart(ctx, agentId)
/actions/start

Action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentActionSuspend

> Agent AgentActionSuspend(ctx, agentId)
/actions/suspend

Action suspend

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentActionTransfer

> Agent AgentActionTransfer(ctx, agentId, agentActionTransfer)
/actions/transfer

Action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**agentActionTransfer** | [**AgentActionTransfer**](AgentActionTransfer.md)|  | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentActionUpdateEnabledServices

> Agent AgentActionUpdateEnabledServices(ctx, agentId, agentActionUpdateEnabledServices)
/actions/update_enabledServices

Action update_enabledServices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**agentActionUpdateEnabledServices** | [**AgentActionUpdateEnabledServices**](AgentActionUpdateEnabledServices.md)|  | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentCreate

> Agent AgentCreate(ctx, agentCreate)
Create

Create agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentCreate** | [**AgentCreate**](AgentCreate.md)|  | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentDelete

> AgentDelete(ctx, agentId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentDeleteAccessrightsIdentity

> Agent AgentDeleteAccessrightsIdentity(ctx, agentId, identity)
/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**identity** | **string**| identity | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentDeleteCredentialId

> Agent AgentDeleteCredentialId(ctx, agentId, id)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**id** | **string**| id | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentDeleteCredentialcertificateId

> CredentialCertificate AgentDeleteCredentialcertificateId(ctx, agentId, id)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**id** | **string**| id | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentDeleteTagKey

> map[string]string AgentDeleteTagKey(ctx, agentId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
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


## AgentGetCredentialId

> CredentialPassword AgentGetCredentialId(ctx, agentId, id)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**id** | **string**| id | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetCredentialcertificateId

> CredentialCertificate AgentGetCredentialcertificateId(ctx, agentId, id)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**id** | **string**| id | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetInspect

> map[string]interface{} AgentGetInspect(ctx, agentId)
/inspect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetResourceResourceIdinspect

> map[string]interface{} AgentGetResourceResourceIdinspect(ctx, agentId, resourceId)
/resource/:resourceId/inspect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**resourceId** | **string**| resourceId | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetServicesServiceId

> AgentServices AgentGetServicesServiceId(ctx, agentId, serviceId)
/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**serviceId** | **string**| serviceId | 

### Return type

[**AgentServices**](agent.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetTag

> map[string]string AgentGetTag(ctx, agentId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

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


## AgentList

> []Agent AgentList(ctx, optional)
List

List agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentListAccessrights

> []string AgentListAccessrights(ctx, agentId)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

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


## AgentListCredential

> []CredentialPassword AgentListCredential(ctx, agentId)
/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**[]CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentListCredentialcertificate

> []CredentialCertificate AgentListCredentialcertificate(ctx, agentId)
/credential/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**[]CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentListEnabledServices

> []string AgentListEnabledServices(ctx, agentId)
/enabledServices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

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


## AgentListQueue

> []Event AgentListQueue(ctx, agentId)
/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

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


## AgentListResource

> []AgentResource AgentListResource(ctx, agentId)
/resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**[]AgentResource**](agentResource.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentListResourceResourceIdqueue

> []AgentResourceEvent AgentListResourceResourceIdqueue(ctx, agentId, resourceId)
/resource/:resourceId/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**resourceId** | **string**| resourceId | 

### Return type

[**[]AgentResourceEvent**](agentResourceEvent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentListServices

> []AgentServices AgentListServices(ctx, agentId)
/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**[]AgentServices**](agent.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPatchCredentialId

> CredentialPassword AgentPatchCredentialId(ctx, agentId, id, agentPatchCredentialId)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**id** | **string**| id | 
**agentPatchCredentialId** | [**AgentPatchCredentialId**](AgentPatchCredentialId.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPatchCredentialcertificateId

> CredentialCertificate AgentPatchCredentialcertificateId(ctx, agentId, id, agentPatchCredentialcertificateId)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**id** | **string**| id | 
**agentPatchCredentialcertificateId** | [**AgentPatchCredentialcertificateId**](AgentPatchCredentialcertificateId.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPatchTag

> map[string]string AgentPatchTag(ctx, agentId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
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


## AgentPostAccessrights

> Agent AgentPostAccessrights(ctx, agentId, agentPostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**agentPostAccessrights** | [**AgentPostAccessrights**](AgentPostAccessrights.md)|  | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPostCredential

> CredentialPassword AgentPostCredential(ctx, agentId, agentPostCredential)
/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**agentPostCredential** | [**AgentPostCredential**](AgentPostCredential.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPostCredentialcertificate

> CredentialCertificate AgentPostCredentialcertificate(ctx, agentId, agentPostCredentialcertificate)
/credential/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**agentPostCredentialcertificate** | [**AgentPostCredentialcertificate**](AgentPostCredentialcertificate.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPostResourceResourceIdactionsrecreate

> AgentResource AgentPostResourceResourceIdactionsrecreate(ctx, agentId, resourceId)
/resource/:resourceId/actions/recreate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**resourceId** | **string**| resourceId | 

### Return type

[**AgentResource**](agentResource.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPutEnabledServices

> []string AgentPutEnabledServices(ctx, agentId, requestBody)
/enabledServices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**requestBody** | [**[]string**](string.md)|  | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentShow

> Agent AgentShow(ctx, agentId)
Get

Returns a single agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentUpdate

> Agent AgentUpdate(ctx, agentId, agentUpdate)
Update

Returns modified agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string**| ID of agent | 
**agentUpdate** | [**AgentUpdate**](AgentUpdate.md)|  | 

### Return type

[**Agent**](agent.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

