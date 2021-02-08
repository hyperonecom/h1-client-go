# \InsightProjectJournalApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsightProjectJournalCreate**](InsightProjectJournalApi.md#InsightProjectJournalCreate) | **Post** /insight/{locationId}/project/{projectId}/journal | Create insight/journal
[**InsightProjectJournalCredentialCreate**](InsightProjectJournalApi.md#InsightProjectJournalCredentialCreate) | **Post** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential | Create insight/journal.credential
[**InsightProjectJournalCredentialDelete**](InsightProjectJournalApi.md#InsightProjectJournalCredentialDelete) | **Delete** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential/{credentialId} | Delete insight/journal.credential
[**InsightProjectJournalCredentialGet**](InsightProjectJournalApi.md#InsightProjectJournalCredentialGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential/{credentialId} | Get insight/journal.credential
[**InsightProjectJournalCredentialList**](InsightProjectJournalApi.md#InsightProjectJournalCredentialList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential | List insight/journal.credential
[**InsightProjectJournalCredentialPatch**](InsightProjectJournalApi.md#InsightProjectJournalCredentialPatch) | **Patch** /insight/{locationId}/project/{projectId}/journal/{journalId}/credential/{credentialId} | Update insight/journal.credential
[**InsightProjectJournalDelete**](InsightProjectJournalApi.md#InsightProjectJournalDelete) | **Delete** /insight/{locationId}/project/{projectId}/journal/{journalId} | Delete insight/journal
[**InsightProjectJournalEventGet**](InsightProjectJournalApi.md#InsightProjectJournalEventGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/event/{eventId} | Get insight/journal.event
[**InsightProjectJournalEventList**](InsightProjectJournalApi.md#InsightProjectJournalEventList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/event | List insight/journal.event
[**InsightProjectJournalGet**](InsightProjectJournalApi.md#InsightProjectJournalGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId} | Get insight/journal
[**InsightProjectJournalList**](InsightProjectJournalApi.md#InsightProjectJournalList) | **Get** /insight/{locationId}/project/{projectId}/journal | List insight/journal
[**InsightProjectJournalLogGet**](InsightProjectJournalApi.md#InsightProjectJournalLogGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/log | Get insight/journal.log
[**InsightProjectJournalServiceGet**](InsightProjectJournalApi.md#InsightProjectJournalServiceGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/service/{serviceId} | Get insight/journal.service
[**InsightProjectJournalServiceList**](InsightProjectJournalApi.md#InsightProjectJournalServiceList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/service | List insight/journal.service
[**InsightProjectJournalTagCreate**](InsightProjectJournalApi.md#InsightProjectJournalTagCreate) | **Post** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag | Create insight/journal.tag
[**InsightProjectJournalTagDelete**](InsightProjectJournalApi.md#InsightProjectJournalTagDelete) | **Delete** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag/{tagId} | Delete insight/journal.tag
[**InsightProjectJournalTagGet**](InsightProjectJournalApi.md#InsightProjectJournalTagGet) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag/{tagId} | Get insight/journal.tag
[**InsightProjectJournalTagList**](InsightProjectJournalApi.md#InsightProjectJournalTagList) | **Get** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag | List insight/journal.tag
[**InsightProjectJournalTagPut**](InsightProjectJournalApi.md#InsightProjectJournalTagPut) | **Put** /insight/{locationId}/project/{projectId}/journal/{journalId}/tag | Replace insight/journal.tag
[**InsightProjectJournalTransfer**](InsightProjectJournalApi.md#InsightProjectJournalTransfer) | **Post** /insight/{locationId}/project/{projectId}/journal/{journalId}/actions/transfer | Transfer insight/journal
[**InsightProjectJournalUpdate**](InsightProjectJournalApi.md#InsightProjectJournalUpdate) | **Patch** /insight/{locationId}/project/{projectId}/journal/{journalId} | Update insight/journal



## InsightProjectJournalCreate

> Journal InsightProjectJournalCreate(ctx, projectId, locationId).InsightProjectJournalCreate(insightProjectJournalCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create insight/journal



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
    insightProjectJournalCreate := *openapiclient.NewInsightProjectJournalCreate("Name_example") // InsightProjectJournalCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalCreate(context.Background(), projectId, locationId).InsightProjectJournalCreate(insightProjectJournalCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalCreate`: Journal
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **insightProjectJournalCreate** | [**InsightProjectJournalCreate**](InsightProjectJournalCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialCreate

> JournalCredential InsightProjectJournalCredentialCreate(ctx, projectId, locationId, journalId).JournalCredential(journalCredential).Execute()

Create insight/journal.credential



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
    journalId := "journalId_example" // string | Journal Id
    journalCredential := *openapiclient.NewJournalCredential("Name_example", "Type_example", "Value_example") // JournalCredential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalCredentialCreate(context.Background(), projectId, locationId, journalId).JournalCredential(journalCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalCredentialCreate`: JournalCredential
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **journalCredential** | [**JournalCredential**](JournalCredential.md) |  | 

### Return type

[**JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialDelete

> Journal InsightProjectJournalCredentialDelete(ctx, projectId, locationId, journalId, credentialId).Execute()

Delete insight/journal.credential



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
    journalId := "journalId_example" // string | Journal Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalCredentialDelete(context.Background(), projectId, locationId, journalId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalCredentialDelete`: Journal
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialGet

> JournalCredential InsightProjectJournalCredentialGet(ctx, projectId, locationId, journalId, credentialId).Execute()

Get insight/journal.credential



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
    journalId := "journalId_example" // string | Journal Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalCredentialGet(context.Background(), projectId, locationId, journalId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalCredentialGet`: JournalCredential
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialList

> []JournalCredential InsightProjectJournalCredentialList(ctx, projectId, locationId, journalId).Execute()

List insight/journal.credential



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
    journalId := "journalId_example" // string | Journal Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalCredentialList(context.Background(), projectId, locationId, journalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalCredentialList`: []JournalCredential
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalCredentialPatch

> JournalCredential InsightProjectJournalCredentialPatch(ctx, projectId, locationId, journalId, credentialId).InsightProjectJournalCredentialPatch(insightProjectJournalCredentialPatch).Execute()

Update insight/journal.credential



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
    journalId := "journalId_example" // string | Journal Id
    credentialId := "credentialId_example" // string | credentialId
    insightProjectJournalCredentialPatch := *openapiclient.NewInsightProjectJournalCredentialPatch("Name_example") // InsightProjectJournalCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalCredentialPatch(context.Background(), projectId, locationId, journalId, credentialId).InsightProjectJournalCredentialPatch(insightProjectJournalCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalCredentialPatch`: JournalCredential
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **insightProjectJournalCredentialPatch** | [**InsightProjectJournalCredentialPatch**](InsightProjectJournalCredentialPatch.md) |  | 

### Return type

[**JournalCredential**](journal.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalDelete

> InsightProjectJournalDelete(ctx, projectId, locationId, journalId).Execute()

Delete insight/journal



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
    journalId := "journalId_example" // string | Journal Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalDelete(context.Background(), projectId, locationId, journalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalDelete``: %v\n", err)
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
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalDeleteRequest struct via the builder pattern


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


## InsightProjectJournalEventGet

> Event InsightProjectJournalEventGet(ctx, projectId, locationId, journalId, eventId).Execute()

Get insight/journal.event



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
    journalId := "journalId_example" // string | Journal Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalEventGet(context.Background(), projectId, locationId, journalId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalEventGetRequest struct via the builder pattern


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


## InsightProjectJournalEventList

> []Event InsightProjectJournalEventList(ctx, projectId, locationId, journalId).Limit(limit).Skip(skip).Execute()

List insight/journal.event



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
    journalId := "journalId_example" // string | Journal Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalEventList(context.Background(), projectId, locationId, journalId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalEventListRequest struct via the builder pattern


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


## InsightProjectJournalGet

> Journal InsightProjectJournalGet(ctx, projectId, locationId, journalId).Execute()

Get insight/journal



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
    journalId := "journalId_example" // string | Journal Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalGet(context.Background(), projectId, locationId, journalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalGet`: Journal
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalList

> []Journal InsightProjectJournalList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List insight/journal



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
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalList`: []Journal
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalLogGet

> InsightProjectJournalLogGet(ctx, projectId, locationId, journalId).Since(since).Until(until).Follow(follow).Tail(tail).Tag(tag).Execute()

Get insight/journal.log



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
    locationId := "locationId_example" // string | Location Id
    journalId := "journalId_example" // string | Journal Id
    since := time.Now() // time.Time | since (optional)
    until := time.Now() // time.Time | until (optional)
    follow := true // bool | follow (optional) (default to false)
    tail := float32(8.14) // float32 | tail (optional)
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | tag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalLogGet(context.Background(), projectId, locationId, journalId).Since(since).Until(until).Follow(follow).Tail(tail).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalLogGet``: %v\n", err)
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
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalLogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **since** | **time.Time** | since | 
 **until** | **time.Time** | until | 
 **follow** | **bool** | follow | [default to false]
 **tail** | **float32** | tail | 
 **tag** | [**[]Tag**](tag.md) | tag | 

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


## InsightProjectJournalServiceGet

> ResourceService InsightProjectJournalServiceGet(ctx, projectId, locationId, journalId, serviceId).Execute()

Get insight/journal.service



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
    journalId := "journalId_example" // string | Journal Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalServiceGet(context.Background(), projectId, locationId, journalId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalServiceGetRequest struct via the builder pattern


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


## InsightProjectJournalServiceList

> []ResourceService InsightProjectJournalServiceList(ctx, projectId, locationId, journalId).Execute()

List insight/journal.service



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
    journalId := "journalId_example" // string | Journal Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalServiceList(context.Background(), projectId, locationId, journalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalServiceListRequest struct via the builder pattern


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


## InsightProjectJournalTagCreate

> Tag InsightProjectJournalTagCreate(ctx, projectId, locationId, journalId).Tag(tag).Execute()

Create insight/journal.tag



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
    journalId := "journalId_example" // string | Journal Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalTagCreate(context.Background(), projectId, locationId, journalId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalTagCreateRequest struct via the builder pattern


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


## InsightProjectJournalTagDelete

> InsightProjectJournalTagDelete(ctx, projectId, locationId, journalId, tagId).Execute()

Delete insight/journal.tag



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
    journalId := "journalId_example" // string | Journal Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalTagDelete(context.Background(), projectId, locationId, journalId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalTagDelete``: %v\n", err)
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
**journalId** | **string** | Journal Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalTagDeleteRequest struct via the builder pattern


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


## InsightProjectJournalTagGet

> Tag InsightProjectJournalTagGet(ctx, projectId, locationId, journalId, tagId).Execute()

Get insight/journal.tag



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
    journalId := "journalId_example" // string | Journal Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalTagGet(context.Background(), projectId, locationId, journalId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalTagGetRequest struct via the builder pattern


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


## InsightProjectJournalTagList

> []Tag InsightProjectJournalTagList(ctx, projectId, locationId, journalId).Execute()

List insight/journal.tag



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
    journalId := "journalId_example" // string | Journal Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalTagList(context.Background(), projectId, locationId, journalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalTagListRequest struct via the builder pattern


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


## InsightProjectJournalTagPut

> []Tag InsightProjectJournalTagPut(ctx, projectId, locationId, journalId).Tag(tag).Execute()

Replace insight/journal.tag



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
    journalId := "journalId_example" // string | Journal Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalTagPut(context.Background(), projectId, locationId, journalId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalTagPutRequest struct via the builder pattern


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


## InsightProjectJournalTransfer

> Journal InsightProjectJournalTransfer(ctx, projectId, locationId, journalId).InsightProjectJournalTransfer(insightProjectJournalTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer insight/journal



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
    journalId := "journalId_example" // string | Journal Id
    insightProjectJournalTransfer := *openapiclient.NewInsightProjectJournalTransfer("Project_example") // InsightProjectJournalTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalTransfer(context.Background(), projectId, locationId, journalId).InsightProjectJournalTransfer(insightProjectJournalTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalTransfer`: Journal
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **insightProjectJournalTransfer** | [**InsightProjectJournalTransfer**](InsightProjectJournalTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightProjectJournalUpdate

> Journal InsightProjectJournalUpdate(ctx, projectId, locationId, journalId).InsightProjectJournalUpdate(insightProjectJournalUpdate).Execute()

Update insight/journal



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
    journalId := "journalId_example" // string | Journal Id
    insightProjectJournalUpdate := *openapiclient.NewInsightProjectJournalUpdate() // InsightProjectJournalUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightProjectJournalApi.InsightProjectJournalUpdate(context.Background(), projectId, locationId, journalId).InsightProjectJournalUpdate(insightProjectJournalUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightProjectJournalApi.InsightProjectJournalUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightProjectJournalUpdate`: Journal
    fmt.Fprintf(os.Stdout, "Response from `InsightProjectJournalApi.InsightProjectJournalUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**journalId** | **string** | Journal Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightProjectJournalUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **insightProjectJournalUpdate** | [**InsightProjectJournalUpdate**](InsightProjectJournalUpdate.md) |  | 

### Return type

[**Journal**](journal.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

