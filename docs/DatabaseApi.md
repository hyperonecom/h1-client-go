# \DatabaseApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatabaseActionStart**](DatabaseApi.md#DatabaseActionStart) | **Post** /database/{databaseId}/actions/start | /actions/start
[**DatabaseActionStop**](DatabaseApi.md#DatabaseActionStop) | **Post** /database/{databaseId}/actions/stop | /actions/stop
[**DatabaseActionTransfer**](DatabaseApi.md#DatabaseActionTransfer) | **Post** /database/{databaseId}/actions/transfer | /actions/transfer
[**DatabaseCreate**](DatabaseApi.md#DatabaseCreate) | **Post** /database | Create
[**DatabaseDelete**](DatabaseApi.md#DatabaseDelete) | **Delete** /database/{databaseId} | Delete
[**DatabaseDeleteAccessrightsIdentity**](DatabaseApi.md#DatabaseDeleteAccessrightsIdentity) | **Delete** /database/{databaseId}/accessrights/{identity} | /accessrights/:identity
[**DatabaseDeleteCredentialId**](DatabaseApi.md#DatabaseDeleteCredentialId) | **Delete** /database/{databaseId}/credential/{id} | /credential/:id
[**DatabaseDeleteCredentialpasswordId**](DatabaseApi.md#DatabaseDeleteCredentialpasswordId) | **Delete** /database/{databaseId}/credential/password/{id} | /credential/password/:id
[**DatabaseDeleteTagKey**](DatabaseApi.md#DatabaseDeleteTagKey) | **Delete** /database/{databaseId}/tag/{key} | /tag/:key
[**DatabaseGetCredentialId**](DatabaseApi.md#DatabaseGetCredentialId) | **Get** /database/{databaseId}/credential/{id} | /credential/:id
[**DatabaseGetCredentialpasswordId**](DatabaseApi.md#DatabaseGetCredentialpasswordId) | **Get** /database/{databaseId}/credential/password/{id} | /credential/password/:id
[**DatabaseGetServicesServiceId**](DatabaseApi.md#DatabaseGetServicesServiceId) | **Get** /database/{databaseId}/services/{serviceId} | /services/:serviceId
[**DatabaseGetTag**](DatabaseApi.md#DatabaseGetTag) | **Get** /database/{databaseId}/tag | /tag
[**DatabaseList**](DatabaseApi.md#DatabaseList) | **Get** /database | List
[**DatabaseListAccessrights**](DatabaseApi.md#DatabaseListAccessrights) | **Get** /database/{databaseId}/accessrights | /accessrights
[**DatabaseListCredential**](DatabaseApi.md#DatabaseListCredential) | **Get** /database/{databaseId}/credential | /credential
[**DatabaseListCredentialpassword**](DatabaseApi.md#DatabaseListCredentialpassword) | **Get** /database/{databaseId}/credential/password | /credential/password
[**DatabaseListQueue**](DatabaseApi.md#DatabaseListQueue) | **Get** /database/{databaseId}/queue | /queue
[**DatabaseListServices**](DatabaseApi.md#DatabaseListServices) | **Get** /database/{databaseId}/services | /services
[**DatabasePatchCredentialId**](DatabaseApi.md#DatabasePatchCredentialId) | **Patch** /database/{databaseId}/credential/{id} | /credential/:id
[**DatabasePatchCredentialpasswordId**](DatabaseApi.md#DatabasePatchCredentialpasswordId) | **Patch** /database/{databaseId}/credential/password/{id} | /credential/password/:id
[**DatabasePatchTag**](DatabaseApi.md#DatabasePatchTag) | **Patch** /database/{databaseId}/tag | /tag
[**DatabasePostAccessrights**](DatabaseApi.md#DatabasePostAccessrights) | **Post** /database/{databaseId}/accessrights | /accessrights
[**DatabasePostCredential**](DatabaseApi.md#DatabasePostCredential) | **Post** /database/{databaseId}/credential | /credential
[**DatabasePostCredentialpassword**](DatabaseApi.md#DatabasePostCredentialpassword) | **Post** /database/{databaseId}/credential/password | /credential/password
[**DatabaseShow**](DatabaseApi.md#DatabaseShow) | **Get** /database/{databaseId} | Get
[**DatabaseUpdate**](DatabaseApi.md#DatabaseUpdate) | **Patch** /database/{databaseId} | Update



## DatabaseActionStart

> Database DatabaseActionStart(ctx, databaseId)
/actions/start

Action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseActionStop

> Database DatabaseActionStop(ctx, databaseId)
/actions/stop

Action stop

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseActionTransfer

> Database DatabaseActionTransfer(ctx, databaseId, databaseActionTransfer)
/actions/transfer

Action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**databaseActionTransfer** | [**DatabaseActionTransfer**](DatabaseActionTransfer.md)|  | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseCreate

> Database DatabaseCreate(ctx, databaseCreate)
Create

Create database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseCreate** | [**DatabaseCreate**](DatabaseCreate.md)|  | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseDelete

> DatabaseDelete(ctx, databaseId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

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


## DatabaseDeleteAccessrightsIdentity

> Database DatabaseDeleteAccessrightsIdentity(ctx, databaseId, identity)
/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**identity** | **string**| identity | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseDeleteCredentialId

> Database DatabaseDeleteCredentialId(ctx, databaseId, id)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**id** | **string**| id | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseDeleteCredentialpasswordId

> Database DatabaseDeleteCredentialpasswordId(ctx, databaseId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**id** | **string**| id | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseDeleteTagKey

> map[string]string DatabaseDeleteTagKey(ctx, databaseId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
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


## DatabaseGetCredentialId

> CredentialPassword DatabaseGetCredentialId(ctx, databaseId, id)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
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


## DatabaseGetCredentialpasswordId

> CredentialPassword DatabaseGetCredentialpasswordId(ctx, databaseId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
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


## DatabaseGetServicesServiceId

> DatabaseServices DatabaseGetServicesServiceId(ctx, databaseId, serviceId)
/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**serviceId** | **string**| serviceId | 

### Return type

[**DatabaseServices**](database.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseGetTag

> map[string]string DatabaseGetTag(ctx, databaseId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

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


## DatabaseList

> []Database DatabaseList(ctx, optional)
List

List database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DatabaseListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatabaseListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseListAccessrights

> []string DatabaseListAccessrights(ctx, databaseId)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

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


## DatabaseListCredential

> []CredentialPassword DatabaseListCredential(ctx, databaseId)
/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

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


## DatabaseListCredentialpassword

> []CredentialPassword DatabaseListCredentialpassword(ctx, databaseId)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

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


## DatabaseListQueue

> []Event DatabaseListQueue(ctx, databaseId)
/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

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


## DatabaseListServices

> []DatabaseServices DatabaseListServices(ctx, databaseId)
/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

### Return type

[**[]DatabaseServices**](database.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabasePatchCredentialId

> CredentialPassword DatabasePatchCredentialId(ctx, databaseId, id, databasePatchCredentialId)
/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**id** | **string**| id | 
**databasePatchCredentialId** | [**DatabasePatchCredentialId**](DatabasePatchCredentialId.md)|  | 

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


## DatabasePatchCredentialpasswordId

> CredentialPassword DatabasePatchCredentialpasswordId(ctx, databaseId, id, databasePatchCredentialpasswordId)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**id** | **string**| id | 
**databasePatchCredentialpasswordId** | [**DatabasePatchCredentialpasswordId**](DatabasePatchCredentialpasswordId.md)|  | 

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


## DatabasePatchTag

> map[string]string DatabasePatchTag(ctx, databaseId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
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


## DatabasePostAccessrights

> Database DatabasePostAccessrights(ctx, databaseId, databasePostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**databasePostAccessrights** | [**DatabasePostAccessrights**](DatabasePostAccessrights.md)|  | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabasePostCredential

> CredentialPassword DatabasePostCredential(ctx, databaseId, databasePostCredential)
/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**databasePostCredential** | [**DatabasePostCredential**](DatabasePostCredential.md)|  | 

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


## DatabasePostCredentialpassword

> CredentialPassword DatabasePostCredentialpassword(ctx, databaseId, databasePostCredentialpassword)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**databasePostCredentialpassword** | [**DatabasePostCredentialpassword**](DatabasePostCredentialpassword.md)|  | 

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


## DatabaseShow

> Database DatabaseShow(ctx, databaseId)
Get

Returns a single database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseUpdate

> Database DatabaseUpdate(ctx, databaseId, databaseUpdate)
Update

Returns modified database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string**| ID of database | 
**databaseUpdate** | [**DatabaseUpdate**](DatabaseUpdate.md)|  | 

### Return type

[**Database**](database.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

