# \NetworkingProjectFirewallApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkingProjectFirewallCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall | Create networking/firewall
[**NetworkingProjectFirewallDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId} | Delete networking/firewall
[**NetworkingProjectFirewallEgressCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress | Create networking/firewall.egress
[**NetworkingProjectFirewallEgressDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress/{egressId} | Delete networking/firewall.egress
[**NetworkingProjectFirewallEgressGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress/{egressId} | Get networking/firewall.egress
[**NetworkingProjectFirewallEgressList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress | List networking/firewall.egress
[**NetworkingProjectFirewallEgressPut**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEgressPut) | **Put** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/egress | Replace networking/firewall.egress
[**NetworkingProjectFirewallEventGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEventGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/event/{eventId} | Get networking/firewall.event
[**NetworkingProjectFirewallEventList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallEventList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/event | List networking/firewall.event
[**NetworkingProjectFirewallGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId} | Get networking/firewall
[**NetworkingProjectFirewallIngressCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress | Create networking/firewall.ingress
[**NetworkingProjectFirewallIngressDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress/{ingressId} | Delete networking/firewall.ingress
[**NetworkingProjectFirewallIngressGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress/{ingressId} | Get networking/firewall.ingress
[**NetworkingProjectFirewallIngressList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress | List networking/firewall.ingress
[**NetworkingProjectFirewallIngressPut**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallIngressPut) | **Put** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/ingress | Replace networking/firewall.ingress
[**NetworkingProjectFirewallList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallList) | **Get** /networking/{locationId}/project/{projectId}/firewall | List networking/firewall
[**NetworkingProjectFirewallServiceGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallServiceGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/service/{serviceId} | Get networking/firewall.service
[**NetworkingProjectFirewallServiceList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallServiceList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/service | List networking/firewall.service
[**NetworkingProjectFirewallTagCreate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagCreate) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag | Create networking/firewall.tag
[**NetworkingProjectFirewallTagDelete**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagDelete) | **Delete** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag/{tagId} | Delete networking/firewall.tag
[**NetworkingProjectFirewallTagGet**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagGet) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag/{tagId} | Get networking/firewall.tag
[**NetworkingProjectFirewallTagList**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagList) | **Get** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag | List networking/firewall.tag
[**NetworkingProjectFirewallTagPut**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTagPut) | **Put** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/tag | Replace networking/firewall.tag
[**NetworkingProjectFirewallTransfer**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallTransfer) | **Post** /networking/{locationId}/project/{projectId}/firewall/{firewallId}/actions/transfer | Transfer networking/firewall
[**NetworkingProjectFirewallUpdate**](NetworkingProjectFirewallApi.md#NetworkingProjectFirewallUpdate) | **Patch** /networking/{locationId}/project/{projectId}/firewall/{firewallId} | Update networking/firewall



## NetworkingProjectFirewallCreate

> Firewall NetworkingProjectFirewallCreate(ctx, projectId, locationId, networkingProjectFirewallCreate, optional)

Create networking/firewall

Create firewall

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**networkingProjectFirewallCreate** | [**NetworkingProjectFirewallCreate**](NetworkingProjectFirewallCreate.md)|  | 
 **optional** | ***NetworkingProjectFirewallCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectFirewallCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallDelete

> NetworkingProjectFirewallDelete(ctx, projectId, locationId, firewallId)

Delete networking/firewall

Delete firewall

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 

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


## NetworkingProjectFirewallEgressCreate

> NetworkingRule NetworkingProjectFirewallEgressCreate(ctx, projectId, locationId, firewallId, networkingRule)

Create networking/firewall.egress

Create networking/firewall.egress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**networkingRule** | [**NetworkingRule**](NetworkingRule.md)|  | 

### Return type

[**NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEgressDelete

> NetworkingProjectFirewallEgressDelete(ctx, projectId, locationId, firewallId, egressId)

Delete networking/firewall.egress

Delete networking/firewall.egress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**egressId** | **string**| egressId | 

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


## NetworkingProjectFirewallEgressGet

> NetworkingRule NetworkingProjectFirewallEgressGet(ctx, projectId, locationId, firewallId, egressId)

Get networking/firewall.egress

Get networking/firewall.egress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**egressId** | **string**| egressId | 

### Return type

[**NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEgressList

> []NetworkingRule NetworkingProjectFirewallEgressList(ctx, projectId, locationId, firewallId)

List networking/firewall.egress

List networking/firewall.egress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 

### Return type

[**[]NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEgressPut

> []NetworkingRule NetworkingProjectFirewallEgressPut(ctx, projectId, locationId, firewallId, networkingRule)

Replace networking/firewall.egress

Replace networking/firewall.egress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**networkingRule** | [**[]NetworkingRule**](networking.rule.md)|  | 

### Return type

[**[]NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallEventGet

> Event NetworkingProjectFirewallEventGet(ctx, projectId, locationId, firewallId, eventId)

Get networking/firewall.event

Get networking/firewall.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
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


## NetworkingProjectFirewallEventList

> []Event NetworkingProjectFirewallEventList(ctx, projectId, locationId, firewallId, optional)

List networking/firewall.event

List networking/firewall.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
 **optional** | ***NetworkingProjectFirewallEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectFirewallEventListOpts struct


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


## NetworkingProjectFirewallGet

> Firewall NetworkingProjectFirewallGet(ctx, projectId, locationId, firewallId)

Get networking/firewall

Returns a single firewall

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressCreate

> NetworkingRule NetworkingProjectFirewallIngressCreate(ctx, projectId, locationId, firewallId, networkingRule)

Create networking/firewall.ingress

Create networking/firewall.ingress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**networkingRule** | [**NetworkingRule**](NetworkingRule.md)|  | 

### Return type

[**NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressDelete

> NetworkingProjectFirewallIngressDelete(ctx, projectId, locationId, firewallId, ingressId)

Delete networking/firewall.ingress

Delete networking/firewall.ingress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**ingressId** | **string**| ingressId | 

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


## NetworkingProjectFirewallIngressGet

> NetworkingRule NetworkingProjectFirewallIngressGet(ctx, projectId, locationId, firewallId, ingressId)

Get networking/firewall.ingress

Get networking/firewall.ingress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**ingressId** | **string**| ingressId | 

### Return type

[**NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressList

> []NetworkingRule NetworkingProjectFirewallIngressList(ctx, projectId, locationId, firewallId)

List networking/firewall.ingress

List networking/firewall.ingress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 

### Return type

[**[]NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallIngressPut

> []NetworkingRule NetworkingProjectFirewallIngressPut(ctx, projectId, locationId, firewallId, networkingRule)

Replace networking/firewall.ingress

Replace networking/firewall.ingress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**networkingRule** | [**[]NetworkingRule**](networking.rule.md)|  | 

### Return type

[**[]NetworkingRule**](networking.rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallList

> []Firewall NetworkingProjectFirewallList(ctx, projectId, locationId, optional)

List networking/firewall

List firewall

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
 **optional** | ***NetworkingProjectFirewallListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectFirewallListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Firewall**](firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallServiceGet

> ResourceService NetworkingProjectFirewallServiceGet(ctx, projectId, locationId, firewallId, serviceId)

Get networking/firewall.service

Get networking/firewall.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
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


## NetworkingProjectFirewallServiceList

> []ResourceService NetworkingProjectFirewallServiceList(ctx, projectId, locationId, firewallId)

List networking/firewall.service

List networking/firewall.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 

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


## NetworkingProjectFirewallTagCreate

> Tag NetworkingProjectFirewallTagCreate(ctx, projectId, locationId, firewallId, tag)

Create networking/firewall.tag

Create networking/firewall.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
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


## NetworkingProjectFirewallTagDelete

> NetworkingProjectFirewallTagDelete(ctx, projectId, locationId, firewallId, tagId)

Delete networking/firewall.tag

Delete networking/firewall.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
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


## NetworkingProjectFirewallTagGet

> Tag NetworkingProjectFirewallTagGet(ctx, projectId, locationId, firewallId, tagId)

Get networking/firewall.tag

Get networking/firewall.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
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


## NetworkingProjectFirewallTagList

> []Tag NetworkingProjectFirewallTagList(ctx, projectId, locationId, firewallId)

List networking/firewall.tag

List networking/firewall.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 

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


## NetworkingProjectFirewallTagPut

> []Tag NetworkingProjectFirewallTagPut(ctx, projectId, locationId, firewallId, tag)

Replace networking/firewall.tag

Replace networking/firewall.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
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


## NetworkingProjectFirewallTransfer

> Firewall NetworkingProjectFirewallTransfer(ctx, projectId, locationId, firewallId, networkingProjectFirewallTransfer, optional)

Transfer networking/firewall

action transfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**networkingProjectFirewallTransfer** | [**NetworkingProjectFirewallTransfer**](NetworkingProjectFirewallTransfer.md)|  | 
 **optional** | ***NetworkingProjectFirewallTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkingProjectFirewallTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkingProjectFirewallUpdate

> Firewall NetworkingProjectFirewallUpdate(ctx, projectId, locationId, firewallId, networkingProjectFirewallUpdate)

Update networking/firewall

Returns modified firewall

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**locationId** | **string**| Location Id | 
**firewallId** | **string**| Firewall Id | 
**networkingProjectFirewallUpdate** | [**NetworkingProjectFirewallUpdate**](NetworkingProjectFirewallUpdate.md)|  | 

### Return type

[**Firewall**](firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

