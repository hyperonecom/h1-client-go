# \StorageProjectBucketApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectBucketGet**](StorageProjectBucketApi.md#StorageProjectBucketGet) | **Get** /storage/{locationId}/project/{projectId}/bucket/{bucketId} | Get storage/bucket
[**StorageProjectBucketList**](StorageProjectBucketApi.md#StorageProjectBucketList) | **Get** /storage/{locationId}/project/{projectId}/bucket | List storage/bucket
[**StorageProjectBucketObjectDelete**](StorageProjectBucketApi.md#StorageProjectBucketObjectDelete) | **Delete** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object/{objectId} | Delete storage/bucket.object
[**StorageProjectBucketObjectDownload**](StorageProjectBucketApi.md#StorageProjectBucketObjectDownload) | **Post** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object/{objectId}/actions/download | Download storage/bucket.object
[**StorageProjectBucketObjectGet**](StorageProjectBucketApi.md#StorageProjectBucketObjectGet) | **Get** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object/{objectId} | Get storage/bucket.object
[**StorageProjectBucketObjectList**](StorageProjectBucketApi.md#StorageProjectBucketObjectList) | **Get** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/object | List storage/bucket.object
[**StorageProjectBucketUpload**](StorageProjectBucketApi.md#StorageProjectBucketUpload) | **Post** /storage/{locationId}/project/{projectId}/bucket/{bucketId}/actions/upload | Upload storage/bucket



## StorageProjectBucketGet

> Bucket StorageProjectBucketGet(ctx, projectId, locationId, bucketId).Execute()

Get storage/bucket



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
    bucketId := "bucketId_example" // string | Bucket Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectBucketApi.StorageProjectBucketGet(context.Background(), projectId, locationId, bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectBucketApi.StorageProjectBucketGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectBucketGet`: Bucket
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectBucketApi.StorageProjectBucketGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**bucketId** | **string** | Bucket Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectBucketGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Bucket**](Bucket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketList

> []Bucket StorageProjectBucketList(ctx, projectId, locationId).Execute()

List storage/bucket



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectBucketApi.StorageProjectBucketList(context.Background(), projectId, locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectBucketApi.StorageProjectBucketList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectBucketList`: []Bucket
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectBucketApi.StorageProjectBucketList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectBucketListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Bucket**](Bucket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketObjectDelete

> StorageProjectBucketObjectDelete(ctx, projectId, locationId, bucketId, objectId).Execute()

Delete storage/bucket.object



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
    bucketId := "bucketId_example" // string | Bucket Id
    objectId := "objectId_example" // string | objectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectBucketApi.StorageProjectBucketObjectDelete(context.Background(), projectId, locationId, bucketId, objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectBucketApi.StorageProjectBucketObjectDelete``: %v\n", err)
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
**bucketId** | **string** | Bucket Id | 
**objectId** | **string** | objectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectBucketObjectDeleteRequest struct via the builder pattern


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


## StorageProjectBucketObjectDownload

> StorageProjectBucketObjectDownload(ctx, projectId, locationId, bucketId, objectId).Execute()

Download storage/bucket.object



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
    bucketId := "bucketId_example" // string | Bucket Id
    objectId := "objectId_example" // string | objectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectBucketApi.StorageProjectBucketObjectDownload(context.Background(), projectId, locationId, bucketId, objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectBucketApi.StorageProjectBucketObjectDownload``: %v\n", err)
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
**bucketId** | **string** | Bucket Id | 
**objectId** | **string** | objectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectBucketObjectDownloadRequest struct via the builder pattern


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


## StorageProjectBucketObjectGet

> StorageObject StorageProjectBucketObjectGet(ctx, projectId, locationId, bucketId, objectId).Execute()

Get storage/bucket.object



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
    bucketId := "bucketId_example" // string | Bucket Id
    objectId := "objectId_example" // string | objectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectBucketApi.StorageProjectBucketObjectGet(context.Background(), projectId, locationId, bucketId, objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectBucketApi.StorageProjectBucketObjectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectBucketObjectGet`: StorageObject
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectBucketApi.StorageProjectBucketObjectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**bucketId** | **string** | Bucket Id | 
**objectId** | **string** | objectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectBucketObjectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StorageObject**](StorageObject.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketObjectList

> []StorageObject StorageProjectBucketObjectList(ctx, projectId, locationId, bucketId).Execute()

List storage/bucket.object



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
    bucketId := "bucketId_example" // string | Bucket Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectBucketApi.StorageProjectBucketObjectList(context.Background(), projectId, locationId, bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectBucketApi.StorageProjectBucketObjectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectBucketObjectList`: []StorageObject
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectBucketApi.StorageProjectBucketObjectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**bucketId** | **string** | Bucket Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectBucketObjectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]StorageObject**](StorageObject.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectBucketUpload

> Bucket StorageProjectBucketUpload(ctx, projectId, locationId, bucketId).StorageProjectBucketUpload(storageProjectBucketUpload).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Upload storage/bucket



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
    bucketId := "bucketId_example" // string | Bucket Id
    storageProjectBucketUpload := *openapiclient.NewStorageProjectBucketUpload("Name_example") // StorageProjectBucketUpload | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectBucketApi.StorageProjectBucketUpload(context.Background(), projectId, locationId, bucketId).StorageProjectBucketUpload(storageProjectBucketUpload).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectBucketApi.StorageProjectBucketUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectBucketUpload`: Bucket
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectBucketApi.StorageProjectBucketUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**bucketId** | **string** | Bucket Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectBucketUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectBucketUpload** | [**StorageProjectBucketUpload**](StorageProjectBucketUpload.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

