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

> []Billing IamOrganisationBillingList(ctx, organisationId, optional)

List iam/organisation.billing

List iam/organisation.billing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***IamOrganisationBillingListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationBillingListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Time**| start | 
 **end** | **optional.Time**| end | 
 **resourceType** | **optional.String**| resource.type | 

### Return type

[**[]Billing**](billing.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationCreate

> Organisation IamOrganisationCreate(ctx, iamOrganisationCreate, optional)

Create iam/organisation

Create organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iamOrganisationCreate** | [**IamOrganisationCreate**](IamOrganisationCreate.md)|  | 
 **optional** | ***IamOrganisationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationDelete

> IamOrganisationDelete(ctx, organisationId)

Delete iam/organisation

Delete organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

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


## IamOrganisationEventGet

> Event IamOrganisationEventGet(ctx, organisationId, eventId)

Get iam/organisation.event

Get iam/organisation.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**eventId** | **string**| eventId | 

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


## IamOrganisationEventList

> []Event IamOrganisationEventList(ctx, organisationId, optional)

List iam/organisation.event

List iam/organisation.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***IamOrganisationEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationEventListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Float32**| $limit | [default to 100]
 **skip** | **optional.Float32**| $skip | 

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


## IamOrganisationGet

> Organisation IamOrganisationGet(ctx, organisationId)

Get iam/organisation

Returns a single organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvitationAccept

> Invitation IamOrganisationInvitationAccept(ctx, organisationId, invitationId, iamOrganisationInvitationAccept)

Accept iam/organisation.invitation

action accept

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**invitationId** | **string**| invitationId | 
**iamOrganisationInvitationAccept** | [**IamOrganisationInvitationAccept**](IamOrganisationInvitationAccept.md)|  | 

### Return type

[**Invitation**](invitation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvitationDelete

> IamOrganisationInvitationDelete(ctx, organisationId, invitationId)

Delete iam/organisation.invitation

Delete iam/organisation.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**invitationId** | **string**| invitationId | 

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

> Invitation IamOrganisationInvitationGet(ctx, organisationId, invitationId)

Get iam/organisation.invitation

Get iam/organisation.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**invitationId** | **string**| invitationId | 

### Return type

[**Invitation**](invitation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvitationList

> []Invitation IamOrganisationInvitationList(ctx, organisationId, optional)

List iam/organisation.invitation

List iam/organisation.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***IamOrganisationInvitationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationInvitationListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | **optional.String**| resource | 

### Return type

[**[]Invitation**](invitation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvoiceDownload

> *os.File IamOrganisationInvoiceDownload(ctx, organisationId, invoiceId)

Download iam/organisation.invoice

action download

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**invoiceId** | **string**| invoiceId | 

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

> Invoice IamOrganisationInvoiceGet(ctx, organisationId, invoiceId)

Get iam/organisation.invoice

Get iam/organisation.invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**invoiceId** | **string**| invoiceId | 

### Return type

[**Invoice**](invoice.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationInvoiceList

> []Invoice IamOrganisationInvoiceList(ctx, organisationId)

List iam/organisation.invoice

List iam/organisation.invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

### Return type

[**[]Invoice**](invoice.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationList

> []Organisation IamOrganisationList(ctx, optional)

List iam/organisation

List organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IamOrganisationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **billingCompany** | **optional.String**| Filter by billing.company | 
 **limit** | **optional.Float32**| Filter by $limit | 
 **active** | **optional.Bool**| Filter by active | [default to false]

### Return type

[**[]Organisation**](organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationOwnershipCreate

> Organisation IamOrganisationOwnershipCreate(ctx, organisationId, iamOrganisationOwnershipCreate)

Create iam/organisation.ownership

Create iam/organisation.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**iamOrganisationOwnershipCreate** | [**IamOrganisationOwnershipCreate**](IamOrganisationOwnershipCreate.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationOwnershipDelete

> IamOrganisationOwnershipDelete(ctx, organisationId, ownershipId)

Delete iam/organisation.ownership

Delete iam/organisation.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**ownershipId** | **string**| ownershipId | 

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

> Ownership IamOrganisationOwnershipGet(ctx, organisationId, ownershipId)

Get iam/organisation.ownership

Get iam/organisation.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**ownershipId** | **string**| ownershipId | 

### Return type

[**Ownership**](ownership.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationOwnershipList

> []Ownership IamOrganisationOwnershipList(ctx, organisationId)

List iam/organisation.ownership

List iam/organisation.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

### Return type

[**[]Ownership**](ownership.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPaymentAllocate

> Payment IamOrganisationPaymentAllocate(ctx, organisationId, paymentId, iamOrganisationPaymentAllocate)

Allocate iam/organisation.payment

action allocate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**paymentId** | **string**| paymentId | 
**iamOrganisationPaymentAllocate** | [**IamOrganisationPaymentAllocate**](IamOrganisationPaymentAllocate.md)|  | 

### Return type

[**Payment**](payment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPaymentGet

> Payment IamOrganisationPaymentGet(ctx, organisationId, paymentId)

Get iam/organisation.payment

Get iam/organisation.payment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**paymentId** | **string**| paymentId | 

### Return type

[**Payment**](payment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationPaymentList

> []Payment IamOrganisationPaymentList(ctx, organisationId)

List iam/organisation.payment

List iam/organisation.payment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

### Return type

[**[]Payment**](payment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationProformaCreate

> Proforma IamOrganisationProformaCreate(ctx, organisationId, iamOrganisationProformaCreate)

Create iam/organisation.proforma

Create iam/organisation.proforma

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**iamOrganisationProformaCreate** | [**IamOrganisationProformaCreate**](IamOrganisationProformaCreate.md)|  | 

### Return type

[**Proforma**](proforma.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationProformaDownload

> *os.File IamOrganisationProformaDownload(ctx, organisationId, proformaId)

Download iam/organisation.proforma

action download

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**proformaId** | **string**| proformaId | 

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

> Proforma IamOrganisationProformaGet(ctx, organisationId, proformaId)

Get iam/organisation.proforma

Get iam/organisation.proforma

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**proformaId** | **string**| proformaId | 

### Return type

[**Proforma**](proforma.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationProformaList

> []Proforma IamOrganisationProformaList(ctx, organisationId)

List iam/organisation.proforma

List iam/organisation.proforma

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

### Return type

[**[]Proforma**](proforma.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationServiceGet

> ResourceService IamOrganisationServiceGet(ctx, organisationId, serviceId)

Get iam/organisation.service

Get iam/organisation.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**serviceId** | **string**| serviceId | 

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


## IamOrganisationServiceList

> []ResourceService IamOrganisationServiceList(ctx, organisationId)

List iam/organisation.service

List iam/organisation.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

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


## IamOrganisationTransferAccept

> Transfer IamOrganisationTransferAccept(ctx, organisationId, transferId, iamOrganisationTransferAccept)

Accept iam/organisation.transfer

action accept

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**transferId** | **string**| transferId | 
**iamOrganisationTransferAccept** | [**IamOrganisationTransferAccept**](IamOrganisationTransferAccept.md)|  | 

### Return type

[**Transfer**](transfer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationTransferGet

> Transfer IamOrganisationTransferGet(ctx, organisationId, transferId)

Get iam/organisation.transfer

Get iam/organisation.transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**transferId** | **string**| transferId | 

### Return type

[**Transfer**](transfer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationTransferList

> []Transfer IamOrganisationTransferList(ctx, organisationId)

List iam/organisation.transfer

List iam/organisation.transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 

### Return type

[**[]Transfer**](transfer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationUpdate

> Organisation IamOrganisationUpdate(ctx, organisationId, iamOrganisationUpdate)

Update iam/organisation

Returns modified organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**iamOrganisationUpdate** | [**IamOrganisationUpdate**](IamOrganisationUpdate.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

