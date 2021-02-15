# \DatabaseProjectInstanceApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatabaseProjectInstanceConnectGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceConnectGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/connect/{connectId} | Get database/instance.connect
[**DatabaseProjectInstanceConnectList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceConnectList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/connect | List database/instance.connect
[**DatabaseProjectInstanceCreate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCreate) | **Post** /database/{locationId}/project/{projectId}/instance | Create database/instance
[**DatabaseProjectInstanceCredentialCreate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialCreate) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential | Create database/instance.credential
[**DatabaseProjectInstanceCredentialDelete**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialDelete) | **Delete** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential/{credentialId} | Delete database/instance.credential
[**DatabaseProjectInstanceCredentialGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential/{credentialId} | Get database/instance.credential
[**DatabaseProjectInstanceCredentialList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential | List database/instance.credential
[**DatabaseProjectInstanceCredentialPatch**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceCredentialPatch) | **Patch** /database/{locationId}/project/{projectId}/instance/{instanceId}/credential/{credentialId} | Update database/instance.credential
[**DatabaseProjectInstanceDelete**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceDelete) | **Delete** /database/{locationId}/project/{projectId}/instance/{instanceId} | Delete database/instance
[**DatabaseProjectInstanceEventGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceEventGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/event/{eventId} | Get database/instance.event
[**DatabaseProjectInstanceEventList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceEventList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/event | List database/instance.event
[**DatabaseProjectInstanceGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId} | Get database/instance
[**DatabaseProjectInstanceList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceList) | **Get** /database/{locationId}/project/{projectId}/instance | List database/instance
[**DatabaseProjectInstanceMetricGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceMetricGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/metric/{metricId} | Get database/instance.metric
[**DatabaseProjectInstanceMetricList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceMetricList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/metric | List database/instance.metric
[**DatabaseProjectInstanceMetricPointList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceMetricPointList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/metric/{metricId}/point | List database/instance.point
[**DatabaseProjectInstanceServiceGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceServiceGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/service/{serviceId} | Get database/instance.service
[**DatabaseProjectInstanceServiceList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceServiceList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/service | List database/instance.service
[**DatabaseProjectInstanceStart**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceStart) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/actions/start | Start database/instance
[**DatabaseProjectInstanceStop**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceStop) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/actions/stop | Stop database/instance
[**DatabaseProjectInstanceTagCreate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagCreate) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag | Create database/instance.tag
[**DatabaseProjectInstanceTagDelete**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagDelete) | **Delete** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Delete database/instance.tag
[**DatabaseProjectInstanceTagGet**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagGet) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Get database/instance.tag
[**DatabaseProjectInstanceTagList**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagList) | **Get** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag | List database/instance.tag
[**DatabaseProjectInstanceTagPut**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTagPut) | **Put** /database/{locationId}/project/{projectId}/instance/{instanceId}/tag | Replace database/instance.tag
[**DatabaseProjectInstanceTransfer**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceTransfer) | **Post** /database/{locationId}/project/{projectId}/instance/{instanceId}/actions/transfer | Transfer database/instance
[**DatabaseProjectInstanceUpdate**](DatabaseProjectInstanceApi.md#DatabaseProjectInstanceUpdate) | **Patch** /database/{locationId}/project/{projectId}/instance/{instanceId} | Update database/instance



## DatabaseProjectInstanceConnectGet

> ResourceConnect DatabaseProjectInstanceConnectGet(ctx, projectId, locationId, instanceId, connectId).Execute()

Get database/instance.connect



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
    instanceId := "instanceId_example" // string | Instance Id
    connectId := "connectId_example" // string | connectId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceConnectGet(context.Background(), projectId, locationId, instanceId, connectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceConnectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceConnectGet`: ResourceConnect
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceConnectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**connectId** | **string** | connectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceConnectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ResourceConnect**](resource.connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceConnectList

> []ResourceConnect DatabaseProjectInstanceConnectList(ctx, projectId, locationId, instanceId).Execute()

List database/instance.connect



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
    instanceId := "instanceId_example" // string | Instance Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceConnectList(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceConnectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceConnectList`: []ResourceConnect
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceConnectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceConnectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ResourceConnect**](resource.connect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCreate

> Database DatabaseProjectInstanceCreate(ctx, projectId, locationId).DatabaseProjectInstanceCreate(databaseProjectInstanceCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create database/instance



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
    databaseProjectInstanceCreate := *openapiclient.NewDatabaseProjectInstanceCreate("Name_example", "Service_example") // DatabaseProjectInstanceCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceCreate(context.Background(), projectId, locationId).DatabaseProjectInstanceCreate(databaseProjectInstanceCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceCreate`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **databaseProjectInstanceCreate** | [**DatabaseProjectInstanceCreate**](DatabaseProjectInstanceCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialCreate

> DatabaseCredential DatabaseProjectInstanceCredentialCreate(ctx, projectId, locationId, instanceId).DatabaseCredential(databaseCredential).Execute()

Create database/instance.credential



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
    instanceId := "instanceId_example" // string | Instance Id
    databaseCredential := *openapiclient.NewDatabaseCredential("Name_example", "Type_example", "Value_example") // DatabaseCredential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialCreate(context.Background(), projectId, locationId, instanceId).DatabaseCredential(databaseCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceCredentialCreate`: DatabaseCredential
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **databaseCredential** | [**DatabaseCredential**](DatabaseCredential.md) |  | 

### Return type

[**DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialDelete

> Database DatabaseProjectInstanceCredentialDelete(ctx, projectId, locationId, instanceId, credentialId).Execute()

Delete database/instance.credential



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
    instanceId := "instanceId_example" // string | Instance Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialDelete(context.Background(), projectId, locationId, instanceId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceCredentialDelete`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialGet

> DatabaseCredential DatabaseProjectInstanceCredentialGet(ctx, projectId, locationId, instanceId, credentialId).Execute()

Get database/instance.credential



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
    instanceId := "instanceId_example" // string | Instance Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialGet(context.Background(), projectId, locationId, instanceId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceCredentialGet`: DatabaseCredential
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialList

> []DatabaseCredential DatabaseProjectInstanceCredentialList(ctx, projectId, locationId, instanceId).Execute()

List database/instance.credential



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
    instanceId := "instanceId_example" // string | Instance Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialList(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceCredentialList`: []DatabaseCredential
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceCredentialPatch

> DatabaseCredential DatabaseProjectInstanceCredentialPatch(ctx, projectId, locationId, instanceId, credentialId).DatabaseProjectInstanceCredentialPatch(databaseProjectInstanceCredentialPatch).Execute()

Update database/instance.credential



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
    instanceId := "instanceId_example" // string | Instance Id
    credentialId := "credentialId_example" // string | credentialId
    databaseProjectInstanceCredentialPatch := *openapiclient.NewDatabaseProjectInstanceCredentialPatch("Name_example") // DatabaseProjectInstanceCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialPatch(context.Background(), projectId, locationId, instanceId, credentialId).DatabaseProjectInstanceCredentialPatch(databaseProjectInstanceCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceCredentialPatch`: DatabaseCredential
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **databaseProjectInstanceCredentialPatch** | [**DatabaseProjectInstanceCredentialPatch**](DatabaseProjectInstanceCredentialPatch.md) |  | 

### Return type

[**DatabaseCredential**](database.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceDelete

> DatabaseProjectInstanceDelete(ctx, projectId, locationId, instanceId).Execute()

Delete database/instance



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
    instanceId := "instanceId_example" // string | Instance Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceDelete(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceDelete``: %v\n", err)
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
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceDeleteRequest struct via the builder pattern


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


## DatabaseProjectInstanceEventGet

> Event DatabaseProjectInstanceEventGet(ctx, projectId, locationId, instanceId, eventId).Execute()

Get database/instance.event



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
    instanceId := "instanceId_example" // string | Instance Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceEventGet(context.Background(), projectId, locationId, instanceId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceEventGetRequest struct via the builder pattern


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


## DatabaseProjectInstanceEventList

> []Event DatabaseProjectInstanceEventList(ctx, projectId, locationId, instanceId).Limit(limit).Skip(skip).Execute()

List database/instance.event



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
    instanceId := "instanceId_example" // string | Instance Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceEventList(context.Background(), projectId, locationId, instanceId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceEventListRequest struct via the builder pattern


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


## DatabaseProjectInstanceGet

> Database DatabaseProjectInstanceGet(ctx, projectId, locationId, instanceId).Execute()

Get database/instance



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
    instanceId := "instanceId_example" // string | Instance Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceGet(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceGet`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceList

> []Database DatabaseProjectInstanceList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List database/instance



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceList`: []Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceMetricGet

> Metric DatabaseProjectInstanceMetricGet(ctx, projectId, locationId, instanceId, metricId).Execute()

Get database/instance.metric



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
    instanceId := "instanceId_example" // string | Instance Id
    metricId := "metricId_example" // string | metricId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricGet(context.Background(), projectId, locationId, instanceId, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceMetricGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceMetricGetRequest struct via the builder pattern


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


## DatabaseProjectInstanceMetricList

> []Metric DatabaseProjectInstanceMetricList(ctx, projectId, locationId, instanceId).Execute()

List database/instance.metric



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
    instanceId := "instanceId_example" // string | Instance Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricList(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceMetricList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceMetricListRequest struct via the builder pattern


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


## DatabaseProjectInstanceMetricPointList

> []Point DatabaseProjectInstanceMetricPointList(ctx, projectId, locationId, instanceId, metricId).Interval(interval).Timespan(timespan).Execute()

List database/instance.point



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
    instanceId := "instanceId_example" // string | Instance Id
    metricId := "metricId_example" // string | metricId
    interval := "interval_example" // string | interval (optional)
    timespan := "timespan_example" // string | timespan (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricPointList(context.Background(), projectId, locationId, instanceId, metricId).Interval(interval).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricPointList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceMetricPointList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceMetricPointList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceMetricPointListRequest struct via the builder pattern


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


## DatabaseProjectInstanceServiceGet

> ResourceService DatabaseProjectInstanceServiceGet(ctx, projectId, locationId, instanceId, serviceId).Execute()

Get database/instance.service



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
    instanceId := "instanceId_example" // string | Instance Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceServiceGet(context.Background(), projectId, locationId, instanceId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceServiceGetRequest struct via the builder pattern


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


## DatabaseProjectInstanceServiceList

> []ResourceService DatabaseProjectInstanceServiceList(ctx, projectId, locationId, instanceId).Execute()

List database/instance.service



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
    instanceId := "instanceId_example" // string | Instance Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceServiceList(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceServiceListRequest struct via the builder pattern


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


## DatabaseProjectInstanceStart

> Database DatabaseProjectInstanceStart(ctx, projectId, locationId, instanceId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Start database/instance



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
    instanceId := "instanceId_example" // string | Instance Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceStart(context.Background(), projectId, locationId, instanceId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceStart`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceStop

> Database DatabaseProjectInstanceStop(ctx, projectId, locationId, instanceId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Stop database/instance



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
    instanceId := "instanceId_example" // string | Instance Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceStop(context.Background(), projectId, locationId, instanceId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceStop`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceTagCreate

> Tag DatabaseProjectInstanceTagCreate(ctx, projectId, locationId, instanceId).Tag(tag).Execute()

Create database/instance.tag



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
    instanceId := "instanceId_example" // string | Instance Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceTagCreate(context.Background(), projectId, locationId, instanceId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceTagCreateRequest struct via the builder pattern


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


## DatabaseProjectInstanceTagDelete

> DatabaseProjectInstanceTagDelete(ctx, projectId, locationId, instanceId, tagId).Execute()

Delete database/instance.tag



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
    instanceId := "instanceId_example" // string | Instance Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceTagDelete(context.Background(), projectId, locationId, instanceId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagDelete``: %v\n", err)
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
**instanceId** | **string** | Instance Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceTagDeleteRequest struct via the builder pattern


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


## DatabaseProjectInstanceTagGet

> Tag DatabaseProjectInstanceTagGet(ctx, projectId, locationId, instanceId, tagId).Execute()

Get database/instance.tag



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
    instanceId := "instanceId_example" // string | Instance Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceTagGet(context.Background(), projectId, locationId, instanceId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceTagGetRequest struct via the builder pattern


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


## DatabaseProjectInstanceTagList

> []Tag DatabaseProjectInstanceTagList(ctx, projectId, locationId, instanceId).Execute()

List database/instance.tag



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
    instanceId := "instanceId_example" // string | Instance Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceTagList(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceTagListRequest struct via the builder pattern


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


## DatabaseProjectInstanceTagPut

> []Tag DatabaseProjectInstanceTagPut(ctx, projectId, locationId, instanceId).Tag(tag).Execute()

Replace database/instance.tag



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
    instanceId := "instanceId_example" // string | Instance Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceTagPut(context.Background(), projectId, locationId, instanceId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceTagPutRequest struct via the builder pattern


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


## DatabaseProjectInstanceTransfer

> Database DatabaseProjectInstanceTransfer(ctx, projectId, locationId, instanceId).DatabaseProjectInstanceTransfer(databaseProjectInstanceTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer database/instance



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
    instanceId := "instanceId_example" // string | Instance Id
    databaseProjectInstanceTransfer := *openapiclient.NewDatabaseProjectInstanceTransfer("Project_example") // DatabaseProjectInstanceTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceTransfer(context.Background(), projectId, locationId, instanceId).DatabaseProjectInstanceTransfer(databaseProjectInstanceTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceTransfer`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **databaseProjectInstanceTransfer** | [**DatabaseProjectInstanceTransfer**](DatabaseProjectInstanceTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseProjectInstanceUpdate

> Database DatabaseProjectInstanceUpdate(ctx, projectId, locationId, instanceId).DatabaseProjectInstanceUpdate(databaseProjectInstanceUpdate).Execute()

Update database/instance



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
    instanceId := "instanceId_example" // string | Instance Id
    databaseProjectInstanceUpdate := *openapiclient.NewDatabaseProjectInstanceUpdate() // DatabaseProjectInstanceUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseProjectInstanceApi.DatabaseProjectInstanceUpdate(context.Background(), projectId, locationId, instanceId).DatabaseProjectInstanceUpdate(databaseProjectInstanceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseProjectInstanceApi.DatabaseProjectInstanceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseProjectInstanceUpdate`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseProjectInstanceApi.DatabaseProjectInstanceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**instanceId** | **string** | Instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseProjectInstanceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **databaseProjectInstanceUpdate** | [**DatabaseProjectInstanceUpdate**](DatabaseProjectInstanceUpdate.md) |  | 

### Return type

[**Database**](database.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

