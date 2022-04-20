# \IamProjectSaApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectSaCreate**](IamProjectSaApi.md#IamProjectSaCreate) | **Post** /iam/project/{projectId}/sa | Create iam/sa
[**IamProjectSaCredentialCreate**](IamProjectSaApi.md#IamProjectSaCredentialCreate) | **Post** /iam/project/{projectId}/sa/{saId}/credential | Create iam/sa.credential
[**IamProjectSaCredentialDelete**](IamProjectSaApi.md#IamProjectSaCredentialDelete) | **Delete** /iam/project/{projectId}/sa/{saId}/credential/{credentialId} | Delete iam/sa.credential
[**IamProjectSaCredentialGet**](IamProjectSaApi.md#IamProjectSaCredentialGet) | **Get** /iam/project/{projectId}/sa/{saId}/credential/{credentialId} | Get iam/sa.credential
[**IamProjectSaCredentialList**](IamProjectSaApi.md#IamProjectSaCredentialList) | **Get** /iam/project/{projectId}/sa/{saId}/credential | List iam/sa.credential
[**IamProjectSaCredentialPatch**](IamProjectSaApi.md#IamProjectSaCredentialPatch) | **Patch** /iam/project/{projectId}/sa/{saId}/credential/{credentialId} | Update iam/sa.credential
[**IamProjectSaDelete**](IamProjectSaApi.md#IamProjectSaDelete) | **Delete** /iam/project/{projectId}/sa/{saId} | Delete iam/sa
[**IamProjectSaEventGet**](IamProjectSaApi.md#IamProjectSaEventGet) | **Get** /iam/project/{projectId}/sa/{saId}/event/{eventId} | Get iam/sa.event
[**IamProjectSaEventList**](IamProjectSaApi.md#IamProjectSaEventList) | **Get** /iam/project/{projectId}/sa/{saId}/event | List iam/sa.event
[**IamProjectSaGet**](IamProjectSaApi.md#IamProjectSaGet) | **Get** /iam/project/{projectId}/sa/{saId} | Get iam/sa
[**IamProjectSaList**](IamProjectSaApi.md#IamProjectSaList) | **Get** /iam/project/{projectId}/sa | List iam/sa
[**IamProjectSaServiceGet**](IamProjectSaApi.md#IamProjectSaServiceGet) | **Get** /iam/project/{projectId}/sa/{saId}/service/{serviceId} | Get iam/sa.service
[**IamProjectSaServiceList**](IamProjectSaApi.md#IamProjectSaServiceList) | **Get** /iam/project/{projectId}/sa/{saId}/service | List iam/sa.service
[**IamProjectSaTagCreate**](IamProjectSaApi.md#IamProjectSaTagCreate) | **Post** /iam/project/{projectId}/sa/{saId}/tag | Create iam/sa.tag
[**IamProjectSaTagDelete**](IamProjectSaApi.md#IamProjectSaTagDelete) | **Delete** /iam/project/{projectId}/sa/{saId}/tag/{tagId} | Delete iam/sa.tag
[**IamProjectSaTagGet**](IamProjectSaApi.md#IamProjectSaTagGet) | **Get** /iam/project/{projectId}/sa/{saId}/tag/{tagId} | Get iam/sa.tag
[**IamProjectSaTagList**](IamProjectSaApi.md#IamProjectSaTagList) | **Get** /iam/project/{projectId}/sa/{saId}/tag | List iam/sa.tag
[**IamProjectSaTagPut**](IamProjectSaApi.md#IamProjectSaTagPut) | **Put** /iam/project/{projectId}/sa/{saId}/tag | Replace iam/sa.tag
[**IamProjectSaUpdate**](IamProjectSaApi.md#IamProjectSaUpdate) | **Patch** /iam/project/{projectId}/sa/{saId} | Update iam/sa



## IamProjectSaCreate

> Sa IamProjectSaCreate(ctx, projectId).IamProjectSaCreate(iamProjectSaCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create iam/sa



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
    iamProjectSaCreate := *openapiclient.NewIamProjectSaCreate("Name_example") // IamProjectSaCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaCreate(context.Background(), projectId).IamProjectSaCreate(iamProjectSaCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaCreate`: Sa
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectSaCreate** | [**IamProjectSaCreate**](IamProjectSaCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Sa**](Sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialCreate

> SaCredential IamProjectSaCredentialCreate(ctx, projectId, saId).SaCredential(saCredential).Execute()

Create iam/sa.credential



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
    saId := "saId_example" // string | Sa Id
    saCredential := *openapiclient.NewSaCredential("Name_example", "Type_example", "Value_example") // SaCredential | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaCredentialCreate(context.Background(), projectId, saId).SaCredential(saCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaCredentialCreate`: SaCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **saCredential** | [**SaCredential**](SaCredential.md) |  | 

### Return type

[**SaCredential**](SaCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialDelete

> Sa IamProjectSaCredentialDelete(ctx, projectId, saId, credentialId).Execute()

Delete iam/sa.credential



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
    saId := "saId_example" // string | Sa Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaCredentialDelete(context.Background(), projectId, saId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaCredentialDelete`: Sa
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Sa**](Sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialGet

> SaCredential IamProjectSaCredentialGet(ctx, projectId, saId, credentialId).Execute()

Get iam/sa.credential



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
    saId := "saId_example" // string | Sa Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaCredentialGet(context.Background(), projectId, saId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaCredentialGet`: SaCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SaCredential**](SaCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialList

> []SaCredential IamProjectSaCredentialList(ctx, projectId, saId).Execute()

List iam/sa.credential



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
    saId := "saId_example" // string | Sa Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaCredentialList(context.Background(), projectId, saId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaCredentialList`: []SaCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SaCredential**](SaCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaCredentialPatch

> SaCredential IamProjectSaCredentialPatch(ctx, projectId, saId, credentialId).IamProjectSaCredentialPatch(iamProjectSaCredentialPatch).Execute()

Update iam/sa.credential



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
    saId := "saId_example" // string | Sa Id
    credentialId := "credentialId_example" // string | credentialId
    iamProjectSaCredentialPatch := *openapiclient.NewIamProjectSaCredentialPatch("Name_example") // IamProjectSaCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaCredentialPatch(context.Background(), projectId, saId, credentialId).IamProjectSaCredentialPatch(iamProjectSaCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaCredentialPatch`: SaCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iamProjectSaCredentialPatch** | [**IamProjectSaCredentialPatch**](IamProjectSaCredentialPatch.md) |  | 

### Return type

[**SaCredential**](SaCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaDelete

> Sa IamProjectSaDelete(ctx, projectId, saId).Execute()

Delete iam/sa



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
    saId := "saId_example" // string | Sa Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaDelete(context.Background(), projectId, saId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaDelete`: Sa
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sa**](Sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaEventGet

> Event IamProjectSaEventGet(ctx, projectId, saId, eventId).Execute()

Get iam/sa.event



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
    saId := "saId_example" // string | Sa Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaEventGet(context.Background(), projectId, saId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaEventGetRequest struct via the builder pattern


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


## IamProjectSaEventList

> []Event IamProjectSaEventList(ctx, projectId, saId).Limit(limit).Skip(skip).Execute()

List iam/sa.event



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
    saId := "saId_example" // string | Sa Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaEventList(context.Background(), projectId, saId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaEventListRequest struct via the builder pattern


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


## IamProjectSaGet

> Sa IamProjectSaGet(ctx, projectId, saId).Execute()

Get iam/sa



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
    saId := "saId_example" // string | Sa Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaGet(context.Background(), projectId, saId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaGet`: Sa
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sa**](Sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaList

> []Sa IamProjectSaList(ctx, projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List iam/sa



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
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaList(context.Background(), projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaList`: []Sa
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Sa**](Sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectSaServiceGet

> ResourceService IamProjectSaServiceGet(ctx, projectId, saId, serviceId).Execute()

Get iam/sa.service



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
    saId := "saId_example" // string | Sa Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaServiceGet(context.Background(), projectId, saId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaServiceGetRequest struct via the builder pattern


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


## IamProjectSaServiceList

> []ResourceService IamProjectSaServiceList(ctx, projectId, saId).Execute()

List iam/sa.service



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
    saId := "saId_example" // string | Sa Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaServiceList(context.Background(), projectId, saId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaServiceListRequest struct via the builder pattern


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


## IamProjectSaTagCreate

> Tag IamProjectSaTagCreate(ctx, projectId, saId).Tag(tag).Execute()

Create iam/sa.tag



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
    saId := "saId_example" // string | Sa Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaTagCreate(context.Background(), projectId, saId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaTagCreateRequest struct via the builder pattern


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


## IamProjectSaTagDelete

> IamProjectSaTagDelete(ctx, projectId, saId, tagId).Execute()

Delete iam/sa.tag



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
    saId := "saId_example" // string | Sa Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaTagDelete(context.Background(), projectId, saId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaTagDeleteRequest struct via the builder pattern


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


## IamProjectSaTagGet

> Tag IamProjectSaTagGet(ctx, projectId, saId, tagId).Execute()

Get iam/sa.tag



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
    saId := "saId_example" // string | Sa Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaTagGet(context.Background(), projectId, saId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaTagGetRequest struct via the builder pattern


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


## IamProjectSaTagList

> []Tag IamProjectSaTagList(ctx, projectId, saId).Execute()

List iam/sa.tag



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
    saId := "saId_example" // string | Sa Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaTagList(context.Background(), projectId, saId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaTagListRequest struct via the builder pattern


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


## IamProjectSaTagPut

> []Tag IamProjectSaTagPut(ctx, projectId, saId).Tag(tag).Execute()

Replace iam/sa.tag



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
    saId := "saId_example" // string | Sa Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaTagPut(context.Background(), projectId, saId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaTagPutRequest struct via the builder pattern


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


## IamProjectSaUpdate

> Sa IamProjectSaUpdate(ctx, projectId, saId).IamProjectSaUpdate(iamProjectSaUpdate).Execute()

Update iam/sa



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
    saId := "saId_example" // string | Sa Id
    iamProjectSaUpdate := *openapiclient.NewIamProjectSaUpdate() // IamProjectSaUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectSaApi.IamProjectSaUpdate(context.Background(), projectId, saId).IamProjectSaUpdate(iamProjectSaUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectSaApi.IamProjectSaUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectSaUpdate`: Sa
    fmt.Fprintf(os.Stdout, "Response from `IamProjectSaApi.IamProjectSaUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**saId** | **string** | Sa Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectSaUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectSaUpdate** | [**IamProjectSaUpdate**](IamProjectSaUpdate.md) |  | 

### Return type

[**Sa**](Sa.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

