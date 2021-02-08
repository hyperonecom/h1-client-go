# \RecoveryProjectBackupApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecoveryProjectBackupCreate**](RecoveryProjectBackupApi.md#RecoveryProjectBackupCreate) | **Post** /recovery/{locationId}/project/{projectId}/backup | Create recovery/backup
[**RecoveryProjectBackupDelete**](RecoveryProjectBackupApi.md#RecoveryProjectBackupDelete) | **Delete** /recovery/{locationId}/project/{projectId}/backup/{backupId} | Delete recovery/backup
[**RecoveryProjectBackupEventGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupEventGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/event/{eventId} | Get recovery/backup.event
[**RecoveryProjectBackupEventList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupEventList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/event | List recovery/backup.event
[**RecoveryProjectBackupExport**](RecoveryProjectBackupApi.md#RecoveryProjectBackupExport) | **Post** /recovery/{locationId}/project/{projectId}/backup/{backupId}/actions/export | Export recovery/backup
[**RecoveryProjectBackupGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId} | Get recovery/backup
[**RecoveryProjectBackupList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupList) | **Get** /recovery/{locationId}/project/{projectId}/backup | List recovery/backup
[**RecoveryProjectBackupMetricGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupMetricGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/metric/{metricId} | Get recovery/backup.metric
[**RecoveryProjectBackupMetricList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupMetricList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/metric | List recovery/backup.metric
[**RecoveryProjectBackupMetricPointList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupMetricPointList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/metric/{metricId}/point | List recovery/backup.point
[**RecoveryProjectBackupServiceGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupServiceGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/service/{serviceId} | Get recovery/backup.service
[**RecoveryProjectBackupServiceList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupServiceList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/service | List recovery/backup.service
[**RecoveryProjectBackupTagCreate**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagCreate) | **Post** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag | Create recovery/backup.tag
[**RecoveryProjectBackupTagDelete**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagDelete) | **Delete** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag/{tagId} | Delete recovery/backup.tag
[**RecoveryProjectBackupTagGet**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagGet) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag/{tagId} | Get recovery/backup.tag
[**RecoveryProjectBackupTagList**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagList) | **Get** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag | List recovery/backup.tag
[**RecoveryProjectBackupTagPut**](RecoveryProjectBackupApi.md#RecoveryProjectBackupTagPut) | **Put** /recovery/{locationId}/project/{projectId}/backup/{backupId}/tag | Replace recovery/backup.tag
[**RecoveryProjectBackupUpdate**](RecoveryProjectBackupApi.md#RecoveryProjectBackupUpdate) | **Patch** /recovery/{locationId}/project/{projectId}/backup/{backupId} | Update recovery/backup



## RecoveryProjectBackupCreate

> Backup RecoveryProjectBackupCreate(ctx, projectId, locationId).RecoveryProjectBackupCreate(recoveryProjectBackupCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create recovery/backup



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
    recoveryProjectBackupCreate := *openapiclient.NewRecoveryProjectBackupCreate("Name_example", "TODO") // RecoveryProjectBackupCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupCreate(context.Background(), projectId, locationId).RecoveryProjectBackupCreate(recoveryProjectBackupCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupCreate`: Backup
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recoveryProjectBackupCreate** | [**RecoveryProjectBackupCreate**](RecoveryProjectBackupCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupDelete

> RecoveryProjectBackupDelete(ctx, projectId, locationId, backupId).Execute()

Delete recovery/backup



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
    backupId := "backupId_example" // string | Backup Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupDelete(context.Background(), projectId, locationId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupDelete``: %v\n", err)
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
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupDeleteRequest struct via the builder pattern


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


## RecoveryProjectBackupEventGet

> Event RecoveryProjectBackupEventGet(ctx, projectId, locationId, backupId, eventId).Execute()

Get recovery/backup.event



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
    backupId := "backupId_example" // string | Backup Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupEventGet(context.Background(), projectId, locationId, backupId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupEventGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupEventList

> []Event RecoveryProjectBackupEventList(ctx, projectId, locationId, backupId).Limit(limit).Skip(skip).Execute()

List recovery/backup.event



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
    backupId := "backupId_example" // string | Backup Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupEventList(context.Background(), projectId, locationId, backupId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **float32** | $limit | [default to 100]
 **skip** | **float32** | $skip | 

### Return type

[**[]Event**](event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupExport

> Backup RecoveryProjectBackupExport(ctx, projectId, locationId, backupId).RecoveryProjectBackupExport(recoveryProjectBackupExport).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Export recovery/backup



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
    backupId := "backupId_example" // string | Backup Id
    recoveryProjectBackupExport := *openapiclient.NewRecoveryProjectBackupExport("Bucket_example") // RecoveryProjectBackupExport | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupExport(context.Background(), projectId, locationId, backupId).RecoveryProjectBackupExport(recoveryProjectBackupExport).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupExport`: Backup
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **recoveryProjectBackupExport** | [**RecoveryProjectBackupExport**](RecoveryProjectBackupExport.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupGet

> Backup RecoveryProjectBackupGet(ctx, projectId, locationId, backupId).Execute()

Get recovery/backup



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
    backupId := "backupId_example" // string | Backup Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupGet(context.Background(), projectId, locationId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupGet`: Backup
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupList

> []Backup RecoveryProjectBackupList(ctx, projectId, locationId).Name(name).Source(source).TagValue(tagValue).TagKey(tagKey).Execute()

List recovery/backup



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
    source := "source_example" // string | Filter by source (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupList(context.Background(), projectId, locationId).Name(name).Source(source).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupList`: []Backup
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **source** | **string** | Filter by source | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupMetricGet

> Metric RecoveryProjectBackupMetricGet(ctx, projectId, locationId, backupId, metricId).Execute()

Get recovery/backup.metric



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
    backupId := "backupId_example" // string | Backup Id
    metricId := "metricId_example" // string | metricId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupMetricGet(context.Background(), projectId, locationId, backupId, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupMetricGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupMetricGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupMetricGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupMetricGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Metric**](metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupMetricList

> []Metric RecoveryProjectBackupMetricList(ctx, projectId, locationId, backupId).Execute()

List recovery/backup.metric



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
    backupId := "backupId_example" // string | Backup Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupMetricList(context.Background(), projectId, locationId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupMetricList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupMetricList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupMetricList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupMetricListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Metric**](metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupMetricPointList

> []Point RecoveryProjectBackupMetricPointList(ctx, projectId, locationId, backupId, metricId).Interval(interval).Timespan(timespan).Execute()

List recovery/backup.point



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
    backupId := "backupId_example" // string | Backup Id
    metricId := "metricId_example" // string | metricId
    interval := "interval_example" // string | interval (optional)
    timespan := "timespan_example" // string | timespan (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupMetricPointList(context.Background(), projectId, locationId, backupId, metricId).Interval(interval).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupMetricPointList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupMetricPointList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupMetricPointList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupMetricPointListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **interval** | **string** | interval | 
 **timespan** | **string** | timespan | 

### Return type

[**[]Point**](point.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupServiceGet

> ResourceService RecoveryProjectBackupServiceGet(ctx, projectId, locationId, backupId, serviceId).Execute()

Get recovery/backup.service



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
    backupId := "backupId_example" // string | Backup Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupServiceGet(context.Background(), projectId, locationId, backupId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ResourceService**](resourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupServiceList

> []ResourceService RecoveryProjectBackupServiceList(ctx, projectId, locationId, backupId).Execute()

List recovery/backup.service



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
    backupId := "backupId_example" // string | Backup Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupServiceList(context.Background(), projectId, locationId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ResourceService**](resourceService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupTagCreate

> Tag RecoveryProjectBackupTagCreate(ctx, projectId, locationId, backupId).Tag(tag).Execute()

Create recovery/backup.tag



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
    backupId := "backupId_example" // string | Backup Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupTagCreate(context.Background(), projectId, locationId, backupId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupTagCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | [**Tag**](Tag.md) |  | 

### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupTagDelete

> RecoveryProjectBackupTagDelete(ctx, projectId, locationId, backupId, tagId).Execute()

Delete recovery/backup.tag



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
    backupId := "backupId_example" // string | Backup Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupTagDelete(context.Background(), projectId, locationId, backupId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupTagDelete``: %v\n", err)
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
**backupId** | **string** | Backup Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupTagDeleteRequest struct via the builder pattern


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


## RecoveryProjectBackupTagGet

> Tag RecoveryProjectBackupTagGet(ctx, projectId, locationId, backupId, tagId).Execute()

Get recovery/backup.tag



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
    backupId := "backupId_example" // string | Backup Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupTagGet(context.Background(), projectId, locationId, backupId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupTagGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupTagList

> []Tag RecoveryProjectBackupTagList(ctx, projectId, locationId, backupId).Execute()

List recovery/backup.tag



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
    backupId := "backupId_example" // string | Backup Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupTagList(context.Background(), projectId, locationId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupTagListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupTagPut

> []Tag RecoveryProjectBackupTagPut(ctx, projectId, locationId, backupId).Tag(tag).Execute()

Replace recovery/backup.tag



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
    backupId := "backupId_example" // string | Backup Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupTagPut(context.Background(), projectId, locationId, backupId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupTagPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | [**[]Tag**](tag.md) |  | 

### Return type

[**[]Tag**](tag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoveryProjectBackupUpdate

> Backup RecoveryProjectBackupUpdate(ctx, projectId, locationId, backupId).RecoveryProjectBackupUpdate(recoveryProjectBackupUpdate).Execute()

Update recovery/backup



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
    backupId := "backupId_example" // string | Backup Id
    recoveryProjectBackupUpdate := *openapiclient.NewRecoveryProjectBackupUpdate() // RecoveryProjectBackupUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoveryProjectBackupApi.RecoveryProjectBackupUpdate(context.Background(), projectId, locationId, backupId).RecoveryProjectBackupUpdate(recoveryProjectBackupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoveryProjectBackupApi.RecoveryProjectBackupUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoveryProjectBackupUpdate`: Backup
    fmt.Fprintf(os.Stdout, "Response from `RecoveryProjectBackupApi.RecoveryProjectBackupUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**backupId** | **string** | Backup Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoveryProjectBackupUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **recoveryProjectBackupUpdate** | [**RecoveryProjectBackupUpdate**](RecoveryProjectBackupUpdate.md) |  | 

### Return type

[**Backup**](backup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

