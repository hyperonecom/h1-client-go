# \IamOrganisationApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganisationBillingList**](IamOrganisationApi.md#OrganisationBillingList) | **Get** /iam/organisation/{organisationId}/billing | List iam/organisation.billing
[**OrganisationCreate**](IamOrganisationApi.md#OrganisationCreate) | **Post** /iam/organisation | Create iam/organisation
[**OrganisationDelete**](IamOrganisationApi.md#OrganisationDelete) | **Delete** /iam/organisation/{organisationId} | Delete iam/organisation
[**OrganisationEventGet**](IamOrganisationApi.md#OrganisationEventGet) | **Get** /iam/organisation/{organisationId}/event/{eventId} | Get iam/organisation.event
[**OrganisationEventList**](IamOrganisationApi.md#OrganisationEventList) | **Get** /iam/organisation/{organisationId}/event | List iam/organisation.event
[**OrganisationGet**](IamOrganisationApi.md#OrganisationGet) | **Get** /iam/organisation/{organisationId} | Get iam/organisation
[**OrganisationInvitationAccept**](IamOrganisationApi.md#OrganisationInvitationAccept) | **Post** /iam/organisation/{organisationId}/invitation/{invitationId}/actions/accept | Create iam/organisation.actions
[**OrganisationInvitationDelete**](IamOrganisationApi.md#OrganisationInvitationDelete) | **Delete** /iam/organisation/{organisationId}/invitation/{invitationId} | Delete iam/organisation.invitation
[**OrganisationInvitationGet**](IamOrganisationApi.md#OrganisationInvitationGet) | **Get** /iam/organisation/{organisationId}/invitation/{invitationId} | Get iam/organisation.invitation
[**OrganisationInvitationList**](IamOrganisationApi.md#OrganisationInvitationList) | **Get** /iam/organisation/{organisationId}/invitation | List iam/organisation.invitation
[**OrganisationInvoiceDownload**](IamOrganisationApi.md#OrganisationInvoiceDownload) | **Post** /iam/organisation/{organisationId}/invoice/{invoiceId}/actions/download | Create iam/organisation.actions
[**OrganisationInvoiceGet**](IamOrganisationApi.md#OrganisationInvoiceGet) | **Get** /iam/organisation/{organisationId}/invoice/{invoiceId} | Get iam/organisation.invoice
[**OrganisationInvoiceList**](IamOrganisationApi.md#OrganisationInvoiceList) | **Get** /iam/organisation/{organisationId}/invoice | List iam/organisation.invoice
[**OrganisationList**](IamOrganisationApi.md#OrganisationList) | **Get** /iam/organisation | List iam/organisation
[**OrganisationOwnershipCreate**](IamOrganisationApi.md#OrganisationOwnershipCreate) | **Post** /iam/organisation/{organisationId}/ownership | Create iam/organisation.ownership
[**OrganisationOwnershipDelete**](IamOrganisationApi.md#OrganisationOwnershipDelete) | **Delete** /iam/organisation/{organisationId}/ownership/{ownershipId} | Delete iam/organisation.ownership
[**OrganisationOwnershipGet**](IamOrganisationApi.md#OrganisationOwnershipGet) | **Get** /iam/organisation/{organisationId}/ownership/{ownershipId} | Get iam/organisation.ownership
[**OrganisationOwnershipList**](IamOrganisationApi.md#OrganisationOwnershipList) | **Get** /iam/organisation/{organisationId}/ownership | List iam/organisation.ownership
[**OrganisationPaymentAllocate**](IamOrganisationApi.md#OrganisationPaymentAllocate) | **Post** /iam/organisation/{organisationId}/payment/{paymentId}/actions/allocate | Create iam/organisation.actions
[**OrganisationPaymentGet**](IamOrganisationApi.md#OrganisationPaymentGet) | **Get** /iam/organisation/{organisationId}/payment/{paymentId} | Get iam/organisation.payment
[**OrganisationPaymentList**](IamOrganisationApi.md#OrganisationPaymentList) | **Get** /iam/organisation/{organisationId}/payment | List iam/organisation.payment
[**OrganisationProformaCreate**](IamOrganisationApi.md#OrganisationProformaCreate) | **Post** /iam/organisation/{organisationId}/proforma | Create iam/organisation.proforma
[**OrganisationProformaDownload**](IamOrganisationApi.md#OrganisationProformaDownload) | **Post** /iam/organisation/{organisationId}/proforma/{proformaId}/actions/download | Create iam/organisation.actions
[**OrganisationProformaGet**](IamOrganisationApi.md#OrganisationProformaGet) | **Get** /iam/organisation/{organisationId}/proforma/{proformaId} | Get iam/organisation.proforma
[**OrganisationProformaList**](IamOrganisationApi.md#OrganisationProformaList) | **Get** /iam/organisation/{organisationId}/proforma | List iam/organisation.proforma
[**OrganisationTransferAccept**](IamOrganisationApi.md#OrganisationTransferAccept) | **Post** /iam/organisation/{organisationId}/actions/transfer_accept | Transfer accept iam/organisation
[**OrganisationUpdate**](IamOrganisationApi.md#OrganisationUpdate) | **Patch** /iam/organisation/{organisationId} | Update iam/organisation



## OrganisationBillingList

> []Billing OrganisationBillingList(ctx, organisationId, optional)

List iam/organisation.billing

List iam/organisation.billing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***OrganisationBillingListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationBillingListOpts struct


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


## OrganisationCreate

> Organisation OrganisationCreate(ctx, organisationCreate, optional)

Create iam/organisation

Create organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationCreate** | [**OrganisationCreate**](OrganisationCreate.md)|  | 
 **optional** | ***OrganisationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationCreateOpts struct


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


## OrganisationDelete

> OrganisationDelete(ctx, organisationId)

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


## OrganisationEventGet

> Event OrganisationEventGet(ctx, organisationId, eventId)

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


## OrganisationEventList

> []Event OrganisationEventList(ctx, organisationId, optional)

List iam/organisation.event

List iam/organisation.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***OrganisationEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationEventListOpts struct


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


## OrganisationGet

> Organisation OrganisationGet(ctx, organisationId)

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


## OrganisationInvitationAccept

> Invitation OrganisationInvitationAccept(ctx, organisationId, invitationId, organisationInvitationAccept)

Create iam/organisation.actions

Create iam/organisation.actions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**invitationId** | **string**| invitationId | 
**organisationInvitationAccept** | [**OrganisationInvitationAccept**](OrganisationInvitationAccept.md)|  | 

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


## OrganisationInvitationDelete

> OrganisationInvitationDelete(ctx, organisationId, invitationId)

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


## OrganisationInvitationGet

> Invitation OrganisationInvitationGet(ctx, organisationId, invitationId)

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


## OrganisationInvitationList

> []Invitation OrganisationInvitationList(ctx, organisationId, optional)

List iam/organisation.invitation

List iam/organisation.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***OrganisationInvitationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationInvitationListOpts struct


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


## OrganisationInvoiceDownload

> *os.File OrganisationInvoiceDownload(ctx, organisationId, invoiceId)

Create iam/organisation.actions

Create iam/organisation.actions

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


## OrganisationInvoiceGet

> Proforma OrganisationInvoiceGet(ctx, organisationId, invoiceId)

Get iam/organisation.invoice

Get iam/organisation.invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**invoiceId** | **string**| invoiceId | 

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


## OrganisationInvoiceList

> []Invoice OrganisationInvoiceList(ctx, organisationId)

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


## OrganisationList

> []Organisation OrganisationList(ctx, optional)

List iam/organisation

List organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganisationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationListOpts struct


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


## OrganisationOwnershipCreate

> Organisation OrganisationOwnershipCreate(ctx, organisationId, organisationOwnershipCreate)

Create iam/organisation.ownership

Create iam/organisation.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**organisationOwnershipCreate** | [**OrganisationOwnershipCreate**](OrganisationOwnershipCreate.md)|  | 

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


## OrganisationOwnershipDelete

> OrganisationOwnershipDelete(ctx, organisationId, ownershipId)

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


## OrganisationOwnershipGet

> Ownership OrganisationOwnershipGet(ctx, organisationId, ownershipId)

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


## OrganisationOwnershipList

> []Ownership OrganisationOwnershipList(ctx, organisationId)

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


## OrganisationPaymentAllocate

> Payment OrganisationPaymentAllocate(ctx, organisationId, paymentId, organisationPaymentAllocate)

Create iam/organisation.actions

Create iam/organisation.actions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**paymentId** | **string**| paymentId | 
**organisationPaymentAllocate** | [**OrganisationPaymentAllocate**](OrganisationPaymentAllocate.md)|  | 

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


## OrganisationPaymentGet

> Payment OrganisationPaymentGet(ctx, organisationId, paymentId)

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


## OrganisationPaymentList

> []Payment OrganisationPaymentList(ctx, organisationId)

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


## OrganisationProformaCreate

> Proforma OrganisationProformaCreate(ctx, organisationId, organisationProformaCreate)

Create iam/organisation.proforma

Create iam/organisation.proforma

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**organisationProformaCreate** | [**OrganisationProformaCreate**](OrganisationProformaCreate.md)|  | 

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


## OrganisationProformaDownload

> *os.File OrganisationProformaDownload(ctx, organisationId, proformaId)

Create iam/organisation.actions

Create iam/organisation.actions

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


## OrganisationProformaGet

> Proforma OrganisationProformaGet(ctx, organisationId, proformaId)

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


## OrganisationProformaList

> []Proforma OrganisationProformaList(ctx, organisationId)

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


## OrganisationTransferAccept

> Organisation OrganisationTransferAccept(ctx, organisationId, organisationTransferAccept, optional)

Transfer accept iam/organisation

action transfer_accept

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**organisationTransferAccept** | [**OrganisationTransferAccept**](OrganisationTransferAccept.md)|  | 
 **optional** | ***OrganisationTransferAcceptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationTransferAcceptOpts struct


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


## OrganisationUpdate

> Organisation OrganisationUpdate(ctx, organisationId, organisationUpdate)

Update iam/organisation

Returns modified organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**organisationUpdate** | [**OrganisationUpdate**](OrganisationUpdate.md)|  | 

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

