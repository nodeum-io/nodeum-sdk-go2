# \TaskDestinationsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskDestination**](TaskDestinationsApi.md#CreateTaskDestination) | **Post** /tasks/{task_id}/task_destinations | Creates a new task destination.
[**DestroyTaskDestination**](TaskDestinationsApi.md#DestroyTaskDestination) | **Delete** /tasks/{task_id}/task_destinations/{task_destination_id} | Destroys a specific task destination.
[**IndexTaskDestinations**](TaskDestinationsApi.md#IndexTaskDestinations) | **Get** /tasks/{task_id}/task_destinations | Lists all task destinations.
[**ShowTaskDestination**](TaskDestinationsApi.md#ShowTaskDestination) | **Get** /tasks/{task_id}/task_destinations/{task_destination_id} | Displays a specific task destination.
[**UpdateTaskDestination**](TaskDestinationsApi.md#UpdateTaskDestination) | **Put** /tasks/{task_id}/task_destinations/{task_destination_id} | Updates a specific task destination.



## CreateTaskDestination

> TaskDestinationDown CreateTaskDestination(ctx, taskId, taskDestinationBody)

Creates a new task destination.

**API Key Scope**: task_destinations / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskDestinationBody** | [**TaskDestinationUp**](TaskDestinationUp.md)|  | 

### Return type

[**TaskDestinationDown**](task_destination_down.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTaskDestination

> DestroyTaskDestination(ctx, taskId, taskDestinationId)

Destroys a specific task destination.

**API Key Scope**: task_destinations / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskDestinationId** | **int32**| Numeric ID of task destination. | 

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


## IndexTaskDestinations

> TaskDestinationCollection IndexTaskDestinations(ctx, taskId, optional)

Lists all task destinations.

**API Key Scope**: task_destinations / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
 **optional** | ***IndexTaskDestinationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskDestinationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **folderId** | **optional.String**| Filter on folder id | 
 **tapeId** | **optional.String**| Filter on tape id | 
 **poolId** | **optional.String**| Filter on a pool id | 

### Return type

[**TaskDestinationCollection**](task_destination_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskDestination

> TaskDestinationDown ShowTaskDestination(ctx, taskId, taskDestinationId)

Displays a specific task destination.

**API Key Scope**: task_destinations / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskDestinationId** | **int32**| Numeric ID of task destination. | 

### Return type

[**TaskDestinationDown**](task_destination_down.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskDestination

> TaskDestinationDown UpdateTaskDestination(ctx, taskId, taskDestinationId, taskDestinationBody)

Updates a specific task destination.

**API Key Scope**: task_destinations / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskDestinationId** | **int32**| Numeric ID of task destination. | 
**taskDestinationBody** | [**TaskDestinationUp**](TaskDestinationUp.md)|  | 

### Return type

[**TaskDestinationDown**](task_destination_down.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

