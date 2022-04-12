# \IamProjectPolicyApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectPolicyActorCreate**](IamProjectPolicyApi.md#IamProjectPolicyActorCreate) | **Post** /iam/project/{projectId}/policy/{policyId}/actor | Create iam/policy.actor
[**IamProjectPolicyActorDelete**](IamProjectPolicyApi.md#IamProjectPolicyActorDelete) | **Delete** /iam/project/{projectId}/policy/{policyId}/actor/{actorId} | Delete iam/policy.actor
[**IamProjectPolicyActorGet**](IamProjectPolicyApi.md#IamProjectPolicyActorGet) | **Get** /iam/project/{projectId}/policy/{policyId}/actor/{actorId} | Get iam/policy.actor
[**IamProjectPolicyActorList**](IamProjectPolicyApi.md#IamProjectPolicyActorList) | **Get** /iam/project/{projectId}/policy/{policyId}/actor | List iam/policy.actor
[**IamProjectPolicyCreate**](IamProjectPolicyApi.md#IamProjectPolicyCreate) | **Post** /iam/project/{projectId}/policy | Create iam/policy
[**IamProjectPolicyDelete**](IamProjectPolicyApi.md#IamProjectPolicyDelete) | **Delete** /iam/project/{projectId}/policy/{policyId} | Delete iam/policy
[**IamProjectPolicyEventGet**](IamProjectPolicyApi.md#IamProjectPolicyEventGet) | **Get** /iam/project/{projectId}/policy/{policyId}/event/{eventId} | Get iam/policy.event
[**IamProjectPolicyEventList**](IamProjectPolicyApi.md#IamProjectPolicyEventList) | **Get** /iam/project/{projectId}/policy/{policyId}/event | List iam/policy.event
[**IamProjectPolicyGet**](IamProjectPolicyApi.md#IamProjectPolicyGet) | **Get** /iam/project/{projectId}/policy/{policyId} | Get iam/policy
[**IamProjectPolicyList**](IamProjectPolicyApi.md#IamProjectPolicyList) | **Get** /iam/project/{projectId}/policy | List iam/policy
[**IamProjectPolicyServiceGet**](IamProjectPolicyApi.md#IamProjectPolicyServiceGet) | **Get** /iam/project/{projectId}/policy/{policyId}/service/{serviceId} | Get iam/policy.service
[**IamProjectPolicyServiceList**](IamProjectPolicyApi.md#IamProjectPolicyServiceList) | **Get** /iam/project/{projectId}/policy/{policyId}/service | List iam/policy.service
[**IamProjectPolicyTagCreate**](IamProjectPolicyApi.md#IamProjectPolicyTagCreate) | **Post** /iam/project/{projectId}/policy/{policyId}/tag | Create iam/policy.tag
[**IamProjectPolicyTagDelete**](IamProjectPolicyApi.md#IamProjectPolicyTagDelete) | **Delete** /iam/project/{projectId}/policy/{policyId}/tag/{tagId} | Delete iam/policy.tag
[**IamProjectPolicyTagGet**](IamProjectPolicyApi.md#IamProjectPolicyTagGet) | **Get** /iam/project/{projectId}/policy/{policyId}/tag/{tagId} | Get iam/policy.tag
[**IamProjectPolicyTagList**](IamProjectPolicyApi.md#IamProjectPolicyTagList) | **Get** /iam/project/{projectId}/policy/{policyId}/tag | List iam/policy.tag
[**IamProjectPolicyTagPut**](IamProjectPolicyApi.md#IamProjectPolicyTagPut) | **Put** /iam/project/{projectId}/policy/{policyId}/tag | Replace iam/policy.tag
[**IamProjectPolicyUpdate**](IamProjectPolicyApi.md#IamProjectPolicyUpdate) | **Patch** /iam/project/{projectId}/policy/{policyId} | Update iam/policy



## IamProjectPolicyActorCreate

> IamActorOrGroup IamProjectPolicyActorCreate(ctx, projectId, policyId).IamActorOrGroup(iamActorOrGroup).Execute()

Create iam/policy.actor



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
    policyId := "policyId_example" // string | Policy Id
    iamActorOrGroup := *openapiclient.NewIamActorOrGroup("TODO") // IamActorOrGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyActorCreate(context.Background(), projectId, policyId).IamActorOrGroup(iamActorOrGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyActorCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyActorCreate`: IamActorOrGroup
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyActorCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyActorCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamActorOrGroup** | [**IamActorOrGroup**](IamActorOrGroup.md) |  | 

### Return type

[**IamActorOrGroup**](IamActorOrGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyActorDelete

> IamActorOrGroup IamProjectPolicyActorDelete(ctx, projectId, policyId, actorId).Execute()

Delete iam/policy.actor



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
    policyId := "policyId_example" // string | Policy Id
    actorId := "actorId_example" // string | actorId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyActorDelete(context.Background(), projectId, policyId, actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyActorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyActorDelete`: IamActorOrGroup
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyActorDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 
**actorId** | **string** | actorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyActorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamActorOrGroup**](IamActorOrGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyActorGet

> IamActorOrGroup IamProjectPolicyActorGet(ctx, projectId, policyId, actorId).Execute()

Get iam/policy.actor



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
    policyId := "policyId_example" // string | Policy Id
    actorId := "actorId_example" // string | actorId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyActorGet(context.Background(), projectId, policyId, actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyActorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyActorGet`: IamActorOrGroup
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyActorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 
**actorId** | **string** | actorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyActorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamActorOrGroup**](IamActorOrGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyActorList

> []IamActorOrGroup IamProjectPolicyActorList(ctx, projectId, policyId).Execute()

List iam/policy.actor



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
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyActorList(context.Background(), projectId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyActorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyActorList`: []IamActorOrGroup
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyActorList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyActorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IamActorOrGroup**](IamActorOrGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyCreate

> Policy IamProjectPolicyCreate(ctx, projectId).IamProjectPolicyCreate(iamProjectPolicyCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create iam/policy



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
    iamProjectPolicyCreate := *openapiclient.NewIamProjectPolicyCreate("Name_example", "Resource_example", "Role_example") // IamProjectPolicyCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyCreate(context.Background(), projectId).IamProjectPolicyCreate(iamProjectPolicyCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyCreate`: Policy
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectPolicyCreate** | [**IamProjectPolicyCreate**](IamProjectPolicyCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Policy**](Policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyDelete

> IamProjectPolicyDelete(ctx, projectId, policyId).Execute()

Delete iam/policy



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
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyDelete(context.Background(), projectId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyDeleteRequest struct via the builder pattern


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


## IamProjectPolicyEventGet

> Event IamProjectPolicyEventGet(ctx, projectId, policyId, eventId).Execute()

Get iam/policy.event



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
    policyId := "policyId_example" // string | Policy Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyEventGet(context.Background(), projectId, policyId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyEventGetRequest struct via the builder pattern


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


## IamProjectPolicyEventList

> []Event IamProjectPolicyEventList(ctx, projectId, policyId).Limit(limit).Skip(skip).Execute()

List iam/policy.event



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
    policyId := "policyId_example" // string | Policy Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyEventList(context.Background(), projectId, policyId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyEventListRequest struct via the builder pattern


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


## IamProjectPolicyGet

> Policy IamProjectPolicyGet(ctx, projectId, policyId).Execute()

Get iam/policy



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
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyGet(context.Background(), projectId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyGet`: Policy
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Policy**](Policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyList

> []Policy IamProjectPolicyList(ctx, projectId).Name(name).Resource(resource).TagValue(tagValue).TagKey(tagKey).Execute()

List iam/policy



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
    resource := "resource_example" // string | Filter by resource (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyList(context.Background(), projectId).Name(name).Resource(resource).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyList`: []Policy
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **resource** | **string** | Filter by resource | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Policy**](Policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPolicyServiceGet

> ResourceService IamProjectPolicyServiceGet(ctx, projectId, policyId, serviceId).Execute()

Get iam/policy.service



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
    policyId := "policyId_example" // string | Policy Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyServiceGet(context.Background(), projectId, policyId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyServiceGetRequest struct via the builder pattern


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


## IamProjectPolicyServiceList

> []ResourceService IamProjectPolicyServiceList(ctx, projectId, policyId).Execute()

List iam/policy.service



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
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyServiceList(context.Background(), projectId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyServiceListRequest struct via the builder pattern


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


## IamProjectPolicyTagCreate

> Tag IamProjectPolicyTagCreate(ctx, projectId, policyId).Tag(tag).Execute()

Create iam/policy.tag



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
    policyId := "policyId_example" // string | Policy Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyTagCreate(context.Background(), projectId, policyId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyTagCreateRequest struct via the builder pattern


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


## IamProjectPolicyTagDelete

> IamProjectPolicyTagDelete(ctx, projectId, policyId, tagId).Execute()

Delete iam/policy.tag



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
    policyId := "policyId_example" // string | Policy Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyTagDelete(context.Background(), projectId, policyId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyTagDeleteRequest struct via the builder pattern


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


## IamProjectPolicyTagGet

> Tag IamProjectPolicyTagGet(ctx, projectId, policyId, tagId).Execute()

Get iam/policy.tag



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
    policyId := "policyId_example" // string | Policy Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyTagGet(context.Background(), projectId, policyId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyTagGetRequest struct via the builder pattern


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


## IamProjectPolicyTagList

> []Tag IamProjectPolicyTagList(ctx, projectId, policyId).Execute()

List iam/policy.tag



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
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyTagList(context.Background(), projectId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyTagListRequest struct via the builder pattern


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


## IamProjectPolicyTagPut

> []Tag IamProjectPolicyTagPut(ctx, projectId, policyId).Tag(tag).Execute()

Replace iam/policy.tag



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
    policyId := "policyId_example" // string | Policy Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyTagPut(context.Background(), projectId, policyId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyTagPutRequest struct via the builder pattern


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


## IamProjectPolicyUpdate

> Policy IamProjectPolicyUpdate(ctx, projectId, policyId).IamProjectPolicyUpdate(iamProjectPolicyUpdate).Execute()

Update iam/policy



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
    policyId := "policyId_example" // string | Policy Id
    iamProjectPolicyUpdate := *openapiclient.NewIamProjectPolicyUpdate() // IamProjectPolicyUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectPolicyApi.IamProjectPolicyUpdate(context.Background(), projectId, policyId).IamProjectPolicyUpdate(iamProjectPolicyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectPolicyApi.IamProjectPolicyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPolicyUpdate`: Policy
    fmt.Fprintf(os.Stdout, "Response from `IamProjectPolicyApi.IamProjectPolicyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectPolicyUpdate** | [**IamProjectPolicyUpdate**](IamProjectPolicyUpdate.md) |  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

