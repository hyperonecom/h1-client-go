# \NetworkingProjectNetadpApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectNetadpCreate**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpCreate) | **Post** /networking/{locationId}/project/{projectId}/netadp | Create networking/netadp
[**NetworkingProjectNetadpDelete**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpDelete) | **Delete** /networking/{locationId}/project/{projectId}/netadp/{netadpId} | Delete networking/netadp
[**NetworkingProjectNetadpEventGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpEventGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/event/{eventId} | Get networking/netadp.event
[**NetworkingProjectNetadpEventList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpEventList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/event | List networking/netadp.event
[**NetworkingProjectNetadpGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId} | Get networking/netadp
[**NetworkingProjectNetadpList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpList) | **Get** /networking/{locationId}/project/{projectId}/netadp | List networking/netadp
[**NetworkingProjectNetadpMetricGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpMetricGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/metric/{metricId} | Get networking/netadp.metric
[**NetworkingProjectNetadpMetricList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpMetricList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/metric | List networking/netadp.metric
[**NetworkingProjectNetadpMetricSeriesList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpMetricSeriesList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/metric/{metricId}/series | List networking/netadp.series
[**NetworkingProjectNetadpServiceGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpServiceGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/service/{serviceId} | Get networking/netadp.service
[**NetworkingProjectNetadpServiceList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpServiceList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/service | List networking/netadp.service
[**NetworkingProjectNetadpTagCreate**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagCreate) | **Post** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag | Create networking/netadp.tag
[**NetworkingProjectNetadpTagDelete**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag/{tagId} | Delete networking/netadp.tag
[**NetworkingProjectNetadpTagGet**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagGet) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag/{tagId} | Get networking/netadp.tag
[**NetworkingProjectNetadpTagList**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagList) | **Get** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag | List networking/netadp.tag
[**NetworkingProjectNetadpTagPut**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpTagPut) | **Put** /networking/{locationId}/project/{projectId}/netadp/{netadpId}/tag | Replace networking/netadp.tag
[**NetworkingProjectNetadpUpdate**](NetworkingProjectNetadpApi.md#NetworkingProjectNetadpUpdate) | **Patch** /networking/{locationId}/project/{projectId}/netadp/{netadpId} | Update networking/netadp



## NetworkingProjectNetadpCreate

> Netadp NetworkingProjectNetadpCreate(ctx, projectId, locationId).NetworkingProjectNetadpCreate(networkingProjectNetadpCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create networking/netadp



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
    networkingProjectNetadpCreate := *openapiclient.NewNetworkingProjectNetadpCreate("Network_example", "Vm_example") // NetworkingProjectNetadpCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpCreate(context.Background(), projectId, locationId).NetworkingProjectNetadpCreate(networkingProjectNetadpCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpCreate`: Netadp
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkingProjectNetadpCreate** | [**NetworkingProjectNetadpCreate**](NetworkingProjectNetadpCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Netadp**](Netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetadpDelete

> NetworkingProjectNetadpDelete(ctx, projectId, locationId, netadpId).Execute()

Delete networking/netadp



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
    netadpId := "netadpId_example" // string | Netadp Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpDelete(context.Background(), projectId, locationId, netadpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpDelete``: %v\n", err)
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
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpDeleteRequest struct via the builder pattern


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


## NetworkingProjectNetadpEventGet

> Event NetworkingProjectNetadpEventGet(ctx, projectId, locationId, netadpId, eventId).Execute()

Get networking/netadp.event



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
    netadpId := "netadpId_example" // string | Netadp Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpEventGet(context.Background(), projectId, locationId, netadpId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpEventGetRequest struct via the builder pattern


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


## NetworkingProjectNetadpEventList

> []Event NetworkingProjectNetadpEventList(ctx, projectId, locationId, netadpId).Limit(limit).Skip(skip).Execute()

List networking/netadp.event



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
    netadpId := "netadpId_example" // string | Netadp Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpEventList(context.Background(), projectId, locationId, netadpId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpEventListRequest struct via the builder pattern


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


## NetworkingProjectNetadpGet

> Netadp NetworkingProjectNetadpGet(ctx, projectId, locationId, netadpId).Execute()

Get networking/netadp



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
    netadpId := "netadpId_example" // string | Netadp Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpGet(context.Background(), projectId, locationId, netadpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpGet`: Netadp
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Netadp**](Netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetadpList

> []Netadp NetworkingProjectNetadpList(ctx, projectId, locationId).AssignedResource(assignedResource).AssignedId(assignedId).Network(network).TagValue(tagValue).TagKey(tagKey).Execute()

List networking/netadp



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
    assignedResource := "assignedResource_example" // string | Filter by assigned.resource (optional)
    assignedId := "assignedId_example" // string | Filter by assigned.id (optional)
    network := "network_example" // string | Filter by network (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpList(context.Background(), projectId, locationId).AssignedResource(assignedResource).AssignedId(assignedId).Network(network).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpList`: []Netadp
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignedResource** | **string** | Filter by assigned.resource | 
 **assignedId** | **string** | Filter by assigned.id | 
 **network** | **string** | Filter by network | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Netadp**](Netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectNetadpMetricGet

> Metric NetworkingProjectNetadpMetricGet(ctx, projectId, locationId, netadpId, metricId).Execute()

Get networking/netadp.metric



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
    netadpId := "netadpId_example" // string | Netadp Id
    metricId := "metricId_example" // string | metricId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricGet(context.Background(), projectId, locationId, netadpId, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpMetricGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpMetricGetRequest struct via the builder pattern


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


## NetworkingProjectNetadpMetricList

> []Metric NetworkingProjectNetadpMetricList(ctx, projectId, locationId, netadpId).Execute()

List networking/netadp.metric



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
    netadpId := "netadpId_example" // string | Netadp Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricList(context.Background(), projectId, locationId, netadpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpMetricList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpMetricListRequest struct via the builder pattern


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


## NetworkingProjectNetadpMetricSeriesList

> []Point NetworkingProjectNetadpMetricSeriesList(ctx, projectId, locationId, netadpId, metricId).Interval(interval).Timespan(timespan).Aligner(aligner).Execute()

List networking/netadp.series



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
    netadpId := "netadpId_example" // string | Netadp Id
    metricId := "metricId_example" // string | metricId
    interval := "interval_example" // string | interval (optional)
    timespan := "timespan_example" // string | timespan (optional)
    aligner := "aligner_example" // string | aligner (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricSeriesList(context.Background(), projectId, locationId, netadpId, metricId).Interval(interval).Timespan(timespan).Aligner(aligner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricSeriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpMetricSeriesList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpMetricSeriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpMetricSeriesListRequest struct via the builder pattern


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


## NetworkingProjectNetadpServiceGet

> ResourceService NetworkingProjectNetadpServiceGet(ctx, projectId, locationId, netadpId, serviceId).Execute()

Get networking/netadp.service



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
    netadpId := "netadpId_example" // string | Netadp Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpServiceGet(context.Background(), projectId, locationId, netadpId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpServiceGetRequest struct via the builder pattern


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


## NetworkingProjectNetadpServiceList

> []ResourceService NetworkingProjectNetadpServiceList(ctx, projectId, locationId, netadpId).Execute()

List networking/netadp.service



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
    netadpId := "netadpId_example" // string | Netadp Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpServiceList(context.Background(), projectId, locationId, netadpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpServiceListRequest struct via the builder pattern


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


## NetworkingProjectNetadpTagCreate

> Tag NetworkingProjectNetadpTagCreate(ctx, projectId, locationId, netadpId).Tag(tag).Execute()

Create networking/netadp.tag



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
    netadpId := "netadpId_example" // string | Netadp Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpTagCreate(context.Background(), projectId, locationId, netadpId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpTagCreateRequest struct via the builder pattern


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


## NetworkingProjectNetadpTagDelete

> NetworkingProjectNetadpTagDelete(ctx, projectId, locationId, netadpId, tagId).Execute()

Delete networking/netadp.tag



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
    netadpId := "netadpId_example" // string | Netadp Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpTagDelete(context.Background(), projectId, locationId, netadpId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagDelete``: %v\n", err)
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
**netadpId** | **string** | Netadp Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpTagDeleteRequest struct via the builder pattern


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


## NetworkingProjectNetadpTagGet

> Tag NetworkingProjectNetadpTagGet(ctx, projectId, locationId, netadpId, tagId).Execute()

Get networking/netadp.tag



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
    netadpId := "netadpId_example" // string | Netadp Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpTagGet(context.Background(), projectId, locationId, netadpId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpTagGetRequest struct via the builder pattern


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


## NetworkingProjectNetadpTagList

> []Tag NetworkingProjectNetadpTagList(ctx, projectId, locationId, netadpId).Execute()

List networking/netadp.tag



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
    netadpId := "netadpId_example" // string | Netadp Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpTagList(context.Background(), projectId, locationId, netadpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpTagListRequest struct via the builder pattern


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


## NetworkingProjectNetadpTagPut

> []Tag NetworkingProjectNetadpTagPut(ctx, projectId, locationId, netadpId).Tag(tag).Execute()

Replace networking/netadp.tag



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
    netadpId := "netadpId_example" // string | Netadp Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpTagPut(context.Background(), projectId, locationId, netadpId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpTagPutRequest struct via the builder pattern


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


## NetworkingProjectNetadpUpdate

> Netadp NetworkingProjectNetadpUpdate(ctx, projectId, locationId, netadpId).NetworkingProjectNetadpUpdate(networkingProjectNetadpUpdate).Execute()

Update networking/netadp



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
    netadpId := "netadpId_example" // string | Netadp Id
    networkingProjectNetadpUpdate := *openapiclient.NewNetworkingProjectNetadpUpdate() // NetworkingProjectNetadpUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkingProjectNetadpApi.NetworkingProjectNetadpUpdate(context.Background(), projectId, locationId, netadpId).NetworkingProjectNetadpUpdate(networkingProjectNetadpUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkingProjectNetadpApi.NetworkingProjectNetadpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkingProjectNetadpUpdate`: Netadp
    fmt.Fprintf(os.Stdout, "Response from `NetworkingProjectNetadpApi.NetworkingProjectNetadpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**netadpId** | **string** | Netadp Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkingProjectNetadpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkingProjectNetadpUpdate** | [**NetworkingProjectNetadpUpdate**](NetworkingProjectNetadpUpdate.md) |  | 

### Return type

[**Netadp**](Netadp.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

