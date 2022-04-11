# \DnsProjectZoneApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsProjectZoneCreate**](DnsProjectZoneApi.md#DnsProjectZoneCreate) | **Post** /dns/{locationId}/project/{projectId}/zone | Create dns/zone
[**DnsProjectZoneDelete**](DnsProjectZoneApi.md#DnsProjectZoneDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId} | Delete dns/zone
[**DnsProjectZoneEventGet**](DnsProjectZoneApi.md#DnsProjectZoneEventGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/event/{eventId} | Get dns/zone.event
[**DnsProjectZoneEventList**](DnsProjectZoneApi.md#DnsProjectZoneEventList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/event | List dns/zone.event
[**DnsProjectZoneGet**](DnsProjectZoneApi.md#DnsProjectZoneGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId} | Get dns/zone
[**DnsProjectZoneList**](DnsProjectZoneApi.md#DnsProjectZoneList) | **Get** /dns/{locationId}/project/{projectId}/zone | List dns/zone
[**DnsProjectZoneRecordsetCreate**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetCreate) | **Post** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset | Create dns/zone.recordset
[**DnsProjectZoneRecordsetDelete**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId} | Delete dns/zone.recordset
[**DnsProjectZoneRecordsetGet**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId} | Get dns/zone.recordset
[**DnsProjectZoneRecordsetList**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset | List dns/zone.recordset
[**DnsProjectZoneRecordsetPatch**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetPatch) | **Patch** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId} | Update dns/zone.recordset
[**DnsProjectZoneRecordsetRecordCreate**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordCreate) | **Post** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record | Create dns/zone.record
[**DnsProjectZoneRecordsetRecordDelete**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record/{recordId} | Delete dns/zone.record
[**DnsProjectZoneRecordsetRecordGet**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record/{recordId} | Get dns/zone.record
[**DnsProjectZoneRecordsetRecordList**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record | List dns/zone.record
[**DnsProjectZoneRecordsetRecordPut**](DnsProjectZoneApi.md#DnsProjectZoneRecordsetRecordPut) | **Put** /dns/{locationId}/project/{projectId}/zone/{zoneId}/recordset/{recordsetId}/record | Replace dns/zone.record
[**DnsProjectZoneServiceGet**](DnsProjectZoneApi.md#DnsProjectZoneServiceGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/service/{serviceId} | Get dns/zone.service
[**DnsProjectZoneServiceList**](DnsProjectZoneApi.md#DnsProjectZoneServiceList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/service | List dns/zone.service
[**DnsProjectZoneTagCreate**](DnsProjectZoneApi.md#DnsProjectZoneTagCreate) | **Post** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag | Create dns/zone.tag
[**DnsProjectZoneTagDelete**](DnsProjectZoneApi.md#DnsProjectZoneTagDelete) | **Delete** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag/{tagId} | Delete dns/zone.tag
[**DnsProjectZoneTagGet**](DnsProjectZoneApi.md#DnsProjectZoneTagGet) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag/{tagId} | Get dns/zone.tag
[**DnsProjectZoneTagList**](DnsProjectZoneApi.md#DnsProjectZoneTagList) | **Get** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag | List dns/zone.tag
[**DnsProjectZoneTagPut**](DnsProjectZoneApi.md#DnsProjectZoneTagPut) | **Put** /dns/{locationId}/project/{projectId}/zone/{zoneId}/tag | Replace dns/zone.tag
[**DnsProjectZoneUpdate**](DnsProjectZoneApi.md#DnsProjectZoneUpdate) | **Patch** /dns/{locationId}/project/{projectId}/zone/{zoneId} | Update dns/zone



## DnsProjectZoneCreate

> Zone DnsProjectZoneCreate(ctx, projectId, locationId).DnsProjectZoneCreate(dnsProjectZoneCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create dns/zone



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
    dnsProjectZoneCreate := *openapiclient.NewDnsProjectZoneCreate("Name_example", "Service_example", "DnsName_example") // DnsProjectZoneCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneCreate(context.Background(), projectId, locationId).DnsProjectZoneCreate(dnsProjectZoneCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneCreate`: Zone
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsProjectZoneCreate** | [**DnsProjectZoneCreate**](DnsProjectZoneCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Zone**](Zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneDelete

> DnsProjectZoneDelete(ctx, projectId, locationId, zoneId).Execute()

Delete dns/zone



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
    zoneId := "zoneId_example" // string | Zone Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneDelete(context.Background(), projectId, locationId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneDelete``: %v\n", err)
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
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneDeleteRequest struct via the builder pattern


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


## DnsProjectZoneEventGet

> Event DnsProjectZoneEventGet(ctx, projectId, locationId, zoneId, eventId).Execute()

Get dns/zone.event



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
    zoneId := "zoneId_example" // string | Zone Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneEventGet(context.Background(), projectId, locationId, zoneId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneEventGetRequest struct via the builder pattern


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


## DnsProjectZoneEventList

> []Event DnsProjectZoneEventList(ctx, projectId, locationId, zoneId).Limit(limit).Skip(skip).Execute()

List dns/zone.event



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
    zoneId := "zoneId_example" // string | Zone Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneEventList(context.Background(), projectId, locationId, zoneId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneEventListRequest struct via the builder pattern


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


## DnsProjectZoneGet

> Zone DnsProjectZoneGet(ctx, projectId, locationId, zoneId).Execute()

Get dns/zone



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
    zoneId := "zoneId_example" // string | Zone Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneGet(context.Background(), projectId, locationId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneGet`: Zone
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Zone**](Zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneList

> []Zone DnsProjectZoneList(ctx, projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()

List dns/zone



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneList(context.Background(), projectId, locationId).Name(name).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneList`: []Zone
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Zone**](Zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetCreate

> DnsRecordset DnsProjectZoneRecordsetCreate(ctx, projectId, locationId, zoneId).DnsRecordset(dnsRecordset).Execute()

Create dns/zone.recordset



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
    zoneId := "zoneId_example" // string | Zone Id
    dnsRecordset := *openapiclient.NewDnsRecordset("Type_example") // DnsRecordset | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetCreate(context.Background(), projectId, locationId, zoneId).DnsRecordset(dnsRecordset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetCreate`: DnsRecordset
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dnsRecordset** | [**DnsRecordset**](DnsRecordset.md) |  | 

### Return type

[**DnsRecordset**](DnsRecordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetDelete

> Zone DnsProjectZoneRecordsetDelete(ctx, projectId, locationId, zoneId, recordsetId).Execute()

Delete dns/zone.recordset



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetDelete(context.Background(), projectId, locationId, zoneId, recordsetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetDelete`: Zone
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Zone**](Zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetGet

> DnsRecordset DnsProjectZoneRecordsetGet(ctx, projectId, locationId, zoneId, recordsetId).Execute()

Get dns/zone.recordset



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetGet(context.Background(), projectId, locationId, zoneId, recordsetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetGet`: DnsRecordset
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DnsRecordset**](DnsRecordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetList

> []DnsRecordset DnsProjectZoneRecordsetList(ctx, projectId, locationId, zoneId).Execute()

List dns/zone.recordset



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
    zoneId := "zoneId_example" // string | Zone Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetList(context.Background(), projectId, locationId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetList`: []DnsRecordset
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]DnsRecordset**](DnsRecordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetPatch

> DnsRecordset DnsProjectZoneRecordsetPatch(ctx, projectId, locationId, zoneId, recordsetId).DnsProjectZoneRecordsetPatch(dnsProjectZoneRecordsetPatch).Execute()

Update dns/zone.recordset



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId
    dnsProjectZoneRecordsetPatch := *openapiclient.NewDnsProjectZoneRecordsetPatch() // DnsProjectZoneRecordsetPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetPatch(context.Background(), projectId, locationId, zoneId, recordsetId).DnsProjectZoneRecordsetPatch(dnsProjectZoneRecordsetPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetPatch`: DnsRecordset
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dnsProjectZoneRecordsetPatch** | [**DnsProjectZoneRecordsetPatch**](DnsProjectZoneRecordsetPatch.md) |  | 

### Return type

[**DnsRecordset**](DnsRecordset.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordCreate

> DnsRecord DnsProjectZoneRecordsetRecordCreate(ctx, projectId, locationId, zoneId, recordsetId).DnsRecord(dnsRecord).Execute()

Create dns/zone.record



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId
    dnsRecord := *openapiclient.NewDnsRecord("Content_example") // DnsRecord | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetRecordCreate(context.Background(), projectId, locationId, zoneId, recordsetId).DnsRecord(dnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetRecordCreate`: DnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetRecordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dnsRecord** | [**DnsRecord**](DnsRecord.md) |  | 

### Return type

[**DnsRecord**](DnsRecord.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordDelete

> DnsProjectZoneRecordsetRecordDelete(ctx, projectId, locationId, zoneId, recordsetId, recordId).Execute()

Delete dns/zone.record



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId
    recordId := "recordId_example" // string | recordId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetRecordDelete(context.Background(), projectId, locationId, zoneId, recordsetId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordDelete``: %v\n", err)
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
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 
**recordId** | **string** | recordId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetRecordDeleteRequest struct via the builder pattern


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


## DnsProjectZoneRecordsetRecordGet

> DnsRecord DnsProjectZoneRecordsetRecordGet(ctx, projectId, locationId, zoneId, recordsetId, recordId).Execute()

Get dns/zone.record



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId
    recordId := "recordId_example" // string | recordId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetRecordGet(context.Background(), projectId, locationId, zoneId, recordsetId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetRecordGet`: DnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 
**recordId** | **string** | recordId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**DnsRecord**](DnsRecord.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordList

> []DnsRecord DnsProjectZoneRecordsetRecordList(ctx, projectId, locationId, zoneId, recordsetId).Execute()

List dns/zone.record



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetRecordList(context.Background(), projectId, locationId, zoneId, recordsetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetRecordList`: []DnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetRecordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]DnsRecord**](DnsRecord.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneRecordsetRecordPut

> []DnsRecord DnsProjectZoneRecordsetRecordPut(ctx, projectId, locationId, zoneId, recordsetId).DnsRecord(dnsRecord).Execute()

Replace dns/zone.record



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
    zoneId := "zoneId_example" // string | Zone Id
    recordsetId := "recordsetId_example" // string | recordsetId
    dnsRecord := []openapiclient.DnsRecord{*openapiclient.NewDnsRecord("Content_example")} // []DnsRecord | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneRecordsetRecordPut(context.Background(), projectId, locationId, zoneId, recordsetId).DnsRecord(dnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneRecordsetRecordPut`: []DnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneRecordsetRecordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**recordsetId** | **string** | recordsetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneRecordsetRecordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dnsRecord** | [**[]DnsRecord**](DnsRecord.md) |  | 

### Return type

[**[]DnsRecord**](DnsRecord.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsProjectZoneServiceGet

> ResourceService DnsProjectZoneServiceGet(ctx, projectId, locationId, zoneId, serviceId).Execute()

Get dns/zone.service



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
    zoneId := "zoneId_example" // string | Zone Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneServiceGet(context.Background(), projectId, locationId, zoneId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneServiceGetRequest struct via the builder pattern


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


## DnsProjectZoneServiceList

> []ResourceService DnsProjectZoneServiceList(ctx, projectId, locationId, zoneId).Execute()

List dns/zone.service



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
    zoneId := "zoneId_example" // string | Zone Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneServiceList(context.Background(), projectId, locationId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneServiceListRequest struct via the builder pattern


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


## DnsProjectZoneTagCreate

> Tag DnsProjectZoneTagCreate(ctx, projectId, locationId, zoneId).Tag(tag).Execute()

Create dns/zone.tag



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
    zoneId := "zoneId_example" // string | Zone Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneTagCreate(context.Background(), projectId, locationId, zoneId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneTagCreateRequest struct via the builder pattern


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


## DnsProjectZoneTagDelete

> DnsProjectZoneTagDelete(ctx, projectId, locationId, zoneId, tagId).Execute()

Delete dns/zone.tag



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
    zoneId := "zoneId_example" // string | Zone Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneTagDelete(context.Background(), projectId, locationId, zoneId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneTagDelete``: %v\n", err)
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
**zoneId** | **string** | Zone Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneTagDeleteRequest struct via the builder pattern


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


## DnsProjectZoneTagGet

> Tag DnsProjectZoneTagGet(ctx, projectId, locationId, zoneId, tagId).Execute()

Get dns/zone.tag



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
    zoneId := "zoneId_example" // string | Zone Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneTagGet(context.Background(), projectId, locationId, zoneId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneTagGetRequest struct via the builder pattern


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


## DnsProjectZoneTagList

> []Tag DnsProjectZoneTagList(ctx, projectId, locationId, zoneId).Execute()

List dns/zone.tag



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
    zoneId := "zoneId_example" // string | Zone Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneTagList(context.Background(), projectId, locationId, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneTagListRequest struct via the builder pattern


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


## DnsProjectZoneTagPut

> []Tag DnsProjectZoneTagPut(ctx, projectId, locationId, zoneId).Tag(tag).Execute()

Replace dns/zone.tag



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
    zoneId := "zoneId_example" // string | Zone Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneTagPut(context.Background(), projectId, locationId, zoneId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneTagPutRequest struct via the builder pattern


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


## DnsProjectZoneUpdate

> Zone DnsProjectZoneUpdate(ctx, projectId, locationId, zoneId).DnsProjectZoneUpdate(dnsProjectZoneUpdate).Execute()

Update dns/zone



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
    zoneId := "zoneId_example" // string | Zone Id
    dnsProjectZoneUpdate := *openapiclient.NewDnsProjectZoneUpdate() // DnsProjectZoneUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsProjectZoneApi.DnsProjectZoneUpdate(context.Background(), projectId, locationId, zoneId).DnsProjectZoneUpdate(dnsProjectZoneUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsProjectZoneApi.DnsProjectZoneUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsProjectZoneUpdate`: Zone
    fmt.Fprintf(os.Stdout, "Response from `DnsProjectZoneApi.DnsProjectZoneUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**zoneId** | **string** | Zone Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsProjectZoneUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dnsProjectZoneUpdate** | [**DnsProjectZoneUpdate**](DnsProjectZoneUpdate.md) |  | 

### Return type

[**Zone**](Zone.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

