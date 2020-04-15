# \CloudBucketsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndexCloudBuckets**](CloudBucketsApi.md#IndexCloudBuckets) | **Get** /cloud_buckets | Lists all cloud buckets.
[**IndexCloudBucketsByCloudConnector**](CloudBucketsApi.md#IndexCloudBucketsByCloudConnector) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets | Lists all cloud buckets.
[**IndexCloudBucketsByPool**](CloudBucketsApi.md#IndexCloudBucketsByPool) | **Get** /pools/{pool_id}/cloud_buckets | Lists all cloud buckets from pool.
[**MountStatusCloudBucket**](CloudBucketsApi.md#MountStatusCloudBucket) | **Get** /cloud_buckets/{cloud_bucket_id}/mount | Get mount status of Cloud bucket.
[**MountStatusCloudBucketByCloudConnector**](CloudBucketsApi.md#MountStatusCloudBucketByCloudConnector) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets/{cloud_bucket_id}/mount | Get mount status of Cloud bucket.
[**MountStatusCloudBucketByPool**](CloudBucketsApi.md#MountStatusCloudBucketByPool) | **Get** /pools/{pool_id}/cloud_buckets/{cloud_bucket_id}/mount | Get mount status of Cloud bucket.
[**ShowCloudBucket**](CloudBucketsApi.md#ShowCloudBucket) | **Get** /cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
[**ShowCloudBucketByCloudConnector**](CloudBucketsApi.md#ShowCloudBucketByCloudConnector) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
[**ShowCloudBucketByPool**](CloudBucketsApi.md#ShowCloudBucketByPool) | **Get** /pools/{pool_id}/cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
[**SyncCloudBuckets**](CloudBucketsApi.md#SyncCloudBuckets) | **Put** /cloud_connectors/{cloud_connector_id}/cloud_buckets/-/sync | Synchronize internal cloud buckets with their remote equivalent.
[**SyncResultCloudBuckets**](CloudBucketsApi.md#SyncResultCloudBuckets) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets/-/sync | Check result of cloud connector sync job.
[**UpdateCloudBucket**](CloudBucketsApi.md#UpdateCloudBucket) | **Put** /cloud_buckets/{cloud_bucket_id} | Updates a specific cloud bucket.
[**UpdateCloudBucketByCloudConnector**](CloudBucketsApi.md#UpdateCloudBucketByCloudConnector) | **Put** /cloud_connectors/{cloud_connector_id}/cloud_buckets/{cloud_bucket_id} | Updates a specific cloud bucket.
[**UpdateCloudBucketByPool**](CloudBucketsApi.md#UpdateCloudBucketByPool) | **Put** /pools/{pool_id}/cloud_buckets/{cloud_bucket_id} | Updates a specific cloud bucket.



## IndexCloudBuckets

> CloudBucketCollection IndexCloudBuckets(ctx, optional)

Lists all cloud buckets.

**API Key Scope**: cloud_buckets / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexCloudBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexCloudBucketsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **cloudConnectorId** | **optional.String**| Filter on cloud connector id | 
 **poolId** | **optional.String**| Filter on a pool id | 
 **name** | **optional.String**| Filter on name | 
 **location** | **optional.String**| Filter on location | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**CloudBucketCollection**](cloud_bucket_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexCloudBucketsByCloudConnector

> CloudBucketCollection IndexCloudBucketsByCloudConnector(ctx, cloudConnectorId, optional)

Lists all cloud buckets.

**API Key Scope**: cloud_buckets / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
 **optional** | ***IndexCloudBucketsByCloudConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexCloudBucketsByCloudConnectorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **poolId** | **optional.String**| Filter on a pool id | 
 **name** | **optional.String**| Filter on name | 
 **location** | **optional.String**| Filter on location | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**CloudBucketCollection**](cloud_bucket_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexCloudBucketsByPool

> CloudBucketCollection IndexCloudBucketsByPool(ctx, poolId, optional)

Lists all cloud buckets from pool.

**API Key Scope**: cloud_buckets / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
 **optional** | ***IndexCloudBucketsByPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexCloudBucketsByPoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **cloudConnectorId** | **optional.String**| Filter on cloud connector id | 
 **name** | **optional.String**| Filter on name | 
 **location** | **optional.String**| Filter on location | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**CloudBucketCollection**](cloud_bucket_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountStatusCloudBucket

> MountStatus MountStatusCloudBucket(ctx, cloudBucketId)

Get mount status of Cloud bucket.

**API Key Scope**: cloud_buckets / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountStatusCloudBucketByCloudConnector

> MountStatus MountStatusCloudBucketByCloudConnector(ctx, cloudConnectorId, cloudBucketId)

Get mount status of Cloud bucket.

**API Key Scope**: cloud_buckets / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountStatusCloudBucketByPool

> MountStatus MountStatusCloudBucketByPool(ctx, poolId, cloudBucketId)

Get mount status of Cloud bucket.

**API Key Scope**: cloud_buckets / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCloudBucket

> CloudBucket ShowCloudBucket(ctx, cloudBucketId)

Displays a specific cloud bucket.

**API Key Scope**: cloud_buckets / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCloudBucketByCloudConnector

> CloudBucket ShowCloudBucketByCloudConnector(ctx, cloudConnectorId, cloudBucketId)

Displays a specific cloud bucket.

**API Key Scope**: cloud_buckets / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCloudBucketByPool

> CloudBucket ShowCloudBucketByPool(ctx, poolId, cloudBucketId)

Displays a specific cloud bucket.

**API Key Scope**: cloud_buckets / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncCloudBuckets

> ActiveJobStatus SyncCloudBuckets(ctx, cloudConnectorId)

Synchronize internal cloud buckets with their remote equivalent.

**API Key Scope**: cloud_buckets / sync

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 

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


## SyncResultCloudBuckets

> CloudBucketSimpleCollection SyncResultCloudBuckets(ctx, cloudConnectorId, jobId)

Check result of cloud connector sync job.

**API Key Scope**: cloud_buckets / sync

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
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


## UpdateCloudBucket

> CloudBucket UpdateCloudBucket(ctx, cloudBucketId, cloudBucketBody)

Updates a specific cloud bucket.

**API Key Scope**: cloud_buckets / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 
**cloudBucketBody** | [**CloudBucket**](CloudBucket.md)|  | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudBucketByCloudConnector

> CloudBucket UpdateCloudBucketByCloudConnector(ctx, cloudConnectorId, cloudBucketId, cloudBucketBody)

Updates a specific cloud bucket.

**API Key Scope**: cloud_buckets / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 
**cloudBucketBody** | [**CloudBucket**](CloudBucket.md)|  | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudBucketByPool

> CloudBucket UpdateCloudBucketByPool(ctx, poolId, cloudBucketId, cloudBucketBody)

Updates a specific cloud bucket.

**API Key Scope**: cloud_buckets / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 
**cloudBucketBody** | [**CloudBucket**](CloudBucket.md)|  | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

