# \IamUserApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamUserCredentialAuthtokenDelete**](IamUserApi.md#IamUserCredentialAuthtokenDelete) | **Delete** /iam/user/{userId}/credential/authtoken/{authtokenId} | Delete iam/user.credential
[**IamUserCredentialAuthtokenGet**](IamUserApi.md#IamUserCredentialAuthtokenGet) | **Get** /iam/user/{userId}/credential/authtoken/{authtokenId} | Get iam/user.credential
[**IamUserCredentialAuthtokenList**](IamUserApi.md#IamUserCredentialAuthtokenList) | **Get** /iam/user/{userId}/credential/authtoken | List iam/user.credential
[**IamUserCredentialCreate**](IamUserApi.md#IamUserCredentialCreate) | **Post** /iam/user/{userId}/credential | Create iam/user.credential
[**IamUserCredentialDelete**](IamUserApi.md#IamUserCredentialDelete) | **Delete** /iam/user/{userId}/credential/{credentialId} | Delete iam/user.credential
[**IamUserCredentialGet**](IamUserApi.md#IamUserCredentialGet) | **Get** /iam/user/{userId}/credential/{credentialId} | Get iam/user.credential
[**IamUserCredentialList**](IamUserApi.md#IamUserCredentialList) | **Get** /iam/user/{userId}/credential | List iam/user.credential
[**IamUserCredentialPatch**](IamUserApi.md#IamUserCredentialPatch) | **Patch** /iam/user/{userId}/credential/{credentialId} | Update iam/user.credential
[**IamUserGet**](IamUserApi.md#IamUserGet) | **Get** /iam/user/{userId} | Get iam/user
[**IamUserServiceGet**](IamUserApi.md#IamUserServiceGet) | **Get** /iam/user/{userId}/service/{serviceId} | Get iam/user.service
[**IamUserServiceList**](IamUserApi.md#IamUserServiceList) | **Get** /iam/user/{userId}/service | List iam/user.service
[**IamUserUpdate**](IamUserApi.md#IamUserUpdate) | **Patch** /iam/user/{userId} | Update iam/user



## IamUserCredentialAuthtokenDelete

> IamUserCredentialAuthtokenDelete(ctx, userId, authtokenId).Execute()

Delete iam/user.credential



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
    userId := "userId_example" // string | User Id
    authtokenId := "authtokenId_example" // string | authtokenId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialAuthtokenDelete(context.Background(), userId, authtokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialAuthtokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**authtokenId** | **string** | authtokenId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialAuthtokenDeleteRequest struct via the builder pattern


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


## IamUserCredentialAuthtokenGet

> AuthToken IamUserCredentialAuthtokenGet(ctx, userId, authtokenId).Execute()

Get iam/user.credential



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
    userId := "userId_example" // string | User Id
    authtokenId := "authtokenId_example" // string | authtokenId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialAuthtokenGet(context.Background(), userId, authtokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialAuthtokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserCredentialAuthtokenGet`: AuthToken
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserCredentialAuthtokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**authtokenId** | **string** | authtokenId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialAuthtokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthToken**](authToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserCredentialAuthtokenList

> []AuthToken IamUserCredentialAuthtokenList(ctx, userId).Execute()

List iam/user.credential



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
    userId := "userId_example" // string | User Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialAuthtokenList(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialAuthtokenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserCredentialAuthtokenList`: []AuthToken
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserCredentialAuthtokenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialAuthtokenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AuthToken**](authToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserCredentialCreate

> UserCredential IamUserCredentialCreate(ctx, userId).UserCredential(userCredential).Execute()

Create iam/user.credential



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
    userId := "userId_example" // string | User Id
    userCredential := *openapiclient.NewUserCredential("Name_example", "Type_example", "Value_example") // UserCredential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialCreate(context.Background(), userId).UserCredential(userCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserCredentialCreate`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCredential** | [**UserCredential**](UserCredential.md) |  | 

### Return type

[**UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserCredentialDelete

> User IamUserCredentialDelete(ctx, userId, credentialId).Execute()

Delete iam/user.credential



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
    userId := "userId_example" // string | User Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialDelete(context.Background(), userId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserCredentialDelete`: User
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](user.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserCredentialGet

> UserCredential IamUserCredentialGet(ctx, userId, credentialId).Execute()

Get iam/user.credential



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
    userId := "userId_example" // string | User Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialGet(context.Background(), userId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserCredentialGet`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserCredentialList

> []UserCredential IamUserCredentialList(ctx, userId).Execute()

List iam/user.credential



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
    userId := "userId_example" // string | User Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialList(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserCredentialList`: []UserCredential
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserCredentialPatch

> UserCredential IamUserCredentialPatch(ctx, userId, credentialId).IamUserCredentialPatch(iamUserCredentialPatch).Execute()

Update iam/user.credential



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
    userId := "userId_example" // string | User Id
    credentialId := "credentialId_example" // string | credentialId
    iamUserCredentialPatch := *openapiclient.NewIamUserCredentialPatch("Name_example") // IamUserCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserCredentialPatch(context.Background(), userId, credentialId).IamUserCredentialPatch(iamUserCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserCredentialPatch`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamUserCredentialPatch** | [**IamUserCredentialPatch**](IamUserCredentialPatch.md) |  | 

### Return type

[**UserCredential**](user.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserGet

> User IamUserGet(ctx, userId).Execute()

Get iam/user



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
    userId := "userId_example" // string | User Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserGet(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserGet`: User
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](user.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserServiceGet

> ResourceService IamUserServiceGet(ctx, userId, serviceId).Execute()

Get iam/user.service



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
    userId := "userId_example" // string | User Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserServiceGet(context.Background(), userId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserServiceGetRequest struct via the builder pattern


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


## IamUserServiceList

> []ResourceService IamUserServiceList(ctx, userId).Execute()

List iam/user.service



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
    userId := "userId_example" // string | User Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserServiceList(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserServiceListRequest struct via the builder pattern


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


## IamUserUpdate

> User IamUserUpdate(ctx, userId).IamUserUpdate(iamUserUpdate).Execute()

Update iam/user



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
    userId := "userId_example" // string | User Id
    iamUserUpdate := *openapiclient.NewIamUserUpdate() // IamUserUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IamUserApi.IamUserUpdate(context.Background(), userId).IamUserUpdate(iamUserUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamUserUpdate** | [**IamUserUpdate**](IamUserUpdate.md) |  | 

### Return type

[**User**](user.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

