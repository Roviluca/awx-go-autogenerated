# \CustomInventoryScriptsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomInventoryScriptsInventoryScriptsCopyCreate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsCopyCreate) | **Post** /api/v2/inventory_scripts/{id}/copy/ | 
[**CustomInventoryScriptsInventoryScriptsCopyList**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsCopyList) | **Get** /api/v2/inventory_scripts/{id}/copy/ | 
[**CustomInventoryScriptsInventoryScriptsCreate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsCreate) | **Post** /api/v2/inventory_scripts/ |  Create a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsDelete**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsDelete) | **Delete** /api/v2/inventory_scripts/{id}/ |  Delete a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsList**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsList) | **Get** /api/v2/inventory_scripts/ |  List Custom Inventory Scripts
[**CustomInventoryScriptsInventoryScriptsObjectRolesList**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsObjectRolesList) | **Get** /api/v2/inventory_scripts/{id}/object_roles/ |  List Roles for a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsPartialUpdate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsPartialUpdate) | **Patch** /api/v2/inventory_scripts/{id}/ |  Update a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsRead**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsRead) | **Get** /api/v2/inventory_scripts/{id}/ |  Retrieve a Custom Inventory Script
[**CustomInventoryScriptsInventoryScriptsUpdate**](CustomInventoryScriptsApi.md#CustomInventoryScriptsInventoryScriptsUpdate) | **Put** /api/v2/inventory_scripts/{id}/ |  Update a Custom Inventory Script


# **CustomInventoryScriptsInventoryScriptsCopyCreate**
> CustomInventoryScriptsInventoryScriptsCopyCreate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data26**](Data26.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsCopyList**
> CustomInventoryScriptsInventoryScriptsCopyList(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyListOpts struct

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

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsCreate**
> CustomInventoryScriptsInventoryScriptsCreate(ctx, optional)
 Create a Custom Inventory Script

 Make a POST request to this resource with the following custom inventory script fields to create a new custom inventory script:          * `name`: Name of this custom inventory script. (string, required) * `description`: Optional description of this custom inventory script. (string, default=`\"\"`) * `script`:  (string, required) * `organization`: Organization owning this inventory script (id, required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data23**](Data23.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsDelete**
> CustomInventoryScriptsInventoryScriptsDelete(ctx, id, optional)
 Delete a Custom Inventory Script

 Make a DELETE request to this resource to delete this custom inventory script.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsList**
> CustomInventoryScriptsInventoryScriptsList(ctx, optional)
 List Custom Inventory Scripts

 Make a GET request to this resource to retrieve the list of custom inventory scripts.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of custom inventory scripts found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more custom inventory script records.    ## Results  Each custom inventory script data structure includes the following fields:  * `id`: Database ID for this custom inventory script. (integer) * `type`: Data type for this custom inventory script. (choice) * `url`: URL for this custom inventory script. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this custom inventory script was created. (datetime) * `modified`: Timestamp when this custom inventory script was last modified. (datetime) * `name`: Name of this custom inventory script. (string) * `description`: Optional description of this custom inventory script. (string) * `script`:  (string) * `organization`: Organization owning this inventory script (id)    ## Sorting  To specify that custom inventory scripts are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsListOpts struct

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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsObjectRolesList**
> CustomInventoryScriptsInventoryScriptsObjectRolesList(ctx, id, optional)
 List Roles for a Custom Inventory Script

 Make a GET request to this resource to retrieve a list of roles associated with the selected custom inventory script.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of roles found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * `id`: Database ID for this role. (integer) * `type`: Data type for this role. (choice) * `url`: URL for this role. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `name`: Name of this role. (field) * `description`: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsObjectRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsObjectRolesListOpts struct

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

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsPartialUpdate**
> CustomInventoryScriptsInventoryScriptsPartialUpdate(ctx, id, optional)
 Update a Custom Inventory Script

 Make a PUT or PATCH request to this resource to update this custom inventory script.  The following fields may be modified:          * `name`: Name of this custom inventory script. (string, required) * `description`: Optional description of this custom inventory script. (string, default=`\"\"`) * `script`:  (string, required) * `organization`: Organization owning this inventory script (id, required)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data25**](Data25.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsRead**
> CustomInventoryScriptsInventoryScriptsRead(ctx, id, optional)
 Retrieve a Custom Inventory Script

 Make GET request to this resource to retrieve a single custom inventory script record containing the following fields:  * `id`: Database ID for this custom inventory script. (integer) * `type`: Data type for this custom inventory script. (choice) * `url`: URL for this custom inventory script. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this custom inventory script was created. (datetime) * `modified`: Timestamp when this custom inventory script was last modified. (datetime) * `name`: Name of this custom inventory script. (string) * `description`: Optional description of this custom inventory script. (string) * `script`:  (string) * `organization`: Organization owning this inventory script (id)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsReadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomInventoryScriptsInventoryScriptsUpdate**
> CustomInventoryScriptsInventoryScriptsUpdate(ctx, id, optional)
 Update a Custom Inventory Script

 Make a PUT or PATCH request to this resource to update this custom inventory script.  The following fields may be modified:          * `name`: Name of this custom inventory script. (string, required) * `description`: Optional description of this custom inventory script. (string, default=`\"\"`) * `script`:  (string, required) * `organization`: Organization owning this inventory script (id, required)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data24**](Data24.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

