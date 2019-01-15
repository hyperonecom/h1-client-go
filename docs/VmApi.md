# \VmApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionVmFlavour**](VmApi.md#ActionVmFlavour) | **Post** /vm/{vmId}/actions/flavour | /actions/flavour
[**ActionVmImage**](VmApi.md#ActionVmImage) | **Post** /vm/{vmId}/actions/image | /actions/image
[**ActionVmPasswordReset**](VmApi.md#ActionVmPasswordReset) | **Post** /vm/{vmId}/actions/password_reset | /actions/password_reset
[**ActionVmRename**](VmApi.md#ActionVmRename) | **Post** /vm/{vmId}/actions/rename | /actions/rename
[**ActionVmRestart**](VmApi.md#ActionVmRestart) | **Post** /vm/{vmId}/actions/restart | /actions/restart
[**ActionVmStart**](VmApi.md#ActionVmStart) | **Post** /vm/{vmId}/actions/start | /actions/start
[**ActionVmStop**](VmApi.md#ActionVmStop) | **Post** /vm/{vmId}/actions/stop | /actions/stop
[**ActionVmTurnoff**](VmApi.md#ActionVmTurnoff) | **Post** /vm/{vmId}/actions/turnoff | /actions/turnoff
[**CreateVm**](VmApi.md#CreateVm) | **Post** /vm | Create
[**DeleteVm**](VmApi.md#DeleteVm) | **Delete** /vm/{vmId} | Delete
[**ListVm**](VmApi.md#ListVm) | **Get** /vm | List
[**OperationVmDeleteaccessrightsIdentity**](VmApi.md#OperationVmDeleteaccessrightsIdentity) | **Delete** /vm/{vmId}/accessrights/{identity} | /accessrights/:identity
[**OperationVmDeletehddDiskId**](VmApi.md#OperationVmDeletehddDiskId) | **Delete** /vm/{vmId}/hdd/{diskId} | /hdd/:diskId
[**OperationVmDeletenetadp**](VmApi.md#OperationVmDeletenetadp) | **Delete** /vm/{vmId}/netadp | /netadp
[**OperationVmDeletetagKey**](VmApi.md#OperationVmDeletetagKey) | **Delete** /vm/{vmId}/tag/{key} | /tag/:key
[**OperationVmGetservicesServiceId**](VmApi.md#OperationVmGetservicesServiceId) | **Get** /vm/{vmId}/services/{serviceId} | /services/:serviceId
[**OperationVmGettag**](VmApi.md#OperationVmGettag) | **Get** /vm/{vmId}/tag/ | /tag/
[**OperationVmListaccessrights**](VmApi.md#OperationVmListaccessrights) | **Get** /vm/{vmId}/accessrights/ | /accessrights/
[**OperationVmListhdd**](VmApi.md#OperationVmListhdd) | **Get** /vm/{vmId}/hdd | /hdd
[**OperationVmListnetadp**](VmApi.md#OperationVmListnetadp) | **Get** /vm/{vmId}/netadp | /netadp
[**OperationVmListqueue**](VmApi.md#OperationVmListqueue) | **Get** /vm/{vmId}/queue/ | /queue/
[**OperationVmListservices**](VmApi.md#OperationVmListservices) | **Get** /vm/{vmId}/services/ | /services/
[**OperationVmPatchtag**](VmApi.md#OperationVmPatchtag) | **Patch** /vm/{vmId}/tag/ | /tag/
[**OperationVmPostaccessrights**](VmApi.md#OperationVmPostaccessrights) | **Post** /vm/{vmId}/accessrights/ | /accessrights/
[**OperationVmPosthdd**](VmApi.md#OperationVmPosthdd) | **Post** /vm/{vmId}/hdd | /hdd
[**OperationVmPostnetadp**](VmApi.md#OperationVmPostnetadp) | **Post** /vm/{vmId}/netadp | /netadp
[**ShowVm**](VmApi.md#ShowVm) | **Get** /vm/{vmId} | Get
[**UpdateVm**](VmApi.md#UpdateVm) | **Patch** /vm/{vmId} | Update


# **ActionVmFlavour**
> Vm ActionVmFlavour(ctx, vmId)
/actions/flavour

Action flavour

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmImage**
> Vm ActionVmImage(ctx, vmId, optional)
/actions/image

Action image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
 **optional** | ***ActionVmImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionVmImageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject28** | [**optional.Interface of InlineObject28**](InlineObject28.md)|  | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmPasswordReset**
> Vm ActionVmPasswordReset(ctx, vmId, optional)
/actions/password_reset

Action password_reset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
 **optional** | ***ActionVmPasswordResetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionVmPasswordResetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject27** | [**optional.Interface of InlineObject27**](InlineObject27.md)|  | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmRename**
> Vm ActionVmRename(ctx, vmId)
/actions/rename

Action rename

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmRestart**
> Vm ActionVmRestart(ctx, vmId)
/actions/restart

Action restart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmStart**
> Vm ActionVmStart(ctx, vmId)
/actions/start

Action start

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmStop**
> Vm ActionVmStop(ctx, vmId)
/actions/stop

Action stop

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionVmTurnoff**
> Vm ActionVmTurnoff(ctx, vmId)
/actions/turnoff

Action turnoff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVm**
> Vm CreateVm(ctx, vmCreate)
Create

Create vm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmCreate** | [**VmCreate**](VmCreate.md)|  | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVm**
> DeleteVm(ctx, vmId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVm**
> []Vm ListVm(ctx, optional)
List

List vm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListVmOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmDeleteaccessrightsIdentity**
> Vm OperationVmDeleteaccessrightsIdentity(ctx, vmId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
  **identity** | **string**| identity | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmDeletehddDiskId**
> Vm OperationVmDeletehddDiskId(ctx, vmId, diskId)
/hdd/:diskId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
  **diskId** | **string**| diskId | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmDeletenetadp**
> Vm OperationVmDeletenetadp(ctx, vmId)
/netadp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmDeletetagKey**
> map[string]string OperationVmDeletetagKey(ctx, vmId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmGetservicesServiceId**
> VmServices OperationVmGetservicesServiceId(ctx, vmId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
  **serviceId** | **string**| serviceId | 

### Return type

[**VmServices**](vm.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmGettag**
> map[string]string OperationVmGettag(ctx, vmId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmListaccessrights**
> []string OperationVmListaccessrights(ctx, vmId)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmListhdd**
> []Hdd OperationVmListhdd(ctx, vmId)
/hdd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**[]Hdd**](hdd.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmListnetadp**
> []Netadp OperationVmListnetadp(ctx, vmId)
/netadp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**[]Netadp**](netadp.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmListqueue**
> []Event OperationVmListqueue(ctx, vmId)
/queue/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmListservices**
> []VmServices OperationVmListservices(ctx, vmId)
/services/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**[]VmServices**](vm.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmPatchtag**
> map[string]string OperationVmPatchtag(ctx, vmId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmPostaccessrights**
> string OperationVmPostaccessrights(ctx, vmId, optional)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
 **optional** | ***OperationVmPostaccessrightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVmPostaccessrightsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject31** | [**optional.Interface of InlineObject31**](InlineObject31.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmPosthdd**
> Vm OperationVmPosthdd(ctx, vmId, optional)
/hdd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
 **optional** | ***OperationVmPosthddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVmPosthddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject29** | [**optional.Interface of InlineObject29**](InlineObject29.md)|  | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationVmPostnetadp**
> Vm OperationVmPostnetadp(ctx, vmId, optional)
/netadp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
 **optional** | ***OperationVmPostnetadpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationVmPostnetadpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject30** | [**optional.Interface of InlineObject30**](InlineObject30.md)|  | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowVm**
> Vm ShowVm(ctx, vmId)
Get

Returns a single vm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVm**
> Vm UpdateVm(ctx, vmId, optional)
Update

Returns modified vm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| ID of vm | 
 **optional** | ***UpdateVmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateVmOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject26** | [**optional.Interface of InlineObject26**](InlineObject26.md)|  | 

### Return type

[**Vm**](vm.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

