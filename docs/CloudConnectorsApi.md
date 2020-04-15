# \CloudConnectorsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudConnector**](CloudConnectorsApi.md#CreateCloudConnector) | **Post** /cloud_connectors | Creates a new cloud connector.
[**DestroyCloudConnector**](CloudConnectorsApi.md#DestroyCloudConnector) | **Delete** /cloud_connectors/{cloud_connector_id} | Destroys a specific cloud connector.
[**IndexCloudConnectors**](CloudConnectorsApi.md#IndexCloudConnectors) | **Get** /cloud_connectors | Lists all cloud connectors.
[**ShowCloudConnector**](CloudConnectorsApi.md#ShowCloudConnector) | **Get** /cloud_connectors/{cloud_connector_id} | Displays a specific cloud connector.
[**TestCloudConnector**](CloudConnectorsApi.md#TestCloudConnector) | **Put** /cloud_connectors/-/test | Test an unsaved cloud connector.
[**TestResultCloudConnector**](CloudConnectorsApi.md#TestResultCloudConnector) | **Get** /cloud_connectors/-/test | Check result of cloud connector test job.
[**UpdateCloudConnector**](CloudConnectorsApi.md#UpdateCloudConnector) | **Put** /cloud_connectors/{cloud_connector_id} | Updates a specific cloud connector.



## CreateCloudConnector

> CloudConnector CreateCloudConnector(ctx, cloudConnectorBody)

Creates a new cloud connector.

**API Key Scope**: cloud_connectors / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorBody** | [**CloudConnector**](CloudConnector.md)|  | 

### Return type

[**CloudConnector**](cloud_connector.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyCloudConnector

> DestroyCloudConnector(ctx, cloudConnectorId)

Destroys a specific cloud connector.

**API Key Scope**: cloud_connectors / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 

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


## IndexCloudConnectors

> CloudConnectorCollection IndexCloudConnectors(ctx, optional)

Lists all cloud connectors.

**API Key Scope**: cloud_connectors / index   Optional API Key Explicit Scope: cloud_connectors / get_secret_key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexCloudConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexCloudConnectorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **url** | **optional.String**| Filter on url | 
 **urlProxy** | **optional.String**| Filter on url proxy | 
 **provider** | **optional.String**| Filter on provider | 
 **region** | **optional.String**| Filter on region | 
 **accessKey** | **optional.String**| Filter on access key | 

### Return type

[**CloudConnectorCollection**](cloud_connector_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCloudConnector

> CloudConnector ShowCloudConnector(ctx, cloudConnectorId)

Displays a specific cloud connector.

**API Key Scope**: cloud_connectors / show   Optional API Key Explicit Scope: cloud_connectors / get_secret_key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 

### Return type

[**CloudConnector**](cloud_connector.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCloudConnector

> ActiveJobStatus TestCloudConnector(ctx, cloudConnectorBody)

Test an unsaved cloud connector.

**API Key Scope**: cloud_connectors / test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorBody** | [**CloudConnector**](CloudConnector.md)|  | 

### Return type

[**ActiveJobStatus**](active_job_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, queued, working, failed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestResultCloudConnector

> CloudBucketSimpleCollection TestResultCloudConnector(ctx, jobId)

Check result of cloud connector test job.

**API Key Scope**: cloud_connectors / test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**| ID of active job | 

### Return type

[**CloudBucketSimpleCollection**](cloud_bucket_simple_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, queued, working, failed, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudConnector

> CloudConnector UpdateCloudConnector(ctx, cloudConnectorId, cloudConnectorBody)

Updates a specific cloud connector.

**API Key Scope**: cloud_connectors / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
**cloudConnectorBody** | [**CloudConnector**](CloudConnector.md)|  | 

### Return type

[**CloudConnector**](cloud_connector.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

