# \ContainersApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainer**](ContainersApi.md#CreateContainer) | **Post** /containers | Creates a new container.
[**CreateContainerPrivilege**](ContainersApi.md#CreateContainerPrivilege) | **Post** /containers/{container_id}/container_privileges | Creates a new privilege on the container.
[**DestroyContainer**](ContainersApi.md#DestroyContainer) | **Delete** /containers/{container_id} | Destroys a specific container.
[**DestroyContainerPrivilege**](ContainersApi.md#DestroyContainerPrivilege) | **Delete** /containers/{container_id}/container_privileges/{container_privilege_id} | Destroys a specific privilege.
[**IndexContainerPrivileges**](ContainersApi.md#IndexContainerPrivileges) | **Get** /containers/{container_id}/container_privileges | Lists all privilege on the container.
[**IndexContainers**](ContainersApi.md#IndexContainers) | **Get** /containers | Lists all containers.
[**ShowContainer**](ContainersApi.md#ShowContainer) | **Get** /containers/{container_id} | Displays a specific container.
[**ShowContainerPrivilege**](ContainersApi.md#ShowContainerPrivilege) | **Get** /containers/{container_id}/container_privileges/{container_privilege_id} | Displays a specific privilege.
[**UpdateContainer**](ContainersApi.md#UpdateContainer) | **Put** /containers/{container_id} | Updates a specific container.
[**UpdateContainerPrivilege**](ContainersApi.md#UpdateContainerPrivilege) | **Put** /containers/{container_id}/container_privileges/{container_privilege_id} | Updates a specific privilege.



## CreateContainer

> Container CreateContainer(ctx, containerBody)

Creates a new container.

It **does not** yet create the file structure and configure the samba connection. Use API v1 instead.  **API Key Scope**: containers / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerBody** | [**Container**](Container.md)|  | 

### Return type

[**Container**](container.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerPrivilege

> ContainerPrivilege CreateContainerPrivilege(ctx, containerId, containerPrivilegeBody)

Creates a new privilege on the container.

**API Key Scope**: container_privileges / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
**containerPrivilegeBody** | [**ContainerPrivilege**](ContainerPrivilege.md)|  | 

### Return type

[**ContainerPrivilege**](container_privilege.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyContainer

> DestroyContainer(ctx, containerId)

Destroys a specific container.

**API Key Scope**: containers / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyContainerPrivilege

> DestroyContainerPrivilege(ctx, containerId, containerPrivilegeId)

Destroys a specific privilege.

**API Key Scope**: container_privileges / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
**containerPrivilegeId** | **int32**| Numeric ID of container privilege. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexContainerPrivileges

> ContainerPrivilegeCollection IndexContainerPrivileges(ctx, containerId, optional)

Lists all privilege on the container.

**API Key Scope**: container_privileges / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
 **optional** | ***IndexContainerPrivilegesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexContainerPrivilegesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **privilege** | **optional.String**| Filter on privilege | 
 **type_** | **optional.String**| Filter on type | 

### Return type

[**ContainerPrivilegeCollection**](container_privilege_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexContainers

> ContainerCollection IndexContainers(ctx, optional)

Lists all containers.

**API Key Scope**: containers / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexContainersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexContainersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **comment** | **optional.String**| Filter on comment | 
 **quotaTotalSize** | **optional.String**| Filter on quota total size | 
 **quotaOnCache** | **optional.String**| Filter on quota on cache | 
 **statTotalFiles** | **optional.String**| Filter on stat total files | 
 **statTotalSize** | **optional.String**| Filter on stat total size | 
 **statSizeOnCache** | **optional.String**| Filter on stat size on cache | 
 **guestRight** | **optional.String**| Filter on guest right | 
 **lastUpdate** | **optional.String**| Filter on last update | 

### Return type

[**ContainerCollection**](container_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowContainer

> Container ShowContainer(ctx, containerId)

Displays a specific container.

**API Key Scope**: containers / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 

### Return type

[**Container**](container.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowContainerPrivilege

> ContainerPrivilege ShowContainerPrivilege(ctx, containerId, containerPrivilegeId)

Displays a specific privilege.

**API Key Scope**: container_privileges / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
**containerPrivilegeId** | **int32**| Numeric ID of container privilege. | 

### Return type

[**ContainerPrivilege**](container_privilege.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainer

> Container UpdateContainer(ctx, containerId, containerBody)

Updates a specific container.

It **does not** yet create the file structure and configure the samba connection. Use API v1 instead.  **API Key Scope**: containers / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
**containerBody** | [**Container**](Container.md)|  | 

### Return type

[**Container**](container.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerPrivilege

> ContainerPrivilege UpdateContainerPrivilege(ctx, containerId, containerPrivilegeId, containerPrivilegeBody)

Updates a specific privilege.

**API Key Scope**: container_privileges / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
**containerPrivilegeId** | **int32**| Numeric ID of container privilege. | 
**containerPrivilegeBody** | [**ContainerPrivilege**](ContainerPrivilege.md)|  | 

### Return type

[**ContainerPrivilege**](container_privilege.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

