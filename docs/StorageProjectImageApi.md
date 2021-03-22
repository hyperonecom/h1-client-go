# \StorageProjectImageApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectImageCreate**](StorageProjectImageApi.md#StorageProjectImageCreate) | **Post** /storage/{locationId}/project/{projectId}/image | Create storage/image
[**StorageProjectImageDelete**](StorageProjectImageApi.md#StorageProjectImageDelete) | **Delete** /storage/{locationId}/project/{projectId}/image/{imageId} | Delete storage/image
[**StorageProjectImageDiskList**](StorageProjectImageApi.md#StorageProjectImageDiskList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/disk | List storage/image.disk
[**StorageProjectImageEventGet**](StorageProjectImageApi.md#StorageProjectImageEventGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/event/{eventId} | Get storage/image.event
[**StorageProjectImageEventList**](StorageProjectImageApi.md#StorageProjectImageEventList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/event | List storage/image.event
[**StorageProjectImageGet**](StorageProjectImageApi.md#StorageProjectImageGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId} | Get storage/image
[**StorageProjectImageList**](StorageProjectImageApi.md#StorageProjectImageList) | **Get** /storage/{locationId}/project/{projectId}/image | List storage/image
[**StorageProjectImageServiceGet**](StorageProjectImageApi.md#StorageProjectImageServiceGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/service/{serviceId} | Get storage/image.service
[**StorageProjectImageServiceList**](StorageProjectImageApi.md#StorageProjectImageServiceList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/service | List storage/image.service
[**StorageProjectImageTagCreate**](StorageProjectImageApi.md#StorageProjectImageTagCreate) | **Post** /storage/{locationId}/project/{projectId}/image/{imageId}/tag | Create storage/image.tag
[**StorageProjectImageTagDelete**](StorageProjectImageApi.md#StorageProjectImageTagDelete) | **Delete** /storage/{locationId}/project/{projectId}/image/{imageId}/tag/{tagId} | Delete storage/image.tag
[**StorageProjectImageTagGet**](StorageProjectImageApi.md#StorageProjectImageTagGet) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/tag/{tagId} | Get storage/image.tag
[**StorageProjectImageTagList**](StorageProjectImageApi.md#StorageProjectImageTagList) | **Get** /storage/{locationId}/project/{projectId}/image/{imageId}/tag | List storage/image.tag
[**StorageProjectImageTagPut**](StorageProjectImageApi.md#StorageProjectImageTagPut) | **Put** /storage/{locationId}/project/{projectId}/image/{imageId}/tag | Replace storage/image.tag
[**StorageProjectImageTransfer**](StorageProjectImageApi.md#StorageProjectImageTransfer) | **Post** /storage/{locationId}/project/{projectId}/image/{imageId}/actions/transfer | Transfer storage/image
[**StorageProjectImageUpdate**](StorageProjectImageApi.md#StorageProjectImageUpdate) | **Patch** /storage/{locationId}/project/{projectId}/image/{imageId} | Update storage/image



## StorageProjectImageCreate

> Image StorageProjectImageCreate(ctx, projectId, locationId).StorageProjectImageCreate(storageProjectImageCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create storage/image



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
    storageProjectImageCreate := *openapiclient.NewStorageProjectImageCreate("Name_example") // StorageProjectImageCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageCreate(context.Background(), projectId, locationId).StorageProjectImageCreate(storageProjectImageCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageCreate`: Image
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storageProjectImageCreate** | [**StorageProjectImageCreate**](StorageProjectImageCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageDelete

> StorageProjectImageDelete(ctx, projectId, locationId, imageId).Execute()

Delete storage/image



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
    imageId := "imageId_example" // string | Image Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageDelete(context.Background(), projectId, locationId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageDelete``: %v\n", err)
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
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageDeleteRequest struct via the builder pattern


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


## StorageProjectImageDiskList

> []Disk StorageProjectImageDiskList(ctx, projectId, locationId, imageId).Execute()

List storage/image.disk



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
    imageId := "imageId_example" // string | Image Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageDiskList(context.Background(), projectId, locationId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageDiskList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageDiskList`: []Disk
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageDiskList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageDiskListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## StorageProjectImageEventGet

> Event StorageProjectImageEventGet(ctx, projectId, locationId, imageId, eventId).Execute()

Get storage/image.event



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
    imageId := "imageId_example" // string | Image Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageEventGet(context.Background(), projectId, locationId, imageId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageEventGetRequest struct via the builder pattern


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


## StorageProjectImageEventList

> []Event StorageProjectImageEventList(ctx, projectId, locationId, imageId).Limit(limit).Skip(skip).Execute()

List storage/image.event



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
    imageId := "imageId_example" // string | Image Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageEventList(context.Background(), projectId, locationId, imageId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageEventListRequest struct via the builder pattern


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


## StorageProjectImageGet

> Image StorageProjectImageGet(ctx, projectId, locationId, imageId).Execute()

Get storage/image



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
    imageId := "imageId_example" // string | Image Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageGet(context.Background(), projectId, locationId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageGet`: Image
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageList

> []Image StorageProjectImageList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List storage/image



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
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageList`: []Image
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageServiceGet

> ResourceService StorageProjectImageServiceGet(ctx, projectId, locationId, imageId, serviceId).Execute()

Get storage/image.service



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
    imageId := "imageId_example" // string | Image Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageServiceGet(context.Background(), projectId, locationId, imageId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageServiceGetRequest struct via the builder pattern


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


## StorageProjectImageServiceList

> []ResourceService StorageProjectImageServiceList(ctx, projectId, locationId, imageId).Execute()

List storage/image.service



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
    imageId := "imageId_example" // string | Image Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageServiceList(context.Background(), projectId, locationId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageServiceListRequest struct via the builder pattern


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


## StorageProjectImageTagCreate

> Tag StorageProjectImageTagCreate(ctx, projectId, locationId, imageId).Tag(tag).Execute()

Create storage/image.tag



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
    imageId := "imageId_example" // string | Image Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageTagCreate(context.Background(), projectId, locationId, imageId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageTagCreateRequest struct via the builder pattern


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


## StorageProjectImageTagDelete

> StorageProjectImageTagDelete(ctx, projectId, locationId, imageId, tagId).Execute()

Delete storage/image.tag



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
    imageId := "imageId_example" // string | Image Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageTagDelete(context.Background(), projectId, locationId, imageId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageTagDelete``: %v\n", err)
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
**imageId** | **string** | Image Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageTagDeleteRequest struct via the builder pattern


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


## StorageProjectImageTagGet

> Tag StorageProjectImageTagGet(ctx, projectId, locationId, imageId, tagId).Execute()

Get storage/image.tag



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
    imageId := "imageId_example" // string | Image Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageTagGet(context.Background(), projectId, locationId, imageId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageTagGetRequest struct via the builder pattern


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


## StorageProjectImageTagList

> []Tag StorageProjectImageTagList(ctx, projectId, locationId, imageId).Execute()

List storage/image.tag



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
    imageId := "imageId_example" // string | Image Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageTagList(context.Background(), projectId, locationId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageTagListRequest struct via the builder pattern


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


## StorageProjectImageTagPut

> []Tag StorageProjectImageTagPut(ctx, projectId, locationId, imageId).Tag(tag).Execute()

Replace storage/image.tag



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
    imageId := "imageId_example" // string | Image Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageTagPut(context.Background(), projectId, locationId, imageId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageTagPutRequest struct via the builder pattern


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


## StorageProjectImageTransfer

> Image StorageProjectImageTransfer(ctx, projectId, locationId, imageId).StorageProjectImageTransfer(storageProjectImageTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer storage/image



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
    imageId := "imageId_example" // string | Image Id
    storageProjectImageTransfer := *openapiclient.NewStorageProjectImageTransfer("Project_example") // StorageProjectImageTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageTransfer(context.Background(), projectId, locationId, imageId).StorageProjectImageTransfer(storageProjectImageTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageTransfer`: Image
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectImageTransfer** | [**StorageProjectImageTransfer**](StorageProjectImageTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectImageUpdate

> Image StorageProjectImageUpdate(ctx, projectId, locationId, imageId).StorageProjectImageUpdate(storageProjectImageUpdate).Execute()

Update storage/image



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
    imageId := "imageId_example" // string | Image Id
    storageProjectImageUpdate := *openapiclient.NewStorageProjectImageUpdate() // StorageProjectImageUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StorageProjectImageApi.StorageProjectImageUpdate(context.Background(), projectId, locationId, imageId).StorageProjectImageUpdate(storageProjectImageUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectImageApi.StorageProjectImageUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectImageUpdate`: Image
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectImageApi.StorageProjectImageUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**imageId** | **string** | Image Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectImageUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectImageUpdate** | [**StorageProjectImageUpdate**](StorageProjectImageUpdate.md) |  | 

### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

