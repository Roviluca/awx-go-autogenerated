/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type CustomInventoryScriptsApiService service

/*
CustomInventoryScriptsApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyCreateOpts - Optional Parameters:
     * @param "Body" (optional.Interface of Body26) -

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyCreateOpts struct {
	Body optional.Interface
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsCopyCreate(ctx context.Context, id string, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyCreateOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/{id}/copy/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyListOpts - Optional Parameters:
     * @param "Page" (optional.Int32) -  A page number within the paginated result set.
     * @param "PageSize" (optional.Int32) -  Number of results to return per page.
     * @param "Search" (optional.String) -  A search term.

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyListOpts struct {
	Page     optional.Int32
	PageSize optional.Int32
	Search   optional.String
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsCopyList(ctx context.Context, id string, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCopyListOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/{id}/copy/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService  Create a Custom Inventory Script
 Make a POST request to this resource with the following custom inventory script fields to create a new custom inventory script:          * &#x60;name&#x60;: Name of this custom inventory script. (string, required) * &#x60;description&#x60;: Optional description of this custom inventory script. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;) * &#x60;script&#x60;:  (string, required) * &#x60;organization&#x60;: Organization owning this inventory script (id, required)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCreateOpts - Optional Parameters:
     * @param "Body" (optional.Interface of Body23) -

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCreateOpts struct {
	Body optional.Interface
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsCreate(ctx context.Context, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsCreateOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService  Delete a Custom Inventory Script
 Make a DELETE request to this resource to delete this custom inventory script.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsDeleteOpts - Optional Parameters:
     * @param "Search" (optional.String) -  A search term.

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsDeleteOpts struct {
	Search optional.String
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsDelete(ctx context.Context, id string, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsDeleteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService  List Custom Inventory Scripts
 Make a GET request to this resource to retrieve the list of custom inventory scripts.  The resulting data structure contains:      {         \&quot;count\&quot;: 99,         \&quot;next\&quot;: null,         \&quot;previous\&quot;: null,         \&quot;results\&quot;: [             ...         ]     }  The &#x60;count&#x60; field indicates the total number of custom inventory scripts found for the given query.  The &#x60;next&#x60; and &#x60;previous&#x60; fields provides links to additional results if there are more than will fit on a single page.  The &#x60;results&#x60; list contains zero or more custom inventory script records.    ## Results  Each custom inventory script data structure includes the following fields:  * &#x60;id&#x60;: Database ID for this custom inventory script. (integer) * &#x60;type&#x60;: Data type for this custom inventory script. (choice) * &#x60;url&#x60;: URL for this custom inventory script. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this custom inventory script was created. (datetime) * &#x60;modified&#x60;: Timestamp when this custom inventory script was last modified. (datetime) * &#x60;name&#x60;: Name of this custom inventory script. (string) * &#x60;description&#x60;: Optional description of this custom inventory script. (string) * &#x60;script&#x60;:  (string) * &#x60;organization&#x60;: Organization owning this inventory script (id)    ## Sorting  To specify that custom inventory scripts are returned in a particular order, use the &#x60;order_by&#x60; query string parameter on the GET request.      ?order_by&#x3D;name  Prefix the field name with a dash &#x60;-&#x60; to sort in reverse:      ?order_by&#x3D;-name  Multiple sorting fields may be specified by separating the field names with a comma &#x60;,&#x60;:      ?order_by&#x3D;name,some_other_field  ## Pagination  Use the &#x60;page_size&#x60; query string parameter to change the number of results returned for each request.  Use the &#x60;page&#x60; query string parameter to retrieve a particular page of results.      ?page_size&#x3D;100&amp;page&#x3D;2  The &#x60;previous&#x60; and &#x60;next&#x60; links returned with the results will set these query string parameters automatically.  ## Searching  Use the &#x60;search&#x60; query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search&#x3D;findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search&#x3D;findme
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsListOpts - Optional Parameters:
     * @param "Page" (optional.Int32) -  A page number within the paginated result set.
     * @param "PageSize" (optional.Int32) -  Number of results to return per page.
     * @param "Search" (optional.String) -  A search term.

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsListOpts struct {
	Page     optional.Int32
	PageSize optional.Int32
	Search   optional.String
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsList(ctx context.Context, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsListOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService  List Roles for a Custom Inventory Script
 Make a GET request to this resource to retrieve a list of roles associated with the selected custom inventory script.  The resulting data structure contains:      {         \&quot;count\&quot;: 99,         \&quot;next\&quot;: null,         \&quot;previous\&quot;: null,         \&quot;results\&quot;: [             ...         ]     }  The &#x60;count&#x60; field indicates the total number of roles found for the given query.  The &#x60;next&#x60; and &#x60;previous&#x60; fields provides links to additional results if there are more than will fit on a single page.  The &#x60;results&#x60; list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * &#x60;id&#x60;: Database ID for this role. (integer) * &#x60;type&#x60;: Data type for this role. (choice) * &#x60;url&#x60;: URL for this role. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;name&#x60;: Name of this role. (field) * &#x60;description&#x60;: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the &#x60;order_by&#x60; query string parameter on the GET request.      ?order_by&#x3D;name  Prefix the field name with a dash &#x60;-&#x60; to sort in reverse:      ?order_by&#x3D;-name  Multiple sorting fields may be specified by separating the field names with a comma &#x60;,&#x60;:      ?order_by&#x3D;name,some_other_field  ## Pagination  Use the &#x60;page_size&#x60; query string parameter to change the number of results returned for each request.  Use the &#x60;page&#x60; query string parameter to retrieve a particular page of results.      ?page_size&#x3D;100&amp;page&#x3D;2  The &#x60;previous&#x60; and &#x60;next&#x60; links returned with the results will set these query string parameters automatically.  ## Searching  Use the &#x60;search&#x60; query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search&#x3D;findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search&#x3D;findme
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsObjectRolesListOpts - Optional Parameters:
     * @param "Page" (optional.Int32) -  A page number within the paginated result set.
     * @param "PageSize" (optional.Int32) -  Number of results to return per page.
     * @param "Search" (optional.String) -  A search term.

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsObjectRolesListOpts struct {
	Page     optional.Int32
	PageSize optional.Int32
	Search   optional.String
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsObjectRolesList(ctx context.Context, id string, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsObjectRolesListOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/{id}/object_roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService  Update a Custom Inventory Script
 Make a PUT or PATCH request to this resource to update this custom inventory script.  The following fields may be modified:          * &#x60;name&#x60;: Name of this custom inventory script. (string, required) * &#x60;description&#x60;: Optional description of this custom inventory script. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;) * &#x60;script&#x60;:  (string, required) * &#x60;organization&#x60;: Organization owning this inventory script (id, required)         For a PATCH request, include only the fields that are being modified.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsPartialUpdateOpts - Optional Parameters:
     * @param "Body" (optional.Interface of Body25) -
     * @param "Search" (optional.String) -  A search term.

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsPartialUpdateOpts struct {
	Body   optional.Interface
	Search optional.String
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsPartialUpdate(ctx context.Context, id string, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsPartialUpdateOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService  Retrieve a Custom Inventory Script
 Make GET request to this resource to retrieve a single custom inventory script record containing the following fields:  * &#x60;id&#x60;: Database ID for this custom inventory script. (integer) * &#x60;type&#x60;: Data type for this custom inventory script. (choice) * &#x60;url&#x60;: URL for this custom inventory script. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this custom inventory script was created. (datetime) * &#x60;modified&#x60;: Timestamp when this custom inventory script was last modified. (datetime) * &#x60;name&#x60;: Name of this custom inventory script. (string) * &#x60;description&#x60;: Optional description of this custom inventory script. (string) * &#x60;script&#x60;:  (string) * &#x60;organization&#x60;: Organization owning this inventory script (id)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsReadOpts - Optional Parameters:
     * @param "Search" (optional.String) -  A search term.

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsReadOpts struct {
	Search optional.String
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsRead(ctx context.Context, id string, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsReadOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CustomInventoryScriptsApiService  Update a Custom Inventory Script
 Make a PUT or PATCH request to this resource to update this custom inventory script.  The following fields may be modified:          * &#x60;name&#x60;: Name of this custom inventory script. (string, required) * &#x60;description&#x60;: Optional description of this custom inventory script. (string, default&#x3D;&#x60;\&quot;\&quot;&#x60;) * &#x60;script&#x60;:  (string, required) * &#x60;organization&#x60;: Organization owning this inventory script (id, required)       For a PUT request, include **all** fields in the request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsUpdateOpts - Optional Parameters:
     * @param "Body" (optional.Interface of Body24) -
     * @param "Search" (optional.String) -  A search term.

*/

type CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsUpdateOpts struct {
	Body   optional.Interface
	Search optional.String
}

func (a *CustomInventoryScriptsApiService) CustomInventoryScriptsInventoryScriptsUpdate(ctx context.Context, id string, localVarOptionals *CustomInventoryScriptsApiCustomInventoryScriptsInventoryScriptsUpdateOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/inventory_scripts/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
		localVarHttpResponse.Body.Close()
		if err != nil {
			return localVarHttpResponse, err
		}
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
