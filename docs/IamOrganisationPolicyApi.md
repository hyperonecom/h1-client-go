# \IamOrganisationPolicyApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamOrganisationPolicyActorCreate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorCreate) | **Post** /iam/organisation/{organisationId}/policy/{policyId}/actor | Create iam/policy.actor
[**IamOrganisationPolicyActorDelete**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorDelete) | **Delete** /iam/organisation/{organisationId}/policy/{policyId}/actor/{actorId} | Delete iam/policy.actor
[**IamOrganisationPolicyActorGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/actor/{actorId} | Get iam/policy.actor
[**IamOrganisationPolicyActorList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyActorList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/actor | List iam/policy.actor
[**IamOrganisationPolicyCreate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyCreate) | **Post** /iam/organisation/{organisationId}/policy | Create iam/policy
[**IamOrganisationPolicyDelete**](IamOrganisationPolicyApi.md#IamOrganisationPolicyDelete) | **Delete** /iam/organisation/{organisationId}/policy/{policyId} | Delete iam/policy
[**IamOrganisationPolicyEventGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyEventGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/event/{eventId} | Get iam/policy.event
[**IamOrganisationPolicyEventList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyEventList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/event | List iam/policy.event
[**IamOrganisationPolicyGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId} | Get iam/policy
[**IamOrganisationPolicyList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyList) | **Get** /iam/organisation/{organisationId}/policy | List iam/policy
[**IamOrganisationPolicyServiceGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyServiceGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/service/{serviceId} | Get iam/policy.service
[**IamOrganisationPolicyServiceList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyServiceList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/service | List iam/policy.service
[**IamOrganisationPolicyTagCreate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagCreate) | **Post** /iam/organisation/{organisationId}/policy/{policyId}/tag | Create iam/policy.tag
[**IamOrganisationPolicyTagDelete**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagDelete) | **Delete** /iam/organisation/{organisationId}/policy/{policyId}/tag/{tagId} | Delete iam/policy.tag
[**IamOrganisationPolicyTagGet**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagGet) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/tag/{tagId} | Get iam/policy.tag
[**IamOrganisationPolicyTagList**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagList) | **Get** /iam/organisation/{organisationId}/policy/{policyId}/tag | List iam/policy.tag
[**IamOrganisationPolicyTagPut**](IamOrganisationPolicyApi.md#IamOrganisationPolicyTagPut) | **Put** /iam/organisation/{organisationId}/policy/{policyId}/tag | Replace iam/policy.tag
[**IamOrganisationPolicyUpdate**](IamOrganisationPolicyApi.md#IamOrganisationPolicyUpdate) | **Patch** /iam/organisation/{organisationId}/policy/{policyId} | Update iam/policy



## IamOrganisationPolicyActorCreate

> IamActor IamOrganisationPolicyActorCreate(ctx, organisationId, policyId).IamActor(iamActor).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    iamActor := *openapiclient.NewIamActor("Value_example") // IamActor | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyActorCreate(context.Background(), organisationId, policyId).IamActor(iamActor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyActorCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyActorCreate`: IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyActorCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyActorCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamActor** | [**IamActor**](IamActor.md) |  | 

### Return type

[**IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPolicyActorDelete

> IamActor IamOrganisationPolicyActorDelete(ctx, organisationId, policyId, actorId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    actorId := "actorId_example" // string | actorId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyActorDelete(context.Background(), organisationId, policyId, actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyActorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyActorDelete`: IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyActorDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 
**actorId** | **string** | actorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyActorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPolicyActorGet

> IamActor IamOrganisationPolicyActorGet(ctx, organisationId, policyId, actorId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    actorId := "actorId_example" // string | actorId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyActorGet(context.Background(), organisationId, policyId, actorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyActorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyActorGet`: IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyActorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 
**actorId** | **string** | actorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyActorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPolicyActorList

> []IamActor IamOrganisationPolicyActorList(ctx, organisationId, policyId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyActorList(context.Background(), organisationId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyActorList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyActorList`: []IamActor
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyActorList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyActorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IamActor**](iam.actor.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPolicyCreate

> Policy IamOrganisationPolicyCreate(ctx, organisationId).IamProjectPolicyCreate(iamProjectPolicyCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    iamProjectPolicyCreate := *openapiclient.NewIamProjectPolicyCreate("Name_example", "Role_example", "Resource_example") // IamProjectPolicyCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyCreate(context.Background(), organisationId).IamProjectPolicyCreate(iamProjectPolicyCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyCreate`: Policy
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectPolicyCreate** | [**IamProjectPolicyCreate**](IamProjectPolicyCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPolicyDelete

> IamOrganisationPolicyDelete(ctx, organisationId, policyId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyDelete(context.Background(), organisationId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyDeleteRequest struct via the builder pattern


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


## IamOrganisationPolicyEventGet

> Event IamOrganisationPolicyEventGet(ctx, organisationId, policyId, eventId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyEventGet(context.Background(), organisationId, policyId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyEventGetRequest struct via the builder pattern


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


## IamOrganisationPolicyEventList

> []Event IamOrganisationPolicyEventList(ctx, organisationId, policyId).Limit(limit).Skip(skip).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyEventList(context.Background(), organisationId, policyId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyEventListRequest struct via the builder pattern


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


## IamOrganisationPolicyGet

> Policy IamOrganisationPolicyGet(ctx, organisationId, policyId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyGet(context.Background(), organisationId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyGet`: Policy
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPolicyList

> []Policy IamOrganisationPolicyList(ctx, organisationId).Name(name).Resource(resource).TagValue(tagValue).TagKey(tagKey).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    name := "name_example" // string | Filter by name (optional)
    resource := "resource_example" // string | Filter by resource (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyList(context.Background(), organisationId).Name(name).Resource(resource).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyList`: []Policy
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **resource** | **string** | Filter by resource | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPolicyServiceGet

> ResourceService IamOrganisationPolicyServiceGet(ctx, organisationId, policyId, serviceId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyServiceGet(context.Background(), organisationId, policyId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyServiceGetRequest struct via the builder pattern


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


## IamOrganisationPolicyServiceList

> []ResourceService IamOrganisationPolicyServiceList(ctx, organisationId, policyId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyServiceList(context.Background(), organisationId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyServiceListRequest struct via the builder pattern


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


## IamOrganisationPolicyTagCreate

> Tag IamOrganisationPolicyTagCreate(ctx, organisationId, policyId).Tag(tag).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyTagCreate(context.Background(), organisationId, policyId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyTagCreateRequest struct via the builder pattern


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


## IamOrganisationPolicyTagDelete

> IamOrganisationPolicyTagDelete(ctx, organisationId, policyId, tagId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyTagDelete(context.Background(), organisationId, policyId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyTagDeleteRequest struct via the builder pattern


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


## IamOrganisationPolicyTagGet

> Tag IamOrganisationPolicyTagGet(ctx, organisationId, policyId, tagId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyTagGet(context.Background(), organisationId, policyId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyTagGetRequest struct via the builder pattern


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


## IamOrganisationPolicyTagList

> []Tag IamOrganisationPolicyTagList(ctx, organisationId, policyId).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyTagList(context.Background(), organisationId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyTagListRequest struct via the builder pattern


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


## IamOrganisationPolicyTagPut

> []Tag IamOrganisationPolicyTagPut(ctx, organisationId, policyId).Tag(tag).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyTagPut(context.Background(), organisationId, policyId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyTagPutRequest struct via the builder pattern


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


## IamOrganisationPolicyUpdate

> Policy IamOrganisationPolicyUpdate(ctx, organisationId, policyId).IamProjectPolicyUpdate(iamProjectPolicyUpdate).Execute()

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
    organisationId := "organisationId_example" // string | Organisation Id
    policyId := "policyId_example" // string | Policy Id
    iamProjectPolicyUpdate := *openapiclient.NewIamProjectPolicyUpdate() // IamProjectPolicyUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamOrganisationPolicyApi.IamOrganisationPolicyUpdate(context.Background(), organisationId, policyId).IamProjectPolicyUpdate(iamProjectPolicyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationPolicyApi.IamOrganisationPolicyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPolicyUpdate`: Policy
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationPolicyApi.IamOrganisationPolicyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectPolicyUpdate** | [**IamProjectPolicyUpdate**](IamProjectPolicyUpdate.md) |  | 

### Return type

[**Policy**](policy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

