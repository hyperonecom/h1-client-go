# \LogArchiveApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionLogArchiveTransfer**](LogArchiveApi.md#ActionLogArchiveTransfer) | **Post** /logArchive/{logArchiveId}/actions/transfer | /actions/transfer
[**CreateLogArchive**](LogArchiveApi.md#CreateLogArchive) | **Post** /logArchive | Create
[**DeleteLogArchive**](LogArchiveApi.md#DeleteLogArchive) | **Delete** /logArchive/{logArchiveId} | Delete
[**ListLogArchive**](LogArchiveApi.md#ListLogArchive) | **Get** /logArchive | List
[**OperationLogArchiveDeleteaccessrightsIdentity**](LogArchiveApi.md#OperationLogArchiveDeleteaccessrightsIdentity) | **Delete** /logArchive/{logArchiveId}/accessrights/{identity} | /accessrights/:identity
[**OperationLogArchiveDeletecredentialcertificateId**](LogArchiveApi.md#OperationLogArchiveDeletecredentialcertificateId) | **Delete** /logArchive/{logArchiveId}/credential/certificate/{id} | /credential/certificate/:id
[**OperationLogArchiveDeletecredentialpasswordId**](LogArchiveApi.md#OperationLogArchiveDeletecredentialpasswordId) | **Delete** /logArchive/{logArchiveId}/credential/password/{id} | /credential/password/:id
[**OperationLogArchiveDeletetagKey**](LogArchiveApi.md#OperationLogArchiveDeletetagKey) | **Delete** /logArchive/{logArchiveId}/tag/{key} | /tag/:key
[**OperationLogArchiveGetcredentialcertificateId**](LogArchiveApi.md#OperationLogArchiveGetcredentialcertificateId) | **Get** /logArchive/{logArchiveId}/credential/certificate/{id} | /credential/certificate/:id
[**OperationLogArchiveGetcredentialpasswordId**](LogArchiveApi.md#OperationLogArchiveGetcredentialpasswordId) | **Get** /logArchive/{logArchiveId}/credential/password/{id} | /credential/password/:id
[**OperationLogArchiveGetservicesServiceId**](LogArchiveApi.md#OperationLogArchiveGetservicesServiceId) | **Get** /logArchive/{logArchiveId}/services/{serviceId} | /services/:serviceId
[**OperationLogArchiveGettag**](LogArchiveApi.md#OperationLogArchiveGettag) | **Get** /logArchive/{logArchiveId}/tag/ | /tag/
[**OperationLogArchiveListaccessrights**](LogArchiveApi.md#OperationLogArchiveListaccessrights) | **Get** /logArchive/{logArchiveId}/accessrights/ | /accessrights/
[**OperationLogArchiveListcredentialcertificate**](LogArchiveApi.md#OperationLogArchiveListcredentialcertificate) | **Get** /logArchive/{logArchiveId}/credential/certificate | /credential/certificate
[**OperationLogArchiveListcredentialpassword**](LogArchiveApi.md#OperationLogArchiveListcredentialpassword) | **Get** /logArchive/{logArchiveId}/credential/password | /credential/password
[**OperationLogArchiveListqueue**](LogArchiveApi.md#OperationLogArchiveListqueue) | **Get** /logArchive/{logArchiveId}/queue/ | /queue/
[**OperationLogArchiveListservices**](LogArchiveApi.md#OperationLogArchiveListservices) | **Get** /logArchive/{logArchiveId}/services/ | /services/
[**OperationLogArchivePatchcredentialcertificateId**](LogArchiveApi.md#OperationLogArchivePatchcredentialcertificateId) | **Patch** /logArchive/{logArchiveId}/credential/certificate/{id} | /credential/certificate/:id
[**OperationLogArchivePatchcredentialpasswordId**](LogArchiveApi.md#OperationLogArchivePatchcredentialpasswordId) | **Patch** /logArchive/{logArchiveId}/credential/password/{id} | /credential/password/:id
[**OperationLogArchivePatchtag**](LogArchiveApi.md#OperationLogArchivePatchtag) | **Patch** /logArchive/{logArchiveId}/tag/ | /tag/
[**OperationLogArchivePostaccessrights**](LogArchiveApi.md#OperationLogArchivePostaccessrights) | **Post** /logArchive/{logArchiveId}/accessrights/ | /accessrights/
[**OperationLogArchivePostcredentialcertificate**](LogArchiveApi.md#OperationLogArchivePostcredentialcertificate) | **Post** /logArchive/{logArchiveId}/credential/certificate | /credential/certificate
[**OperationLogArchivePostcredentialpassword**](LogArchiveApi.md#OperationLogArchivePostcredentialpassword) | **Post** /logArchive/{logArchiveId}/credential/password | /credential/password
[**ShowLogArchive**](LogArchiveApi.md#ShowLogArchive) | **Get** /logArchive/{logArchiveId} | Get
[**UpdateLogArchive**](LogArchiveApi.md#UpdateLogArchive) | **Patch** /logArchive/{logArchiveId} | Update


# **ActionLogArchiveTransfer**
> LogArchive ActionLogArchiveTransfer(ctx, logArchiveId, optional)
/actions/transfer

Action transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
 **optional** | ***ActionLogArchiveTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionLogArchiveTransferOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject47** | [**optional.Interface of InlineObject47**](InlineObject47.md)|  | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLogArchive**
> LogArchive CreateLogArchive(ctx, logArchiveCreate)
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

# **DeleteLogArchive**
> DeleteLogArchive(ctx, logArchiveId)
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

# **ListLogArchive**
> []LogArchive ListLogArchive(ctx, optional)
List

List logArchive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListLogArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListLogArchiveOpts struct

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

# **OperationLogArchiveDeleteaccessrightsIdentity**
> LogArchive OperationLogArchiveDeleteaccessrightsIdentity(ctx, logArchiveId, identity)
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

# **OperationLogArchiveDeletecredentialcertificateId**
> LogArchive OperationLogArchiveDeletecredentialcertificateId(ctx, logArchiveId, id)
/credential/certificate/:id

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

# **OperationLogArchiveDeletecredentialpasswordId**
> LogArchive OperationLogArchiveDeletecredentialpasswordId(ctx, logArchiveId, id)
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

# **OperationLogArchiveDeletetagKey**
> map[string]string OperationLogArchiveDeletetagKey(ctx, logArchiveId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchiveGetcredentialcertificateId**
> CredentialCertificate OperationLogArchiveGetcredentialcertificateId(ctx, logArchiveId, id)
/credential/certificate/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **id** | **string**| id | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchiveGetcredentialpasswordId**
> CredentialPassword OperationLogArchiveGetcredentialpasswordId(ctx, logArchiveId, id)
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

# **OperationLogArchiveGetservicesServiceId**
> LogArchiveServices OperationLogArchiveGetservicesServiceId(ctx, logArchiveId, serviceId)
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

# **OperationLogArchiveGettag**
> map[string]string OperationLogArchiveGettag(ctx, logArchiveId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchiveListaccessrights**
> []string OperationLogArchiveListaccessrights(ctx, logArchiveId)
/accessrights/

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

# **OperationLogArchiveListcredentialcertificate**
> []CredentialCertificate OperationLogArchiveListcredentialcertificate(ctx, logArchiveId)
/credential/certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 

### Return type

[**[]CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchiveListcredentialpassword**
> []CredentialPassword OperationLogArchiveListcredentialpassword(ctx, logArchiveId)
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

# **OperationLogArchiveListqueue**
> []Event OperationLogArchiveListqueue(ctx, logArchiveId)
/queue/

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

# **OperationLogArchiveListservices**
> []LogArchiveServices OperationLogArchiveListservices(ctx, logArchiveId)
/services/

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

# **OperationLogArchivePatchcredentialcertificateId**
> CredentialCertificate OperationLogArchivePatchcredentialcertificateId(ctx, logArchiveId, id, optional)
/credential/certificate/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **id** | **string**| id | 
 **optional** | ***OperationLogArchivePatchcredentialcertificateIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationLogArchivePatchcredentialcertificateIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject51** | [**optional.Interface of InlineObject51**](InlineObject51.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchivePatchcredentialpasswordId**
> CredentialPassword OperationLogArchivePatchcredentialpasswordId(ctx, logArchiveId, id, optional)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **id** | **string**| id | 
 **optional** | ***OperationLogArchivePatchcredentialpasswordIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationLogArchivePatchcredentialpasswordIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject49** | [**optional.Interface of InlineObject49**](InlineObject49.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchivePatchtag**
> map[string]string OperationLogArchivePatchtag(ctx, logArchiveId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchivePostaccessrights**
> string OperationLogArchivePostaccessrights(ctx, logArchiveId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
 **optional** | ***OperationLogArchivePostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationLogArchivePostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject52** | [**optional.Interface of InlineObject52**](InlineObject52.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchivePostcredentialcertificate**
> CredentialCertificate OperationLogArchivePostcredentialcertificate(ctx, logArchiveId, optional)
/credential/certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
 **optional** | ***OperationLogArchivePostcredentialcertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationLogArchivePostcredentialcertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject50** | [**optional.Interface of InlineObject50**](InlineObject50.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationLogArchivePostcredentialpassword**
> CredentialPassword OperationLogArchivePostcredentialpassword(ctx, logArchiveId, optional)
/credential/password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
 **optional** | ***OperationLogArchivePostcredentialpasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationLogArchivePostcredentialpasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject48** | [**optional.Interface of InlineObject48**](InlineObject48.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowLogArchive**
> LogArchive ShowLogArchive(ctx, logArchiveId)
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

# **UpdateLogArchive**
> LogArchive UpdateLogArchive(ctx, logArchiveId, optional)
Update

Returns modified logArchive

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logArchiveId** | **string**| ID of logArchive | 
 **optional** | ***UpdateLogArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateLogArchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject46** | [**optional.Interface of InlineObject46**](InlineObject46.md)|  | 

### Return type

[**LogArchive**](logArchive.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

