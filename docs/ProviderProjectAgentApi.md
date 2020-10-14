# \ProviderProjectAgentApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProviderProjectAgentCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentCreate) | **Post** /provider/{locationId}/project/{projectId}/agent | Create provider/agent
[**ProviderProjectAgentCredentialCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialCreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential | Create provider/agent.credential
[**ProviderProjectAgentCredentialDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential/{credentialId} | Delete provider/agent.credential
[**ProviderProjectAgentCredentialGet**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential/{credentialId} | Get provider/agent.credential
[**ProviderProjectAgentCredentialList**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential | List provider/agent.credential
[**ProviderProjectAgentCredentialPatch**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialPatch) | **Patch** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential/{credentialId} | Update provider/agent.credential
[**ProviderProjectAgentDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId} | Delete provider/agent
[**ProviderProjectAgentEnabledServiceCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceCreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService | Create provider/agent.enabledService
[**ProviderProjectAgentEnabledServiceDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService/{enabledServiceId} | Delete provider/agent.enabledService
[**ProviderProjectAgentEnabledServiceGet**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService/{enabledServiceId} | Get provider/agent.enabledService
[**ProviderProjectAgentEnabledServiceList**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService | List provider/agent.enabledService
[**ProviderProjectAgentEventGet**](ProviderProjectAgentApi.md#ProviderProjectAgentEventGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/event/{eventId} | Get provider/agent.event
[**ProviderProjectAgentEventList**](ProviderProjectAgentApi.md#ProviderProjectAgentEventList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/event | List provider/agent.event
[**ProviderProjectAgentGet**](ProviderProjectAgentApi.md#ProviderProjectAgentGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId} | Get provider/agent
[**ProviderProjectAgentInspect**](ProviderProjectAgentApi.md#ProviderProjectAgentInspect) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/inspect | Inspect provider/agent
[**ProviderProjectAgentList**](ProviderProjectAgentApi.md#ProviderProjectAgentList) | **Get** /provider/{locationId}/project/{projectId}/agent | List provider/agent
[**ProviderProjectAgentResourceEventList**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceEventList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource/{resourceId}/event | List provider/agent.event
[**ProviderProjectAgentResourceInspect**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceInspect) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource/{resourceId}/actions/inspect | Inspect provider/agent.resource
[**ProviderProjectAgentResourceList**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource | List provider/agent.resource
[**ProviderProjectAgentResourceRecreate**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceRecreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource/{resourceId}/actions/recreate | Recreate provider/agent.resource
[**ProviderProjectAgentServiceGet**](ProviderProjectAgentApi.md#ProviderProjectAgentServiceGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/service/{serviceId} | Get provider/agent.service
[**ProviderProjectAgentServiceList**](ProviderProjectAgentApi.md#ProviderProjectAgentServiceList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/service | List provider/agent.service
[**ProviderProjectAgentStart**](ProviderProjectAgentApi.md#ProviderProjectAgentStart) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/start | Start provider/agent
[**ProviderProjectAgentSuspend**](ProviderProjectAgentApi.md#ProviderProjectAgentSuspend) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/suspend | Suspend provider/agent
[**ProviderProjectAgentTagCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentTagCreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag | Create provider/agent.tag
[**ProviderProjectAgentTagDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentTagDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag/{tagId} | Delete provider/agent.tag
[**ProviderProjectAgentTagGet**](ProviderProjectAgentApi.md#ProviderProjectAgentTagGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag/{tagId} | Get provider/agent.tag
[**ProviderProjectAgentTagList**](ProviderProjectAgentApi.md#ProviderProjectAgentTagList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag | List provider/agent.tag
[**ProviderProjectAgentTagPut**](ProviderProjectAgentApi.md#ProviderProjectAgentTagPut) | **Put** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag | Replace provider/agent.tag
[**ProviderProjectAgentTransfer**](ProviderProjectAgentApi.md#ProviderProjectAgentTransfer) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/transfer | Transfer provider/agent
[**ProviderProjectAgentUpdate**](ProviderProjectAgentApi.md#ProviderProjectAgentUpdate) | **Patch** /provider/{locationId}/project/{projectId}/agent/{agentId} | Update provider/agent



## ProviderProjectAgentCreate

> Agent ProviderProjectAgentCreate(ctx, projectId, locationId, providerProjectAgentCreate, optional)

Create provider/agent

Create agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**providerProjectAgentCreate** | [**ProviderProjectAgentCreate**](ProviderProjectAgentCreate.md)|  | 
 **optional** | ***ProviderProjectAgentCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialCreate

> AgentCredential ProviderProjectAgentCredentialCreate(ctx, projectId, locationId, agentId, agentCredential)

Create provider/agent.credential

Create provider/agent.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**agentCredential** | [**AgentCredential**](AgentCredential.md)|  | 

### Return type

[**AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialDelete

> Agent ProviderProjectAgentCredentialDelete(ctx, projectId, locationId, agentId, credentialId)

Delete provider/agent.credential

Delete provider/agent.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialGet

> AgentCredential ProviderProjectAgentCredentialGet(ctx, projectId, locationId, agentId, credentialId)

Get provider/agent.credential

Get provider/agent.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialList

> []AgentCredential ProviderProjectAgentCredentialList(ctx, projectId, locationId, agentId)

List provider/agent.credential

List provider/agent.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 

### Return type

[**[]AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialPatch

> AgentCredential ProviderProjectAgentCredentialPatch(ctx, projectId, locationId, agentId, credentialId, providerProjectAgentCredentialPatch)

Update provider/agent.credential

Update provider/agent.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**credentialId** | **string**| credentialId | 
**providerProjectAgentCredentialPatch** | [**ProviderProjectAgentCredentialPatch**](ProviderProjectAgentCredentialPatch.md)|  | 

### Return type

[**AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentDelete

> ProviderProjectAgentDelete(ctx, projectId, locationId, agentId)

Delete provider/agent

Delete agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEnabledServiceCreate

> EnabledService ProviderProjectAgentEnabledServiceCreate(ctx, projectId, locationId, agentId, enabledService)

Create provider/agent.enabledService

Create provider/agent.enabledService

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**enabledService** | [**EnabledService**](EnabledService.md)|  | 

### Return type

[**EnabledService**](enabledService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEnabledServiceDelete

> Agent ProviderProjectAgentEnabledServiceDelete(ctx, projectId, locationId, agentId, enabledServiceId)

Delete provider/agent.enabledService

Delete provider/agent.enabledService

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**enabledServiceId** | **string**| enabledServiceId | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEnabledServiceGet

> EnabledService ProviderProjectAgentEnabledServiceGet(ctx, projectId, locationId, agentId, enabledServiceId)

Get provider/agent.enabledService

Get provider/agent.enabledService

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**enabledServiceId** | **string**| enabledServiceId | 

### Return type

[**EnabledService**](enabledService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEnabledServiceList

> []EnabledService ProviderProjectAgentEnabledServiceList(ctx, projectId, locationId, agentId)

List provider/agent.enabledService

List provider/agent.enabledService

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 

### Return type

[**[]EnabledService**](enabledService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEventGet

> Event ProviderProjectAgentEventGet(ctx, projectId, locationId, agentId, eventId)

Get provider/agent.event

Get provider/agent.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**eventId** | **string**| eventId | 

### Return type

[**Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEventList

> []Event ProviderProjectAgentEventList(ctx, projectId, locationId, agentId, optional)

List provider/agent.event

List provider/agent.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
 **optional** | ***ProviderProjectAgentEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentEventListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Float32**| $limit | [default to 100]
 **skip** | **optional.Float32**| $skip | 

### Return type

[**[]Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentGet

> Agent ProviderProjectAgentGet(ctx, projectId, locationId, agentId)

Get provider/agent

Returns a single agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentInspect

> Agent ProviderProjectAgentInspect(ctx, projectId, locationId, agentId, optional)

Inspect provider/agent

action inspect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
 **optional** | ***ProviderProjectAgentInspectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentInspectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentList

> []Agent ProviderProjectAgentList(ctx, projectId, locationId, optional)

List provider/agent

List agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***ProviderProjectAgentListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceEventList

> []AgentResourceEvent ProviderProjectAgentResourceEventList(ctx, projectId, locationId, agentId, resourceId, optional)

List provider/agent.event

List provider/agent.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**resourceId** | **string**| resourceId | 
 **optional** | ***ProviderProjectAgentResourceEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentResourceEventListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float32**| $limit | 
 **skip** | **optional.Float32**| $skip | 

### Return type

[**[]AgentResourceEvent**](agentResourceEvent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceInspect

> interface{} ProviderProjectAgentResourceInspect(ctx, projectId, locationId, agentId, resourceId)

Inspect provider/agent.resource

action inspect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**resourceId** | **string**| resourceId | 

### Return type

**interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceList

> []AgentResource ProviderProjectAgentResourceList(ctx, projectId, locationId, agentId)

List provider/agent.resource

List provider/agent.resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 

### Return type

[**[]AgentResource**](agentResource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceRecreate

> AgentResource ProviderProjectAgentResourceRecreate(ctx, projectId, locationId, agentId, resourceId)

Recreate provider/agent.resource

action recreate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**resourceId** | **string**| resourceId | 

### Return type

[**AgentResource**](agentResource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentServiceGet

> ResourceService ProviderProjectAgentServiceGet(ctx, projectId, locationId, agentId, serviceId)

Get provider/agent.service

Get provider/agent.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**serviceId** | **string**| serviceId | 

### Return type

[**ResourceService**](resourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentServiceList

> []ResourceService ProviderProjectAgentServiceList(ctx, projectId, locationId, agentId)

List provider/agent.service

List provider/agent.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 

### Return type

[**[]ResourceService**](resourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentStart

> Agent ProviderProjectAgentStart(ctx, projectId, locationId, agentId, optional)

Start provider/agent

action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
 **optional** | ***ProviderProjectAgentStartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentStartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentSuspend

> Agent ProviderProjectAgentSuspend(ctx, projectId, locationId, agentId, optional)

Suspend provider/agent

action suspend

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
 **optional** | ***ProviderProjectAgentSuspendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentSuspendOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentTagCreate

> Tag ProviderProjectAgentTagCreate(ctx, projectId, locationId, agentId, tag)

Create provider/agent.tag

Create provider/agent.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**tag** | [**Tag**](Tag.md)|  | 

### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentTagDelete

> ProviderProjectAgentTagDelete(ctx, projectId, locationId, agentId, tagId)

Delete provider/agent.tag

Delete provider/agent.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**tagId** | **string**| tagId | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentTagGet

> Tag ProviderProjectAgentTagGet(ctx, projectId, locationId, agentId, tagId)

Get provider/agent.tag

Get provider/agent.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**tagId** | **string**| tagId | 

### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentTagList

> []Tag ProviderProjectAgentTagList(ctx, projectId, locationId, agentId)

List provider/agent.tag

List provider/agent.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 

### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentTagPut

> []Tag ProviderProjectAgentTagPut(ctx, projectId, locationId, agentId, tag)

Replace provider/agent.tag

Replace provider/agent.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**tag** | [**[]Tag**](tag.md)|  | 

### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentTransfer

> Agent ProviderProjectAgentTransfer(ctx, projectId, locationId, agentId, providerProjectAgentTransfer, optional)

Transfer provider/agent

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**providerProjectAgentTransfer** | [**ProviderProjectAgentTransfer**](ProviderProjectAgentTransfer.md)|  | 
 **optional** | ***ProviderProjectAgentTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProviderProjectAgentTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentUpdate

> Agent ProviderProjectAgentUpdate(ctx, projectId, locationId, agentId, providerProjectAgentUpdate)

Update provider/agent

Returns modified agent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**agentId** | **string**| Agent Id | 
**providerProjectAgentUpdate** | [**ProviderProjectAgentUpdate**](ProviderProjectAgentUpdate.md)|  | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

