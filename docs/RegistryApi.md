# \RegistryApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryActionStart**](RegistryApi.md#RegistryActionStart) | **Post** /registry/{registryId}/actions/start | /actions/start
[**RegistryActionStop**](RegistryApi.md#RegistryActionStop) | **Post** /registry/{registryId}/actions/stop | /actions/stop
[**RegistryActionTransfer**](RegistryApi.md#RegistryActionTransfer) | **Post** /registry/{registryId}/actions/transfer | /actions/transfer
[**RegistryActionUpdateDomain**](RegistryApi.md#RegistryActionUpdateDomain) | **Post** /registry/{registryId}/actions/update_domain | /actions/update_domain
[**RegistryCreate**](RegistryApi.md#RegistryCreate) | **Post** /registry | Create
[**RegistryDelete**](RegistryApi.md#RegistryDelete) | **Delete** /registry/{registryId} | Delete
[**RegistryDeleteAccessrightsIdentity**](RegistryApi.md#RegistryDeleteAccessrightsIdentity) | **Delete** /registry/{registryId}/accessrights/{identity} | /accessrights/:identity
[**RegistryDeleteCredentialId**](RegistryApi.md#RegistryDeleteCredentialId) | **Delete** /registry/{registryId}/credential/{id} | /credential/:id
[**RegistryDeleteCredentialpasswordId**](RegistryApi.md#RegistryDeleteCredentialpasswordId) | **Delete** /registry/{registryId}/credential/password/{id} | /credential/password/:id
[**RegistryDeleteRepositoryRepositoryIdtagTagId**](RegistryApi.md#RegistryDeleteRepositoryRepositoryIdtagTagId) | **Delete** /registry/{registryId}/repository/{repositoryId}/tag/{tagId} | /repository/:repositoryId/tag/:tagId
[**RegistryDeleteTagKey**](RegistryApi.md#RegistryDeleteTagKey) | **Delete** /registry/{registryId}/tag/{key} | /tag/:key
[**RegistryGetCredentialId**](RegistryApi.md#RegistryGetCredentialId) | **Get** /registry/{registryId}/credential/{id} | /credential/:id
[**RegistryGetCredentialpasswordId**](RegistryApi.md#RegistryGetCredentialpasswordId) | **Get** /registry/{registryId}/credential/password/{id} | /credential/password/:id
[**RegistryGetRepositoryRepositoryId**](RegistryApi.md#RegistryGetRepositoryRepositoryId) | **Get** /registry/{registryId}/repository/{repositoryId} | /repository/:repositoryId
[**RegistryGetRepositoryRepositoryIdtagTagId**](RegistryApi.md#RegistryGetRepositoryRepositoryIdtagTagId) | **Get** /registry/{registryId}/repository/{repositoryId}/tag/{tagId} | /repository/:repositoryId/tag/:tagId
[**RegistryGetServicesServiceId**](RegistryApi.md#RegistryGetServicesServiceId) | **Get** /registry/{registryId}/services/{serviceId} | /services/:serviceId
[**RegistryGetTag**](RegistryApi.md#RegistryGetTag) | **Get** /registry/{registryId}/tag | /tag
[**RegistryList**](RegistryApi.md#RegistryList) | **Get** /registry | List
[**RegistryListAccessrights**](RegistryApi.md#RegistryListAccessrights) | **Get** /registry/{registryId}/accessrights | /accessrights
[**RegistryListCredential**](RegistryApi.md#RegistryListCredential) | **Get** /registry/{registryId}/credential | /credential
[**RegistryListCredentialpassword**](RegistryApi.md#RegistryListCredentialpassword) | **Get** /registry/{registryId}/credential/password | /credential/password
[**RegistryListQueue**](RegistryApi.md#RegistryListQueue) | **Get** /registry/{registryId}/queue | /queue
[**RegistryListRepository**](RegistryApi.md#RegistryListRepository) | **Get** /registry/{registryId}/repository | /repository
[**RegistryListRepositoryRepositoryIdtag**](RegistryApi.md#RegistryListRepositoryRepositoryIdtag) | **Get** /registry/{registryId}/repository/{repositoryId}/tag | /repository/:repositoryId/tag
[**RegistryListServices**](RegistryApi.md#RegistryListServices) | **Get** /registry/{registryId}/services | /services
[**RegistryPatchCredentialId**](RegistryApi.md#RegistryPatchCredentialId) | **Patch** /registry/{registryId}/credential/{id} | /credential/:id
[**RegistryPatchCredentialpasswordId**](RegistryApi.md#RegistryPatchCredentialpasswordId) | **Patch** /registry/{registryId}/credential/password/{id} | /credential/password/:id
[**RegistryPatchTag**](RegistryApi.md#RegistryPatchTag) | **Patch** /registry/{registryId}/tag | /tag
[**RegistryPostAccessrights**](RegistryApi.md#RegistryPostAccessrights) | **Post** /registry/{registryId}/accessrights | /accessrights
[**RegistryPostCredential**](RegistryApi.md#RegistryPostCredential) | **Post** /registry/{registryId}/credential | /credential
[**RegistryPostCredentialpassword**](RegistryApi.md#RegistryPostCredentialpassword) | **Post** /registry/{registryId}/credential/password | /credential/password
[**RegistryShow**](RegistryApi.md#RegistryShow) | **Get** /registry/{registryId} | Get
[**RegistryUpdate**](RegistryApi.md#RegistryUpdate) | **Patch** /registry/{registryId} | Update



## RegistryActionStart

> Registry RegistryActionStart(ctx, registryId)
/actions/start

Action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryActionStop

> Registry RegistryActionStop(ctx, registryId)
/actions/stop

Action stop

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryActionTransfer

> Registry RegistryActionTransfer(ctx, registryId, registryActionTransfer)
/actions/transfer

Action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**registryActionTransfer** | [**RegistryActionTransfer**](RegistryActionTransfer.md)|  | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryActionUpdateDomain

> Registry RegistryActionUpdateDomain(ctx, registryId, registryActionUpdateDomain)
/actions/update_domain

Action update_domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**registryActionUpdateDomain** | [**RegistryActionUpdateDomain**](RegistryActionUpdateDomain.md)|  | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryCreate

> Registry RegistryCreate(ctx, registryCreate)
Create

Create registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryCreate** | [**RegistryCreate**](RegistryCreate.md)|  | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryDelete

> RegistryDelete(ctx, registryId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

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


## RegistryDeleteAccessrightsIdentity

> Registry RegistryDeleteAccessrightsIdentity(ctx, registryId, identity)
/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**identity** | **string**| identity | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryDeleteCredentialId

> Registry RegistryDeleteCredentialId(ctx, registryId, id)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**id** | **string**| id | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryDeleteCredentialpasswordId

> Registry RegistryDeleteCredentialpasswordId(ctx, registryId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**id** | **string**| id | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryDeleteRepositoryRepositoryIdtagTagId

> RegistryDeleteRepositoryRepositoryIdtagTagId(ctx, registryId, repositoryId, tagId)
/repository/:repositoryId/tag/:tagId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**repositoryId** | **string**| repositoryId | 
**tagId** | **string**| tagId | 

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


## RegistryDeleteTagKey

> map[string]string RegistryDeleteTagKey(ctx, registryId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
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


## RegistryGetCredentialId

> CredentialPassword RegistryGetCredentialId(ctx, registryId, id)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
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


## RegistryGetCredentialpasswordId

> CredentialPassword RegistryGetCredentialpasswordId(ctx, registryId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
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


## RegistryGetRepositoryRepositoryId

> RegistryGetRepositoryRepositoryId(ctx, registryId, repositoryId)
/repository/:repositoryId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**repositoryId** | **string**| repositoryId | 

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


## RegistryGetRepositoryRepositoryIdtagTagId

> RegistryGetRepositoryRepositoryIdtagTagId(ctx, registryId, repositoryId, tagId)
/repository/:repositoryId/tag/:tagId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**repositoryId** | **string**| repositoryId | 
**tagId** | **string**| tagId | 

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


## RegistryGetServicesServiceId

> RegistryServices RegistryGetServicesServiceId(ctx, registryId, serviceId)
/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**serviceId** | **string**| serviceId | 

### Return type

[**RegistryServices**](registry.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryGetTag

> map[string]string RegistryGetTag(ctx, registryId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

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


## RegistryList

> []Registry RegistryList(ctx, optional)
List

List registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegistryListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegistryListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryListAccessrights

> []string RegistryListAccessrights(ctx, registryId)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

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


## RegistryListCredential

> []CredentialPassword RegistryListCredential(ctx, registryId)
/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

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


## RegistryListCredentialpassword

> []CredentialPassword RegistryListCredentialpassword(ctx, registryId)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

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


## RegistryListQueue

> []Event RegistryListQueue(ctx, registryId)
/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

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


## RegistryListRepository

> RegistryListRepository(ctx, registryId)
/repository

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

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


## RegistryListRepositoryRepositoryIdtag

> RegistryListRepositoryRepositoryIdtag(ctx, registryId, repositoryId)
/repository/:repositoryId/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**repositoryId** | **string**| repositoryId | 

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


## RegistryListServices

> []RegistryServices RegistryListServices(ctx, registryId)
/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

### Return type

[**[]RegistryServices**](registry.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryPatchCredentialId

> CredentialPassword RegistryPatchCredentialId(ctx, registryId, id, registryPatchCredentialId)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**id** | **string**| id | 
**registryPatchCredentialId** | [**RegistryPatchCredentialId**](RegistryPatchCredentialId.md)|  | 

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


## RegistryPatchCredentialpasswordId

> CredentialPassword RegistryPatchCredentialpasswordId(ctx, registryId, id, registryPatchCredentialpasswordId)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**id** | **string**| id | 
**registryPatchCredentialpasswordId** | [**RegistryPatchCredentialpasswordId**](RegistryPatchCredentialpasswordId.md)|  | 

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


## RegistryPatchTag

> map[string]string RegistryPatchTag(ctx, registryId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
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


## RegistryPostAccessrights

> Registry RegistryPostAccessrights(ctx, registryId, registryPostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**registryPostAccessrights** | [**RegistryPostAccessrights**](RegistryPostAccessrights.md)|  | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryPostCredential

> CredentialPassword RegistryPostCredential(ctx, registryId, registryPostCredential)
/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**registryPostCredential** | [**RegistryPostCredential**](RegistryPostCredential.md)|  | 

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


## RegistryPostCredentialpassword

> CredentialPassword RegistryPostCredentialpassword(ctx, registryId, registryPostCredentialpassword)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**registryPostCredentialpassword** | [**RegistryPostCredentialpassword**](RegistryPostCredentialpassword.md)|  | 

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


## RegistryShow

> Registry RegistryShow(ctx, registryId)
Get

Returns a single registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryUpdate

> Registry RegistryUpdate(ctx, registryId, registryUpdate)
Update

Returns modified registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string**| ID of registry | 
**registryUpdate** | [**RegistryUpdate**](RegistryUpdate.md)|  | 

### Return type

[**Registry**](registry.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

