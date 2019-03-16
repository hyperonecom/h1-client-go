# \FirewallApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FirewallActionAttach**](FirewallApi.md#FirewallActionAttach) | **Post** /firewall/{firewallId}/actions/attach | /actions/attach
[**FirewallActionDetach**](FirewallApi.md#FirewallActionDetach) | **Post** /firewall/{firewallId}/actions/detach | /actions/detach
[**FirewallActionTransfer**](FirewallApi.md#FirewallActionTransfer) | **Post** /firewall/{firewallId}/actions/transfer | /actions/transfer
[**FirewallCreate**](FirewallApi.md#FirewallCreate) | **Post** /firewall | Create
[**FirewallDelete**](FirewallApi.md#FirewallDelete) | **Delete** /firewall/{firewallId} | Delete
[**FirewallDeleteAccessrightsIdentity**](FirewallApi.md#FirewallDeleteAccessrightsIdentity) | **Delete** /firewall/{firewallId}/accessrights/{identity} | /accessrights/:identity
[**FirewallDeleteEgressRuleId**](FirewallApi.md#FirewallDeleteEgressRuleId) | **Delete** /firewall/{firewallId}/egress/{ruleId} | /egress/:ruleId
[**FirewallDeleteIngressRuleId**](FirewallApi.md#FirewallDeleteIngressRuleId) | **Delete** /firewall/{firewallId}/ingress/{ruleId} | /ingress/:ruleId
[**FirewallDeleteTagKey**](FirewallApi.md#FirewallDeleteTagKey) | **Delete** /firewall/{firewallId}/tag/{key} | /tag/:key
[**FirewallGetEgressRuleId**](FirewallApi.md#FirewallGetEgressRuleId) | **Get** /firewall/{firewallId}/egress/{ruleId} | /egress/:ruleId
[**FirewallGetIngressRuleId**](FirewallApi.md#FirewallGetIngressRuleId) | **Get** /firewall/{firewallId}/ingress/{ruleId} | /ingress/:ruleId
[**FirewallGetServicesServiceId**](FirewallApi.md#FirewallGetServicesServiceId) | **Get** /firewall/{firewallId}/services/{serviceId} | /services/:serviceId
[**FirewallGetTag**](FirewallApi.md#FirewallGetTag) | **Get** /firewall/{firewallId}/tag | /tag
[**FirewallList**](FirewallApi.md#FirewallList) | **Get** /firewall | List
[**FirewallListAccessrights**](FirewallApi.md#FirewallListAccessrights) | **Get** /firewall/{firewallId}/accessrights | /accessrights
[**FirewallListEgress**](FirewallApi.md#FirewallListEgress) | **Get** /firewall/{firewallId}/egress | /egress
[**FirewallListIngress**](FirewallApi.md#FirewallListIngress) | **Get** /firewall/{firewallId}/ingress | /ingress
[**FirewallListQueue**](FirewallApi.md#FirewallListQueue) | **Get** /firewall/{firewallId}/queue | /queue
[**FirewallListServices**](FirewallApi.md#FirewallListServices) | **Get** /firewall/{firewallId}/services | /services
[**FirewallPatchTag**](FirewallApi.md#FirewallPatchTag) | **Patch** /firewall/{firewallId}/tag | /tag
[**FirewallPostAccessrights**](FirewallApi.md#FirewallPostAccessrights) | **Post** /firewall/{firewallId}/accessrights | /accessrights
[**FirewallPostEgress**](FirewallApi.md#FirewallPostEgress) | **Post** /firewall/{firewallId}/egress | /egress
[**FirewallPostIngress**](FirewallApi.md#FirewallPostIngress) | **Post** /firewall/{firewallId}/ingress | /ingress
[**FirewallShow**](FirewallApi.md#FirewallShow) | **Get** /firewall/{firewallId} | Get
[**FirewallUpdate**](FirewallApi.md#FirewallUpdate) | **Patch** /firewall/{firewallId} | Update


# **FirewallActionAttach**
> Firewall FirewallActionAttach(ctx, firewallId, firewallActionAttach)
/actions/attach

Action attach

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **firewallActionAttach** | [**FirewallActionAttach**](FirewallActionAttach.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallActionDetach**
> Firewall FirewallActionDetach(ctx, firewallId)
/actions/detach

Action detach

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallActionTransfer**
> Firewall FirewallActionTransfer(ctx, firewallId, firewallActionTransfer)
/actions/transfer

Action transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **firewallActionTransfer** | [**FirewallActionTransfer**](FirewallActionTransfer.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallCreate**
> Firewall FirewallCreate(ctx, firewallCreate)
Create

Create firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallCreate** | [**FirewallCreate**](FirewallCreate.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallDelete**
> FirewallDelete(ctx, firewallId)
Delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

 (empty response body)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallDeleteAccessrightsIdentity**
> Firewall FirewallDeleteAccessrightsIdentity(ctx, firewallId, identity)
/accessrights/:identity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **identity** | **string**| identity | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallDeleteEgressRuleId**
> InlineResponse2001 FirewallDeleteEgressRuleId(ctx, firewallId, ruleId)
/egress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallDeleteIngressRuleId**
> InlineResponse2001 FirewallDeleteIngressRuleId(ctx, firewallId, ruleId)
/ingress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallDeleteTagKey**
> map[string]interface{} FirewallDeleteTagKey(ctx, firewallId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **key** | **string**| key | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGetEgressRuleId**
> InlineResponse2001 FirewallGetEgressRuleId(ctx, firewallId, ruleId)
/egress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGetIngressRuleId**
> InlineResponse2001 FirewallGetIngressRuleId(ctx, firewallId, ruleId)
/ingress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGetServicesServiceId**
> FirewallServices FirewallGetServicesServiceId(ctx, firewallId, serviceId)
/services/:serviceId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **serviceId** | **string**| serviceId | 

### Return type

[**FirewallServices**](firewall.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGetTag**
> map[string]interface{} FirewallGetTag(ctx, firewallId)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallList**
> []Firewall FirewallList(ctx, optional)
List

List firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **tag** | [**optional.Interface of map[string]string**](string.md)| Filter by tag | 

### Return type

[**[]Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallListAccessrights**
> []string FirewallListAccessrights(ctx, firewallId)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

**[]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallListEgress**
> []InlineResponse2001 FirewallListEgress(ctx, firewallId)
/egress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallListIngress**
> []InlineResponse2001 FirewallListIngress(ctx, firewallId)
/ingress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallListQueue**
> []Event FirewallListQueue(ctx, firewallId)
/queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**[]Event**](event.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallListServices**
> []FirewallServices FirewallListServices(ctx, firewallId)
/services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**[]FirewallServices**](firewall.services.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallPatchTag**
> map[string]interface{} FirewallPatchTag(ctx, firewallId, requestBody)
/tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallPostAccessrights**
> Firewall FirewallPostAccessrights(ctx, firewallId, firewallPostAccessrights)
/accessrights

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **firewallPostAccessrights** | [**FirewallPostAccessrights**](FirewallPostAccessrights.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallPostEgress**
> InlineResponse2001 FirewallPostEgress(ctx, firewallId, firewallPostEgress)
/egress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **firewallPostEgress** | [**FirewallPostEgress**](FirewallPostEgress.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallPostIngress**
> InlineResponse2001 FirewallPostIngress(ctx, firewallId, firewallPostIngress)
/ingress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **firewallPostIngress** | [**FirewallPostIngress**](FirewallPostIngress.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallShow**
> Firewall FirewallShow(ctx, firewallId)
Get

Returns a single firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallUpdate**
> Firewall FirewallUpdate(ctx, firewallId, firewallUpdate)
Update

Returns modified firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **firewallUpdate** | [**FirewallUpdate**](FirewallUpdate.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

