# \NetworkingProjectIpApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectIpAssociate**](NetworkingProjectIpApi.md#NetworkingProjectIpAssociate) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/associate | Associate networking/ip
[**NetworkingProjectIpCreate**](NetworkingProjectIpApi.md#NetworkingProjectIpCreate) | **Post** /networking/{locationId}/project/{projectId}/ip | Create networking/ip
[**NetworkingProjectIpDelete**](NetworkingProjectIpApi.md#NetworkingProjectIpDelete) | **Delete** /networking/{locationId}/project/{projectId}/ip/{ipId} | Delete networking/ip
[**NetworkingProjectIpDisassociate**](NetworkingProjectIpApi.md#NetworkingProjectIpDisassociate) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/disassociate | Disassociate networking/ip
[**NetworkingProjectIpEventGet**](NetworkingProjectIpApi.md#NetworkingProjectIpEventGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/event/{eventId} | Get networking/ip.event
[**NetworkingProjectIpEventList**](NetworkingProjectIpApi.md#NetworkingProjectIpEventList) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/event | List networking/ip.event
[**NetworkingProjectIpGet**](NetworkingProjectIpApi.md#NetworkingProjectIpGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId} | Get networking/ip
[**NetworkingProjectIpList**](NetworkingProjectIpApi.md#NetworkingProjectIpList) | **Get** /networking/{locationId}/project/{projectId}/ip | List networking/ip
[**NetworkingProjectIpPersist**](NetworkingProjectIpApi.md#NetworkingProjectIpPersist) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/persist | Persist networking/ip
[**NetworkingProjectIpServiceGet**](NetworkingProjectIpApi.md#NetworkingProjectIpServiceGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/service/{serviceId} | Get networking/ip.service
[**NetworkingProjectIpServiceList**](NetworkingProjectIpApi.md#NetworkingProjectIpServiceList) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/service | List networking/ip.service
[**NetworkingProjectIpTagCreate**](NetworkingProjectIpApi.md#NetworkingProjectIpTagCreate) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag | Create networking/ip.tag
[**NetworkingProjectIpTagDelete**](NetworkingProjectIpApi.md#NetworkingProjectIpTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag/{tagId} | Delete networking/ip.tag
[**NetworkingProjectIpTagGet**](NetworkingProjectIpApi.md#NetworkingProjectIpTagGet) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag/{tagId} | Get networking/ip.tag
[**NetworkingProjectIpTagList**](NetworkingProjectIpApi.md#NetworkingProjectIpTagList) | **Get** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag | List networking/ip.tag
[**NetworkingProjectIpTagPut**](NetworkingProjectIpApi.md#NetworkingProjectIpTagPut) | **Put** /networking/{locationId}/project/{projectId}/ip/{ipId}/tag | Replace networking/ip.tag
[**NetworkingProjectIpTransfer**](NetworkingProjectIpApi.md#NetworkingProjectIpTransfer) | **Post** /networking/{locationId}/project/{projectId}/ip/{ipId}/actions/transfer | Transfer networking/ip
[**NetworkingProjectIpUpdate**](NetworkingProjectIpApi.md#NetworkingProjectIpUpdate) | **Patch** /networking/{locationId}/project/{projectId}/ip/{ipId} | Update networking/ip



## NetworkingProjectIpAssociate

> Ip NetworkingProjectIpAssociate(ctx, projectId, locationId, ipId).NetworkingProjectIpAssociate(networkingProjectIpAssociate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Associate networking/ip



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
    ipId := "ipId_example" // string | Ip Id
    networkingProjectIpAssociate := *openapiclient.NewNetworkingProjectIpAssociate("Ip_example") // NetworkingProjectIpAssociate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpAssociate(context.Background(), projectId, locationId, ipId).NetworkingProjectIpAssociate(networkingProjectIpAssociate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpAssociate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpAssociate`: Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpAssociate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpAssociateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectIpAssociate** | [**NetworkingProjectIpAssociate**](NetworkingProjectIpAssociate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpCreate

> Ip NetworkingProjectIpCreate(ctx, projectId, locationId).NetworkingProjectIpCreate(networkingProjectIpCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create networking/ip



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
    networkingProjectIpCreate := *openapiclient.NewNetworkingProjectIpCreate() // NetworkingProjectIpCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpCreate(context.Background(), projectId, locationId).NetworkingProjectIpCreate(networkingProjectIpCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpCreate`: Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkingProjectIpCreate** | [**NetworkingProjectIpCreate**](NetworkingProjectIpCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpDelete

> InlineResponseDefault NetworkingProjectIpDelete(ctx, projectId, locationId, ipId).Execute()

Delete networking/ip



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
    ipId := "ipId_example" // string | Ip Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpDelete(context.Background(), projectId, locationId, ipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpDelete`: InlineResponseDefault
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpDeleteRequest struct via the builder pattern


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


## NetworkingProjectIpDisassociate

> Ip NetworkingProjectIpDisassociate(ctx, projectId, locationId, ipId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Disassociate networking/ip



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
    ipId := "ipId_example" // string | Ip Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpDisassociate(context.Background(), projectId, locationId, ipId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpDisassociate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpDisassociate`: Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpDisassociate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpDisassociateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpEventGet

> Event NetworkingProjectIpEventGet(ctx, projectId, locationId, ipId, eventId).Execute()

Get networking/ip.event



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
    ipId := "ipId_example" // string | Ip Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpEventGet(context.Background(), projectId, locationId, ipId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpEventGetRequest struct via the builder pattern


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


## NetworkingProjectIpEventList

> []Event NetworkingProjectIpEventList(ctx, projectId, locationId, ipId).Limit(limit).Skip(skip).Execute()

List networking/ip.event



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
    ipId := "ipId_example" // string | Ip Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpEventList(context.Background(), projectId, locationId, ipId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpEventListRequest struct via the builder pattern


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


## NetworkingProjectIpGet

> Ip NetworkingProjectIpGet(ctx, projectId, locationId, ipId).Execute()

Get networking/ip



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
    ipId := "ipId_example" // string | Ip Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpGet(context.Background(), projectId, locationId, ipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpGet`: Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpList

> []Ip NetworkingProjectIpList(ctx, projectId, locationId).Network(network).AssociatedNetadp(associatedNetadp).TagValue(tagValue).TagKey(tagKey).Execute()

List networking/ip



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
    network := "network_example" // string | Filter by network (optional)
    associatedNetadp := "associatedNetadp_example" // string | Filter by associated.netadp (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpList(context.Background(), projectId, locationId).Network(network).AssociatedNetadp(associatedNetadp).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpList`: []Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **network** | **string** | Filter by network | 
 **associatedNetadp** | **string** | Filter by associated.netadp | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpPersist

> Ip NetworkingProjectIpPersist(ctx, projectId, locationId, ipId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Persist networking/ip



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
    ipId := "ipId_example" // string | Ip Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpPersist(context.Background(), projectId, locationId, ipId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpPersist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpPersist`: Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpPersist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpPersistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpServiceGet

> ResourceService NetworkingProjectIpServiceGet(ctx, projectId, locationId, ipId, serviceId).Execute()

Get networking/ip.service



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
    ipId := "ipId_example" // string | Ip Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpServiceGet(context.Background(), projectId, locationId, ipId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpServiceGetRequest struct via the builder pattern


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


## NetworkingProjectIpServiceList

> []ResourceService NetworkingProjectIpServiceList(ctx, projectId, locationId, ipId).Execute()

List networking/ip.service



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
    ipId := "ipId_example" // string | Ip Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpServiceList(context.Background(), projectId, locationId, ipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpServiceListRequest struct via the builder pattern


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


## NetworkingProjectIpTagCreate

> Tag NetworkingProjectIpTagCreate(ctx, projectId, locationId, ipId).Tag(tag).Execute()

Create networking/ip.tag



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
    ipId := "ipId_example" // string | Ip Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpTagCreate(context.Background(), projectId, locationId, ipId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpTagCreateRequest struct via the builder pattern


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


## NetworkingProjectIpTagDelete

> NetworkingProjectIpTagDelete(ctx, projectId, locationId, ipId, tagId).Execute()

Delete networking/ip.tag



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
    ipId := "ipId_example" // string | Ip Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpTagDelete(context.Background(), projectId, locationId, ipId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpTagDelete``: %v\n", err)
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
**ipId** | **string** | Ip Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpTagDeleteRequest struct via the builder pattern


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


## NetworkingProjectIpTagGet

> Tag NetworkingProjectIpTagGet(ctx, projectId, locationId, ipId, tagId).Execute()

Get networking/ip.tag



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
    ipId := "ipId_example" // string | Ip Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpTagGet(context.Background(), projectId, locationId, ipId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpTagGetRequest struct via the builder pattern


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


## NetworkingProjectIpTagList

> []Tag NetworkingProjectIpTagList(ctx, projectId, locationId, ipId).Execute()

List networking/ip.tag



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
    ipId := "ipId_example" // string | Ip Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpTagList(context.Background(), projectId, locationId, ipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpTagListRequest struct via the builder pattern


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


## NetworkingProjectIpTagPut

> []Tag NetworkingProjectIpTagPut(ctx, projectId, locationId, ipId).Tag(tag).Execute()

Replace networking/ip.tag



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
    ipId := "ipId_example" // string | Ip Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpTagPut(context.Background(), projectId, locationId, ipId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpTagPutRequest struct via the builder pattern


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


## NetworkingProjectIpTransfer

> Ip NetworkingProjectIpTransfer(ctx, projectId, locationId, ipId).NetworkingProjectIpTransfer(networkingProjectIpTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer networking/ip



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
    ipId := "ipId_example" // string | Ip Id
    networkingProjectIpTransfer := *openapiclient.NewNetworkingProjectIpTransfer("Project_example") // NetworkingProjectIpTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpTransfer(context.Background(), projectId, locationId, ipId).NetworkingProjectIpTransfer(networkingProjectIpTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpTransfer`: Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectIpTransfer** | [**NetworkingProjectIpTransfer**](NetworkingProjectIpTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectIpUpdate

> Ip NetworkingProjectIpUpdate(ctx, projectId, locationId, ipId).NetworkingProjectIpUpdate(networkingProjectIpUpdate).Execute()

Update networking/ip



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
    ipId := "ipId_example" // string | Ip Id
    networkingProjectIpUpdate := *openapiclient.NewNetworkingProjectIpUpdate() // NetworkingProjectIpUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectIpApi.NetworkingProjectIpUpdate(context.Background(), projectId, locationId, ipId).NetworkingProjectIpUpdate(networkingProjectIpUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectIpApi.NetworkingProjectIpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectIpUpdate`: Ip
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectIpApi.NetworkingProjectIpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**ipId** | **string** | Ip Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectIpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectIpUpdate** | [**NetworkingProjectIpUpdate**](NetworkingProjectIpUpdate.md) |  | 

### Return type

[**Ip**](Ip.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

