# \UserApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserDelete**](UserApi.md#UserDelete) | **Delete** /user/{userId} | Delete
[**UserDeleteCredentialcertificateId**](UserApi.md#UserDeleteCredentialcertificateId) | **Delete** /user/{userId}/credential/certificate/{id} | /credential/certificate/:id
[**UserDeleteCredentialpasswordId**](UserApi.md#UserDeleteCredentialpasswordId) | **Delete** /user/{userId}/credential/password/{id} | /credential/password/:id
[**UserDeleteTagKey**](UserApi.md#UserDeleteTagKey) | **Delete** /user/{userId}/tag/{key} | /tag/:key
[**UserGetCredentialcertificateId**](UserApi.md#UserGetCredentialcertificateId) | **Get** /user/{userId}/credential/certificate/{id} | /credential/certificate/:id
[**UserGetCredentialpasswordId**](UserApi.md#UserGetCredentialpasswordId) | **Get** /user/{userId}/credential/password/{id} | /credential/password/:id
[**UserGetTag**](UserApi.md#UserGetTag) | **Get** /user/{userId}/tag | /tag
[**UserList**](UserApi.md#UserList) | **Get** /user | List
[**UserListCredentialcertificate**](UserApi.md#UserListCredentialcertificate) | **Get** /user/{userId}/credential/certificate | /credential/certificate
[**UserListCredentialpassword**](UserApi.md#UserListCredentialpassword) | **Get** /user/{userId}/credential/password | /credential/password
[**UserPatchCredentialcertificateId**](UserApi.md#UserPatchCredentialcertificateId) | **Patch** /user/{userId}/credential/certificate/{id} | /credential/certificate/:id
[**UserPatchCredentialpasswordId**](UserApi.md#UserPatchCredentialpasswordId) | **Patch** /user/{userId}/credential/password/{id} | /credential/password/:id
[**UserPatchTag**](UserApi.md#UserPatchTag) | **Patch** /user/{userId}/tag | /tag
[**UserPostCredentialcertificate**](UserApi.md#UserPostCredentialcertificate) | **Post** /user/{userId}/credential/certificate | /credential/certificate
[**UserPostCredentialpassword**](UserApi.md#UserPostCredentialpassword) | **Post** /user/{userId}/credential/password | /credential/password
[**UserShow**](UserApi.md#UserShow) | **Get** /user/{userId} | Get



## UserDelete

> UserDelete(ctx, userId)
Delete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

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


## UserDeleteCredentialcertificateId

> CredentialCertificate UserDeleteCredentialcertificateId(ctx, userId, id)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserDeleteCredentialpasswordId

> User UserDeleteCredentialpasswordId(ctx, userId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**id** | **string**| id | 

### Return type

[**User**](user.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeleteTagKey

> map[string]interface{} UserDeleteTagKey(ctx, userId, key)
/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserGetCredentialcertificateId

> CredentialCertificate UserGetCredentialcertificateId(ctx, userId, id)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserGetCredentialpasswordId

> CredentialPassword UserGetCredentialpasswordId(ctx, userId, id)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserGetTag

> map[string]interface{} UserGetTag(ctx, userId)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

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


## UserList

> []User UserList(ctx, )
List

List user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]User**](user.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserListCredentialcertificate

> []CredentialCertificate UserListCredentialcertificate(ctx, userId)
/credential/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

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


## UserListCredentialpassword

> []CredentialPassword UserListCredentialpassword(ctx, userId)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

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


## UserPatchCredentialcertificateId

> CredentialCertificate UserPatchCredentialcertificateId(ctx, userId, id, userPatchCredentialcertificateId)
/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**id** | **string**| id | 
**userPatchCredentialcertificateId** | [**UserPatchCredentialcertificateId**](UserPatchCredentialcertificateId.md)|  | 

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


## UserPatchCredentialpasswordId

> CredentialPassword UserPatchCredentialpasswordId(ctx, userId, id, userPatchCredentialpasswordId)
/credential/password/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**id** | **string**| id | 
**userPatchCredentialpasswordId** | [**UserPatchCredentialpasswordId**](UserPatchCredentialpasswordId.md)|  | 

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


## UserPatchTag

> map[string]interface{} UserPatchTag(ctx, userId, requestBody)
/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserPostCredentialcertificate

> CredentialCertificate UserPostCredentialcertificate(ctx, userId, userPostCredentialcertificate)
/credential/certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**userPostCredentialcertificate** | [**UserPostCredentialcertificate**](UserPostCredentialcertificate.md)|  | 

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


## UserPostCredentialpassword

> CredentialPassword UserPostCredentialpassword(ctx, userId, userPostCredentialpassword)
/credential/password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**userPostCredentialpassword** | [**UserPostCredentialpassword**](UserPostCredentialpassword.md)|  | 

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


## UserShow

> User UserShow(ctx, userId)
Get

Returns a single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

### Return type

[**User**](user.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
