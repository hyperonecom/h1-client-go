# \NetworkingProjectNetworkApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectNetworkCreate**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkCreate) | **Post** /networking/{locationId}/project/{projectId}/network | Create networking/network
[**NetworkingProjectNetworkDelete**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkDelete) | **Delete** /networking/{locationId}/project/{projectId}/network/{networkId} | Delete networking/network
[**NetworkingProjectNetworkEventGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkEventGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/event/{eventId} | Get networking/network.event
[**NetworkingProjectNetworkEventList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkEventList) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/event | List networking/network.event
[**NetworkingProjectNetworkGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId} | Get networking/network
[**NetworkingProjectNetworkList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkList) | **Get** /networking/{locationId}/project/{projectId}/network | List networking/network
[**NetworkingProjectNetworkServiceGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkServiceGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/service/{serviceId} | Get networking/network.service
[**NetworkingProjectNetworkServiceList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkServiceList) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/service | List networking/network.service
[**NetworkingProjectNetworkTagCreate**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagCreate) | **Post** /networking/{locationId}/project/{projectId}/network/{networkId}/tag | Create networking/network.tag
[**NetworkingProjectNetworkTagDelete**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/network/{networkId}/tag/{tagId} | Delete networking/network.tag
[**NetworkingProjectNetworkTagGet**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagGet) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/tag/{tagId} | Get networking/network.tag
[**NetworkingProjectNetworkTagList**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagList) | **Get** /networking/{locationId}/project/{projectId}/network/{networkId}/tag | List networking/network.tag
[**NetworkingProjectNetworkTagPut**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkTagPut) | **Put** /networking/{locationId}/project/{projectId}/network/{networkId}/tag | Replace networking/network.tag
[**NetworkingProjectNetworkUpdate**](NetworkingProjectNetworkApi.md#NetworkingProjectNetworkUpdate) | **Patch** /networking/{locationId}/project/{projectId}/network/{networkId} | Update networking/network



## NetworkingProjectNetworkCreate

> Network NetworkingProjectNetworkCreate(ctx, projectId, locationId).NetworkingProjectNetworkCreate(networkingProjectNetworkCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create networking/network



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
    networkingProjectNetworkCreate := *openapiclient.NewNetworkingProjectNetworkCreate("Name_example") // NetworkingProjectNetworkCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkCreate(context.Background(), projectId, locationId).NetworkingProjectNetworkCreate(networkingProjectNetworkCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkCreate`: Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkingProjectNetworkCreate** | [**NetworkingProjectNetworkCreate**](NetworkingProjectNetworkCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetworkDelete

> InlineResponseDefault NetworkingProjectNetworkDelete(ctx, projectId, locationId, networkId).Execute()

Delete networking/network



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
    networkId := "networkId_example" // string | Network Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkDelete(context.Background(), projectId, locationId, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkDelete`: InlineResponseDefault
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkDeleteRequest struct via the builder pattern


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


## NetworkingProjectNetworkEventGet

> Event NetworkingProjectNetworkEventGet(ctx, projectId, locationId, networkId, eventId).Execute()

Get networking/network.event



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
    networkId := "networkId_example" // string | Network Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkEventGet(context.Background(), projectId, locationId, networkId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkEventGetRequest struct via the builder pattern


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


## NetworkingProjectNetworkEventList

> []Event NetworkingProjectNetworkEventList(ctx, projectId, locationId, networkId).Limit(limit).Skip(skip).Execute()

List networking/network.event



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
    networkId := "networkId_example" // string | Network Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkEventList(context.Background(), projectId, locationId, networkId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkEventListRequest struct via the builder pattern


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


## NetworkingProjectNetworkGet

> Network NetworkingProjectNetworkGet(ctx, projectId, locationId, networkId).Execute()

Get networking/network



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
    networkId := "networkId_example" // string | Network Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkGet(context.Background(), projectId, locationId, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkGet`: Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetworkList

> []Network NetworkingProjectNetworkList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List networking/network



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
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkList`: []Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetworkServiceGet

> ResourceService NetworkingProjectNetworkServiceGet(ctx, projectId, locationId, networkId, serviceId).Execute()

Get networking/network.service



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
    networkId := "networkId_example" // string | Network Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkServiceGet(context.Background(), projectId, locationId, networkId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkServiceGetRequest struct via the builder pattern


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


## NetworkingProjectNetworkServiceList

> []ResourceService NetworkingProjectNetworkServiceList(ctx, projectId, locationId, networkId).Execute()

List networking/network.service



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
    networkId := "networkId_example" // string | Network Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkServiceList(context.Background(), projectId, locationId, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkServiceListRequest struct via the builder pattern


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


## NetworkingProjectNetworkTagCreate

> Tag NetworkingProjectNetworkTagCreate(ctx, projectId, locationId, networkId).Tag(tag).Execute()

Create networking/network.tag



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
    networkId := "networkId_example" // string | Network Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkTagCreate(context.Background(), projectId, locationId, networkId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkTagCreateRequest struct via the builder pattern


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


## NetworkingProjectNetworkTagDelete

> NetworkingProjectNetworkTagDelete(ctx, projectId, locationId, networkId, tagId).Execute()

Delete networking/network.tag



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
    networkId := "networkId_example" // string | Network Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkTagDelete(context.Background(), projectId, locationId, networkId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagDelete``: %v\n", err)
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
**networkId** | **string** | Network Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkTagDeleteRequest struct via the builder pattern


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


## NetworkingProjectNetworkTagGet

> Tag NetworkingProjectNetworkTagGet(ctx, projectId, locationId, networkId, tagId).Execute()

Get networking/network.tag



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
    networkId := "networkId_example" // string | Network Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkTagGet(context.Background(), projectId, locationId, networkId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkTagGetRequest struct via the builder pattern


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


## NetworkingProjectNetworkTagList

> []Tag NetworkingProjectNetworkTagList(ctx, projectId, locationId, networkId).Execute()

List networking/network.tag



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
    networkId := "networkId_example" // string | Network Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkTagList(context.Background(), projectId, locationId, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkTagListRequest struct via the builder pattern


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


## NetworkingProjectNetworkTagPut

> []Tag NetworkingProjectNetworkTagPut(ctx, projectId, locationId, networkId).Tag(tag).Execute()

Replace networking/network.tag



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
    networkId := "networkId_example" // string | Network Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkTagPut(context.Background(), projectId, locationId, networkId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkTagPutRequest struct via the builder pattern


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


## NetworkingProjectNetworkUpdate

> Network NetworkingProjectNetworkUpdate(ctx, projectId, locationId, networkId).NetworkingProjectNetworkUpdate(networkingProjectNetworkUpdate).Execute()

Update networking/network



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
    networkId := "networkId_example" // string | Network Id
    networkingProjectNetworkUpdate := *openapiclient.NewNetworkingProjectNetworkUpdate() // NetworkingProjectNetworkUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetworkApi.NetworkingProjectNetworkUpdate(context.Background(), projectId, locationId, networkId).NetworkingProjectNetworkUpdate(networkingProjectNetworkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetworkApi.NetworkingProjectNetworkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetworkUpdate`: Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetworkApi.NetworkingProjectNetworkUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**networkId** | **string** | Network Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetworkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectNetworkUpdate** | [**NetworkingProjectNetworkUpdate**](NetworkingProjectNetworkUpdate.md) |  | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

