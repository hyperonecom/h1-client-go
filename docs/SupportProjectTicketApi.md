# \SupportProjectTicketApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupportProjectTicketClose**](SupportProjectTicketApi.md#SupportProjectTicketClose) | **Post** /support/project/{projectId}/ticket/{ticketId}/actions/close | Close support/ticket
[**SupportProjectTicketCreate**](SupportProjectTicketApi.md#SupportProjectTicketCreate) | **Post** /support/project/{projectId}/ticket | Create support/ticket
[**SupportProjectTicketGet**](SupportProjectTicketApi.md#SupportProjectTicketGet) | **Get** /support/project/{projectId}/ticket/{ticketId} | Get support/ticket
[**SupportProjectTicketList**](SupportProjectTicketApi.md#SupportProjectTicketList) | **Get** /support/project/{projectId}/ticket | List support/ticket
[**SupportProjectTicketMessageCreate**](SupportProjectTicketApi.md#SupportProjectTicketMessageCreate) | **Post** /support/project/{projectId}/ticket/{ticketId}/message | Create support/ticket.message
[**SupportProjectTicketMessageGet**](SupportProjectTicketApi.md#SupportProjectTicketMessageGet) | **Get** /support/project/{projectId}/ticket/{ticketId}/message/{messageId} | Get support/ticket.message
[**SupportProjectTicketMessageList**](SupportProjectTicketApi.md#SupportProjectTicketMessageList) | **Get** /support/project/{projectId}/ticket/{ticketId}/message | List support/ticket.message



## SupportProjectTicketClose

> Ticket SupportProjectTicketClose(ctx, projectId, ticketId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Close support/ticket



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
    ticketId := "ticketId_example" // string | Ticket Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportProjectTicketApi.SupportProjectTicketClose(context.Background(), projectId, ticketId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportProjectTicketApi.SupportProjectTicketClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportProjectTicketClose`: Ticket
    fmt.Fprintf(os.Stdout, "Response from `SupportProjectTicketApi.SupportProjectTicketClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**ticketId** | **string** | Ticket Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportProjectTicketCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketCreate

> Ticket SupportProjectTicketCreate(ctx, projectId).SupportProjectTicketCreate(supportProjectTicketCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create support/ticket



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
    supportProjectTicketCreate := *openapiclient.NewSupportProjectTicketCreate("Type_example", "Subject_example", "Message_example") // SupportProjectTicketCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportProjectTicketApi.SupportProjectTicketCreate(context.Background(), projectId).SupportProjectTicketCreate(supportProjectTicketCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportProjectTicketApi.SupportProjectTicketCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportProjectTicketCreate`: Ticket
    fmt.Fprintf(os.Stdout, "Response from `SupportProjectTicketApi.SupportProjectTicketCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportProjectTicketCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportProjectTicketCreate** | [**SupportProjectTicketCreate**](SupportProjectTicketCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketGet

> Ticket SupportProjectTicketGet(ctx, projectId, ticketId).Execute()

Get support/ticket



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
    ticketId := "ticketId_example" // string | Ticket Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportProjectTicketApi.SupportProjectTicketGet(context.Background(), projectId, ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportProjectTicketApi.SupportProjectTicketGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportProjectTicketGet`: Ticket
    fmt.Fprintf(os.Stdout, "Response from `SupportProjectTicketApi.SupportProjectTicketGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**ticketId** | **string** | Ticket Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportProjectTicketGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ticket**](Ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketList

> []Ticket SupportProjectTicketList(ctx, projectId).State(state).Execute()

List support/ticket



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
    state := "state_example" // string | Filter by state (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportProjectTicketApi.SupportProjectTicketList(context.Background(), projectId).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportProjectTicketApi.SupportProjectTicketList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportProjectTicketList`: []Ticket
    fmt.Fprintf(os.Stdout, "Response from `SupportProjectTicketApi.SupportProjectTicketList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportProjectTicketListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Filter by state | 

### Return type

[**[]Ticket**](Ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketMessageCreate

> SupportMessage SupportProjectTicketMessageCreate(ctx, projectId, ticketId).SupportMessage(supportMessage).Execute()

Create support/ticket.message



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
    ticketId := "ticketId_example" // string | Ticket Id
    supportMessage := *openapiclient.NewSupportMessage("Type_example") // SupportMessage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportProjectTicketApi.SupportProjectTicketMessageCreate(context.Background(), projectId, ticketId).SupportMessage(supportMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportProjectTicketApi.SupportProjectTicketMessageCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportProjectTicketMessageCreate`: SupportMessage
    fmt.Fprintf(os.Stdout, "Response from `SupportProjectTicketApi.SupportProjectTicketMessageCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**ticketId** | **string** | Ticket Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportProjectTicketMessageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportMessage** | [**SupportMessage**](SupportMessage.md) |  | 

### Return type

[**SupportMessage**](SupportMessage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketMessageGet

> SupportMessage SupportProjectTicketMessageGet(ctx, projectId, ticketId, messageId).Execute()

Get support/ticket.message



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
    ticketId := "ticketId_example" // string | Ticket Id
    messageId := "messageId_example" // string | messageId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportProjectTicketApi.SupportProjectTicketMessageGet(context.Background(), projectId, ticketId, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportProjectTicketApi.SupportProjectTicketMessageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportProjectTicketMessageGet`: SupportMessage
    fmt.Fprintf(os.Stdout, "Response from `SupportProjectTicketApi.SupportProjectTicketMessageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**ticketId** | **string** | Ticket Id | 
**messageId** | **string** | messageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportProjectTicketMessageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SupportMessage**](SupportMessage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketMessageList

> []SupportMessage SupportProjectTicketMessageList(ctx, projectId, ticketId).Execute()

List support/ticket.message



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
    ticketId := "ticketId_example" // string | Ticket Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportProjectTicketApi.SupportProjectTicketMessageList(context.Background(), projectId, ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportProjectTicketApi.SupportProjectTicketMessageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportProjectTicketMessageList`: []SupportMessage
    fmt.Fprintf(os.Stdout, "Response from `SupportProjectTicketApi.SupportProjectTicketMessageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**ticketId** | **string** | Ticket Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupportProjectTicketMessageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SupportMessage**](SupportMessage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

