# \OrganisationApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionOrganisationTransferAccept**](OrganisationApi.md#ActionOrganisationTransferAccept) | **Post** /organisation/{organisationId}/actions/transfer_accept | /actions/transfer_accept
[**CreateOrganisation**](OrganisationApi.md#CreateOrganisation) | **Post** /organisation | Create
[**ListOrganisation**](OrganisationApi.md#ListOrganisation) | **Get** /organisation | List
[**OperationOrganisationDeleteaccessrightsIdentity**](OrganisationApi.md#OperationOrganisationDeleteaccessrightsIdentity) | **Delete** /organisation/{organisationId}/accessrights/{identity} | /accessrights/:identity
[**OperationOrganisationDeletetagKey**](OrganisationApi.md#OperationOrganisationDeletetagKey) | **Delete** /organisation/{organisationId}/tag/{key} | /tag/:key
[**OperationOrganisationGettag**](OrganisationApi.md#OperationOrganisationGettag) | **Get** /organisation/{organisationId}/tag/ | /tag/
[**OperationOrganisationListaccessrights**](OrganisationApi.md#OperationOrganisationListaccessrights) | **Get** /organisation/{organisationId}/accessrights/ | /accessrights/
[**OperationOrganisationListqueue**](OrganisationApi.md#OperationOrganisationListqueue) | **Get** /organisation/{organisationId}/queue/ | /queue/
[**OperationOrganisationPatchtag**](OrganisationApi.md#OperationOrganisationPatchtag) | **Patch** /organisation/{organisationId}/tag/ | /tag/
[**OperationOrganisationPostaccessrights**](OrganisationApi.md#OperationOrganisationPostaccessrights) | **Post** /organisation/{organisationId}/accessrights/ | /accessrights/
[**ShowOrganisation**](OrganisationApi.md#ShowOrganisation) | **Get** /organisation/{organisationId} | Get
[**UpdateOrganisation**](OrganisationApi.md#UpdateOrganisation) | **Patch** /organisation/{organisationId} | Update


# **ActionOrganisationTransferAccept**
> Organisation ActionOrganisationTransferAccept(ctx, organisationId, optional)
/actions/transfer_accept

Action transfer_accept

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 
 **optional** | ***ActionOrganisationTransferAcceptOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionOrganisationTransferAcceptOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject5** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrganisation**
> Organisation CreateOrganisation(ctx, optional)
Create

Create organisation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateOrganisationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOrganisationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject3** | [**optional.Interface of InlineObject3**](InlineObject3.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrganisation**
> []Organisation ListOrganisation(ctx, optional)
List

List organisation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOrganisationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOrganisationOpts struct

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

# **OperationOrganisationDeleteaccessrightsIdentity**
> Organisation OperationOrganisationDeleteaccessrightsIdentity(ctx, organisationId, identity)
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

# **OperationOrganisationDeletetagKey**
> map[string]string OperationOrganisationDeletetagKey(ctx, organisationId, key)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationOrganisationGettag**
> map[string]string OperationOrganisationGettag(ctx, organisationId)
/tag/

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationOrganisationListaccessrights**
> []OrganisationAccessRights OperationOrganisationListaccessrights(ctx, organisationId)
/accessrights/

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

# **OperationOrganisationListqueue**
> []Event OperationOrganisationListqueue(ctx, organisationId)
/queue/

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

# **OperationOrganisationPatchtag**
> map[string]string OperationOrganisationPatchtag(ctx, organisationId, requestBody)
/tag/

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationOrganisationPostaccessrights**
> OrganisationAccessRights OperationOrganisationPostaccessrights(ctx, organisationId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 
 **optional** | ***OperationOrganisationPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationOrganisationPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject6** | [**optional.Interface of InlineObject6**](InlineObject6.md)|  | 

### Return type

[**OrganisationAccessRights**](organisation.accessRights.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowOrganisation**
> Organisation ShowOrganisation(ctx, organisationId)
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

# **UpdateOrganisation**
> Organisation UpdateOrganisation(ctx, organisationId, optional)
Update

Returns modified organisation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organisationId** | **string**| ID of organisation | 
 **optional** | ***UpdateOrganisationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateOrganisationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject4** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

### Return type

[**Organisation**](organisation.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

