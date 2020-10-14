# \IamUserApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamUserCredentialAuthtokenDelete**](IamUserApi.md#IamUserCredentialAuthtokenDelete) | **Delete** /iam/user/{userId}/credential/authtoken/{authtokenId} | Delete iam/user.credential
[**IamUserCredentialAuthtokenGet**](IamUserApi.md#IamUserCredentialAuthtokenGet) | **Get** /iam/user/{userId}/credential/authtoken/{authtokenId} | Get iam/user.credential
[**IamUserCredentialAuthtokenList**](IamUserApi.md#IamUserCredentialAuthtokenList) | **Get** /iam/user/{userId}/credential/authtoken | List iam/user.credential
[**IamUserCredentialCreate**](IamUserApi.md#IamUserCredentialCreate) | **Post** /iam/user/{userId}/credential | Create iam/user.credential
[**IamUserCredentialDelete**](IamUserApi.md#IamUserCredentialDelete) | **Delete** /iam/user/{userId}/credential/{credentialId} | Delete iam/user.credential
[**IamUserCredentialGet**](IamUserApi.md#IamUserCredentialGet) | **Get** /iam/user/{userId}/credential/{credentialId} | Get iam/user.credential
[**IamUserCredentialList**](IamUserApi.md#IamUserCredentialList) | **Get** /iam/user/{userId}/credential | List iam/user.credential
[**IamUserCredentialPatch**](IamUserApi.md#IamUserCredentialPatch) | **Patch** /iam/user/{userId}/credential/{credentialId} | Update iam/user.credential
[**IamUserGet**](IamUserApi.md#IamUserGet) | **Get** /iam/user/{userId} | Get iam/user
[**IamUserServiceGet**](IamUserApi.md#IamUserServiceGet) | **Get** /iam/user/{userId}/service/{serviceId} | Get iam/user.service
[**IamUserServiceList**](IamUserApi.md#IamUserServiceList) | **Get** /iam/user/{userId}/service | List iam/user.service
[**IamUserUpdate**](IamUserApi.md#IamUserUpdate) | **Patch** /iam/user/{userId} | Update iam/user



## IamUserCredentialAuthtokenDelete

> IamUserCredentialAuthtokenDelete(ctx, userId, authtokenId)

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


## IamUserCredentialAuthtokenGet

> AuthToken IamUserCredentialAuthtokenGet(ctx, userId, authtokenId)

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


## IamUserCredentialAuthtokenList

> []AuthToken IamUserCredentialAuthtokenList(ctx, userId)

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


## IamUserCredentialCreate

> UserCredential IamUserCredentialCreate(ctx, userId, userCredential)

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


## IamUserCredentialDelete

> User IamUserCredentialDelete(ctx, userId, credentialId)

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


## IamUserCredentialGet

> UserCredential IamUserCredentialGet(ctx, userId, credentialId)

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


## IamUserCredentialList

> []UserCredential IamUserCredentialList(ctx, userId)

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


## IamUserCredentialPatch

> UserCredential IamUserCredentialPatch(ctx, userId, credentialId, iamUserCredentialPatch)

Update iam/user.credential

Update iam/user.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**credentialId** | **string**| credentialId | 
**iamUserCredentialPatch** | [**IamUserCredentialPatch**](IamUserCredentialPatch.md)|  | 

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


## IamUserGet

> User IamUserGet(ctx, userId)

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


## IamUserServiceGet

> ResourceService IamUserServiceGet(ctx, userId, serviceId)

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


## IamUserServiceList

> []ResourceService IamUserServiceList(ctx, userId)

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


## IamUserUpdate

> User IamUserUpdate(ctx, userId, iamUserUpdate)

Update iam/user

Returns modified user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User Id | 
**iamUserUpdate** | [**IamUserUpdate**](IamUserUpdate.md)|  | 

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

