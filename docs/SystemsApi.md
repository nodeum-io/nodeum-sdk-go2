# \SystemsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadResetVars**](SystemsApi.md#DownloadResetVars) | **Post** /systems/reset/generate_vars | Creates a YAML file with selected tables and downloads it
[**ResultDownloadTraces**](SystemsApi.md#ResultDownloadTraces) | **Get** /systems/download_traces | Check result of a download traces job.
[**TriggerDownloadTraces**](SystemsApi.md#TriggerDownloadTraces) | **Put** /systems/download_traces | Trigger a download traces request.



## DownloadResetVars

> *os.File DownloadResetVars(ctx, resetForm)

Creates a YAML file with selected tables and downloads it

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resetForm** | [**Reset**](Reset.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResultDownloadTraces

> *os.File ResultDownloadTraces(ctx, jobId)

Check result of a download traces job.

**API Key Scope**: systems / download_traces

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**| ID of active job | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/tar+gzip, queued, working, failed, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerDownloadTraces

> ActiveJobStatus TriggerDownloadTraces(ctx, type_)

Trigger a download traces request.

**API Key Scope**: systems / download_traces

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| Type of traces to download | 

### Return type

[**ActiveJobStatus**](active_job_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, queued, working, failed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

