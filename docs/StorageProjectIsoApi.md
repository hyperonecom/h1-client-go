# \StorageProjectIsoApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProjectIsoCreate**](StorageProjectIsoApi.md#StorageProjectIsoCreate) | **Post** /storage/{locationId}/project/{projectId}/iso | Create storage/iso
[**StorageProjectIsoDelete**](StorageProjectIsoApi.md#StorageProjectIsoDelete) | **Delete** /storage/{locationId}/project/{projectId}/iso/{isoId} | Delete storage/iso
[**StorageProjectIsoDetach**](StorageProjectIsoApi.md#StorageProjectIsoDetach) | **Post** /storage/{locationId}/project/{projectId}/iso/{isoId}/actions/detach | Detach storage/iso
[**StorageProjectIsoEventGet**](StorageProjectIsoApi.md#StorageProjectIsoEventGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/event/{eventId} | Get storage/iso.event
[**StorageProjectIsoEventList**](StorageProjectIsoApi.md#StorageProjectIsoEventList) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/event | List storage/iso.event
[**StorageProjectIsoGet**](StorageProjectIsoApi.md#StorageProjectIsoGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId} | Get storage/iso
[**StorageProjectIsoList**](StorageProjectIsoApi.md#StorageProjectIsoList) | **Get** /storage/{locationId}/project/{projectId}/iso | List storage/iso
[**StorageProjectIsoServiceGet**](StorageProjectIsoApi.md#StorageProjectIsoServiceGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/service/{serviceId} | Get storage/iso.service
[**StorageProjectIsoServiceList**](StorageProjectIsoApi.md#StorageProjectIsoServiceList) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/service | List storage/iso.service
[**StorageProjectIsoTagCreate**](StorageProjectIsoApi.md#StorageProjectIsoTagCreate) | **Post** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag | Create storage/iso.tag
[**StorageProjectIsoTagDelete**](StorageProjectIsoApi.md#StorageProjectIsoTagDelete) | **Delete** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag/{tagId} | Delete storage/iso.tag
[**StorageProjectIsoTagGet**](StorageProjectIsoApi.md#StorageProjectIsoTagGet) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag/{tagId} | Get storage/iso.tag
[**StorageProjectIsoTagList**](StorageProjectIsoApi.md#StorageProjectIsoTagList) | **Get** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag | List storage/iso.tag
[**StorageProjectIsoTagPut**](StorageProjectIsoApi.md#StorageProjectIsoTagPut) | **Put** /storage/{locationId}/project/{projectId}/iso/{isoId}/tag | Replace storage/iso.tag
[**StorageProjectIsoTransfer**](StorageProjectIsoApi.md#StorageProjectIsoTransfer) | **Post** /storage/{locationId}/project/{projectId}/iso/{isoId}/actions/transfer | Transfer storage/iso
[**StorageProjectIsoUpdate**](StorageProjectIsoApi.md#StorageProjectIsoUpdate) | **Patch** /storage/{locationId}/project/{projectId}/iso/{isoId} | Update storage/iso



## StorageProjectIsoCreate

> Iso StorageProjectIsoCreate(ctx, projectId, locationId).StorageProjectIsoCreate(storageProjectIsoCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create storage/iso



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
    storageProjectIsoCreate := *openapiclient.NewStorageProjectIsoCreate("Name_example", "Source_example") // StorageProjectIsoCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoCreate(context.Background(), projectId, locationId).StorageProjectIsoCreate(storageProjectIsoCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoCreate`: Iso
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storageProjectIsoCreate** | [**StorageProjectIsoCreate**](StorageProjectIsoCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Iso**](Iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoDelete

> InlineResponseDefault StorageProjectIsoDelete(ctx, projectId, locationId, isoId).Execute()

Delete storage/iso



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
    isoId := "isoId_example" // string | Iso Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoDelete(context.Background(), projectId, locationId, isoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoDelete`: InlineResponseDefault
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponseDefault**](InlineResponseDefault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoDetach

> Iso StorageProjectIsoDetach(ctx, projectId, locationId, isoId).StorageProjectIsoDetach(storageProjectIsoDetach).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Detach storage/iso



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
    isoId := "isoId_example" // string | Iso Id
    storageProjectIsoDetach := *openapiclient.NewStorageProjectIsoDetach("Vm_example") // StorageProjectIsoDetach | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoDetach(context.Background(), projectId, locationId, isoId).StorageProjectIsoDetach(storageProjectIsoDetach).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoDetach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoDetach`: Iso
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoDetach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoDetachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectIsoDetach** | [**StorageProjectIsoDetach**](StorageProjectIsoDetach.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Iso**](Iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoEventGet

> Event StorageProjectIsoEventGet(ctx, projectId, locationId, isoId, eventId).Execute()

Get storage/iso.event



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
    isoId := "isoId_example" // string | Iso Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoEventGet(context.Background(), projectId, locationId, isoId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoEventGetRequest struct via the builder pattern


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


## StorageProjectIsoEventList

> []Event StorageProjectIsoEventList(ctx, projectId, locationId, isoId).Limit(limit).Skip(skip).Execute()

List storage/iso.event



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
    isoId := "isoId_example" // string | Iso Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoEventList(context.Background(), projectId, locationId, isoId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoEventListRequest struct via the builder pattern


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


## StorageProjectIsoGet

> Iso StorageProjectIsoGet(ctx, projectId, locationId, isoId).Execute()

Get storage/iso



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
    isoId := "isoId_example" // string | Iso Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoGet(context.Background(), projectId, locationId, isoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoGet`: Iso
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Iso**](Iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoList

> []Iso StorageProjectIsoList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List storage/iso



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoList`: []Iso
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Iso**](Iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoServiceGet

> ResourceService StorageProjectIsoServiceGet(ctx, projectId, locationId, isoId, serviceId).Execute()

Get storage/iso.service



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
    isoId := "isoId_example" // string | Iso Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoServiceGet(context.Background(), projectId, locationId, isoId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoServiceGetRequest struct via the builder pattern


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


## StorageProjectIsoServiceList

> []ResourceService StorageProjectIsoServiceList(ctx, projectId, locationId, isoId).Execute()

List storage/iso.service



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
    isoId := "isoId_example" // string | Iso Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoServiceList(context.Background(), projectId, locationId, isoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoServiceListRequest struct via the builder pattern


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


## StorageProjectIsoTagCreate

> Tag StorageProjectIsoTagCreate(ctx, projectId, locationId, isoId).Tag(tag).Execute()

Create storage/iso.tag



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
    isoId := "isoId_example" // string | Iso Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoTagCreate(context.Background(), projectId, locationId, isoId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoTagCreateRequest struct via the builder pattern


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


## StorageProjectIsoTagDelete

> StorageProjectIsoTagDelete(ctx, projectId, locationId, isoId, tagId).Execute()

Delete storage/iso.tag



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
    isoId := "isoId_example" // string | Iso Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoTagDelete(context.Background(), projectId, locationId, isoId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoTagDelete``: %v\n", err)
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
**isoId** | **string** | Iso Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoTagDeleteRequest struct via the builder pattern


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


## StorageProjectIsoTagGet

> Tag StorageProjectIsoTagGet(ctx, projectId, locationId, isoId, tagId).Execute()

Get storage/iso.tag



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
    isoId := "isoId_example" // string | Iso Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoTagGet(context.Background(), projectId, locationId, isoId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoTagGetRequest struct via the builder pattern


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


## StorageProjectIsoTagList

> []Tag StorageProjectIsoTagList(ctx, projectId, locationId, isoId).Execute()

List storage/iso.tag



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
    isoId := "isoId_example" // string | Iso Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoTagList(context.Background(), projectId, locationId, isoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoTagListRequest struct via the builder pattern


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


## StorageProjectIsoTagPut

> []Tag StorageProjectIsoTagPut(ctx, projectId, locationId, isoId).Tag(tag).Execute()

Replace storage/iso.tag



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
    isoId := "isoId_example" // string | Iso Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoTagPut(context.Background(), projectId, locationId, isoId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoTagPutRequest struct via the builder pattern


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


## StorageProjectIsoTransfer

> Iso StorageProjectIsoTransfer(ctx, projectId, locationId, isoId).StorageProjectIsoTransfer(storageProjectIsoTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer storage/iso



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
    isoId := "isoId_example" // string | Iso Id
    storageProjectIsoTransfer := *openapiclient.NewStorageProjectIsoTransfer("Project_example") // StorageProjectIsoTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoTransfer(context.Background(), projectId, locationId, isoId).StorageProjectIsoTransfer(storageProjectIsoTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoTransfer`: Iso
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectIsoTransfer** | [**StorageProjectIsoTransfer**](StorageProjectIsoTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Iso**](Iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageProjectIsoUpdate

> Iso StorageProjectIsoUpdate(ctx, projectId, locationId, isoId).StorageProjectIsoUpdate(storageProjectIsoUpdate).Execute()

Update storage/iso



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
    isoId := "isoId_example" // string | Iso Id
    storageProjectIsoUpdate := *openapiclient.NewStorageProjectIsoUpdate() // StorageProjectIsoUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageProjectIsoApi.StorageProjectIsoUpdate(context.Background(), projectId, locationId, isoId).StorageProjectIsoUpdate(storageProjectIsoUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageProjectIsoApi.StorageProjectIsoUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageProjectIsoUpdate`: Iso
    fmt.Fprintf(os.Stdout, "Response from `StorageProjectIsoApi.StorageProjectIsoUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**isoId** | **string** | Iso Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageProjectIsoUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageProjectIsoUpdate** | [**StorageProjectIsoUpdate**](StorageProjectIsoUpdate.md) |  | 

### Return type

[**Iso**](Iso.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

