# \NasSharesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNasShareByNas**](NasSharesApi.md#CreateNasShareByNas) | **Post** /nas/{nas_id}/nas_shares | Creates a new NAS share.
[**DestroyNasShare**](NasSharesApi.md#DestroyNasShare) | **Delete** /nas_shares/{nas_share_id} | Destroys a specific NAS share.
[**DestroyNasShareByNas**](NasSharesApi.md#DestroyNasShareByNas) | **Delete** /nas/{nas_id}/nas_shares/{nas_share_id} | Destroys a specific NAS share.
[**DestroyNasShareByPool**](NasSharesApi.md#DestroyNasShareByPool) | **Delete** /pools/{pool_id}/nas_shares/{nas_share_id} | Destroys a specific NAS share.
[**IndexNasShares**](NasSharesApi.md#IndexNasShares) | **Get** /nas_shares | Lists all NAS shares.
[**IndexNasSharesByNas**](NasSharesApi.md#IndexNasSharesByNas) | **Get** /nas/{nas_id}/nas_shares | Lists all NAS shares.
[**IndexNasSharesByPool**](NasSharesApi.md#IndexNasSharesByPool) | **Get** /pools/{pool_id}/nas_shares | Lists all NAS shares from pool.
[**MountStatusNasShare**](NasSharesApi.md#MountStatusNasShare) | **Get** /nas_shares/{nas_share_id}/mount | Get mount status of NAS Share.
[**MountStatusNasShareByNas**](NasSharesApi.md#MountStatusNasShareByNas) | **Get** /nas/{nas_id}/nas_shares/{nas_share_id}/mount | Get mount status of NAS Share.
[**MountStatusNasShareByPool**](NasSharesApi.md#MountStatusNasShareByPool) | **Get** /pools/{pool_id}/nas_shares/{nas_share_id}/mount | Get mount status of NAS Share.
[**ShowNasShare**](NasSharesApi.md#ShowNasShare) | **Get** /nas_shares/{nas_share_id} | Displays a specific NAS share.
[**ShowNasShareByNas**](NasSharesApi.md#ShowNasShareByNas) | **Get** /nas/{nas_id}/nas_shares/{nas_share_id} | Displays a specific NAS share.
[**ShowNasShareByPool**](NasSharesApi.md#ShowNasShareByPool) | **Get** /pools/{pool_id}/nas_shares/{nas_share_id} | Displays a specific NAS share.
[**TestNasShare**](NasSharesApi.md#TestNasShare) | **Put** /nas/{nas_id}/nas_shares/-/test | Test an unsaved NAS Share.
[**TestResultNasShare**](NasSharesApi.md#TestResultNasShare) | **Get** /nas/{nas_id}/nas_shares/-/test | Check result of a NAS Share test job.
[**UpdateNasShare**](NasSharesApi.md#UpdateNasShare) | **Put** /nas_shares/{nas_share_id} | Updates a specific NAS share.
[**UpdateNasShareByNas**](NasSharesApi.md#UpdateNasShareByNas) | **Put** /nas/{nas_id}/nas_shares/{nas_share_id} | Updates a specific NAS share.
[**UpdateNasShareByPool**](NasSharesApi.md#UpdateNasShareByPool) | **Put** /pools/{pool_id}/nas_shares/{nas_share_id} | Updates a specific NAS share.



## CreateNasShareByNas

> NasShare CreateNasShareByNas(ctx, nasId, nasShareBody)

Creates a new NAS share.

**API Key Scope**: nas_shares / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyNasShare

> DestroyNasShare(ctx, nasShareId)

Destroys a specific NAS share.

**API Key Scope**: nas_shares / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

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


## DestroyNasShareByNas

> DestroyNasShareByNas(ctx, nasId, nasShareId)

Destroys a specific NAS share.

**API Key Scope**: nas_shares / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

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


## DestroyNasShareByPool

> DestroyNasShareByPool(ctx, poolId, nasShareId)

Destroys a specific NAS share.

**API Key Scope**: nas_shares / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

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


## IndexNasShares

> NasShareCollection IndexNasShares(ctx, optional)

Lists all NAS shares.

**API Key Scope**: nas_shares / index   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexNasSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexNasSharesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **path** | **optional.String**| Filter on path | 
 **options** | **optional.String**| Filter on options | 
 **username** | **optional.String**| Filter on username | 
 **nasId** | **optional.String**| Filter on NAS id | 
 **poolId** | **optional.String**| Filter on a pool id | 

### Return type

[**NasShareCollection**](nas_share_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexNasSharesByNas

> NasShareCollection IndexNasSharesByNas(ctx, nasId, optional)

Lists all NAS shares.

**API Key Scope**: nas_shares / index   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
 **optional** | ***IndexNasSharesByNasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexNasSharesByNasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **path** | **optional.String**| Filter on path | 
 **options** | **optional.String**| Filter on options | 
 **username** | **optional.String**| Filter on username | 
 **poolId** | **optional.String**| Filter on a pool id | 

### Return type

[**NasShareCollection**](nas_share_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexNasSharesByPool

> NasShareCollection IndexNasSharesByPool(ctx, poolId, optional)

Lists all NAS shares from pool.

**API Key Scope**: nas_shares / index   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
 **optional** | ***IndexNasSharesByPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexNasSharesByPoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **path** | **optional.String**| Filter on path | 
 **options** | **optional.String**| Filter on options | 
 **username** | **optional.String**| Filter on username | 
 **nasId** | **optional.String**| Filter on NAS id | 

### Return type

[**NasShareCollection**](nas_share_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountStatusNasShare

> MountStatus MountStatusNasShare(ctx, nasShareId)

Get mount status of NAS Share.

**API Key Scope**: nas_shares / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

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


## MountStatusNasShareByNas

> MountStatus MountStatusNasShareByNas(ctx, nasId, nasShareId)

Get mount status of NAS Share.

**API Key Scope**: nas_shares / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

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


## MountStatusNasShareByPool

> MountStatus MountStatusNasShareByPool(ctx, poolId, nasShareId)

Get mount status of NAS Share.

**API Key Scope**: nas_shares / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

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


## ShowNasShare

> NasShare ShowNasShare(ctx, nasShareId)

Displays a specific NAS share.

**API Key Scope**: nas_shares / show   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowNasShareByNas

> NasShare ShowNasShareByNas(ctx, nasId, nasShareId)

Displays a specific NAS share.

**API Key Scope**: nas_shares / show   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowNasShareByPool

> NasShare ShowNasShareByPool(ctx, poolId, nasShareId)

Displays a specific NAS share.

**API Key Scope**: nas_shares / show   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestNasShare

> ActiveJobStatus TestNasShare(ctx, nasId, nasShareBody)

Test an unsaved NAS Share.

**API Key Scope**: nas_shares / test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

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


## TestResultNasShare

> ActiveJobStatus TestResultNasShare(ctx, nasId, jobId)

Check result of a NAS Share test job.

**API Key Scope**: nas_shares / test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**jobId** | **string**| ID of active job | 

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


## UpdateNasShare

> NasShare UpdateNasShare(ctx, nasShareId, nasShareBody)

Updates a specific NAS share.

**API Key Scope**: nas_shares / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **string**| Numeric ID or name of NAS share. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNasShareByNas

> NasShare UpdateNasShareByNas(ctx, nasId, nasShareId, nasShareBody)

Updates a specific NAS share.

**API Key Scope**: nas_shares / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNasShareByPool

> NasShare UpdateNasShareByPool(ctx, poolId, nasShareId, nasShareBody)

Updates a specific NAS share.

**API Key Scope**: nas_shares / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Numeric ID, or name of pool. | 
**nasShareId** | **string**| Numeric ID or name of NAS share. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

