# \IamOrganisationRoleApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamOrganisationRoleCreate**](IamOrganisationRoleApi.md#IamOrganisationRoleCreate) | **Post** /iam/organisation/{organisationId}/role | Create iam/role
[**IamOrganisationRoleDelete**](IamOrganisationRoleApi.md#IamOrganisationRoleDelete) | **Delete** /iam/organisation/{organisationId}/role/{roleId} | Delete iam/role
[**IamOrganisationRoleEventGet**](IamOrganisationRoleApi.md#IamOrganisationRoleEventGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/event/{eventId} | Get iam/role.event
[**IamOrganisationRoleEventList**](IamOrganisationRoleApi.md#IamOrganisationRoleEventList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/event | List iam/role.event
[**IamOrganisationRoleGet**](IamOrganisationRoleApi.md#IamOrganisationRoleGet) | **Get** /iam/organisation/{organisationId}/role/{roleId} | Get iam/role
[**IamOrganisationRoleList**](IamOrganisationRoleApi.md#IamOrganisationRoleList) | **Get** /iam/organisation/{organisationId}/role | List iam/role
[**IamOrganisationRolePermissionCreate**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionCreate) | **Post** /iam/organisation/{organisationId}/role/{roleId}/permission | Create iam/role.permission
[**IamOrganisationRolePermissionDelete**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionDelete) | **Delete** /iam/organisation/{organisationId}/role/{roleId}/permission/{permissionId} | Delete iam/role.permission
[**IamOrganisationRolePermissionGet**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/permission/{permissionId} | Get iam/role.permission
[**IamOrganisationRolePermissionList**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/permission | List iam/role.permission
[**IamOrganisationRolePermissionPut**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionPut) | **Put** /iam/organisation/{organisationId}/role/{roleId}/permission | Replace iam/role.permission
[**IamOrganisationRoleServiceGet**](IamOrganisationRoleApi.md#IamOrganisationRoleServiceGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/service/{serviceId} | Get iam/role.service
[**IamOrganisationRoleServiceList**](IamOrganisationRoleApi.md#IamOrganisationRoleServiceList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/service | List iam/role.service
[**IamOrganisationRoleTagCreate**](IamOrganisationRoleApi.md#IamOrganisationRoleTagCreate) | **Post** /iam/organisation/{organisationId}/role/{roleId}/tag | Create iam/role.tag
[**IamOrganisationRoleTagDelete**](IamOrganisationRoleApi.md#IamOrganisationRoleTagDelete) | **Delete** /iam/organisation/{organisationId}/role/{roleId}/tag/{tagId} | Delete iam/role.tag
[**IamOrganisationRoleTagGet**](IamOrganisationRoleApi.md#IamOrganisationRoleTagGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/tag/{tagId} | Get iam/role.tag
[**IamOrganisationRoleTagList**](IamOrganisationRoleApi.md#IamOrganisationRoleTagList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/tag | List iam/role.tag
[**IamOrganisationRoleTagPut**](IamOrganisationRoleApi.md#IamOrganisationRoleTagPut) | **Put** /iam/organisation/{organisationId}/role/{roleId}/tag | Replace iam/role.tag
[**IamOrganisationRoleUpdate**](IamOrganisationRoleApi.md#IamOrganisationRoleUpdate) | **Patch** /iam/organisation/{organisationId}/role/{roleId} | Update iam/role



## IamOrganisationRoleCreate

> Role IamOrganisationRoleCreate(ctx, organisationId).IamProjectRoleCreate(iamProjectRoleCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    iamProjectRoleCreate := *openapiclient.NewIamProjectRoleCreate("Name_example") // IamProjectRoleCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleCreate(context.Background(), organisationId).IamProjectRoleCreate(iamProjectRoleCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleCreate`: Role
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectRoleCreate** | [**IamProjectRoleCreate**](IamProjectRoleCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Role**](Role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRoleDelete

> Role IamOrganisationRoleDelete(ctx, organisationId, roleId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleDelete(context.Background(), organisationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleDelete`: Role
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRoleEventGet

> Event IamOrganisationRoleEventGet(ctx, organisationId, roleId, eventId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleEventGet(context.Background(), organisationId, roleId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleEventGetRequest struct via the builder pattern


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


## IamOrganisationRoleEventList

> []Event IamOrganisationRoleEventList(ctx, organisationId, roleId).Limit(limit).Skip(skip).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleEventList(context.Background(), organisationId, roleId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleEventListRequest struct via the builder pattern


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


## IamOrganisationRoleGet

> Role IamOrganisationRoleGet(ctx, organisationId, roleId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleGet(context.Background(), organisationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleGet`: Role
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRoleList

> []Role IamOrganisationRoleList(ctx, organisationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    name := "name_example" // string | Filter by name (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleList(context.Background(), organisationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleList`: []Role
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Role**](Role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionCreate

> IamPermission IamOrganisationRolePermissionCreate(ctx, organisationId, roleId).IamPermission(iamPermission).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    iamPermission := *openapiclient.NewIamPermission("Value_example") // IamPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRolePermissionCreate(context.Background(), organisationId, roleId).IamPermission(iamPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRolePermissionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRolePermissionCreate`: IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRolePermissionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRolePermissionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamPermission** | [**IamPermission**](IamPermission.md) |  | 

### Return type

[**IamPermission**](IamPermission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionDelete

> IamPermission IamOrganisationRolePermissionDelete(ctx, organisationId, roleId, permissionId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    permissionId := "permissionId_example" // string | permissionId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRolePermissionDelete(context.Background(), organisationId, roleId, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRolePermissionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRolePermissionDelete`: IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRolePermissionDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 
**permissionId** | **string** | permissionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRolePermissionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamPermission**](IamPermission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionGet

> IamPermission IamOrganisationRolePermissionGet(ctx, organisationId, roleId, permissionId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    permissionId := "permissionId_example" // string | permissionId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRolePermissionGet(context.Background(), organisationId, roleId, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRolePermissionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRolePermissionGet`: IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRolePermissionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 
**permissionId** | **string** | permissionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRolePermissionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamPermission**](IamPermission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionList

> []IamPermission IamOrganisationRolePermissionList(ctx, organisationId, roleId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRolePermissionList(context.Background(), organisationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRolePermissionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRolePermissionList`: []IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRolePermissionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRolePermissionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IamPermission**](IamPermission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionPut

> []IamPermission IamOrganisationRolePermissionPut(ctx, organisationId, roleId).IamPermission(iamPermission).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    iamPermission := []openapiclient.IamPermission{*openapiclient.NewIamPermission("Value_example")} // []IamPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRolePermissionPut(context.Background(), organisationId, roleId).IamPermission(iamPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRolePermissionPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRolePermissionPut`: []IamPermission
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRolePermissionPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRolePermissionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamPermission** | [**[]IamPermission**](IamPermission.md) |  | 

### Return type

[**[]IamPermission**](IamPermission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRoleServiceGet

> ResourceService IamOrganisationRoleServiceGet(ctx, organisationId, roleId, serviceId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleServiceGet(context.Background(), organisationId, roleId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleServiceGetRequest struct via the builder pattern


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


## IamOrganisationRoleServiceList

> []ResourceService IamOrganisationRoleServiceList(ctx, organisationId, roleId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleServiceList(context.Background(), organisationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleServiceListRequest struct via the builder pattern


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


## IamOrganisationRoleTagCreate

> Tag IamOrganisationRoleTagCreate(ctx, organisationId, roleId).Tag(tag).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleTagCreate(context.Background(), organisationId, roleId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleTagCreateRequest struct via the builder pattern


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


## IamOrganisationRoleTagDelete

> IamOrganisationRoleTagDelete(ctx, organisationId, roleId, tagId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleTagDelete(context.Background(), organisationId, roleId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleTagDeleteRequest struct via the builder pattern


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


## IamOrganisationRoleTagGet

> Tag IamOrganisationRoleTagGet(ctx, organisationId, roleId, tagId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleTagGet(context.Background(), organisationId, roleId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleTagGetRequest struct via the builder pattern


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


## IamOrganisationRoleTagList

> []Tag IamOrganisationRoleTagList(ctx, organisationId, roleId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleTagList(context.Background(), organisationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleTagListRequest struct via the builder pattern


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


## IamOrganisationRoleTagPut

> []Tag IamOrganisationRoleTagPut(ctx, organisationId, roleId).Tag(tag).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleTagPut(context.Background(), organisationId, roleId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleTagPutRequest struct via the builder pattern


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


## IamOrganisationRoleUpdate

> Role IamOrganisationRoleUpdate(ctx, organisationId, roleId).IamProjectRoleUpdate(iamProjectRoleUpdate).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    roleId := "roleId_example" // string | Role Id
    iamProjectRoleUpdate := *openapiclient.NewIamProjectRoleUpdate() // IamProjectRoleUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationRoleApi.IamOrganisationRoleUpdate(context.Background(), organisationId, roleId).IamProjectRoleUpdate(iamProjectRoleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationRoleApi.IamOrganisationRoleUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationRoleUpdate`: Role
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationRoleApi.IamOrganisationRoleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**roleId** | **string** | Role Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationRoleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectRoleUpdate** | [**IamProjectRoleUpdate**](IamProjectRoleUpdate.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

