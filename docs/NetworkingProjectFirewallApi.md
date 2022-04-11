# \NetworkingProjectFirewallApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectFirewallCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall | Create networking/firewall
[**NetworkingProjectFirewallDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId} | Delete networking/firewall
[**NetworkingProjectFirewallEgressCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress | Create networking/firewall.egress
[**NetworkingProjectFirewallEgressDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress/{egressId} | Delete networking/firewall.egress
[**NetworkingProjectFirewallEgressGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress/{egressId} | Get networking/firewall.egress
[**NetworkingProjectFirewallEgressList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress | List networking/firewall.egress
[**NetworkingProjectFirewallEgressPut**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressPut) | **Put** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress | Replace networking/firewall.egress
[**NetworkingProjectFirewallEventGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEventGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/event/{eventId} | Get networking/firewall.event
[**NetworkingProjectFirewallEventList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEventList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/event | List networking/firewall.event
[**NetworkingProjectFirewallGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId} | Get networking/firewall
[**NetworkingProjectFirewallIngressCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress | Create networking/firewall.ingress
[**NetworkingProjectFirewallIngressDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress/{ingressId} | Delete networking/firewall.ingress
[**NetworkingProjectFirewallIngressGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress/{ingressId} | Get networking/firewall.ingress
[**NetworkingProjectFirewallIngressList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress | List networking/firewall.ingress
[**NetworkingProjectFirewallIngressPut**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressPut) | **Put** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress | Replace networking/firewall.ingress
[**NetworkingProjectFirewallList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallList) | **Get** /networking/{locationId}/project/{projectId}/firewall | List networking/firewall
[**NetworkingProjectFirewallServiceGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallServiceGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/service/{serviceId} | Get networking/firewall.service
[**NetworkingProjectFirewallServiceList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallServiceList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/service | List networking/firewall.service
[**NetworkingProjectFirewallTagCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag | Create networking/firewall.tag
[**NetworkingProjectFirewallTagDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag/{tagId} | Delete networking/firewall.tag
[**NetworkingProjectFirewallTagGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag/{tagId} | Get networking/firewall.tag
[**NetworkingProjectFirewallTagList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag | List networking/firewall.tag
[**NetworkingProjectFirewallTagPut**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagPut) | **Put** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag | Replace networking/firewall.tag
[**NetworkingProjectFirewallTransfer**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTransfer) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/actions/transfer | Transfer networking/firewall
[**NetworkingProjectFirewallUpdate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallUpdate) | **Patch** /networking/{locationId}/project/{projectId}/firewall/{firewallId} | Update networking/firewall



## NetworkingProjectFirewallCreate

> Firewall NetworkingProjectFirewallCreate(ctx, projectId, locationId).NetworkingProjectFirewallCreate(networkingProjectFirewallCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create networking/firewall



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
    networkingProjectFirewallCreate := *openapiclient.NewNetworkingProjectFirewallCreate("Name_example") // NetworkingProjectFirewallCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallCreate(context.Background(), projectId, locationId).NetworkingProjectFirewallCreate(networkingProjectFirewallCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallCreate`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkingProjectFirewallCreate** | [**NetworkingProjectFirewallCreate**](NetworkingProjectFirewallCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallDelete

> NetworkingProjectFirewallDelete(ctx, projectId, locationId, firewallId).Execute()

Delete networking/firewall



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
    firewallId := "firewallId_example" // string | Firewall Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallDelete(context.Background(), projectId, locationId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallDelete``: %v\n", err)
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
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallDeleteRequest struct via the builder pattern


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


## NetworkingProjectFirewallEgressCreate

> NetworkingRule NetworkingProjectFirewallEgressCreate(ctx, projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()

Create networking/firewall.egress



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
    firewallId := "firewallId_example" // string | Firewall Id
    networkingRule := *openapiclient.NewNetworkingRule("Name_example", "Action_example", float32(123), []string{"Filter_example"}) // NetworkingRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressCreate(context.Background(), projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallEgressCreate`: NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallEgressCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingRule** | [**NetworkingRule**](NetworkingRule.md) |  | 

### Return type

[**NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEgressDelete

> NetworkingProjectFirewallEgressDelete(ctx, projectId, locationId, firewallId, egressId).Execute()

Delete networking/firewall.egress



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
    firewallId := "firewallId_example" // string | Firewall Id
    egressId := "egressId_example" // string | egressId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressDelete(context.Background(), projectId, locationId, firewallId, egressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressDelete``: %v\n", err)
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
**firewallId** | **string** | Firewall Id | 
**egressId** | **string** | egressId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallEgressDeleteRequest struct via the builder pattern


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


## NetworkingProjectFirewallEgressGet

> NetworkingRule NetworkingProjectFirewallEgressGet(ctx, projectId, locationId, firewallId, egressId).Execute()

Get networking/firewall.egress



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
    firewallId := "firewallId_example" // string | Firewall Id
    egressId := "egressId_example" // string | egressId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressGet(context.Background(), projectId, locationId, firewallId, egressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallEgressGet`: NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 
**egressId** | **string** | egressId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallEgressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEgressList

> []NetworkingRule NetworkingProjectFirewallEgressList(ctx, projectId, locationId, firewallId).Execute()

List networking/firewall.egress



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
    firewallId := "firewallId_example" // string | Firewall Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressList(context.Background(), projectId, locationId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallEgressList`: []NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallEgressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEgressPut

> []NetworkingRule NetworkingProjectFirewallEgressPut(ctx, projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()

Replace networking/firewall.egress



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
    firewallId := "firewallId_example" // string | Firewall Id
    networkingRule := []openapiclient.NetworkingRule{*openapiclient.NewNetworkingRule("Name_example", "Action_example", float32(123), []string{"Filter_example"})} // []NetworkingRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressPut(context.Background(), projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallEgressPut`: []NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallEgressPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallEgressPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingRule** | [**[]NetworkingRule**](NetworkingRule.md) |  | 

### Return type

[**[]NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEventGet

> Event NetworkingProjectFirewallEventGet(ctx, projectId, locationId, firewallId, eventId).Execute()

Get networking/firewall.event



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
    firewallId := "firewallId_example" // string | Firewall Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallEventGet(context.Background(), projectId, locationId, firewallId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallEventGetRequest struct via the builder pattern


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


## NetworkingProjectFirewallEventList

> []Event NetworkingProjectFirewallEventList(ctx, projectId, locationId, firewallId).Limit(limit).Skip(skip).Execute()

List networking/firewall.event



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
    firewallId := "firewallId_example" // string | Firewall Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallEventList(context.Background(), projectId, locationId, firewallId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallEventListRequest struct via the builder pattern


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


## NetworkingProjectFirewallGet

> Firewall NetworkingProjectFirewallGet(ctx, projectId, locationId, firewallId).Execute()

Get networking/firewall



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
    firewallId := "firewallId_example" // string | Firewall Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallGet(context.Background(), projectId, locationId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallGet`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressCreate

> NetworkingRule NetworkingProjectFirewallIngressCreate(ctx, projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()

Create networking/firewall.ingress



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
    firewallId := "firewallId_example" // string | Firewall Id
    networkingRule := *openapiclient.NewNetworkingRule("Name_example", "Action_example", float32(123), []string{"Filter_example"}) // NetworkingRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressCreate(context.Background(), projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallIngressCreate`: NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallIngressCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingRule** | [**NetworkingRule**](NetworkingRule.md) |  | 

### Return type

[**NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressDelete

> NetworkingProjectFirewallIngressDelete(ctx, projectId, locationId, firewallId, ingressId).Execute()

Delete networking/firewall.ingress



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
    firewallId := "firewallId_example" // string | Firewall Id
    ingressId := "ingressId_example" // string | ingressId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressDelete(context.Background(), projectId, locationId, firewallId, ingressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressDelete``: %v\n", err)
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
**firewallId** | **string** | Firewall Id | 
**ingressId** | **string** | ingressId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallIngressDeleteRequest struct via the builder pattern


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


## NetworkingProjectFirewallIngressGet

> NetworkingRule NetworkingProjectFirewallIngressGet(ctx, projectId, locationId, firewallId, ingressId).Execute()

Get networking/firewall.ingress



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
    firewallId := "firewallId_example" // string | Firewall Id
    ingressId := "ingressId_example" // string | ingressId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressGet(context.Background(), projectId, locationId, firewallId, ingressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallIngressGet`: NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 
**ingressId** | **string** | ingressId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallIngressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressList

> []NetworkingRule NetworkingProjectFirewallIngressList(ctx, projectId, locationId, firewallId).Execute()

List networking/firewall.ingress



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
    firewallId := "firewallId_example" // string | Firewall Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressList(context.Background(), projectId, locationId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallIngressList`: []NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallIngressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressPut

> []NetworkingRule NetworkingProjectFirewallIngressPut(ctx, projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()

Replace networking/firewall.ingress



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
    firewallId := "firewallId_example" // string | Firewall Id
    networkingRule := []openapiclient.NetworkingRule{*openapiclient.NewNetworkingRule("Name_example", "Action_example", float32(123), []string{"Filter_example"})} // []NetworkingRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressPut(context.Background(), projectId, locationId, firewallId).NetworkingRule(networkingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallIngressPut`: []NetworkingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallIngressPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallIngressPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingRule** | [**[]NetworkingRule**](NetworkingRule.md) |  | 

### Return type

[**[]NetworkingRule**](NetworkingRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallList

> []Firewall NetworkingProjectFirewallList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List networking/firewall



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
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallList`: []Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallServiceGet

> ResourceService NetworkingProjectFirewallServiceGet(ctx, projectId, locationId, firewallId, serviceId).Execute()

Get networking/firewall.service



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
    firewallId := "firewallId_example" // string | Firewall Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallServiceGet(context.Background(), projectId, locationId, firewallId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallServiceGetRequest struct via the builder pattern


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


## NetworkingProjectFirewallServiceList

> []ResourceService NetworkingProjectFirewallServiceList(ctx, projectId, locationId, firewallId).Execute()

List networking/firewall.service



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
    firewallId := "firewallId_example" // string | Firewall Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallServiceList(context.Background(), projectId, locationId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallServiceListRequest struct via the builder pattern


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


## NetworkingProjectFirewallTagCreate

> Tag NetworkingProjectFirewallTagCreate(ctx, projectId, locationId, firewallId).Tag(tag).Execute()

Create networking/firewall.tag



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
    firewallId := "firewallId_example" // string | Firewall Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallTagCreate(context.Background(), projectId, locationId, firewallId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallTagCreateRequest struct via the builder pattern


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


## NetworkingProjectFirewallTagDelete

> NetworkingProjectFirewallTagDelete(ctx, projectId, locationId, firewallId, tagId).Execute()

Delete networking/firewall.tag



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
    firewallId := "firewallId_example" // string | Firewall Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallTagDelete(context.Background(), projectId, locationId, firewallId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagDelete``: %v\n", err)
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
**firewallId** | **string** | Firewall Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallTagDeleteRequest struct via the builder pattern


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


## NetworkingProjectFirewallTagGet

> Tag NetworkingProjectFirewallTagGet(ctx, projectId, locationId, firewallId, tagId).Execute()

Get networking/firewall.tag



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
    firewallId := "firewallId_example" // string | Firewall Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallTagGet(context.Background(), projectId, locationId, firewallId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallTagGetRequest struct via the builder pattern


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


## NetworkingProjectFirewallTagList

> []Tag NetworkingProjectFirewallTagList(ctx, projectId, locationId, firewallId).Execute()

List networking/firewall.tag



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
    firewallId := "firewallId_example" // string | Firewall Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallTagList(context.Background(), projectId, locationId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallTagListRequest struct via the builder pattern


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


## NetworkingProjectFirewallTagPut

> []Tag NetworkingProjectFirewallTagPut(ctx, projectId, locationId, firewallId).Tag(tag).Execute()

Replace networking/firewall.tag



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
    firewallId := "firewallId_example" // string | Firewall Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallTagPut(context.Background(), projectId, locationId, firewallId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallTagPutRequest struct via the builder pattern


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


## NetworkingProjectFirewallTransfer

> Firewall NetworkingProjectFirewallTransfer(ctx, projectId, locationId, firewallId).NetworkingProjectFirewallTransfer(networkingProjectFirewallTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer networking/firewall



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
    firewallId := "firewallId_example" // string | Firewall Id
    networkingProjectFirewallTransfer := *openapiclient.NewNetworkingProjectFirewallTransfer("Project_example") // NetworkingProjectFirewallTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallTransfer(context.Background(), projectId, locationId, firewallId).NetworkingProjectFirewallTransfer(networkingProjectFirewallTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallTransfer`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectFirewallTransfer** | [**NetworkingProjectFirewallTransfer**](NetworkingProjectFirewallTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallUpdate

> Firewall NetworkingProjectFirewallUpdate(ctx, projectId, locationId, firewallId).NetworkingProjectFirewallUpdate(networkingProjectFirewallUpdate).Execute()

Update networking/firewall



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
    firewallId := "firewallId_example" // string | Firewall Id
    networkingProjectFirewallUpdate := *openapiclient.NewNetworkingProjectFirewallUpdate() // NetworkingProjectFirewallUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectFirewallApi.NetworkingProjectFirewallUpdate(context.Background(), projectId, locationId, firewallId).NetworkingProjectFirewallUpdate(networkingProjectFirewallUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectFirewallApi.NetworkingProjectFirewallUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectFirewallUpdate`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectFirewallApi.NetworkingProjectFirewallUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**firewallId** | **string** | Firewall Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectFirewallUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectFirewallUpdate** | [**NetworkingProjectFirewallUpdate**](NetworkingProjectFirewallUpdate.md) |  | 

### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

