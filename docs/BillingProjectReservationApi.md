# \BillingProjectReservationApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingProjectReservationAssign**](BillingProjectReservationApi.md#BillingProjectReservationAssign) | **Post** /billing/project/{projectId}/reservation/{reservationId}/actions/assign | Assign billing/reservation
[**BillingProjectReservationCreate**](BillingProjectReservationApi.md#BillingProjectReservationCreate) | **Post** /billing/project/{projectId}/reservation | Create billing/reservation
[**BillingProjectReservationDelete**](BillingProjectReservationApi.md#BillingProjectReservationDelete) | **Delete** /billing/project/{projectId}/reservation/{reservationId} | Delete billing/reservation
[**BillingProjectReservationEventGet**](BillingProjectReservationApi.md#BillingProjectReservationEventGet) | **Get** /billing/project/{projectId}/reservation/{reservationId}/event/{eventId} | Get billing/reservation.event
[**BillingProjectReservationEventList**](BillingProjectReservationApi.md#BillingProjectReservationEventList) | **Get** /billing/project/{projectId}/reservation/{reservationId}/event | List billing/reservation.event
[**BillingProjectReservationExtend**](BillingProjectReservationApi.md#BillingProjectReservationExtend) | **Post** /billing/project/{projectId}/reservation/{reservationId}/actions/extend | Extend billing/reservation
[**BillingProjectReservationGet**](BillingProjectReservationApi.md#BillingProjectReservationGet) | **Get** /billing/project/{projectId}/reservation/{reservationId} | Get billing/reservation
[**BillingProjectReservationList**](BillingProjectReservationApi.md#BillingProjectReservationList) | **Get** /billing/project/{projectId}/reservation | List billing/reservation
[**BillingProjectReservationServiceGet**](BillingProjectReservationApi.md#BillingProjectReservationServiceGet) | **Get** /billing/project/{projectId}/reservation/{reservationId}/service/{serviceId} | Get billing/reservation.service
[**BillingProjectReservationServiceList**](BillingProjectReservationApi.md#BillingProjectReservationServiceList) | **Get** /billing/project/{projectId}/reservation/{reservationId}/service | List billing/reservation.service
[**BillingProjectReservationTagCreate**](BillingProjectReservationApi.md#BillingProjectReservationTagCreate) | **Post** /billing/project/{projectId}/reservation/{reservationId}/tag | Create billing/reservation.tag
[**BillingProjectReservationTagDelete**](BillingProjectReservationApi.md#BillingProjectReservationTagDelete) | **Delete** /billing/project/{projectId}/reservation/{reservationId}/tag/{tagId} | Delete billing/reservation.tag
[**BillingProjectReservationTagGet**](BillingProjectReservationApi.md#BillingProjectReservationTagGet) | **Get** /billing/project/{projectId}/reservation/{reservationId}/tag/{tagId} | Get billing/reservation.tag
[**BillingProjectReservationTagList**](BillingProjectReservationApi.md#BillingProjectReservationTagList) | **Get** /billing/project/{projectId}/reservation/{reservationId}/tag | List billing/reservation.tag
[**BillingProjectReservationTagPut**](BillingProjectReservationApi.md#BillingProjectReservationTagPut) | **Put** /billing/project/{projectId}/reservation/{reservationId}/tag | Replace billing/reservation.tag
[**BillingProjectReservationUpdate**](BillingProjectReservationApi.md#BillingProjectReservationUpdate) | **Patch** /billing/project/{projectId}/reservation/{reservationId} | Update billing/reservation



## BillingProjectReservationAssign

> Reservation BillingProjectReservationAssign(ctx, projectId, reservationId).BillingProjectReservationAssign(billingProjectReservationAssign).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Assign billing/reservation



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
    reservationId := "reservationId_example" // string | Reservation Id
    billingProjectReservationAssign := *openapiclient.NewBillingProjectReservationAssign() // BillingProjectReservationAssign | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationAssign(context.Background(), projectId, reservationId).BillingProjectReservationAssign(billingProjectReservationAssign).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationAssign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationAssign`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationAssign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **billingProjectReservationAssign** | [**BillingProjectReservationAssign**](BillingProjectReservationAssign.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationCreate

> Reservation BillingProjectReservationCreate(ctx, projectId).BillingProjectReservationCreate(billingProjectReservationCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create billing/reservation



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
    billingProjectReservationCreate := *openapiclient.NewBillingProjectReservationCreate("Name_example", "Service_example") // BillingProjectReservationCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationCreate(context.Background(), projectId).BillingProjectReservationCreate(billingProjectReservationCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationCreate`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingProjectReservationCreate** | [**BillingProjectReservationCreate**](BillingProjectReservationCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationDelete

> BillingProjectReservationDelete(ctx, projectId, reservationId).Execute()

Delete billing/reservation



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
    reservationId := "reservationId_example" // string | Reservation Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationDelete(context.Background(), projectId, reservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationDeleteRequest struct via the builder pattern


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


## BillingProjectReservationEventGet

> Event BillingProjectReservationEventGet(ctx, projectId, reservationId, eventId).Execute()

Get billing/reservation.event



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
    reservationId := "reservationId_example" // string | Reservation Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationEventGet(context.Background(), projectId, reservationId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationEventGetRequest struct via the builder pattern


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


## BillingProjectReservationEventList

> []Event BillingProjectReservationEventList(ctx, projectId, reservationId).Limit(limit).Skip(skip).Execute()

List billing/reservation.event



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
    reservationId := "reservationId_example" // string | Reservation Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationEventList(context.Background(), projectId, reservationId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationEventListRequest struct via the builder pattern


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


## BillingProjectReservationExtend

> Reservation BillingProjectReservationExtend(ctx, projectId, reservationId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Extend billing/reservation



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
    reservationId := "reservationId_example" // string | Reservation Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationExtend(context.Background(), projectId, reservationId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationExtend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationExtend`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationExtend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationExtendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationGet

> Reservation BillingProjectReservationGet(ctx, projectId, reservationId).Execute()

Get billing/reservation



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
    reservationId := "reservationId_example" // string | Reservation Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationGet(context.Background(), projectId, reservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationGet`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationList

> []Reservation BillingProjectReservationList(ctx, projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List billing/reservation



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationList(context.Background(), projectId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationList`: []Reservation
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingProjectReservationServiceGet

> ResourceService BillingProjectReservationServiceGet(ctx, projectId, reservationId, serviceId).Execute()

Get billing/reservation.service



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
    reservationId := "reservationId_example" // string | Reservation Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationServiceGet(context.Background(), projectId, reservationId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationServiceGetRequest struct via the builder pattern


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


## BillingProjectReservationServiceList

> []ResourceService BillingProjectReservationServiceList(ctx, projectId, reservationId).Execute()

List billing/reservation.service



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
    reservationId := "reservationId_example" // string | Reservation Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationServiceList(context.Background(), projectId, reservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationServiceListRequest struct via the builder pattern


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


## BillingProjectReservationTagCreate

> Tag BillingProjectReservationTagCreate(ctx, projectId, reservationId).Tag(tag).Execute()

Create billing/reservation.tag



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
    reservationId := "reservationId_example" // string | Reservation Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationTagCreate(context.Background(), projectId, reservationId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationTagCreateRequest struct via the builder pattern


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


## BillingProjectReservationTagDelete

> BillingProjectReservationTagDelete(ctx, projectId, reservationId, tagId).Execute()

Delete billing/reservation.tag



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
    reservationId := "reservationId_example" // string | Reservation Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationTagDelete(context.Background(), projectId, reservationId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationTagDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationTagDeleteRequest struct via the builder pattern


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


## BillingProjectReservationTagGet

> Tag BillingProjectReservationTagGet(ctx, projectId, reservationId, tagId).Execute()

Get billing/reservation.tag



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
    reservationId := "reservationId_example" // string | Reservation Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationTagGet(context.Background(), projectId, reservationId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationTagGetRequest struct via the builder pattern


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


## BillingProjectReservationTagList

> []Tag BillingProjectReservationTagList(ctx, projectId, reservationId).Execute()

List billing/reservation.tag



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
    reservationId := "reservationId_example" // string | Reservation Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationTagList(context.Background(), projectId, reservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationTagListRequest struct via the builder pattern


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


## BillingProjectReservationTagPut

> []Tag BillingProjectReservationTagPut(ctx, projectId, reservationId).Tag(tag).Execute()

Replace billing/reservation.tag



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
    reservationId := "reservationId_example" // string | Reservation Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationTagPut(context.Background(), projectId, reservationId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationTagPutRequest struct via the builder pattern


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


## BillingProjectReservationUpdate

> Reservation BillingProjectReservationUpdate(ctx, projectId, reservationId).BillingProjectReservationUpdate(billingProjectReservationUpdate).Execute()

Update billing/reservation



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
    reservationId := "reservationId_example" // string | Reservation Id
    billingProjectReservationUpdate := *openapiclient.NewBillingProjectReservationUpdate() // BillingProjectReservationUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingProjectReservationApi.BillingProjectReservationUpdate(context.Background(), projectId, reservationId).BillingProjectReservationUpdate(billingProjectReservationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingProjectReservationApi.BillingProjectReservationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingProjectReservationUpdate`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `BillingProjectReservationApi.BillingProjectReservationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**reservationId** | **string** | Reservation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingProjectReservationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **billingProjectReservationUpdate** | [**BillingProjectReservationUpdate**](BillingProjectReservationUpdate.md) |  | 

### Return type

[**Reservation**](reservation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

