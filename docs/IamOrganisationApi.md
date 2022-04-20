# \IamOrganisationApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamOrganisationBillingList**](IamOrganisationApi.md#IamOrganisationBillingList) | **Get** /iam/organisation/{organisationId}/billing | List iam/organisation.billing
[**IamOrganisationCreate**](IamOrganisationApi.md#IamOrganisationCreate) | **Post** /iam/organisation | Create iam/organisation
[**IamOrganisationDelete**](IamOrganisationApi.md#IamOrganisationDelete) | **Delete** /iam/organisation/{organisationId} | Delete iam/organisation
[**IamOrganisationEventGet**](IamOrganisationApi.md#IamOrganisationEventGet) | **Get** /iam/organisation/{organisationId}/event/{eventId} | Get iam/organisation.event
[**IamOrganisationEventList**](IamOrganisationApi.md#IamOrganisationEventList) | **Get** /iam/organisation/{organisationId}/event | List iam/organisation.event
[**IamOrganisationGet**](IamOrganisationApi.md#IamOrganisationGet) | **Get** /iam/organisation/{organisationId} | Get iam/organisation
[**IamOrganisationInvitationAccept**](IamOrganisationApi.md#IamOrganisationInvitationAccept) | **Post** /iam/organisation/{organisationId}/invitation/{invitationId}/actions/accept | Accept iam/organisation.invitation
[**IamOrganisationInvitationDelete**](IamOrganisationApi.md#IamOrganisationInvitationDelete) | **Delete** /iam/organisation/{organisationId}/invitation/{invitationId} | Delete iam/organisation.invitation
[**IamOrganisationInvitationGet**](IamOrganisationApi.md#IamOrganisationInvitationGet) | **Get** /iam/organisation/{organisationId}/invitation/{invitationId} | Get iam/organisation.invitation
[**IamOrganisationInvitationList**](IamOrganisationApi.md#IamOrganisationInvitationList) | **Get** /iam/organisation/{organisationId}/invitation | List iam/organisation.invitation
[**IamOrganisationInvoiceDownload**](IamOrganisationApi.md#IamOrganisationInvoiceDownload) | **Post** /iam/organisation/{organisationId}/invoice/{invoiceId}/actions/download | Download iam/organisation.invoice
[**IamOrganisationInvoiceGet**](IamOrganisationApi.md#IamOrganisationInvoiceGet) | **Get** /iam/organisation/{organisationId}/invoice/{invoiceId} | Get iam/organisation.invoice
[**IamOrganisationInvoiceList**](IamOrganisationApi.md#IamOrganisationInvoiceList) | **Get** /iam/organisation/{organisationId}/invoice | List iam/organisation.invoice
[**IamOrganisationList**](IamOrganisationApi.md#IamOrganisationList) | **Get** /iam/organisation | List iam/organisation
[**IamOrganisationOwnershipCreate**](IamOrganisationApi.md#IamOrganisationOwnershipCreate) | **Post** /iam/organisation/{organisationId}/ownership | Create iam/organisation.ownership
[**IamOrganisationOwnershipDelete**](IamOrganisationApi.md#IamOrganisationOwnershipDelete) | **Delete** /iam/organisation/{organisationId}/ownership/{ownershipId} | Delete iam/organisation.ownership
[**IamOrganisationOwnershipGet**](IamOrganisationApi.md#IamOrganisationOwnershipGet) | **Get** /iam/organisation/{organisationId}/ownership/{ownershipId} | Get iam/organisation.ownership
[**IamOrganisationOwnershipList**](IamOrganisationApi.md#IamOrganisationOwnershipList) | **Get** /iam/organisation/{organisationId}/ownership | List iam/organisation.ownership
[**IamOrganisationPaymentAllocate**](IamOrganisationApi.md#IamOrganisationPaymentAllocate) | **Post** /iam/organisation/{organisationId}/payment/{paymentId}/actions/allocate | Allocate iam/organisation.payment
[**IamOrganisationPaymentGet**](IamOrganisationApi.md#IamOrganisationPaymentGet) | **Get** /iam/organisation/{organisationId}/payment/{paymentId} | Get iam/organisation.payment
[**IamOrganisationPaymentList**](IamOrganisationApi.md#IamOrganisationPaymentList) | **Get** /iam/organisation/{organisationId}/payment | List iam/organisation.payment
[**IamOrganisationProformaCreate**](IamOrganisationApi.md#IamOrganisationProformaCreate) | **Post** /iam/organisation/{organisationId}/proforma | Create iam/organisation.proforma
[**IamOrganisationProformaDownload**](IamOrganisationApi.md#IamOrganisationProformaDownload) | **Post** /iam/organisation/{organisationId}/proforma/{proformaId}/actions/download | Download iam/organisation.proforma
[**IamOrganisationProformaGet**](IamOrganisationApi.md#IamOrganisationProformaGet) | **Get** /iam/organisation/{organisationId}/proforma/{proformaId} | Get iam/organisation.proforma
[**IamOrganisationProformaList**](IamOrganisationApi.md#IamOrganisationProformaList) | **Get** /iam/organisation/{organisationId}/proforma | List iam/organisation.proforma
[**IamOrganisationServiceGet**](IamOrganisationApi.md#IamOrganisationServiceGet) | **Get** /iam/organisation/{organisationId}/service/{serviceId} | Get iam/organisation.service
[**IamOrganisationServiceList**](IamOrganisationApi.md#IamOrganisationServiceList) | **Get** /iam/organisation/{organisationId}/service | List iam/organisation.service
[**IamOrganisationTransferAccept**](IamOrganisationApi.md#IamOrganisationTransferAccept) | **Post** /iam/organisation/{organisationId}/transfer/{transferId}/actions/accept | Accept iam/organisation.transfer
[**IamOrganisationTransferGet**](IamOrganisationApi.md#IamOrganisationTransferGet) | **Get** /iam/organisation/{organisationId}/transfer/{transferId} | Get iam/organisation.transfer
[**IamOrganisationTransferList**](IamOrganisationApi.md#IamOrganisationTransferList) | **Get** /iam/organisation/{organisationId}/transfer | List iam/organisation.transfer
[**IamOrganisationUpdate**](IamOrganisationApi.md#IamOrganisationUpdate) | **Patch** /iam/organisation/{organisationId} | Update iam/organisation



## IamOrganisationBillingList

> []Billing IamOrganisationBillingList(ctx, organisationId).Start(start).End(end).ResourceType(resourceType).Execute()

List iam/organisation.billing



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
    organisationId := "organisationId_example" // string | Organisation Id
    start := time.Now() // time.Time | start (optional)
    end := time.Now() // time.Time | end (optional)
    resourceType := "resourceType_example" // string | resource.type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationBillingList(context.Background(), organisationId).Start(start).End(end).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationBillingList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationBillingList`: []Billing
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationBillingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationBillingListRequest struct via the builder pattern


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


## IamOrganisationCreate

> Organisation IamOrganisationCreate(ctx).IamOrganisationCreate(iamOrganisationCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create iam/organisation



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
    iamOrganisationCreate := *openapiclient.NewIamOrganisationCreate("Name_example") // IamOrganisationCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationCreate(context.Background()).IamOrganisationCreate(iamOrganisationCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationCreate`: Organisation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamOrganisationCreate** | [**IamOrganisationCreate**](IamOrganisationCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Organisation**](Organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationDelete

> InlineResponseDefault IamOrganisationDelete(ctx, organisationId).Execute()

Delete iam/organisation



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationDelete(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationDelete`: InlineResponseDefault
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponseDefault**](InlineResponseDefault.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationEventGet

> Event IamOrganisationEventGet(ctx, organisationId, eventId).Execute()

Get iam/organisation.event



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
    organisationId := "organisationId_example" // string | Organisation Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationEventGet(context.Background(), organisationId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationEventGetRequest struct via the builder pattern


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


## IamOrganisationEventList

> []Event IamOrganisationEventList(ctx, organisationId).Limit(limit).Skip(skip).Execute()

List iam/organisation.event



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
    organisationId := "organisationId_example" // string | Organisation Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationEventList(context.Background(), organisationId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationEventListRequest struct via the builder pattern


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


## IamOrganisationGet

> Organisation IamOrganisationGet(ctx, organisationId).Execute()

Get iam/organisation



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationGet(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationGet`: Organisation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organisation**](Organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvitationAccept

> Invitation IamOrganisationInvitationAccept(ctx, organisationId, invitationId).IamOrganisationInvitationAccept(iamOrganisationInvitationAccept).Execute()

Accept iam/organisation.invitation



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
    organisationId := "organisationId_example" // string | Organisation Id
    invitationId := "invitationId_example" // string | invitationId
    iamOrganisationInvitationAccept := *openapiclient.NewIamOrganisationInvitationAccept("Token_example") // IamOrganisationInvitationAccept | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationInvitationAccept(context.Background(), organisationId, invitationId).IamOrganisationInvitationAccept(iamOrganisationInvitationAccept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationInvitationAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationInvitationAccept`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationInvitationAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**invitationId** | **string** | invitationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationInvitationAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamOrganisationInvitationAccept** | [**IamOrganisationInvitationAccept**](IamOrganisationInvitationAccept.md) |  | 

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


## IamOrganisationInvitationDelete

> IamOrganisationInvitationDelete(ctx, organisationId, invitationId).Execute()

Delete iam/organisation.invitation



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
    organisationId := "organisationId_example" // string | Organisation Id
    invitationId := "invitationId_example" // string | invitationId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationInvitationDelete(context.Background(), organisationId, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationInvitationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**invitationId** | **string** | invitationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationInvitationDeleteRequest struct via the builder pattern


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


## IamOrganisationInvitationGet

> Invitation IamOrganisationInvitationGet(ctx, organisationId, invitationId).Execute()

Get iam/organisation.invitation



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
    organisationId := "organisationId_example" // string | Organisation Id
    invitationId := "invitationId_example" // string | invitationId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationInvitationGet(context.Background(), organisationId, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationInvitationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationInvitationGet`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationInvitationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**invitationId** | **string** | invitationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationInvitationGetRequest struct via the builder pattern


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


## IamOrganisationInvitationList

> []Invitation IamOrganisationInvitationList(ctx, organisationId).Resource(resource).Execute()

List iam/organisation.invitation



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
    organisationId := "organisationId_example" // string | Organisation Id
    resource := "resource_example" // string | resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationInvitationList(context.Background(), organisationId).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationInvitationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationInvitationList`: []Invitation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationInvitationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationInvitationListRequest struct via the builder pattern


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


## IamOrganisationInvoiceDownload

> *os.File IamOrganisationInvoiceDownload(ctx, organisationId, invoiceId).Execute()

Download iam/organisation.invoice



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
    organisationId := "organisationId_example" // string | Organisation Id
    invoiceId := "invoiceId_example" // string | invoiceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationInvoiceDownload(context.Background(), organisationId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationInvoiceDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationInvoiceDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationInvoiceDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**invoiceId** | **string** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationInvoiceDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvoiceGet

> Invoice IamOrganisationInvoiceGet(ctx, organisationId, invoiceId).Execute()

Get iam/organisation.invoice



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
    organisationId := "organisationId_example" // string | Organisation Id
    invoiceId := "invoiceId_example" // string | invoiceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationInvoiceGet(context.Background(), organisationId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationInvoiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationInvoiceGet`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationInvoiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**invoiceId** | **string** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationInvoiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Invoice**](Invoice.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvoiceList

> []Invoice IamOrganisationInvoiceList(ctx, organisationId).Execute()

List iam/organisation.invoice



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationInvoiceList(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationInvoiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationInvoiceList`: []Invoice
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationInvoiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationInvoiceListRequest struct via the builder pattern


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


## IamOrganisationList

> []Organisation IamOrganisationList(ctx).Name(name).BillingCompany(billingCompany).Limit(limit).Active(active).Execute()

List iam/organisation



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
    billingCompany := "billingCompany_example" // string | Filter by billing.company (optional)
    limit := float32(8.14) // float32 | Filter by $limit (optional)
    active := true // bool | Filter by active (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationList(context.Background()).Name(name).BillingCompany(billingCompany).Limit(limit).Active(active).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationList`: []Organisation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **billingCompany** | **string** | Filter by billing.company | 
 **limit** | **float32** | Filter by $limit | 
 **active** | **bool** | Filter by active | [default to false]

### Return type

[**[]Organisation**](Organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationOwnershipCreate

> Organisation IamOrganisationOwnershipCreate(ctx, organisationId).IamOrganisationOwnershipCreate(iamOrganisationOwnershipCreate).Execute()

Create iam/organisation.ownership



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
    organisationId := "organisationId_example" // string | Organisation Id
    iamOrganisationOwnershipCreate := *openapiclient.NewIamOrganisationOwnershipCreate("Email_example") // IamOrganisationOwnershipCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationOwnershipCreate(context.Background(), organisationId).IamOrganisationOwnershipCreate(iamOrganisationOwnershipCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationOwnershipCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationOwnershipCreate`: Organisation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationOwnershipCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationOwnershipCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamOrganisationOwnershipCreate** | [**IamOrganisationOwnershipCreate**](IamOrganisationOwnershipCreate.md) |  | 

### Return type

[**Organisation**](Organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationOwnershipDelete

> IamOrganisationOwnershipDelete(ctx, organisationId, ownershipId).Execute()

Delete iam/organisation.ownership



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
    organisationId := "organisationId_example" // string | Organisation Id
    ownershipId := "ownershipId_example" // string | ownershipId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationOwnershipDelete(context.Background(), organisationId, ownershipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationOwnershipDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**ownershipId** | **string** | ownershipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationOwnershipDeleteRequest struct via the builder pattern


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


## IamOrganisationOwnershipGet

> Ownership IamOrganisationOwnershipGet(ctx, organisationId, ownershipId).Execute()

Get iam/organisation.ownership



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
    organisationId := "organisationId_example" // string | Organisation Id
    ownershipId := "ownershipId_example" // string | ownershipId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationOwnershipGet(context.Background(), organisationId, ownershipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationOwnershipGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationOwnershipGet`: Ownership
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationOwnershipGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**ownershipId** | **string** | ownershipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationOwnershipGetRequest struct via the builder pattern


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


## IamOrganisationOwnershipList

> []Ownership IamOrganisationOwnershipList(ctx, organisationId).Execute()

List iam/organisation.ownership



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationOwnershipList(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationOwnershipList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationOwnershipList`: []Ownership
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationOwnershipList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationOwnershipListRequest struct via the builder pattern


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


## IamOrganisationPaymentAllocate

> Payment IamOrganisationPaymentAllocate(ctx, organisationId, paymentId).IamOrganisationPaymentAllocate(iamOrganisationPaymentAllocate).Execute()

Allocate iam/organisation.payment



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
    organisationId := "organisationId_example" // string | Organisation Id
    paymentId := "paymentId_example" // string | paymentId
    iamOrganisationPaymentAllocate := *openapiclient.NewIamOrganisationPaymentAllocate("Project_example") // IamOrganisationPaymentAllocate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationPaymentAllocate(context.Background(), organisationId, paymentId).IamOrganisationPaymentAllocate(iamOrganisationPaymentAllocate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationPaymentAllocate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPaymentAllocate`: Payment
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationPaymentAllocate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**paymentId** | **string** | paymentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPaymentAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamOrganisationPaymentAllocate** | [**IamOrganisationPaymentAllocate**](IamOrganisationPaymentAllocate.md) |  | 

### Return type

[**Payment**](Payment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPaymentGet

> Payment IamOrganisationPaymentGet(ctx, organisationId, paymentId).Execute()

Get iam/organisation.payment



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
    organisationId := "organisationId_example" // string | Organisation Id
    paymentId := "paymentId_example" // string | paymentId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationPaymentGet(context.Background(), organisationId, paymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationPaymentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPaymentGet`: Payment
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationPaymentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**paymentId** | **string** | paymentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPaymentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Payment**](Payment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPaymentList

> []Payment IamOrganisationPaymentList(ctx, organisationId).Execute()

List iam/organisation.payment



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationPaymentList(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationPaymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationPaymentList`: []Payment
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationPaymentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationPaymentListRequest struct via the builder pattern


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


## IamOrganisationProformaCreate

> Proforma IamOrganisationProformaCreate(ctx, organisationId).IamOrganisationProformaCreate(iamOrganisationProformaCreate).Execute()

Create iam/organisation.proforma



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
    organisationId := "organisationId_example" // string | Organisation Id
    iamOrganisationProformaCreate := *openapiclient.NewIamOrganisationProformaCreate(float32(123), "Project_example") // IamOrganisationProformaCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationProformaCreate(context.Background(), organisationId).IamOrganisationProformaCreate(iamOrganisationProformaCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationProformaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationProformaCreate`: Proforma
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationProformaCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationProformaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamOrganisationProformaCreate** | [**IamOrganisationProformaCreate**](IamOrganisationProformaCreate.md) |  | 

### Return type

[**Proforma**](Proforma.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationProformaDownload

> *os.File IamOrganisationProformaDownload(ctx, organisationId, proformaId).Execute()

Download iam/organisation.proforma



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
    organisationId := "organisationId_example" // string | Organisation Id
    proformaId := "proformaId_example" // string | proformaId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationProformaDownload(context.Background(), organisationId, proformaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationProformaDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationProformaDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationProformaDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**proformaId** | **string** | proformaId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationProformaDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationProformaGet

> Proforma IamOrganisationProformaGet(ctx, organisationId, proformaId).Execute()

Get iam/organisation.proforma



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
    organisationId := "organisationId_example" // string | Organisation Id
    proformaId := "proformaId_example" // string | proformaId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationProformaGet(context.Background(), organisationId, proformaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationProformaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationProformaGet`: Proforma
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationProformaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**proformaId** | **string** | proformaId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationProformaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Proforma**](Proforma.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationProformaList

> []Proforma IamOrganisationProformaList(ctx, organisationId).Execute()

List iam/organisation.proforma



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationProformaList(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationProformaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationProformaList`: []Proforma
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationProformaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationProformaListRequest struct via the builder pattern


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


## IamOrganisationServiceGet

> ResourceService IamOrganisationServiceGet(ctx, organisationId, serviceId).Execute()

Get iam/organisation.service



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
    organisationId := "organisationId_example" // string | Organisation Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationServiceGet(context.Background(), organisationId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationServiceGetRequest struct via the builder pattern


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


## IamOrganisationServiceList

> []ResourceService IamOrganisationServiceList(ctx, organisationId).Execute()

List iam/organisation.service



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationServiceList(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationServiceListRequest struct via the builder pattern


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


## IamOrganisationTransferAccept

> Transfer IamOrganisationTransferAccept(ctx, organisationId, transferId).IamOrganisationTransferAccept(iamOrganisationTransferAccept).Execute()

Accept iam/organisation.transfer



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
    organisationId := "organisationId_example" // string | Organisation Id
    transferId := "transferId_example" // string | transferId
    iamOrganisationTransferAccept := *openapiclient.NewIamOrganisationTransferAccept("Payment_example") // IamOrganisationTransferAccept | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationTransferAccept(context.Background(), organisationId, transferId).IamOrganisationTransferAccept(iamOrganisationTransferAccept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationTransferAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationTransferAccept`: Transfer
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationTransferAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**transferId** | **string** | transferId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationTransferAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamOrganisationTransferAccept** | [**IamOrganisationTransferAccept**](IamOrganisationTransferAccept.md) |  | 

### Return type

[**Transfer**](Transfer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationTransferGet

> Transfer IamOrganisationTransferGet(ctx, organisationId, transferId).Execute()

Get iam/organisation.transfer



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
    organisationId := "organisationId_example" // string | Organisation Id
    transferId := "transferId_example" // string | transferId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationTransferGet(context.Background(), organisationId, transferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationTransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationTransferGet`: Transfer
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationTransferGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 
**transferId** | **string** | transferId | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Transfer**](Transfer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationTransferList

> []Transfer IamOrganisationTransferList(ctx, organisationId).Execute()

List iam/organisation.transfer



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
    organisationId := "organisationId_example" // string | Organisation Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationTransferList(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationTransferList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationTransferList`: []Transfer
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationTransferList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Transfer**](Transfer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationUpdate

> Organisation IamOrganisationUpdate(ctx, organisationId).IamOrganisationUpdate(iamOrganisationUpdate).Execute()

Update iam/organisation



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
    organisationId := "organisationId_example" // string | Organisation Id
    iamOrganisationUpdate := *openapiclient.NewIamOrganisationUpdate() // IamOrganisationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamOrganisationApi.IamOrganisationUpdate(context.Background(), organisationId).IamOrganisationUpdate(iamOrganisationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamOrganisationApi.IamOrganisationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamOrganisationUpdate`: Organisation
    fmt.Fprintf(os.Stdout, "Response from `IamOrganisationApi.IamOrganisationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamOrganisationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamOrganisationUpdate** | [**IamOrganisationUpdate**](IamOrganisationUpdate.md) |  | 

### Return type

[**Organisation**](Organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

