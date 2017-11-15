/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package shortcode_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method SendDedicatedShortcode
 */
type SendDedicatedShortcodeInput struct {
    Shortcode             int64           //Your dedicated shortcode
    To                    float64         //The number to send your SMS to
    Body                  string          //The body of your message
    ResponseType          string          //Response type format xml or json
    Method                models_pkg.HttpActionEnum //Callback status method, POST or GET
    Messagestatuscallback *string         //Callback url for SMS status
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
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    PageSize        *int64          //Number of individual resources listed in the response per page
    From            *string         //Only list SMS messages sent from this number
    Shortcode       *string         //Only list SMS messages sent to Shortcode
    DateReceived    *string         //Only list SMS messages sent in the specified date MAKE REQUEST
}

/*
 * Client structure as interface implementation
 */
type SHORTCODE_IMPL struct { }

/**
 * TODO: type endpoint description here
 * @param  SendDedicatedShortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) SendDedicatedShortcode (input *SendDedicatedShortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/senddedicatedsms.{ResponseType}"

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
        "user-agent" : "message360-api",
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
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, message360_lib.Config.BasicAuthUserName, message360_lib.Config.BasicAuthPassword)
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
    _queryBuilder := message360_lib.BASEURI;

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
        "user-agent" : "message360-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "MessageSid" : input.MessageSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, message360_lib.Config.BasicAuthUserName, message360_lib.Config.BasicAuthPassword)
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
    _queryBuilder := message360_lib.BASEURI;

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
        "user-agent" : "message360-api",
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
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, message360_lib.Config.BasicAuthUserName, message360_lib.Config.BasicAuthPassword)
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
 * Retrive a list of inbound Sms Short Code messages associated with your message360 account.
 * @param  ListInboundShortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) ListInboundShortcode (input *ListInboundShortcodeInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/getinboundsms.{ResponseType}"

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
        "user-agent" : "message360-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Page" : apihelper_pkg.ToString(*input.Page, "1"),
        "PageSize" : apihelper_pkg.ToString(*input.PageSize, "10"),
        "From" : input.From,
        "Shortcode" : input.Shortcode,
        "DateReceived" : input.DateReceived,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, message360_lib.Config.BasicAuthUserName, message360_lib.Config.BasicAuthPassword)
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

