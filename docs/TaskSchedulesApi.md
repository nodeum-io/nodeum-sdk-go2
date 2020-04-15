# \TaskSchedulesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskSchedule**](TaskSchedulesApi.md#CreateTaskSchedule) | **Post** /tasks/{task_id}/task_schedule | Creates a new task schedule. Only one should be created.
[**DestroyTaskSchedule**](TaskSchedulesApi.md#DestroyTaskSchedule) | **Delete** /tasks/{task_id}/task_schedule | Destroys the task schedule.
[**IndexTaskSchedules**](TaskSchedulesApi.md#IndexTaskSchedules) | **Get** /task_schedules | Lists all task schedules.
[**ShowTaskSchedule**](TaskSchedulesApi.md#ShowTaskSchedule) | **Get** /tasks/{task_id}/task_schedule | Displays the task schedule.
[**UpdateTaskSchedule**](TaskSchedulesApi.md#UpdateTaskSchedule) | **Put** /tasks/{task_id}/task_schedule | Updates the existing task schedule.



## CreateTaskSchedule

> TaskSchedule CreateTaskSchedule(ctx, taskId, taskScheduleBody)

Creates a new task schedule. Only one should be created.

**API Key Scope**: task_schedules / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskScheduleBody** | [**TaskSchedule**](TaskSchedule.md)|  | 

### Return type

[**TaskSchedule**](task_schedule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTaskSchedule

> DestroyTaskSchedule(ctx, taskId)

Destroys the task schedule.

**API Key Scope**: task_schedules / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 

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


## IndexTaskSchedules

> TaskScheduleCollection IndexTaskSchedules(ctx, optional)

Lists all task schedules.

**API Key Scope**: task_schedules / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTaskSchedulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskSchedulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withNext** | **optional.Bool**| Display the next scheduled date, and information about missing executions. | [default to true]
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **rrule** | **optional.String**| Filter on RRule of schedule | 
 **done** | **optional.String**| Filter on done schedule | 
 **taskId** | **optional.String**| Filter on task id | 

### Return type

[**TaskScheduleCollection**](task_schedule_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskSchedule

> TaskSchedule ShowTaskSchedule(ctx, taskId)

Displays the task schedule.

**API Key Scope**: task_schedules / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 

### Return type

[**TaskSchedule**](task_schedule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskSchedule

> TaskSchedule UpdateTaskSchedule(ctx, taskId, taskScheduleBody)

Updates the existing task schedule.

**API Key Scope**: task_schedules / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskScheduleBody** | [**TaskSchedule**](TaskSchedule.md)|  | 

### Return type

[**TaskSchedule**](task_schedule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

