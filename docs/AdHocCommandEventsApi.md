# {{classname}}

All URIs are relative to *https://null/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdHocCommandEventsAdHocCommandEventsList**](AdHocCommandEventsApi.md#AdHocCommandEventsAdHocCommandEventsList) | **Get** /api/v2/ad_hoc_command_events/ |  List Ad Hoc Command Events
[**AdHocCommandEventsAdHocCommandEventsRead**](AdHocCommandEventsApi.md#AdHocCommandEventsAdHocCommandEventsRead) | **Get** /api/v2/ad_hoc_command_events/{id}/ |  Retrieve an Ad Hoc Command Event

# **AdHocCommandEventsAdHocCommandEventsList**
> AdHocCommandEventsAdHocCommandEventsList(ctx, optional)
 List Ad Hoc Command Events

 Make a GET request to this resource to retrieve the list of ad hoc command events.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of ad hoc command events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more ad hoc command event records.    ## Results  Each ad hoc command event data structure includes the following fields:  * `id`: Database ID for this ad hoc command event. (integer) * `type`: Data type for this ad hoc command event. (choice) * `url`: URL for this ad hoc command event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command event was created. (datetime) * `modified`: Timestamp when this ad hoc command event was last modified. (datetime) * `ad_hoc_command`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_ok`: Host OK     - `runner_on_unreachable`: Host Unreachable     - `runner_on_skipped`: Host Skipped     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)    ## Sorting  To specify that ad hoc command events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdHocCommandEventsApiAdHocCommandEventsAdHocCommandEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandEventsApiAdHocCommandEventsAdHocCommandEventsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdHocCommandEventsAdHocCommandEventsRead**
> AdHocCommandEventsAdHocCommandEventsRead(ctx, id, optional)
 Retrieve an Ad Hoc Command Event

 Make GET request to this resource to retrieve a single ad hoc command event record containing the following fields:  * `id`: Database ID for this ad hoc command event. (integer) * `type`: Data type for this ad hoc command event. (choice) * `url`: URL for this ad hoc command event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command event was created. (datetime) * `modified`: Timestamp when this ad hoc command event was last modified. (datetime) * `ad_hoc_command`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_ok`: Host OK     - `runner_on_unreachable`: Host Unreachable     - `runner_on_skipped`: Host Skipped     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandEventsApiAdHocCommandEventsAdHocCommandEventsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandEventsApiAdHocCommandEventsAdHocCommandEventsReadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

