# \IamOrganisationRoleApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamOrganisationRoleCreate**](IamOrganisationRoleApi.md#IamOrganisationRoleCreate) | **Post** /iam/organisation/{organisationId}/role | Create iam/role
[**IamOrganisationRoleDelete**](IamOrganisationRoleApi.md#IamOrganisationRoleDelete) | **Delete** /iam/organisation/{organisationId}/role/{roleId} | Delete iam/role
[**IamOrganisationRoleEventGet**](IamOrganisationRoleApi.md#IamOrganisationRoleEventGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/event/{eventId} | Get iam/role.event
[**IamOrganisationRoleEventList**](IamOrganisationRoleApi.md#IamOrganisationRoleEventList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/event | List iam/role.event
[**IamOrganisationRoleGet**](IamOrganisationRoleApi.md#IamOrganisationRoleGet) | **Get** /iam/organisation/{organisationId}/role/{roleId} | Get iam/role
[**IamOrganisationRoleList**](IamOrganisationRoleApi.md#IamOrganisationRoleList) | **Get** /iam/organisation/{organisationId}/role | List iam/role
[**IamOrganisationRolePermissionCreate**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionCreate) | **Post** /iam/organisation/{organisationId}/role/{roleId}/permission | Create iam/role.permission
[**IamOrganisationRolePermissionDelete**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionDelete) | **Delete** /iam/organisation/{organisationId}/role/{roleId}/permission/{permissionId} | Delete iam/role.permission
[**IamOrganisationRolePermissionGet**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/permission/{permissionId} | Get iam/role.permission
[**IamOrganisationRolePermissionList**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/permission | List iam/role.permission
[**IamOrganisationRolePermissionPut**](IamOrganisationRoleApi.md#IamOrganisationRolePermissionPut) | **Put** /iam/organisation/{organisationId}/role/{roleId}/permission | Replace iam/role.permission
[**IamOrganisationRoleServiceGet**](IamOrganisationRoleApi.md#IamOrganisationRoleServiceGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/service/{serviceId} | Get iam/role.service
[**IamOrganisationRoleServiceList**](IamOrganisationRoleApi.md#IamOrganisationRoleServiceList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/service | List iam/role.service
[**IamOrganisationRoleTagCreate**](IamOrganisationRoleApi.md#IamOrganisationRoleTagCreate) | **Post** /iam/organisation/{organisationId}/role/{roleId}/tag | Create iam/role.tag
[**IamOrganisationRoleTagDelete**](IamOrganisationRoleApi.md#IamOrganisationRoleTagDelete) | **Delete** /iam/organisation/{organisationId}/role/{roleId}/tag/{tagId} | Delete iam/role.tag
[**IamOrganisationRoleTagGet**](IamOrganisationRoleApi.md#IamOrganisationRoleTagGet) | **Get** /iam/organisation/{organisationId}/role/{roleId}/tag/{tagId} | Get iam/role.tag
[**IamOrganisationRoleTagList**](IamOrganisationRoleApi.md#IamOrganisationRoleTagList) | **Get** /iam/organisation/{organisationId}/role/{roleId}/tag | List iam/role.tag
[**IamOrganisationRoleTagPut**](IamOrganisationRoleApi.md#IamOrganisationRoleTagPut) | **Put** /iam/organisation/{organisationId}/role/{roleId}/tag | Replace iam/role.tag
[**IamOrganisationRoleUpdate**](IamOrganisationRoleApi.md#IamOrganisationRoleUpdate) | **Patch** /iam/organisation/{organisationId}/role/{roleId} | Update iam/role



## IamOrganisationRoleCreate

> Role IamOrganisationRoleCreate(ctx, organisationId, iamProjectRoleCreate, optional)

Create iam/role

Create role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**iamProjectRoleCreate** | [**IamProjectRoleCreate**](IamProjectRoleCreate.md)|  | 
 **optional** | ***IamOrganisationRoleCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationRoleCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotency key | 

### Return type

[**Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRoleDelete

> IamOrganisationRoleDelete(ctx, organisationId, roleId)

Delete iam/role

Delete role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 

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


## IamOrganisationRoleEventGet

> Event IamOrganisationRoleEventGet(ctx, organisationId, roleId, eventId)

Get iam/role.event

Get iam/role.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
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


## IamOrganisationRoleEventList

> []Event IamOrganisationRoleEventList(ctx, organisationId, roleId, optional)

List iam/role.event

List iam/role.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
 **optional** | ***IamOrganisationRoleEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationRoleEventListOpts struct


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


## IamOrganisationRoleGet

> Role IamOrganisationRoleGet(ctx, organisationId, roleId)

Get iam/role

Returns a single role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 

### Return type

[**Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRoleList

> []Role IamOrganisationRoleList(ctx, organisationId, optional)

List iam/role

List role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
 **optional** | ***IamOrganisationRoleListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamOrganisationRoleListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Filter by name | 
 **tagValue** | **optional.String**| Filter by tag.value | 
 **tagKey** | **optional.String**| Filter by tag.key | 

### Return type

[**[]Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionCreate

> IamPermission IamOrganisationRolePermissionCreate(ctx, organisationId, roleId, iamPermission)

Create iam/role.permission

Create iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
**iamPermission** | [**IamPermission**](IamPermission.md)|  | 

### Return type

[**IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionDelete

> IamPermission IamOrganisationRolePermissionDelete(ctx, organisationId, roleId, permissionId)

Delete iam/role.permission

Delete iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
**permissionId** | **string**| permissionId | 

### Return type

[**IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionGet

> IamPermission IamOrganisationRolePermissionGet(ctx, organisationId, roleId, permissionId)

Get iam/role.permission

Get iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
**permissionId** | **string**| permissionId | 

### Return type

[**IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionList

> []IamPermission IamOrganisationRolePermissionList(ctx, organisationId, roleId)

List iam/role.permission

List iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 

### Return type

[**[]IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRolePermissionPut

> []IamPermission IamOrganisationRolePermissionPut(ctx, organisationId, roleId, iamPermission)

Replace iam/role.permission

Replace iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
**iamPermission** | [**[]IamPermission**](iam.permission.md)|  | 

### Return type

[**[]IamPermission**](iam.permission.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamOrganisationRoleServiceGet

> ResourceService IamOrganisationRoleServiceGet(ctx, organisationId, roleId, serviceId)

Get iam/role.service

Get iam/role.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
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


## IamOrganisationRoleServiceList

> []ResourceService IamOrganisationRoleServiceList(ctx, organisationId, roleId)

List iam/role.service

List iam/role.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 

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


## IamOrganisationRoleTagCreate

> Tag IamOrganisationRoleTagCreate(ctx, organisationId, roleId, tag)

Create iam/role.tag

Create iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
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


## IamOrganisationRoleTagDelete

> IamOrganisationRoleTagDelete(ctx, organisationId, roleId, tagId)

Delete iam/role.tag

Delete iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
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


## IamOrganisationRoleTagGet

> Tag IamOrganisationRoleTagGet(ctx, organisationId, roleId, tagId)

Get iam/role.tag

Get iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
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


## IamOrganisationRoleTagList

> []Tag IamOrganisationRoleTagList(ctx, organisationId, roleId)

List iam/role.tag

List iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 

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


## IamOrganisationRoleTagPut

> []Tag IamOrganisationRoleTagPut(ctx, organisationId, roleId, tag)

Replace iam/role.tag

Replace iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
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


## IamOrganisationRoleUpdate

> Role IamOrganisationRoleUpdate(ctx, organisationId, roleId, iamProjectRoleUpdate)

Update iam/role

Returns modified role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string**| Organisation Id | 
**roleId** | **string**| Role Id | 
**iamProjectRoleUpdate** | [**IamProjectRoleUpdate**](IamProjectRoleUpdate.md)|  | 

### Return type

[**Role**](role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

