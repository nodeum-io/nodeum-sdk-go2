# \TaskExecutionsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndexTaskExecutions**](TaskExecutionsApi.md#IndexTaskExecutions) | **Get** /task_executions | Lists all task executions.
[**IndexTaskExecutionsByTask**](TaskExecutionsApi.md#IndexTaskExecutionsByTask) | **Get** /tasks/{task_id}/task_executions | Lists all task executions.
[**ShowTaskExecution**](TaskExecutionsApi.md#ShowTaskExecution) | **Get** /task_executions/{task_execution_id} | Displays a specific task execution.
[**ShowTaskExecutionByTask**](TaskExecutionsApi.md#ShowTaskExecutionByTask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id} | Displays a specific task execution.



## IndexTaskExecutions

> TaskExecutionCollection IndexTaskExecutions(ctx, optional)

Lists all task executions.

**API Key Scope**: task_executions / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTaskExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskExecutionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **completeList** | **optional.Bool**| If &#x60;false&#x60;, only includes the last correct execution of a task. | [default to true]
 **id** | **optional.String**| Filter on id | 
 **taskId** | **optional.String**| Filter on task id | 
 **name** | **optional.String**| Filter on name | 
 **workflowType** | **optional.String**| Filter on task workflow type | 
 **workflowAction** | **optional.String**| Filter on task workflow action | 
 **sourceType** | **optional.String**| Filter on task source type | 
 **destinationType** | **optional.String**| Filter on task destination type | 
 **status** | **optional.String**| Filter on status | 
 **logTime** | **optional.String**| Filter on log time | 
 **jobStarted** | **optional.String**| Filter on job started | 
 **jobFinished** | **optional.String**| Filter on job finished | 
 **creationDate** | **optional.String**| Filter on creation date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **toProcessSize** | **optional.String**| Filter on to process size | 
 **processedSize** | **optional.String**| Filter on processed size | 
 **toProcessFiles** | **optional.String**| Filter on to process files | 
 **processedFiles** | **optional.String**| Filter on processed files | 
 **finalizedFiles** | **optional.String**| Filter on finalized files | 
 **estimationTime** | **optional.String**| Filter on estimation time | 
 **bandwidth** | **optional.String**| Filter on bandwidth | 

### Return type

[**TaskExecutionCollection**](task_execution_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexTaskExecutionsByTask

> TaskExecutionCollection IndexTaskExecutionsByTask(ctx, optional)

Lists all task executions.

**API Key Scope**: task_executions / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTaskExecutionsByTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskExecutionsByTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **optional.String**| Filter on task id | 
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **workflowType** | **optional.String**| Filter on task workflow type | 
 **workflowAction** | **optional.String**| Filter on task workflow action | 
 **sourceType** | **optional.String**| Filter on task source type | 
 **destinationType** | **optional.String**| Filter on task destination type | 
 **status** | **optional.String**| Filter on status | 
 **logTime** | **optional.String**| Filter on log time | 
 **jobStarted** | **optional.String**| Filter on job started | 
 **jobFinished** | **optional.String**| Filter on job finished | 
 **toProcessSize** | **optional.String**| Filter on to process size | 
 **processedSize** | **optional.String**| Filter on processed size | 
 **toProcessFiles** | **optional.String**| Filter on to process files | 
 **processedFiles** | **optional.String**| Filter on processed files | 
 **finalizedFiles** | **optional.String**| Filter on finalized files | 
 **estimationTime** | **optional.String**| Filter on estimation time | 
 **bandwidth** | **optional.String**| Filter on bandwidth | 

### Return type

[**TaskExecutionCollection**](task_execution_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskExecution

> TaskExecution ShowTaskExecution(ctx, taskExecutionId)

Displays a specific task execution.

**API Key Scope**: task_executions / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskExecutionId** | **string**| Numeric ID of task execution. | 

### Return type

[**TaskExecution**](task_execution.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskExecutionByTask

> TaskExecution ShowTaskExecutionByTask(ctx, taskId, taskExecutionId)

Displays a specific task execution.

**API Key Scope**: task_executions / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskExecutionId** | **string**| Numeric ID of task execution. | 

### Return type

[**TaskExecution**](task_execution.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

