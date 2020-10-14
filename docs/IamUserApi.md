# \IamUserApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCredentialAuthtokenDelete**](IamUserApi.md#UserCredentialAuthtokenDelete) | **Delete** /iam/user/{userId}/credential/authtoken/{authtokenId} | Delete iam/user.credential
[**UserCredentialAuthtokenGet**](IamUserApi.md#UserCredentialAuthtokenGet) | **Get** /iam/user/{userId}/credential/authtoken/{authtokenId} | Get iam/user.credential
[**UserCredentialAuthtokenList**](IamUserApi.md#UserCredentialAuthtokenList) | **Get** /iam/user/{userId}/credential/authtoken | List iam/user.credential
[**UserCredentialCreate**](IamUserApi.md#UserCredentialCreate) | **Post** /iam/user/{userId}/credential | Create iam/user.credential
[**UserCredentialDelete**](IamUserApi.md#UserCredentialDelete) | **Delete** /iam/user/{userId}/credential/{credentialId} | Delete iam/user.credential
[**UserCredentialGet**](IamUserApi.md#UserCredentialGet) | **Get** /iam/user/{userId}/credential/{credentialId} | Get iam/user.credential
[**UserCredentialList**](IamUserApi.md#UserCredentialList) | **Get** /iam/user/{userId}/credential | List iam/user.credential
[**UserCredentialPatch**](IamUserApi.md#UserCredentialPatch) | **Patch** /iam/user/{userId}/credential/{credentialId} | Update iam/user.credential
[**UserGet**](IamUserApi.md#UserGet) | **Get** /iam/user/{userId} | Get iam/user
[**UserServiceGet**](IamUserApi.md#UserServiceGet) | **Get** /iam/user/{userId}/service/{serviceId} | Get iam/user.service
[**UserServiceList**](IamUserApi.md#UserServiceList) | **Get** /iam/user/{userId}/service | List iam/user.service
[**UserUpdate**](IamUserApi.md#UserUpdate) | **Patch** /iam/user/{userId} | Update iam/user



## UserCredentialAuthtokenDelete

> UserCredentialAuthtokenDelete(ctx, userId, authtokenId)

Delete iam/user.credential

Delete iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**authtokenId** | **string**| authtokenId | 

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


## UserCredentialAuthtokenGet

> AuthToken UserCredentialAuthtokenGet(ctx, userId, authtokenId)

Get iam/user.credential

Get iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**authtokenId** | **string**| authtokenId | 

### Return type

[**AuthToken**](authToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialAuthtokenList

> []AuthToken UserCredentialAuthtokenList(ctx, userId)

List iam/user.credential

List iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 

### Return type

[**[]AuthToken**](authToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialCreate

> UserCredential UserCredentialCreate(ctx, userId, userCredential)

Create iam/user.credential

Create iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**userCredential** | [**UserCredential**](UserCredential.md)|  | 

### Return type

[**UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialDelete

> User UserCredentialDelete(ctx, userId, credentialId)

Delete iam/user.credential

Delete iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**User**](user.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialGet

> UserCredential UserCredentialGet(ctx, userId, credentialId)

Get iam/user.credential

Get iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialList

> []UserCredential UserCredentialList(ctx, userId)

List iam/user.credential

List iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 

### Return type

[**[]UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialPatch

> UserCredential UserCredentialPatch(ctx, userId, credentialId, userCredentialPatch)

Update iam/user.credential

Update iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**credentialId** | **string**| credentialId | 
**userCredentialPatch** | [**UserCredentialPatch**](UserCredentialPatch.md)|  | 

### Return type

[**UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGet

> User UserGet(ctx, userId)

Get iam/user

Returns a single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 

### Return type

[**User**](user.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserServiceGet

> ResourceService UserServiceGet(ctx, userId, serviceId)

Get iam/user.service

Get iam/user.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
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


## UserServiceList

> []ResourceService UserServiceList(ctx, userId)

List iam/user.service

List iam/user.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 

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


## UserUpdate

> User UserUpdate(ctx, userId, userUpdate)

Update iam/user

Returns modified user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**userUpdate** | [**UserUpdate**](UserUpdate.md)|  | 

### Return type

[**User**](user.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

