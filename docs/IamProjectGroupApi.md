# \IamProjectGroupApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectGroupActorCreate**](IamProjectGroupApi.md#IamProjectGroupActorCreate) | **Post** /iam/project/{projectId}/group/{groupId}/actor | Create iam/group.actor
[**IamProjectGroupActorDelete**](IamProjectGroupApi.md#IamProjectGroupActorDelete) | **Delete** /iam/project/{projectId}/group/{groupId}/actor/{actorId} | Delete iam/group.actor
[**IamProjectGroupActorGet**](IamProjectGroupApi.md#IamProjectGroupActorGet) | **Get** /iam/project/{projectId}/group/{groupId}/actor/{actorId} | Get iam/group.actor
[**IamProjectGroupActorList**](IamProjectGroupApi.md#IamProjectGroupActorList) | **Get** /iam/project/{projectId}/group/{groupId}/actor | List iam/group.actor
[**IamProjectGroupCreate**](IamProjectGroupApi.md#IamProjectGroupCreate) | **Post** /iam/project/{projectId}/group | Create iam/group
[**IamProjectGroupDelete**](IamProjectGroupApi.md#IamProjectGroupDelete) | **Delete** /iam/project/{projectId}/group/{groupId} | Delete iam/group
[**IamProjectGroupEventGet**](IamProjectGroupApi.md#IamProjectGroupEventGet) | **Get** /iam/project/{projectId}/group/{groupId}/event/{eventId} | Get iam/group.event
[**IamProjectGroupEventList**](IamProjectGroupApi.md#IamProjectGroupEventList) | **Get** /iam/project/{projectId}/group/{groupId}/event | List iam/group.event
[**IamProjectGroupGet**](IamProjectGroupApi.md#IamProjectGroupGet) | **Get** /iam/project/{projectId}/group/{groupId} | Get iam/group
[**IamProjectGroupList**](IamProjectGroupApi.md#IamProjectGroupList) | **Get** /iam/project/{projectId}/group | List iam/group
[**IamProjectGroupServiceGet**](IamProjectGroupApi.md#IamProjectGroupServiceGet) | **Get** /iam/project/{projectId}/group/{groupId}/service/{serviceId} | Get iam/group.service
[**IamProjectGroupServiceList**](IamProjectGroupApi.md#IamProjectGroupServiceList) | **Get** /iam/project/{projectId}/group/{groupId}/service | List iam/group.service
[**IamProjectGroupTagCreate**](IamProjectGroupApi.md#IamProjectGroupTagCreate) | **Post** /iam/project/{projectId}/group/{groupId}/tag | Create iam/group.tag
[**IamProjectGroupTagDelete**](IamProjectGroupApi.md#IamProjectGroupTagDelete) | **Delete** /iam/project/{projectId}/group/{groupId}/tag/{tagId} | Delete iam/group.tag
[**IamProjectGroupTagGet**](IamProjectGroupApi.md#IamProjectGroupTagGet) | **Get** /iam/project/{projectId}/group/{groupId}/tag/{tagId} | Get iam/group.tag
[**IamProjectGroupTagList**](IamProjectGroupApi.md#IamProjectGroupTagList) | **Get** /iam/project/{projectId}/group/{groupId}/tag | List iam/group.tag
[**IamProjectGroupTagPut**](IamProjectGroupApi.md#IamProjectGroupTagPut) | **Put** /iam/project/{projectId}/group/{groupId}/tag | Replace iam/group.tag
[**IamProjectGroupUpdate**](IamProjectGroupApi.md#IamProjectGroupUpdate) | **Patch** /iam/project/{projectId}/group/{groupId} | Update iam/group



## IamProjectGroupActorCreate

> IamActor IamProjectGroupActorCreate(ctx, projectId, groupId).IamActor(iamActor).Execute()

Create iam/group.actor



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
    groupId := "groupId_example" // string | Group Id
    iamActor := *openapiclient.NewIamActor("Value_example") // IamActor | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupActorCreate(context.Background(), projectId, groupId).IamActor(iamActor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupActorCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupActorCreate`: IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupActorCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupActorCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamActor** | [**IamActor**](IamActor.md) |  | 

### Return type

[**IamActor**](IamActor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectGroupActorDelete

> IamActor IamProjectGroupActorDelete(ctx, projectId, groupId, actorId).Execute()

Delete iam/group.actor



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
    groupId := "groupId_example" // string | Group Id
    actorId := "actorId_example" // string | actorId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupActorDelete(context.Background(), projectId, groupId, actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupActorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupActorDelete`: IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupActorDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 
**actorId** | **string** | actorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupActorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamActor**](IamActor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectGroupActorGet

> IamActor IamProjectGroupActorGet(ctx, projectId, groupId, actorId).Execute()

Get iam/group.actor



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
    groupId := "groupId_example" // string | Group Id
    actorId := "actorId_example" // string | actorId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupActorGet(context.Background(), projectId, groupId, actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupActorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupActorGet`: IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupActorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 
**actorId** | **string** | actorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupActorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamActor**](IamActor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectGroupActorList

> []IamActor IamProjectGroupActorList(ctx, projectId, groupId).Execute()

List iam/group.actor



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
    groupId := "groupId_example" // string | Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupActorList(context.Background(), projectId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupActorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupActorList`: []IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupActorList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupActorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IamActor**](IamActor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectGroupCreate

> Group IamProjectGroupCreate(ctx, projectId).IamProjectGroupCreate(iamProjectGroupCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create iam/group



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
    iamProjectGroupCreate := *openapiclient.NewIamProjectGroupCreate("Name_example") // IamProjectGroupCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupCreate(context.Background(), projectId).IamProjectGroupCreate(iamProjectGroupCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupCreate`: Group
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectGroupCreate** | [**IamProjectGroupCreate**](IamProjectGroupCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectGroupDelete

> InlineResponseDefault IamProjectGroupDelete(ctx, projectId, groupId).Execute()

Delete iam/group



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
    groupId := "groupId_example" // string | Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupDelete(context.Background(), projectId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupDelete`: InlineResponseDefault
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupDeleteRequest struct via the builder pattern


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


## IamProjectGroupEventGet

> Event IamProjectGroupEventGet(ctx, projectId, groupId, eventId).Execute()

Get iam/group.event



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
    groupId := "groupId_example" // string | Group Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupEventGet(context.Background(), projectId, groupId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupEventGetRequest struct via the builder pattern


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


## IamProjectGroupEventList

> []Event IamProjectGroupEventList(ctx, projectId, groupId).Limit(limit).Skip(skip).Execute()

List iam/group.event



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
    groupId := "groupId_example" // string | Group Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupEventList(context.Background(), projectId, groupId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupEventListRequest struct via the builder pattern


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


## IamProjectGroupGet

> Group IamProjectGroupGet(ctx, projectId, groupId).Execute()

Get iam/group



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
    groupId := "groupId_example" // string | Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupGet(context.Background(), projectId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupGet`: Group
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectGroupList

> []Group IamProjectGroupList(ctx, projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List iam/group



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
    name := "name_example" // string | Filter by name (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupList(context.Background(), projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupList`: []Group
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectGroupServiceGet

> ResourceService IamProjectGroupServiceGet(ctx, projectId, groupId, serviceId).Execute()

Get iam/group.service



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
    groupId := "groupId_example" // string | Group Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupServiceGet(context.Background(), projectId, groupId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupServiceGetRequest struct via the builder pattern


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


## IamProjectGroupServiceList

> []ResourceService IamProjectGroupServiceList(ctx, projectId, groupId).Execute()

List iam/group.service



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
    groupId := "groupId_example" // string | Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupServiceList(context.Background(), projectId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupServiceListRequest struct via the builder pattern


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


## IamProjectGroupTagCreate

> Tag IamProjectGroupTagCreate(ctx, projectId, groupId).Tag(tag).Execute()

Create iam/group.tag



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
    groupId := "groupId_example" // string | Group Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupTagCreate(context.Background(), projectId, groupId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupTagCreateRequest struct via the builder pattern


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


## IamProjectGroupTagDelete

> IamProjectGroupTagDelete(ctx, projectId, groupId, tagId).Execute()

Delete iam/group.tag



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
    groupId := "groupId_example" // string | Group Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupTagDelete(context.Background(), projectId, groupId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupTagDeleteRequest struct via the builder pattern


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


## IamProjectGroupTagGet

> Tag IamProjectGroupTagGet(ctx, projectId, groupId, tagId).Execute()

Get iam/group.tag



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
    groupId := "groupId_example" // string | Group Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupTagGet(context.Background(), projectId, groupId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupTagGetRequest struct via the builder pattern


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


## IamProjectGroupTagList

> []Tag IamProjectGroupTagList(ctx, projectId, groupId).Execute()

List iam/group.tag



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
    groupId := "groupId_example" // string | Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupTagList(context.Background(), projectId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupTagListRequest struct via the builder pattern


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


## IamProjectGroupTagPut

> []Tag IamProjectGroupTagPut(ctx, projectId, groupId).Tag(tag).Execute()

Replace iam/group.tag



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
    groupId := "groupId_example" // string | Group Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupTagPut(context.Background(), projectId, groupId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupTagPutRequest struct via the builder pattern


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


## IamProjectGroupUpdate

> Group IamProjectGroupUpdate(ctx, projectId, groupId).IamProjectGroupUpdate(iamProjectGroupUpdate).Execute()

Update iam/group



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
    groupId := "groupId_example" // string | Group Id
    iamProjectGroupUpdate := *openapiclient.NewIamProjectGroupUpdate() // IamProjectGroupUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectGroupApi.IamProjectGroupUpdate(context.Background(), projectId, groupId).IamProjectGroupUpdate(iamProjectGroupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectGroupApi.IamProjectGroupUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGroupUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `IamProjectGroupApi.IamProjectGroupUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**groupId** | **string** | Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGroupUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectGroupUpdate** | [**IamProjectGroupUpdate**](IamProjectGroupUpdate.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

