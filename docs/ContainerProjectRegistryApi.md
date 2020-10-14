# \ContainerProjectRegistryApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContainerProjectRegistryCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCreate) | **Post** /container/{locationId}/project/{projectId}/registry | Create container/registry
[**ContainerProjectRegistryCredentialCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialCreate) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/credential | Create container/registry.credential
[**ContainerProjectRegistryCredentialDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/credential/{credentialId} | Delete container/registry.credential
[**ContainerProjectRegistryCredentialGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/credential/{credentialId} | Get container/registry.credential
[**ContainerProjectRegistryCredentialList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/credential | List container/registry.credential
[**ContainerProjectRegistryCredentialPatch**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialPatch) | **Patch** /container/{locationId}/project/{projectId}/registry/{registryId}/credential/{credentialId} | Update container/registry.credential
[**ContainerProjectRegistryDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId} | Delete container/registry
[**ContainerProjectRegistryDomainCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainCreate) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/domain | Create container/registry.domain
[**ContainerProjectRegistryDomainDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/domain/{domainId} | Delete container/registry.domain
[**ContainerProjectRegistryDomainGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/domain/{domainId} | Get container/registry.domain
[**ContainerProjectRegistryDomainList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/domain | List container/registry.domain
[**ContainerProjectRegistryEventGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryEventGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/event/{eventId} | Get container/registry.event
[**ContainerProjectRegistryEventList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryEventList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/event | List container/registry.event
[**ContainerProjectRegistryGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId} | Get container/registry
[**ContainerProjectRegistryList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryList) | **Get** /container/{locationId}/project/{projectId}/registry | List container/registry
[**ContainerProjectRegistryRepositoryGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId} | Get container/registry.repository
[**ContainerProjectRegistryRepositoryImageDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryImageDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId}/image/{imageId} | Delete container/registry.image
[**ContainerProjectRegistryRepositoryImageGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryImageGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId}/image/{imageId} | Get container/registry.image
[**ContainerProjectRegistryRepositoryImageList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryImageList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId}/image | List container/registry.image
[**ContainerProjectRegistryRepositoryList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository | List container/registry.repository
[**ContainerProjectRegistryServiceGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryServiceGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/service/{serviceId} | Get container/registry.service
[**ContainerProjectRegistryServiceList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryServiceList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/service | List container/registry.service
[**ContainerProjectRegistryStart**](ContainerProjectRegistryApi.md#ContainerProjectRegistryStart) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/actions/start | Start container/registry
[**ContainerProjectRegistryStop**](ContainerProjectRegistryApi.md#ContainerProjectRegistryStop) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/actions/stop | Stop container/registry
[**ContainerProjectRegistryTagCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagCreate) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/tag | Create container/registry.tag
[**ContainerProjectRegistryTagDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/tag/{tagId} | Delete container/registry.tag
[**ContainerProjectRegistryTagGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/tag/{tagId} | Get container/registry.tag
[**ContainerProjectRegistryTagList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/tag | List container/registry.tag
[**ContainerProjectRegistryTagPut**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagPut) | **Put** /container/{locationId}/project/{projectId}/registry/{registryId}/tag | Replace container/registry.tag
[**ContainerProjectRegistryTransfer**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTransfer) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/actions/transfer | Transfer container/registry
[**ContainerProjectRegistryUpdate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryUpdate) | **Patch** /container/{locationId}/project/{projectId}/registry/{registryId} | Update container/registry



## ContainerProjectRegistryCreate

> Registry ContainerProjectRegistryCreate(ctx, projectId, locationId, containerProjectRegistryCreate, optional)

Create container/registry

Create registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**containerProjectRegistryCreate** | [**ContainerProjectRegistryCreate**](ContainerProjectRegistryCreate.md)|  | 
 **optional** | ***ContainerProjectRegistryCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContainerProjectRegistryCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialCreate

> RegistryCredential ContainerProjectRegistryCredentialCreate(ctx, projectId, locationId, registryId, registryCredential)

Create container/registry.credential

Create container/registry.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**registryCredential** | [**RegistryCredential**](RegistryCredential.md)|  | 

### Return type

[**RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialDelete

> Registry ContainerProjectRegistryCredentialDelete(ctx, projectId, locationId, registryId, credentialId)

Delete container/registry.credential

Delete container/registry.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialGet

> RegistryCredential ContainerProjectRegistryCredentialGet(ctx, projectId, locationId, registryId, credentialId)

Get container/registry.credential

Get container/registry.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialList

> []RegistryCredential ContainerProjectRegistryCredentialList(ctx, projectId, locationId, registryId)

List container/registry.credential

List container/registry.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 

### Return type

[**[]RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialPatch

> RegistryCredential ContainerProjectRegistryCredentialPatch(ctx, projectId, locationId, registryId, credentialId, containerProjectRegistryCredentialPatch)

Update container/registry.credential

Update container/registry.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**credentialId** | **string**| credentialId | 
**containerProjectRegistryCredentialPatch** | [**ContainerProjectRegistryCredentialPatch**](ContainerProjectRegistryCredentialPatch.md)|  | 

### Return type

[**RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryDelete

> ContainerProjectRegistryDelete(ctx, projectId, locationId, registryId)

Delete container/registry

Delete registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 

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


## ContainerProjectRegistryDomainCreate

> Domain ContainerProjectRegistryDomainCreate(ctx, projectId, locationId, registryId, domain)

Create container/registry.domain

Create container/registry.domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**domain** | [**Domain**](Domain.md)|  | 

### Return type

[**Domain**](domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryDomainDelete

> ContainerProjectRegistryDomainDelete(ctx, projectId, locationId, registryId, domainId)

Delete container/registry.domain

Delete container/registry.domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**domainId** | **string**| domainId | 

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


## ContainerProjectRegistryDomainGet

> Domain ContainerProjectRegistryDomainGet(ctx, projectId, locationId, registryId, domainId)

Get container/registry.domain

Get container/registry.domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**domainId** | **string**| domainId | 

### Return type

[**Domain**](domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryDomainList

> []Domain ContainerProjectRegistryDomainList(ctx, projectId, locationId, registryId)

List container/registry.domain

List container/registry.domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 

### Return type

[**[]Domain**](domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryEventGet

> Event ContainerProjectRegistryEventGet(ctx, projectId, locationId, registryId, eventId)

Get container/registry.event

Get container/registry.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
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


## ContainerProjectRegistryEventList

> []Event ContainerProjectRegistryEventList(ctx, projectId, locationId, registryId, optional)

List container/registry.event

List container/registry.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
 **optional** | ***ContainerProjectRegistryEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContainerProjectRegistryEventListOpts struct


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


## ContainerProjectRegistryGet

> Registry ContainerProjectRegistryGet(ctx, projectId, locationId, registryId)

Get container/registry

Returns a single registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryList

> []Registry ContainerProjectRegistryList(ctx, projectId, locationId, optional)

List container/registry

List registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***ContainerProjectRegistryListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContainerProjectRegistryListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryRepositoryGet

> ContainerProjectRegistryRepositoryGet(ctx, projectId, locationId, registryId, repositoryId)

Get container/registry.repository

Get container/registry.repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**repositoryId** | **string**| repositoryId | 

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


## ContainerProjectRegistryRepositoryImageDelete

> ContainerProjectRegistryRepositoryImageDelete(ctx, projectId, locationId, registryId, repositoryId, imageId)

Delete container/registry.image

Delete container/registry.image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**repositoryId** | **string**| repositoryId | 
**imageId** | **string**| imageId | 

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


## ContainerProjectRegistryRepositoryImageGet

> ContainerProjectRegistryRepositoryImageGet(ctx, projectId, locationId, registryId, repositoryId, imageId)

Get container/registry.image

Get container/registry.image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**repositoryId** | **string**| repositoryId | 
**imageId** | **string**| imageId | 

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


## ContainerProjectRegistryRepositoryImageList

> ContainerProjectRegistryRepositoryImageList(ctx, projectId, locationId, registryId, repositoryId)

List container/registry.image

List container/registry.image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**repositoryId** | **string**| repositoryId | 

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


## ContainerProjectRegistryRepositoryList

> ContainerProjectRegistryRepositoryList(ctx, projectId, locationId, registryId)

List container/registry.repository

List container/registry.repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 

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


## ContainerProjectRegistryServiceGet

> ResourceService ContainerProjectRegistryServiceGet(ctx, projectId, locationId, registryId, serviceId)

Get container/registry.service

Get container/registry.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
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


## ContainerProjectRegistryServiceList

> []ResourceService ContainerProjectRegistryServiceList(ctx, projectId, locationId, registryId)

List container/registry.service

List container/registry.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 

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


## ContainerProjectRegistryStart

> Registry ContainerProjectRegistryStart(ctx, projectId, locationId, registryId, optional)

Start container/registry

action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
 **optional** | ***ContainerProjectRegistryStartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContainerProjectRegistryStartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryStop

> Registry ContainerProjectRegistryStop(ctx, projectId, locationId, registryId, optional)

Stop container/registry

action stop

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
 **optional** | ***ContainerProjectRegistryStopOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContainerProjectRegistryStopOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryTagCreate

> Tag ContainerProjectRegistryTagCreate(ctx, projectId, locationId, registryId, tag)

Create container/registry.tag

Create container/registry.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
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


## ContainerProjectRegistryTagDelete

> ContainerProjectRegistryTagDelete(ctx, projectId, locationId, registryId, tagId)

Delete container/registry.tag

Delete container/registry.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
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


## ContainerProjectRegistryTagGet

> Tag ContainerProjectRegistryTagGet(ctx, projectId, locationId, registryId, tagId)

Get container/registry.tag

Get container/registry.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
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


## ContainerProjectRegistryTagList

> []Tag ContainerProjectRegistryTagList(ctx, projectId, locationId, registryId)

List container/registry.tag

List container/registry.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 

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


## ContainerProjectRegistryTagPut

> []Tag ContainerProjectRegistryTagPut(ctx, projectId, locationId, registryId, tag)

Replace container/registry.tag

Replace container/registry.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
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


## ContainerProjectRegistryTransfer

> Registry ContainerProjectRegistryTransfer(ctx, projectId, locationId, registryId, containerProjectRegistryTransfer, optional)

Transfer container/registry

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**containerProjectRegistryTransfer** | [**ContainerProjectRegistryTransfer**](ContainerProjectRegistryTransfer.md)|  | 
 **optional** | ***ContainerProjectRegistryTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContainerProjectRegistryTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryUpdate

> Registry ContainerProjectRegistryUpdate(ctx, projectId, locationId, registryId, containerProjectRegistryUpdate)

Update container/registry

Returns modified registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**registryId** | **string**| Registry Id | 
**containerProjectRegistryUpdate** | [**ContainerProjectRegistryUpdate**](ContainerProjectRegistryUpdate.md)|  | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

