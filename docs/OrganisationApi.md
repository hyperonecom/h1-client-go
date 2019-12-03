# \OrganisationApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganisationActionPaymentAssign**](OrganisationApi.md#OrganisationActionPaymentAssign) | **Post** /organisation/{organisationId}/actions/payment_assign | /actions/payment_assign
[**OrganisationActionTransferAccept**](OrganisationApi.md#OrganisationActionTransferAccept) | **Post** /organisation/{organisationId}/actions/transfer_accept | /actions/transfer_accept
[**OrganisationCreate**](OrganisationApi.md#OrganisationCreate) | **Post** /organisation | Create
[**OrganisationDeleteAccessrightsId**](OrganisationApi.md#OrganisationDeleteAccessrightsId) | **Delete** /organisation/{organisationId}/accessrights/{id} | /accessrights/:id
[**OrganisationDeleteTagKey**](OrganisationApi.md#OrganisationDeleteTagKey) | **Delete** /organisation/{organisationId}/tag/{key} | /tag/:key
[**OrganisationGetTag**](OrganisationApi.md#OrganisationGetTag) | **Get** /organisation/{organisationId}/tag | /tag
[**OrganisationList**](OrganisationApi.md#OrganisationList) | **Get** /organisation | List
[**OrganisationListAccessrights**](OrganisationApi.md#OrganisationListAccessrights) | **Get** /organisation/{organisationId}/accessrights | /accessrights
[**OrganisationListBilling**](OrganisationApi.md#OrganisationListBilling) | **Get** /organisation/{organisationId}/billing | /billing
[**OrganisationListQueue**](OrganisationApi.md#OrganisationListQueue) | **Get** /organisation/{organisationId}/queue | /queue
[**OrganisationPatchTag**](OrganisationApi.md#OrganisationPatchTag) | **Patch** /organisation/{organisationId}/tag | /tag
[**OrganisationPostAccessrights**](OrganisationApi.md#OrganisationPostAccessrights) | **Post** /organisation/{organisationId}/accessrights | /accessrights
[**OrganisationPutTag**](OrganisationApi.md#OrganisationPutTag) | **Put** /organisation/{organisationId}/tag | /tag
[**OrganisationShow**](OrganisationApi.md#OrganisationShow) | **Get** /organisation/{organisationId} | Get
[**OrganisationUpdate**](OrganisationApi.md#OrganisationUpdate) | **Patch** /organisation/{organisationId} | Update



## OrganisationActionPaymentAssign

> Organisation OrganisationActionPaymentAssign(ctx, organisationId, organisationActionPaymentAssign)

/actions/payment_assign

Action payment_assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**organisationActionPaymentAssign** | [**OrganisationActionPaymentAssign**](OrganisationActionPaymentAssign.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationActionTransferAccept

> Organisation OrganisationActionTransferAccept(ctx, organisationId, organisationActionTransferAccept)

/actions/transfer_accept

Action transfer_accept

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**organisationActionTransferAccept** | [**OrganisationActionTransferAccept**](OrganisationActionTransferAccept.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationCreate

> Organisation OrganisationCreate(ctx, organisationCreate)

Create

Create organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationCreate** | [**OrganisationCreate**](OrganisationCreate.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationDeleteAccessrightsId

> OrganisationDeleteAccessrightsId(ctx, organisationId, id)

/accessrights/:id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationDeleteTagKey

> map[string]string OrganisationDeleteTagKey(ctx, organisationId, key)

/tag/:key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationGetTag

> map[string]string OrganisationGetTag(ctx, organisationId)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationList

> []Organisation OrganisationList(ctx, optional)

List

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
 **active** | **optional.Bool**| Filter by active | 
 **accessRightsId** | **optional.String**| Filter by accessRights.id | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationListAccessrights

> []AccessrightsUserRole OrganisationListAccessrights(ctx, organisationId)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 

### Return type

[**[]AccessrightsUserRole**](accessrightsUserRole.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationListBilling

> []Billing OrganisationListBilling(ctx, organisationId, optional)

/billing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
 **optional** | ***OrganisationListBillingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationListBillingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Time**| start | 
 **end** | **optional.Time**| end | 
 **resourceType** | **optional.String**| resource.type | 

### Return type

[**[]Billing**](billing.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationListQueue

> []Event OrganisationListQueue(ctx, organisationId, optional)

/queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
 **optional** | ***OrganisationListQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganisationListQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Float32**| $limit | 
 **skip** | **optional.Float32**| $skip | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationPatchTag

> map[string]string OrganisationPatchTag(ctx, organisationId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationPostAccessrights

> OrganisationPostAccessrights(ctx, organisationId, organisationPostAccessrights)

/accessrights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**organisationPostAccessrights** | [**OrganisationPostAccessrights**](OrganisationPostAccessrights.md)|  | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationPutTag

> map[string]string OrganisationPutTag(ctx, organisationId, requestBody)

/tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationShow

> Organisation OrganisationShow(ctx, organisationId)

Get

Returns a single organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationUpdate

> Organisation OrganisationUpdate(ctx, organisationId, organisationUpdate)

Update

Returns modified organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| ID of organisation | 
**organisationUpdate** | [**OrganisationUpdate**](OrganisationUpdate.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

