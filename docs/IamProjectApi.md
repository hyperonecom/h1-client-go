# \IamProjectApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectBillingList**](IamProjectApi.md#ProjectBillingList) | **Get** /iam/project/{projectId}/billing | List iam/project.billing
[**ProjectCreate**](IamProjectApi.md#ProjectCreate) | **Post** /iam/project | Create iam/project
[**ProjectCredentialStoreCreate**](IamProjectApi.md#ProjectCredentialStoreCreate) | **Post** /iam/project/{projectId}/credentialStore | Create iam/project.credentialStore
[**ProjectCredentialStoreDelete**](IamProjectApi.md#ProjectCredentialStoreDelete) | **Delete** /iam/project/{projectId}/credentialStore/{credentialId} | Delete iam/project.credentialStore
[**ProjectCredentialStoreGet**](IamProjectApi.md#ProjectCredentialStoreGet) | **Get** /iam/project/{projectId}/credentialStore/{credentialId} | Get iam/project.credentialStore
[**ProjectCredentialStoreList**](IamProjectApi.md#ProjectCredentialStoreList) | **Get** /iam/project/{projectId}/credentialStore | List iam/project.credentialStore
[**ProjectCredentialStorePatch**](IamProjectApi.md#ProjectCredentialStorePatch) | **Patch** /iam/project/{projectId}/credentialStore/{credentialId} | Update iam/project.credentialStore
[**ProjectDelete**](IamProjectApi.md#ProjectDelete) | **Delete** /iam/project/{projectId} | Delete iam/project
[**ProjectEventGet**](IamProjectApi.md#ProjectEventGet) | **Get** /iam/project/{projectId}/event/{eventId} | Get iam/project.event
[**ProjectEventList**](IamProjectApi.md#ProjectEventList) | **Get** /iam/project/{projectId}/event | List iam/project.event
[**ProjectGet**](IamProjectApi.md#ProjectGet) | **Get** /iam/project/{projectId} | Get iam/project
[**ProjectInvitationAccept**](IamProjectApi.md#ProjectInvitationAccept) | **Post** /iam/project/{projectId}/invitation/{invitationId}/actions/accept | Create iam/project.actions
[**ProjectInvitationDelete**](IamProjectApi.md#ProjectInvitationDelete) | **Delete** /iam/project/{projectId}/invitation/{invitationId} | Delete iam/project.invitation
[**ProjectInvitationGet**](IamProjectApi.md#ProjectInvitationGet) | **Get** /iam/project/{projectId}/invitation/{invitationId} | Get iam/project.invitation
[**ProjectInvitationList**](IamProjectApi.md#ProjectInvitationList) | **Get** /iam/project/{projectId}/invitation | List iam/project.invitation
[**ProjectInvoiceList**](IamProjectApi.md#ProjectInvoiceList) | **Get** /iam/project/{projectId}/invoice | List iam/project.invoice
[**ProjectList**](IamProjectApi.md#ProjectList) | **Get** /iam/project | List iam/project
[**ProjectOwnershipCreate**](IamProjectApi.md#ProjectOwnershipCreate) | **Post** /iam/project/{projectId}/ownership | Create iam/project.ownership
[**ProjectOwnershipDelete**](IamProjectApi.md#ProjectOwnershipDelete) | **Delete** /iam/project/{projectId}/ownership/{ownershipId} | Delete iam/project.ownership
[**ProjectOwnershipGet**](IamProjectApi.md#ProjectOwnershipGet) | **Get** /iam/project/{projectId}/ownership/{ownershipId} | Get iam/project.ownership
[**ProjectOwnershipList**](IamProjectApi.md#ProjectOwnershipList) | **Get** /iam/project/{projectId}/ownership | List iam/project.ownership
[**ProjectPaymentList**](IamProjectApi.md#ProjectPaymentList) | **Get** /iam/project/{projectId}/payment | List iam/project.payment
[**ProjectProformaList**](IamProjectApi.md#ProjectProformaList) | **Get** /iam/project/{projectId}/proforma | List iam/project.proforma
[**ProjectQuotaGet**](IamProjectApi.md#ProjectQuotaGet) | **Get** /iam/project/{projectId}/quota/{quotaId} | Get iam/project.quota
[**ProjectQuotaLimitPatch**](IamProjectApi.md#ProjectQuotaLimitPatch) | **Patch** /iam/project/{projectId}/quota/{quotaId}/limit | Update iam/project.limit
[**ProjectQuotaList**](IamProjectApi.md#ProjectQuotaList) | **Get** /iam/project/{projectId}/quota | List iam/project.quota
[**ProjectServiceGet**](IamProjectApi.md#ProjectServiceGet) | **Get** /iam/project/{projectId}/service/{serviceId} | Get iam/project.service
[**ProjectServiceList**](IamProjectApi.md#ProjectServiceList) | **Get** /iam/project/{projectId}/service | List iam/project.service
[**ProjectTagCreate**](IamProjectApi.md#ProjectTagCreate) | **Post** /iam/project/{projectId}/tag | Create iam/project.tag
[**ProjectTagDelete**](IamProjectApi.md#ProjectTagDelete) | **Delete** /iam/project/{projectId}/tag/{tagId} | Delete iam/project.tag
[**ProjectTagGet**](IamProjectApi.md#ProjectTagGet) | **Get** /iam/project/{projectId}/tag/{tagId} | Get iam/project.tag
[**ProjectTagList**](IamProjectApi.md#ProjectTagList) | **Get** /iam/project/{projectId}/tag | List iam/project.tag
[**ProjectTagPut**](IamProjectApi.md#ProjectTagPut) | **Put** /iam/project/{projectId}/tag | Replace iam/project.tag
[**ProjectUpdate**](IamProjectApi.md#ProjectUpdate) | **Patch** /iam/project/{projectId} | Update iam/project



## ProjectBillingList

> []Billing ProjectBillingList(ctx, projectId, optional)

List iam/project.billing

List iam/project.billing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***ProjectBillingListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectBillingListOpts struct


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


## ProjectCreate

> Project ProjectCreate(ctx, projectCreate, optional)

Create iam/project

Create project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectCreate** | [**ProjectCreate**](ProjectCreate.md)|  | 
 **optional** | ***ProjectCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Project**](project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectCredentialStoreCreate

> ProjectCredential ProjectCredentialStoreCreate(ctx, projectId, projectCredential)

Create iam/project.credentialStore

Create iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**projectCredential** | [**ProjectCredential**](ProjectCredential.md)|  | 

### Return type

[**ProjectCredential**](project.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectCredentialStoreDelete

> Project ProjectCredentialStoreDelete(ctx, projectId, credentialId)

Delete iam/project.credentialStore

Delete iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**Project**](project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectCredentialStoreGet

> ProjectCredential ProjectCredentialStoreGet(ctx, projectId, credentialId)

Get iam/project.credentialStore

Get iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**credentialId** | **string**| credentialId | 

### Return type

[**ProjectCredential**](project.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectCredentialStoreList

> []ProjectCredential ProjectCredentialStoreList(ctx, projectId)

List iam/project.credentialStore

List iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

[**[]ProjectCredential**](project.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectCredentialStorePatch

> ProjectCredential ProjectCredentialStorePatch(ctx, projectId, credentialId, projectCredentialStorePatch)

Update iam/project.credentialStore

Update iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**credentialId** | **string**| credentialId | 
**projectCredentialStorePatch** | [**ProjectCredentialStorePatch**](ProjectCredentialStorePatch.md)|  | 

### Return type

[**ProjectCredential**](project.credential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectDelete

> ProjectDelete(ctx, projectId)

Delete iam/project

Delete project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

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


## ProjectEventGet

> Event ProjectEventGet(ctx, projectId, eventId)

Get iam/project.event

Get iam/project.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## ProjectEventList

> []Event ProjectEventList(ctx, projectId, optional)

List iam/project.event

List iam/project.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***ProjectEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectEventListOpts struct


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


## ProjectGet

> Project ProjectGet(ctx, projectId)

Get iam/project

Returns a single project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

[**Project**](project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectInvitationAccept

> Invitation ProjectInvitationAccept(ctx, projectId, invitationId, projectInvitationAccept)

Create iam/project.actions

Create iam/project.actions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**invitationId** | **string**| invitationId | 
**projectInvitationAccept** | [**ProjectInvitationAccept**](ProjectInvitationAccept.md)|  | 

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


## ProjectInvitationDelete

> ProjectInvitationDelete(ctx, projectId, invitationId)

Delete iam/project.invitation

Delete iam/project.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## ProjectInvitationGet

> Invitation ProjectInvitationGet(ctx, projectId, invitationId)

Get iam/project.invitation

Get iam/project.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## ProjectInvitationList

> []Invitation ProjectInvitationList(ctx, projectId, optional)

List iam/project.invitation

List iam/project.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***ProjectInvitationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectInvitationListOpts struct


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


## ProjectInvoiceList

> []Invoice ProjectInvoiceList(ctx, projectId)

List iam/project.invoice

List iam/project.invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

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


## ProjectList

> []Project ProjectList(ctx, optional)

List iam/project

List project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **limit** | **optional.Float32**| Filter by $limit | 
 **active** | **optional.Bool**| Filter by active | [default to false]
 **organisation** | **optional.String**| Filter by organisation | 
 **lean** | **optional.Bool**| return a lightweight version of the resource | [default to false]
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Project**](project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectOwnershipCreate

> Project ProjectOwnershipCreate(ctx, projectId, projectOwnershipCreate)

Create iam/project.ownership

Create iam/project.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**projectOwnershipCreate** | [**ProjectOwnershipCreate**](ProjectOwnershipCreate.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectOwnershipDelete

> ProjectOwnershipDelete(ctx, projectId, ownershipId)

Delete iam/project.ownership

Delete iam/project.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## ProjectOwnershipGet

> Ownership ProjectOwnershipGet(ctx, projectId, ownershipId)

Get iam/project.ownership

Get iam/project.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## ProjectOwnershipList

> []Ownership ProjectOwnershipList(ctx, projectId)

List iam/project.ownership

List iam/project.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

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


## ProjectPaymentList

> []Payment ProjectPaymentList(ctx, projectId)

List iam/project.payment

List iam/project.payment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

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


## ProjectProformaList

> []Proforma ProjectProformaList(ctx, projectId)

List iam/project.proforma

List iam/project.proforma

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

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


## ProjectQuotaGet

> Quota ProjectQuotaGet(ctx, projectId, quotaId)

Get iam/project.quota

Get iam/project.quota

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**quotaId** | **string**| quotaId | 

### Return type

[**Quota**](quota.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectQuotaLimitPatch

> QuotaLimit ProjectQuotaLimitPatch(ctx, projectId, quotaId, projectQuotaLimitPatch)

Update iam/project.limit

Update iam/project.limit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**quotaId** | **string**| quotaId | 
**projectQuotaLimitPatch** | [**ProjectQuotaLimitPatch**](ProjectQuotaLimitPatch.md)|  | 

### Return type

[**QuotaLimit**](quotaLimit.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectQuotaList

> []Quota ProjectQuotaList(ctx, projectId)

List iam/project.quota

List iam/project.quota

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

[**[]Quota**](quota.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceGet

> ResourceService ProjectServiceGet(ctx, projectId, serviceId)

Get iam/project.service

Get iam/project.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## ProjectServiceList

> []ResourceService ProjectServiceList(ctx, projectId)

List iam/project.service

List iam/project.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

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


## ProjectTagCreate

> Tag ProjectTagCreate(ctx, projectId, tag)

Create iam/project.tag

Create iam/project.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**tag** | [**Tag**](Tag.md)|  | 

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


## ProjectTagDelete

> ProjectTagDelete(ctx, projectId, tagId)

Delete iam/project.tag

Delete iam/project.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**tagId** | **string**| tagId | 

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


## ProjectTagGet

> Tag ProjectTagGet(ctx, projectId, tagId)

Get iam/project.tag

Get iam/project.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**tagId** | **string**| tagId | 

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


## ProjectTagList

> []Tag ProjectTagList(ctx, projectId)

List iam/project.tag

List iam/project.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

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


## ProjectTagPut

> []Tag ProjectTagPut(ctx, projectId, tag)

Replace iam/project.tag

Replace iam/project.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**tag** | [**[]Tag**](tag.md)|  | 

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


## ProjectUpdate

> Project ProjectUpdate(ctx, projectId, projectUpdate)

Update iam/project

Returns modified project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**projectUpdate** | [**ProjectUpdate**](ProjectUpdate.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

