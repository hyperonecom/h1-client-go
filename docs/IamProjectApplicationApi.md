# \IamProjectApplicationApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectApplicationCreate**](IamProjectApplicationApi.md#IamProjectApplicationCreate) | **Post** /iam/project/{projectId}/application | Create iam/application
[**IamProjectApplicationCredentialCreate**](IamProjectApplicationApi.md#IamProjectApplicationCredentialCreate) | **Post** /iam/project/{projectId}/application/{applicationId}/credential | Create iam/application.credential
[**IamProjectApplicationCredentialDelete**](IamProjectApplicationApi.md#IamProjectApplicationCredentialDelete) | **Delete** /iam/project/{projectId}/application/{applicationId}/credential/{credentialId} | Delete iam/application.credential
[**IamProjectApplicationCredentialGet**](IamProjectApplicationApi.md#IamProjectApplicationCredentialGet) | **Get** /iam/project/{projectId}/application/{applicationId}/credential/{credentialId} | Get iam/application.credential
[**IamProjectApplicationCredentialList**](IamProjectApplicationApi.md#IamProjectApplicationCredentialList) | **Get** /iam/project/{projectId}/application/{applicationId}/credential | List iam/application.credential
[**IamProjectApplicationCredentialPatch**](IamProjectApplicationApi.md#IamProjectApplicationCredentialPatch) | **Patch** /iam/project/{projectId}/application/{applicationId}/credential/{credentialId} | Update iam/application.credential
[**IamProjectApplicationDelete**](IamProjectApplicationApi.md#IamProjectApplicationDelete) | **Delete** /iam/project/{projectId}/application/{applicationId} | Delete iam/application
[**IamProjectApplicationEventGet**](IamProjectApplicationApi.md#IamProjectApplicationEventGet) | **Get** /iam/project/{projectId}/application/{applicationId}/event/{eventId} | Get iam/application.event
[**IamProjectApplicationEventList**](IamProjectApplicationApi.md#IamProjectApplicationEventList) | **Get** /iam/project/{projectId}/application/{applicationId}/event | List iam/application.event
[**IamProjectApplicationGet**](IamProjectApplicationApi.md#IamProjectApplicationGet) | **Get** /iam/project/{projectId}/application/{applicationId} | Get iam/application
[**IamProjectApplicationList**](IamProjectApplicationApi.md#IamProjectApplicationList) | **Get** /iam/project/{projectId}/application | List iam/application
[**IamProjectApplicationRedirectCreate**](IamProjectApplicationApi.md#IamProjectApplicationRedirectCreate) | **Post** /iam/project/{projectId}/application/{applicationId}/redirect | Create iam/application.redirect
[**IamProjectApplicationRedirectDelete**](IamProjectApplicationApi.md#IamProjectApplicationRedirectDelete) | **Delete** /iam/project/{projectId}/application/{applicationId}/redirect/{redirectId} | Delete iam/application.redirect
[**IamProjectApplicationRedirectGet**](IamProjectApplicationApi.md#IamProjectApplicationRedirectGet) | **Get** /iam/project/{projectId}/application/{applicationId}/redirect/{redirectId} | Get iam/application.redirect
[**IamProjectApplicationRedirectList**](IamProjectApplicationApi.md#IamProjectApplicationRedirectList) | **Get** /iam/project/{projectId}/application/{applicationId}/redirect | List iam/application.redirect
[**IamProjectApplicationRedirectPatch**](IamProjectApplicationApi.md#IamProjectApplicationRedirectPatch) | **Patch** /iam/project/{projectId}/application/{applicationId}/redirect/{redirectId} | Update iam/application.redirect
[**IamProjectApplicationServiceGet**](IamProjectApplicationApi.md#IamProjectApplicationServiceGet) | **Get** /iam/project/{projectId}/application/{applicationId}/service/{serviceId} | Get iam/application.service
[**IamProjectApplicationServiceList**](IamProjectApplicationApi.md#IamProjectApplicationServiceList) | **Get** /iam/project/{projectId}/application/{applicationId}/service | List iam/application.service
[**IamProjectApplicationTagCreate**](IamProjectApplicationApi.md#IamProjectApplicationTagCreate) | **Post** /iam/project/{projectId}/application/{applicationId}/tag | Create iam/application.tag
[**IamProjectApplicationTagDelete**](IamProjectApplicationApi.md#IamProjectApplicationTagDelete) | **Delete** /iam/project/{projectId}/application/{applicationId}/tag/{tagId} | Delete iam/application.tag
[**IamProjectApplicationTagGet**](IamProjectApplicationApi.md#IamProjectApplicationTagGet) | **Get** /iam/project/{projectId}/application/{applicationId}/tag/{tagId} | Get iam/application.tag
[**IamProjectApplicationTagList**](IamProjectApplicationApi.md#IamProjectApplicationTagList) | **Get** /iam/project/{projectId}/application/{applicationId}/tag | List iam/application.tag
[**IamProjectApplicationTagPut**](IamProjectApplicationApi.md#IamProjectApplicationTagPut) | **Put** /iam/project/{projectId}/application/{applicationId}/tag | Replace iam/application.tag
[**IamProjectApplicationUpdate**](IamProjectApplicationApi.md#IamProjectApplicationUpdate) | **Patch** /iam/project/{projectId}/application/{applicationId} | Update iam/application



## IamProjectApplicationCreate

> Application IamProjectApplicationCreate(ctx, projectId).IamProjectApplicationCreate(iamProjectApplicationCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create iam/application



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
    iamProjectApplicationCreate := *openapiclient.NewIamProjectApplicationCreate("Name_example", "Service_example") // IamProjectApplicationCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationCreate(context.Background(), projectId).IamProjectApplicationCreate(iamProjectApplicationCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationCreate`: Application
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectApplicationCreate** | [**IamProjectApplicationCreate**](IamProjectApplicationCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Application**](Application.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationCredentialCreate

> ApplicationCredential IamProjectApplicationCredentialCreate(ctx, projectId, applicationId).ApplicationCredential(applicationCredential).Execute()

Create iam/application.credential



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
    applicationId := "applicationId_example" // string | Application Id
    applicationCredential := *openapiclient.NewApplicationCredential("Name_example", "Type_example", "Value_example") // ApplicationCredential | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationCredentialCreate(context.Background(), projectId, applicationId).ApplicationCredential(applicationCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationCredentialCreate`: ApplicationCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationCredential** | [**ApplicationCredential**](ApplicationCredential.md) |  | 

### Return type

[**ApplicationCredential**](ApplicationCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationCredentialDelete

> Application IamProjectApplicationCredentialDelete(ctx, projectId, applicationId, credentialId).Execute()

Delete iam/application.credential



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
    applicationId := "applicationId_example" // string | Application Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationCredentialDelete(context.Background(), projectId, applicationId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationCredentialDelete`: Application
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Application**](Application.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationCredentialGet

> ApplicationCredential IamProjectApplicationCredentialGet(ctx, projectId, applicationId, credentialId).Execute()

Get iam/application.credential



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
    applicationId := "applicationId_example" // string | Application Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationCredentialGet(context.Background(), projectId, applicationId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationCredentialGet`: ApplicationCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicationCredential**](ApplicationCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationCredentialList

> []ApplicationCredential IamProjectApplicationCredentialList(ctx, projectId, applicationId).Execute()

List iam/application.credential



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
    applicationId := "applicationId_example" // string | Application Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationCredentialList(context.Background(), projectId, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationCredentialList`: []ApplicationCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ApplicationCredential**](ApplicationCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationCredentialPatch

> ApplicationCredential IamProjectApplicationCredentialPatch(ctx, projectId, applicationId, credentialId).IamProjectApplicationCredentialPatch(iamProjectApplicationCredentialPatch).Execute()

Update iam/application.credential



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
    applicationId := "applicationId_example" // string | Application Id
    credentialId := "credentialId_example" // string | credentialId
    iamProjectApplicationCredentialPatch := *openapiclient.NewIamProjectApplicationCredentialPatch("Name_example") // IamProjectApplicationCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationCredentialPatch(context.Background(), projectId, applicationId, credentialId).IamProjectApplicationCredentialPatch(iamProjectApplicationCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationCredentialPatch`: ApplicationCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iamProjectApplicationCredentialPatch** | [**IamProjectApplicationCredentialPatch**](IamProjectApplicationCredentialPatch.md) |  | 

### Return type

[**ApplicationCredential**](ApplicationCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationDelete

> Application IamProjectApplicationDelete(ctx, projectId, applicationId).Execute()

Delete iam/application



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
    applicationId := "applicationId_example" // string | Application Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationDelete(context.Background(), projectId, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationDelete`: Application
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Application**](Application.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationEventGet

> Event IamProjectApplicationEventGet(ctx, projectId, applicationId, eventId).Execute()

Get iam/application.event



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
    applicationId := "applicationId_example" // string | Application Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationEventGet(context.Background(), projectId, applicationId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationEventGetRequest struct via the builder pattern


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


## IamProjectApplicationEventList

> []Event IamProjectApplicationEventList(ctx, projectId, applicationId).Limit(limit).Skip(skip).Execute()

List iam/application.event



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
    applicationId := "applicationId_example" // string | Application Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationEventList(context.Background(), projectId, applicationId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationEventListRequest struct via the builder pattern


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


## IamProjectApplicationGet

> Application IamProjectApplicationGet(ctx, projectId, applicationId).Execute()

Get iam/application



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
    applicationId := "applicationId_example" // string | Application Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationGet(context.Background(), projectId, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationGet`: Application
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Application**](Application.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationList

> []Application IamProjectApplicationList(ctx, projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List iam/application



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
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationList(context.Background(), projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationList`: []Application
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Application**](Application.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationRedirectCreate

> IamRedirect IamProjectApplicationRedirectCreate(ctx, projectId, applicationId).IamRedirect(iamRedirect).Execute()

Create iam/application.redirect



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
    applicationId := "applicationId_example" // string | Application Id
    iamRedirect := *openapiclient.NewIamRedirect("Name_example", "Value_example") // IamRedirect | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationRedirectCreate(context.Background(), projectId, applicationId).IamRedirect(iamRedirect).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationRedirectCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationRedirectCreate`: IamRedirect
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationRedirectCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationRedirectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamRedirect** | [**IamRedirect**](IamRedirect.md) |  | 

### Return type

[**IamRedirect**](IamRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationRedirectDelete

> Application IamProjectApplicationRedirectDelete(ctx, projectId, applicationId, redirectId).Execute()

Delete iam/application.redirect



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
    applicationId := "applicationId_example" // string | Application Id
    redirectId := "redirectId_example" // string | redirectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationRedirectDelete(context.Background(), projectId, applicationId, redirectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationRedirectDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationRedirectDelete`: Application
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationRedirectDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**redirectId** | **string** | redirectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationRedirectDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Application**](Application.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationRedirectGet

> IamRedirect IamProjectApplicationRedirectGet(ctx, projectId, applicationId, redirectId).Execute()

Get iam/application.redirect



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
    applicationId := "applicationId_example" // string | Application Id
    redirectId := "redirectId_example" // string | redirectId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationRedirectGet(context.Background(), projectId, applicationId, redirectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationRedirectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationRedirectGet`: IamRedirect
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationRedirectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**redirectId** | **string** | redirectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationRedirectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IamRedirect**](IamRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationRedirectList

> []IamRedirect IamProjectApplicationRedirectList(ctx, projectId, applicationId).Execute()

List iam/application.redirect



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
    applicationId := "applicationId_example" // string | Application Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationRedirectList(context.Background(), projectId, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationRedirectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationRedirectList`: []IamRedirect
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationRedirectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationRedirectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IamRedirect**](IamRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationRedirectPatch

> IamRedirect IamProjectApplicationRedirectPatch(ctx, projectId, applicationId, redirectId).IamProjectApplicationRedirectPatch(iamProjectApplicationRedirectPatch).Execute()

Update iam/application.redirect



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
    applicationId := "applicationId_example" // string | Application Id
    redirectId := "redirectId_example" // string | redirectId
    iamProjectApplicationRedirectPatch := *openapiclient.NewIamProjectApplicationRedirectPatch("Name_example") // IamProjectApplicationRedirectPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationRedirectPatch(context.Background(), projectId, applicationId, redirectId).IamProjectApplicationRedirectPatch(iamProjectApplicationRedirectPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationRedirectPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationRedirectPatch`: IamRedirect
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationRedirectPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**redirectId** | **string** | redirectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationRedirectPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iamProjectApplicationRedirectPatch** | [**IamProjectApplicationRedirectPatch**](IamProjectApplicationRedirectPatch.md) |  | 

### Return type

[**IamRedirect**](IamRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectApplicationServiceGet

> ResourceService IamProjectApplicationServiceGet(ctx, projectId, applicationId, serviceId).Execute()

Get iam/application.service



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
    applicationId := "applicationId_example" // string | Application Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationServiceGet(context.Background(), projectId, applicationId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationServiceGetRequest struct via the builder pattern


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


## IamProjectApplicationServiceList

> []ResourceService IamProjectApplicationServiceList(ctx, projectId, applicationId).Execute()

List iam/application.service



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
    applicationId := "applicationId_example" // string | Application Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationServiceList(context.Background(), projectId, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationServiceListRequest struct via the builder pattern


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


## IamProjectApplicationTagCreate

> Tag IamProjectApplicationTagCreate(ctx, projectId, applicationId).Tag(tag).Execute()

Create iam/application.tag



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
    applicationId := "applicationId_example" // string | Application Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationTagCreate(context.Background(), projectId, applicationId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationTagCreateRequest struct via the builder pattern


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


## IamProjectApplicationTagDelete

> IamProjectApplicationTagDelete(ctx, projectId, applicationId, tagId).Execute()

Delete iam/application.tag



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
    applicationId := "applicationId_example" // string | Application Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationTagDelete(context.Background(), projectId, applicationId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationTagDeleteRequest struct via the builder pattern


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


## IamProjectApplicationTagGet

> Tag IamProjectApplicationTagGet(ctx, projectId, applicationId, tagId).Execute()

Get iam/application.tag



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
    applicationId := "applicationId_example" // string | Application Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationTagGet(context.Background(), projectId, applicationId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationTagGetRequest struct via the builder pattern


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


## IamProjectApplicationTagList

> []Tag IamProjectApplicationTagList(ctx, projectId, applicationId).Execute()

List iam/application.tag



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
    applicationId := "applicationId_example" // string | Application Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationTagList(context.Background(), projectId, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationTagListRequest struct via the builder pattern


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


## IamProjectApplicationTagPut

> []Tag IamProjectApplicationTagPut(ctx, projectId, applicationId).Tag(tag).Execute()

Replace iam/application.tag



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
    applicationId := "applicationId_example" // string | Application Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationTagPut(context.Background(), projectId, applicationId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationTagPutRequest struct via the builder pattern


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


## IamProjectApplicationUpdate

> Application IamProjectApplicationUpdate(ctx, projectId, applicationId).IamProjectApplicationUpdate(iamProjectApplicationUpdate).Execute()

Update iam/application



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
    applicationId := "applicationId_example" // string | Application Id
    iamProjectApplicationUpdate := *openapiclient.NewIamProjectApplicationUpdate() // IamProjectApplicationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApplicationApi.IamProjectApplicationUpdate(context.Background(), projectId, applicationId).IamProjectApplicationUpdate(iamProjectApplicationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApplicationApi.IamProjectApplicationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectApplicationUpdate`: Application
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApplicationApi.IamProjectApplicationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectApplicationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectApplicationUpdate** | [**IamProjectApplicationUpdate**](IamProjectApplicationUpdate.md) |  | 

### Return type

[**Application**](Application.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

