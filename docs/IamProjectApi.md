# \IamProjectApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectBillingList**](IamProjectApi.md#IamProjectBillingList) | **Get** /iam/project/{projectId}/billing | List iam/project.billing
[**IamProjectCreate**](IamProjectApi.md#IamProjectCreate) | **Post** /iam/project | Create iam/project
[**IamProjectCredentialStoreCreate**](IamProjectApi.md#IamProjectCredentialStoreCreate) | **Post** /iam/project/{projectId}/credentialStore | Create iam/project.credentialStore
[**IamProjectCredentialStoreDelete**](IamProjectApi.md#IamProjectCredentialStoreDelete) | **Delete** /iam/project/{projectId}/credentialStore/{credentialStoreId} | Delete iam/project.credentialStore
[**IamProjectCredentialStoreGet**](IamProjectApi.md#IamProjectCredentialStoreGet) | **Get** /iam/project/{projectId}/credentialStore/{credentialStoreId} | Get iam/project.credentialStore
[**IamProjectCredentialStoreList**](IamProjectApi.md#IamProjectCredentialStoreList) | **Get** /iam/project/{projectId}/credentialStore | List iam/project.credentialStore
[**IamProjectCredentialStorePatch**](IamProjectApi.md#IamProjectCredentialStorePatch) | **Patch** /iam/project/{projectId}/credentialStore/{credentialStoreId} | Update iam/project.credentialStore
[**IamProjectDelete**](IamProjectApi.md#IamProjectDelete) | **Delete** /iam/project/{projectId} | Delete iam/project
[**IamProjectEventGet**](IamProjectApi.md#IamProjectEventGet) | **Get** /iam/project/{projectId}/event/{eventId} | Get iam/project.event
[**IamProjectEventList**](IamProjectApi.md#IamProjectEventList) | **Get** /iam/project/{projectId}/event | List iam/project.event
[**IamProjectGet**](IamProjectApi.md#IamProjectGet) | **Get** /iam/project/{projectId} | Get iam/project
[**IamProjectInvitationAccept**](IamProjectApi.md#IamProjectInvitationAccept) | **Post** /iam/project/{projectId}/invitation/{invitationId}/actions/accept | Accept iam/project.invitation
[**IamProjectInvitationDelete**](IamProjectApi.md#IamProjectInvitationDelete) | **Delete** /iam/project/{projectId}/invitation/{invitationId} | Delete iam/project.invitation
[**IamProjectInvitationGet**](IamProjectApi.md#IamProjectInvitationGet) | **Get** /iam/project/{projectId}/invitation/{invitationId} | Get iam/project.invitation
[**IamProjectInvitationList**](IamProjectApi.md#IamProjectInvitationList) | **Get** /iam/project/{projectId}/invitation | List iam/project.invitation
[**IamProjectInvoiceList**](IamProjectApi.md#IamProjectInvoiceList) | **Get** /iam/project/{projectId}/invoice | List iam/project.invoice
[**IamProjectList**](IamProjectApi.md#IamProjectList) | **Get** /iam/project | List iam/project
[**IamProjectOwnershipCreate**](IamProjectApi.md#IamProjectOwnershipCreate) | **Post** /iam/project/{projectId}/ownership | Create iam/project.ownership
[**IamProjectOwnershipDelete**](IamProjectApi.md#IamProjectOwnershipDelete) | **Delete** /iam/project/{projectId}/ownership/{ownershipId} | Delete iam/project.ownership
[**IamProjectOwnershipGet**](IamProjectApi.md#IamProjectOwnershipGet) | **Get** /iam/project/{projectId}/ownership/{ownershipId} | Get iam/project.ownership
[**IamProjectOwnershipList**](IamProjectApi.md#IamProjectOwnershipList) | **Get** /iam/project/{projectId}/ownership | List iam/project.ownership
[**IamProjectPaymentList**](IamProjectApi.md#IamProjectPaymentList) | **Get** /iam/project/{projectId}/payment | List iam/project.payment
[**IamProjectProformaList**](IamProjectApi.md#IamProjectProformaList) | **Get** /iam/project/{projectId}/proforma | List iam/project.proforma
[**IamProjectQuotaGet**](IamProjectApi.md#IamProjectQuotaGet) | **Get** /iam/project/{projectId}/quota/{quotaId} | Get iam/project.quota
[**IamProjectQuotaLimitPatch**](IamProjectApi.md#IamProjectQuotaLimitPatch) | **Patch** /iam/project/{projectId}/quota/{quotaId}/limit | Update iam/project.limit
[**IamProjectQuotaList**](IamProjectApi.md#IamProjectQuotaList) | **Get** /iam/project/{projectId}/quota | List iam/project.quota
[**IamProjectServiceGet**](IamProjectApi.md#IamProjectServiceGet) | **Get** /iam/project/{projectId}/service/{serviceId} | Get iam/project.service
[**IamProjectServiceList**](IamProjectApi.md#IamProjectServiceList) | **Get** /iam/project/{projectId}/service | List iam/project.service
[**IamProjectTagCreate**](IamProjectApi.md#IamProjectTagCreate) | **Post** /iam/project/{projectId}/tag | Create iam/project.tag
[**IamProjectTagDelete**](IamProjectApi.md#IamProjectTagDelete) | **Delete** /iam/project/{projectId}/tag/{tagId} | Delete iam/project.tag
[**IamProjectTagGet**](IamProjectApi.md#IamProjectTagGet) | **Get** /iam/project/{projectId}/tag/{tagId} | Get iam/project.tag
[**IamProjectTagList**](IamProjectApi.md#IamProjectTagList) | **Get** /iam/project/{projectId}/tag | List iam/project.tag
[**IamProjectTagPut**](IamProjectApi.md#IamProjectTagPut) | **Put** /iam/project/{projectId}/tag | Replace iam/project.tag
[**IamProjectThresholdCreate**](IamProjectApi.md#IamProjectThresholdCreate) | **Post** /iam/project/{projectId}/threshold | Create iam/project.threshold
[**IamProjectThresholdDelete**](IamProjectApi.md#IamProjectThresholdDelete) | **Delete** /iam/project/{projectId}/threshold/{thresholdId} | Delete iam/project.threshold
[**IamProjectThresholdGet**](IamProjectApi.md#IamProjectThresholdGet) | **Get** /iam/project/{projectId}/threshold/{thresholdId} | Get iam/project.threshold
[**IamProjectThresholdList**](IamProjectApi.md#IamProjectThresholdList) | **Get** /iam/project/{projectId}/threshold | List iam/project.threshold
[**IamProjectTransfer**](IamProjectApi.md#IamProjectTransfer) | **Post** /iam/project/{projectId}/actions/transfer | Transfer iam/project
[**IamProjectUpdate**](IamProjectApi.md#IamProjectUpdate) | **Patch** /iam/project/{projectId} | Update iam/project



## IamProjectBillingList

> []Billing IamProjectBillingList(ctx, projectId, optional)

List iam/project.billing

List iam/project.billing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***IamProjectBillingListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectBillingListOpts struct


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


## IamProjectCreate

> Project IamProjectCreate(ctx, iamProjectCreate, optional)

Create iam/project

Create project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iamProjectCreate** | [**IamProjectCreate**](IamProjectCreate.md)|  | 
 **optional** | ***IamProjectCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectCreateOpts struct


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


## IamProjectCredentialStoreCreate

> ProjectCredential IamProjectCredentialStoreCreate(ctx, projectId, projectCredential)

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


## IamProjectCredentialStoreDelete

> Project IamProjectCredentialStoreDelete(ctx, projectId, credentialStoreId)

Delete iam/project.credentialStore

Delete iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**credentialStoreId** | **string**| credentialStoreId | 

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


## IamProjectCredentialStoreGet

> ProjectCredential IamProjectCredentialStoreGet(ctx, projectId, credentialStoreId)

Get iam/project.credentialStore

Get iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**credentialStoreId** | **string**| credentialStoreId | 

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


## IamProjectCredentialStoreList

> []ProjectCredential IamProjectCredentialStoreList(ctx, projectId)

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


## IamProjectCredentialStorePatch

> ProjectCredential IamProjectCredentialStorePatch(ctx, projectId, credentialStoreId, iamProjectCredentialStorePatch)

Update iam/project.credentialStore

Update iam/project.credentialStore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**credentialStoreId** | **string**| credentialStoreId | 
**iamProjectCredentialStorePatch** | [**IamProjectCredentialStorePatch**](IamProjectCredentialStorePatch.md)|  | 

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


## IamProjectDelete

> IamProjectDelete(ctx, projectId)

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


## IamProjectEventGet

> Event IamProjectEventGet(ctx, projectId, eventId)

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


## IamProjectEventList

> []Event IamProjectEventList(ctx, projectId, optional)

List iam/project.event

List iam/project.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***IamProjectEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectEventListOpts struct


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


## IamProjectGet

> Project IamProjectGet(ctx, projectId)

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


## IamProjectInvitationAccept

> Invitation IamProjectInvitationAccept(ctx, projectId, invitationId, iamProjectInvitationAccept)

Accept iam/project.invitation

action accept

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**invitationId** | **string**| invitationId | 
**iamProjectInvitationAccept** | [**IamProjectInvitationAccept**](IamProjectInvitationAccept.md)|  | 

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


## IamProjectInvitationDelete

> IamProjectInvitationDelete(ctx, projectId, invitationId)

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


## IamProjectInvitationGet

> Invitation IamProjectInvitationGet(ctx, projectId, invitationId)

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


## IamProjectInvitationList

> []Invitation IamProjectInvitationList(ctx, projectId, optional)

List iam/project.invitation

List iam/project.invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***IamProjectInvitationListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectInvitationListOpts struct


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


## IamProjectInvoiceList

> []Invoice IamProjectInvoiceList(ctx, projectId)

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


## IamProjectList

> []Project IamProjectList(ctx, optional)

List iam/project

List project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IamProjectListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectListOpts struct


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


## IamProjectOwnershipCreate

> Project IamProjectOwnershipCreate(ctx, projectId, iamProjectOwnershipCreate)

Create iam/project.ownership

Create iam/project.ownership

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**iamProjectOwnershipCreate** | [**IamProjectOwnershipCreate**](IamProjectOwnershipCreate.md)|  | 

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


## IamProjectOwnershipDelete

> IamProjectOwnershipDelete(ctx, projectId, ownershipId)

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


## IamProjectOwnershipGet

> Ownership IamProjectOwnershipGet(ctx, projectId, ownershipId)

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


## IamProjectOwnershipList

> []Ownership IamProjectOwnershipList(ctx, projectId)

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


## IamProjectPaymentList

> []Payment IamProjectPaymentList(ctx, projectId)

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


## IamProjectProformaList

> []Proforma IamProjectProformaList(ctx, projectId)

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


## IamProjectQuotaGet

> Quota IamProjectQuotaGet(ctx, projectId, quotaId)

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


## IamProjectQuotaLimitPatch

> QuotaLimit IamProjectQuotaLimitPatch(ctx, projectId, quotaId, iamProjectQuotaLimitPatch)

Update iam/project.limit

Update iam/project.limit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**quotaId** | **string**| quotaId | 
**iamProjectQuotaLimitPatch** | [**IamProjectQuotaLimitPatch**](IamProjectQuotaLimitPatch.md)|  | 

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


## IamProjectQuotaList

> []Quota IamProjectQuotaList(ctx, projectId)

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


## IamProjectServiceGet

> ResourceService IamProjectServiceGet(ctx, projectId, serviceId)

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


## IamProjectServiceList

> []ResourceService IamProjectServiceList(ctx, projectId)

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


## IamProjectTagCreate

> Tag IamProjectTagCreate(ctx, projectId, tag)

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


## IamProjectTagDelete

> IamProjectTagDelete(ctx, projectId, tagId)

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


## IamProjectTagGet

> Tag IamProjectTagGet(ctx, projectId, tagId)

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


## IamProjectTagList

> []Tag IamProjectTagList(ctx, projectId)

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


## IamProjectTagPut

> []Tag IamProjectTagPut(ctx, projectId, tag)

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


## IamProjectThresholdCreate

> ProjectThreshold IamProjectThresholdCreate(ctx, projectId, iamProjectThresholdCreate)

Create iam/project.threshold

Create iam/project.threshold

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**iamProjectThresholdCreate** | [**IamProjectThresholdCreate**](IamProjectThresholdCreate.md)|  | 

### Return type

[**ProjectThreshold**](project.threshold.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectThresholdDelete

> IamProjectThresholdDelete(ctx, projectId, thresholdId)

Delete iam/project.threshold

Delete iam/project.threshold

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**thresholdId** | **string**| thresholdId | 

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


## IamProjectThresholdGet

> ProjectThreshold IamProjectThresholdGet(ctx, projectId, thresholdId)

Get iam/project.threshold

Get iam/project.threshold

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**thresholdId** | **string**| thresholdId | 

### Return type

[**ProjectThreshold**](project.threshold.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectThresholdList

> []ProjectThreshold IamProjectThresholdList(ctx, projectId)

List iam/project.threshold

List iam/project.threshold

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

[**[]ProjectThreshold**](project.threshold.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamProjectTransfer

> Project IamProjectTransfer(ctx, projectId, iamProjectTransfer, optional)

Transfer iam/project

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**iamProjectTransfer** | [**IamProjectTransfer**](IamProjectTransfer.md)|  | 
 **optional** | ***IamProjectTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectTransferOpts struct


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


## IamProjectUpdate

> Project IamProjectUpdate(ctx, projectId, iamProjectUpdate)

Update iam/project

Returns modified project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**iamProjectUpdate** | [**IamProjectUpdate**](IamProjectUpdate.md)|  | 

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

