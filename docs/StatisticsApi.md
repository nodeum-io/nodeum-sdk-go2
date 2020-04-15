# \StatisticsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticsByFileExtension**](StatisticsApi.md#StatisticsByFileExtension) | **Get** /statistics/by_file_extension | TODO
[**StatisticsByGroupOwner**](StatisticsApi.md#StatisticsByGroupOwner) | **Get** /statistics/by_group_owner | TODO
[**StatisticsByPrimaryName**](StatisticsApi.md#StatisticsByPrimaryName) | **Get** /statistics/by_primary_name | TODO
[**StatisticsBySecondaryStorage**](StatisticsApi.md#StatisticsBySecondaryStorage) | **Get** /statistics/by_secondary_storage | TODO
[**StatisticsBySize**](StatisticsApi.md#StatisticsBySize) | **Get** /statistics/by_size | TODO
[**StatisticsByUserOwner**](StatisticsApi.md#StatisticsByUserOwner) | **Get** /statistics/by_user_owner | TODO



## StatisticsByFileExtension

> ByFileExtensionFacet StatisticsByFileExtension(ctx, optional)

TODO

**API Key Scope**: statistics / by_file_extension

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByFileExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StatisticsByFileExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**ByFileExtensionFacet**](by_file_extension_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsByGroupOwner

> ByGroupOwnerFacet StatisticsByGroupOwner(ctx, optional)

TODO

**API Key Scope**: statistics / by_group_owner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByGroupOwnerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StatisticsByGroupOwnerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**ByGroupOwnerFacet**](by_group_owner_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsByPrimaryName

> ByPrimaryFacet StatisticsByPrimaryName(ctx, optional)

TODO

**API Key Scope**: statistics / by_primary_name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByPrimaryNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StatisticsByPrimaryNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**ByPrimaryFacet**](by_primary_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsBySecondaryStorage

> BySecondaryFacet StatisticsBySecondaryStorage(ctx, optional)

TODO

**API Key Scope**: statistics / by_secondary_storage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsBySecondaryStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StatisticsBySecondaryStorageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**BySecondaryFacet**](by_secondary_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsBySize

> BySizeFacet StatisticsBySize(ctx, optional)

TODO

**API Key Scope**: statistics / by_size

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsBySizeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StatisticsBySizeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**BySizeFacet**](by_size_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticsByUserOwner

> ByUserOwnerFacet StatisticsByUserOwner(ctx, optional)

TODO

**API Key Scope**: statistics / by_user_owner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByUserOwnerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StatisticsByUserOwnerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**ByUserOwnerFacet**](by_user_owner_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

