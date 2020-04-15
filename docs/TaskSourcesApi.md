# \TaskSourcesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskSource**](TaskSourcesApi.md#CreateTaskSource) | **Post** /tasks/{task_id}/task_sources | Creates a new task source.
[**DestroyTaskSource**](TaskSourcesApi.md#DestroyTaskSource) | **Delete** /tasks/{task_id}/task_sources/{task_source_id} | Destroys a specific task source.
[**IndexTaskSources**](TaskSourcesApi.md#IndexTaskSources) | **Get** /tasks/{task_id}/task_sources | Lists all task sources.
[**ShowTaskSource**](TaskSourcesApi.md#ShowTaskSource) | **Get** /tasks/{task_id}/task_sources/{task_source_id} | Displays a specific task source.
[**UpdateTaskSource**](TaskSourcesApi.md#UpdateTaskSource) | **Put** /tasks/{task_id}/task_sources/{task_source_id} | Updates a specific task source.



## CreateTaskSource

> TaskSourceDown CreateTaskSource(ctx, taskId, taskSourceBody)

Creates a new task source.

**API Key Scope**: task_sources / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskSourceBody** | [**TaskSourceUp**](TaskSourceUp.md)|  | 

### Return type

[**TaskSourceDown**](task_source_down.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTaskSource

> DestroyTaskSource(ctx, taskId, taskSourceId)

Destroys a specific task source.

**API Key Scope**: task_sources / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskSourceId** | **int32**| Numeric ID of task source. | 

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


## IndexTaskSources

> TaskSourceCollection IndexTaskSources(ctx, taskId, optional)

Lists all task sources.

**API Key Scope**: task_sources / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
 **optional** | ***IndexTaskSourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskSourcesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **fileId** | **optional.String**| Filter on file id | 
 **importFileId** | **optional.String**| Filter on import file id | 
 **tapeId** | **optional.String**| Filter on tape id | 
 **poolId** | **optional.String**| Filter on a pool id | 

### Return type

[**TaskSourceCollection**](task_source_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskSource

> TaskSourceDown ShowTaskSource(ctx, taskId, taskSourceId)

Displays a specific task source.

**API Key Scope**: task_sources / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskSourceId** | **int32**| Numeric ID of task source. | 

### Return type

[**TaskSourceDown**](task_source_down.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskSource

> TaskSourceDown UpdateTaskSource(ctx, taskId, taskSourceId, taskSourceBody)

Updates a specific task source.

**API Key Scope**: task_sources / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskSourceId** | **int32**| Numeric ID of task source. | 
**taskSourceBody** | [**TaskSourceUp**](TaskSourceUp.md)|  | 

### Return type

[**TaskSourceDown**](task_source_down.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

