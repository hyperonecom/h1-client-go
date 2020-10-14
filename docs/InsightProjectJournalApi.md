# \InsightProjectJournalApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsightProjectJournalCreate**](InsightProjectJournalApi.md#InsightProjectJournalCreate) | **Post** /insight/{locationId}/project/{projectId}/journal | Create insight/journal
[**InsightProjectJournalCredentialCreate**](InsightProjectJournalApi.md#InsightProjectJournalCredentialCreate) | **Post** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential | Create insight/journal.credential
[**InsightProjectJournalCredentialDelete**](InsightProjectJournalApi.md#InsightProjectJournalCredentialDelete) | **Delete** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential/{credentialId} | Delete insight/journal.credential
[**InsightProjectJournalCredentialGet**](InsightProjectJournalApi.md#InsightProjectJournalCredentialGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential/{credentialId} | Get insight/journal.credential
[**InsightProjectJournalCredentialList**](InsightProjectJournalApi.md#InsightProjectJournalCredentialList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential | List insight/journal.credential
[**InsightProjectJournalCredentialPatch**](InsightProjectJournalApi.md#InsightProjectJournalCredentialPatch) | **Patch** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential/{credentialId} | Update insight/journal.credential
[**InsightProjectJournalDelete**](InsightProjectJournalApi.md#InsightProjectJournalDelete) | **Delete** /insight/{locationId}/project/{projectId}/journal/{journalId} | Delete insight/journal
[**InsightProjectJournalEventGet**](InsightProjectJournalApi.md#InsightProjectJournalEventGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/event/{eventId} | Get insight/journal.event
[**InsightProjectJournalEventList**](InsightProjectJournalApi.md#InsightProjectJournalEventList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/event | List insight/journal.event
[**InsightProjectJournalGet**](InsightProjectJournalApi.md#InsightProjectJournalGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId} | Get insight/journal
[**InsightProjectJournalList**](InsightProjectJournalApi.md#InsightProjectJournalList) | **Get** /insight/{locationId}/project/{projectId}/journal | List insight/journal
[**InsightProjectJournalLogGet**](InsightProjectJournalApi.md#InsightProjectJournalLogGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/log | Get insight/journal.log
[**InsightProjectJournalServiceGet**](InsightProjectJournalApi.md#InsightProjectJournalServiceGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/service/{serviceId} | Get insight/journal.service
[**InsightProjectJournalServiceList**](InsightProjectJournalApi.md#InsightProjectJournalServiceList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/service | List insight/journal.service
[**InsightProjectJournalTagCreate**](InsightProjectJournalApi.md#InsightProjectJournalTagCreate) | **Post** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag | Create insight/journal.tag
[**InsightProjectJournalTagDelete**](InsightProjectJournalApi.md#InsightProjectJournalTagDelete) | **Delete** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag/{tagId} | Delete insight/journal.tag
[**InsightProjectJournalTagGet**](InsightProjectJournalApi.md#InsightProjectJournalTagGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag/{tagId} | Get insight/journal.tag
[**InsightProjectJournalTagList**](InsightProjectJournalApi.md#InsightProjectJournalTagList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag | List insight/journal.tag
[**InsightProjectJournalTagPut**](InsightProjectJournalApi.md#InsightProjectJournalTagPut) | **Put** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag | Replace insight/journal.tag
[**InsightProjectJournalTransfer**](InsightProjectJournalApi.md#InsightProjectJournalTransfer) | **Post** /insight/{locationId}/project/{projectId}/journal/{journalId}/actions/transfer | Transfer insight/journal
[**InsightProjectJournalUpdate**](InsightProjectJournalApi.md#InsightProjectJournalUpdate) | **Patch** /insight/{locationId}/project/{projectId}/journal/{journalId} | Update insight/journal



## InsightProjectJournalCreate

> Journal InsightProjectJournalCreate(ctx, projectId, locationId, insightProjectJournalCreate, optional)

Create insight/journal

Create journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**insightProjectJournalCreate** | [**InsightProjectJournalCreate**](InsightProjectJournalCreate.md)|  | 
 **optional** | ***InsightProjectJournalCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InsightProjectJournalCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialCreate

> JournalCredential InsightProjectJournalCredentialCreate(ctx, projectId, locationId, journalId, journalCredential)

Create insight/journal.credential

Create insight/journal.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
**journalCredential** | [**JournalCredential**](JournalCredential.md)|  | 

### Return type

[**JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialDelete

> Journal InsightProjectJournalCredentialDelete(ctx, projectId, locationId, journalId, credentialId)

Delete insight/journal.credential

Delete insight/journal.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialGet

> JournalCredential InsightProjectJournalCredentialGet(ctx, projectId, locationId, journalId, credentialId)

Get insight/journal.credential

Get insight/journal.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialList

> []JournalCredential InsightProjectJournalCredentialList(ctx, projectId, locationId, journalId)

List insight/journal.credential

List insight/journal.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 

### Return type

[**[]JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialPatch

> JournalCredential InsightProjectJournalCredentialPatch(ctx, projectId, locationId, journalId, credentialId, insightProjectJournalCredentialPatch)

Update insight/journal.credential

Update insight/journal.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
**credentialId** | **string**| credentialId | 
**insightProjectJournalCredentialPatch** | [**InsightProjectJournalCredentialPatch**](InsightProjectJournalCredentialPatch.md)|  | 

### Return type

[**JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalDelete

> InsightProjectJournalDelete(ctx, projectId, locationId, journalId)

Delete insight/journal

Delete journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 

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


## InsightProjectJournalEventGet

> Event InsightProjectJournalEventGet(ctx, projectId, locationId, journalId, eventId)

Get insight/journal.event

Get insight/journal.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
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


## InsightProjectJournalEventList

> []Event InsightProjectJournalEventList(ctx, projectId, locationId, journalId, optional)

List insight/journal.event

List insight/journal.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
 **optional** | ***InsightProjectJournalEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InsightProjectJournalEventListOpts struct


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


## InsightProjectJournalGet

> Journal InsightProjectJournalGet(ctx, projectId, locationId, journalId)

Get insight/journal

Returns a single journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalList

> []Journal InsightProjectJournalList(ctx, projectId, locationId, optional)

List insight/journal

List journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***InsightProjectJournalListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InsightProjectJournalListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalLogGet

> InsightProjectJournalLogGet(ctx, projectId, locationId, journalId, optional)

Get insight/journal.log

websocket is also supported

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
 **optional** | ***InsightProjectJournalLogGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InsightProjectJournalLogGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **since** | **optional.Time**| since | 
 **until** | **optional.Time**| until | 
 **follow** | **optional.Bool**| follow | [default to false]
 **tail** | **optional.Float32**| tail | 
 **tag** | [**optional.Interface of []Tag**](Tag.md)| tag | 

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


## InsightProjectJournalServiceGet

> ResourceService InsightProjectJournalServiceGet(ctx, projectId, locationId, journalId, serviceId)

Get insight/journal.service

Get insight/journal.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
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


## InsightProjectJournalServiceList

> []ResourceService InsightProjectJournalServiceList(ctx, projectId, locationId, journalId)

List insight/journal.service

List insight/journal.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 

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


## InsightProjectJournalTagCreate

> Tag InsightProjectJournalTagCreate(ctx, projectId, locationId, journalId, tag)

Create insight/journal.tag

Create insight/journal.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
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


## InsightProjectJournalTagDelete

> InsightProjectJournalTagDelete(ctx, projectId, locationId, journalId, tagId)

Delete insight/journal.tag

Delete insight/journal.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
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


## InsightProjectJournalTagGet

> Tag InsightProjectJournalTagGet(ctx, projectId, locationId, journalId, tagId)

Get insight/journal.tag

Get insight/journal.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
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


## InsightProjectJournalTagList

> []Tag InsightProjectJournalTagList(ctx, projectId, locationId, journalId)

List insight/journal.tag

List insight/journal.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 

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


## InsightProjectJournalTagPut

> []Tag InsightProjectJournalTagPut(ctx, projectId, locationId, journalId, tag)

Replace insight/journal.tag

Replace insight/journal.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
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


## InsightProjectJournalTransfer

> Journal InsightProjectJournalTransfer(ctx, projectId, locationId, journalId, insightProjectJournalTransfer, optional)

Transfer insight/journal

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
**insightProjectJournalTransfer** | [**InsightProjectJournalTransfer**](InsightProjectJournalTransfer.md)|  | 
 **optional** | ***InsightProjectJournalTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InsightProjectJournalTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalUpdate

> Journal InsightProjectJournalUpdate(ctx, projectId, locationId, journalId, insightProjectJournalUpdate)

Update insight/journal

Returns modified journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**journalId** | **string**| Journal Id | 
**insightProjectJournalUpdate** | [**InsightProjectJournalUpdate**](InsightProjectJournalUpdate.md)|  | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

