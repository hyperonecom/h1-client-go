# \ComputeProjectVmApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeProjectVmConnectGet**](ComputeProjectVmApi.md#ComputeProjectVmConnectGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/connect/{connectId} | Get compute/vm.connect
[**ComputeProjectVmConnectList**](ComputeProjectVmApi.md#ComputeProjectVmConnectList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/connect | List compute/vm.connect
[**ComputeProjectVmConnectOpen**](ComputeProjectVmApi.md#ComputeProjectVmConnectOpen) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/connect/{connectId}/actions/open | Open compute/vm.connect
[**ComputeProjectVmCreate**](ComputeProjectVmApi.md#ComputeProjectVmCreate) | **Post** /compute/{locationId}/project/{projectId}/vm | Create compute/vm
[**ComputeProjectVmDelete**](ComputeProjectVmApi.md#ComputeProjectVmDelete) | **Delete** /compute/{locationId}/project/{projectId}/vm/{vmId} | Delete compute/vm
[**ComputeProjectVmDiskCreate**](ComputeProjectVmApi.md#ComputeProjectVmDiskCreate) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/disk | Create compute/vm.disk
[**ComputeProjectVmDiskList**](ComputeProjectVmApi.md#ComputeProjectVmDiskList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/disk | List compute/vm.disk
[**ComputeProjectVmEventGet**](ComputeProjectVmApi.md#ComputeProjectVmEventGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/event/{eventId} | Get compute/vm.event
[**ComputeProjectVmEventList**](ComputeProjectVmApi.md#ComputeProjectVmEventList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/event | List compute/vm.event
[**ComputeProjectVmFlavour**](ComputeProjectVmApi.md#ComputeProjectVmFlavour) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/flavour | Flavour compute/vm
[**ComputeProjectVmGet**](ComputeProjectVmApi.md#ComputeProjectVmGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId} | Get compute/vm
[**ComputeProjectVmIsoCreate**](ComputeProjectVmApi.md#ComputeProjectVmIsoCreate) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/iso | Create compute/vm.iso
[**ComputeProjectVmIsoList**](ComputeProjectVmApi.md#ComputeProjectVmIsoList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/iso | List compute/vm.iso
[**ComputeProjectVmList**](ComputeProjectVmApi.md#ComputeProjectVmList) | **Get** /compute/{locationId}/project/{projectId}/vm | List compute/vm
[**ComputeProjectVmMetricGet**](ComputeProjectVmApi.md#ComputeProjectVmMetricGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/metric/{metricId} | Get compute/vm.metric
[**ComputeProjectVmMetricList**](ComputeProjectVmApi.md#ComputeProjectVmMetricList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/metric | List compute/vm.metric
[**ComputeProjectVmMetricPointList**](ComputeProjectVmApi.md#ComputeProjectVmMetricPointList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/metric/{metricId}/point | List compute/vm.point
[**ComputeProjectVmPasswordReset**](ComputeProjectVmApi.md#ComputeProjectVmPasswordReset) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/password_reset | Password reset compute/vm
[**ComputeProjectVmRestart**](ComputeProjectVmApi.md#ComputeProjectVmRestart) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/restart | Restart compute/vm
[**ComputeProjectVmSerialport**](ComputeProjectVmApi.md#ComputeProjectVmSerialport) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/serialport | Serialport compute/vm
[**ComputeProjectVmServiceGet**](ComputeProjectVmApi.md#ComputeProjectVmServiceGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/service/{serviceId} | Get compute/vm.service
[**ComputeProjectVmServiceList**](ComputeProjectVmApi.md#ComputeProjectVmServiceList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/service | List compute/vm.service
[**ComputeProjectVmStart**](ComputeProjectVmApi.md#ComputeProjectVmStart) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/start | Start compute/vm
[**ComputeProjectVmStop**](ComputeProjectVmApi.md#ComputeProjectVmStop) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/stop | Stop compute/vm
[**ComputeProjectVmTagCreate**](ComputeProjectVmApi.md#ComputeProjectVmTagCreate) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag | Create compute/vm.tag
[**ComputeProjectVmTagDelete**](ComputeProjectVmApi.md#ComputeProjectVmTagDelete) | **Delete** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag/{tagId} | Delete compute/vm.tag
[**ComputeProjectVmTagGet**](ComputeProjectVmApi.md#ComputeProjectVmTagGet) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag/{tagId} | Get compute/vm.tag
[**ComputeProjectVmTagList**](ComputeProjectVmApi.md#ComputeProjectVmTagList) | **Get** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag | List compute/vm.tag
[**ComputeProjectVmTagPut**](ComputeProjectVmApi.md#ComputeProjectVmTagPut) | **Put** /compute/{locationId}/project/{projectId}/vm/{vmId}/tag | Replace compute/vm.tag
[**ComputeProjectVmTurnoff**](ComputeProjectVmApi.md#ComputeProjectVmTurnoff) | **Post** /compute/{locationId}/project/{projectId}/vm/{vmId}/actions/turnoff | Turnoff compute/vm
[**ComputeProjectVmUpdate**](ComputeProjectVmApi.md#ComputeProjectVmUpdate) | **Patch** /compute/{locationId}/project/{projectId}/vm/{vmId} | Update compute/vm



## ComputeProjectVmConnectGet

> Connect ComputeProjectVmConnectGet(ctx, projectId, locationId, vmId, connectId).Execute()

Get compute/vm.connect



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
    vmId := "vmId_example" // string | Vm Id
    connectId := "connectId_example" // string | connectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmConnectGet(context.Background(), projectId, locationId, vmId, connectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmConnectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmConnectGet`: Connect
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmConnectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 
**connectId** | **string** | connectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmConnectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Connect**](Connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmConnectList

> []Connect ComputeProjectVmConnectList(ctx, projectId, locationId, vmId).Execute()

List compute/vm.connect



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmConnectList(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmConnectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmConnectList`: []Connect
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmConnectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmConnectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Connect**](Connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmConnectOpen

> ComputeProjectVmConnectOpen(ctx, projectId, locationId, vmId, connectId).ComputeProjectVmConnectOpen(computeProjectVmConnectOpen).Execute()

Open compute/vm.connect



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
    vmId := "vmId_example" // string | Vm Id
    connectId := "connectId_example" // string | connectId
    computeProjectVmConnectOpen := *openapiclient.NewComputeProjectVmConnectOpen() // ComputeProjectVmConnectOpen | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmConnectOpen(context.Background(), projectId, locationId, vmId, connectId).ComputeProjectVmConnectOpen(computeProjectVmConnectOpen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmConnectOpen``: %v\n", err)
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
**vmId** | **string** | Vm Id | 
**connectId** | **string** | connectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmConnectOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **computeProjectVmConnectOpen** | [**ComputeProjectVmConnectOpen**](ComputeProjectVmConnectOpen.md) |  | 

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


## ComputeProjectVmCreate

> Vm ComputeProjectVmCreate(ctx, projectId, locationId).ComputeProjectVmCreate(computeProjectVmCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create compute/vm



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
    computeProjectVmCreate := *openapiclient.NewComputeProjectVmCreate("Name_example", "Service_example") // ComputeProjectVmCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmCreate(context.Background(), projectId, locationId).ComputeProjectVmCreate(computeProjectVmCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmCreate`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **computeProjectVmCreate** | [**ComputeProjectVmCreate**](ComputeProjectVmCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmDelete

> ComputeProjectVmDelete(ctx, projectId, locationId, vmId).Execute()

Delete compute/vm



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmDelete(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmDelete``: %v\n", err)
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
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmDeleteRequest struct via the builder pattern


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


## ComputeProjectVmDiskCreate

> Disk ComputeProjectVmDiskCreate(ctx, projectId, locationId, vmId).ComputeProjectVmDiskCreate(computeProjectVmDiskCreate).Execute()

Create compute/vm.disk



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
    vmId := "vmId_example" // string | Vm Id
    computeProjectVmDiskCreate := *openapiclient.NewComputeProjectVmDiskCreate("Disk_example") // ComputeProjectVmDiskCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmDiskCreate(context.Background(), projectId, locationId, vmId).ComputeProjectVmDiskCreate(computeProjectVmDiskCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmDiskCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmDiskCreate`: Disk
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmDiskCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmDiskCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **computeProjectVmDiskCreate** | [**ComputeProjectVmDiskCreate**](ComputeProjectVmDiskCreate.md) |  | 

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


## ComputeProjectVmDiskList

> []Disk ComputeProjectVmDiskList(ctx, projectId, locationId, vmId).Execute()

List compute/vm.disk



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmDiskList(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmDiskList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmDiskList`: []Disk
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmDiskList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmDiskListRequest struct via the builder pattern


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


## ComputeProjectVmEventGet

> Event ComputeProjectVmEventGet(ctx, projectId, locationId, vmId, eventId).Execute()

Get compute/vm.event



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
    vmId := "vmId_example" // string | Vm Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmEventGet(context.Background(), projectId, locationId, vmId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmEventGetRequest struct via the builder pattern


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


## ComputeProjectVmEventList

> []Event ComputeProjectVmEventList(ctx, projectId, locationId, vmId).Limit(limit).Skip(skip).Execute()

List compute/vm.event



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
    vmId := "vmId_example" // string | Vm Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmEventList(context.Background(), projectId, locationId, vmId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmEventListRequest struct via the builder pattern


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


## ComputeProjectVmFlavour

> Vm ComputeProjectVmFlavour(ctx, projectId, locationId, vmId).ComputeProjectVmFlavour(computeProjectVmFlavour).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Flavour compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    computeProjectVmFlavour := *openapiclient.NewComputeProjectVmFlavour("Service_example") // ComputeProjectVmFlavour | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmFlavour(context.Background(), projectId, locationId, vmId).ComputeProjectVmFlavour(computeProjectVmFlavour).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmFlavour``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmFlavour`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmFlavour`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmFlavourRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **computeProjectVmFlavour** | [**ComputeProjectVmFlavour**](ComputeProjectVmFlavour.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmGet

> Vm ComputeProjectVmGet(ctx, projectId, locationId, vmId).Execute()

Get compute/vm



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmGet(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmGet`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmIsoCreate

> Iso ComputeProjectVmIsoCreate(ctx, projectId, locationId, vmId).ComputeProjectVmIsoCreate(computeProjectVmIsoCreate).Execute()

Create compute/vm.iso



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
    vmId := "vmId_example" // string | Vm Id
    computeProjectVmIsoCreate := *openapiclient.NewComputeProjectVmIsoCreate() // ComputeProjectVmIsoCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmIsoCreate(context.Background(), projectId, locationId, vmId).ComputeProjectVmIsoCreate(computeProjectVmIsoCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmIsoCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmIsoCreate`: Iso
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmIsoCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmIsoCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **computeProjectVmIsoCreate** | [**ComputeProjectVmIsoCreate**](ComputeProjectVmIsoCreate.md) |  | 

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


## ComputeProjectVmIsoList

> []Iso ComputeProjectVmIsoList(ctx, projectId, locationId, vmId).Execute()

List compute/vm.iso



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmIsoList(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmIsoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmIsoList`: []Iso
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmIsoList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmIsoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## ComputeProjectVmList

> []Vm ComputeProjectVmList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List compute/vm



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
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmList`: []Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmMetricGet

> Metric ComputeProjectVmMetricGet(ctx, projectId, locationId, vmId, metricId).Execute()

Get compute/vm.metric



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
    vmId := "vmId_example" // string | Vm Id
    metricId := "metricId_example" // string | metricId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmMetricGet(context.Background(), projectId, locationId, vmId, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmMetricGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmMetricGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmMetricGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmMetricGetRequest struct via the builder pattern


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


## ComputeProjectVmMetricList

> []Metric ComputeProjectVmMetricList(ctx, projectId, locationId, vmId).Execute()

List compute/vm.metric



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmMetricList(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmMetricList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmMetricList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmMetricList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmMetricListRequest struct via the builder pattern


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


## ComputeProjectVmMetricPointList

> []Point ComputeProjectVmMetricPointList(ctx, projectId, locationId, vmId, metricId).Interval(interval).Timespan(timespan).Execute()

List compute/vm.point



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
    vmId := "vmId_example" // string | Vm Id
    metricId := "metricId_example" // string | metricId
    interval := "interval_example" // string | interval (optional)
    timespan := "timespan_example" // string | timespan (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmMetricPointList(context.Background(), projectId, locationId, vmId, metricId).Interval(interval).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmMetricPointList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmMetricPointList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmMetricPointList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmMetricPointListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **interval** | **string** | interval | 
 **timespan** | **string** | timespan | 

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


## ComputeProjectVmPasswordReset

> Vm ComputeProjectVmPasswordReset(ctx, projectId, locationId, vmId).ComputeProjectVmPasswordReset(computeProjectVmPasswordReset).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Password reset compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    computeProjectVmPasswordReset := *openapiclient.NewComputeProjectVmPasswordReset("UserName_example", "Modulus_example", "Exponent_example") // ComputeProjectVmPasswordReset | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmPasswordReset(context.Background(), projectId, locationId, vmId).ComputeProjectVmPasswordReset(computeProjectVmPasswordReset).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmPasswordReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmPasswordReset`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmPasswordReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **computeProjectVmPasswordReset** | [**ComputeProjectVmPasswordReset**](ComputeProjectVmPasswordReset.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmRestart

> Vm ComputeProjectVmRestart(ctx, projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Restart compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmRestart(context.Background(), projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmRestart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmRestart`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmSerialport

> *os.File ComputeProjectVmSerialport(ctx, projectId, locationId, vmId).ComputeProjectVmSerialport(computeProjectVmSerialport).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Serialport compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    computeProjectVmSerialport := *openapiclient.NewComputeProjectVmSerialport() // ComputeProjectVmSerialport | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmSerialport(context.Background(), projectId, locationId, vmId).ComputeProjectVmSerialport(computeProjectVmSerialport).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmSerialport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmSerialport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmSerialport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmSerialportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **computeProjectVmSerialport** | [**ComputeProjectVmSerialport**](ComputeProjectVmSerialport.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmServiceGet

> ResourceService ComputeProjectVmServiceGet(ctx, projectId, locationId, vmId, serviceId).Execute()

Get compute/vm.service



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
    vmId := "vmId_example" // string | Vm Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmServiceGet(context.Background(), projectId, locationId, vmId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmServiceGetRequest struct via the builder pattern


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


## ComputeProjectVmServiceList

> []ResourceService ComputeProjectVmServiceList(ctx, projectId, locationId, vmId).Execute()

List compute/vm.service



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmServiceList(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmServiceListRequest struct via the builder pattern


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


## ComputeProjectVmStart

> Vm ComputeProjectVmStart(ctx, projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Start compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmStart(context.Background(), projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmStart`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmStop

> Vm ComputeProjectVmStop(ctx, projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Stop compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmStop(context.Background(), projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmStop`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmTagCreate

> Tag ComputeProjectVmTagCreate(ctx, projectId, locationId, vmId).Tag(tag).Execute()

Create compute/vm.tag



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
    vmId := "vmId_example" // string | Vm Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmTagCreate(context.Background(), projectId, locationId, vmId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmTagCreateRequest struct via the builder pattern


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


## ComputeProjectVmTagDelete

> ComputeProjectVmTagDelete(ctx, projectId, locationId, vmId, tagId).Execute()

Delete compute/vm.tag



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
    vmId := "vmId_example" // string | Vm Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmTagDelete(context.Background(), projectId, locationId, vmId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmTagDelete``: %v\n", err)
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
**vmId** | **string** | Vm Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmTagDeleteRequest struct via the builder pattern


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


## ComputeProjectVmTagGet

> Tag ComputeProjectVmTagGet(ctx, projectId, locationId, vmId, tagId).Execute()

Get compute/vm.tag



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
    vmId := "vmId_example" // string | Vm Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmTagGet(context.Background(), projectId, locationId, vmId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmTagGetRequest struct via the builder pattern


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


## ComputeProjectVmTagList

> []Tag ComputeProjectVmTagList(ctx, projectId, locationId, vmId).Execute()

List compute/vm.tag



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
    vmId := "vmId_example" // string | Vm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmTagList(context.Background(), projectId, locationId, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmTagListRequest struct via the builder pattern


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


## ComputeProjectVmTagPut

> []Tag ComputeProjectVmTagPut(ctx, projectId, locationId, vmId).Tag(tag).Execute()

Replace compute/vm.tag



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
    vmId := "vmId_example" // string | Vm Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmTagPut(context.Background(), projectId, locationId, vmId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmTagPutRequest struct via the builder pattern


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


## ComputeProjectVmTurnoff

> Vm ComputeProjectVmTurnoff(ctx, projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Turnoff compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmTurnoff(context.Background(), projectId, locationId, vmId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmTurnoff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmTurnoff`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmTurnoff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmTurnoffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectVmUpdate

> Vm ComputeProjectVmUpdate(ctx, projectId, locationId, vmId).ComputeProjectVmUpdate(computeProjectVmUpdate).Execute()

Update compute/vm



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
    vmId := "vmId_example" // string | Vm Id
    computeProjectVmUpdate := *openapiclient.NewComputeProjectVmUpdate() // ComputeProjectVmUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectVmApi.ComputeProjectVmUpdate(context.Background(), projectId, locationId, vmId).ComputeProjectVmUpdate(computeProjectVmUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectVmApi.ComputeProjectVmUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectVmUpdate`: Vm
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectVmApi.ComputeProjectVmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**vmId** | **string** | Vm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectVmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **computeProjectVmUpdate** | [**ComputeProjectVmUpdate**](ComputeProjectVmUpdate.md) |  | 

### Return type

[**Vm**](Vm.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

