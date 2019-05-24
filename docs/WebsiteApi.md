# \WebsiteApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebsiteActionStart**](WebsiteApi.md#WebsiteActionStart) | **Post** /website/{websiteId}/actions/start | /actions/start
[**WebsiteActionStop**](WebsiteApi.md#WebsiteActionStop) | **Post** /website/{websiteId}/actions/stop | /actions/stop
[**WebsiteCreate**](WebsiteApi.md#WebsiteCreate) | **Post** /website | Create
[**WebsiteDelete**](WebsiteApi.md#WebsiteDelete) | **Delete** /website/{websiteId} | Delete
[**WebsiteDeleteAccessrightsIdentity**](WebsiteApi.md#WebsiteDeleteAccessrightsIdentity) | **Delete** /website/{websiteId}/accessrights/{identity} | /accessrights/:identity
[**WebsiteDeleteCredentialcertificateId**](WebsiteApi.md#WebsiteDeleteCredentialcertificateId) | **Delete** /website/{websiteId}/credential/certificate/{id} | /credential/certificate/:id
[**WebsiteDeleteCredentialpasswordId**](WebsiteApi.md#WebsiteDeleteCredentialpasswordId) | **Delete** /website/{websiteId}/credential/password/{id} | /credential/password/:id
[**WebsiteDeleteTagKey**](WebsiteApi.md#WebsiteDeleteTagKey) | **Delete** /website/{websiteId}/tag/{key} | /tag/:key
[**WebsiteGetCredentialcertificateId**](WebsiteApi.md#WebsiteGetCredentialcertificateId) | **Get** /website/{websiteId}/credential/certificate/{id} | /credential/certificate/:id
[**WebsiteGetCredentialpasswordId**](WebsiteApi.md#WebsiteGetCredentialpasswordId) | **Get** /website/{websiteId}/credential/password/{id} | /credential/password/:id
[**WebsiteGetServicesServiceId**](WebsiteApi.md#WebsiteGetServicesServiceId) | **Get** /website/{websiteId}/services/{serviceId} | /services/:serviceId
[**WebsiteGetTag**](WebsiteApi.md#WebsiteGetTag) | **Get** /website/{websiteId}/tag | /tag
[**WebsiteList**](WebsiteApi.md#WebsiteList) | **Get** /website | List
[**WebsiteListAccessrights**](WebsiteApi.md#WebsiteListAccessrights) | **Get** /website/{websiteId}/accessrights | /accessrights
[**WebsiteListCredentialcertificate**](WebsiteApi.md#WebsiteListCredentialcertificate) | **Get** /website/{websiteId}/credential/certificate | /credential/certificate
[**WebsiteListCredentialpassword**](WebsiteApi.md#WebsiteListCredentialpassword) | **Get** /website/{websiteId}/credential/password | /credential/password
[**WebsiteListQueue**](WebsiteApi.md#WebsiteListQueue) | **Get** /website/{websiteId}/queue | /queue
[**WebsiteListServices**](WebsiteApi.md#WebsiteListServices) | **Get** /website/{websiteId}/services | /services
[**WebsitePatchCredentialcertificateId**](WebsiteApi.md#WebsitePatchCredentialcertificateId) | **Patch** /website/{websiteId}/credential/certificate/{id} | /credential/certificate/:id
[**WebsitePatchCredentialpasswordId**](WebsiteApi.md#WebsitePatchCredentialpasswordId) | **Patch** /website/{websiteId}/credential/password/{id} | /credential/password/:id
[**WebsitePatchTag**](WebsiteApi.md#WebsitePatchTag) | **Patch** /website/{websiteId}/tag | /tag
[**WebsitePostAccessrights**](WebsiteApi.md#WebsitePostAccessrights) | **Post** /website/{websiteId}/accessrights | /accessrights
[**WebsitePostCredentialcertificate**](WebsiteApi.md#WebsitePostCredentialcertificate) | **Post** /website/{websiteId}/credential/certificate | /credential/certificate
[**WebsitePostCredentialpassword**](WebsiteApi.md#WebsitePostCredentialpassword) | **Post** /website/{websiteId}/credential/password | /credential/password
[**WebsiteShow**](WebsiteApi.md#WebsiteShow) | **Get** /website/{websiteId} | Get
[**WebsiteUpdate**](WebsiteApi.md#WebsiteUpdate) | **Patch** /website/{websiteId} | Update



## WebsiteActionStart

> Website WebsiteActionStart(ctx, websiteId)
/actions/start

Action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteActionStop

> Website WebsiteActionStop(ctx, websiteId)
/actions/stop

Action stop

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteCreate

> Website WebsiteCreate(ctx, websiteCreate)
Create

Create website

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteCreate** | [**WebsiteCreate**](WebsiteCreate.md)|  | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteDelete

> WebsiteDelete(ctx, websiteId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

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


## WebsiteDeleteAccessrightsIdentity

> Website WebsiteDeleteAccessrightsIdentity(ctx, websiteId, identity)
/accessrights/:identity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**identity** | **string**| identity | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteDeleteCredentialcertificateId

> CredentialCertificate WebsiteDeleteCredentialcertificateId(ctx, websiteId, id)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
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


## WebsiteDeleteCredentialpasswordId

> Website WebsiteDeleteCredentialpasswordId(ctx, websiteId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**id** | **string**| id | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteDeleteTagKey

> map[string]interface{} WebsiteDeleteTagKey(ctx, websiteId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteGetCredentialcertificateId

> CredentialCertificate WebsiteGetCredentialcertificateId(ctx, websiteId, id)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
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


## WebsiteGetCredentialpasswordId

> CredentialPassword WebsiteGetCredentialpasswordId(ctx, websiteId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
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


## WebsiteGetServicesServiceId

> WebsiteServices WebsiteGetServicesServiceId(ctx, websiteId, serviceId)
/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**serviceId** | **string**| serviceId | 

### Return type

[**WebsiteServices**](website.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteGetTag

> map[string]interface{} WebsiteGetTag(ctx, websiteId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteList

> []Website WebsiteList(ctx, optional)
List

List website

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebsiteListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebsiteListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteListAccessrights

> []string WebsiteListAccessrights(ctx, websiteId)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

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


## WebsiteListCredentialcertificate

> []CredentialCertificate WebsiteListCredentialcertificate(ctx, websiteId)
/credential/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

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


## WebsiteListCredentialpassword

> []CredentialPassword WebsiteListCredentialpassword(ctx, websiteId)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

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


## WebsiteListQueue

> []Event WebsiteListQueue(ctx, websiteId)
/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

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


## WebsiteListServices

> []WebsiteServices WebsiteListServices(ctx, websiteId)
/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

### Return type

[**[]WebsiteServices**](website.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsitePatchCredentialcertificateId

> CredentialCertificate WebsitePatchCredentialcertificateId(ctx, websiteId, id, websitePatchCredentialcertificateId)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**id** | **string**| id | 
**websitePatchCredentialcertificateId** | [**WebsitePatchCredentialcertificateId**](WebsitePatchCredentialcertificateId.md)|  | 

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


## WebsitePatchCredentialpasswordId

> CredentialPassword WebsitePatchCredentialpasswordId(ctx, websiteId, id, websitePatchCredentialpasswordId)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**id** | **string**| id | 
**websitePatchCredentialpasswordId** | [**WebsitePatchCredentialpasswordId**](WebsitePatchCredentialpasswordId.md)|  | 

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


## WebsitePatchTag

> map[string]interface{} WebsitePatchTag(ctx, websiteId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsitePostAccessrights

> Website WebsitePostAccessrights(ctx, websiteId, websitePostAccessrights)
/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**websitePostAccessrights** | [**WebsitePostAccessrights**](WebsitePostAccessrights.md)|  | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsitePostCredentialcertificate

> CredentialCertificate WebsitePostCredentialcertificate(ctx, websiteId, websitePostCredentialcertificate)
/credential/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**websitePostCredentialcertificate** | [**WebsitePostCredentialcertificate**](WebsitePostCredentialcertificate.md)|  | 

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


## WebsitePostCredentialpassword

> CredentialPassword WebsitePostCredentialpassword(ctx, websiteId, websitePostCredentialpassword)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**websitePostCredentialpassword** | [**WebsitePostCredentialpassword**](WebsitePostCredentialpassword.md)|  | 

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


## WebsiteShow

> Website WebsiteShow(ctx, websiteId)
Get

Returns a single website

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebsiteUpdate

> Website WebsiteUpdate(ctx, websiteId, websiteUpdate)
Update

Returns modified website

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string**| ID of website | 
**websiteUpdate** | [**WebsiteUpdate**](WebsiteUpdate.md)|  | 

### Return type

[**Website**](website.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
