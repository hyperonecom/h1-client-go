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

> ResourceConnect StorageProjectVaultConnectGet(ctx, projectId, locationId, vaultId, connectId).Execute()

Get storage/vault.connect



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    connectId := "connectId_example" // string | connectId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultConnectGet(context.Background(), projectId, locationId, vaultId, connectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultConnectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultConnectGet`: ResourceConnect
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultConnectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**connectId** | **string** | connectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultConnectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ResourceConnect**](ResourceConnect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultConnectList

> []ResourceConnect StorageProjectVaultConnectList(ctx, projectId, locationId, vaultId).Execute()

List storage/vault.connect



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultConnectList(context.Background(), projectId, locationId, vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultConnectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultConnectList`: []ResourceConnect
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultConnectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultConnectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ResourceConnect**](ResourceConnect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCreate

> Vault StorageProjectVaultCreate(ctx, projectId, locationId).StorageProjectVaultCreate(storageProjectVaultCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    storageProjectVaultCreate := *openapiclient.NewStorageProjectVaultCreate("Name_example", float32(123)) // StorageProjectVaultCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultCreate(context.Background(), projectId, locationId).StorageProjectVaultCreate(storageProjectVaultCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultCreate`: Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storageProjectVaultCreate** | [**StorageProjectVaultCreate**](StorageProjectVaultCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialCreate

> VaultCredential StorageProjectVaultCredentialCreate(ctx, projectId, locationId, vaultId).VaultCredential(vaultCredential).Execute()

Create storage/vault.credential



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    vaultCredential := *openapiclient.NewVaultCredential("Name_example", "Type_example", "Value_example") // VaultCredential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultCredentialCreate(context.Background(), projectId, locationId, vaultId).VaultCredential(vaultCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultCredentialCreate`: VaultCredential
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vaultCredential** | [**VaultCredential**](VaultCredential.md) |  | 

### Return type

[**VaultCredential**](VaultCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialDelete

> Vault StorageProjectVaultCredentialDelete(ctx, projectId, locationId, vaultId, credentialId).Execute()

Delete storage/vault.credential



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultCredentialDelete(context.Background(), projectId, locationId, vaultId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultCredentialDelete`: Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialGet

> VaultCredential StorageProjectVaultCredentialGet(ctx, projectId, locationId, vaultId, credentialId).Execute()

Get storage/vault.credential



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultCredentialGet(context.Background(), projectId, locationId, vaultId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultCredentialGet`: VaultCredential
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**VaultCredential**](VaultCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialList

> []VaultCredential StorageProjectVaultCredentialList(ctx, projectId, locationId, vaultId).Execute()

List storage/vault.credential



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultCredentialList(context.Background(), projectId, locationId, vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultCredentialList`: []VaultCredential
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]VaultCredential**](VaultCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultCredentialPatch

> VaultCredential StorageProjectVaultCredentialPatch(ctx, projectId, locationId, vaultId, credentialId).StorageProjectVaultCredentialPatch(storageProjectVaultCredentialPatch).Execute()

Update storage/vault.credential



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    credentialId := "credentialId_example" // string | credentialId
    storageProjectVaultCredentialPatch := *openapiclient.NewStorageProjectVaultCredentialPatch("Name_example") // StorageProjectVaultCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultCredentialPatch(context.Background(), projectId, locationId, vaultId, credentialId).StorageProjectVaultCredentialPatch(storageProjectVaultCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultCredentialPatch`: VaultCredential
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **storageProjectVaultCredentialPatch** | [**StorageProjectVaultCredentialPatch**](StorageProjectVaultCredentialPatch.md) |  | 

### Return type

[**VaultCredential**](VaultCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultDelete

> StorageProjectVaultDelete(ctx, projectId, locationId, vaultId).StorageProjectVaultDelete(storageProjectVaultDelete).Execute()

Delete storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    storageProjectVaultDelete := *openapiclient.NewStorageProjectVaultDelete() // StorageProjectVaultDelete | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultDelete(context.Background(), projectId, locationId, vaultId).StorageProjectVaultDelete(storageProjectVaultDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectVaultDelete** | [**StorageProjectVaultDelete**](StorageProjectVaultDelete.md) |  | 

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

> Event StorageProjectVaultEventGet(ctx, projectId, locationId, vaultId, eventId).Execute()

Get storage/vault.event



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultEventGet(context.Background(), projectId, locationId, vaultId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultEventGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Event**](Event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultEventList

> []Event StorageProjectVaultEventList(ctx, projectId, locationId, vaultId).Limit(limit).Skip(skip).Execute()

List storage/vault.event



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultEventList(context.Background(), projectId, locationId, vaultId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **float32** | $limit | [default to 100]
 **skip** | **float32** | $skip | 

### Return type

[**[]Event**](Event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultGet

> Vault StorageProjectVaultGet(ctx, projectId, locationId, vaultId).Execute()

Get storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultGet(context.Background(), projectId, locationId, vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultGet`: Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultList

> []Vault StorageProjectVaultList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    name := "name_example" // string | Filter by name (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultList`: []Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultResize

> Vault StorageProjectVaultResize(ctx, projectId, locationId, vaultId).StorageProjectVaultResize(storageProjectVaultResize).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Resize storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    storageProjectVaultResize := *openapiclient.NewStorageProjectVaultResize(float32(123)) // StorageProjectVaultResize | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultResize(context.Background(), projectId, locationId, vaultId).StorageProjectVaultResize(storageProjectVaultResize).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultResize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultResize`: Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultResize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultResizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectVaultResize** | [**StorageProjectVaultResize**](StorageProjectVaultResize.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultServiceGet

> ResourceService StorageProjectVaultServiceGet(ctx, projectId, locationId, vaultId, serviceId).Execute()

Get storage/vault.service



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultServiceGet(context.Background(), projectId, locationId, vaultId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ResourceService**](ResourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultServiceList

> []ResourceService StorageProjectVaultServiceList(ctx, projectId, locationId, vaultId).Execute()

List storage/vault.service



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultServiceList(context.Background(), projectId, locationId, vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ResourceService**](ResourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultSnapshotCreate

> StorageSnapshot StorageProjectVaultSnapshotCreate(ctx, projectId, locationId, vaultId).StorageProjectVaultSnapshotCreate(storageProjectVaultSnapshotCreate).Execute()

Create storage/vault.snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    storageProjectVaultSnapshotCreate := *openapiclient.NewStorageProjectVaultSnapshotCreate("Name_example") // StorageProjectVaultSnapshotCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultSnapshotCreate(context.Background(), projectId, locationId, vaultId).StorageProjectVaultSnapshotCreate(storageProjectVaultSnapshotCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultSnapshotCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultSnapshotCreate`: StorageSnapshot
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultSnapshotCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultSnapshotCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectVaultSnapshotCreate** | [**StorageProjectVaultSnapshotCreate**](StorageProjectVaultSnapshotCreate.md) |  | 

### Return type

[**StorageSnapshot**](StorageSnapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultSnapshotDelete

> StorageSnapshot StorageProjectVaultSnapshotDelete(ctx, projectId, locationId, vaultId, snapshotId).Execute()

Delete storage/vault.snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    snapshotId := "snapshotId_example" // string | snapshotId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultSnapshotDelete(context.Background(), projectId, locationId, vaultId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultSnapshotDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultSnapshotDelete`: StorageSnapshot
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultSnapshotDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**snapshotId** | **string** | snapshotId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultSnapshotDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StorageSnapshot**](StorageSnapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultSnapshotGet

> StorageSnapshot StorageProjectVaultSnapshotGet(ctx, projectId, locationId, vaultId, snapshotId).Execute()

Get storage/vault.snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    snapshotId := "snapshotId_example" // string | snapshotId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultSnapshotGet(context.Background(), projectId, locationId, vaultId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultSnapshotGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultSnapshotGet`: StorageSnapshot
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultSnapshotGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**snapshotId** | **string** | snapshotId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultSnapshotGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StorageSnapshot**](StorageSnapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultSnapshotList

> []StorageSnapshot StorageProjectVaultSnapshotList(ctx, projectId, locationId, vaultId).Execute()

List storage/vault.snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultSnapshotList(context.Background(), projectId, locationId, vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultSnapshotList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultSnapshotList`: []StorageSnapshot
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultSnapshotList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultSnapshotListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]StorageSnapshot**](StorageSnapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultStart

> Vault StorageProjectVaultStart(ctx, projectId, locationId, vaultId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Start storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultStart(context.Background(), projectId, locationId, vaultId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultStart`: Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultStop

> Vault StorageProjectVaultStop(ctx, projectId, locationId, vaultId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Stop storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultStop(context.Background(), projectId, locationId, vaultId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultStop`: Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultTagCreate

> Tag StorageProjectVaultTagCreate(ctx, projectId, locationId, vaultId).Tag(tag).Execute()

Create storage/vault.tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultTagCreate(context.Background(), projectId, locationId, vaultId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultTagCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | [**Tag**](Tag.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultTagDelete

> StorageProjectVaultTagDelete(ctx, projectId, locationId, vaultId, tagId).Execute()

Delete storage/vault.tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultTagDelete(context.Background(), projectId, locationId, vaultId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultTagDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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

> Tag StorageProjectVaultTagGet(ctx, projectId, locationId, vaultId, tagId).Execute()

Get storage/vault.tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultTagGet(context.Background(), projectId, locationId, vaultId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultTagGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Tag**](Tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultTagList

> []Tag StorageProjectVaultTagList(ctx, projectId, locationId, vaultId).Execute()

List storage/vault.tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultTagList(context.Background(), projectId, locationId, vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultTagListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Tag**](Tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultTagPut

> []Tag StorageProjectVaultTagPut(ctx, projectId, locationId, vaultId).Tag(tag).Execute()

Replace storage/vault.tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultTagPut(context.Background(), projectId, locationId, vaultId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultTagPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | [**[]Tag**](Tag.md) |  | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectVaultUpdate

> Vault StorageProjectVaultUpdate(ctx, projectId, locationId, vaultId).StorageProjectVaultUpdate(storageProjectVaultUpdate).Execute()

Update storage/vault



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    locationId := "locationId_example" // string | Location Id
    vaultId := "vaultId_example" // string | Vault Id
    storageProjectVaultUpdate := *openapiclient.NewStorageProjectVaultUpdate() // StorageProjectVaultUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectVaultApi.StorageProjectVaultUpdate(context.Background(), projectId, locationId, vaultId).StorageProjectVaultUpdate(storageProjectVaultUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectVaultApi.StorageProjectVaultUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectVaultUpdate`: Vault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectVaultApi.StorageProjectVaultUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vaultId** | **string** | Vault Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectVaultUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectVaultUpdate** | [**StorageProjectVaultUpdate**](StorageProjectVaultUpdate.md) |  | 

### Return type

[**Vault**](Vault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

