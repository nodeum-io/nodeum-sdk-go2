# \TapeLibrariesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTapeLibrary**](TapeLibrariesApi.md#CreateTapeLibrary) | **Post** /tape_libraries | Creates a new tape library.
[**DestroyTapeLibrary**](TapeLibrariesApi.md#DestroyTapeLibrary) | **Delete** /tape_libraries/{tape_library_id} | Destroys a specific tape library.
[**IndexTapeLibraries**](TapeLibrariesApi.md#IndexTapeLibraries) | **Get** /tape_libraries | Lists all tape libraries.
[**IndexTapeLibraryDevices**](TapeLibrariesApi.md#IndexTapeLibraryDevices) | **Get** /tape_libraries/-/devices | Lists tape libraries devices.
[**ShowTapeLibrary**](TapeLibrariesApi.md#ShowTapeLibrary) | **Get** /tape_libraries/{tape_library_id} | Displays a specific tape library.
[**UpdateTapeLibrary**](TapeLibrariesApi.md#UpdateTapeLibrary) | **Put** /tape_libraries/{tape_library_id} | Updates a specific tape library.



## CreateTapeLibrary

> TapeLibrary CreateTapeLibrary(ctx, tapeLibraryBody)

Creates a new tape library.

**API Key Scope**: tape_libraries / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryBody** | [**TapeLibrary**](TapeLibrary.md)|  | 

### Return type

[**TapeLibrary**](tape_library.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTapeLibrary

> DestroyTapeLibrary(ctx, tapeLibraryId)

Destroys a specific tape library.

**API Key Scope**: tape_libraries / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 

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


## IndexTapeLibraries

> TapeLibraryCollection IndexTapeLibraries(ctx, optional)

Lists all tape libraries.

**API Key Scope**: tape_libraries / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTapeLibrariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTapeLibrariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **serial** | **optional.String**| Filter on serial | 
 **comment** | **optional.String**| Filter on comment | 
 **protocol** | **optional.String**| Filter on protocol | 
 **vendor** | **optional.String**| Filter on vendor | 
 **product** | **optional.String**| Filter on product | 
 **firmware** | **optional.String**| Filter on firmware | 
 **device** | **optional.String**| Filter on device | 
 **libso** | **optional.String**| Filter on libso | 
 **acs** | **optional.String**| Filter on acs | 
 **status** | **optional.String**| Filter on status | 
 **storageSlots** | **optional.String**| Filter on storage slots | 
 **storageSlotsAddress** | **optional.String**| Filter on storage slots address | 
 **ioSlots** | **optional.String**| Filter on io slots | 
 **ioSlotsAddress** | **optional.String**| Filter on io slots address | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**TapeLibraryCollection**](tape_library_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexTapeLibraryDevices

> TapeLibraryDeviceCollection IndexTapeLibraryDevices(ctx, jobId)

Lists tape libraries devices.

**API Key Scope**: tape_libraries / devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**| ID of active job | 

### Return type

[**TapeLibraryDeviceCollection**](tape_library_device_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, queued, working, failed, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTapeLibrary

> TapeLibrary ShowTapeLibrary(ctx, tapeLibraryId)

Displays a specific tape library.

**API Key Scope**: tape_libraries / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 

### Return type

[**TapeLibrary**](tape_library.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTapeLibrary

> TapeLibrary UpdateTapeLibrary(ctx, tapeLibraryId, tapeLibraryBody)

Updates a specific tape library.

**API Key Scope**: tape_libraries / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
**tapeLibraryBody** | [**TapeLibrary**](TapeLibrary.md)|  | 

### Return type

[**TapeLibrary**](tape_library.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

