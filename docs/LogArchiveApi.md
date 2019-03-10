# \LogArchiveApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogArchiveActionTransfer**](LogArchiveApi.md#LogArchiveActionTransfer) | **Post** /logArchive/{logArchiveId}/actions/transfer | /actions/transfer
[**LogArchiveCreate**](LogArchiveApi.md#LogArchiveCreate) | **Post** /logArchive | Create
[**LogArchiveDelete**](LogArchiveApi.md#LogArchiveDelete) | **Delete** /logArchive/{logArchiveId} | Delete
[**LogArchiveDeleteAccessrightsIdentity**](LogArchiveApi.md#LogArchiveDeleteAccessrightsIdentity) | **Delete** /logArchive/{logArchiveId}/accessrights/{identity} | /accessrights/:identity
[**LogArchiveDeleteCredentialpasswordId**](LogArchiveApi.md#LogArchiveDeleteCredentialpasswordId) | **Delete** /logArchive/{logArchiveId}/credential/password/{id} | /credential/password/:id
[**LogArchiveDeleteTagKey**](LogArchiveApi.md#LogArchiveDeleteTagKey) | **Delete** /logArchive/{logArchiveId}/tag/{key} | /tag/:key
[**LogArchiveGetCredentialpasswordId**](LogArchiveApi.md#LogArchiveGetCredentialpasswordId) | **Get** /logArchive/{logArchiveId}/credential/password/{id} | /credential/password/:id
[**LogArchiveGetServicesServiceId**](LogArchiveApi.md#LogArchiveGetServicesServiceId) | **Get** /logArchive/{logArchiveId}/services/{serviceId} | /services/:serviceId
[**LogArchiveGetTag**](LogArchiveApi.md#LogArchiveGetTag) | **Get** /logArchive/{logArchiveId}/tag | /tag
[**LogArchiveList**](LogArchiveApi.md#LogArchiveList) | **Get** /logArchive | List
[**LogArchiveListAccessrights**](LogArchiveApi.md#LogArchiveListAccessrights) | **Get** /logArchive/{logArchiveId}/accessrights | /accessrights
[**LogArchiveListCredentialpassword**](LogArchiveApi.md#LogArchiveListCredentialpassword) | **Get** /logArchive/{logArchiveId}/credential/password | /credential/password
[**LogArchiveListQueue**](LogArchiveApi.md#LogArchiveListQueue) | **Get** /logArchive/{logArchiveId}/queue | /queue
[**LogArchiveListServices**](LogArchiveApi.md#LogArchiveListServices) | **Get** /logArchive/{logArchiveId}/services | /services
[**LogArchivePatchCredentialpasswordId**](LogArchiveApi.md#LogArchivePatchCredentialpasswordId) | **Patch** /logArchive/{logArchiveId}/credential/password/{id} | /credential/password/:id
[**LogArchivePatchTag**](LogArchiveApi.md#LogArchivePatchTag) | **Patch** /logArchive/{logArchiveId}/tag | /tag
[**LogArchivePostAccessrights**](LogArchiveApi.md#LogArchivePostAccessrights) | **Post** /logArchive/{logArchiveId}/accessrights | /accessrights
[**LogArchivePostCredentialpassword**](LogArchiveApi.md#LogArchivePostCredentialpassword) | **Post** /logArchive/{logArchiveId}/credential/password | /credential/password
[**LogArchiveShow**](LogArchiveApi.md#LogArchiveShow) | **Get** /logArchive/{logArchiveId} | Get
[**LogArchiveUpdate**](LogArchiveApi.md#LogArchiveUpdate) | **Patch** /logArchive/{logArchiveId} | Update


# **LogArchiveActionTransfer**
> LogArchive LogArchiveActionTransfer(ctx, logArchiveId, logArchiveActionTransfer)
/actions/transfer

Action transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **logArchiveActionTransfer** | [**LogArchiveActionTransfer**](LogArchiveActionTransfer.md)|  | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveCreate**
> LogArchive LogArchiveCreate(ctx, logArchiveCreate)
Create

Create logArchive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveCreate** | [**LogArchiveCreate**](LogArchiveCreate.md)|  | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveDelete**
> LogArchiveDelete(ctx, logArchiveId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveDeleteAccessrightsIdentity**
> LogArchive LogArchiveDeleteAccessrightsIdentity(ctx, logArchiveId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **identity** | **string**| identity | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveDeleteCredentialpasswordId**
> LogArchive LogArchiveDeleteCredentialpasswordId(ctx, logArchiveId, id)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **id** | **string**| id | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveDeleteTagKey**
> map[string]interface{} LogArchiveDeleteTagKey(ctx, logArchiveId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveGetCredentialpasswordId**
> CredentialPassword LogArchiveGetCredentialpasswordId(ctx, logArchiveId, id)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **id** | **string**| id | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveGetServicesServiceId**
> LogArchiveServices LogArchiveGetServicesServiceId(ctx, logArchiveId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **serviceId** | **string**| serviceId | 

### Return type

[**LogArchiveServices**](logArchive.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveGetTag**
> map[string]interface{} LogArchiveGetTag(ctx, logArchiveId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveList**
> []LogArchive LogArchiveList(ctx, optional)
List

List logArchive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogArchiveListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogArchiveListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveListAccessrights**
> []string LogArchiveListAccessrights(ctx, logArchiveId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveListCredentialpassword**
> []CredentialPassword LogArchiveListCredentialpassword(ctx, logArchiveId)
/credential/password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

[**[]CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveListQueue**
> []Event LogArchiveListQueue(ctx, logArchiveId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveListServices**
> []LogArchiveServices LogArchiveListServices(ctx, logArchiveId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

[**[]LogArchiveServices**](logArchive.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchivePatchCredentialpasswordId**
> CredentialPassword LogArchivePatchCredentialpasswordId(ctx, logArchiveId, id, logArchivePatchCredentialpasswordId)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **id** | **string**| id | 
  **logArchivePatchCredentialpasswordId** | [**LogArchivePatchCredentialpasswordId**](LogArchivePatchCredentialpasswordId.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchivePatchTag**
> map[string]interface{} LogArchivePatchTag(ctx, logArchiveId, requestBody)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchivePostAccessrights**
> LogArchive LogArchivePostAccessrights(ctx, logArchiveId, logArchivePostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **logArchivePostAccessrights** | [**LogArchivePostAccessrights**](LogArchivePostAccessrights.md)|  | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchivePostCredentialpassword**
> CredentialPassword LogArchivePostCredentialpassword(ctx, logArchiveId, logArchivePostCredentialpassword)
/credential/password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **logArchivePostCredentialpassword** | [**LogArchivePostCredentialpassword**](LogArchivePostCredentialpassword.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveShow**
> LogArchive LogArchiveShow(ctx, logArchiveId)
Get

Returns a single logArchive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogArchiveUpdate**
> LogArchive LogArchiveUpdate(ctx, logArchiveId, logArchiveUpdate)
Update

Returns modified logArchive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **logArchiveUpdate** | [**LogArchiveUpdate**](LogArchiveUpdate.md)|  | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

