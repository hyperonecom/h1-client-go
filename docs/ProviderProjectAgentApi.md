# \ProviderProjectAgentApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProviderProjectAgentCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentCreate) | **Post** /provider/{locationId}/project/{projectId}/agent | Create provider/agent
[**ProviderProjectAgentCredentialCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialCreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential | Create provider/agent.credential
[**ProviderProjectAgentCredentialDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential/{credentialId} | Delete provider/agent.credential
[**ProviderProjectAgentCredentialGet**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential/{credentialId} | Get provider/agent.credential
[**ProviderProjectAgentCredentialList**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential | List provider/agent.credential
[**ProviderProjectAgentCredentialPatch**](ProviderProjectAgentApi.md#ProviderProjectAgentCredentialPatch) | **Patch** /provider/{locationId}/project/{projectId}/agent/{agentId}/credential/{credentialId} | Update provider/agent.credential
[**ProviderProjectAgentDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId} | Delete provider/agent
[**ProviderProjectAgentEnabledServiceCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceCreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService | Create provider/agent.enabledService
[**ProviderProjectAgentEnabledServiceDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService/{enabledServiceId} | Delete provider/agent.enabledService
[**ProviderProjectAgentEnabledServiceGet**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService/{enabledServiceId} | Get provider/agent.enabledService
[**ProviderProjectAgentEnabledServiceList**](ProviderProjectAgentApi.md#ProviderProjectAgentEnabledServiceList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/enabledService | List provider/agent.enabledService
[**ProviderProjectAgentEventGet**](ProviderProjectAgentApi.md#ProviderProjectAgentEventGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/event/{eventId} | Get provider/agent.event
[**ProviderProjectAgentEventList**](ProviderProjectAgentApi.md#ProviderProjectAgentEventList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/event | List provider/agent.event
[**ProviderProjectAgentGet**](ProviderProjectAgentApi.md#ProviderProjectAgentGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId} | Get provider/agent
[**ProviderProjectAgentInspect**](ProviderProjectAgentApi.md#ProviderProjectAgentInspect) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/inspect | Inspect provider/agent
[**ProviderProjectAgentList**](ProviderProjectAgentApi.md#ProviderProjectAgentList) | **Get** /provider/{locationId}/project/{projectId}/agent | List provider/agent
[**ProviderProjectAgentMetricGet**](ProviderProjectAgentApi.md#ProviderProjectAgentMetricGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/metric/{metricId} | Get provider/agent.metric
[**ProviderProjectAgentMetricList**](ProviderProjectAgentApi.md#ProviderProjectAgentMetricList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/metric | List provider/agent.metric
[**ProviderProjectAgentMetricPointList**](ProviderProjectAgentApi.md#ProviderProjectAgentMetricPointList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/metric/{metricId}/point | List provider/agent.point
[**ProviderProjectAgentResourceEventList**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceEventList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource/{resourceId}/event | List provider/agent.event
[**ProviderProjectAgentResourceGet**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource/{resourceId} | Get provider/agent.resource
[**ProviderProjectAgentResourceInspect**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceInspect) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource/{resourceId}/actions/inspect | Inspect provider/agent.resource
[**ProviderProjectAgentResourceList**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource | List provider/agent.resource
[**ProviderProjectAgentResourceRecreate**](ProviderProjectAgentApi.md#ProviderProjectAgentResourceRecreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/resource/{resourceId}/actions/recreate | Recreate provider/agent.resource
[**ProviderProjectAgentServiceGet**](ProviderProjectAgentApi.md#ProviderProjectAgentServiceGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/service/{serviceId} | Get provider/agent.service
[**ProviderProjectAgentServiceList**](ProviderProjectAgentApi.md#ProviderProjectAgentServiceList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/service | List provider/agent.service
[**ProviderProjectAgentStart**](ProviderProjectAgentApi.md#ProviderProjectAgentStart) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/start | Start provider/agent
[**ProviderProjectAgentSuspend**](ProviderProjectAgentApi.md#ProviderProjectAgentSuspend) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/suspend | Suspend provider/agent
[**ProviderProjectAgentTagCreate**](ProviderProjectAgentApi.md#ProviderProjectAgentTagCreate) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag | Create provider/agent.tag
[**ProviderProjectAgentTagDelete**](ProviderProjectAgentApi.md#ProviderProjectAgentTagDelete) | **Delete** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag/{tagId} | Delete provider/agent.tag
[**ProviderProjectAgentTagGet**](ProviderProjectAgentApi.md#ProviderProjectAgentTagGet) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag/{tagId} | Get provider/agent.tag
[**ProviderProjectAgentTagList**](ProviderProjectAgentApi.md#ProviderProjectAgentTagList) | **Get** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag | List provider/agent.tag
[**ProviderProjectAgentTagPut**](ProviderProjectAgentApi.md#ProviderProjectAgentTagPut) | **Put** /provider/{locationId}/project/{projectId}/agent/{agentId}/tag | Replace provider/agent.tag
[**ProviderProjectAgentTransfer**](ProviderProjectAgentApi.md#ProviderProjectAgentTransfer) | **Post** /provider/{locationId}/project/{projectId}/agent/{agentId}/actions/transfer | Transfer provider/agent
[**ProviderProjectAgentUpdate**](ProviderProjectAgentApi.md#ProviderProjectAgentUpdate) | **Patch** /provider/{locationId}/project/{projectId}/agent/{agentId} | Update provider/agent



## ProviderProjectAgentCreate

> Agent ProviderProjectAgentCreate(ctx, projectId, locationId).ProviderProjectAgentCreate(providerProjectAgentCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Create provider/agent



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
    providerProjectAgentCreate := *openapiclient.NewProviderProjectAgentCreate("Name_example", "Service_example") // ProviderProjectAgentCreate | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentCreate(context.Background(), projectId, locationId).ProviderProjectAgentCreate(providerProjectAgentCreate).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentCreate`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **providerProjectAgentCreate** | [**ProviderProjectAgentCreate**](ProviderProjectAgentCreate.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialCreate

> AgentCredential ProviderProjectAgentCredentialCreate(ctx, projectId, locationId, agentId).AgentCredential(agentCredential).Execute()

Create provider/agent.credential



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
    agentId := "agentId_example" // string | Agent Id
    agentCredential := *openapiclient.NewAgentCredential("Name_example", "Type_example", "Value_example") // AgentCredential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentCredentialCreate(context.Background(), projectId, locationId, agentId).AgentCredential(agentCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentCredentialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentCredentialCreate`: AgentCredential
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentCredentialCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **agentCredential** | [**AgentCredential**](AgentCredential.md) |  | 

### Return type

[**AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialDelete

> Agent ProviderProjectAgentCredentialDelete(ctx, projectId, locationId, agentId, credentialId).Execute()

Delete provider/agent.credential



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
    agentId := "agentId_example" // string | Agent Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentCredentialDelete(context.Background(), projectId, locationId, agentId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentCredentialDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentCredentialDelete`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentCredentialDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentCredentialDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialGet

> AgentCredential ProviderProjectAgentCredentialGet(ctx, projectId, locationId, agentId, credentialId).Execute()

Get provider/agent.credential



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
    agentId := "agentId_example" // string | Agent Id
    credentialId := "credentialId_example" // string | credentialId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentCredentialGet(context.Background(), projectId, locationId, agentId, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentCredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentCredentialGet`: AgentCredential
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentCredentialGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentCredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialList

> []AgentCredential ProviderProjectAgentCredentialList(ctx, projectId, locationId, agentId).Execute()

List provider/agent.credential



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentCredentialList(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentCredentialList`: []AgentCredential
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentCredentialPatch

> AgentCredential ProviderProjectAgentCredentialPatch(ctx, projectId, locationId, agentId, credentialId).ProviderProjectAgentCredentialPatch(providerProjectAgentCredentialPatch).Execute()

Update provider/agent.credential



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
    agentId := "agentId_example" // string | Agent Id
    credentialId := "credentialId_example" // string | credentialId
    providerProjectAgentCredentialPatch := *openapiclient.NewProviderProjectAgentCredentialPatch("Name_example") // ProviderProjectAgentCredentialPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentCredentialPatch(context.Background(), projectId, locationId, agentId, credentialId).ProviderProjectAgentCredentialPatch(providerProjectAgentCredentialPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentCredentialPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentCredentialPatch`: AgentCredential
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentCredentialPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**credentialId** | **string** | credentialId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentCredentialPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **providerProjectAgentCredentialPatch** | [**ProviderProjectAgentCredentialPatch**](ProviderProjectAgentCredentialPatch.md) |  | 

### Return type

[**AgentCredential**](agent.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentDelete

> ProviderProjectAgentDelete(ctx, projectId, locationId, agentId).Execute()

Delete provider/agent



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentDelete(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentDelete``: %v\n", err)
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
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentDeleteRequest struct via the builder pattern


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


## ProviderProjectAgentEnabledServiceCreate

> EnabledService ProviderProjectAgentEnabledServiceCreate(ctx, projectId, locationId, agentId).EnabledService(enabledService).Execute()

Create provider/agent.enabledService



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
    agentId := "agentId_example" // string | Agent Id
    enabledService := *openapiclient.NewEnabledService("Id_example", "Name_example") // EnabledService | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceCreate(context.Background(), projectId, locationId, agentId).EnabledService(enabledService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentEnabledServiceCreate`: EnabledService
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentEnabledServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enabledService** | [**EnabledService**](EnabledService.md) |  | 

### Return type

[**EnabledService**](enabledService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEnabledServiceDelete

> Agent ProviderProjectAgentEnabledServiceDelete(ctx, projectId, locationId, agentId, enabledServiceId).Execute()

Delete provider/agent.enabledService



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
    agentId := "agentId_example" // string | Agent Id
    enabledServiceId := "enabledServiceId_example" // string | enabledServiceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceDelete(context.Background(), projectId, locationId, agentId, enabledServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentEnabledServiceDelete`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**enabledServiceId** | **string** | enabledServiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentEnabledServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEnabledServiceGet

> EnabledService ProviderProjectAgentEnabledServiceGet(ctx, projectId, locationId, agentId, enabledServiceId).Execute()

Get provider/agent.enabledService



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
    agentId := "agentId_example" // string | Agent Id
    enabledServiceId := "enabledServiceId_example" // string | enabledServiceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceGet(context.Background(), projectId, locationId, agentId, enabledServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentEnabledServiceGet`: EnabledService
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**enabledServiceId** | **string** | enabledServiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentEnabledServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**EnabledService**](enabledService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEnabledServiceList

> []EnabledService ProviderProjectAgentEnabledServiceList(ctx, projectId, locationId, agentId).Execute()

List provider/agent.enabledService



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceList(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentEnabledServiceList`: []EnabledService
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentEnabledServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentEnabledServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]EnabledService**](enabledService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentEventGet

> Event ProviderProjectAgentEventGet(ctx, projectId, locationId, agentId, eventId).Execute()

Get provider/agent.event



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
    agentId := "agentId_example" // string | Agent Id
    eventId := "eventId_example" // string | eventId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentEventGet(context.Background(), projectId, locationId, agentId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentEventGet`: Event
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**eventId** | **string** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentEventGetRequest struct via the builder pattern


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


## ProviderProjectAgentEventList

> []Event ProviderProjectAgentEventList(ctx, projectId, locationId, agentId).Limit(limit).Skip(skip).Execute()

List provider/agent.event



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
    agentId := "agentId_example" // string | Agent Id
    limit := float32(8.14) // float32 | $limit (optional) (default to 100)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentEventList(context.Background(), projectId, locationId, agentId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentEventList`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentEventListRequest struct via the builder pattern


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


## ProviderProjectAgentGet

> Agent ProviderProjectAgentGet(ctx, projectId, locationId, agentId).Execute()

Get provider/agent



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentGet(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentGet`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentInspect

> map[string]interface{} ProviderProjectAgentInspect(ctx, projectId, locationId, agentId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Inspect provider/agent



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
    agentId := "agentId_example" // string | Agent Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentInspect(context.Background(), projectId, locationId, agentId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentInspect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentInspect`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentInspect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentInspectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

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


## ProviderProjectAgentList

> []Agent ProviderProjectAgentList(ctx, projectId, locationId).Name(name).EnabledServices(enabledServices).TagValue(tagValue).TagKey(tagKey).Execute()

List provider/agent



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
    enabledServices := "enabledServices_example" // string | Filter by enabledServices (optional)
    tagValue := "tagValue_example" // string | Filter by tag.value (optional)
    tagKey := "tagKey_example" // string | Filter by tag.key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentList(context.Background(), projectId, locationId).Name(name).EnabledServices(enabledServices).TagValue(tagValue).TagKey(tagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentList`: []Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Filter by name | 
 **enabledServices** | **string** | Filter by enabledServices | 
 **tagValue** | **string** | Filter by tag.value | 
 **tagKey** | **string** | Filter by tag.key | 

### Return type

[**[]Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentMetricGet

> Metric ProviderProjectAgentMetricGet(ctx, projectId, locationId, agentId, metricId).Execute()

Get provider/agent.metric



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
    agentId := "agentId_example" // string | Agent Id
    metricId := "metricId_example" // string | metricId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentMetricGet(context.Background(), projectId, locationId, agentId, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentMetricGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentMetricGet`: Metric
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentMetricGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentMetricGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Metric**](metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentMetricList

> []Metric ProviderProjectAgentMetricList(ctx, projectId, locationId, agentId).Execute()

List provider/agent.metric



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentMetricList(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentMetricList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentMetricList`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentMetricList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentMetricListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Metric**](metric.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentMetricPointList

> []Point ProviderProjectAgentMetricPointList(ctx, projectId, locationId, agentId, metricId).Interval(interval).Timespan(timespan).Execute()

List provider/agent.point



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
    agentId := "agentId_example" // string | Agent Id
    metricId := "metricId_example" // string | metricId
    interval := "interval_example" // string | interval (optional)
    timespan := "timespan_example" // string | timespan (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentMetricPointList(context.Background(), projectId, locationId, agentId, metricId).Interval(interval).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentMetricPointList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentMetricPointList`: []Point
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentMetricPointList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**metricId** | **string** | metricId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentMetricPointListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **interval** | **string** | interval | 
 **timespan** | **string** | timespan | 

### Return type

[**[]Point**](point.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceEventList

> []ProviderAgentResourceEvent ProviderProjectAgentResourceEventList(ctx, projectId, locationId, agentId, resourceId).Limit(limit).Skip(skip).Execute()

List provider/agent.event



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
    agentId := "agentId_example" // string | Agent Id
    resourceId := "resourceId_example" // string | resourceId
    limit := float32(8.14) // float32 | $limit (optional)
    skip := float32(8.14) // float32 | $skip (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentResourceEventList(context.Background(), projectId, locationId, agentId, resourceId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentResourceEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentResourceEventList`: []ProviderAgentResourceEvent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentResourceEventList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**resourceId** | **string** | resourceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentResourceEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **float32** | $limit | 
 **skip** | **float32** | $skip | 

### Return type

[**[]ProviderAgentResourceEvent**](provider.agent.resource.event.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceGet

> ProviderAgentResource ProviderProjectAgentResourceGet(ctx, projectId, locationId, agentId, resourceId).Execute()

Get provider/agent.resource



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
    agentId := "agentId_example" // string | Agent Id
    resourceId := "resourceId_example" // string | resourceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentResourceGet(context.Background(), projectId, locationId, agentId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentResourceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentResourceGet`: ProviderAgentResource
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentResourceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**resourceId** | **string** | resourceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentResourceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ProviderAgentResource**](provider.agent.resource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceInspect

> interface{} ProviderProjectAgentResourceInspect(ctx, projectId, locationId, agentId, resourceId).Execute()

Inspect provider/agent.resource



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
    agentId := "agentId_example" // string | Agent Id
    resourceId := "resourceId_example" // string | resourceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentResourceInspect(context.Background(), projectId, locationId, agentId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentResourceInspect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentResourceInspect`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentResourceInspect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**resourceId** | **string** | resourceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentResourceInspectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

**interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceList

> []ProviderAgentResource ProviderProjectAgentResourceList(ctx, projectId, locationId, agentId).Execute()

List provider/agent.resource



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentResourceList(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentResourceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentResourceList`: []ProviderAgentResource
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentResourceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentResourceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]ProviderAgentResource**](provider.agent.resource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentResourceRecreate

> ProviderAgentResource ProviderProjectAgentResourceRecreate(ctx, projectId, locationId, agentId, resourceId).Execute()

Recreate provider/agent.resource



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
    agentId := "agentId_example" // string | Agent Id
    resourceId := "resourceId_example" // string | resourceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentResourceRecreate(context.Background(), projectId, locationId, agentId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentResourceRecreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentResourceRecreate`: ProviderAgentResource
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentResourceRecreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**resourceId** | **string** | resourceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentResourceRecreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ProviderAgentResource**](provider.agent.resource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentServiceGet

> ResourceService ProviderProjectAgentServiceGet(ctx, projectId, locationId, agentId, serviceId).Execute()

Get provider/agent.service



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
    agentId := "agentId_example" // string | Agent Id
    serviceId := "serviceId_example" // string | serviceId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentServiceGet(context.Background(), projectId, locationId, agentId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentServiceGet`: ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**serviceId** | **string** | serviceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentServiceGetRequest struct via the builder pattern


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


## ProviderProjectAgentServiceList

> []ResourceService ProviderProjectAgentServiceList(ctx, projectId, locationId, agentId).Execute()

List provider/agent.service



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentServiceList(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentServiceList`: []ResourceService
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentServiceListRequest struct via the builder pattern


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


## ProviderProjectAgentStart

> Agent ProviderProjectAgentStart(ctx, projectId, locationId, agentId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Start provider/agent



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
    agentId := "agentId_example" // string | Agent Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentStart(context.Background(), projectId, locationId, agentId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentStart`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentSuspend

> Agent ProviderProjectAgentSuspend(ctx, projectId, locationId, agentId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Suspend provider/agent



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
    agentId := "agentId_example" // string | Agent Id
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentSuspend(context.Background(), projectId, locationId, agentId).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentSuspend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentSuspend`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentSuspend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentTagCreate

> Tag ProviderProjectAgentTagCreate(ctx, projectId, locationId, agentId).Tag(tag).Execute()

Create provider/agent.tag



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
    agentId := "agentId_example" // string | Agent Id
    tag := *openapiclient.NewTag("Id_example", "Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentTagCreate(context.Background(), projectId, locationId, agentId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentTagCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentTagCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentTagCreateRequest struct via the builder pattern


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


## ProviderProjectAgentTagDelete

> ProviderProjectAgentTagDelete(ctx, projectId, locationId, agentId, tagId).Execute()

Delete provider/agent.tag



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
    agentId := "agentId_example" // string | Agent Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentTagDelete(context.Background(), projectId, locationId, agentId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentTagDelete``: %v\n", err)
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
**agentId** | **string** | Agent Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentTagDeleteRequest struct via the builder pattern


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


## ProviderProjectAgentTagGet

> Tag ProviderProjectAgentTagGet(ctx, projectId, locationId, agentId, tagId).Execute()

Get provider/agent.tag



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
    agentId := "agentId_example" // string | Agent Id
    tagId := "tagId_example" // string | tagId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentTagGet(context.Background(), projectId, locationId, agentId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentTagGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentTagGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentTagGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 
**tagId** | **string** | tagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentTagGetRequest struct via the builder pattern


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


## ProviderProjectAgentTagList

> []Tag ProviderProjectAgentTagList(ctx, projectId, locationId, agentId).Execute()

List provider/agent.tag



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
    agentId := "agentId_example" // string | Agent Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentTagList(context.Background(), projectId, locationId, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentTagList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentTagList`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentTagList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentTagListRequest struct via the builder pattern


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


## ProviderProjectAgentTagPut

> []Tag ProviderProjectAgentTagPut(ctx, projectId, locationId, agentId).Tag(tag).Execute()

Replace provider/agent.tag



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
    agentId := "agentId_example" // string | Agent Id
    tag := []openapiclient.Tag{*openapiclient.NewTag("Id_example", "Key_example", "Value_example")} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentTagPut(context.Background(), projectId, locationId, agentId).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentTagPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentTagPut`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentTagPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentTagPutRequest struct via the builder pattern


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


## ProviderProjectAgentTransfer

> Agent ProviderProjectAgentTransfer(ctx, projectId, locationId, agentId).ProviderProjectAgentTransfer(providerProjectAgentTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()

Transfer provider/agent



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
    agentId := "agentId_example" // string | Agent Id
    providerProjectAgentTransfer := *openapiclient.NewProviderProjectAgentTransfer("Project_example") // ProviderProjectAgentTransfer | 
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency key (optional)
    xDryRun := "xDryRun_example" // string | Dry run (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentTransfer(context.Background(), projectId, locationId, agentId).ProviderProjectAgentTransfer(providerProjectAgentTransfer).XIdempotencyKey(xIdempotencyKey).XDryRun(xDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentTransfer`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **providerProjectAgentTransfer** | [**ProviderProjectAgentTransfer**](ProviderProjectAgentTransfer.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency key | 
 **xDryRun** | **string** | Dry run | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderProjectAgentUpdate

> Agent ProviderProjectAgentUpdate(ctx, projectId, locationId, agentId).ProviderProjectAgentUpdate(providerProjectAgentUpdate).Execute()

Update provider/agent



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
    agentId := "agentId_example" // string | Agent Id
    providerProjectAgentUpdate := *openapiclient.NewProviderProjectAgentUpdate() // ProviderProjectAgentUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderProjectAgentApi.ProviderProjectAgentUpdate(context.Background(), projectId, locationId, agentId).ProviderProjectAgentUpdate(providerProjectAgentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderProjectAgentApi.ProviderProjectAgentUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderProjectAgentUpdate`: Agent
    fmt.Fprintf(os.Stdout, "Response from `ProviderProjectAgentApi.ProviderProjectAgentUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project Id | 
**locationId** | **string** | Location Id | 
**agentId** | **string** | Agent Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderProjectAgentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **providerProjectAgentUpdate** | [**ProviderProjectAgentUpdate**](ProviderProjectAgentUpdate.md) |  | 

### Return type

[**Agent**](agent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

