# \ContainerProjectRegistryApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContainerProjectRegistryCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCreate) | **Post** /container/{locationId}/project/{projectId}/registry | Create container/registry
[**ContainerProjectRegistryCredentialCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialCreate) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/credential | Create container/registry.credential
[**ContainerProjectRegistryCredentialDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/credential/{credentialId} | Delete container/registry.credential
[**ContainerProjectRegistryCredentialGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/credential/{credentialId} | Get container/registry.credential
[**ContainerProjectRegistryCredentialList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/credential | List container/registry.credential
[**ContainerProjectRegistryCredentialPatch**](ContainerProjectRegistryApi.md#ContainerProjectRegistryCredentialPatch) | **Patch** /container/{locationId}/project/{projectId}/registry/{registryId}/credential/{credentialId} | Update container/registry.credential
[**ContainerProjectRegistryDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId} | Delete container/registry
[**ContainerProjectRegistryDomainCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainCreate) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/domain | Create container/registry.domain
[**ContainerProjectRegistryDomainDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/domain/{domainId} | Delete container/registry.domain
[**ContainerProjectRegistryDomainGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/domain/{domainId} | Get container/registry.domain
[**ContainerProjectRegistryDomainList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryDomainList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/domain | List container/registry.domain
[**ContainerProjectRegistryEventGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryEventGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/event/{eventId} | Get container/registry.event
[**ContainerProjectRegistryEventList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryEventList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/event | List container/registry.event
[**ContainerProjectRegistryGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId} | Get container/registry
[**ContainerProjectRegistryList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryList) | **Get** /container/{locationId}/project/{projectId}/registry | List container/registry
[**ContainerProjectRegistryRepositoryGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId} | Get container/registry.repository
[**ContainerProjectRegistryRepositoryImageDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryImageDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId}/image/{imageId} | Delete container/registry.image
[**ContainerProjectRegistryRepositoryImageGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryImageGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId}/image/{imageId} | Get container/registry.image
[**ContainerProjectRegistryRepositoryImageList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryImageList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository/{repositoryId}/image | List container/registry.image
[**ContainerProjectRegistryRepositoryList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryRepositoryList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/repository | List container/registry.repository
[**ContainerProjectRegistryServiceGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryServiceGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/service/{serviceId} | Get container/registry.service
[**ContainerProjectRegistryServiceList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryServiceList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/service | List container/registry.service
[**ContainerProjectRegistryStart**](ContainerProjectRegistryApi.md#ContainerProjectRegistryStart) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/actions/start | Start container/registry
[**ContainerProjectRegistryStop**](ContainerProjectRegistryApi.md#ContainerProjectRegistryStop) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/actions/stop | Stop container/registry
[**ContainerProjectRegistryTagCreate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagCreate) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/tag | Create container/registry.tag
[**ContainerProjectRegistryTagDelete**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagDelete) | **Delete** /container/{locationId}/project/{projectId}/registry/{registryId}/tag/{tagId} | Delete container/registry.tag
[**ContainerProjectRegistryTagGet**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagGet) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/tag/{tagId} | Get container/registry.tag
[**ContainerProjectRegistryTagList**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagList) | **Get** /container/{locationId}/project/{projectId}/registry/{registryId}/tag | List container/registry.tag
[**ContainerProjectRegistryTagPut**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTagPut) | **Put** /container/{locationId}/project/{projectId}/registry/{registryId}/tag | Replace container/registry.tag
[**ContainerProjectRegistryTransfer**](ContainerProjectRegistryApi.md#ContainerProjectRegistryTransfer) | **Post** /container/{locationId}/project/{projectId}/registry/{registryId}/actions/transfer | Transfer container/registry
[**ContainerProjectRegistryUpdate**](ContainerProjectRegistryApi.md#ContainerProjectRegistryUpdate) | **Patch** /container/{locationId}/project/{projectId}/registry/{registryId} | Update container/registry



## ContainerProjectRegistryCreate

> Registry ContainerProjectRegistryCreate(ctx, projectId, locationId).ContainerProjectRegistryCreate(containerProjectRegistryCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create container/registry



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
    containerProjectRegistryCreate := *openapiclient.NewContainerProjectRegistryCreate("Name_example", "Service_example") // ContainerProjectRegistryCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryCreate(context.Background(), projectId, locationId).ContainerProjectRegistryCreate(containerProjectRegistryCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryCreate`: Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **containerProjectRegistryCreate** | [**ContainerProjectRegistryCreate**](ContainerProjectRegistryCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialCreate

> RegistryCredential ContainerProjectRegistryCredentialCreate(ctx, projectId, locationId, registryId).RegistryCredential(registryCredential).Execute()

Create container/registry.credential



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
    registryId := "registryId_example" // string | Registry Id
    registryCredential := *openapiclient.NewRegistryCredential("Name_example", "Type_example", "Value_example") // RegistryCredential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryCredentialCreate(context.Background(), projectId, locationId, registryId).RegistryCredential(registryCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryCredentialCreate`: RegistryCredential
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **registryCredential** | [**RegistryCredential**](RegistryCredential.md) |  | 

### Return type

[**RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialDelete

> Registry ContainerProjectRegistryCredentialDelete(ctx, projectId, locationId, registryId, credentialId).Execute()

Delete container/registry.credential



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
    registryId := "registryId_example" // string | Registry Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryCredentialDelete(context.Background(), projectId, locationId, registryId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryCredentialDelete`: Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialGet

> RegistryCredential ContainerProjectRegistryCredentialGet(ctx, projectId, locationId, registryId, credentialId).Execute()

Get container/registry.credential



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
    registryId := "registryId_example" // string | Registry Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryCredentialGet(context.Background(), projectId, locationId, registryId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryCredentialGet`: RegistryCredential
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialList

> []RegistryCredential ContainerProjectRegistryCredentialList(ctx, projectId, locationId, registryId).Execute()

List container/registry.credential



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
    registryId := "registryId_example" // string | Registry Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryCredentialList(context.Background(), projectId, locationId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryCredentialList`: []RegistryCredential
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryCredentialPatch

> RegistryCredential ContainerProjectRegistryCredentialPatch(ctx, projectId, locationId, registryId, credentialId).ContainerProjectRegistryCredentialPatch(containerProjectRegistryCredentialPatch).Execute()

Update container/registry.credential



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
    registryId := "registryId_example" // string | Registry Id
    credentialId := "credentialId_example" // string | credentialId
    containerProjectRegistryCredentialPatch := *openapiclient.NewContainerProjectRegistryCredentialPatch("Name_example") // ContainerProjectRegistryCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryCredentialPatch(context.Background(), projectId, locationId, registryId, credentialId).ContainerProjectRegistryCredentialPatch(containerProjectRegistryCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryCredentialPatch`: RegistryCredential
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **containerProjectRegistryCredentialPatch** | [**ContainerProjectRegistryCredentialPatch**](ContainerProjectRegistryCredentialPatch.md) |  | 

### Return type

[**RegistryCredential**](registry.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryDelete

> ContainerProjectRegistryDelete(ctx, projectId, locationId, registryId).Execute()

Delete container/registry



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
    registryId := "registryId_example" // string | Registry Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryDelete(context.Background(), projectId, locationId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryDelete``: %v\n", err)
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
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryDeleteRequest struct via the builder pattern


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


## ContainerProjectRegistryDomainCreate

> Domain ContainerProjectRegistryDomainCreate(ctx, projectId, locationId, registryId).Domain(domain).Execute()

Create container/registry.domain



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
    registryId := "registryId_example" // string | Registry Id
    domain := *openapiclient.NewDomain("Value_example") // Domain | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryDomainCreate(context.Background(), projectId, locationId, registryId).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryDomainCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryDomainCreate`: Domain
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryDomainCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **domain** | [**Domain**](Domain.md) |  | 

### Return type

[**Domain**](domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryDomainDelete

> ContainerProjectRegistryDomainDelete(ctx, projectId, locationId, registryId, domainId).Execute()

Delete container/registry.domain



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
    registryId := "registryId_example" // string | Registry Id
    domainId := "domainId_example" // string | domainId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryDomainDelete(context.Background(), projectId, locationId, registryId, domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryDomainDelete``: %v\n", err)
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
**registryId** | **string** | Registry Id | 
**domainId** | **string** | domainId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryDomainDeleteRequest struct via the builder pattern


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


## ContainerProjectRegistryDomainGet

> Domain ContainerProjectRegistryDomainGet(ctx, projectId, locationId, registryId, domainId).Execute()

Get container/registry.domain



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
    registryId := "registryId_example" // string | Registry Id
    domainId := "domainId_example" // string | domainId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryDomainGet(context.Background(), projectId, locationId, registryId, domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryDomainGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryDomainGet`: Domain
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryDomainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**domainId** | **string** | domainId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryDomainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Domain**](domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryDomainList

> []Domain ContainerProjectRegistryDomainList(ctx, projectId, locationId, registryId).Execute()

List container/registry.domain



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
    registryId := "registryId_example" // string | Registry Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryDomainList(context.Background(), projectId, locationId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryDomainList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryDomainList`: []Domain
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryDomainList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryDomainListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Domain**](domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryEventGet

> Event ContainerProjectRegistryEventGet(ctx, projectId, locationId, registryId, eventId).Execute()

Get container/registry.event



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
    registryId := "registryId_example" // string | Registry Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryEventGet(context.Background(), projectId, locationId, registryId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryEventGetRequest struct via the builder pattern


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


## ContainerProjectRegistryEventList

> []Event ContainerProjectRegistryEventList(ctx, projectId, locationId, registryId).Limit(limit).Skip(skip).Execute()

List container/registry.event



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
    registryId := "registryId_example" // string | Registry Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryEventList(context.Background(), projectId, locationId, registryId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryEventListRequest struct via the builder pattern


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


## ContainerProjectRegistryGet

> Registry ContainerProjectRegistryGet(ctx, projectId, locationId, registryId).Execute()

Get container/registry



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
    registryId := "registryId_example" // string | Registry Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryGet(context.Background(), projectId, locationId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryGet`: Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryList

> []Registry ContainerProjectRegistryList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List container/registry



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
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryList`: []Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryRepositoryGet

> ContainerRepository ContainerProjectRegistryRepositoryGet(ctx, projectId, locationId, registryId, repositoryId).Execute()

Get container/registry.repository



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
    registryId := "registryId_example" // string | Registry Id
    repositoryId := "repositoryId_example" // string | repositoryId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryGet(context.Background(), projectId, locationId, registryId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryRepositoryGet`: ContainerRepository
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**repositoryId** | **string** | repositoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryRepositoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ContainerRepository**](container.repository.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryRepositoryImageDelete

> ContainerProjectRegistryRepositoryImageDelete(ctx, projectId, locationId, registryId, repositoryId, imageId).Execute()

Delete container/registry.image



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
    registryId := "registryId_example" // string | Registry Id
    repositoryId := "repositoryId_example" // string | repositoryId
    imageId := "imageId_example" // string | imageId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageDelete(context.Background(), projectId, locationId, registryId, repositoryId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageDelete``: %v\n", err)
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
**registryId** | **string** | Registry Id | 
**repositoryId** | **string** | repositoryId | 
**imageId** | **string** | imageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryRepositoryImageDeleteRequest struct via the builder pattern


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


## ContainerProjectRegistryRepositoryImageGet

> ContainerImage ContainerProjectRegistryRepositoryImageGet(ctx, projectId, locationId, registryId, repositoryId, imageId).Execute()

Get container/registry.image



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
    registryId := "registryId_example" // string | Registry Id
    repositoryId := "repositoryId_example" // string | repositoryId
    imageId := "imageId_example" // string | imageId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageGet(context.Background(), projectId, locationId, registryId, repositoryId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryRepositoryImageGet`: ContainerImage
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**repositoryId** | **string** | repositoryId | 
**imageId** | **string** | imageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryRepositoryImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**ContainerImage**](container.image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryRepositoryImageList

> []ContainerImage ContainerProjectRegistryRepositoryImageList(ctx, projectId, locationId, registryId, repositoryId).Execute()

List container/registry.image



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
    registryId := "registryId_example" // string | Registry Id
    repositoryId := "repositoryId_example" // string | repositoryId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageList(context.Background(), projectId, locationId, registryId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryRepositoryImageList`: []ContainerImage
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryImageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**repositoryId** | **string** | repositoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryRepositoryImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]ContainerImage**](container.image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryRepositoryList

> []ContainerRepository ContainerProjectRegistryRepositoryList(ctx, projectId, locationId, registryId).Execute()

List container/registry.repository



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
    registryId := "registryId_example" // string | Registry Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryList(context.Background(), projectId, locationId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryRepositoryList`: []ContainerRepository
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryRepositoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryRepositoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ContainerRepository**](container.repository.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryServiceGet

> ResourceService ContainerProjectRegistryServiceGet(ctx, projectId, locationId, registryId, serviceId).Execute()

Get container/registry.service



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
    registryId := "registryId_example" // string | Registry Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryServiceGet(context.Background(), projectId, locationId, registryId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryServiceGetRequest struct via the builder pattern


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


## ContainerProjectRegistryServiceList

> []ResourceService ContainerProjectRegistryServiceList(ctx, projectId, locationId, registryId).Execute()

List container/registry.service



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
    registryId := "registryId_example" // string | Registry Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryServiceList(context.Background(), projectId, locationId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryServiceListRequest struct via the builder pattern


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


## ContainerProjectRegistryStart

> Registry ContainerProjectRegistryStart(ctx, projectId, locationId, registryId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Start container/registry



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
    registryId := "registryId_example" // string | Registry Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryStart(context.Background(), projectId, locationId, registryId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryStart`: Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryStop

> Registry ContainerProjectRegistryStop(ctx, projectId, locationId, registryId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Stop container/registry



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
    registryId := "registryId_example" // string | Registry Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryStop(context.Background(), projectId, locationId, registryId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryStop`: Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryTagCreate

> Tag ContainerProjectRegistryTagCreate(ctx, projectId, locationId, registryId).Tag(tag).Execute()

Create container/registry.tag



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
    registryId := "registryId_example" // string | Registry Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryTagCreate(context.Background(), projectId, locationId, registryId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryTagCreateRequest struct via the builder pattern


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


## ContainerProjectRegistryTagDelete

> ContainerProjectRegistryTagDelete(ctx, projectId, locationId, registryId, tagId).Execute()

Delete container/registry.tag



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
    registryId := "registryId_example" // string | Registry Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryTagDelete(context.Background(), projectId, locationId, registryId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryTagDelete``: %v\n", err)
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
**registryId** | **string** | Registry Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryTagDeleteRequest struct via the builder pattern


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


## ContainerProjectRegistryTagGet

> Tag ContainerProjectRegistryTagGet(ctx, projectId, locationId, registryId, tagId).Execute()

Get container/registry.tag



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
    registryId := "registryId_example" // string | Registry Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryTagGet(context.Background(), projectId, locationId, registryId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryTagGetRequest struct via the builder pattern


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


## ContainerProjectRegistryTagList

> []Tag ContainerProjectRegistryTagList(ctx, projectId, locationId, registryId).Execute()

List container/registry.tag



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
    registryId := "registryId_example" // string | Registry Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryTagList(context.Background(), projectId, locationId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryTagListRequest struct via the builder pattern


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


## ContainerProjectRegistryTagPut

> []Tag ContainerProjectRegistryTagPut(ctx, projectId, locationId, registryId).Tag(tag).Execute()

Replace container/registry.tag



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
    registryId := "registryId_example" // string | Registry Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryTagPut(context.Background(), projectId, locationId, registryId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryTagPutRequest struct via the builder pattern


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


## ContainerProjectRegistryTransfer

> Registry ContainerProjectRegistryTransfer(ctx, projectId, locationId, registryId).ContainerProjectRegistryTransfer(containerProjectRegistryTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer container/registry



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
    registryId := "registryId_example" // string | Registry Id
    containerProjectRegistryTransfer := *openapiclient.NewContainerProjectRegistryTransfer("Project_example") // ContainerProjectRegistryTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryTransfer(context.Background(), projectId, locationId, registryId).ContainerProjectRegistryTransfer(containerProjectRegistryTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryTransfer`: Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **containerProjectRegistryTransfer** | [**ContainerProjectRegistryTransfer**](ContainerProjectRegistryTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContainerProjectRegistryUpdate

> Registry ContainerProjectRegistryUpdate(ctx, projectId, locationId, registryId).ContainerProjectRegistryUpdate(containerProjectRegistryUpdate).Execute()

Update container/registry



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
    registryId := "registryId_example" // string | Registry Id
    containerProjectRegistryUpdate := *openapiclient.NewContainerProjectRegistryUpdate() // ContainerProjectRegistryUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContainerProjectRegistryApi.ContainerProjectRegistryUpdate(context.Background(), projectId, locationId, registryId).ContainerProjectRegistryUpdate(containerProjectRegistryUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerProjectRegistryApi.ContainerProjectRegistryUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContainerProjectRegistryUpdate`: Registry
    fmt.Fprintf(os.Stdout, "Response from `ContainerProjectRegistryApi.ContainerProjectRegistryUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**registryId** | **string** | Registry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerProjectRegistryUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **containerProjectRegistryUpdate** | [**ContainerProjectRegistryUpdate**](ContainerProjectRegistryUpdate.md) |  | 

### Return type

[**Registry**](registry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

