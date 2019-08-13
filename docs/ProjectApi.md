# \ProjectApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectCreate**](ProjectApi.md#ProjectCreate) | **Post** /project | Create
[**ProjectDelete**](ProjectApi.md#ProjectDelete) | **Delete** /project/{projectId} | Delete
[**ProjectDeleteAccessrightsId**](ProjectApi.md#ProjectDeleteAccessrightsId) | **Delete** /project/{projectId}/accessrights/{id} | /accessrights/:id
[**ProjectDeleteCredentialStoreId**](ProjectApi.md#ProjectDeleteCredentialStoreId) | **Delete** /project/{projectId}/credentialStore/{id} | /credentialStore/:id
[**ProjectDeleteCredentialStorecertificateId**](ProjectApi.md#ProjectDeleteCredentialStorecertificateId) | **Delete** /project/{projectId}/credentialStore/certificate/{id} | /credentialStore/certificate/:id
[**ProjectDeleteTagKey**](ProjectApi.md#ProjectDeleteTagKey) | **Delete** /project/{projectId}/tag/{key} | /tag/:key
[**ProjectGetCredentialStoreId**](ProjectApi.md#ProjectGetCredentialStoreId) | **Get** /project/{projectId}/credentialStore/{id} | /credentialStore/:id
[**ProjectGetCredentialStorecertificateId**](ProjectApi.md#ProjectGetCredentialStorecertificateId) | **Get** /project/{projectId}/credentialStore/certificate/{id} | /credentialStore/certificate/:id
[**ProjectGetServicesServiceId**](ProjectApi.md#ProjectGetServicesServiceId) | **Get** /project/{projectId}/services/{serviceId} | /services/:serviceId
[**ProjectGetTag**](ProjectApi.md#ProjectGetTag) | **Get** /project/{projectId}/tag | /tag
[**ProjectList**](ProjectApi.md#ProjectList) | **Get** /project | List
[**ProjectListAccessrights**](ProjectApi.md#ProjectListAccessrights) | **Get** /project/{projectId}/accessrights | /accessrights
[**ProjectListCredentialStore**](ProjectApi.md#ProjectListCredentialStore) | **Get** /project/{projectId}/credentialStore | /credentialStore
[**ProjectListCredentialStorecertificate**](ProjectApi.md#ProjectListCredentialStorecertificate) | **Get** /project/{projectId}/credentialStore/certificate | /credentialStore/certificate
[**ProjectListLimit**](ProjectApi.md#ProjectListLimit) | **Get** /project/{projectId}/limit | /limit
[**ProjectListQueue**](ProjectApi.md#ProjectListQueue) | **Get** /project/{projectId}/queue | /queue
[**ProjectListServices**](ProjectApi.md#ProjectListServices) | **Get** /project/{projectId}/services | /services
[**ProjectPatchCredentialStoreId**](ProjectApi.md#ProjectPatchCredentialStoreId) | **Patch** /project/{projectId}/credentialStore/{id} | /credentialStore/:id
[**ProjectPatchCredentialStorecertificateId**](ProjectApi.md#ProjectPatchCredentialStorecertificateId) | **Patch** /project/{projectId}/credentialStore/certificate/{id} | /credentialStore/certificate/:id
[**ProjectPatchTag**](ProjectApi.md#ProjectPatchTag) | **Patch** /project/{projectId}/tag | /tag
[**ProjectPostAccessrights**](ProjectApi.md#ProjectPostAccessrights) | **Post** /project/{projectId}/accessrights | /accessrights
[**ProjectPostCredentialStore**](ProjectApi.md#ProjectPostCredentialStore) | **Post** /project/{projectId}/credentialStore | /credentialStore
[**ProjectPostCredentialStorecertificate**](ProjectApi.md#ProjectPostCredentialStorecertificate) | **Post** /project/{projectId}/credentialStore/certificate | /credentialStore/certificate
[**ProjectShow**](ProjectApi.md#ProjectShow) | **Get** /project/{projectId} | Get
[**ProjectUpdate**](ProjectApi.md#ProjectUpdate) | **Patch** /project/{projectId} | Update



## ProjectCreate

> Project ProjectCreate(ctx, projectCreate)
Create

Create project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectCreate** | [**ProjectCreate**](ProjectCreate.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectDelete

> ProjectDelete(ctx, projectId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

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


## ProjectDeleteAccessrightsId

> ProjectDeleteAccessrightsId(ctx, projectId, id)
/accessrights/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**id** | **string**| id | 

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


## ProjectDeleteCredentialStoreId

> Project ProjectDeleteCredentialStoreId(ctx, projectId, id)
/credentialStore/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**id** | **string**| id | 

### Return type

[**Project**](project.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectDeleteCredentialStorecertificateId

> CredentialCertificate ProjectDeleteCredentialStorecertificateId(ctx, projectId, id)
/credentialStore/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
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


## ProjectDeleteTagKey

> map[string]string ProjectDeleteTagKey(ctx, projectId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
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


## ProjectGetCredentialStoreId

> CredentialPassword ProjectGetCredentialStoreId(ctx, projectId, id)
/credentialStore/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
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


## ProjectGetCredentialStorecertificateId

> CredentialCertificate ProjectGetCredentialStorecertificateId(ctx, projectId, id)
/credentialStore/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
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


## ProjectGetServicesServiceId

> ProjectServices ProjectGetServicesServiceId(ctx, projectId, serviceId)
/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**serviceId** | **string**| serviceId | 

### Return type

[**ProjectServices**](project.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectGetTag

> map[string]string ProjectGetTag(ctx, projectId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

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


## ProjectList

> []Project ProjectList(ctx, optional)
List

List project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **limit** | **optional.Float32**| Filter by $limit | 
 **active** | **optional.Bool**| Filter by active | 
 **organisation** | **optional.String**| Filter by organisation | 
 **accessRightsId** | **optional.String**| Filter by accessRights.id | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Project**](project.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectListAccessrights

> []AccessrightsUserRole ProjectListAccessrights(ctx, projectId)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

### Return type

[**[]AccessrightsUserRole**](accessrightsUserRole.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectListCredentialStore

> []CredentialPassword ProjectListCredentialStore(ctx, projectId)
/credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

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


## ProjectListCredentialStorecertificate

> []CredentialCertificate ProjectListCredentialStorecertificate(ctx, projectId)
/credentialStore/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

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


## ProjectListLimit

> []ProjectLimit ProjectListLimit(ctx, projectId)
/limit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

### Return type

[**[]ProjectLimit**](project.limit.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectListQueue

> []Event ProjectListQueue(ctx, projectId)
/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

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


## ProjectListServices

> []ProjectServices ProjectListServices(ctx, projectId)
/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

### Return type

[**[]ProjectServices**](project.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectPatchCredentialStoreId

> CredentialPassword ProjectPatchCredentialStoreId(ctx, projectId, id, projectPatchCredentialStoreId)
/credentialStore/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**id** | **string**| id | 
**projectPatchCredentialStoreId** | [**ProjectPatchCredentialStoreId**](ProjectPatchCredentialStoreId.md)|  | 

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


## ProjectPatchCredentialStorecertificateId

> CredentialCertificate ProjectPatchCredentialStorecertificateId(ctx, projectId, id, projectPatchCredentialStorecertificateId)
/credentialStore/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**id** | **string**| id | 
**projectPatchCredentialStorecertificateId** | [**ProjectPatchCredentialStorecertificateId**](ProjectPatchCredentialStorecertificateId.md)|  | 

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


## ProjectPatchTag

> map[string]string ProjectPatchTag(ctx, projectId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
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


## ProjectPostAccessrights

> ProjectPostAccessrights(ctx, projectId, projectPostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**projectPostAccessrights** | [**ProjectPostAccessrights**](ProjectPostAccessrights.md)|  | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectPostCredentialStore

> CredentialPassword ProjectPostCredentialStore(ctx, projectId, projectPostCredentialStore)
/credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**projectPostCredentialStore** | [**ProjectPostCredentialStore**](ProjectPostCredentialStore.md)|  | 

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


## ProjectPostCredentialStorecertificate

> CredentialCertificate ProjectPostCredentialStorecertificate(ctx, projectId, projectPostCredentialStorecertificate)
/credentialStore/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**projectPostCredentialStorecertificate** | [**ProjectPostCredentialStorecertificate**](ProjectPostCredentialStorecertificate.md)|  | 

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


## ProjectShow

> Project ProjectShow(ctx, projectId)
Get

Returns a single project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 

### Return type

[**Project**](project.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectUpdate

> Project ProjectUpdate(ctx, projectId, projectUpdate)
Update

Returns modified project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| ID of project | 
**projectUpdate** | [**ProjectUpdate**](ProjectUpdate.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

