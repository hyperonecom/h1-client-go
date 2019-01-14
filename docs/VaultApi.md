# \VaultApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionVaultResize**](VaultApi.md#ActionVaultResize) | **Post** /vault/{vaultId}/actions/resize | /actions/resize
[**ActionVaultSnapshot**](VaultApi.md#ActionVaultSnapshot) | **Post** /vault/{vaultId}/actions/snapshot | /actions/snapshot
[**ActionVaultStart**](VaultApi.md#ActionVaultStart) | **Post** /vault/{vaultId}/actions/start | /actions/start
[**ActionVaultStop**](VaultApi.md#ActionVaultStop) | **Post** /vault/{vaultId}/actions/stop | /actions/stop
[**CreateVault**](VaultApi.md#CreateVault) | **Post** /vault | Create
[**DeleteVault**](VaultApi.md#DeleteVault) | **Delete** /vault/{vaultId} | Delete
[**ListVault**](VaultApi.md#ListVault) | **Get** /vault | List
[**OperationVaultDeleteaccessrightsIdentity**](VaultApi.md#OperationVaultDeleteaccessrightsIdentity) | **Delete** /vault/{vaultId}/accessrights/{identity} | /accessrights/:identity
[**OperationVaultDeletecredentialcertificateId**](VaultApi.md#OperationVaultDeletecredentialcertificateId) | **Delete** /vault/{vaultId}/credential/certificate/{id} | /credential/certificate/:id
[**OperationVaultDeletecredentialpasswordId**](VaultApi.md#OperationVaultDeletecredentialpasswordId) | **Delete** /vault/{vaultId}/credential/password/{id} | /credential/password/:id
[**OperationVaultDeletetagKey**](VaultApi.md#OperationVaultDeletetagKey) | **Delete** /vault/{vaultId}/tag/{key} | /tag/:key
[**OperationVaultGetcredentialcertificateId**](VaultApi.md#OperationVaultGetcredentialcertificateId) | **Get** /vault/{vaultId}/credential/certificate/{id} | /credential/certificate/:id
[**OperationVaultGetcredentialpasswordId**](VaultApi.md#OperationVaultGetcredentialpasswordId) | **Get** /vault/{vaultId}/credential/password/{id} | /credential/password/:id
[**OperationVaultGetservicesServiceId**](VaultApi.md#OperationVaultGetservicesServiceId) | **Get** /vault/{vaultId}/services/{serviceId} | /services/:serviceId
[**OperationVaultGettag**](VaultApi.md#OperationVaultGettag) | **Get** /vault/{vaultId}/tag/ | /tag/
[**OperationVaultListaccessrights**](VaultApi.md#OperationVaultListaccessrights) | **Get** /vault/{vaultId}/accessrights/ | /accessrights/
[**OperationVaultListcredentialcertificate**](VaultApi.md#OperationVaultListcredentialcertificate) | **Get** /vault/{vaultId}/credential/certificate | /credential/certificate
[**OperationVaultListcredentialpassword**](VaultApi.md#OperationVaultListcredentialpassword) | **Get** /vault/{vaultId}/credential/password | /credential/password
[**OperationVaultListqueue**](VaultApi.md#OperationVaultListqueue) | **Get** /vault/{vaultId}/queue/ | /queue/
[**OperationVaultListservices**](VaultApi.md#OperationVaultListservices) | **Get** /vault/{vaultId}/services/ | /services/
[**OperationVaultPatchcredentialcertificateId**](VaultApi.md#OperationVaultPatchcredentialcertificateId) | **Patch** /vault/{vaultId}/credential/certificate/{id} | /credential/certificate/:id
[**OperationVaultPatchcredentialpasswordId**](VaultApi.md#OperationVaultPatchcredentialpasswordId) | **Patch** /vault/{vaultId}/credential/password/{id} | /credential/password/:id
[**OperationVaultPatchtag**](VaultApi.md#OperationVaultPatchtag) | **Patch** /vault/{vaultId}/tag/ | /tag/
[**OperationVaultPostaccessrights**](VaultApi.md#OperationVaultPostaccessrights) | **Post** /vault/{vaultId}/accessrights/ | /accessrights/
[**OperationVaultPostcredentialcertificate**](VaultApi.md#OperationVaultPostcredentialcertificate) | **Post** /vault/{vaultId}/credential/certificate | /credential/certificate
[**OperationVaultPostcredentialpassword**](VaultApi.md#OperationVaultPostcredentialpassword) | **Post** /vault/{vaultId}/credential/password | /credential/password
[**ShowVault**](VaultApi.md#ShowVault) | **Get** /vault/{vaultId} | Get
[**UpdateVault**](VaultApi.md#UpdateVault) | **Patch** /vault/{vaultId} | Update


# **ActionVaultResize**
> Vault ActionVaultResize(ctx, vaultId)
/actions/resize

Action resize

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

# **ActionVaultSnapshot**
> Vault ActionVaultSnapshot(ctx, vaultId)
/actions/snapshot

Action snapshot

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

# **ActionVaultStart**
> Vault ActionVaultStart(ctx, vaultId)
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

# **ActionVaultStop**
> Vault ActionVaultStop(ctx, vaultId)
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

# **CreateVault**
> Vault CreateVault(ctx, optional)
Create

Create vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateVaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateVaultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject47** | [**optional.Interface of InlineObject47**](InlineObject47.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVault**
> DeleteVault(ctx, vaultId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVault**
> []Vault ListVault(ctx, optional)
List

List vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListVaultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultDeleteaccessrightsIdentity**
> Vault OperationVaultDeleteaccessrightsIdentity(ctx, vaultId, identity)
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

# **OperationVaultDeletecredentialcertificateId**
> Vault OperationVaultDeletecredentialcertificateId(ctx, vaultId, id)
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

# **OperationVaultDeletecredentialpasswordId**
> Vault OperationVaultDeletecredentialpasswordId(ctx, vaultId, id)
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

# **OperationVaultDeletetagKey**
> map[string]string OperationVaultDeletetagKey(ctx, vaultId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultGetcredentialcertificateId**
> CredentialCertificate OperationVaultGetcredentialcertificateId(ctx, vaultId, id)
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

# **OperationVaultGetcredentialpasswordId**
> CredentialPassword OperationVaultGetcredentialpasswordId(ctx, vaultId, id)
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

# **OperationVaultGetservicesServiceId**
> VaultServices OperationVaultGetservicesServiceId(ctx, vaultId, serviceId)
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

# **OperationVaultGettag**
> map[string]string OperationVaultGettag(ctx, vaultId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultListaccessrights**
> []string OperationVaultListaccessrights(ctx, vaultId)
/accessrights/

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

# **OperationVaultListcredentialcertificate**
> []CredentialCertificate OperationVaultListcredentialcertificate(ctx, vaultId)
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

# **OperationVaultListcredentialpassword**
> []CredentialPassword OperationVaultListcredentialpassword(ctx, vaultId)
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

# **OperationVaultListqueue**
> []Event OperationVaultListqueue(ctx, vaultId)
/queue/

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

# **OperationVaultListservices**
> []VaultServices OperationVaultListservices(ctx, vaultId)
/services/

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

# **OperationVaultPatchcredentialcertificateId**
> CredentialCertificate OperationVaultPatchcredentialcertificateId(ctx, vaultId, id, optional)
/credential/certificate/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 
 **optional** | ***OperationVaultPatchcredentialcertificateIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVaultPatchcredentialcertificateIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject52** | [**optional.Interface of InlineObject52**](InlineObject52.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultPatchcredentialpasswordId**
> CredentialPassword OperationVaultPatchcredentialpasswordId(ctx, vaultId, id, optional)
/credential/password/:id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **id** | **string**| id | 
 **optional** | ***OperationVaultPatchcredentialpasswordIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVaultPatchcredentialpasswordIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject50** | [**optional.Interface of InlineObject50**](InlineObject50.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultPatchtag**
> map[string]string OperationVaultPatchtag(ctx, vaultId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultPostaccessrights**
> string OperationVaultPostaccessrights(ctx, vaultId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
 **optional** | ***OperationVaultPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVaultPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject53** | [**optional.Interface of InlineObject53**](InlineObject53.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultPostcredentialcertificate**
> CredentialCertificate OperationVaultPostcredentialcertificate(ctx, vaultId, optional)
/credential/certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
 **optional** | ***OperationVaultPostcredentialcertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVaultPostcredentialcertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject51** | [**optional.Interface of InlineObject51**](InlineObject51.md)|  | 

### Return type

[**CredentialCertificate**](credential.certificate.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVaultPostcredentialpassword**
> CredentialPassword OperationVaultPostcredentialpassword(ctx, vaultId, optional)
/credential/password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
 **optional** | ***OperationVaultPostcredentialpasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVaultPostcredentialpasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject49** | [**optional.Interface of InlineObject49**](InlineObject49.md)|  | 

### Return type

[**CredentialPassword**](credential.password.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowVault**
> Vault ShowVault(ctx, vaultId)
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

# **UpdateVault**
> Vault UpdateVault(ctx, vaultId, optional)
Update

Returns modified vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **string**| ID of vault | 
 **optional** | ***UpdateVaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateVaultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject48** | [**optional.Interface of InlineObject48**](InlineObject48.md)|  | 

### Return type

[**Vault**](vault.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

