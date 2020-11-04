# \IamOrganisationPolicyApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamOrganisationPolicyActorCreate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorCreate) | **Post** /iam/organisation/{organisationId}/policy/{policyId}/actor | Create iam/policy.actor
[**IamOrganisationPolicyActorDelete**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorDelete) | **Delete** /iam/organisation/{organisationId}/policy/{policyId}/actor/{actorId} | Delete iam/policy.actor
[**IamOrganisationPolicyActorGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/actor/{actorId} | Get iam/policy.actor
[**IamOrganisationPolicyActorList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/actor | List iam/policy.actor
[**IamOrganisationPolicyCreate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyCreate) | **Post** /iam/organisation/{organisationId}/policy | Create iam/policy
[**IamOrganisationPolicyDelete**](IamOrganisationPolicyApi.md#IamOrganisationPolicyDelete) | **Delete** /iam/organisation/{organisationId}/policy/{policyId} | Delete iam/policy
[**IamOrganisationPolicyEventGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyEventGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/event/{eventId} | Get iam/policy.event
[**IamOrganisationPolicyEventList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyEventList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/event | List iam/policy.event
[**IamOrganisationPolicyGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId} | Get iam/policy
[**IamOrganisationPolicyList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyList) | **Get** /iam/organisation/{organisationId}/policy | List iam/policy
[**IamOrganisationPolicyServiceGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyServiceGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/service/{serviceId} | Get iam/policy.service
[**IamOrganisationPolicyServiceList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyServiceList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/service | List iam/policy.service
[**IamOrganisationPolicyTagCreate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagCreate) | **Post** /iam/organisation/{organisationId}/policy/{policyId}/tag | Create iam/policy.tag
[**IamOrganisationPolicyTagDelete**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagDelete) | **Delete** /iam/organisation/{organisationId}/policy/{policyId}/tag/{tagId} | Delete iam/policy.tag
[**IamOrganisationPolicyTagGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/tag/{tagId} | Get iam/policy.tag
[**IamOrganisationPolicyTagList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/tag | List iam/policy.tag
[**IamOrganisationPolicyTagPut**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagPut) | **Put** /iam/organisation/{organisationId}/policy/{policyId}/tag | Replace iam/policy.tag
[**IamOrganisationPolicyUpdate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyUpdate) | **Patch** /iam/organisation/{organisationId}/policy/{policyId} | Update iam/policy



## IamOrganisationPolicyActorCreate

> IamActor IamOrganisationPolicyActorCreate(ctx, organisationId, policyId, iamActor)

Create iam/policy.actor

Create iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyActorDelete

> IamActor IamOrganisationPolicyActorDelete(ctx, organisationId, policyId, actorId)

Delete iam/policy.actor

Delete iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyActorGet

> IamActor IamOrganisationPolicyActorGet(ctx, organisationId, policyId, actorId)

Get iam/policy.actor

Get iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyActorList

> []IamActor IamOrganisationPolicyActorList(ctx, organisationId, policyId)

List iam/policy.actor

List iam/policy.actor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyCreate

> Policy IamOrganisationPolicyCreate(ctx, organisationId, iamProjectPolicyCreate, optional)

Create iam/policy

Create policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**iamProjectPolicyCreate** | [**IamProjectPolicyCreate**](IamProjectPolicyCreate.md)|  | 
 **optional** | ***IamOrganisationPolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationPolicyCreateOpts struct


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


## IamOrganisationPolicyDelete

> IamOrganisationPolicyDelete(ctx, organisationId, policyId)

Delete iam/policy

Delete policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyEventGet

> Event IamOrganisationPolicyEventGet(ctx, organisationId, policyId, eventId)

Get iam/policy.event

Get iam/policy.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyEventList

> []Event IamOrganisationPolicyEventList(ctx, organisationId, policyId, optional)

List iam/policy.event

List iam/policy.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**policyId** | **string**| Policy Id | 
 **optional** | ***IamOrganisationPolicyEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationPolicyEventListOpts struct


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


## IamOrganisationPolicyGet

> Policy IamOrganisationPolicyGet(ctx, organisationId, policyId)

Get iam/policy

Returns a single policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyList

> []Policy IamOrganisationPolicyList(ctx, organisationId, optional)

List iam/policy

List policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***IamOrganisationPolicyListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationPolicyListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Filter by name | 
 **resource** | **optional.String**| Filter by resource | 
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


## IamOrganisationPolicyServiceGet

> ResourceService IamOrganisationPolicyServiceGet(ctx, organisationId, policyId, serviceId)

Get iam/policy.service

Get iam/policy.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyServiceList

> []ResourceService IamOrganisationPolicyServiceList(ctx, organisationId, policyId)

List iam/policy.service

List iam/policy.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyTagCreate

> Tag IamOrganisationPolicyTagCreate(ctx, organisationId, policyId, tag)

Create iam/policy.tag

Create iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyTagDelete

> IamOrganisationPolicyTagDelete(ctx, organisationId, policyId, tagId)

Delete iam/policy.tag

Delete iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyTagGet

> Tag IamOrganisationPolicyTagGet(ctx, organisationId, policyId, tagId)

Get iam/policy.tag

Get iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyTagList

> []Tag IamOrganisationPolicyTagList(ctx, organisationId, policyId)

List iam/policy.tag

List iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyTagPut

> []Tag IamOrganisationPolicyTagPut(ctx, organisationId, policyId, tag)

Replace iam/policy.tag

Replace iam/policy.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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


## IamOrganisationPolicyUpdate

> Policy IamOrganisationPolicyUpdate(ctx, organisationId, policyId, iamProjectPolicyUpdate)

Update iam/policy

Returns modified policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
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

