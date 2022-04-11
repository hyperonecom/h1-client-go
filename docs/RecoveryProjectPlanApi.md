# \RecoveryProjectPlanApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecoveryProjectPlanCreate**](RecoveryProjectPlanApi.md#RecoveryProjectPlanCreate) | **Post** /recovery/{locationId}/project/{projectId}/plan | Create recovery/plan
[**RecoveryProjectPlanDelete**](RecoveryProjectPlanApi.md#RecoveryProjectPlanDelete) | **Delete** /recovery/{locationId}/project/{projectId}/plan/{planId} | Delete recovery/plan
[**RecoveryProjectPlanEventGet**](RecoveryProjectPlanApi.md#RecoveryProjectPlanEventGet) | **Get** /recovery/{locationId}/project/{projectId}/plan/{planId}/event/{eventId} | Get recovery/plan.event
[**RecoveryProjectPlanEventList**](RecoveryProjectPlanApi.md#RecoveryProjectPlanEventList) | **Get** /recovery/{locationId}/project/{projectId}/plan/{planId}/event | List recovery/plan.event
[**RecoveryProjectPlanGet**](RecoveryProjectPlanApi.md#RecoveryProjectPlanGet) | **Get** /recovery/{locationId}/project/{projectId}/plan/{planId} | Get recovery/plan
[**RecoveryProjectPlanList**](RecoveryProjectPlanApi.md#RecoveryProjectPlanList) | **Get** /recovery/{locationId}/project/{projectId}/plan | List recovery/plan
[**RecoveryProjectPlanServiceGet**](RecoveryProjectPlanApi.md#RecoveryProjectPlanServiceGet) | **Get** /recovery/{locationId}/project/{projectId}/plan/{planId}/service/{serviceId} | Get recovery/plan.service
[**RecoveryProjectPlanServiceList**](RecoveryProjectPlanApi.md#RecoveryProjectPlanServiceList) | **Get** /recovery/{locationId}/project/{projectId}/plan/{planId}/service | List recovery/plan.service
[**RecoveryProjectPlanTagCreate**](RecoveryProjectPlanApi.md#RecoveryProjectPlanTagCreate) | **Post** /recovery/{locationId}/project/{projectId}/plan/{planId}/tag | Create recovery/plan.tag
[**RecoveryProjectPlanTagDelete**](RecoveryProjectPlanApi.md#RecoveryProjectPlanTagDelete) | **Delete** /recovery/{locationId}/project/{projectId}/plan/{planId}/tag/{tagId} | Delete recovery/plan.tag
[**RecoveryProjectPlanTagGet**](RecoveryProjectPlanApi.md#RecoveryProjectPlanTagGet) | **Get** /recovery/{locationId}/project/{projectId}/plan/{planId}/tag/{tagId} | Get recovery/plan.tag
[**RecoveryProjectPlanTagList**](RecoveryProjectPlanApi.md#RecoveryProjectPlanTagList) | **Get** /recovery/{locationId}/project/{projectId}/plan/{planId}/tag | List recovery/plan.tag
[**RecoveryProjectPlanTagPut**](RecoveryProjectPlanApi.md#RecoveryProjectPlanTagPut) | **Put** /recovery/{locationId}/project/{projectId}/plan/{planId}/tag | Replace recovery/plan.tag
[**RecoveryProjectPlanUpdate**](RecoveryProjectPlanApi.md#RecoveryProjectPlanUpdate) | **Patch** /recovery/{locationId}/project/{projectId}/plan/{planId} | Update recovery/plan



## RecoveryProjectPlanCreate

> Plan RecoveryProjectPlanCreate(ctx, projectId, locationId).RecoveryProjectPlanCreate(recoveryProjectPlanCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create recovery/plan



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
    recoveryProjectPlanCreate := *openapiclient.NewRecoveryProjectPlanCreate("Name_example") // RecoveryProjectPlanCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanCreate(context.Background(), projectId, locationId).RecoveryProjectPlanCreate(recoveryProjectPlanCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanCreate`: Plan
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recoveryProjectPlanCreate** | [**RecoveryProjectPlanCreate**](RecoveryProjectPlanCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Plan**](Plan.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectPlanDelete

> RecoveryProjectPlanDelete(ctx, projectId, locationId, planId).Execute()

Delete recovery/plan



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
    planId := "planId_example" // string | Plan Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanDelete(context.Background(), projectId, locationId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanDelete``: %v\n", err)
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
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanDeleteRequest struct via the builder pattern


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


## RecoveryProjectPlanEventGet

> Event RecoveryProjectPlanEventGet(ctx, projectId, locationId, planId, eventId).Execute()

Get recovery/plan.event



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
    planId := "planId_example" // string | Plan Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanEventGet(context.Background(), projectId, locationId, planId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanEventGetRequest struct via the builder pattern


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


## RecoveryProjectPlanEventList

> []Event RecoveryProjectPlanEventList(ctx, projectId, locationId, planId).Limit(limit).Skip(skip).Execute()

List recovery/plan.event



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
    planId := "planId_example" // string | Plan Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanEventList(context.Background(), projectId, locationId, planId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanEventListRequest struct via the builder pattern


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


## RecoveryProjectPlanGet

> Plan RecoveryProjectPlanGet(ctx, projectId, locationId, planId).Execute()

Get recovery/plan



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
    planId := "planId_example" // string | Plan Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanGet(context.Background(), projectId, locationId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanGet`: Plan
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Plan**](Plan.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectPlanList

> []Plan RecoveryProjectPlanList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List recovery/plan



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
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanList`: []Plan
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Plan**](Plan.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectPlanServiceGet

> ResourceService RecoveryProjectPlanServiceGet(ctx, projectId, locationId, planId, serviceId).Execute()

Get recovery/plan.service



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
    planId := "planId_example" // string | Plan Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanServiceGet(context.Background(), projectId, locationId, planId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanServiceGetRequest struct via the builder pattern


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


## RecoveryProjectPlanServiceList

> []ResourceService RecoveryProjectPlanServiceList(ctx, projectId, locationId, planId).Execute()

List recovery/plan.service



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
    planId := "planId_example" // string | Plan Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanServiceList(context.Background(), projectId, locationId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanServiceListRequest struct via the builder pattern


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


## RecoveryProjectPlanTagCreate

> Tag RecoveryProjectPlanTagCreate(ctx, projectId, locationId, planId).Tag(tag).Execute()

Create recovery/plan.tag



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
    planId := "planId_example" // string | Plan Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanTagCreate(context.Background(), projectId, locationId, planId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanTagCreateRequest struct via the builder pattern


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


## RecoveryProjectPlanTagDelete

> RecoveryProjectPlanTagDelete(ctx, projectId, locationId, planId, tagId).Execute()

Delete recovery/plan.tag



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
    planId := "planId_example" // string | Plan Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanTagDelete(context.Background(), projectId, locationId, planId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanTagDelete``: %v\n", err)
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
**planId** | **string** | Plan Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanTagDeleteRequest struct via the builder pattern


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


## RecoveryProjectPlanTagGet

> Tag RecoveryProjectPlanTagGet(ctx, projectId, locationId, planId, tagId).Execute()

Get recovery/plan.tag



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
    planId := "planId_example" // string | Plan Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanTagGet(context.Background(), projectId, locationId, planId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanTagGetRequest struct via the builder pattern


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


## RecoveryProjectPlanTagList

> []Tag RecoveryProjectPlanTagList(ctx, projectId, locationId, planId).Execute()

List recovery/plan.tag



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
    planId := "planId_example" // string | Plan Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanTagList(context.Background(), projectId, locationId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanTagListRequest struct via the builder pattern


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


## RecoveryProjectPlanTagPut

> []Tag RecoveryProjectPlanTagPut(ctx, projectId, locationId, planId).Tag(tag).Execute()

Replace recovery/plan.tag



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
    planId := "planId_example" // string | Plan Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanTagPut(context.Background(), projectId, locationId, planId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanTagPutRequest struct via the builder pattern


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


## RecoveryProjectPlanUpdate

> Plan RecoveryProjectPlanUpdate(ctx, projectId, locationId, planId).RecoveryProjectPlanUpdate(recoveryProjectPlanUpdate).Execute()

Update recovery/plan



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
    planId := "planId_example" // string | Plan Id
    recoveryProjectPlanUpdate := *openapiclient.NewRecoveryProjectPlanUpdate() // RecoveryProjectPlanUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecoveryProjectPlanApi.RecoveryProjectPlanUpdate(context.Background(), projectId, locationId, planId).RecoveryProjectPlanUpdate(recoveryProjectPlanUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectPlanApi.RecoveryProjectPlanUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectPlanUpdate`: Plan
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectPlanApi.RecoveryProjectPlanUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**planId** | **string** | Plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectPlanUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **recoveryProjectPlanUpdate** | [**RecoveryProjectPlanUpdate**](RecoveryProjectPlanUpdate.md) |  | 

### Return type

[**Plan**](Plan.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

