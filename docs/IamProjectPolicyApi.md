# \IamProjectPolicyApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectPolicyActorCreate**](IamProjectPolicyApi.md#IamProjectPolicyActorCreate) | **Post** /iam/project/{projectId}/policy/{policyId}/actor | Create iam/policy.actor
[**IamProjectPolicyActorDelete**](IamProjectPolicyApi.md#IamProjectPolicyActorDelete) | **Delete** /iam/project/{projectId}/policy/{policyId}/actor/{actorId} | Delete iam/policy.actor
[**IamProjectPolicyActorGet**](IamProjectPolicyApi.md#IamProjectPolicyActorGet) | **Get** /iam/project/{projectId}/policy/{policyId}/actor/{actorId} | Get iam/policy.actor
[**IamProjectPolicyActorList**](IamProjectPolicyApi.md#IamProjectPolicyActorList) | **Get** /iam/project/{projectId}/policy/{policyId}/actor | List iam/policy.actor
[**IamProjectPolicyCreate**](IamProjectPolicyApi.md#IamProjectPolicyCreate) | **Post** /iam/project/{projectId}/policy | Create iam/policy
[**IamProjectPolicyDelete**](IamProjectPolicyApi.md#IamProjectPolicyDelete) | **Delete** /iam/project/{projectId}/policy/{policyId} | Delete iam/policy
[**IamProjectPolicyEventGet**](IamProjectPolicyApi.md#IamProjectPolicyEventGet) | **Get** /iam/project/{projectId}/policy/{policyId}/event/{eventId} | Get iam/policy.event
[**IamProjectPolicyEventList**](IamProjectPolicyApi.md#IamProjectPolicyEventList) | **Get** /iam/project/{projectId}/policy/{policyId}/event | List iam/policy.event
[**IamProjectPolicyGet**](IamProjectPolicyApi.md#IamProjectPolicyGet) | **Get** /iam/project/{projectId}/policy/{policyId} | Get iam/policy
[**IamProjectPolicyList**](IamProjectPolicyApi.md#IamProjectPolicyList) | **Get** /iam/project/{projectId}/policy | List iam/policy
[**IamProjectPolicyServiceGet**](IamProjectPolicyApi.md#IamProjectPolicyServiceGet) | **Get** /iam/project/{projectId}/policy/{policyId}/service/{serviceId} | Get iam/policy.service
[**IamProjectPolicyServiceList**](IamProjectPolicyApi.md#IamProjectPolicyServiceList) | **Get** /iam/project/{projectId}/policy/{policyId}/service | List iam/policy.service
[**IamProjectPolicyTagCreate**](IamProjectPolicyApi.md#IamProjectPolicyTagCreate) | **Post** /iam/project/{projectId}/policy/{policyId}/tag | Create iam/policy.tag
[**IamProjectPolicyTagDelete**](IamProjectPolicyApi.md#IamProjectPolicyTagDelete) | **Delete** /iam/project/{projectId}/policy/{policyId}/tag/{tagId} | Delete iam/policy.tag
[**IamProjectPolicyTagGet**](IamProjectPolicyApi.md#IamProjectPolicyTagGet) | **Get** /iam/project/{projectId}/policy/{policyId}/tag/{tagId} | Get iam/policy.tag
[**IamProjectPolicyTagList**](IamProjectPolicyApi.md#IamProjectPolicyTagList) | **Get** /iam/project/{projectId}/policy/{policyId}/tag | List iam/policy.tag
[**IamProjectPolicyTagPut**](IamProjectPolicyApi.md#IamProjectPolicyTagPut) | **Put** /iam/project/{projectId}/policy/{policyId}/tag | Replace iam/policy.tag
[**IamProjectPolicyUpdate**](IamProjectPolicyApi.md#IamProjectPolicyUpdate) | **Patch** /iam/project/{projectId}/policy/{policyId} | Update iam/policy



## IamProjectPolicyActorCreate

> IamActor IamProjectPolicyActorCreate(ctx, projectId, policyId, iamActor)

Create iam/policy.actor

Create iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
**iamActor** | [**IamActor**](IamActor.md)|  | 

### Return type

[**IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyActorDelete

> IamActor IamProjectPolicyActorDelete(ctx, projectId, policyId, actorId)

Delete iam/policy.actor

Delete iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
**actorId** | **string**| actorId | 

### Return type

[**IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyActorGet

> IamActor IamProjectPolicyActorGet(ctx, projectId, policyId, actorId)

Get iam/policy.actor

Get iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
**actorId** | **string**| actorId | 

### Return type

[**IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyActorList

> []IamActor IamProjectPolicyActorList(ctx, projectId, policyId)

List iam/policy.actor

List iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 

### Return type

[**[]IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyCreate

> Policy IamProjectPolicyCreate(ctx, projectId, iamProjectPolicyCreate, optional)

Create iam/policy

Create policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**iamProjectPolicyCreate** | [**IamProjectPolicyCreate**](IamProjectPolicyCreate.md)|  | 
 **optional** | ***IamProjectPolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectPolicyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyDelete

> IamProjectPolicyDelete(ctx, projectId, policyId)

Delete iam/policy

Delete policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 

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


## IamProjectPolicyEventGet

> Event IamProjectPolicyEventGet(ctx, projectId, policyId, eventId)

Get iam/policy.event

Get iam/policy.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
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


## IamProjectPolicyEventList

> []Event IamProjectPolicyEventList(ctx, projectId, policyId, optional)

List iam/policy.event

List iam/policy.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
 **optional** | ***IamProjectPolicyEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectPolicyEventListOpts struct


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


## IamProjectPolicyGet

> Policy IamProjectPolicyGet(ctx, projectId, policyId)

Get iam/policy

Returns a single policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 

### Return type

[**Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyList

> []Policy IamProjectPolicyList(ctx, projectId, optional)

List iam/policy

List policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***IamProjectPolicyListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectPolicyListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyServiceGet

> ResourceService IamProjectPolicyServiceGet(ctx, projectId, policyId, serviceId)

Get iam/policy.service

Get iam/policy.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
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


## IamProjectPolicyServiceList

> []ResourceService IamProjectPolicyServiceList(ctx, projectId, policyId)

List iam/policy.service

List iam/policy.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 

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


## IamProjectPolicyTagCreate

> Tag IamProjectPolicyTagCreate(ctx, projectId, policyId, tag)

Create iam/policy.tag

Create iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
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


## IamProjectPolicyTagDelete

> IamProjectPolicyTagDelete(ctx, projectId, policyId, tagId)

Delete iam/policy.tag

Delete iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
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


## IamProjectPolicyTagGet

> Tag IamProjectPolicyTagGet(ctx, projectId, policyId, tagId)

Get iam/policy.tag

Get iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
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


## IamProjectPolicyTagList

> []Tag IamProjectPolicyTagList(ctx, projectId, policyId)

List iam/policy.tag

List iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 

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


## IamProjectPolicyTagPut

> []Tag IamProjectPolicyTagPut(ctx, projectId, policyId, tag)

Replace iam/policy.tag

Replace iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
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


## IamProjectPolicyUpdate

> Policy IamProjectPolicyUpdate(ctx, projectId, policyId, iamProjectPolicyUpdate)

Update iam/policy

Returns modified policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**policyId** | **string**| Policy Id | 
**iamProjectPolicyUpdate** | [**IamProjectPolicyUpdate**](IamProjectPolicyUpdate.md)|  | 

### Return type

[**Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

