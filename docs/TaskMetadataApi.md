# \TaskMetadataApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskMetadatum**](TaskMetadataApi.md#CreateTaskMetadatum) | **Post** /tasks/{task_id}/task_metadata | Creates a new task metadatum.
[**DestroyTaskMetadatum**](TaskMetadataApi.md#DestroyTaskMetadatum) | **Delete** /tasks/{task_id}/task_metadata/{task_metadatum_id} | Destroys a specific task metadatum.
[**IndexTaskMetadata**](TaskMetadataApi.md#IndexTaskMetadata) | **Get** /tasks/{task_id}/task_metadata | Lists all task metadata.
[**ShowTaskMetadatum**](TaskMetadataApi.md#ShowTaskMetadatum) | **Get** /tasks/{task_id}/task_metadata/{task_metadatum_id} | Displays a specific task metadatum.
[**UpdateTaskMetadatum**](TaskMetadataApi.md#UpdateTaskMetadatum) | **Put** /tasks/{task_id}/task_metadata/{task_metadatum_id} | Updates a specific task metadatum.



## CreateTaskMetadatum

> TaskMetadatum CreateTaskMetadatum(ctx, taskId, taskMetadatumBody)

Creates a new task metadatum.

**API Key Scope**: task_metadata / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskMetadatumBody** | [**TaskMetadatum**](TaskMetadatum.md)|  | 

### Return type

[**TaskMetadatum**](task_metadatum.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTaskMetadatum

> DestroyTaskMetadatum(ctx, taskId, taskMetadatumId)

Destroys a specific task metadatum.

**API Key Scope**: task_metadata / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskMetadatumId** | **int32**| Numeric ID of task metadatum. | 

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


## IndexTaskMetadata

> TaskMetadatumCollection IndexTaskMetadata(ctx, taskId, optional)

Lists all task metadata.

**API Key Scope**: task_metadata / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
 **optional** | ***IndexTaskMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTaskMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **key** | **optional.String**| Filter on key | 
 **value** | **optional.String**| Filter on value | 

### Return type

[**TaskMetadatumCollection**](task_metadatum_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTaskMetadatum

> TaskMetadatum ShowTaskMetadatum(ctx, taskId, taskMetadatumId)

Displays a specific task metadatum.

**API Key Scope**: task_metadata / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskMetadatumId** | **int32**| Numeric ID of task metadatum. | 

### Return type

[**TaskMetadatum**](task_metadatum.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskMetadatum

> TaskMetadatum UpdateTaskMetadatum(ctx, taskId, taskMetadatumId, taskMetadatumBody)

Updates a specific task metadatum.

**API Key Scope**: task_metadata / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskMetadatumId** | **int32**| Numeric ID of task metadatum. | 
**taskMetadatumBody** | [**TaskMetadatum**](TaskMetadatum.md)|  | 

### Return type

[**TaskMetadatum**](task_metadatum.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

