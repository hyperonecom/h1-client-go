# \StorageProjectVaultApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectVaultConnectGet**](StorageProjectVaultApi.md#StorageProjectVaultConnectGet) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/connect/{connectId} | Get storage/vault.connect
[**StorageProjectVaultConnectList**](StorageProjectVaultApi.md#StorageProjectVaultConnectList) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/connect | List storage/vault.connect
[**StorageProjectVaultCreate**](StorageProjectVaultApi.md#StorageProjectVaultCreate) | **Post** /storage/{locationId}/project/{projectId}/vault | Create storage/vault
[**StorageProjectVaultCredentialCreate**](StorageProjectVaultApi.md#StorageProjectVaultCredentialCreate) | **Post** /storage/{locationId}/project/{projectId}/vault/{vaultId}/credential | Create storage/vault.credential
[**StorageProjectVaultCredentialDelete**](StorageProjectVaultApi.md#StorageProjectVaultCredentialDelete) | **Delete** /storage/{locationId}/project/{projectId}/vault/{vaultId}/credential/{credentialId} | Delete storage/vault.credential
[**StorageProjectVaultCredentialGet**](StorageProjectVaultApi.md#StorageProjectVaultCredentialGet) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/credential/{credentialId} | Get storage/vault.credential
[**StorageProjectVaultCredentialList**](StorageProjectVaultApi.md#StorageProjectVaultCredentialList) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/credential | List storage/vault.credential
[**StorageProjectVaultCredentialPatch**](StorageProjectVaultApi.md#StorageProjectVaultCredentialPatch) | **Patch** /storage/{locationId}/project/{projectId}/vault/{vaultId}/credential/{credentialId} | Update storage/vault.credential
[**StorageProjectVaultDelete**](StorageProjectVaultApi.md#StorageProjectVaultDelete) | **Delete** /storage/{locationId}/project/{projectId}/vault/{vaultId} | Delete storage/vault
[**StorageProjectVaultEventGet**](StorageProjectVaultApi.md#StorageProjectVaultEventGet) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/event/{eventId} | Get storage/vault.event
[**StorageProjectVaultEventList**](StorageProjectVaultApi.md#StorageProjectVaultEventList) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/event | List storage/vault.event
[**StorageProjectVaultGet**](StorageProjectVaultApi.md#StorageProjectVaultGet) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId} | Get storage/vault
[**StorageProjectVaultList**](StorageProjectVaultApi.md#StorageProjectVaultList) | **Get** /storage/{locationId}/project/{projectId}/vault | List storage/vault
[**StorageProjectVaultResize**](StorageProjectVaultApi.md#StorageProjectVaultResize) | **Post** /storage/{locationId}/project/{projectId}/vault/{vaultId}/actions/resize | Resize storage/vault
[**StorageProjectVaultServiceGet**](StorageProjectVaultApi.md#StorageProjectVaultServiceGet) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/service/{serviceId} | Get storage/vault.service
[**StorageProjectVaultServiceList**](StorageProjectVaultApi.md#StorageProjectVaultServiceList) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/service | List storage/vault.service
[**StorageProjectVaultSnapshotCreate**](StorageProjectVaultApi.md#StorageProjectVaultSnapshotCreate) | **Post** /storage/{locationId}/project/{projectId}/vault/{vaultId}/snapshot | Create storage/vault.snapshot
[**StorageProjectVaultSnapshotDelete**](StorageProjectVaultApi.md#StorageProjectVaultSnapshotDelete) | **Delete** /storage/{locationId}/project/{projectId}/vault/{vaultId}/snapshot/{snapshotId} | Delete storage/vault.snapshot
[**StorageProjectVaultSnapshotGet**](StorageProjectVaultApi.md#StorageProjectVaultSnapshotGet) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/snapshot/{snapshotId} | Get storage/vault.snapshot
[**StorageProjectVaultSnapshotList**](StorageProjectVaultApi.md#StorageProjectVaultSnapshotList) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/snapshot | List storage/vault.snapshot
[**StorageProjectVaultStart**](StorageProjectVaultApi.md#StorageProjectVaultStart) | **Post** /storage/{locationId}/project/{projectId}/vault/{vaultId}/actions/start | Start storage/vault
[**StorageProjectVaultStop**](StorageProjectVaultApi.md#StorageProjectVaultStop) | **Post** /storage/{locationId}/project/{projectId}/vault/{vaultId}/actions/stop | Stop storage/vault
[**StorageProjectVaultTagCreate**](StorageProjectVaultApi.md#StorageProjectVaultTagCreate) | **Post** /storage/{locationId}/project/{projectId}/vault/{vaultId}/tag | Create storage/vault.tag
[**StorageProjectVaultTagDelete**](StorageProjectVaultApi.md#StorageProjectVaultTagDelete) | **Delete** /storage/{locationId}/project/{projectId}/vault/{vaultId}/tag/{tagId} | Delete storage/vault.tag
[**StorageProjectVaultTagGet**](StorageProjectVaultApi.md#StorageProjectVaultTagGet) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/tag/{tagId} | Get storage/vault.tag
[**StorageProjectVaultTagList**](StorageProjectVaultApi.md#StorageProjectVaultTagList) | **Get** /storage/{locationId}/project/{projectId}/vault/{vaultId}/tag | List storage/vault.tag
[**StorageProjectVaultTagPut**](StorageProjectVaultApi.md#StorageProjectVaultTagPut) | **Put** /storage/{locationId}/project/{projectId}/vault/{vaultId}/tag | Replace storage/vault.tag
[**StorageProjectVaultUpdate**](StorageProjectVaultApi.md#StorageProjectVaultUpdate) | **Patch** /storage/{locationId}/project/{projectId}/vault/{vaultId} | Update storage/vault



## StorageProjectVaultConnectGet

> ResourceConnect StorageProjectVaultConnectGet(ctx, projectId, locationId, vaultId, connectId)

Get storage/vault.connect

Get storage/vault.connect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
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


## StorageProjectVaultConnectList

> []ResourceConnect StorageProjectVaultConnectList(ctx, projectId, locationId, vaultId)

List storage/vault.connect

List storage/vault.connect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 

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


## StorageProjectVaultCreate

> Vault StorageProjectVaultCreate(ctx, projectId, locationId, storageProjectVaultCreate, optional)

Create storage/vault

Create vault

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**storageProjectVaultCreate** | [**StorageProjectVaultCreate**](StorageProjectVaultCreate.md)|  | 
 **optional** | ***StorageProjectVaultCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectVaultCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialCreate

> VaultCredential StorageProjectVaultCredentialCreate(ctx, projectId, locationId, vaultId, vaultCredential)

Create storage/vault.credential

Create storage/vault.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**vaultCredential** | [**VaultCredential**](VaultCredential.md)|  | 

### Return type

[**VaultCredential**](vault.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialDelete

> Vault StorageProjectVaultCredentialDelete(ctx, projectId, locationId, vaultId, credentialId)

Delete storage/vault.credential

Delete storage/vault.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialGet

> VaultCredential StorageProjectVaultCredentialGet(ctx, projectId, locationId, vaultId, credentialId)

Get storage/vault.credential

Get storage/vault.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**VaultCredential**](vault.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialList

> []VaultCredential StorageProjectVaultCredentialList(ctx, projectId, locationId, vaultId)

List storage/vault.credential

List storage/vault.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 

### Return type

[**[]VaultCredential**](vault.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialPatch

> VaultCredential StorageProjectVaultCredentialPatch(ctx, projectId, locationId, vaultId, credentialId, storageProjectVaultCredentialPatch)

Update storage/vault.credential

Update storage/vault.credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**credentialId** | **string**| credentialId | 
**storageProjectVaultCredentialPatch** | [**StorageProjectVaultCredentialPatch**](StorageProjectVaultCredentialPatch.md)|  | 

### Return type

[**VaultCredential**](vault.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultDelete

> StorageProjectVaultDelete(ctx, projectId, locationId, vaultId, storageProjectVaultDelete)

Delete storage/vault

Delete vault

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**storageProjectVaultDelete** | [**StorageProjectVaultDelete**](StorageProjectVaultDelete.md)|  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultEventGet

> Event StorageProjectVaultEventGet(ctx, projectId, locationId, vaultId, eventId)

Get storage/vault.event

Get storage/vault.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
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


## StorageProjectVaultEventList

> []Event StorageProjectVaultEventList(ctx, projectId, locationId, vaultId, optional)

List storage/vault.event

List storage/vault.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
 **optional** | ***StorageProjectVaultEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectVaultEventListOpts struct


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


## StorageProjectVaultGet

> Vault StorageProjectVaultGet(ctx, projectId, locationId, vaultId)

Get storage/vault

Returns a single vault

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 

### Return type

[**Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultList

> []Vault StorageProjectVaultList(ctx, projectId, locationId, optional)

List storage/vault

List vault

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***StorageProjectVaultListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectVaultListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultResize

> Vault StorageProjectVaultResize(ctx, projectId, locationId, vaultId, storageProjectVaultResize, optional)

Resize storage/vault

action resize

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**storageProjectVaultResize** | [**StorageProjectVaultResize**](StorageProjectVaultResize.md)|  | 
 **optional** | ***StorageProjectVaultResizeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectVaultResizeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultServiceGet

> ResourceService StorageProjectVaultServiceGet(ctx, projectId, locationId, vaultId, serviceId)

Get storage/vault.service

Get storage/vault.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
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


## StorageProjectVaultServiceList

> []ResourceService StorageProjectVaultServiceList(ctx, projectId, locationId, vaultId)

List storage/vault.service

List storage/vault.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 

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


## StorageProjectVaultSnapshotCreate

> StorageSnapshot StorageProjectVaultSnapshotCreate(ctx, projectId, locationId, vaultId, storageProjectVaultSnapshotCreate)

Create storage/vault.snapshot

Create storage/vault.snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**storageProjectVaultSnapshotCreate** | [**StorageProjectVaultSnapshotCreate**](StorageProjectVaultSnapshotCreate.md)|  | 

### Return type

[**StorageSnapshot**](storage.snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultSnapshotDelete

> StorageSnapshot StorageProjectVaultSnapshotDelete(ctx, projectId, locationId, vaultId, snapshotId)

Delete storage/vault.snapshot

Delete storage/vault.snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**snapshotId** | **string**| snapshotId | 

### Return type

[**StorageSnapshot**](storage.snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultSnapshotGet

> StorageSnapshot StorageProjectVaultSnapshotGet(ctx, projectId, locationId, vaultId, snapshotId)

Get storage/vault.snapshot

Get storage/vault.snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**snapshotId** | **string**| snapshotId | 

### Return type

[**StorageSnapshot**](storage.snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultSnapshotList

> []StorageSnapshot StorageProjectVaultSnapshotList(ctx, projectId, locationId, vaultId)

List storage/vault.snapshot

List storage/vault.snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 

### Return type

[**[]StorageSnapshot**](storage.snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultStart

> Vault StorageProjectVaultStart(ctx, projectId, locationId, vaultId, optional)

Start storage/vault

action start

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
 **optional** | ***StorageProjectVaultStartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectVaultStartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultStop

> Vault StorageProjectVaultStop(ctx, projectId, locationId, vaultId, optional)

Stop storage/vault

action stop

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
 **optional** | ***StorageProjectVaultStopOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageProjectVaultStopOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultTagCreate

> Tag StorageProjectVaultTagCreate(ctx, projectId, locationId, vaultId, tag)

Create storage/vault.tag

Create storage/vault.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
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


## StorageProjectVaultTagDelete

> StorageProjectVaultTagDelete(ctx, projectId, locationId, vaultId, tagId)

Delete storage/vault.tag

Delete storage/vault.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
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


## StorageProjectVaultTagGet

> Tag StorageProjectVaultTagGet(ctx, projectId, locationId, vaultId, tagId)

Get storage/vault.tag

Get storage/vault.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
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


## StorageProjectVaultTagList

> []Tag StorageProjectVaultTagList(ctx, projectId, locationId, vaultId)

List storage/vault.tag

List storage/vault.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 

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


## StorageProjectVaultTagPut

> []Tag StorageProjectVaultTagPut(ctx, projectId, locationId, vaultId, tag)

Replace storage/vault.tag

Replace storage/vault.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
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


## StorageProjectVaultUpdate

> Vault StorageProjectVaultUpdate(ctx, projectId, locationId, vaultId, storageProjectVaultUpdate)

Update storage/vault

Returns modified vault

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**vaultId** | **string**| Vault Id | 
**storageProjectVaultUpdate** | [**StorageProjectVaultUpdate**](StorageProjectVaultUpdate.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

