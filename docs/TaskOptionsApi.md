# \TaskOptionsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskOption**](TaskOptionsApi.md#CreateTaskOption) | **Post** /tasks/{task_id}/task_options | Creates a new task option.
[**DestroyTaskOption**](TaskOptionsApi.md#DestroyTaskOption) | **Delete** /tasks/{task_id}/task_options/{task_option_id} | Destroys a specific task option.
[**IndexTaskOptions**](TaskOptionsApi.md#IndexTaskOptions) | **Get** /tasks/{task_id}/task_options | Lists all task options.
[**ShowTaskOption**](TaskOptionsApi.md#ShowTaskOption) | **Get** /tasks/{task_id}/task_options/{task_option_id} | Displays a specific task option.
[**UpdateTaskOption**](TaskOptionsApi.md#UpdateTaskOption) | **Put** /tasks/{task_id}/task_options/{task_option_id} | Updates a specific task option.



## CreateTaskOption

> TaskOption CreateTaskOption(ctx, taskId, taskOptionBody)

Creates a new task option.

**API Key Scope**: task_options / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskOptionBody** | [**TaskOption**](TaskOption.md)|  | 

### Return type

[**TaskOption**](task_option.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTaskOption

> DestroyTaskOption(ctx, taskId, taskOptionId)

Destroys a specific task option.

**API Key Scope**: task_options / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskOptionId** | **int32**| Numeric ID of task option. | 

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


## IndexTaskOptions

> TaskOptionCollection IndexTaskOptions(ctx, taskId, optional)

Lists all task options.

**API Key Scope**: task_options / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
 **optional** | ***IndexTaskOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **type_** | **optional.String**| Filter on type | 
 **value** | **optional.String**| Filter on value | 

### Return type

[**TaskOptionCollection**](task_option_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskOption

> TaskOption ShowTaskOption(ctx, taskId, taskOptionId)

Displays a specific task option.

**API Key Scope**: task_options / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskOptionId** | **int32**| Numeric ID of task option. | 

### Return type

[**TaskOption**](task_option.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskOption

> TaskOption UpdateTaskOption(ctx, taskId, taskOptionId, taskOptionBody)

Updates a specific task option.

**API Key Scope**: task_options / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskOptionId** | **int32**| Numeric ID of task option. | 
**taskOptionBody** | [**TaskOption**](TaskOption.md)|  | 

### Return type

[**TaskOption**](task_option.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

