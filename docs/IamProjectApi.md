# \IamProjectApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectBillingList**](IamProjectApi.md#IamProjectBillingList) | **Get** /iam/project/{projectId}/billing | List iam/project.billing
[**IamProjectCostGet**](IamProjectApi.md#IamProjectCostGet) | **Get** /iam/project/{projectId}/cost/{costId} | Get iam/project.cost
[**IamProjectCostList**](IamProjectApi.md#IamProjectCostList) | **Get** /iam/project/{projectId}/cost | List iam/project.cost
[**IamProjectCostSeriesList**](IamProjectApi.md#IamProjectCostSeriesList) | **Get** /iam/project/{projectId}/cost/{costId}/series | List iam/project.series
[**IamProjectCreate**](IamProjectApi.md#IamProjectCreate) | **Post** /iam/project | Create iam/project
[**IamProjectCredentialStoreCreate**](IamProjectApi.md#IamProjectCredentialStoreCreate) | **Post** /iam/project/{projectId}/credentialStore | Create iam/project.credentialStore
[**IamProjectCredentialStoreDelete**](IamProjectApi.md#IamProjectCredentialStoreDelete) | **Delete** /iam/project/{projectId}/credentialStore/{credentialStoreId} | Delete iam/project.credentialStore
[**IamProjectCredentialStoreGet**](IamProjectApi.md#IamProjectCredentialStoreGet) | **Get** /iam/project/{projectId}/credentialStore/{credentialStoreId} | Get iam/project.credentialStore
[**IamProjectCredentialStoreList**](IamProjectApi.md#IamProjectCredentialStoreList) | **Get** /iam/project/{projectId}/credentialStore | List iam/project.credentialStore
[**IamProjectCredentialStorePatch**](IamProjectApi.md#IamProjectCredentialStorePatch) | **Patch** /iam/project/{projectId}/credentialStore/{credentialStoreId} | Update iam/project.credentialStore
[**IamProjectDelete**](IamProjectApi.md#IamProjectDelete) | **Delete** /iam/project/{projectId} | Delete iam/project
[**IamProjectEventGet**](IamProjectApi.md#IamProjectEventGet) | **Get** /iam/project/{projectId}/event/{eventId} | Get iam/project.event
[**IamProjectEventList**](IamProjectApi.md#IamProjectEventList) | **Get** /iam/project/{projectId}/event | List iam/project.event
[**IamProjectGet**](IamProjectApi.md#IamProjectGet) | **Get** /iam/project/{projectId} | Get iam/project
[**IamProjectInvitationAccept**](IamProjectApi.md#IamProjectInvitationAccept) | **Post** /iam/project/{projectId}/invitation/{invitationId}/actions/accept | Accept iam/project.invitation
[**IamProjectInvitationDelete**](IamProjectApi.md#IamProjectInvitationDelete) | **Delete** /iam/project/{projectId}/invitation/{invitationId} | Delete iam/project.invitation
[**IamProjectInvitationGet**](IamProjectApi.md#IamProjectInvitationGet) | **Get** /iam/project/{projectId}/invitation/{invitationId} | Get iam/project.invitation
[**IamProjectInvitationList**](IamProjectApi.md#IamProjectInvitationList) | **Get** /iam/project/{projectId}/invitation | List iam/project.invitation
[**IamProjectInvoiceList**](IamProjectApi.md#IamProjectInvoiceList) | **Get** /iam/project/{projectId}/invoice | List iam/project.invoice
[**IamProjectList**](IamProjectApi.md#IamProjectList) | **Get** /iam/project | List iam/project
[**IamProjectOwnershipCreate**](IamProjectApi.md#IamProjectOwnershipCreate) | **Post** /iam/project/{projectId}/ownership | Create iam/project.ownership
[**IamProjectOwnershipDelete**](IamProjectApi.md#IamProjectOwnershipDelete) | **Delete** /iam/project/{projectId}/ownership/{ownershipId} | Delete iam/project.ownership
[**IamProjectOwnershipGet**](IamProjectApi.md#IamProjectOwnershipGet) | **Get** /iam/project/{projectId}/ownership/{ownershipId} | Get iam/project.ownership
[**IamProjectOwnershipList**](IamProjectApi.md#IamProjectOwnershipList) | **Get** /iam/project/{projectId}/ownership | List iam/project.ownership
[**IamProjectPaymentList**](IamProjectApi.md#IamProjectPaymentList) | **Get** /iam/project/{projectId}/payment | List iam/project.payment
[**IamProjectProformaList**](IamProjectApi.md#IamProjectProformaList) | **Get** /iam/project/{projectId}/proforma | List iam/project.proforma
[**IamProjectQuotaGet**](IamProjectApi.md#IamProjectQuotaGet) | **Get** /iam/project/{projectId}/quota/{quotaId} | Get iam/project.quota
[**IamProjectQuotaLimitPatch**](IamProjectApi.md#IamProjectQuotaLimitPatch) | **Patch** /iam/project/{projectId}/quota/{quotaId}/limit | Update iam/project.limit
[**IamProjectQuotaList**](IamProjectApi.md#IamProjectQuotaList) | **Get** /iam/project/{projectId}/quota | List iam/project.quota
[**IamProjectServiceGet**](IamProjectApi.md#IamProjectServiceGet) | **Get** /iam/project/{projectId}/service/{serviceId} | Get iam/project.service
[**IamProjectServiceList**](IamProjectApi.md#IamProjectServiceList) | **Get** /iam/project/{projectId}/service | List iam/project.service
[**IamProjectTagCreate**](IamProjectApi.md#IamProjectTagCreate) | **Post** /iam/project/{projectId}/tag | Create iam/project.tag
[**IamProjectTagDelete**](IamProjectApi.md#IamProjectTagDelete) | **Delete** /iam/project/{projectId}/tag/{tagId} | Delete iam/project.tag
[**IamProjectTagGet**](IamProjectApi.md#IamProjectTagGet) | **Get** /iam/project/{projectId}/tag/{tagId} | Get iam/project.tag
[**IamProjectTagList**](IamProjectApi.md#IamProjectTagList) | **Get** /iam/project/{projectId}/tag | List iam/project.tag
[**IamProjectTagPut**](IamProjectApi.md#IamProjectTagPut) | **Put** /iam/project/{projectId}/tag | Replace iam/project.tag
[**IamProjectThresholdCreate**](IamProjectApi.md#IamProjectThresholdCreate) | **Post** /iam/project/{projectId}/threshold | Create iam/project.threshold
[**IamProjectThresholdDelete**](IamProjectApi.md#IamProjectThresholdDelete) | **Delete** /iam/project/{projectId}/threshold/{thresholdId} | Delete iam/project.threshold
[**IamProjectThresholdGet**](IamProjectApi.md#IamProjectThresholdGet) | **Get** /iam/project/{projectId}/threshold/{thresholdId} | Get iam/project.threshold
[**IamProjectThresholdList**](IamProjectApi.md#IamProjectThresholdList) | **Get** /iam/project/{projectId}/threshold | List iam/project.threshold
[**IamProjectTransfer**](IamProjectApi.md#IamProjectTransfer) | **Post** /iam/project/{projectId}/actions/transfer | Transfer iam/project
[**IamProjectUpdate**](IamProjectApi.md#IamProjectUpdate) | **Patch** /iam/project/{projectId} | Update iam/project
[**IamProjectUsageGet**](IamProjectApi.md#IamProjectUsageGet) | **Get** /iam/project/{projectId}/usage/{usageId} | Get iam/project.usage
[**IamProjectUsageList**](IamProjectApi.md#IamProjectUsageList) | **Get** /iam/project/{projectId}/usage | List iam/project.usage
[**IamProjectUsageSeriesList**](IamProjectApi.md#IamProjectUsageSeriesList) | **Get** /iam/project/{projectId}/usage/{usageId}/series | List iam/project.series



## IamProjectBillingList

> []Billing IamProjectBillingList(ctx, projectId).Start(start).End(end).ResourceType(resourceType).Execute()

List iam/project.billing



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project Id
    start := time.Now() // time.Time | start (optional)
    end := time.Now() // time.Time | end (optional)
    resourceType := "resourceType_example" // string | resource.type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectBillingList(context.Background(), projectId).Start(start).End(end).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectBillingList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectBillingList`: []Billing
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectBillingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectBillingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** | start | 
 **end** | **time.Time** | end | 
 **resourceType** | **string** | resource.type | 

### Return type

[**[]Billing**](Billing.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCostGet

> Metric IamProjectCostGet(ctx, projectId, costId).Execute()

Get iam/project.cost



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
    costId := "costId_example" // string | costId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCostGet(context.Background(), projectId, costId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCostGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCostGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**costId** | **string** | costId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Metric**](Metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCostList

> []Metric IamProjectCostList(ctx, projectId).Execute()

List iam/project.cost



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCostList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCostList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCostList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCostList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCostListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Metric**](Metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCostSeriesList

> []Point IamProjectCostSeriesList(ctx, projectId, costId).Timespan(timespan).Execute()

List iam/project.series



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
    costId := "costId_example" // string | costId
    timespan := "timespan_example" // string | timespan (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCostSeriesList(context.Background(), projectId, costId).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCostSeriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCostSeriesList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCostSeriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**costId** | **string** | costId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCostSeriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timespan** | **string** | timespan | 

### Return type

[**[]Point**](Point.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCreate

> Project IamProjectCreate(ctx).IamProjectCreate(iamProjectCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create iam/project



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
    iamProjectCreate := *openapiclient.NewIamProjectCreate("Name_example", "Organisation_example") // IamProjectCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCreate(context.Background()).IamProjectCreate(iamProjectCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCreate`: Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamProjectCreate** | [**IamProjectCreate**](IamProjectCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCredentialStoreCreate

> ProjectCredential IamProjectCredentialStoreCreate(ctx, projectId).ProjectCredential(projectCredential).Execute()

Create iam/project.credentialStore



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
    projectCredential := *openapiclient.NewProjectCredential("Name_example", "Type_example", "Value_example") // ProjectCredential | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCredentialStoreCreate(context.Background(), projectId).ProjectCredential(projectCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCredentialStoreCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCredentialStoreCreate`: ProjectCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCredentialStoreCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCredentialStoreCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectCredential** | [**ProjectCredential**](ProjectCredential.md) |  | 

### Return type

[**ProjectCredential**](ProjectCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCredentialStoreDelete

> Project IamProjectCredentialStoreDelete(ctx, projectId, credentialStoreId).Execute()

Delete iam/project.credentialStore



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
    credentialStoreId := "credentialStoreId_example" // string | credentialStoreId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCredentialStoreDelete(context.Background(), projectId, credentialStoreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCredentialStoreDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCredentialStoreDelete`: Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCredentialStoreDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**credentialStoreId** | **string** | credentialStoreId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCredentialStoreDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCredentialStoreGet

> ProjectCredential IamProjectCredentialStoreGet(ctx, projectId, credentialStoreId).Execute()

Get iam/project.credentialStore



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
    credentialStoreId := "credentialStoreId_example" // string | credentialStoreId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCredentialStoreGet(context.Background(), projectId, credentialStoreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCredentialStoreGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCredentialStoreGet`: ProjectCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCredentialStoreGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**credentialStoreId** | **string** | credentialStoreId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCredentialStoreGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectCredential**](ProjectCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCredentialStoreList

> []ProjectCredential IamProjectCredentialStoreList(ctx, projectId).Execute()

List iam/project.credentialStore



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCredentialStoreList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCredentialStoreList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCredentialStoreList`: []ProjectCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCredentialStoreList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCredentialStoreListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectCredential**](ProjectCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectCredentialStorePatch

> ProjectCredential IamProjectCredentialStorePatch(ctx, projectId, credentialStoreId).IamProjectCredentialStorePatch(iamProjectCredentialStorePatch).Execute()

Update iam/project.credentialStore



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
    credentialStoreId := "credentialStoreId_example" // string | credentialStoreId
    iamProjectCredentialStorePatch := *openapiclient.NewIamProjectCredentialStorePatch("Name_example") // IamProjectCredentialStorePatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectCredentialStorePatch(context.Background(), projectId, credentialStoreId).IamProjectCredentialStorePatch(iamProjectCredentialStorePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectCredentialStorePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectCredentialStorePatch`: ProjectCredential
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectCredentialStorePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**credentialStoreId** | **string** | credentialStoreId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectCredentialStorePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectCredentialStorePatch** | [**IamProjectCredentialStorePatch**](IamProjectCredentialStorePatch.md) |  | 

### Return type

[**ProjectCredential**](ProjectCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectDelete

> Project IamProjectDelete(ctx, projectId).Execute()

Delete iam/project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectDelete(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectDelete`: Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectEventGet

> Event IamProjectEventGet(ctx, projectId, eventId).Execute()

Get iam/project.event



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
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectEventGet(context.Background(), projectId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectEventGetRequest struct via the builder pattern


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


## IamProjectEventList

> []Event IamProjectEventList(ctx, projectId).Limit(limit).Skip(skip).Execute()

List iam/project.event



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
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectEventList(context.Background(), projectId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectEventListRequest struct via the builder pattern


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


## IamProjectGet

> Project IamProjectGet(ctx, projectId).Execute()

Get iam/project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectGet`: Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectInvitationAccept

> Invitation IamProjectInvitationAccept(ctx, projectId, invitationId).IamProjectInvitationAccept(iamProjectInvitationAccept).Execute()

Accept iam/project.invitation



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
    invitationId := "invitationId_example" // string | invitationId
    iamProjectInvitationAccept := *openapiclient.NewIamProjectInvitationAccept("Token_example") // IamProjectInvitationAccept | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectInvitationAccept(context.Background(), projectId, invitationId).IamProjectInvitationAccept(iamProjectInvitationAccept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectInvitationAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectInvitationAccept`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectInvitationAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**invitationId** | **string** | invitationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectInvitationAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectInvitationAccept** | [**IamProjectInvitationAccept**](IamProjectInvitationAccept.md) |  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectInvitationDelete

> IamProjectInvitationDelete(ctx, projectId, invitationId).Execute()

Delete iam/project.invitation



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
    invitationId := "invitationId_example" // string | invitationId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectInvitationDelete(context.Background(), projectId, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectInvitationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**invitationId** | **string** | invitationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectInvitationDeleteRequest struct via the builder pattern


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


## IamProjectInvitationGet

> Invitation IamProjectInvitationGet(ctx, projectId, invitationId).Execute()

Get iam/project.invitation



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
    invitationId := "invitationId_example" // string | invitationId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectInvitationGet(context.Background(), projectId, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectInvitationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectInvitationGet`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectInvitationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**invitationId** | **string** | invitationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectInvitationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Invitation**](Invitation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectInvitationList

> []Invitation IamProjectInvitationList(ctx, projectId).Resource(resource).Execute()

List iam/project.invitation



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
    resource := "resource_example" // string | resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectInvitationList(context.Background(), projectId).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectInvitationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectInvitationList`: []Invitation
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectInvitationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectInvitationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | **string** | resource | 

### Return type

[**[]Invitation**](Invitation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectInvoiceList

> []Invoice IamProjectInvoiceList(ctx, projectId).Execute()

List iam/project.invoice



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectInvoiceList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectInvoiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectInvoiceList`: []Invoice
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectInvoiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectInvoiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectList

> []Project IamProjectList(ctx).Name(name).Limit(limit).Active(active).Organisation(organisation).Lean(lean).TagValue(tagValue).TagKey(tagKey).Execute()

List iam/project



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
    name := "name_example" // string | Filter by name (optional)
    limit := float32(8.14) // float32 | Filter by $limit (optional)
    active := true // bool | Filter by active (optional) (default to false)
    organisation := "organisation_example" // string | Filter by organisation (optional)
    lean := true // bool | return a lightweight version of the resource (optional) (default to false)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectList(context.Background()).Name(name).Limit(limit).Active(active).Organisation(organisation).Lean(lean).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectList`: []Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **limit** | **float32** | Filter by $limit | 
 **active** | **bool** | Filter by active | [default to false]
 **organisation** | **string** | Filter by organisation | 
 **lean** | **bool** | return a lightweight version of the resource | [default to false]
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectOwnershipCreate

> Project IamProjectOwnershipCreate(ctx, projectId).IamProjectOwnershipCreate(iamProjectOwnershipCreate).Execute()

Create iam/project.ownership



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
    iamProjectOwnershipCreate := *openapiclient.NewIamProjectOwnershipCreate("Email_example") // IamProjectOwnershipCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectOwnershipCreate(context.Background(), projectId).IamProjectOwnershipCreate(iamProjectOwnershipCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectOwnershipCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectOwnershipCreate`: Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectOwnershipCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectOwnershipCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectOwnershipCreate** | [**IamProjectOwnershipCreate**](IamProjectOwnershipCreate.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectOwnershipDelete

> IamProjectOwnershipDelete(ctx, projectId, ownershipId).Execute()

Delete iam/project.ownership



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
    ownershipId := "ownershipId_example" // string | ownershipId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectOwnershipDelete(context.Background(), projectId, ownershipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectOwnershipDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**ownershipId** | **string** | ownershipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectOwnershipDeleteRequest struct via the builder pattern


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


## IamProjectOwnershipGet

> Ownership IamProjectOwnershipGet(ctx, projectId, ownershipId).Execute()

Get iam/project.ownership



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
    ownershipId := "ownershipId_example" // string | ownershipId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectOwnershipGet(context.Background(), projectId, ownershipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectOwnershipGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectOwnershipGet`: Ownership
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectOwnershipGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**ownershipId** | **string** | ownershipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectOwnershipGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ownership**](Ownership.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectOwnershipList

> []Ownership IamProjectOwnershipList(ctx, projectId).Execute()

List iam/project.ownership



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectOwnershipList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectOwnershipList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectOwnershipList`: []Ownership
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectOwnershipList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectOwnershipListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Ownership**](Ownership.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectPaymentList

> []Payment IamProjectPaymentList(ctx, projectId).Execute()

List iam/project.payment



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectPaymentList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectPaymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectPaymentList`: []Payment
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectPaymentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectPaymentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Payment**](Payment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectProformaList

> []Proforma IamProjectProformaList(ctx, projectId).Execute()

List iam/project.proforma



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectProformaList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectProformaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectProformaList`: []Proforma
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectProformaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectProformaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Proforma**](Proforma.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectQuotaGet

> Quota IamProjectQuotaGet(ctx, projectId, quotaId).Execute()

Get iam/project.quota



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
    quotaId := "quotaId_example" // string | quotaId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectQuotaGet(context.Background(), projectId, quotaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectQuotaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectQuotaGet`: Quota
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectQuotaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**quotaId** | **string** | quotaId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectQuotaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Quota**](Quota.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectQuotaLimitPatch

> QuotaLimit IamProjectQuotaLimitPatch(ctx, projectId, quotaId).IamProjectQuotaLimitPatch(iamProjectQuotaLimitPatch).Execute()

Update iam/project.limit



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
    quotaId := "quotaId_example" // string | quotaId
    iamProjectQuotaLimitPatch := *openapiclient.NewIamProjectQuotaLimitPatch() // IamProjectQuotaLimitPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectQuotaLimitPatch(context.Background(), projectId, quotaId).IamProjectQuotaLimitPatch(iamProjectQuotaLimitPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectQuotaLimitPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectQuotaLimitPatch`: QuotaLimit
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectQuotaLimitPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**quotaId** | **string** | quotaId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectQuotaLimitPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamProjectQuotaLimitPatch** | [**IamProjectQuotaLimitPatch**](IamProjectQuotaLimitPatch.md) |  | 

### Return type

[**QuotaLimit**](QuotaLimit.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectQuotaList

> []Quota IamProjectQuotaList(ctx, projectId).Execute()

List iam/project.quota



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectQuotaList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectQuotaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectQuotaList`: []Quota
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectQuotaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectQuotaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Quota**](Quota.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectServiceGet

> ResourceService IamProjectServiceGet(ctx, projectId, serviceId).Execute()

Get iam/project.service



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
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectServiceGet(context.Background(), projectId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectServiceGetRequest struct via the builder pattern


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


## IamProjectServiceList

> []ResourceService IamProjectServiceList(ctx, projectId).Execute()

List iam/project.service



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectServiceList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectServiceListRequest struct via the builder pattern


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


## IamProjectTagCreate

> Tag IamProjectTagCreate(ctx, projectId).Tag(tag).Execute()

Create iam/project.tag



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
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectTagCreate(context.Background(), projectId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectTagCreateRequest struct via the builder pattern


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


## IamProjectTagDelete

> IamProjectTagDelete(ctx, projectId, tagId).Execute()

Delete iam/project.tag



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
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectTagDelete(context.Background(), projectId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectTagDeleteRequest struct via the builder pattern


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


## IamProjectTagGet

> Tag IamProjectTagGet(ctx, projectId, tagId).Execute()

Get iam/project.tag



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
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectTagGet(context.Background(), projectId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectTagGetRequest struct via the builder pattern


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


## IamProjectTagList

> []Tag IamProjectTagList(ctx, projectId).Execute()

List iam/project.tag



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectTagList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectTagListRequest struct via the builder pattern


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


## IamProjectTagPut

> []Tag IamProjectTagPut(ctx, projectId).Tag(tag).Execute()

Replace iam/project.tag



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
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectTagPut(context.Background(), projectId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectTagPutRequest struct via the builder pattern


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


## IamProjectThresholdCreate

> ProjectThreshold IamProjectThresholdCreate(ctx, projectId).IamProjectThresholdCreate(iamProjectThresholdCreate).Execute()

Create iam/project.threshold



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
    iamProjectThresholdCreate := *openapiclient.NewIamProjectThresholdCreate() // IamProjectThresholdCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectThresholdCreate(context.Background(), projectId).IamProjectThresholdCreate(iamProjectThresholdCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectThresholdCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectThresholdCreate`: ProjectThreshold
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectThresholdCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectThresholdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectThresholdCreate** | [**IamProjectThresholdCreate**](IamProjectThresholdCreate.md) |  | 

### Return type

[**ProjectThreshold**](ProjectThreshold.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectThresholdDelete

> IamProjectThresholdDelete(ctx, projectId, thresholdId).Execute()

Delete iam/project.threshold



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
    thresholdId := "thresholdId_example" // string | thresholdId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectThresholdDelete(context.Background(), projectId, thresholdId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectThresholdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**thresholdId** | **string** | thresholdId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectThresholdDeleteRequest struct via the builder pattern


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


## IamProjectThresholdGet

> ProjectThreshold IamProjectThresholdGet(ctx, projectId, thresholdId).Execute()

Get iam/project.threshold



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
    thresholdId := "thresholdId_example" // string | thresholdId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectThresholdGet(context.Background(), projectId, thresholdId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectThresholdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectThresholdGet`: ProjectThreshold
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectThresholdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**thresholdId** | **string** | thresholdId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectThresholdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectThreshold**](ProjectThreshold.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectThresholdList

> []ProjectThreshold IamProjectThresholdList(ctx, projectId).Execute()

List iam/project.threshold



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectThresholdList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectThresholdList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectThresholdList`: []ProjectThreshold
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectThresholdList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectThresholdListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectThreshold**](ProjectThreshold.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectTransfer

> Project IamProjectTransfer(ctx, projectId).IamProjectTransfer(iamProjectTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer iam/project



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
    iamProjectTransfer := *openapiclient.NewIamProjectTransfer("Organisation_example") // IamProjectTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectTransfer(context.Background(), projectId).IamProjectTransfer(iamProjectTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectTransfer`: Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectTransfer** | [**IamProjectTransfer**](IamProjectTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectUpdate

> Project IamProjectUpdate(ctx, projectId).IamProjectUpdate(iamProjectUpdate).Execute()

Update iam/project



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
    iamProjectUpdate := *openapiclient.NewIamProjectUpdate() // IamProjectUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectUpdate(context.Background(), projectId).IamProjectUpdate(iamProjectUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectUpdate`: Project
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamProjectUpdate** | [**IamProjectUpdate**](IamProjectUpdate.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectUsageGet

> Metric IamProjectUsageGet(ctx, projectId, usageId).Execute()

Get iam/project.usage



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
    usageId := "usageId_example" // string | usageId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectUsageGet(context.Background(), projectId, usageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectUsageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectUsageGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectUsageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**usageId** | **string** | usageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectUsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Metric**](Metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectUsageList

> []Metric IamProjectUsageList(ctx, projectId).Execute()

List iam/project.usage



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectUsageList(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectUsageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectUsageList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectUsageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectUsageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Metric**](Metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectUsageSeriesList

> []Point IamProjectUsageSeriesList(ctx, projectId, usageId).Timespan(timespan).Color(color).Execute()

List iam/project.series



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
    usageId := "usageId_example" // string | usageId
    timespan := "timespan_example" // string | timespan (optional)
    color := "color_example" // string | color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProjectApi.IamProjectUsageSeriesList(context.Background(), projectId, usageId).Timespan(timespan).Color(color).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProjectApi.IamProjectUsageSeriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamProjectUsageSeriesList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `IamProjectApi.IamProjectUsageSeriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**usageId** | **string** | usageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamProjectUsageSeriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timespan** | **string** | timespan | 
 **color** | **string** | color | 

### Return type

[**[]Point**](Point.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

