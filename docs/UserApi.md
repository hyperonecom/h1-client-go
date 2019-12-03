# \UserApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserDeleteCredentialId**](UserApi.md#UserDeleteCredentialId) | **Delete** /user/{userId}/credential/{id} | /credential/:id
[**UserDeleteCredentialauthtokenId**](UserApi.md#UserDeleteCredentialauthtokenId) | **Delete** /user/{userId}/credential/authtoken/:id | /credential/authtoken/:id
[**UserDeleteCredentialcertificateId**](UserApi.md#UserDeleteCredentialcertificateId) | **Delete** /user/{userId}/credential/certificate/{id} | /credential/certificate/:id
[**UserDeleteCredentialpasswordId**](UserApi.md#UserDeleteCredentialpasswordId) | **Delete** /user/{userId}/credential/password/{id} | /credential/password/:id
[**UserDeleteTagKey**](UserApi.md#UserDeleteTagKey) | **Delete** /user/{userId}/tag/{key} | /tag/:key
[**UserGetCredentialId**](UserApi.md#UserGetCredentialId) | **Get** /user/{userId}/credential/{id} | /credential/:id
[**UserGetCredentialauthtokenId**](UserApi.md#UserGetCredentialauthtokenId) | **Get** /user/{userId}/credential/authtoken/:id | /credential/authtoken/:id
[**UserGetCredentialcertificateId**](UserApi.md#UserGetCredentialcertificateId) | **Get** /user/{userId}/credential/certificate/{id} | /credential/certificate/:id
[**UserGetCredentialpasswordId**](UserApi.md#UserGetCredentialpasswordId) | **Get** /user/{userId}/credential/password/{id} | /credential/password/:id
[**UserGetServicesServiceId**](UserApi.md#UserGetServicesServiceId) | **Get** /user/{userId}/services/{serviceId} | /services/:serviceId
[**UserGetTag**](UserApi.md#UserGetTag) | **Get** /user/{userId}/tag | /tag
[**UserListCredential**](UserApi.md#UserListCredential) | **Get** /user/{userId}/credential | /credential
[**UserListCredentialauthtoken**](UserApi.md#UserListCredentialauthtoken) | **Get** /user/{userId}/credential/authtoken | /credential/authtoken
[**UserListCredentialcertificate**](UserApi.md#UserListCredentialcertificate) | **Get** /user/{userId}/credential/certificate | /credential/certificate
[**UserListCredentialpassword**](UserApi.md#UserListCredentialpassword) | **Get** /user/{userId}/credential/password | /credential/password
[**UserListInvitation**](UserApi.md#UserListInvitation) | **Get** /user/{userId}/invitation | /invitation
[**UserListServices**](UserApi.md#UserListServices) | **Get** /user/{userId}/services | /services
[**UserPatchCredentialId**](UserApi.md#UserPatchCredentialId) | **Patch** /user/{userId}/credential/{id} | /credential/:id
[**UserPatchCredentialcertificateId**](UserApi.md#UserPatchCredentialcertificateId) | **Patch** /user/{userId}/credential/certificate/{id} | /credential/certificate/:id
[**UserPatchCredentialpasswordId**](UserApi.md#UserPatchCredentialpasswordId) | **Patch** /user/{userId}/credential/password/{id} | /credential/password/:id
[**UserPatchTag**](UserApi.md#UserPatchTag) | **Patch** /user/{userId}/tag | /tag
[**UserPostCredential**](UserApi.md#UserPostCredential) | **Post** /user/{userId}/credential | /credential
[**UserPostCredentialcertificate**](UserApi.md#UserPostCredentialcertificate) | **Post** /user/{userId}/credential/certificate | /credential/certificate
[**UserPostCredentialpassword**](UserApi.md#UserPostCredentialpassword) | **Post** /user/{userId}/credential/password | /credential/password
[**UserPostInvitationInvitationIdaccept**](UserApi.md#UserPostInvitationInvitationIdaccept) | **Post** /user/{userId}/invitation/{invitationId}/accept | /invitation/:invitationId/accept
[**UserPostInvitationInvitationIddecline**](UserApi.md#UserPostInvitationInvitationIddecline) | **Post** /user/{userId}/invitation/{invitationId}/decline | /invitation/:invitationId/decline
[**UserPutTag**](UserApi.md#UserPutTag) | **Put** /user/{userId}/tag | /tag
[**UserShow**](UserApi.md#UserShow) | **Get** /user/{userId} | Get
[**UserUpdate**](UserApi.md#UserUpdate) | **Patch** /user/{userId} | Update



## UserDeleteCredentialId

> User UserDeleteCredentialId(ctx, userId, id)

/credential/:id

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


## UserDeleteCredentialauthtokenId

> UserDeleteCredentialauthtokenId(ctx, userId, invitationId)

/credential/authtoken/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**invitationId** | **string**| invitationId | 

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

> InlineResponse2001 UserDeleteCredentialcertificateId(ctx, userId, id)

/credential/certificate/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**id** | **string**| id | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

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

> map[string]string UserDeleteTagKey(ctx, userId, key)

/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserGetCredentialId

> CredentialPassword UserGetCredentialId(ctx, userId, id)

/credential/:id

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


## UserGetCredentialauthtokenId

> AuthToken UserGetCredentialauthtokenId(ctx, userId, invitationId)

/credential/authtoken/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**invitationId** | **string**| invitationId | 

### Return type

[**AuthToken**](authToken.md)

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


## UserGetServicesServiceId

> UserServices UserGetServicesServiceId(ctx, userId, serviceId)

/services/:serviceId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**serviceId** | **string**| serviceId | 

### Return type

[**UserServices**](user.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetTag

> map[string]string UserGetTag(ctx, userId)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

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


## UserListCredential

> []CredentialPassword UserListCredential(ctx, userId)

/credential

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


## UserListCredentialauthtoken

> []AuthToken UserListCredentialauthtoken(ctx, userId)

/credential/authtoken

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

### Return type

[**[]AuthToken**](authToken.md)

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


## UserListInvitation

> []InlineResponse200 UserListInvitation(ctx, userId)

/invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserListServices

> []UserServices UserListServices(ctx, userId)

/services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 

### Return type

[**[]UserServices**](user.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPatchCredentialId

> CredentialPassword UserPatchCredentialId(ctx, userId, id, userPatchCredentialId)

/credential/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**id** | **string**| id | 
**userPatchCredentialId** | [**UserPatchCredentialId**](UserPatchCredentialId.md)|  | 

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

> map[string]string UserPatchTag(ctx, userId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserPostCredential

> CredentialPassword UserPostCredential(ctx, userId, userPostCredential)

/credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**userPostCredential** | [**UserPostCredential**](UserPostCredential.md)|  | 

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


## UserPostInvitationInvitationIdaccept

> UserPostInvitationInvitationIdaccept(ctx, userId, invitationId)

/invitation/:invitationId/accept

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**invitationId** | **string**| invitationId | 

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


## UserPostInvitationInvitationIddecline

> UserPostInvitationInvitationIddecline(ctx, userId, invitationId)

/invitation/:invitationId/decline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**invitationId** | **string**| invitationId | 

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


## UserPutTag

> map[string]string UserPutTag(ctx, userId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
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


## UserUpdate

> User UserUpdate(ctx, userId, userUpdate)

Update

Returns modified user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| ID of user | 
**userUpdate** | [**UserUpdate**](UserUpdate.md)|  | 

### Return type

[**User**](user.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

