# \TaskCallbacksApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskCallback**](TaskCallbacksApi.md#CreateTaskCallback) | **Post** /tasks/{task_id}/task_callbacks | Creates a new task callback.
[**DestroyTaskCallback**](TaskCallbacksApi.md#DestroyTaskCallback) | **Delete** /tasks/{task_id}/task_callbacks/{task_callback_id} | Destroys a specific task callback.
[**IndexTaskCallbacks**](TaskCallbacksApi.md#IndexTaskCallbacks) | **Get** /tasks/{task_id}/task_callbacks | Lists all task callbacks.
[**ShowTaskCallback**](TaskCallbacksApi.md#ShowTaskCallback) | **Get** /tasks/{task_id}/task_callbacks/{task_callback_id} | Displays a specific task callback.
[**UpdateTaskCallback**](TaskCallbacksApi.md#UpdateTaskCallback) | **Put** /tasks/{task_id}/task_callbacks/{task_callback_id} | Updates a specific task callback.



## CreateTaskCallback

> TaskCallback CreateTaskCallback(ctx, taskId, taskCallbackBody)

Creates a new task callback.

**API Key Scope**: task_callbacks / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskCallbackBody** | [**TaskCallback**](TaskCallback.md)|  | 

### Return type

[**TaskCallback**](task_callback.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTaskCallback

> DestroyTaskCallback(ctx, taskId, taskCallbackId)

Destroys a specific task callback.

**API Key Scope**: task_callbacks / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskCallbackId** | **int32**| Numeric ID of task callback. | 

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


## IndexTaskCallbacks

> TaskCallbackCollection IndexTaskCallbacks(ctx, taskId, optional)

Lists all task callbacks.

**API Key Scope**: task_callbacks / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
 **optional** | ***IndexTaskCallbacksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskCallbacksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **type_** | **optional.String**| Filter on type | 
 **script** | **optional.String**| Filter on task callback script | 

### Return type

[**TaskCallbackCollection**](task_callback_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskCallback

> TaskCallback ShowTaskCallback(ctx, taskId, taskCallbackId)

Displays a specific task callback.

**API Key Scope**: task_callbacks / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskCallbackId** | **int32**| Numeric ID of task callback. | 

### Return type

[**TaskCallback**](task_callback.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskCallback

> TaskCallback UpdateTaskCallback(ctx, taskId, taskCallbackId, taskCallbackBody)

Updates a specific task callback.

**API Key Scope**: task_callbacks / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskCallbackId** | **int32**| Numeric ID of task callback. | 
**taskCallbackBody** | [**TaskCallback**](TaskCallback.md)|  | 

### Return type

[**TaskCallback**](task_callback.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

