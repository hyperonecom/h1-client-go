# \IamProjectSaApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectSaCreate**](IamProjectSaApi.md#IamProjectSaCreate) | **Post** /iam/project/{projectId}/sa | Create iam/sa
[**IamProjectSaCredentialCreate**](IamProjectSaApi.md#IamProjectSaCredentialCreate) | **Post** /iam/project/{projectId}/sa/{saId}/credential | Create iam/sa.credential
[**IamProjectSaCredentialDelete**](IamProjectSaApi.md#IamProjectSaCredentialDelete) | **Delete** /iam/project/{projectId}/sa/{saId}/credential/{credentialId} | Delete iam/sa.credential
[**IamProjectSaCredentialGet**](IamProjectSaApi.md#IamProjectSaCredentialGet) | **Get** /iam/project/{projectId}/sa/{saId}/credential/{credentialId} | Get iam/sa.credential
[**IamProjectSaCredentialList**](IamProjectSaApi.md#IamProjectSaCredentialList) | **Get** /iam/project/{projectId}/sa/{saId}/credential | List iam/sa.credential
[**IamProjectSaCredentialPatch**](IamProjectSaApi.md#IamProjectSaCredentialPatch) | **Patch** /iam/project/{projectId}/sa/{saId}/credential/{credentialId} | Update iam/sa.credential
[**IamProjectSaDelete**](IamProjectSaApi.md#IamProjectSaDelete) | **Delete** /iam/project/{projectId}/sa/{saId} | Delete iam/sa
[**IamProjectSaEventGet**](IamProjectSaApi.md#IamProjectSaEventGet) | **Get** /iam/project/{projectId}/sa/{saId}/event/{eventId} | Get iam/sa.event
[**IamProjectSaEventList**](IamProjectSaApi.md#IamProjectSaEventList) | **Get** /iam/project/{projectId}/sa/{saId}/event | List iam/sa.event
[**IamProjectSaGet**](IamProjectSaApi.md#IamProjectSaGet) | **Get** /iam/project/{projectId}/sa/{saId} | Get iam/sa
[**IamProjectSaList**](IamProjectSaApi.md#IamProjectSaList) | **Get** /iam/project/{projectId}/sa | List iam/sa
[**IamProjectSaServiceGet**](IamProjectSaApi.md#IamProjectSaServiceGet) | **Get** /iam/project/{projectId}/sa/{saId}/service/{serviceId} | Get iam/sa.service
[**IamProjectSaServiceList**](IamProjectSaApi.md#IamProjectSaServiceList) | **Get** /iam/project/{projectId}/sa/{saId}/service | List iam/sa.service
[**IamProjectSaTagCreate**](IamProjectSaApi.md#IamProjectSaTagCreate) | **Post** /iam/project/{projectId}/sa/{saId}/tag | Create iam/sa.tag
[**IamProjectSaTagDelete**](IamProjectSaApi.md#IamProjectSaTagDelete) | **Delete** /iam/project/{projectId}/sa/{saId}/tag/{tagId} | Delete iam/sa.tag
[**IamProjectSaTagGet**](IamProjectSaApi.md#IamProjectSaTagGet) | **Get** /iam/project/{projectId}/sa/{saId}/tag/{tagId} | Get iam/sa.tag
[**IamProjectSaTagList**](IamProjectSaApi.md#IamProjectSaTagList) | **Get** /iam/project/{projectId}/sa/{saId}/tag | List iam/sa.tag
[**IamProjectSaTagPut**](IamProjectSaApi.md#IamProjectSaTagPut) | **Put** /iam/project/{projectId}/sa/{saId}/tag | Replace iam/sa.tag
[**IamProjectSaUpdate**](IamProjectSaApi.md#IamProjectSaUpdate) | **Patch** /iam/project/{projectId}/sa/{saId} | Update iam/sa



## IamProjectSaCreate

> Sa IamProjectSaCreate(ctx, projectId, iamProjectSaCreate, optional)

Create iam/sa

Create sa

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**iamProjectSaCreate** | [**IamProjectSaCreate**](IamProjectSaCreate.md)|  | 
 **optional** | ***IamProjectSaCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectSaCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Sa**](sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialCreate

> SaCredential IamProjectSaCredentialCreate(ctx, projectId, saId, saCredential)

Create iam/sa.credential

Create iam/sa.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**saCredential** | [**SaCredential**](SaCredential.md)|  | 

### Return type

[**SaCredential**](sa.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialDelete

> Sa IamProjectSaCredentialDelete(ctx, projectId, saId, credentialId)

Delete iam/sa.credential

Delete iam/sa.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**Sa**](sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialGet

> SaCredential IamProjectSaCredentialGet(ctx, projectId, saId, credentialId)

Get iam/sa.credential

Get iam/sa.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**SaCredential**](sa.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialList

> []SaCredential IamProjectSaCredentialList(ctx, projectId, saId)

List iam/sa.credential

List iam/sa.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 

### Return type

[**[]SaCredential**](sa.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialPatch

> SaCredential IamProjectSaCredentialPatch(ctx, projectId, saId, credentialId, iamProjectSaCredentialPatch)

Update iam/sa.credential

Update iam/sa.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**credentialId** | **string**| credentialId | 
**iamProjectSaCredentialPatch** | [**IamProjectSaCredentialPatch**](IamProjectSaCredentialPatch.md)|  | 

### Return type

[**SaCredential**](sa.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaDelete

> IamProjectSaDelete(ctx, projectId, saId)

Delete iam/sa

Delete sa

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 

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


## IamProjectSaEventGet

> Event IamProjectSaEventGet(ctx, projectId, saId, eventId)

Get iam/sa.event

Get iam/sa.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**eventId** | **string**| eventId | 

### Return type

[**Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaEventList

> []Event IamProjectSaEventList(ctx, projectId, saId, optional)

List iam/sa.event

List iam/sa.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
 **optional** | ***IamProjectSaEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectSaEventListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Float32**| $limit | [default to 100]
 **skip** | **optional.Float32**| $skip | 

### Return type

[**[]Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaGet

> Sa IamProjectSaGet(ctx, projectId, saId)

Get iam/sa

Returns a single sa

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 

### Return type

[**Sa**](sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaList

> []Sa IamProjectSaList(ctx, projectId, optional)

List iam/sa

List sa

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***IamProjectSaListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectSaListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Sa**](sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaServiceGet

> ResourceService IamProjectSaServiceGet(ctx, projectId, saId, serviceId)

Get iam/sa.service

Get iam/sa.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
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


## IamProjectSaServiceList

> []ResourceService IamProjectSaServiceList(ctx, projectId, saId)

List iam/sa.service

List iam/sa.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 

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


## IamProjectSaTagCreate

> Tag IamProjectSaTagCreate(ctx, projectId, saId, tag)

Create iam/sa.tag

Create iam/sa.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**tag** | [**Tag**](Tag.md)|  | 

### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaTagDelete

> IamProjectSaTagDelete(ctx, projectId, saId, tagId)

Delete iam/sa.tag

Delete iam/sa.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**tagId** | **string**| tagId | 

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


## IamProjectSaTagGet

> Tag IamProjectSaTagGet(ctx, projectId, saId, tagId)

Get iam/sa.tag

Get iam/sa.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**tagId** | **string**| tagId | 

### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaTagList

> []Tag IamProjectSaTagList(ctx, projectId, saId)

List iam/sa.tag

List iam/sa.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 

### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaTagPut

> []Tag IamProjectSaTagPut(ctx, projectId, saId, tag)

Replace iam/sa.tag

Replace iam/sa.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**tag** | [**[]Tag**](tag.md)|  | 

### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaUpdate

> Sa IamProjectSaUpdate(ctx, projectId, saId, iamProjectSaUpdate)

Update iam/sa

Returns modified sa

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**saId** | **string**| Sa Id | 
**iamProjectSaUpdate** | [**IamProjectSaUpdate**](IamProjectSaUpdate.md)|  | 

### Return type

[**Sa**](sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

