# \OrganisationApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganisationActionTransferAccept**](OrganisationApi.md#OrganisationActionTransferAccept) | **Post** /organisation/{organisationId}/actions/transfer_accept | /actions/transfer_accept
[**OrganisationCreate**](OrganisationApi.md#OrganisationCreate) | **Post** /organisation | Create
[**OrganisationDeleteAccessrightsIdentity**](OrganisationApi.md#OrganisationDeleteAccessrightsIdentity) | **Delete** /organisation/{organisationId}/accessrights/{identity} | /accessrights/:identity
[**OrganisationDeleteTagKey**](OrganisationApi.md#OrganisationDeleteTagKey) | **Delete** /organisation/{organisationId}/tag/{key} | /tag/:key
[**OrganisationGetTag**](OrganisationApi.md#OrganisationGetTag) | **Get** /organisation/{organisationId}/tag | /tag
[**OrganisationList**](OrganisationApi.md#OrganisationList) | **Get** /organisation | List
[**OrganisationListAccessrights**](OrganisationApi.md#OrganisationListAccessrights) | **Get** /organisation/{organisationId}/accessrights | /accessrights
[**OrganisationListQueue**](OrganisationApi.md#OrganisationListQueue) | **Get** /organisation/{organisationId}/queue | /queue
[**OrganisationPatchTag**](OrganisationApi.md#OrganisationPatchTag) | **Patch** /organisation/{organisationId}/tag | /tag
[**OrganisationPostAccessrights**](OrganisationApi.md#OrganisationPostAccessrights) | **Post** /organisation/{organisationId}/accessrights | /accessrights
[**OrganisationShow**](OrganisationApi.md#OrganisationShow) | **Get** /organisation/{organisationId} | Get
[**OrganisationUpdate**](OrganisationApi.md#OrganisationUpdate) | **Patch** /organisation/{organisationId} | Update


# **OrganisationActionTransferAccept**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationCreate**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationDeleteAccessrightsIdentity**
> Organisation OrganisationDeleteAccessrightsIdentity(ctx, organisationId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 
  **identity** | **string**| identity | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationDeleteTagKey**
> map[string]interface{} OrganisationDeleteTagKey(ctx, organisationId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationGetTag**
> map[string]interface{} OrganisationGetTag(ctx, organisationId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationList**
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
 **limit** | **optional.String**| Filter by $limit | 
 **active** | **optional.String**| Filter by active | 

### Return type

[**[]Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationListAccessrights**
> []OrganisationAccessRights OrganisationListAccessrights(ctx, organisationId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 

### Return type

[**[]OrganisationAccessRights**](organisation.accessRights.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationListQueue**
> []Event OrganisationListQueue(ctx, organisationId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationPatchTag**
> map[string]interface{} OrganisationPatchTag(ctx, organisationId, body)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 
  **body** | **map[string]interface{}**|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationPostAccessrights**
> OrganisationAccessRights OrganisationPostAccessrights(ctx, organisationId, organisationPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 
  **organisationPostAccessrights** | [**OrganisationPostAccessrights**](OrganisationPostAccessrights.md)|  | 

### Return type

[**OrganisationAccessRights**](organisation.accessRights.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationShow**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganisationUpdate**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

