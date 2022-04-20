# \NetworkingProjectNetgwApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectNetgwAttach**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwAttach) | **Post** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/actions/attach | Attach networking/netgw
[**NetworkingProjectNetgwCreate**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwCreate) | **Post** /networking/{locationId}/project/{projectId}/netgw | Create networking/netgw
[**NetworkingProjectNetgwDelete**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwDelete) | **Delete** /networking/{locationId}/project/{projectId}/netgw/{netgwId} | Delete networking/netgw
[**NetworkingProjectNetgwDetach**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwDetach) | **Post** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/actions/detach | Detach networking/netgw
[**NetworkingProjectNetgwEventGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwEventGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/event/{eventId} | Get networking/netgw.event
[**NetworkingProjectNetgwEventList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwEventList) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/event | List networking/netgw.event
[**NetworkingProjectNetgwGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId} | Get networking/netgw
[**NetworkingProjectNetgwList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwList) | **Get** /networking/{locationId}/project/{projectId}/netgw | List networking/netgw
[**NetworkingProjectNetgwServiceGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwServiceGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/service/{serviceId} | Get networking/netgw.service
[**NetworkingProjectNetgwServiceList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwServiceList) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/service | List networking/netgw.service
[**NetworkingProjectNetgwTagCreate**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagCreate) | **Post** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag | Create networking/netgw.tag
[**NetworkingProjectNetgwTagDelete**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag/{tagId} | Delete networking/netgw.tag
[**NetworkingProjectNetgwTagGet**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagGet) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag/{tagId} | Get networking/netgw.tag
[**NetworkingProjectNetgwTagList**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagList) | **Get** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag | List networking/netgw.tag
[**NetworkingProjectNetgwTagPut**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwTagPut) | **Put** /networking/{locationId}/project/{projectId}/netgw/{netgwId}/tag | Replace networking/netgw.tag
[**NetworkingProjectNetgwUpdate**](NetworkingProjectNetgwApi.md#NetworkingProjectNetgwUpdate) | **Patch** /networking/{locationId}/project/{projectId}/netgw/{netgwId} | Update networking/netgw



## NetworkingProjectNetgwAttach

> Netgw NetworkingProjectNetgwAttach(ctx, projectId, locationId, netgwId).NetworkingProjectNetgwAttach(networkingProjectNetgwAttach).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Attach networking/netgw



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
    netgwId := "netgwId_example" // string | Netgw Id
    networkingProjectNetgwAttach := *openapiclient.NewNetworkingProjectNetgwAttach() // NetworkingProjectNetgwAttach | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwAttach(context.Background(), projectId, locationId, netgwId).NetworkingProjectNetgwAttach(networkingProjectNetgwAttach).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwAttach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwAttach`: Netgw
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwAttach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwAttachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectNetgwAttach** | [**NetworkingProjectNetgwAttach**](NetworkingProjectNetgwAttach.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Netgw**](Netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwCreate

> Netgw NetworkingProjectNetgwCreate(ctx, projectId, locationId).NetworkingProjectNetgwCreate(networkingProjectNetgwCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create networking/netgw



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
    networkingProjectNetgwCreate := *openapiclient.NewNetworkingProjectNetgwCreate() // NetworkingProjectNetgwCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwCreate(context.Background(), projectId, locationId).NetworkingProjectNetgwCreate(networkingProjectNetgwCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwCreate`: Netgw
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkingProjectNetgwCreate** | [**NetworkingProjectNetgwCreate**](NetworkingProjectNetgwCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Netgw**](Netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwDelete

> Netgw NetworkingProjectNetgwDelete(ctx, projectId, locationId, netgwId).Execute()

Delete networking/netgw



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
    netgwId := "netgwId_example" // string | Netgw Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwDelete(context.Background(), projectId, locationId, netgwId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwDelete`: Netgw
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Netgw**](Netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwDetach

> Netgw NetworkingProjectNetgwDetach(ctx, projectId, locationId, netgwId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Detach networking/netgw



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
    netgwId := "netgwId_example" // string | Netgw Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwDetach(context.Background(), projectId, locationId, netgwId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwDetach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwDetach`: Netgw
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwDetach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwDetachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Netgw**](Netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwEventGet

> Event NetworkingProjectNetgwEventGet(ctx, projectId, locationId, netgwId, eventId).Execute()

Get networking/netgw.event



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
    netgwId := "netgwId_example" // string | Netgw Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwEventGet(context.Background(), projectId, locationId, netgwId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwEventGetRequest struct via the builder pattern


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


## NetworkingProjectNetgwEventList

> []Event NetworkingProjectNetgwEventList(ctx, projectId, locationId, netgwId).Limit(limit).Skip(skip).Execute()

List networking/netgw.event



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
    netgwId := "netgwId_example" // string | Netgw Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwEventList(context.Background(), projectId, locationId, netgwId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwEventListRequest struct via the builder pattern


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


## NetworkingProjectNetgwGet

> Netgw NetworkingProjectNetgwGet(ctx, projectId, locationId, netgwId).Execute()

Get networking/netgw



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
    netgwId := "netgwId_example" // string | Netgw Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwGet(context.Background(), projectId, locationId, netgwId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwGet`: Netgw
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Netgw**](Netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwList

> []Netgw NetworkingProjectNetgwList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List networking/netgw



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
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwList`: []Netgw
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Netgw**](Netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetgwServiceGet

> ResourceService NetworkingProjectNetgwServiceGet(ctx, projectId, locationId, netgwId, serviceId).Execute()

Get networking/netgw.service



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
    netgwId := "netgwId_example" // string | Netgw Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwServiceGet(context.Background(), projectId, locationId, netgwId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwServiceGetRequest struct via the builder pattern


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


## NetworkingProjectNetgwServiceList

> []ResourceService NetworkingProjectNetgwServiceList(ctx, projectId, locationId, netgwId).Execute()

List networking/netgw.service



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
    netgwId := "netgwId_example" // string | Netgw Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwServiceList(context.Background(), projectId, locationId, netgwId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwServiceListRequest struct via the builder pattern


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


## NetworkingProjectNetgwTagCreate

> Tag NetworkingProjectNetgwTagCreate(ctx, projectId, locationId, netgwId).Tag(tag).Execute()

Create networking/netgw.tag



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
    netgwId := "netgwId_example" // string | Netgw Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwTagCreate(context.Background(), projectId, locationId, netgwId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwTagCreateRequest struct via the builder pattern


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


## NetworkingProjectNetgwTagDelete

> NetworkingProjectNetgwTagDelete(ctx, projectId, locationId, netgwId, tagId).Execute()

Delete networking/netgw.tag



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
    netgwId := "netgwId_example" // string | Netgw Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwTagDelete(context.Background(), projectId, locationId, netgwId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagDelete``: %v\n", err)
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
**netgwId** | **string** | Netgw Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwTagDeleteRequest struct via the builder pattern


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


## NetworkingProjectNetgwTagGet

> Tag NetworkingProjectNetgwTagGet(ctx, projectId, locationId, netgwId, tagId).Execute()

Get networking/netgw.tag



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
    netgwId := "netgwId_example" // string | Netgw Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwTagGet(context.Background(), projectId, locationId, netgwId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwTagGetRequest struct via the builder pattern


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


## NetworkingProjectNetgwTagList

> []Tag NetworkingProjectNetgwTagList(ctx, projectId, locationId, netgwId).Execute()

List networking/netgw.tag



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
    netgwId := "netgwId_example" // string | Netgw Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwTagList(context.Background(), projectId, locationId, netgwId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwTagListRequest struct via the builder pattern


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


## NetworkingProjectNetgwTagPut

> []Tag NetworkingProjectNetgwTagPut(ctx, projectId, locationId, netgwId).Tag(tag).Execute()

Replace networking/netgw.tag



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
    netgwId := "netgwId_example" // string | Netgw Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwTagPut(context.Background(), projectId, locationId, netgwId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwTagPutRequest struct via the builder pattern


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


## NetworkingProjectNetgwUpdate

> Netgw NetworkingProjectNetgwUpdate(ctx, projectId, locationId, netgwId).NetworkingProjectNetgwUpdate(networkingProjectNetgwUpdate).Execute()

Update networking/netgw



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
    netgwId := "netgwId_example" // string | Netgw Id
    networkingProjectNetgwUpdate := *openapiclient.NewNetworkingProjectNetgwUpdate() // NetworkingProjectNetgwUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetgwApi.NetworkingProjectNetgwUpdate(context.Background(), projectId, locationId, netgwId).NetworkingProjectNetgwUpdate(networkingProjectNetgwUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetgwApi.NetworkingProjectNetgwUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetgwUpdate`: Netgw
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetgwApi.NetworkingProjectNetgwUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netgwId** | **string** | Netgw Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetgwUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectNetgwUpdate** | [**NetworkingProjectNetgwUpdate**](NetworkingProjectNetgwUpdate.md) |  | 

### Return type

[**Netgw**](Netgw.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

