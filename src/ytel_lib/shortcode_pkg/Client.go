/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package shortcode_pkg


import(
	"ytel_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"ytel_lib/apihelper_pkg"
	"ytel_lib/configuration_pkg"
)

/*
 * Input structure for the method SendDedicatedShortcode
 */
type SendDedicatedShortcodeInput struct {
    Shortcode             int64           //Your dedicated shortcode
    To                    float64         //The number to send your SMS to
    Body                  string          //The body of your message
    ResponseType          string          //Response type format xml or json
    Method                models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once the Short Code message is sent.GET or POST
    Messagestatuscallback *string         //URL that can be requested to receive notification when Short Code message was sent.
}

/*
 * Input structure for the method ViewShortcode
 */
type ViewShortcodeInput struct {
    MessageSid      string          //The unique identifier for the sms resource
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ListShortcode
 */
type ListShortcodeInput struct {
    ResponseType    string          //Response type format xml or json
    Shortcode       *string         //Only list messages sent from this Short Code
    To              *string         //Only list messages sent to this number
    DateSent        *string         //Only list messages sent with the specified date
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    PageSize        *int64          //The count of objects to return per page.
}

/*
 * Input structure for the method ListInboundShortcode
 */
type ListInboundShortcodeInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //Only list SMS messages sent from this number
    Shortcode       *string         //Only list SMS messages sent to Shortcode
    Datecreated     *string         //Only list SMS messages sent in the specified date MAKE REQUEST
}

/*
 * Input structure for the method ViewDedicatedShortcodeAssignment
 */
type ViewDedicatedShortcodeAssignmentInput struct {
    Shortcode       string          //List of valid Dedicated Short Code to your Ytel account
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method UpdateDedicatedShortCodeAssignment
 */
type UpdateDedicatedShortCodeAssignmentInput struct {
    Shortcode       string          //List of valid dedicated shortcode to your Ytel account.
    ResponseType    string          //Response type format xml or json
    FriendlyName    *string         //User generated name of the dedicated shortcode.
    CallbackMethod  *string         //Specifies the HTTP method used to request the required StatusCallBackUrl once call connects.
    CallbackUrl     *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished.
    FallbackMethod  *string         //Specifies the HTTP method used to request the required FallbackUrl once call connects.
    FallbackUrl     *string         //URL used if any errors occur during execution of InboundXML or at initial request of the required Url provided with the POST.
}

/*
 * Input structure for the method ListShortCodeAssignment
 */
type ListShortCodeAssignmentInput struct {
    ResponseType    string          //Response type format xml or json
    Shortcode       *string         //Only list Short Code Assignment sent from this Short Code
    Page            *string         //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *string         //The count of objects to return per page.
}

/*
 * Client structure as interface implementation
 */
type SHORTCODE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * TODO: type endpoint description here
 * @param  SendDedicatedShortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) SendDedicatedShortcode (input *SendDedicatedShortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/sendsms.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : input.ResponseType,
    })
    if err != nil {
        //error in template param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "shortcode" : input.Shortcode,
        "to" : input.To,
        "body" : input.Body,
        "method" : models_pkg.HttpActionEnumToValue(input.Method),
        "messagestatuscallback" : input.Messagestatuscallback,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * View a single Sms Short Code message.
 * @param  ViewShortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) ViewShortcode (input *ViewShortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/viewsms..{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : input.ResponseType,
    })
    if err != nil {
        //error in template param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "MessageSid" : input.MessageSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a list of Short Code message objects.
 * @param  ListShortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) ListShortcode (input *ListShortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/listsms.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : input.ResponseType,
    })
    if err != nil {
        //error in template param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : input.Shortcode,
        "To" : input.To,
        "DateSent" : input.DateSent,
        "Page" : apihelper_pkg.ToString(*input.Page, "1"),
        "PageSize" : apihelper_pkg.ToString(*input.PageSize, "10"),

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrive a list of inbound Sms Short Code messages associated with your Ytel account.
 * @param  ListInboundShortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) ListInboundShortcode (input *ListInboundShortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/getinboundsms.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : input.ResponseType,
    })
    if err != nil {
        //error in template param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "From" : input.From,
        "Shortcode" : input.Shortcode,
        "Datecreated" : input.Datecreated,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a single Short Code object by its shortcode assignment.
 * @param  ViewDedicatedShortcodeAssignmentInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) ViewDedicatedShortcodeAssignment (input *ViewDedicatedShortcodeAssignmentInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/viewshortcode.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : input.ResponseType,
    })
    if err != nil {
        //error in template param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : input.Shortcode,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Update a dedicated shortcode.
 * @param  UpdateDedicatedShortCodeAssignmentInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) UpdateDedicatedShortCodeAssignment (input *UpdateDedicatedShortCodeAssignmentInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/updateshortcode.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : input.ResponseType,
    })
    if err != nil {
        //error in template param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : input.Shortcode,
        "FriendlyName" : input.FriendlyName,
        "CallbackMethod" : input.CallbackMethod,
        "CallbackUrl" : input.CallbackUrl,
        "FallbackMethod" : input.FallbackMethod,
        "FallbackUrl" : input.FallbackUrl,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a list of Short Code assignment associated with your Ytel account.
 * @param  ListShortCodeAssignmentInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) ListShortCodeAssignment (input *ListShortCodeAssignmentInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/listshortcode.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : input.ResponseType,
    })
    if err != nil {
        //error in template param handling
        return "", err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : input.Shortcode,
        "page" : input.Page,
        "pagesize" : input.Pagesize,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

