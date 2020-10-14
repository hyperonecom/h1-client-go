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

> Ticket SupportProjectTicketClose(ctx, projectId, ticketId, optional)

Close support/ticket

action close

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**ticketId** | **string**| Ticket Id | 
 **optional** | ***SupportProjectTicketCloseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SupportProjectTicketCloseOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Ticket**](ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketCreate

> Ticket SupportProjectTicketCreate(ctx, projectId, supportProjectTicketCreate, optional)

Create support/ticket

Create ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**supportProjectTicketCreate** | [**SupportProjectTicketCreate**](SupportProjectTicketCreate.md)|  | 
 **optional** | ***SupportProjectTicketCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SupportProjectTicketCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Ticket**](ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketGet

> Ticket SupportProjectTicketGet(ctx, projectId, ticketId)

Get support/ticket

Returns a single ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**ticketId** | **string**| Ticket Id | 

### Return type

[**Ticket**](ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketList

> []Ticket SupportProjectTicketList(ctx, projectId, optional)

List support/ticket

List ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***SupportProjectTicketListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SupportProjectTicketListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**| Filter by state | 

### Return type

[**[]Ticket**](ticket.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketMessageCreate

> SupportMessage SupportProjectTicketMessageCreate(ctx, projectId, ticketId, supportMessage)

Create support/ticket.message

Create support/ticket.message

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**ticketId** | **string**| Ticket Id | 
**supportMessage** | [**SupportMessage**](SupportMessage.md)|  | 

### Return type

[**SupportMessage**](support.message.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketMessageGet

> SupportMessage SupportProjectTicketMessageGet(ctx, projectId, ticketId, messageId)

Get support/ticket.message

Get support/ticket.message

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**ticketId** | **string**| Ticket Id | 
**messageId** | **string**| messageId | 

### Return type

[**SupportMessage**](support.message.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportProjectTicketMessageList

> []SupportMessage SupportProjectTicketMessageList(ctx, projectId, ticketId)

List support/ticket.message

List support/ticket.message

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**ticketId** | **string**| Ticket Id | 

### Return type

[**[]SupportMessage**](support.message.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

