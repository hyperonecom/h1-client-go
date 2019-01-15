# \VmhostApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionVmhostMoveDisk**](VmhostApi.md#ActionVmhostMoveDisk) | **Post** /vmhost/{vmhostId}/actions/moveDisk | /actions/moveDisk
[**ActionVmhostMoveISO**](VmhostApi.md#ActionVmhostMoveISO) | **Post** /vmhost/{vmhostId}/actions/moveISO | /actions/moveISO
[**ActionVmhostMoveImage**](VmhostApi.md#ActionVmhostMoveImage) | **Post** /vmhost/{vmhostId}/actions/moveImage | /actions/moveImage
[**ActionVmhostMoveVM**](VmhostApi.md#ActionVmhostMoveVM) | **Post** /vmhost/{vmhostId}/actions/moveVM | /actions/moveVM
[**DeleteVmhost**](VmhostApi.md#DeleteVmhost) | **Delete** /vmhost/{vmhostId} | Delete
[**ListVmhost**](VmhostApi.md#ListVmhost) | **Get** /vmhost | List
[**OperationVmhostDeleteaccessrightsIdentity**](VmhostApi.md#OperationVmhostDeleteaccessrightsIdentity) | **Delete** /vmhost/{vmhostId}/accessrights/{identity} | /accessrights/:identity
[**OperationVmhostDeletetagKey**](VmhostApi.md#OperationVmhostDeletetagKey) | **Delete** /vmhost/{vmhostId}/tag/{key} | /tag/:key
[**OperationVmhostGetservicesServiceId**](VmhostApi.md#OperationVmhostGetservicesServiceId) | **Get** /vmhost/{vmhostId}/services/{serviceId} | /services/:serviceId
[**OperationVmhostGettag**](VmhostApi.md#OperationVmhostGettag) | **Get** /vmhost/{vmhostId}/tag/ | /tag/
[**OperationVmhostListaccessrights**](VmhostApi.md#OperationVmhostListaccessrights) | **Get** /vmhost/{vmhostId}/accessrights/ | /accessrights/
[**OperationVmhostListqueue**](VmhostApi.md#OperationVmhostListqueue) | **Get** /vmhost/{vmhostId}/queue/ | /queue/
[**OperationVmhostListservices**](VmhostApi.md#OperationVmhostListservices) | **Get** /vmhost/{vmhostId}/services/ | /services/
[**OperationVmhostPatchtag**](VmhostApi.md#OperationVmhostPatchtag) | **Patch** /vmhost/{vmhostId}/tag/ | /tag/
[**OperationVmhostPostaccessrights**](VmhostApi.md#OperationVmhostPostaccessrights) | **Post** /vmhost/{vmhostId}/accessrights/ | /accessrights/
[**ShowVmhost**](VmhostApi.md#ShowVmhost) | **Get** /vmhost/{vmhostId} | Get
[**UpdateVmhost**](VmhostApi.md#UpdateVmhost) | **Patch** /vmhost/{vmhostId} | Update


# **ActionVmhostMoveDisk**
> Vmhost ActionVmhostMoveDisk(ctx, vmhostId, optional)
/actions/moveDisk

Action moveDisk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
 **optional** | ***ActionVmhostMoveDiskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionVmhostMoveDiskOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject39** | [**optional.Interface of InlineObject39**](InlineObject39.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmhostMoveISO**
> Vmhost ActionVmhostMoveISO(ctx, vmhostId, optional)
/actions/moveISO

Action moveISO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
 **optional** | ***ActionVmhostMoveISOOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionVmhostMoveISOOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject40** | [**optional.Interface of InlineObject40**](InlineObject40.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmhostMoveImage**
> Vmhost ActionVmhostMoveImage(ctx, vmhostId, optional)
/actions/moveImage

Action moveImage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
 **optional** | ***ActionVmhostMoveImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionVmhostMoveImageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject41** | [**optional.Interface of InlineObject41**](InlineObject41.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmhostMoveVM**
> Vmhost ActionVmhostMoveVM(ctx, vmhostId, optional)
/actions/moveVM

Action moveVM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
 **optional** | ***ActionVmhostMoveVMOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionVmhostMoveVMOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject38** | [**optional.Interface of InlineObject38**](InlineObject38.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVmhost**
> DeleteVmhost(ctx, vmhostId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVmhost**
> []Vmhost ListVmhost(ctx, optional)
List

List vmhost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVmhostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListVmhostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabledServices** | **optional.String**| Filter by enabledServices | 

### Return type

[**[]Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostDeleteaccessrightsIdentity**
> Vmhost OperationVmhostDeleteaccessrightsIdentity(ctx, vmhostId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **identity** | **string**| identity | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostDeletetagKey**
> map[string]string OperationVmhostDeletetagKey(ctx, vmhostId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostGetservicesServiceId**
> VmhostServices OperationVmhostGetservicesServiceId(ctx, vmhostId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **serviceId** | **string**| serviceId | 

### Return type

[**VmhostServices**](vmhost.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostGettag**
> map[string]string OperationVmhostGettag(ctx, vmhostId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostListaccessrights**
> []string OperationVmhostListaccessrights(ctx, vmhostId)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostListqueue**
> []Event OperationVmhostListqueue(ctx, vmhostId)
/queue/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostListservices**
> []VmhostServices OperationVmhostListservices(ctx, vmhostId)
/services/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

[**[]VmhostServices**](vmhost.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostPatchtag**
> map[string]string OperationVmhostPatchtag(ctx, vmhostId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmhostPostaccessrights**
> string OperationVmhostPostaccessrights(ctx, vmhostId, resourceAccessRight)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
  **resourceAccessRight** | [**ResourceAccessRight**](ResourceAccessRight.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowVmhost**
> Vmhost ShowVmhost(ctx, vmhostId)
Get

Returns a single vmhost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVmhost**
> Vmhost UpdateVmhost(ctx, vmhostId, optional)
Update

Returns modified vmhost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmhostId** | **string**| ID of vmhost | 
 **optional** | ***UpdateVmhostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateVmhostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject37** | [**optional.Interface of InlineObject37**](InlineObject37.md)|  | 

### Return type

[**Vmhost**](vmhost.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

