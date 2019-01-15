# \FirewallApi

All URIs are relative to *https://api.hyperone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionFirewallAttach**](FirewallApi.md#ActionFirewallAttach) | **Post** /firewall/{firewallId}/actions/attach | /actions/attach
[**ActionFirewallDetach**](FirewallApi.md#ActionFirewallDetach) | **Post** /firewall/{firewallId}/actions/detach | /actions/detach
[**ActionFirewallTransfer**](FirewallApi.md#ActionFirewallTransfer) | **Post** /firewall/{firewallId}/actions/transfer | /actions/transfer
[**CreateFirewall**](FirewallApi.md#CreateFirewall) | **Post** /firewall | Create
[**DeleteFirewall**](FirewallApi.md#DeleteFirewall) | **Delete** /firewall/{firewallId} | Delete
[**ListFirewall**](FirewallApi.md#ListFirewall) | **Get** /firewall | List
[**OperationFirewallDeleteaccessrightsIdentity**](FirewallApi.md#OperationFirewallDeleteaccessrightsIdentity) | **Delete** /firewall/{firewallId}/accessrights/{identity} | /accessrights/:identity
[**OperationFirewallDeleteegressRuleId**](FirewallApi.md#OperationFirewallDeleteegressRuleId) | **Delete** /firewall/{firewallId}/egress/{ruleId} | /egress/:ruleId
[**OperationFirewallDeleteingressRuleId**](FirewallApi.md#OperationFirewallDeleteingressRuleId) | **Delete** /firewall/{firewallId}/ingress/{ruleId} | /ingress/:ruleId
[**OperationFirewallDeletetagKey**](FirewallApi.md#OperationFirewallDeletetagKey) | **Delete** /firewall/{firewallId}/tag/{key} | /tag/:key
[**OperationFirewallGetegressRuleId**](FirewallApi.md#OperationFirewallGetegressRuleId) | **Get** /firewall/{firewallId}/egress/{ruleId} | /egress/:ruleId
[**OperationFirewallGetingressRuleId**](FirewallApi.md#OperationFirewallGetingressRuleId) | **Get** /firewall/{firewallId}/ingress/{ruleId} | /ingress/:ruleId
[**OperationFirewallGetservicesServiceId**](FirewallApi.md#OperationFirewallGetservicesServiceId) | **Get** /firewall/{firewallId}/services/{serviceId} | /services/:serviceId
[**OperationFirewallGettag**](FirewallApi.md#OperationFirewallGettag) | **Get** /firewall/{firewallId}/tag/ | /tag/
[**OperationFirewallListaccessrights**](FirewallApi.md#OperationFirewallListaccessrights) | **Get** /firewall/{firewallId}/accessrights/ | /accessrights/
[**OperationFirewallListegress**](FirewallApi.md#OperationFirewallListegress) | **Get** /firewall/{firewallId}/egress | /egress
[**OperationFirewallListingress**](FirewallApi.md#OperationFirewallListingress) | **Get** /firewall/{firewallId}/ingress | /ingress
[**OperationFirewallListqueue**](FirewallApi.md#OperationFirewallListqueue) | **Get** /firewall/{firewallId}/queue/ | /queue/
[**OperationFirewallListservices**](FirewallApi.md#OperationFirewallListservices) | **Get** /firewall/{firewallId}/services/ | /services/
[**OperationFirewallPatchtag**](FirewallApi.md#OperationFirewallPatchtag) | **Patch** /firewall/{firewallId}/tag/ | /tag/
[**OperationFirewallPostaccessrights**](FirewallApi.md#OperationFirewallPostaccessrights) | **Post** /firewall/{firewallId}/accessrights/ | /accessrights/
[**OperationFirewallPostegress**](FirewallApi.md#OperationFirewallPostegress) | **Post** /firewall/{firewallId}/egress | /egress
[**OperationFirewallPostingress**](FirewallApi.md#OperationFirewallPostingress) | **Post** /firewall/{firewallId}/ingress | /ingress
[**ShowFirewall**](FirewallApi.md#ShowFirewall) | **Get** /firewall/{firewallId} | Get
[**UpdateFirewall**](FirewallApi.md#UpdateFirewall) | **Patch** /firewall/{firewallId} | Update


# **ActionFirewallAttach**
> Firewall ActionFirewallAttach(ctx, firewallId)
/actions/attach

Action attach

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

# **ActionFirewallDetach**
> Firewall ActionFirewallDetach(ctx, firewallId)
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

# **ActionFirewallTransfer**
> Firewall ActionFirewallTransfer(ctx, firewallId, optional)
/actions/transfer

Action transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
 **optional** | ***ActionFirewallTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionFirewallTransferOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject5** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFirewall**
> Firewall CreateFirewall(ctx, firewallCreate)
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

# **DeleteFirewall**
> DeleteFirewall(ctx, firewallId)
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
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewall**
> []Firewall ListFirewall(ctx, optional)
List

List firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFirewallOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFirewallOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 

### Return type

[**[]Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallDeleteaccessrightsIdentity**
> Firewall OperationFirewallDeleteaccessrightsIdentity(ctx, firewallId, identity)
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

# **OperationFirewallDeleteegressRuleId**
> InlineResponse200 OperationFirewallDeleteegressRuleId(ctx, firewallId, ruleId)
/egress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallDeleteingressRuleId**
> InlineResponse200 OperationFirewallDeleteingressRuleId(ctx, firewallId, ruleId)
/ingress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallDeletetagKey**
> map[string]string OperationFirewallDeletetagKey(ctx, firewallId, key)
/tag/:key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **key** | **string**| key | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallGetegressRuleId**
> InlineResponse200 OperationFirewallGetegressRuleId(ctx, firewallId, ruleId)
/egress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallGetingressRuleId**
> InlineResponse200 OperationFirewallGetingressRuleId(ctx, firewallId, ruleId)
/ingress/:ruleId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **ruleId** | **string**| ruleId | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallGetservicesServiceId**
> FirewallServices OperationFirewallGetservicesServiceId(ctx, firewallId, serviceId)
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

# **OperationFirewallGettag**
> map[string]string OperationFirewallGettag(ctx, firewallId)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallListaccessrights**
> []string OperationFirewallListaccessrights(ctx, firewallId)
/accessrights/

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

# **OperationFirewallListegress**
> []InlineResponse200 OperationFirewallListegress(ctx, firewallId)
/egress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallListingress**
> []InlineResponse200 OperationFirewallListingress(ctx, firewallId)
/ingress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallListqueue**
> []Event OperationFirewallListqueue(ctx, firewallId)
/queue/

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

# **OperationFirewallListservices**
> []FirewallServices OperationFirewallListservices(ctx, firewallId)
/services/

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

# **OperationFirewallPatchtag**
> map[string]string OperationFirewallPatchtag(ctx, firewallId, requestBody)
/tag/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **requestBody** | [**map[string]string**](string.md)|  | 

### Return type

**map[string]string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallPostaccessrights**
> string OperationFirewallPostaccessrights(ctx, firewallId, resourceAccessRight)
/accessrights/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
  **resourceAccessRight** | [**ResourceAccessRight**](ResourceAccessRight.md)|  | 

### Return type

**string**

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallPostegress**
> InlineResponse200 OperationFirewallPostegress(ctx, firewallId, optional)
/egress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
 **optional** | ***OperationFirewallPostegressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationFirewallPostegressOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject7** | [**optional.Interface of InlineObject7**](InlineObject7.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationFirewallPostingress**
> InlineResponse200 OperationFirewallPostingress(ctx, firewallId, optional)
/ingress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
 **optional** | ***OperationFirewallPostingressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationFirewallPostingressOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject6** | [**optional.Interface of InlineObject6**](InlineObject6.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFirewall**
> Firewall ShowFirewall(ctx, firewallId)
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

# **UpdateFirewall**
> Firewall UpdateFirewall(ctx, firewallId, optional)
Update

Returns modified firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallId** | **string**| ID of firewall | 
 **optional** | ***UpdateFirewallOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateFirewallOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject4** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[Project](../README.md#Project), [ServiceAccount](../README.md#ServiceAccount), [Session](../README.md#Session)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

