# \IamProjectRoleApi

All URIs are relative to *https://api.hyperone.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamProjectRoleCreate**](IamProjectRoleApi.md#IamProjectRoleCreate) | **Post** /iam/project/{projectId}/role | Create iam/role
[**IamProjectRoleDelete**](IamProjectRoleApi.md#IamProjectRoleDelete) | **Delete** /iam/project/{projectId}/role/{roleId} | Delete iam/role
[**IamProjectRoleEventGet**](IamProjectRoleApi.md#IamProjectRoleEventGet) | **Get** /iam/project/{projectId}/role/{roleId}/event/{eventId} | Get iam/role.event
[**IamProjectRoleEventList**](IamProjectRoleApi.md#IamProjectRoleEventList) | **Get** /iam/project/{projectId}/role/{roleId}/event | List iam/role.event
[**IamProjectRoleGet**](IamProjectRoleApi.md#IamProjectRoleGet) | **Get** /iam/project/{projectId}/role/{roleId} | Get iam/role
[**IamProjectRoleList**](IamProjectRoleApi.md#IamProjectRoleList) | **Get** /iam/project/{projectId}/role | List iam/role
[**IamProjectRolePermissionCreate**](IamProjectRoleApi.md#IamProjectRolePermissionCreate) | **Post** /iam/project/{projectId}/role/{roleId}/permission | Create iam/role.permission
[**IamProjectRolePermissionDelete**](IamProjectRoleApi.md#IamProjectRolePermissionDelete) | **Delete** /iam/project/{projectId}/role/{roleId}/permission/{permissionId} | Delete iam/role.permission
[**IamProjectRolePermissionGet**](IamProjectRoleApi.md#IamProjectRolePermissionGet) | **Get** /iam/project/{projectId}/role/{roleId}/permission/{permissionId} | Get iam/role.permission
[**IamProjectRolePermissionList**](IamProjectRoleApi.md#IamProjectRolePermissionList) | **Get** /iam/project/{projectId}/role/{roleId}/permission | List iam/role.permission
[**IamProjectRolePermissionPut**](IamProjectRoleApi.md#IamProjectRolePermissionPut) | **Put** /iam/project/{projectId}/role/{roleId}/permission | Replace iam/role.permission
[**IamProjectRoleServiceGet**](IamProjectRoleApi.md#IamProjectRoleServiceGet) | **Get** /iam/project/{projectId}/role/{roleId}/service/{serviceId} | Get iam/role.service
[**IamProjectRoleServiceList**](IamProjectRoleApi.md#IamProjectRoleServiceList) | **Get** /iam/project/{projectId}/role/{roleId}/service | List iam/role.service
[**IamProjectRoleTagCreate**](IamProjectRoleApi.md#IamProjectRoleTagCreate) | **Post** /iam/project/{projectId}/role/{roleId}/tag | Create iam/role.tag
[**IamProjectRoleTagDelete**](IamProjectRoleApi.md#IamProjectRoleTagDelete) | **Delete** /iam/project/{projectId}/role/{roleId}/tag/{tagId} | Delete iam/role.tag
[**IamProjectRoleTagGet**](IamProjectRoleApi.md#IamProjectRoleTagGet) | **Get** /iam/project/{projectId}/role/{roleId}/tag/{tagId} | Get iam/role.tag
[**IamProjectRoleTagList**](IamProjectRoleApi.md#IamProjectRoleTagList) | **Get** /iam/project/{projectId}/role/{roleId}/tag | List iam/role.tag
[**IamProjectRoleTagPut**](IamProjectRoleApi.md#IamProjectRoleTagPut) | **Put** /iam/project/{projectId}/role/{roleId}/tag | Replace iam/role.tag
[**IamProjectRoleUpdate**](IamProjectRoleApi.md#IamProjectRoleUpdate) | **Patch** /iam/project/{projectId}/role/{roleId} | Update iam/role



## IamProjectRoleCreate

> Role IamProjectRoleCreate(ctx, projectId, iamProjectRoleCreate, optional)

Create iam/role

Create role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**iamProjectRoleCreate** | [**IamProjectRoleCreate**](IamProjectRoleCreate.md)|  | 
 **optional** | ***IamProjectRoleCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectRoleCreateOpts struct


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


## IamProjectRoleDelete

> IamProjectRoleDelete(ctx, projectId, roleId)

Delete iam/role

Delete role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleEventGet

> Event IamProjectRoleEventGet(ctx, projectId, roleId, eventId)

Get iam/role.event

Get iam/role.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleEventList

> []Event IamProjectRoleEventList(ctx, projectId, roleId, optional)

List iam/role.event

List iam/role.event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**roleId** | **string**| Role Id | 
 **optional** | ***IamProjectRoleEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectRoleEventListOpts struct


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


## IamProjectRoleGet

> Role IamProjectRoleGet(ctx, projectId, roleId)

Get iam/role

Returns a single role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleList

> []Role IamProjectRoleList(ctx, projectId, optional)

List iam/role

List role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***IamProjectRoleListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IamProjectRoleListOpts struct


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


## IamProjectRolePermissionCreate

> IamPermission IamProjectRolePermissionCreate(ctx, projectId, roleId, iamPermission)

Create iam/role.permission

Create iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRolePermissionDelete

> IamPermission IamProjectRolePermissionDelete(ctx, projectId, roleId, permissionId)

Delete iam/role.permission

Delete iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRolePermissionGet

> IamPermission IamProjectRolePermissionGet(ctx, projectId, roleId, permissionId)

Get iam/role.permission

Get iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRolePermissionList

> []IamPermission IamProjectRolePermissionList(ctx, projectId, roleId)

List iam/role.permission

List iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRolePermissionPut

> []IamPermission IamProjectRolePermissionPut(ctx, projectId, roleId, iamPermission)

Replace iam/role.permission

Replace iam/role.permission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleServiceGet

> ResourceService IamProjectRoleServiceGet(ctx, projectId, roleId, serviceId)

Get iam/role.service

Get iam/role.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleServiceList

> []ResourceService IamProjectRoleServiceList(ctx, projectId, roleId)

List iam/role.service

List iam/role.service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleTagCreate

> Tag IamProjectRoleTagCreate(ctx, projectId, roleId, tag)

Create iam/role.tag

Create iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleTagDelete

> IamProjectRoleTagDelete(ctx, projectId, roleId, tagId)

Delete iam/role.tag

Delete iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleTagGet

> Tag IamProjectRoleTagGet(ctx, projectId, roleId, tagId)

Get iam/role.tag

Get iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleTagList

> []Tag IamProjectRoleTagList(ctx, projectId, roleId)

List iam/role.tag

List iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleTagPut

> []Tag IamProjectRoleTagPut(ctx, projectId, roleId, tag)

Replace iam/role.tag

Replace iam/role.tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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


## IamProjectRoleUpdate

> Role IamProjectRoleUpdate(ctx, projectId, roleId, iamProjectRoleUpdate)

Update iam/role

Returns modified role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
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

