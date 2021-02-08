# \IamProjectRoleApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectRoleCreate**](IamProjectRoleApi.md#IamProjectRoleCreate) | **Post** /iam/project/{projectId}/role | Create iam/role
[**IamProjectRoleDelete**](IamProjectRoleApi.md#IamProjectRoleDelete) | **Delete** /iam/project/{projectId}/role/{roleId} | Delete iam/role
[**IamProjectRoleEventGet**](IamProjectRoleApi.md#IamProjectRoleEventGet) | **Get** /iam/project/{projectId}/role/{roleId}/event/{eventId} | Get iam/role.event
[**IamProjectRoleEventList**](IamProjectRoleApi.md#IamProjectRoleEventList) | **Get** /iam/project/{projectId}/role/{roleId}/event | List iam/role.event
[**IamProjectRoleGet**](IamProjectRoleApi.md#IamProjectRoleGet) | **Get** /iam/project/{projectId}/role/{roleId} | Get iam/role
[**IamProjectRoleList**](IamProjectRoleApi.md#IamProjectRoleList) | **Get** /iam/project/{projectId}/role | List iam/role
[**IamProjectRolePermissionCreate**](IamProjectRoleApi.md#IamProjectRolePermissionCreate) | **Post** /iam/project/{projectId}/role/{roleId}/permission | Create iam/role.permission
[**IamProjectRolePermissionDelete**](IamProjectRoleApi.md#IamProjectRolePermissionDelete) | **Delete** /iam/project/{projectId}/role/{roleId}/permission/{permissionId} | Delete iam/role.permission
[**IamProjectRolePermissionGet**](IamProjectRoleApi.md#IamProjectRolePermissionGet) | **Get** /iam/project/{projectId}/role/{roleId}/permission/{permissionId} | Get iam/role.permission
[**IamProjectRolePermissionList**](IamProjectRoleApi.md#IamProjectRolePermissionList) | **Get** /iam/project/{projectId}/role/{roleId}/permission | List iam/role.permission
[**IamProjectRolePermissionPut**](IamProjectRoleApi.md#IamProjectRolePermissionPut) | **Put** /iam/project/{projectId}/role/{roleId}/permission | Replace iam/role.permission
[**IamProjectRoleServiceGet**](IamProjectRoleApi.md#IamProjectRoleServiceGet) | **Get** /iam/project/{projectId}/role/{roleId}/service/{serviceId} | Get iam/role.service
[**IamProjectRoleServiceList**](IamProjectRoleApi.md#IamProjectRoleServiceList) | **Get** /iam/project/{projectId}/role/{roleId}/service | List iam/role.service
[**IamProjectRoleTagCreate**](IamProjectRoleApi.md#IamProjectRoleTagCreate) | **Post** /iam/project/{projectId}/role/{roleId}/tag | Create iam/role.tag
[**IamProjectRoleTagDelete**](IamProjectRoleApi.md#IamProjectRoleTagDelete) | **Delete** /iam/project/{projectId}/role/{roleId}/tag/{tagId} | Delete iam/role.tag
[**IamProjectRoleTagGet**](IamProjectRoleApi.md#IamProjectRoleTagGet) | **Get** /iam/project/{projectId}/role/{roleId}/tag/{tagId} | Get iam/role.tag
[**IamProjectRoleTagList**](IamProjectRoleApi.md#IamProjectRoleTagList) | **Get** /iam/project/{projectId}/role/{roleId}/tag | List iam/role.tag
[**IamProjectRoleTagPut**](IamProjectRoleApi.md#IamProjectRoleTagPut) | **Put** /iam/project/{projectId}/role/{roleId}/tag | Replace iam/role.tag
[**IamProjectRoleUpdate**](IamProjectRoleApi.md#IamProjectRoleUpdate) | **Patch** /iam/project/{projectId}/role/{roleId} | Update iam/role



## IamProjectRoleCreate

> Role IamProjectRoleCreate(ctx, projectId).IamProjectRoleCreate(iamProjectRoleCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create iam/role



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
    iamProjectRoleCreate := *openapiclient.NewIamProjectRoleCreate("Name_example") // IamProjectRoleCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleCreate(context.Background(), projectId).IamProjectRoleCreate(iamProjectRoleCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleCreate`: Role
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectRoleCreate** | [**IamProjectRoleCreate**](IamProjectRoleCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRoleDelete

> IamProjectRoleDelete(ctx, projectId, roleId).Execute()

Delete iam/role



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
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleDelete(context.Background(), projectId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleDeleteRequest struct via the builder pattern


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


## IamProjectRoleEventGet

> Event IamProjectRoleEventGet(ctx, projectId, roleId, eventId).Execute()

Get iam/role.event



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
    roleId := "roleId_example" // string | Role Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleEventGet(context.Background(), projectId, roleId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleEventGetRequest struct via the builder pattern


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


## IamProjectRoleEventList

> []Event IamProjectRoleEventList(ctx, projectId, roleId).Limit(limit).Skip(skip).Execute()

List iam/role.event



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
    roleId := "roleId_example" // string | Role Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleEventList(context.Background(), projectId, roleId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleEventListRequest struct via the builder pattern


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


## IamProjectRoleGet

> Role IamProjectRoleGet(ctx, projectId, roleId).Execute()

Get iam/role



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
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleGet(context.Background(), projectId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleGet`: Role
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRoleList

> []Role IamProjectRoleList(ctx, projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List iam/role



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleList(context.Background(), projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleList`: []Role
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRolePermissionCreate

> IamPermission IamProjectRolePermissionCreate(ctx, projectId, roleId).IamPermission(iamPermission).Execute()

Create iam/role.permission



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
    roleId := "roleId_example" // string | Role Id
    iamPermission := *openapiclient.NewIamPermission("Value_example") // IamPermission | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRolePermissionCreate(context.Background(), projectId, roleId).IamPermission(iamPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRolePermissionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRolePermissionCreate`: IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRolePermissionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRolePermissionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamPermission** | [**IamPermission**](IamPermission.md) |  | 

### Return type

[**IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRolePermissionDelete

> IamPermission IamProjectRolePermissionDelete(ctx, projectId, roleId, permissionId).Execute()

Delete iam/role.permission



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
    roleId := "roleId_example" // string | Role Id
    permissionId := "permissionId_example" // string | permissionId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRolePermissionDelete(context.Background(), projectId, roleId, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRolePermissionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRolePermissionDelete`: IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRolePermissionDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 
**permissionId** | **string** | permissionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRolePermissionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRolePermissionGet

> IamPermission IamProjectRolePermissionGet(ctx, projectId, roleId, permissionId).Execute()

Get iam/role.permission



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
    roleId := "roleId_example" // string | Role Id
    permissionId := "permissionId_example" // string | permissionId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRolePermissionGet(context.Background(), projectId, roleId, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRolePermissionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRolePermissionGet`: IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRolePermissionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 
**permissionId** | **string** | permissionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRolePermissionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRolePermissionList

> []IamPermission IamProjectRolePermissionList(ctx, projectId, roleId).Execute()

List iam/role.permission



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
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRolePermissionList(context.Background(), projectId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRolePermissionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRolePermissionList`: []IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRolePermissionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRolePermissionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRolePermissionPut

> []IamPermission IamProjectRolePermissionPut(ctx, projectId, roleId).IamPermission(iamPermission).Execute()

Replace iam/role.permission



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
    roleId := "roleId_example" // string | Role Id
    iamPermission := []openapiclient.IamPermission{*openapiclient.NewIamPermission("Value_example")} // []IamPermission | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRolePermissionPut(context.Background(), projectId, roleId).IamPermission(iamPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRolePermissionPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRolePermissionPut`: []IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRolePermissionPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRolePermissionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamPermission** | [**[]IamPermission**](iam.permission.md) |  | 

### Return type

[**[]IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectRoleServiceGet

> ResourceService IamProjectRoleServiceGet(ctx, projectId, roleId, serviceId).Execute()

Get iam/role.service



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
    roleId := "roleId_example" // string | Role Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleServiceGet(context.Background(), projectId, roleId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleServiceGetRequest struct via the builder pattern


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


## IamProjectRoleServiceList

> []ResourceService IamProjectRoleServiceList(ctx, projectId, roleId).Execute()

List iam/role.service



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
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleServiceList(context.Background(), projectId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleServiceListRequest struct via the builder pattern


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


## IamProjectRoleTagCreate

> Tag IamProjectRoleTagCreate(ctx, projectId, roleId).Tag(tag).Execute()

Create iam/role.tag



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
    roleId := "roleId_example" // string | Role Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleTagCreate(context.Background(), projectId, roleId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleTagCreateRequest struct via the builder pattern


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


## IamProjectRoleTagDelete

> IamProjectRoleTagDelete(ctx, projectId, roleId, tagId).Execute()

Delete iam/role.tag



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
    roleId := "roleId_example" // string | Role Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleTagDelete(context.Background(), projectId, roleId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleTagDeleteRequest struct via the builder pattern


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


## IamProjectRoleTagGet

> Tag IamProjectRoleTagGet(ctx, projectId, roleId, tagId).Execute()

Get iam/role.tag



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
    roleId := "roleId_example" // string | Role Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleTagGet(context.Background(), projectId, roleId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleTagGetRequest struct via the builder pattern


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


## IamProjectRoleTagList

> []Tag IamProjectRoleTagList(ctx, projectId, roleId).Execute()

List iam/role.tag



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
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleTagList(context.Background(), projectId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleTagListRequest struct via the builder pattern


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


## IamProjectRoleTagPut

> []Tag IamProjectRoleTagPut(ctx, projectId, roleId).Tag(tag).Execute()

Replace iam/role.tag



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
    roleId := "roleId_example" // string | Role Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleTagPut(context.Background(), projectId, roleId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleTagPutRequest struct via the builder pattern


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


## IamProjectRoleUpdate

> Role IamProjectRoleUpdate(ctx, projectId, roleId).IamProjectRoleUpdate(iamProjectRoleUpdate).Execute()

Update iam/role



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
    roleId := "roleId_example" // string | Role Id
    iamProjectRoleUpdate := *openapiclient.NewIamProjectRoleUpdate() // IamProjectRoleUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamProjectRoleApi.IamProjectRoleUpdate(context.Background(), projectId, roleId).IamProjectRoleUpdate(iamProjectRoleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectRoleApi.IamProjectRoleUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectRoleUpdate`: Role
    fmt.Fprintf(os.Stdout, "Response from `IamProjectRoleApi.IamProjectRoleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectRoleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectRoleUpdate** | [**IamProjectRoleUpdate**](IamProjectRoleUpdate.md) |  | 

### Return type

[**Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

