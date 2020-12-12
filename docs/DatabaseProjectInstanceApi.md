# \DatabaseProjectInstanceApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatabaseProjectInstanceConnectGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceConnectGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/connect/{connectId} | Get database/instance.connect
[**DatabaseProjectInstanceConnectList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceConnectList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/connect | List database/instance.connect
[**DatabaseProjectInstanceCreate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCreate) | **Post** /database/{locationId}/project/{projectId}/instance | Create database/instance
[**DatabaseProjectInstanceCredentialCreate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialCreate) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential | Create database/instance.credential
[**DatabaseProjectInstanceCredentialDelete**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialDelete) | **Delete** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential/{credentialId} | Delete database/instance.credential
[**DatabaseProjectInstanceCredentialGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential/{credentialId} | Get database/instance.credential
[**DatabaseProjectInstanceCredentialList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential | List database/instance.credential
[**DatabaseProjectInstanceCredentialPatch**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialPatch) | **Patch** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential/{credentialId} | Update database/instance.credential
[**DatabaseProjectInstanceDelete**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceDelete) | **Delete** /database/{locationId}/project/{projectId}/instance/{instanceId} | Delete database/instance
[**DatabaseProjectInstanceEventGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceEventGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/event/{eventId} | Get database/instance.event
[**DatabaseProjectInstanceEventList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceEventList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/event | List database/instance.event
[**DatabaseProjectInstanceGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId} | Get database/instance
[**DatabaseProjectInstanceList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceList) | **Get** /database/{locationId}/project/{projectId}/instance | List database/instance
[**DatabaseProjectInstanceServiceGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceServiceGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/service/{serviceId} | Get database/instance.service
[**DatabaseProjectInstanceServiceList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceServiceList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/service | List database/instance.service
[**DatabaseProjectInstanceStart**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceStart) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/actions/start | Start database/instance
[**DatabaseProjectInstanceStop**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceStop) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/actions/stop | Stop database/instance
[**DatabaseProjectInstanceTagCreate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagCreate) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag | Create database/instance.tag
[**DatabaseProjectInstanceTagDelete**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagDelete) | **Delete** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Delete database/instance.tag
[**DatabaseProjectInstanceTagGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Get database/instance.tag
[**DatabaseProjectInstanceTagList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag | List database/instance.tag
[**DatabaseProjectInstanceTagPut**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagPut) | **Put** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag | Replace database/instance.tag
[**DatabaseProjectInstanceTransfer**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTransfer) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/actions/transfer | Transfer database/instance
[**DatabaseProjectInstanceUpdate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceUpdate) | **Patch** /database/{locationId}/project/{projectId}/instance/{instanceId} | Update database/instance



## DatabaseProjectInstanceConnectGet

> ResourceConnect DatabaseProjectInstanceConnectGet(ctx, projectId, locationId, instanceId, connectId)

Get database/instance.connect

Get database/instance.connect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
**connectId** | **string**| connectId | 

### Return type

[**ResourceConnect**](resource.connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceConnectList

> []ResourceConnect DatabaseProjectInstanceConnectList(ctx, projectId, locationId, instanceId)

List database/instance.connect

List database/instance.connect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

### Return type

[**[]ResourceConnect**](resource.connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCreate

> Database DatabaseProjectInstanceCreate(ctx, projectId, locationId, databaseProjectInstanceCreate, optional)

Create database/instance

Create instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**databaseProjectInstanceCreate** | [**DatabaseProjectInstanceCreate**](DatabaseProjectInstanceCreate.md)|  | 
 **optional** | ***DatabaseProjectInstanceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatabaseProjectInstanceCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialCreate

> DatabaseCredential DatabaseProjectInstanceCredentialCreate(ctx, projectId, locationId, instanceId, databaseCredential)

Create database/instance.credential

Create database/instance.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
**databaseCredential** | [**DatabaseCredential**](DatabaseCredential.md)|  | 

### Return type

[**DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialDelete

> Database DatabaseProjectInstanceCredentialDelete(ctx, projectId, locationId, instanceId, credentialId)

Delete database/instance.credential

Delete database/instance.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialGet

> DatabaseCredential DatabaseProjectInstanceCredentialGet(ctx, projectId, locationId, instanceId, credentialId)

Get database/instance.credential

Get database/instance.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialList

> []DatabaseCredential DatabaseProjectInstanceCredentialList(ctx, projectId, locationId, instanceId)

List database/instance.credential

List database/instance.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

### Return type

[**[]DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialPatch

> DatabaseCredential DatabaseProjectInstanceCredentialPatch(ctx, projectId, locationId, instanceId, credentialId, databaseProjectInstanceCredentialPatch)

Update database/instance.credential

Update database/instance.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
**credentialId** | **string**| credentialId | 
**databaseProjectInstanceCredentialPatch** | [**DatabaseProjectInstanceCredentialPatch**](DatabaseProjectInstanceCredentialPatch.md)|  | 

### Return type

[**DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceDelete

> DatabaseProjectInstanceDelete(ctx, projectId, locationId, instanceId)

Delete database/instance

Delete instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

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


## DatabaseProjectInstanceEventGet

> Event DatabaseProjectInstanceEventGet(ctx, projectId, locationId, instanceId, eventId)

Get database/instance.event

Get database/instance.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## DatabaseProjectInstanceEventList

> []Event DatabaseProjectInstanceEventList(ctx, projectId, locationId, instanceId, optional)

List database/instance.event

List database/instance.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
 **optional** | ***DatabaseProjectInstanceEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatabaseProjectInstanceEventListOpts struct


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


## DatabaseProjectInstanceGet

> Database DatabaseProjectInstanceGet(ctx, projectId, locationId, instanceId)

Get database/instance

Returns a single instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceList

> []Database DatabaseProjectInstanceList(ctx, projectId, locationId, optional)

List database/instance

List instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***DatabaseProjectInstanceListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatabaseProjectInstanceListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceServiceGet

> ResourceService DatabaseProjectInstanceServiceGet(ctx, projectId, locationId, instanceId, serviceId)

Get database/instance.service

Get database/instance.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## DatabaseProjectInstanceServiceList

> []ResourceService DatabaseProjectInstanceServiceList(ctx, projectId, locationId, instanceId)

List database/instance.service

List database/instance.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

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


## DatabaseProjectInstanceStart

> Database DatabaseProjectInstanceStart(ctx, projectId, locationId, instanceId, optional)

Start database/instance

action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
 **optional** | ***DatabaseProjectInstanceStartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatabaseProjectInstanceStartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceStop

> Database DatabaseProjectInstanceStop(ctx, projectId, locationId, instanceId, optional)

Stop database/instance

action stop

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
 **optional** | ***DatabaseProjectInstanceStopOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatabaseProjectInstanceStopOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceTagCreate

> Tag DatabaseProjectInstanceTagCreate(ctx, projectId, locationId, instanceId, tag)

Create database/instance.tag

Create database/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## DatabaseProjectInstanceTagDelete

> DatabaseProjectInstanceTagDelete(ctx, projectId, locationId, instanceId, tagId)

Delete database/instance.tag

Delete database/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## DatabaseProjectInstanceTagGet

> Tag DatabaseProjectInstanceTagGet(ctx, projectId, locationId, instanceId, tagId)

Get database/instance.tag

Get database/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## DatabaseProjectInstanceTagList

> []Tag DatabaseProjectInstanceTagList(ctx, projectId, locationId, instanceId)

List database/instance.tag

List database/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 

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


## DatabaseProjectInstanceTagPut

> []Tag DatabaseProjectInstanceTagPut(ctx, projectId, locationId, instanceId, tag)

Replace database/instance.tag

Replace database/instance.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
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


## DatabaseProjectInstanceTransfer

> Database DatabaseProjectInstanceTransfer(ctx, projectId, locationId, instanceId, databaseProjectInstanceTransfer, optional)

Transfer database/instance

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
**databaseProjectInstanceTransfer** | [**DatabaseProjectInstanceTransfer**](DatabaseProjectInstanceTransfer.md)|  | 
 **optional** | ***DatabaseProjectInstanceTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatabaseProjectInstanceTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceUpdate

> Database DatabaseProjectInstanceUpdate(ctx, projectId, locationId, instanceId, databaseProjectInstanceUpdate)

Update database/instance

Returns modified instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**instanceId** | **string**| Instance Id | 
**databaseProjectInstanceUpdate** | [**DatabaseProjectInstanceUpdate**](DatabaseProjectInstanceUpdate.md)|  | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

