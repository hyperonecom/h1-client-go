# \ComputeProjectReplicaApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeProjectReplicaCreate**](ComputeProjectReplicaApi.md#ComputeProjectReplicaCreate) | **Post** /compute/{locationId}/project/{projectId}/replica | Create compute/replica
[**ComputeProjectReplicaDelete**](ComputeProjectReplicaApi.md#ComputeProjectReplicaDelete) | **Delete** /compute/{locationId}/project/{projectId}/replica/{replicaId} | Delete compute/replica
[**ComputeProjectReplicaEventGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaEventGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/event/{eventId} | Get compute/replica.event
[**ComputeProjectReplicaEventList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaEventList) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/event | List compute/replica.event
[**ComputeProjectReplicaGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId} | Get compute/replica
[**ComputeProjectReplicaList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaList) | **Get** /compute/{locationId}/project/{projectId}/replica | List compute/replica
[**ComputeProjectReplicaServiceGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaServiceGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/service/{serviceId} | Get compute/replica.service
[**ComputeProjectReplicaServiceList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaServiceList) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/service | List compute/replica.service
[**ComputeProjectReplicaTagCreate**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagCreate) | **Post** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag | Create compute/replica.tag
[**ComputeProjectReplicaTagDelete**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagDelete) | **Delete** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag/{tagId} | Delete compute/replica.tag
[**ComputeProjectReplicaTagGet**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagGet) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag/{tagId} | Get compute/replica.tag
[**ComputeProjectReplicaTagList**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagList) | **Get** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag | List compute/replica.tag
[**ComputeProjectReplicaTagPut**](ComputeProjectReplicaApi.md#ComputeProjectReplicaTagPut) | **Put** /compute/{locationId}/project/{projectId}/replica/{replicaId}/tag | Replace compute/replica.tag



## ComputeProjectReplicaCreate

> Replica ComputeProjectReplicaCreate(ctx, projectId, locationId).ComputeProjectReplicaCreate(computeProjectReplicaCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create compute/replica



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
    computeProjectReplicaCreate := *openapiclient.NewComputeProjectReplicaCreate("Hostname_example", "Secret_example") // ComputeProjectReplicaCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaCreate(context.Background(), projectId, locationId).ComputeProjectReplicaCreate(computeProjectReplicaCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaCreate`: Replica
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **computeProjectReplicaCreate** | [**ComputeProjectReplicaCreate**](ComputeProjectReplicaCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Replica**](Replica.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectReplicaDelete

> ComputeProjectReplicaDelete(ctx, projectId, locationId, replicaId).Execute()

Delete compute/replica



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
    replicaId := "replicaId_example" // string | Replica Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaDelete(context.Background(), projectId, locationId, replicaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaDelete``: %v\n", err)
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
**replicaId** | **string** | Replica Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaDeleteRequest struct via the builder pattern


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


## ComputeProjectReplicaEventGet

> Event ComputeProjectReplicaEventGet(ctx, projectId, locationId, replicaId, eventId).Execute()

Get compute/replica.event



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
    replicaId := "replicaId_example" // string | Replica Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaEventGet(context.Background(), projectId, locationId, replicaId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaEventGetRequest struct via the builder pattern


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


## ComputeProjectReplicaEventList

> []Event ComputeProjectReplicaEventList(ctx, projectId, locationId, replicaId).Limit(limit).Skip(skip).Execute()

List compute/replica.event



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
    replicaId := "replicaId_example" // string | Replica Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaEventList(context.Background(), projectId, locationId, replicaId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaEventListRequest struct via the builder pattern


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


## ComputeProjectReplicaGet

> Replica ComputeProjectReplicaGet(ctx, projectId, locationId, replicaId).Execute()

Get compute/replica



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
    replicaId := "replicaId_example" // string | Replica Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaGet(context.Background(), projectId, locationId, replicaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaGet`: Replica
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Replica**](Replica.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectReplicaList

> []Replica ComputeProjectReplicaList(ctx, projectId, locationId).Name(name).Execute()

List compute/replica



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaList(context.Background(), projectId, locationId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaList`: []Replica
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 

### Return type

[**[]Replica**](Replica.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeProjectReplicaServiceGet

> ResourceService ComputeProjectReplicaServiceGet(ctx, projectId, locationId, replicaId, serviceId).Execute()

Get compute/replica.service



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
    replicaId := "replicaId_example" // string | Replica Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaServiceGet(context.Background(), projectId, locationId, replicaId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaServiceGetRequest struct via the builder pattern


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


## ComputeProjectReplicaServiceList

> []ResourceService ComputeProjectReplicaServiceList(ctx, projectId, locationId, replicaId).Execute()

List compute/replica.service



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
    replicaId := "replicaId_example" // string | Replica Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaServiceList(context.Background(), projectId, locationId, replicaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaServiceListRequest struct via the builder pattern


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


## ComputeProjectReplicaTagCreate

> Tag ComputeProjectReplicaTagCreate(ctx, projectId, locationId, replicaId).Tag(tag).Execute()

Create compute/replica.tag



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
    replicaId := "replicaId_example" // string | Replica Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaTagCreate(context.Background(), projectId, locationId, replicaId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaTagCreateRequest struct via the builder pattern


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


## ComputeProjectReplicaTagDelete

> ComputeProjectReplicaTagDelete(ctx, projectId, locationId, replicaId, tagId).Execute()

Delete compute/replica.tag



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
    replicaId := "replicaId_example" // string | Replica Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaTagDelete(context.Background(), projectId, locationId, replicaId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaTagDelete``: %v\n", err)
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
**replicaId** | **string** | Replica Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaTagDeleteRequest struct via the builder pattern


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


## ComputeProjectReplicaTagGet

> Tag ComputeProjectReplicaTagGet(ctx, projectId, locationId, replicaId, tagId).Execute()

Get compute/replica.tag



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
    replicaId := "replicaId_example" // string | Replica Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaTagGet(context.Background(), projectId, locationId, replicaId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaTagGetRequest struct via the builder pattern


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


## ComputeProjectReplicaTagList

> []Tag ComputeProjectReplicaTagList(ctx, projectId, locationId, replicaId).Execute()

List compute/replica.tag



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
    replicaId := "replicaId_example" // string | Replica Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaTagList(context.Background(), projectId, locationId, replicaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaTagListRequest struct via the builder pattern


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


## ComputeProjectReplicaTagPut

> []Tag ComputeProjectReplicaTagPut(ctx, projectId, locationId, replicaId).Tag(tag).Execute()

Replace compute/replica.tag



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
    replicaId := "replicaId_example" // string | Replica Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputeProjectReplicaApi.ComputeProjectReplicaTagPut(context.Background(), projectId, locationId, replicaId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputeProjectReplicaApi.ComputeProjectReplicaTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComputeProjectReplicaTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ComputeProjectReplicaApi.ComputeProjectReplicaTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**replicaId** | **string** | Replica Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeProjectReplicaTagPutRequest struct via the builder pattern


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

