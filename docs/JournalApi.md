# \JournalApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JournalActionTransfer**](JournalApi.md#JournalActionTransfer) | **Post** /journal/{journalId}/actions/transfer | /actions/transfer
[**JournalActionUpdateRetention**](JournalApi.md#JournalActionUpdateRetention) | **Post** /journal/{journalId}/actions/update_retention | /actions/update_retention
[**JournalCreate**](JournalApi.md#JournalCreate) | **Post** /journal | Create
[**JournalDelete**](JournalApi.md#JournalDelete) | **Delete** /journal/{journalId} | Delete
[**JournalDeleteAccessrightsIdentity**](JournalApi.md#JournalDeleteAccessrightsIdentity) | **Delete** /journal/{journalId}/accessrights/{identity} | /accessrights/:identity
[**JournalDeleteCredentialId**](JournalApi.md#JournalDeleteCredentialId) | **Delete** /journal/{journalId}/credential/{id} | /credential/:id
[**JournalDeleteCredentialpasswordId**](JournalApi.md#JournalDeleteCredentialpasswordId) | **Delete** /journal/{journalId}/credential/password/{id} | /credential/password/:id
[**JournalDeleteTagKey**](JournalApi.md#JournalDeleteTagKey) | **Delete** /journal/{journalId}/tag/{key} | /tag/:key
[**JournalGetCredentialId**](JournalApi.md#JournalGetCredentialId) | **Get** /journal/{journalId}/credential/{id} | /credential/:id
[**JournalGetCredentialpasswordId**](JournalApi.md#JournalGetCredentialpasswordId) | **Get** /journal/{journalId}/credential/password/{id} | /credential/password/:id
[**JournalGetLog**](JournalApi.md#JournalGetLog) | **Get** /journal/{journalId}/log | /log
[**JournalGetServicesServiceId**](JournalApi.md#JournalGetServicesServiceId) | **Get** /journal/{journalId}/services/{serviceId} | /services/:serviceId
[**JournalGetTag**](JournalApi.md#JournalGetTag) | **Get** /journal/{journalId}/tag | /tag
[**JournalList**](JournalApi.md#JournalList) | **Get** /journal | List
[**JournalListAccessrights**](JournalApi.md#JournalListAccessrights) | **Get** /journal/{journalId}/accessrights | /accessrights
[**JournalListCredential**](JournalApi.md#JournalListCredential) | **Get** /journal/{journalId}/credential | /credential
[**JournalListCredentialpassword**](JournalApi.md#JournalListCredentialpassword) | **Get** /journal/{journalId}/credential/password | /credential/password
[**JournalListQueue**](JournalApi.md#JournalListQueue) | **Get** /journal/{journalId}/queue | /queue
[**JournalListServices**](JournalApi.md#JournalListServices) | **Get** /journal/{journalId}/services | /services
[**JournalPatchCredentialId**](JournalApi.md#JournalPatchCredentialId) | **Patch** /journal/{journalId}/credential/{id} | /credential/:id
[**JournalPatchCredentialpasswordId**](JournalApi.md#JournalPatchCredentialpasswordId) | **Patch** /journal/{journalId}/credential/password/{id} | /credential/password/:id
[**JournalPatchTag**](JournalApi.md#JournalPatchTag) | **Patch** /journal/{journalId}/tag | /tag
[**JournalPostAccessrights**](JournalApi.md#JournalPostAccessrights) | **Post** /journal/{journalId}/accessrights | /accessrights
[**JournalPostCredential**](JournalApi.md#JournalPostCredential) | **Post** /journal/{journalId}/credential | /credential
[**JournalPostCredentialpassword**](JournalApi.md#JournalPostCredentialpassword) | **Post** /journal/{journalId}/credential/password | /credential/password
[**JournalPutTag**](JournalApi.md#JournalPutTag) | **Put** /journal/{journalId}/tag | /tag
[**JournalShow**](JournalApi.md#JournalShow) | **Get** /journal/{journalId} | Get
[**JournalUpdate**](JournalApi.md#JournalUpdate) | **Patch** /journal/{journalId} | Update



## JournalActionTransfer

> Journal JournalActionTransfer(ctx, journalId, journalActionTransfer)

/actions/transfer

Action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**journalActionTransfer** | [**JournalActionTransfer**](JournalActionTransfer.md)|  | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalActionUpdateRetention

> Journal JournalActionUpdateRetention(ctx, journalId, journalActionUpdateRetention)

/actions/update_retention

Action update_retention

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**journalActionUpdateRetention** | [**JournalActionUpdateRetention**](JournalActionUpdateRetention.md)|  | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalCreate

> Journal JournalCreate(ctx, journalCreate)

Create

Create journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalCreate** | [**JournalCreate**](JournalCreate.md)|  | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalDelete

> JournalDelete(ctx, journalId)

Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 

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


## JournalDeleteAccessrightsIdentity

> Journal JournalDeleteAccessrightsIdentity(ctx, journalId, identity)

/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**identity** | **string**| identity | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalDeleteCredentialId

> Journal JournalDeleteCredentialId(ctx, journalId, id)

/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**id** | **string**| id | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalDeleteCredentialpasswordId

> Journal JournalDeleteCredentialpasswordId(ctx, journalId, id)

/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**id** | **string**| id | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalDeleteTagKey

> map[string]string JournalDeleteTagKey(ctx, journalId, key)

/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
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


## JournalGetCredentialId

> CredentialPassword JournalGetCredentialId(ctx, journalId, id)

/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
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


## JournalGetCredentialpasswordId

> CredentialPassword JournalGetCredentialpasswordId(ctx, journalId, id)

/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
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


## JournalGetLog

> JournalGetLog(ctx, journalId, optional)

/log

websocket is also supported

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
 **optional** | ***JournalGetLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JournalGetLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**| since | 
 **until** | **optional.Time**| until | 
 **follow** | **optional.Bool**| follow | 
 **tail** | **optional.Float32**| tail | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| tag | 

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


## JournalGetServicesServiceId

> JournalServices JournalGetServicesServiceId(ctx, journalId, serviceId)

/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**serviceId** | **string**| serviceId | 

### Return type

[**JournalServices**](journal.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalGetTag

> map[string]string JournalGetTag(ctx, journalId)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 

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


## JournalList

> []Journal JournalList(ctx, optional)

List

List journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JournalListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JournalListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalListAccessrights

> []string JournalListAccessrights(ctx, journalId)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 

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


## JournalListCredential

> []CredentialPassword JournalListCredential(ctx, journalId)

/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 

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


## JournalListCredentialpassword

> []CredentialPassword JournalListCredentialpassword(ctx, journalId)

/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 

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


## JournalListQueue

> []Event JournalListQueue(ctx, journalId, optional)

/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
 **optional** | ***JournalListQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JournalListQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Float32**| $limit | 
 **skip** | **optional.Float32**| $skip | 

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


## JournalListServices

> []JournalServices JournalListServices(ctx, journalId)

/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 

### Return type

[**[]JournalServices**](journal.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalPatchCredentialId

> CredentialPassword JournalPatchCredentialId(ctx, journalId, id, journalPatchCredentialId)

/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**id** | **string**| id | 
**journalPatchCredentialId** | [**JournalPatchCredentialId**](JournalPatchCredentialId.md)|  | 

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


## JournalPatchCredentialpasswordId

> CredentialPassword JournalPatchCredentialpasswordId(ctx, journalId, id, journalPatchCredentialpasswordId)

/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**id** | **string**| id | 
**journalPatchCredentialpasswordId** | [**JournalPatchCredentialpasswordId**](JournalPatchCredentialpasswordId.md)|  | 

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


## JournalPatchTag

> map[string]string JournalPatchTag(ctx, journalId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
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


## JournalPostAccessrights

> Journal JournalPostAccessrights(ctx, journalId, journalPostAccessrights)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**journalPostAccessrights** | [**JournalPostAccessrights**](JournalPostAccessrights.md)|  | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalPostCredential

> CredentialPassword JournalPostCredential(ctx, journalId, journalPostCredential)

/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**journalPostCredential** | [**JournalPostCredential**](JournalPostCredential.md)|  | 

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


## JournalPostCredentialpassword

> CredentialPassword JournalPostCredentialpassword(ctx, journalId, journalPostCredentialpassword)

/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**journalPostCredentialpassword** | [**JournalPostCredentialpassword**](JournalPostCredentialpassword.md)|  | 

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


## JournalPutTag

> map[string]string JournalPutTag(ctx, journalId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
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


## JournalShow

> Journal JournalShow(ctx, journalId)

Get

Returns a single journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JournalUpdate

> Journal JournalUpdate(ctx, journalId, journalUpdate)

Update

Returns modified journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string**| ID of journal | 
**journalUpdate** | [**JournalUpdate**](JournalUpdate.md)|  | 

### Return type

[**Journal**](journal.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

