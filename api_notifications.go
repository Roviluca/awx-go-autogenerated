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

type NotificationsApiService service

/*
NotificationsApiService  List Notifications
 Make a GET request to this resource to retrieve the list of notifications.  The resulting data structure contains:      {         \&quot;count\&quot;: 99,         \&quot;next\&quot;: null,         \&quot;previous\&quot;: null,         \&quot;results\&quot;: [             ...         ]     }  The &#x60;count&#x60; field indicates the total number of notifications found for the given query.  The &#x60;next&#x60; and &#x60;previous&#x60; fields provides links to additional results if there are more than will fit on a single page.  The &#x60;results&#x60; list contains zero or more notification records.    ## Results  Each notification data structure includes the following fields:  * &#x60;id&#x60;: Database ID for this notification. (integer) * &#x60;type&#x60;: Data type for this notification. (choice) * &#x60;url&#x60;: URL for this notification. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this notification was created. (datetime) * &#x60;modified&#x60;: Timestamp when this notification was last modified. (datetime) * &#x60;notification_template&#x60;:  (id) * &#x60;error&#x60;:  (string) * &#x60;status&#x60;:  (choice)     - &#x60;pending&#x60;: Pending     - &#x60;successful&#x60;: Successful     - &#x60;failed&#x60;: Failed * &#x60;notifications_sent&#x60;:  (integer) * &#x60;notification_type&#x60;:  (choice)     - &#x60;email&#x60;: Email     - &#x60;grafana&#x60;: Grafana     - &#x60;irc&#x60;: IRC     - &#x60;mattermost&#x60;: Mattermost     - &#x60;pagerduty&#x60;: Pagerduty     - &#x60;rocketchat&#x60;: Rocket.Chat     - &#x60;slack&#x60;: Slack     - &#x60;twilio&#x60;: Twilio     - &#x60;webhook&#x60;: Webhook * &#x60;recipients&#x60;:  (string) * &#x60;subject&#x60;:  (string) * &#x60;body&#x60;: Notification body (json)    ## Sorting  To specify that notifications are returned in a particular order, use the &#x60;order_by&#x60; query string parameter on the GET request.      ?order_by&#x3D;name  Prefix the field name with a dash &#x60;-&#x60; to sort in reverse:      ?order_by&#x3D;-name  Multiple sorting fields may be specified by separating the field names with a comma &#x60;,&#x60;:      ?order_by&#x3D;name,some_other_field  ## Pagination  Use the &#x60;page_size&#x60; query string parameter to change the number of results returned for each request.  Use the &#x60;page&#x60; query string parameter to retrieve a particular page of results.      ?page_size&#x3D;100&amp;page&#x3D;2  The &#x60;previous&#x60; and &#x60;next&#x60; links returned with the results will set these query string parameters automatically.  ## Searching  Use the &#x60;search&#x60; query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search&#x3D;findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search&#x3D;findme
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NotificationsApiNotificationsNotificationsListOpts - Optional Parameters:
     * @param "Page" (optional.Int32) -  A page number within the paginated result set.
     * @param "PageSize" (optional.Int32) -  Number of results to return per page.
     * @param "Search" (optional.String) -  A search term.

*/

type NotificationsApiNotificationsNotificationsListOpts struct {
	Page     optional.Int32
	PageSize optional.Int32
	Search   optional.String
}

func (a *NotificationsApiService) NotificationsNotificationsList(ctx context.Context, localVarOptionals *NotificationsApiNotificationsNotificationsListOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/notifications/"

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
NotificationsApiService  Retrieve a Notification
 Make GET request to this resource to retrieve a single notification record containing the following fields:  * &#x60;id&#x60;: Database ID for this notification. (integer) * &#x60;type&#x60;: Data type for this notification. (choice) * &#x60;url&#x60;: URL for this notification. (string) * &#x60;related&#x60;: Data structure with URLs of related resources. (object) * &#x60;summary_fields&#x60;: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * &#x60;created&#x60;: Timestamp when this notification was created. (datetime) * &#x60;modified&#x60;: Timestamp when this notification was last modified. (datetime) * &#x60;notification_template&#x60;:  (id) * &#x60;error&#x60;:  (string) * &#x60;status&#x60;:  (choice)     - &#x60;pending&#x60;: Pending     - &#x60;successful&#x60;: Successful     - &#x60;failed&#x60;: Failed * &#x60;notifications_sent&#x60;:  (integer) * &#x60;notification_type&#x60;:  (choice)     - &#x60;email&#x60;: Email     - &#x60;grafana&#x60;: Grafana     - &#x60;irc&#x60;: IRC     - &#x60;mattermost&#x60;: Mattermost     - &#x60;pagerduty&#x60;: Pagerduty     - &#x60;rocketchat&#x60;: Rocket.Chat     - &#x60;slack&#x60;: Slack     - &#x60;twilio&#x60;: Twilio     - &#x60;webhook&#x60;: Webhook * &#x60;recipients&#x60;:  (string) * &#x60;subject&#x60;:  (string) * &#x60;body&#x60;: Notification body (json)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *NotificationsApiNotificationsNotificationsReadOpts - Optional Parameters:
     * @param "Search" (optional.String) -  A search term.

*/

type NotificationsApiNotificationsNotificationsReadOpts struct {
	Search optional.String
}

func (a *NotificationsApiService) NotificationsNotificationsRead(ctx context.Context, id string, localVarOptionals *NotificationsApiNotificationsNotificationsReadOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v2/notifications/{id}/"
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
