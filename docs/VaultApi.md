# \VaultApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VaultActionResize**](VaultApi.md#VaultActionResize) | **Post** /vault/{vaultId}/actions/resize | /actions/resize
[**VaultActionSnapshot**](VaultApi.md#VaultActionSnapshot) | **Post** /vault/{vaultId}/actions/snapshot | /actions/snapshot
[**VaultActionStart**](VaultApi.md#VaultActionStart) | **Post** /vault/{vaultId}/actions/start | /actions/start
[**VaultActionStop**](VaultApi.md#VaultActionStop) | **Post** /vault/{vaultId}/actions/stop | /actions/stop
[**VaultCreate**](VaultApi.md#VaultCreate) | **Post** /vault | Create
[**VaultDelete**](VaultApi.md#VaultDelete) | **Delete** /vault/{vaultId} | Delete
[**VaultDeleteAccessrightsIdentity**](VaultApi.md#VaultDeleteAccessrightsIdentity) | **Delete** /vault/{vaultId}/accessrights/{identity} | /accessrights/:identity
[**VaultDeleteCredentialcertificateId**](VaultApi.md#VaultDeleteCredentialcertificateId) | **Delete** /vault/{vaultId}/credential/certificate/{id} | /credential/certificate/:id
[**VaultDeleteCredentialpasswordId**](VaultApi.md#VaultDeleteCredentialpasswordId) | **Delete** /vault/{vaultId}/credential/password/{id} | /credential/password/:id
[**VaultDeleteTagKey**](VaultApi.md#VaultDeleteTagKey) | **Delete** /vault/{vaultId}/tag/{key} | /tag/:key
[**VaultGetCredentialcertificateId**](VaultApi.md#VaultGetCredentialcertificateId) | **Get** /vault/{vaultId}/credential/certificate/{id} | /credential/certificate/:id
[**VaultGetCredentialpasswordId**](VaultApi.md#VaultGetCredentialpasswordId) | **Get** /vault/{vaultId}/credential/password/{id} | /credential/password/:id
[**VaultGetServicesServiceId**](VaultApi.md#VaultGetServicesServiceId) | **Get** /vault/{vaultId}/services/{serviceId} | /services/:serviceId
[**VaultGetTag**](VaultApi.md#VaultGetTag) | **Get** /vault/{vaultId}/tag | /tag
[**VaultList**](VaultApi.md#VaultList) | **Get** /vault | List
[**VaultListAccessrights**](VaultApi.md#VaultListAccessrights) | **Get** /vault/{vaultId}/accessrights | /accessrights
[**VaultListCredentialcertificate**](VaultApi.md#VaultListCredentialcertificate) | **Get** /vault/{vaultId}/credential/certificate | /credential/certificate
[**VaultListCredentialpassword**](VaultApi.md#VaultListCredentialpassword) | **Get** /vault/{vaultId}/credential/password | /credential/password
[**VaultListQueue**](VaultApi.md#VaultListQueue) | **Get** /vault/{vaultId}/queue | /queue
[**VaultListServices**](VaultApi.md#VaultListServices) | **Get** /vault/{vaultId}/services | /services
[**VaultPatchCredentialcertificateId**](VaultApi.md#VaultPatchCredentialcertificateId) | **Patch** /vault/{vaultId}/credential/certificate/{id} | /credential/certificate/:id
[**VaultPatchCredentialpasswordId**](VaultApi.md#VaultPatchCredentialpasswordId) | **Patch** /vault/{vaultId}/credential/password/{id} | /credential/password/:id
[**VaultPatchTag**](VaultApi.md#VaultPatchTag) | **Patch** /vault/{vaultId}/tag | /tag
[**VaultPostAccessrights**](VaultApi.md#VaultPostAccessrights) | **Post** /vault/{vaultId}/accessrights | /accessrights
[**VaultPostCredentialcertificate**](VaultApi.md#VaultPostCredentialcertificate) | **Post** /vault/{vaultId}/credential/certificate | /credential/certificate
[**VaultPostCredentialpassword**](VaultApi.md#VaultPostCredentialpassword) | **Post** /vault/{vaultId}/credential/password | /credential/password
[**VaultShow**](VaultApi.md#VaultShow) | **Get** /vault/{vaultId} | Get
[**VaultUpdate**](VaultApi.md#VaultUpdate) | **Patch** /vault/{vaultId} | Update


# **VaultActionResize**
> Vault VaultActionResize(ctx, vaultId, vaultActionResize)
/actions/resize

Action resize

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **vaultActionResize** | [**VaultActionResize**](VaultActionResize.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultActionSnapshot**
> Vault VaultActionSnapshot(ctx, vaultId, vaultActionSnapshot)
/actions/snapshot

Action snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **vaultActionSnapshot** | [**VaultActionSnapshot**](VaultActionSnapshot.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultActionStart**
> Vault VaultActionStart(ctx, vaultId)
/actions/start

Action start

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultActionStop**
> Vault VaultActionStop(ctx, vaultId)
/actions/stop

Action stop

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultCreate**
> Vault VaultCreate(ctx, vaultCreate)
Create

Create vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultCreate** | [**VaultCreate**](VaultCreate.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDelete**
> VaultDelete(ctx, vaultId, vaultDelete)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **vaultDelete** | [**VaultDelete**](VaultDelete.md)|  | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDeleteAccessrightsIdentity**
> Vault VaultDeleteAccessrightsIdentity(ctx, vaultId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **identity** | **string**| identity | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDeleteCredentialcertificateId**
> Vault VaultDeleteCredentialcertificateId(ctx, vaultId, id)
/credential/certificate/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDeleteCredentialpasswordId**
> Vault VaultDeleteCredentialpasswordId(ctx, vaultId, id)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDeleteTagKey**
> map[string]interface{} VaultDeleteTagKey(ctx, vaultId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetCredentialcertificateId**
> CredentialCertificate VaultGetCredentialcertificateId(ctx, vaultId, id)
/credential/certificate/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetCredentialpasswordId**
> CredentialPassword VaultGetCredentialpasswordId(ctx, vaultId, id)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetServicesServiceId**
> VaultServices VaultGetServicesServiceId(ctx, vaultId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **serviceId** | **string**| serviceId | 

### Return type

[**VaultServices**](vault.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetTag**
> map[string]interface{} VaultGetTag(ctx, vaultId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultList**
> []Vault VaultList(ctx, optional)
List

List vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VaultListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VaultListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultListAccessrights**
> []string VaultListAccessrights(ctx, vaultId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultListCredentialcertificate**
> []CredentialCertificate VaultListCredentialcertificate(ctx, vaultId)
/credential/certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**[]CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultListCredentialpassword**
> []CredentialPassword VaultListCredentialpassword(ctx, vaultId)
/credential/password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**[]CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultListQueue**
> []Event VaultListQueue(ctx, vaultId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultListServices**
> []VaultServices VaultListServices(ctx, vaultId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**[]VaultServices**](vault.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultPatchCredentialcertificateId**
> CredentialCertificate VaultPatchCredentialcertificateId(ctx, vaultId, id, vaultPatchCredentialcertificateId)
/credential/certificate/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 
  **vaultPatchCredentialcertificateId** | [**VaultPatchCredentialcertificateId**](VaultPatchCredentialcertificateId.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultPatchCredentialpasswordId**
> CredentialPassword VaultPatchCredentialpasswordId(ctx, vaultId, id, vaultPatchCredentialpasswordId)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 
  **vaultPatchCredentialpasswordId** | [**VaultPatchCredentialpasswordId**](VaultPatchCredentialpasswordId.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultPatchTag**
> map[string]interface{} VaultPatchTag(ctx, vaultId, requestBody)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultPostAccessrights**
> Vault VaultPostAccessrights(ctx, vaultId, vaultPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **vaultPostAccessrights** | [**VaultPostAccessrights**](VaultPostAccessrights.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultPostCredentialcertificate**
> CredentialCertificate VaultPostCredentialcertificate(ctx, vaultId, vaultPostCredentialcertificate)
/credential/certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **vaultPostCredentialcertificate** | [**VaultPostCredentialcertificate**](VaultPostCredentialcertificate.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultPostCredentialpassword**
> CredentialPassword VaultPostCredentialpassword(ctx, vaultId, vaultPostCredentialpassword)
/credential/password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **vaultPostCredentialpassword** | [**VaultPostCredentialpassword**](VaultPostCredentialpassword.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultShow**
> Vault VaultShow(ctx, vaultId)
Get

Returns a single vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultUpdate**
> Vault VaultUpdate(ctx, vaultId, vaultUpdate)
Update

Returns modified vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **vaultUpdate** | [**VaultUpdate**](VaultUpdate.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

