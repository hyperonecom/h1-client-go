# \DefaultApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1BillingPlanCreate**](DefaultApi.md#V1BillingPlanCreate) | **Post** /billing/{locationId}/project/{projectId}/plan | long text
[**V1BillingPlanDelete**](DefaultApi.md#V1BillingPlanDelete) | **Delete** /billing/{locationId}/project/{projectId}/plan/{planId} | delete plan
[**V1BillingPlanGet**](DefaultApi.md#V1BillingPlanGet) | **Get** /billing/{locationId}/project/{projectId}/plan/{planId} | get plan
[**V1BillingPlanList**](DefaultApi.md#V1BillingPlanList) | **Get** /billing/{locationId}/project/{projectId}/plan | list plan
[**V1BillingPlanNameUpdate**](DefaultApi.md#V1BillingPlanNameUpdate) | **Patch** /billing/{locationId}/project/{projectId}/plan/{planId} | update plan.name
[**V1BillingPlanTagCreate**](DefaultApi.md#V1BillingPlanTagCreate) | **Post** /billing/{locationId}/project/{projectId}/plan/{planId}/tag | create plan.tag
[**V1BillingPlanTagList**](DefaultApi.md#V1BillingPlanTagList) | **Get** /billing/{locationId}/project/{projectId}/plan/{planId}/tag | list plan.tag
[**V1BillingPlanTagPut**](DefaultApi.md#V1BillingPlanTagPut) | **Put** /billing/{locationId}/project/{projectId}/plan/{planId}/tag | put plan.tag



## V1BillingPlanCreate

> InlineResponse2001 V1BillingPlanCreate(ctx, locationId, projectId).InlineObject(inlineObject).Execute()

long text



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id
    inlineObject := *openapiclient.NewInlineObject(*openapiclient.NewBillingLocationIdProjectProjectIdPlanProperties(), "Profile_example", "Name_example") // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanCreate(context.Background(), locationId, projectId).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanCreate`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPlanDelete

> map[string]interface{} V1BillingPlanDelete(ctx, locationId, projectId, planId).Execute()

delete plan



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id
    planId := "planId_example" // string | plan Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanDelete(context.Background(), locationId, projectId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 
**planId** | **string** | plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPlanGet

> InlineResponse2001 V1BillingPlanGet(ctx, locationId, projectId, planId).Execute()

get plan



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id
    planId := "planId_example" // string | plan Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanGet(context.Background(), locationId, projectId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 
**planId** | **string** | plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPlanList

> []InlineResponse2001 V1BillingPlanList(ctx, locationId, projectId).Execute()

list plan



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanList(context.Background(), locationId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanList`: []InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse2001**](InlineResponse2001.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPlanNameUpdate

> InlineResponse2001 V1BillingPlanNameUpdate(ctx, locationId, projectId, planId).InlineObject1(inlineObject1).Execute()

update plan.name



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id
    planId := "planId_example" // string | plan Id
    inlineObject1 := *openapiclient.NewInlineObject1("Name_example") // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanNameUpdate(context.Background(), locationId, projectId, planId).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanNameUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanNameUpdate`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanNameUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 
**planId** | **string** | plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanNameUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPlanTagCreate

> BillingLocationIdProjectProjectIdPlanTag V1BillingPlanTagCreate(ctx, locationId, projectId, planId).InlineObject2(inlineObject2).Execute()

create plan.tag



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id
    planId := "planId_example" // string | plan Id
    inlineObject2 := *openapiclient.NewInlineObject2("Key_example", "Value_example") // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanTagCreate(context.Background(), locationId, projectId, planId).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanTagCreate`: BillingLocationIdProjectProjectIdPlanTag
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 
**planId** | **string** | plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanTagCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**BillingLocationIdProjectProjectIdPlanTag**](BillingLocationIdProjectProjectIdPlanTag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPlanTagList

> []BillingLocationIdProjectProjectIdPlanTag V1BillingPlanTagList(ctx, locationId, projectId, planId).Execute()

list plan.tag



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id
    planId := "planId_example" // string | plan Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanTagList(context.Background(), locationId, projectId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanTagList`: []BillingLocationIdProjectProjectIdPlanTag
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 
**planId** | **string** | plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanTagListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]BillingLocationIdProjectProjectIdPlanTag**](BillingLocationIdProjectProjectIdPlanTag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BillingPlanTagPut

> []BillingLocationIdProjectProjectIdPlanTag V1BillingPlanTagPut(ctx, locationId, projectId, planId).BillingLocationIdProjectProjectIdPlanTag1(billingLocationIdProjectProjectIdPlanTag1).Execute()

put plan.tag



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
    locationId := "locationId_example" // string | location Id
    projectId := "projectId_example" // string | project Id
    planId := "planId_example" // string | plan Id
    billingLocationIdProjectProjectIdPlanTag1 := []openapiclient.BillingLocationIdProjectProjectIdPlanTag1{*openapiclient.NewBillingLocationIdProjectProjectIdPlanTag1("Key_example", "Value_example")} // []BillingLocationIdProjectProjectIdPlanTag1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1BillingPlanTagPut(context.Background(), locationId, projectId, planId).BillingLocationIdProjectProjectIdPlanTag1(billingLocationIdProjectProjectIdPlanTag1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1BillingPlanTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BillingPlanTagPut`: []BillingLocationIdProjectProjectIdPlanTag
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1BillingPlanTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | location Id | 
**projectId** | **string** | project Id | 
**planId** | **string** | plan Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BillingPlanTagPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **billingLocationIdProjectProjectIdPlanTag1** | [**[]BillingLocationIdProjectProjectIdPlanTag1**](BillingLocationIdProjectProjectIdPlanTag1.md) |  | 

### Return type

[**[]BillingLocationIdProjectProjectIdPlanTag**](BillingLocationIdProjectProjectIdPlanTag.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

