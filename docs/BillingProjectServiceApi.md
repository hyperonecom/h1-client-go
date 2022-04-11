# \BillingProjectServiceApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingProjectServiceGet**](BillingProjectServiceApi.md#BillingProjectServiceGet) | **Get** /billing/project/{projectId}/service/{serviceId} | Get billing/service
[**BillingProjectServiceList**](BillingProjectServiceApi.md#BillingProjectServiceList) | **Get** /billing/project/{projectId}/service | List billing/service



## BillingProjectServiceGet

> Service BillingProjectServiceGet(ctx, projectId, serviceId).Execute()

Get billing/service



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
    serviceId := "serviceId_example" // string | Service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingProjectServiceApi.BillingProjectServiceGet(context.Background(), projectId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectServiceApi.BillingProjectServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectServiceGet`: Service
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectServiceApi.BillingProjectServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**serviceId** | **string** | Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Service**](Service.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectServiceList

> []Service BillingProjectServiceList(ctx, projectId).Kind(kind).Name(name).Type_(type_).Execute()

List billing/service



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
    kind := "kind_example" // string | Filter by kind (optional)
    name := "name_example" // string | Filter by name (optional)
    type_ := "type__example" // string | Filter by type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingProjectServiceApi.BillingProjectServiceList(context.Background(), projectId).Kind(kind).Name(name).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectServiceApi.BillingProjectServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectServiceList`: []Service
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectServiceApi.BillingProjectServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kind** | **string** | Filter by kind | 
 **name** | **string** | Filter by name | 
 **type_** | **string** | Filter by type | 

### Return type

[**[]Service**](Service.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

