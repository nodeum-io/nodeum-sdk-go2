/*
 * Nodeum API
 *
 * The Nodeum API makes it easy to tap into the digital data mesh that runs across your organisation. Make requests to our API endpoints and we’ll give you everything you need to interconnect your business workflows with your storage.  All production API requests are made to:  http://nodeumhostname/api/  The current production version of the API is v1.   **REST** The Nodeum API is a RESTful API. This means that the API is designed to allow you to get, create, update, & delete objects with the HTTP verbs GET, POST, PUT, PATCH, & DELETE.  **JSON** The Nodeum API speaks exclusively in JSON. This means that you should always set the Content-Type header to application/json to ensure that your requests are properly accepted and processed by the API.  **Authentication** All API calls require user-password authentication.   **Cross-Origin Resource Sharing** The Nodeum API supports CORS for communicating from Javascript for these endpoints. You will need to specify an Origin URI when creating your application to allow for CORS to be whitelisted for your domain.   **Pagination** Some endpoints such as File Listing return a potentially lengthy array of objects. In order to keep the response sizes manageable the API will take advantage of pagination. Pagination is a mechanism for returning a subset of the results for a request and allowing for subsequent requests to “page” through the rest of the results until the end is reached. Paginated endpoints follow a standard interface that accepts two query parameters, limit and offset, and return a payload that follows a standard form. These parameters names and their behavior are borrowed from SQL LIMIT and OFFSET keywords.  **Versioning** The Nodeum API is constantly being worked on to add features, make improvements, and fix bugs. This means that you should expect changes to be introduced and documented.   However, there are some changes or additions that are considered backwards-compatible and your applications should be flexible enough to handle them. These include:  - Adding new endpoints to the API - Adding new attributes to the response of an existing endpoint - Changing the order of attributes of responses (JSON by definition is an object of unordered key/value pairs)  **Filter parameters** When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.1.0
 * Contact: info@nodeum.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nodeumsdk
// TapeStat struct for TapeStat
type TapeStat struct {
	LogTime string `json:"log_time,omitempty"`
	Barcode string `json:"barcode,omitempty"`
	Mounts int32 `json:"mounts,omitempty"`
	DatasetsWritten int32 `json:"datasets_written,omitempty"`
	DatasetsRead int32 `json:"datasets_read,omitempty"`
	RecoveredWriteDataErrors int32 `json:"recovered_write_data_errors,omitempty"`
	UnrecoveredWriteDataErrors int32 `json:"unrecovered_write_data_errors,omitempty"`
	WriteServoErrors int32 `json:"write_servo_errors,omitempty"`
	UnrecoveredWriteServoErrors int32 `json:"unrecovered_write_servo_errors,omitempty"`
	RecoveredReadErrors int32 `json:"recovered_read_errors,omitempty"`
	UnrecoveredReadErrors int32 `json:"unrecovered_read_errors,omitempty"`
	LastMountUnrecoveredWriteErrors int32 `json:"last_mount_unrecovered_write_errors,omitempty"`
	LastMountUnrecoveredReadErrors int32 `json:"last_mount_unrecovered_read_errors,omitempty"`
	LastMountMbytesWritten int32 `json:"last_mount_mbytes_written,omitempty"`
	LastMountMbytesRead int32 `json:"last_mount_mbytes_read,omitempty"`
	LifetimeMbytesWritten int32 `json:"lifetime_mbytes_written,omitempty"`
	LifetimeMbytesRead int32 `json:"lifetime_mbytes_read,omitempty"`
	LastLoadWriteCompressionRatio int32 `json:"last_load_write_compression_ratio,omitempty"`
	LastLoadReadCompressionRatio int32 `json:"last_load_read_compression_ratio,omitempty"`
	MediumMountTime int32 `json:"medium_mount_time,omitempty"`
	MediumReadyTime int32 `json:"medium_ready_time,omitempty"`
	TotalNativeCapacity int32 `json:"total_native_capacity,omitempty"`
	TotalUsedNativeCapacity int32 `json:"total_used_native_capacity,omitempty"`
	WriteProtect int32 `json:"write_protect,omitempty"`
	Worm int32 `json:"worm,omitempty"`
	BeginningOfMediumPasses int32 `json:"beginning_of_medium_passes,omitempty"`
	MiddleOfTapePasses int32 `json:"middle_of_tape_passes,omitempty"`
	ReadCompressionRatio int32 `json:"read_compression_ratio,omitempty"`
	WriteCompressionRatio int32 `json:"write_compression_ratio,omitempty"`
	MbytesTransferredToAppClient int32 `json:"mbytes_transferred_to_app_client,omitempty"`
	BytesTransferredToAppClient int32 `json:"bytes_transferred_to_app_client,omitempty"`
	MbytesReadFromMedium int32 `json:"mbytes_read_from_medium,omitempty"`
	BytesReadFromMedium int32 `json:"bytes_read_from_medium,omitempty"`
	MbytesTransferredFromAppClient int32 `json:"mbytes_transferred_from_app_client,omitempty"`
	BytesTransferredFromAppClient int32 `json:"bytes_transferred_from_app_client,omitempty"`
	MbytesWrittenToMedium int32 `json:"mbytes_written_to_medium,omitempty"`
	BytesWrittenToMedium int32 `json:"bytes_written_to_medium,omitempty"`
	DataCompressionEnabled int32 `json:"data_compression_enabled,omitempty"`
	WriteRetries int32 `json:"write_retries,omitempty"`
	WritePerms int32 `json:"write_perms,omitempty"`
	SuspendedWrites int32 `json:"suspended_writes,omitempty"`
	FatalSuspendedWrites int32 `json:"fatal_suspended_writes,omitempty"`
	ReadRetries int32 `json:"read_retries,omitempty"`
	ReadPerms int32 `json:"read_perms,omitempty"`
	SuspendedReads int32 `json:"suspended_reads,omitempty"`
	FatalSuspendedReads int32 `json:"fatal_suspended_reads,omitempty"`
	Partition0RemainingCapacity int32 `json:"partition_0_remaining_capacity,omitempty"`
	Partition1RemainingCapacity int32 `json:"partition_1_remaining_capacity,omitempty"`
	Partition0MaximumCapacity int32 `json:"partition_0_maximum_capacity,omitempty"`
	Partition1MaximumCapacity int32 `json:"partition_1_maximum_capacity,omitempty"`
}
