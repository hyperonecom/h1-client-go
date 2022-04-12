# \IamUserApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamUserAuthorizationCode**](IamUserApi.md#IamUserAuthorizationCode) | **Post** /iam/user/{userId}/authorization/{authorizationId}/actions/code | Code iam/user.authorization
[**IamUserAuthorizationCreate**](IamUserApi.md#IamUserAuthorizationCreate) | **Post** /iam/user/{userId}/authorization | Create iam/user.authorization
[**IamUserAuthorizationDelete**](IamUserApi.md#IamUserAuthorizationDelete) | **Delete** /iam/user/{userId}/authorization/{authorizationId} | Delete iam/user.authorization
[**IamUserAuthorizationGet**](IamUserApi.md#IamUserAuthorizationGet) | **Get** /iam/user/{userId}/authorization/{authorizationId} | Get iam/user.authorization
[**IamUserAuthorizationList**](IamUserApi.md#IamUserAuthorizationList) | **Get** /iam/user/{userId}/authorization | List iam/user.authorization
[**IamUserCredentialAuthtokenDelete**](IamUserApi.md#IamUserCredentialAuthtokenDelete) | **Delete** /iam/user/{userId}/credential/authtoken/{authtokenId} | Delete iam/user.credential
[**IamUserCredentialAuthtokenGet**](IamUserApi.md#IamUserCredentialAuthtokenGet) | **Get** /iam/user/{userId}/credential/authtoken/{authtokenId} | Get iam/user.credential
[**IamUserCredentialAuthtokenList**](IamUserApi.md#IamUserCredentialAuthtokenList) | **Get** /iam/user/{userId}/credential/authtoken | List iam/user.credential
[**IamUserCredentialCreate**](IamUserApi.md#IamUserCredentialCreate) | **Post** /iam/user/{userId}/credential | Create iam/user.credential
[**IamUserCredentialDelete**](IamUserApi.md#IamUserCredentialDelete) | **Delete** /iam/user/{userId}/credential/{credentialId} | Delete iam/user.credential
[**IamUserCredentialGet**](IamUserApi.md#IamUserCredentialGet) | **Get** /iam/user/{userId}/credential/{credentialId} | Get iam/user.credential
[**IamUserCredentialList**](IamUserApi.md#IamUserCredentialList) | **Get** /iam/user/{userId}/credential | List iam/user.credential
[**IamUserCredentialPatch**](IamUserApi.md#IamUserCredentialPatch) | **Patch** /iam/user/{userId}/credential/{credentialId} | Update iam/user.credential
[**IamUserDelete**](IamUserApi.md#IamUserDelete) | **Delete** /iam/user/{userId} | Delete iam/user
[**IamUserGet**](IamUserApi.md#IamUserGet) | **Get** /iam/user/{userId} | Get iam/user
[**IamUserServiceGet**](IamUserApi.md#IamUserServiceGet) | **Get** /iam/user/{userId}/service/{serviceId} | Get iam/user.service
[**IamUserServiceList**](IamUserApi.md#IamUserServiceList) | **Get** /iam/user/{userId}/service | List iam/user.service
[**IamUserUpdate**](IamUserApi.md#IamUserUpdate) | **Patch** /iam/user/{userId} | Update iam/user



## IamUserAuthorizationCode

> InlineResponse200 IamUserAuthorizationCode(ctx, userId, authorizationId).IamUserAuthorizationCode(iamUserAuthorizationCode).Execute()

Code iam/user.authorization



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
    authorizationId := "authorizationId_example" // string | authorizationId
    iamUserAuthorizationCode := *openapiclient.NewIamUserAuthorizationCode("Redirect_example") // IamUserAuthorizationCode | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserAuthorizationCode(context.Background(), userId, authorizationId).IamUserAuthorizationCode(iamUserAuthorizationCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserAuthorizationCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserAuthorizationCode`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**authorizationId** | **string** | authorizationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamUserAuthorizationCode** | [**IamUserAuthorizationCode**](IamUserAuthorizationCode.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserAuthorizationCreate

> IamUserAuthorization IamUserAuthorizationCreate(ctx, userId).IamUserAuthorization(iamUserAuthorization).Execute()

Create iam/user.authorization



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
    iamUserAuthorization := *openapiclient.NewIamUserAuthorization("Application_example", "Name_example") // IamUserAuthorization | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserAuthorizationCreate(context.Background(), userId).IamUserAuthorization(iamUserAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserAuthorizationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserAuthorizationCreate`: IamUserAuthorization
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserAuthorizationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserAuthorizationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamUserAuthorization** | [**IamUserAuthorization**](IamUserAuthorization.md) |  | 

### Return type

[**IamUserAuthorization**](IamUserAuthorization.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserAuthorizationDelete

> User IamUserAuthorizationDelete(ctx, userId, authorizationId).Execute()

Delete iam/user.authorization



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
    authorizationId := "authorizationId_example" // string | authorizationId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserAuthorizationDelete(context.Background(), userId, authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserAuthorizationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserAuthorizationDelete`: User
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserAuthorizationDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**authorizationId** | **string** | authorizationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserAuthorizationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserAuthorizationGet

> IamUserAuthorization IamUserAuthorizationGet(ctx, userId, authorizationId).Execute()

Get iam/user.authorization



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
    authorizationId := "authorizationId_example" // string | authorizationId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserAuthorizationGet(context.Background(), userId, authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserAuthorizationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserAuthorizationGet`: IamUserAuthorization
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserAuthorizationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**authorizationId** | **string** | authorizationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserAuthorizationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamUserAuthorization**](IamUserAuthorization.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserAuthorizationList

> []IamUserAuthorization IamUserAuthorizationList(ctx, userId).Execute()

List iam/user.authorization



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserAuthorizationList(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserAuthorizationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamUserAuthorizationList`: []IamUserAuthorization
    fmt.Fprintf(os.Stdout, "Response from `IamUserApi.IamUserAuthorizationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserAuthorizationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IamUserAuthorization**](IamUserAuthorization.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialAuthtokenDelete(context.Background(), userId, authtokenId).Execute()
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialAuthtokenGet(context.Background(), userId, authtokenId).Execute()
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

[**AuthToken**](AuthToken.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialAuthtokenList(context.Background(), userId).Execute()
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

[**[]AuthToken**](AuthToken.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialCreate(context.Background(), userId).UserCredential(userCredential).Execute()
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

[**UserCredential**](UserCredential.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialDelete(context.Background(), userId, credentialId).Execute()
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

[**User**](User.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialGet(context.Background(), userId, credentialId).Execute()
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

[**UserCredential**](UserCredential.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialList(context.Background(), userId).Execute()
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

[**[]UserCredential**](UserCredential.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserCredentialPatch(context.Background(), userId, credentialId).IamUserCredentialPatch(iamUserCredentialPatch).Execute()
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

[**UserCredential**](UserCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamUserDelete

> IamUserDelete(ctx, userId).Execute()

Delete iam/user



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserDelete(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamUserApi.IamUserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamUserDeleteRequest struct via the builder pattern


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserGet(context.Background(), userId).Execute()
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

[**User**](User.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserServiceGet(context.Background(), userId, serviceId).Execute()
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

[**ResourceService**](ResourceService.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserServiceList(context.Background(), userId).Execute()
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

[**[]ResourceService**](ResourceService.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamUserApi.IamUserUpdate(context.Background(), userId).IamUserUpdate(iamUserUpdate).Execute()
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

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

