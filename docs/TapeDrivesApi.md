# \TapeDrivesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTapeDriveByTapeLibrary**](TapeDrivesApi.md#CreateTapeDriveByTapeLibrary) | **Post** /tape_libraries/{tape_library_id}/tape_drives | Creates a new tape drive.
[**DestroyTapeDrive**](TapeDrivesApi.md#DestroyTapeDrive) | **Delete** /tape_drives/{tape_drive_id} | Destroys a specific tape drive.
[**DestroyTapeDriveByTapeLibrary**](TapeDrivesApi.md#DestroyTapeDriveByTapeLibrary) | **Delete** /tape_libraries/{tape_library_id}/tape_drives/{tape_drive_id} | Destroys a specific tape drive.
[**IndexTapeDriveDevices**](TapeDrivesApi.md#IndexTapeDriveDevices) | **Get** /tape_libraries/{tape_library_id}/tape_drives/-/devices | Lists tape drives devices.
[**IndexTapeDrives**](TapeDrivesApi.md#IndexTapeDrives) | **Get** /tape_drives | Lists all tape drives.
[**IndexTapeDrivesByTapeLibrary**](TapeDrivesApi.md#IndexTapeDrivesByTapeLibrary) | **Get** /tape_libraries/{tape_library_id}/tape_drives | Lists all tape drives.
[**ShowTapeDrive**](TapeDrivesApi.md#ShowTapeDrive) | **Get** /tape_drives/{tape_drive_id} | Displays a specific tape drive.
[**ShowTapeDriveByTapeLibrary**](TapeDrivesApi.md#ShowTapeDriveByTapeLibrary) | **Get** /tape_libraries/{tape_library_id}/tape_drives/{tape_drive_id} | Displays a specific tape drive.
[**UpdateTapeDrive**](TapeDrivesApi.md#UpdateTapeDrive) | **Put** /tape_drives/{tape_drive_id} | Updates a specific tape drive.
[**UpdateTapeDriveByTapeLibrary**](TapeDrivesApi.md#UpdateTapeDriveByTapeLibrary) | **Put** /tape_libraries/{tape_library_id}/tape_drives/{tape_drive_id} | Updates a specific tape drive.



## CreateTapeDriveByTapeLibrary

> TapeDrive CreateTapeDriveByTapeLibrary(ctx, tapeLibraryId, tapeDriveBody)

Creates a new tape drive.

**API Key Scope**: tape_drives / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
**tapeDriveBody** | [**TapeDrive**](TapeDrive.md)|  | 

### Return type

[**TapeDrive**](tape_drive.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTapeDrive

> DestroyTapeDrive(ctx, tapeDriveId)

Destroys a specific tape drive.

**API Key Scope**: tape_drives / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeDriveId** | **string**| Numeric ID, serial, or name of tape drive. | 

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


## DestroyTapeDriveByTapeLibrary

> DestroyTapeDriveByTapeLibrary(ctx, tapeLibraryId, tapeDriveId)

Destroys a specific tape drive.

**API Key Scope**: tape_drives / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
**tapeDriveId** | **string**| Numeric ID, serial, or name of tape drive. | 

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


## IndexTapeDriveDevices

> TapeDriveDeviceCollection IndexTapeDriveDevices(ctx, tapeLibraryId, jobId)

Lists tape drives devices.

**API Key Scope**: tape_drives / devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
**jobId** | **string**| ID of active job | 

### Return type

[**TapeDriveDeviceCollection**](tape_drive_device_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, queued, working, failed, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexTapeDrives

> TapeDriveCollection IndexTapeDrives(ctx, optional)

Lists all tape drives.

**API Key Scope**: tape_drives / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTapeDrivesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTapeDrivesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **tapeLibraryId** | **optional.String**| Filter on tape library id | 
 **name** | **optional.String**| Filter on name | 
 **serial** | **optional.String**| Filter on serial | 
 **comment** | **optional.String**| Filter on comment | 
 **scsiAddress** | **optional.String**| Filter on scsi address | 
 **vendor** | **optional.String**| Filter on vendor | 
 **product** | **optional.String**| Filter on product | 
 **firmware** | **optional.String**| Filter on firmware | 
 **device** | **optional.String**| Filter on device | 
 **sgdevice** | **optional.String**| Filter on sgdevice | 
 **libso** | **optional.String**| Filter on libso | 
 **acs** | **optional.String**| Filter on acs | 
 **lsm** | **optional.String**| Filter on lsm | 
 **panel** | **optional.String**| Filter on panel | 
 **transport** | **optional.String**| Filter on transport | 
 **status** | **optional.String**| Filter on status | 
 **full** | **optional.String**| Filter on full | 
 **mountCount** | **optional.String**| Filter on mount count | 
 **useTo** | **optional.String**| Filter on use to | 
 **useBy** | **optional.String**| Filter on use by | 
 **barcode** | **optional.String**| Filter on barcode | 
 **taskId** | **optional.String**| Filter on task id | 
 **useFileProcessedSize** | **optional.String**| Filter on use file processed size | 
 **useFileSizeToProcess** | **optional.String**| Filter on use file size to process | 
 **useFileNameProcessed** | **optional.String**| Filter on use file name processed | 
 **bandwidth** | **optional.String**| Filter on bandwidth | 

### Return type

[**TapeDriveCollection**](tape_drive_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexTapeDrivesByTapeLibrary

> TapeDriveCollection IndexTapeDrivesByTapeLibrary(ctx, tapeLibraryId, optional)

Lists all tape drives.

**API Key Scope**: tape_drives / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
 **optional** | ***IndexTapeDrivesByTapeLibraryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTapeDrivesByTapeLibraryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **serial** | **optional.String**| Filter on serial | 
 **comment** | **optional.String**| Filter on comment | 
 **scsiAddress** | **optional.String**| Filter on scsi address | 
 **vendor** | **optional.String**| Filter on vendor | 
 **product** | **optional.String**| Filter on product | 
 **firmware** | **optional.String**| Filter on firmware | 
 **device** | **optional.String**| Filter on device | 
 **sgdevice** | **optional.String**| Filter on sgdevice | 
 **libso** | **optional.String**| Filter on libso | 
 **acs** | **optional.String**| Filter on acs | 
 **lsm** | **optional.String**| Filter on lsm | 
 **panel** | **optional.String**| Filter on panel | 
 **transport** | **optional.String**| Filter on transport | 
 **status** | **optional.String**| Filter on status | 
 **full** | **optional.String**| Filter on full | 
 **mountCount** | **optional.String**| Filter on mount count | 
 **useTo** | **optional.String**| Filter on use to | 
 **useBy** | **optional.String**| Filter on use by | 
 **barcode** | **optional.String**| Filter on barcode | 
 **taskId** | **optional.String**| Filter on task id | 
 **useFileProcessedSize** | **optional.String**| Filter on use file processed size | 
 **useFileSizeToProcess** | **optional.String**| Filter on use file size to process | 
 **useFileNameProcessed** | **optional.String**| Filter on use file name processed | 
 **bandwidth** | **optional.String**| Filter on bandwidth | 

### Return type

[**TapeDriveCollection**](tape_drive_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTapeDrive

> TapeDrive ShowTapeDrive(ctx, tapeDriveId)

Displays a specific tape drive.

**API Key Scope**: tape_drives / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeDriveId** | **string**| Numeric ID, serial, or name of tape drive. | 

### Return type

[**TapeDrive**](tape_drive.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTapeDriveByTapeLibrary

> TapeDrive ShowTapeDriveByTapeLibrary(ctx, tapeLibraryId, tapeDriveId)

Displays a specific tape drive.

**API Key Scope**: tape_drives / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
**tapeDriveId** | **string**| Numeric ID, serial, or name of tape drive. | 

### Return type

[**TapeDrive**](tape_drive.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTapeDrive

> TapeDrive UpdateTapeDrive(ctx, tapeDriveId, tapeDriveBody)

Updates a specific tape drive.

**API Key Scope**: tape_drives / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeDriveId** | **string**| Numeric ID, serial, or name of tape drive. | 
**tapeDriveBody** | [**TapeDrive**](TapeDrive.md)|  | 

### Return type

[**TapeDrive**](tape_drive.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTapeDriveByTapeLibrary

> TapeDrive UpdateTapeDriveByTapeLibrary(ctx, tapeLibraryId, tapeDriveId, tapeDriveBody)

Updates a specific tape drive.

**API Key Scope**: tape_drives / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
**tapeDriveId** | **string**| Numeric ID, serial, or name of tape drive. | 
**tapeDriveBody** | [**TapeDrive**](TapeDrive.md)|  | 

### Return type

[**TapeDrive**](tape_drive.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

