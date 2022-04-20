# \StorageProjectDiskApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectDiskCreate**](StorageProjectDiskApi.md#StorageProjectDiskCreate) | **Post** /storage/{locationId}/project/{projectId}/disk | Create storage/disk
[**StorageProjectDiskDelete**](StorageProjectDiskApi.md#StorageProjectDiskDelete) | **Delete** /storage/{locationId}/project/{projectId}/disk/{diskId} | Delete storage/disk
[**StorageProjectDiskDetach**](StorageProjectDiskApi.md#StorageProjectDiskDetach) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/detach | Detach storage/disk
[**StorageProjectDiskDownload**](StorageProjectDiskApi.md#StorageProjectDiskDownload) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/download | Download storage/disk
[**StorageProjectDiskEventGet**](StorageProjectDiskApi.md#StorageProjectDiskEventGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/event/{eventId} | Get storage/disk.event
[**StorageProjectDiskEventList**](StorageProjectDiskApi.md#StorageProjectDiskEventList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/event | List storage/disk.event
[**StorageProjectDiskGet**](StorageProjectDiskApi.md#StorageProjectDiskGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId} | Get storage/disk
[**StorageProjectDiskList**](StorageProjectDiskApi.md#StorageProjectDiskList) | **Get** /storage/{locationId}/project/{projectId}/disk | List storage/disk
[**StorageProjectDiskMetricGet**](StorageProjectDiskApi.md#StorageProjectDiskMetricGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/metric/{metricId} | Get storage/disk.metric
[**StorageProjectDiskMetricList**](StorageProjectDiskApi.md#StorageProjectDiskMetricList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/metric | List storage/disk.metric
[**StorageProjectDiskMetricSeriesList**](StorageProjectDiskApi.md#StorageProjectDiskMetricSeriesList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/metric/{metricId}/series | List storage/disk.series
[**StorageProjectDiskResize**](StorageProjectDiskApi.md#StorageProjectDiskResize) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/resize | Resize storage/disk
[**StorageProjectDiskServiceGet**](StorageProjectDiskApi.md#StorageProjectDiskServiceGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/service/{serviceId} | Get storage/disk.service
[**StorageProjectDiskServiceList**](StorageProjectDiskApi.md#StorageProjectDiskServiceList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/service | List storage/disk.service
[**StorageProjectDiskTagCreate**](StorageProjectDiskApi.md#StorageProjectDiskTagCreate) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag | Create storage/disk.tag
[**StorageProjectDiskTagDelete**](StorageProjectDiskApi.md#StorageProjectDiskTagDelete) | **Delete** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag/{tagId} | Delete storage/disk.tag
[**StorageProjectDiskTagGet**](StorageProjectDiskApi.md#StorageProjectDiskTagGet) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag/{tagId} | Get storage/disk.tag
[**StorageProjectDiskTagList**](StorageProjectDiskApi.md#StorageProjectDiskTagList) | **Get** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag | List storage/disk.tag
[**StorageProjectDiskTagPut**](StorageProjectDiskApi.md#StorageProjectDiskTagPut) | **Put** /storage/{locationId}/project/{projectId}/disk/{diskId}/tag | Replace storage/disk.tag
[**StorageProjectDiskTransfer**](StorageProjectDiskApi.md#StorageProjectDiskTransfer) | **Post** /storage/{locationId}/project/{projectId}/disk/{diskId}/actions/transfer | Transfer storage/disk
[**StorageProjectDiskUpdate**](StorageProjectDiskApi.md#StorageProjectDiskUpdate) | **Patch** /storage/{locationId}/project/{projectId}/disk/{diskId} | Update storage/disk



## StorageProjectDiskCreate

> Disk StorageProjectDiskCreate(ctx, projectId, locationId).StorageProjectDiskCreate(storageProjectDiskCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create storage/disk



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
    storageProjectDiskCreate := *openapiclient.NewStorageProjectDiskCreate("Name_example", "Service_example", float32(123)) // StorageProjectDiskCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskCreate(context.Background(), projectId, locationId).StorageProjectDiskCreate(storageProjectDiskCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskCreate`: Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storageProjectDiskCreate** | [**StorageProjectDiskCreate**](StorageProjectDiskCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskDelete

> Disk StorageProjectDiskDelete(ctx, projectId, locationId, diskId).Execute()

Delete storage/disk



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
    diskId := "diskId_example" // string | Disk Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskDelete(context.Background(), projectId, locationId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskDelete`: Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskDetach

> Disk StorageProjectDiskDetach(ctx, projectId, locationId, diskId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Detach storage/disk



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
    diskId := "diskId_example" // string | Disk Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskDetach(context.Background(), projectId, locationId, diskId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskDetach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskDetach`: Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskDetach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskDetachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskDownload

> StorageProjectDiskDownload(ctx, projectId, locationId, diskId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Download storage/disk



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
    diskId := "diskId_example" // string | Disk Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskDownload(context.Background(), projectId, locationId, diskId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskDownload``: %v\n", err)
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
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

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


## StorageProjectDiskEventGet

> Event StorageProjectDiskEventGet(ctx, projectId, locationId, diskId, eventId).Execute()

Get storage/disk.event



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
    diskId := "diskId_example" // string | Disk Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskEventGet(context.Background(), projectId, locationId, diskId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskEventGetRequest struct via the builder pattern


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


## StorageProjectDiskEventList

> []Event StorageProjectDiskEventList(ctx, projectId, locationId, diskId).Limit(limit).Skip(skip).Execute()

List storage/disk.event



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
    diskId := "diskId_example" // string | Disk Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskEventList(context.Background(), projectId, locationId, diskId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskEventListRequest struct via the builder pattern


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


## StorageProjectDiskGet

> Disk StorageProjectDiskGet(ctx, projectId, locationId, diskId).Execute()

Get storage/disk



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
    diskId := "diskId_example" // string | Disk Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskGet(context.Background(), projectId, locationId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskGet`: Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskList

> []Disk StorageProjectDiskList(ctx, projectId, locationId).Name(name).Vm(vm).Image(image).TagValue(tagValue).TagKey(tagKey).Execute()

List storage/disk



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
    vm := "vm_example" // string | Filter by vm (optional)
    image := "image_example" // string | Filter by image (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskList(context.Background(), projectId, locationId).Name(name).Vm(vm).Image(image).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskList`: []Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **vm** | **string** | Filter by vm | 
 **image** | **string** | Filter by image | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskMetricGet

> Metric StorageProjectDiskMetricGet(ctx, projectId, locationId, diskId, metricId).Execute()

Get storage/disk.metric



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
    diskId := "diskId_example" // string | Disk Id
    metricId := "metricId_example" // string | metricId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskMetricGet(context.Background(), projectId, locationId, diskId, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskMetricGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskMetricGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskMetricGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskMetricGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Metric**](Metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskMetricList

> []Metric StorageProjectDiskMetricList(ctx, projectId, locationId, diskId).Execute()

List storage/disk.metric



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
    diskId := "diskId_example" // string | Disk Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskMetricList(context.Background(), projectId, locationId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskMetricList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskMetricList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskMetricList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskMetricListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Metric**](Metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskMetricSeriesList

> []Point StorageProjectDiskMetricSeriesList(ctx, projectId, locationId, diskId, metricId).Interval(interval).Timespan(timespan).Aligner(aligner).Execute()

List storage/disk.series



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
    diskId := "diskId_example" // string | Disk Id
    metricId := "metricId_example" // string | metricId
    interval := "interval_example" // string | interval (optional)
    timespan := "timespan_example" // string | timespan (optional)
    aligner := "aligner_example" // string | aligner (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskMetricSeriesList(context.Background(), projectId, locationId, diskId, metricId).Interval(interval).Timespan(timespan).Aligner(aligner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskMetricSeriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskMetricSeriesList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskMetricSeriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskMetricSeriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **interval** | **string** | interval | 
 **timespan** | **string** | timespan | 
 **aligner** | **string** | aligner | 

### Return type

[**[]Point**](Point.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskResize

> Disk StorageProjectDiskResize(ctx, projectId, locationId, diskId).StorageProjectDiskResize(storageProjectDiskResize).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Resize storage/disk



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
    diskId := "diskId_example" // string | Disk Id
    storageProjectDiskResize := *openapiclient.NewStorageProjectDiskResize(float32(123)) // StorageProjectDiskResize | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskResize(context.Background(), projectId, locationId, diskId).StorageProjectDiskResize(storageProjectDiskResize).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskResize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskResize`: Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskResize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskResizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectDiskResize** | [**StorageProjectDiskResize**](StorageProjectDiskResize.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskServiceGet

> ResourceService StorageProjectDiskServiceGet(ctx, projectId, locationId, diskId, serviceId).Execute()

Get storage/disk.service



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
    diskId := "diskId_example" // string | Disk Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskServiceGet(context.Background(), projectId, locationId, diskId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskServiceGetRequest struct via the builder pattern


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


## StorageProjectDiskServiceList

> []ResourceService StorageProjectDiskServiceList(ctx, projectId, locationId, diskId).Execute()

List storage/disk.service



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
    diskId := "diskId_example" // string | Disk Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskServiceList(context.Background(), projectId, locationId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskServiceListRequest struct via the builder pattern


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


## StorageProjectDiskTagCreate

> Tag StorageProjectDiskTagCreate(ctx, projectId, locationId, diskId).Tag(tag).Execute()

Create storage/disk.tag



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
    diskId := "diskId_example" // string | Disk Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskTagCreate(context.Background(), projectId, locationId, diskId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskTagCreateRequest struct via the builder pattern


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


## StorageProjectDiskTagDelete

> StorageProjectDiskTagDelete(ctx, projectId, locationId, diskId, tagId).Execute()

Delete storage/disk.tag



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
    diskId := "diskId_example" // string | Disk Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskTagDelete(context.Background(), projectId, locationId, diskId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskTagDelete``: %v\n", err)
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
**diskId** | **string** | Disk Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskTagDeleteRequest struct via the builder pattern


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


## StorageProjectDiskTagGet

> Tag StorageProjectDiskTagGet(ctx, projectId, locationId, diskId, tagId).Execute()

Get storage/disk.tag



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
    diskId := "diskId_example" // string | Disk Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskTagGet(context.Background(), projectId, locationId, diskId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskTagGetRequest struct via the builder pattern


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


## StorageProjectDiskTagList

> []Tag StorageProjectDiskTagList(ctx, projectId, locationId, diskId).Execute()

List storage/disk.tag



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
    diskId := "diskId_example" // string | Disk Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskTagList(context.Background(), projectId, locationId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskTagListRequest struct via the builder pattern


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


## StorageProjectDiskTagPut

> []Tag StorageProjectDiskTagPut(ctx, projectId, locationId, diskId).Tag(tag).Execute()

Replace storage/disk.tag



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
    diskId := "diskId_example" // string | Disk Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskTagPut(context.Background(), projectId, locationId, diskId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskTagPutRequest struct via the builder pattern


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


## StorageProjectDiskTransfer

> Disk StorageProjectDiskTransfer(ctx, projectId, locationId, diskId).StorageProjectDiskTransfer(storageProjectDiskTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer storage/disk



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
    diskId := "diskId_example" // string | Disk Id
    storageProjectDiskTransfer := *openapiclient.NewStorageProjectDiskTransfer("Project_example") // StorageProjectDiskTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskTransfer(context.Background(), projectId, locationId, diskId).StorageProjectDiskTransfer(storageProjectDiskTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskTransfer`: Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectDiskTransfer** | [**StorageProjectDiskTransfer**](StorageProjectDiskTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectDiskUpdate

> Disk StorageProjectDiskUpdate(ctx, projectId, locationId, diskId).StorageProjectDiskUpdate(storageProjectDiskUpdate).Execute()

Update storage/disk



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
    diskId := "diskId_example" // string | Disk Id
    storageProjectDiskUpdate := *openapiclient.NewStorageProjectDiskUpdate() // StorageProjectDiskUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectDiskApi.StorageProjectDiskUpdate(context.Background(), projectId, locationId, diskId).StorageProjectDiskUpdate(storageProjectDiskUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectDiskApi.StorageProjectDiskUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectDiskUpdate`: Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectDiskApi.StorageProjectDiskUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**diskId** | **string** | Disk Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectDiskUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectDiskUpdate** | [**StorageProjectDiskUpdate**](StorageProjectDiskUpdate.md) |  | 

### Return type

[**Disk**](Disk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

