# \PoolsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePool**](PoolsApi.md#CreatePool) | **Post** /pools | Creates a new pool.
[**CreatePrimaryScan**](PoolsApi.md#CreatePrimaryScan) | **Post** /pools/{pool_id}/primary_scan | Set a new primary pool scan option.
[**DestroyPool**](PoolsApi.md#DestroyPool) | **Delete** /pools/{pool_id} | Destroys a specific tape pool.
[**DestroyPrimaryScan**](PoolsApi.md#DestroyPrimaryScan) | **Delete** /pools/{pool_id}/primary_scan | Disable the primary pool scan.
[**IndexPools**](PoolsApi.md#IndexPools) | **Get** /pools | Lists all pools.
[**MountPool**](PoolsApi.md#MountPool) | **Put** /pools/{pool_id}/mount | Mount Pool.
[**MountStatusPool**](PoolsApi.md#MountStatusPool) | **Get** /pools/{pool_id}/mount | Get mount status of Pool.
[**ShowPool**](PoolsApi.md#ShowPool) | **Get** /pools/{pool_id} | Displays a specific pool.
[**ShowPrimaryScan**](PoolsApi.md#ShowPrimaryScan) | **Get** /pools/{pool_id}/primary_scan | Displays the primary pool scan status.
[**SyncPrimaryPool**](PoolsApi.md#SyncPrimaryPool) | **Post** /pools/{pool_id}/sync | Synchronize a primary after a scan (for internal use only).
[**UnmountPool**](PoolsApi.md#UnmountPool) | **Delete** /pools/{pool_id}/mount | Unmount Pool.
[**UpdatePool**](PoolsApi.md#UpdatePool) | **Put** /pools/{pool_id} | Updates a specific pool.
[**UpdatePrimaryScan**](PoolsApi.md#UpdatePrimaryScan) | **Put** /pools/{pool_id}/primary_scan | Updates the existing primary pool scan option.



## CreatePool

> Pool CreatePool(ctx, poolBody)

Creates a new pool.

**API Key Scope**: pools / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolBody** | [**PoolUp**](PoolUp.md)|  | 

### Return type

[**Pool**](pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrimaryScan

> PrimaryScan CreatePrimaryScan(ctx, poolId, primaryScanBody)

Set a new primary pool scan option.

**API Key Scope**: primary_scans / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**primaryScanBody** | [**PrimaryScan**](PrimaryScan.md)|  | 

### Return type

[**PrimaryScan**](primary_scan.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyPool

> DestroyPool(ctx, poolId)

Destroys a specific tape pool.

**API Key Scope**: pools / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 

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


## DestroyPrimaryScan

> DestroyPrimaryScan(ctx, poolId)

Disable the primary pool scan.

**API Key Scope**: primary_scans / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 

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


## IndexPools

> PoolCollection IndexPools(ctx, optional)

Lists all pools.

**API Key Scope**: pools / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexPoolsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **comment** | **optional.String**| Filter on comment | 
 **type_** | **optional.String**| Filter on type | 
 **content** | **optional.String**| Filter on content | 
 **primaryId** | **optional.String**| Filter on primary id | 

### Return type

[**PoolCollection**](pool_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountPool

> MountStatus MountPool(ctx, poolId)

Mount Pool.

**API Key Scope**: pools / mount

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 

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


## MountStatusPool

> MountStatus MountStatusPool(ctx, poolId)

Get mount status of Pool.

**API Key Scope**: pools / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 

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


## ShowPool

> Pool ShowPool(ctx, poolId)

Displays a specific pool.

**API Key Scope**: pools / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 

### Return type

[**Pool**](pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPrimaryScan

> PrimaryScan ShowPrimaryScan(ctx, poolId)

Displays the primary pool scan status.

**API Key Scope**: primary_scans / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 

### Return type

[**PrimaryScan**](primary_scan.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPrimaryPool

> SyncPrimaryPool(ctx, poolId, tx)

Synchronize a primary after a scan (for internal use only).

**API Key Scope**: pools / sync_primary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**tx** | **int32**| New transaction number. | 

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


## UnmountPool

> MountStatus UnmountPool(ctx, poolId)

Unmount Pool.

**API Key Scope**: pools / unmount

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 

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


## UpdatePool

> Pool UpdatePool(ctx, poolId, poolBody)

Updates a specific pool.

**API Key Scope**: pools / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**poolBody** | [**PoolUp**](PoolUp.md)|  | 

### Return type

[**Pool**](pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrimaryScan

> PrimaryScan UpdatePrimaryScan(ctx, poolId, primaryScanBody)

Updates the existing primary pool scan option.

**API Key Scope**: primary_scans / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**primaryScanBody** | [**PrimaryScan**](PrimaryScan.md)|  | 

### Return type

[**PrimaryScan**](primary_scan.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

