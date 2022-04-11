# \VmhostProjectInstanceApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmhostProjectInstanceEventGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceEventGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/event/{eventId} | Get vmhost/instance.event
[**VmhostProjectInstanceEventList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceEventList) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/event | List vmhost/instance.event
[**VmhostProjectInstanceGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId} | Get vmhost/instance
[**VmhostProjectInstanceList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceList) | **Get** /vmhost/{locationId}/project/{projectId}/instance | List vmhost/instance
[**VmhostProjectInstanceServiceGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceServiceGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/service/{serviceId} | Get vmhost/instance.service
[**VmhostProjectInstanceServiceList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceServiceList) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/service | List vmhost/instance.service
[**VmhostProjectInstanceTagCreate**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagCreate) | **Post** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag | Create vmhost/instance.tag
[**VmhostProjectInstanceTagDelete**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagDelete) | **Delete** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Delete vmhost/instance.tag
[**VmhostProjectInstanceTagGet**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagGet) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag/{tagId} | Get vmhost/instance.tag
[**VmhostProjectInstanceTagList**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagList) | **Get** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag | List vmhost/instance.tag
[**VmhostProjectInstanceTagPut**](VmhostProjectInstanceApi.md#VmhostProjectInstanceTagPut) | **Put** /vmhost/{locationId}/project/{projectId}/instance/{instanceId}/tag | Replace vmhost/instance.tag



## VmhostProjectInstanceEventGet

> Event VmhostProjectInstanceEventGet(ctx, projectId, locationId, instanceId, eventId).Execute()

Get vmhost/instance.event



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceEventGet(context.Background(), projectId, locationId, instanceId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceEventGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceEventGetRequest struct via the builder pattern


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


## VmhostProjectInstanceEventList

> []Event VmhostProjectInstanceEventList(ctx, projectId, locationId, instanceId).Limit(limit).Skip(skip).Execute()

List vmhost/instance.event



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceEventList(context.Background(), projectId, locationId, instanceId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceEventList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceEventListRequest struct via the builder pattern


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


## VmhostProjectInstanceGet

> Vmhost VmhostProjectInstanceGet(ctx, projectId, locationId, instanceId).Execute()

Get vmhost/instance



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceGet(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceGet`: Vmhost
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Vmhost**](Vmhost.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostProjectInstanceList

> []Vmhost VmhostProjectInstanceList(ctx, projectId, locationId).EnabledServices(enabledServices).Execute()

List vmhost/instance



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
    enabledServices := "enabledServices_example" // string | Filter by enabledServices (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceList(context.Background(), projectId, locationId).EnabledServices(enabledServices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceList`: []Vmhost
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmhostProjectInstanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabledServices** | **string** | Filter by enabledServices | 

### Return type

[**[]Vmhost**](Vmhost.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmhostProjectInstanceServiceGet

> ResourceService VmhostProjectInstanceServiceGet(ctx, projectId, locationId, instanceId, serviceId).Execute()

Get vmhost/instance.service



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceServiceGet(context.Background(), projectId, locationId, instanceId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceServiceGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceServiceGetRequest struct via the builder pattern


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


## VmhostProjectInstanceServiceList

> []ResourceService VmhostProjectInstanceServiceList(ctx, projectId, locationId, instanceId).Execute()

List vmhost/instance.service



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceServiceList(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceServiceList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceServiceListRequest struct via the builder pattern


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


## VmhostProjectInstanceTagCreate

> Tag VmhostProjectInstanceTagCreate(ctx, projectId, locationId, instanceId).Tag(tag).Execute()

Create vmhost/instance.tag



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceTagCreate(context.Background(), projectId, locationId, instanceId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceTagCreate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceTagCreateRequest struct via the builder pattern


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


## VmhostProjectInstanceTagDelete

> VmhostProjectInstanceTagDelete(ctx, projectId, locationId, instanceId, tagId).Execute()

Delete vmhost/instance.tag



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceTagDelete(context.Background(), projectId, locationId, instanceId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceTagDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceTagDeleteRequest struct via the builder pattern


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


## VmhostProjectInstanceTagGet

> Tag VmhostProjectInstanceTagGet(ctx, projectId, locationId, instanceId, tagId).Execute()

Get vmhost/instance.tag



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceTagGet(context.Background(), projectId, locationId, instanceId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceTagGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceTagGetRequest struct via the builder pattern


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


## VmhostProjectInstanceTagList

> []Tag VmhostProjectInstanceTagList(ctx, projectId, locationId, instanceId).Execute()

List vmhost/instance.tag



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceTagList(context.Background(), projectId, locationId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceTagList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceTagListRequest struct via the builder pattern


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


## VmhostProjectInstanceTagPut

> []Tag VmhostProjectInstanceTagPut(ctx, projectId, locationId, instanceId).Tag(tag).Execute()

Replace vmhost/instance.tag



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VmhostProjectInstanceApi.VmhostProjectInstanceTagPut(context.Background(), projectId, locationId, instanceId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VmhostProjectInstanceApi.VmhostProjectInstanceTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmhostProjectInstanceTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `VmhostProjectInstanceApi.VmhostProjectInstanceTagPut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVmhostProjectInstanceTagPutRequest struct via the builder pattern


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

