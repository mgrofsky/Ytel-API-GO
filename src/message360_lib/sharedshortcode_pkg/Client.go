/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package sharedshortcode_pkg


import(
	"errors"
	"message360_lib/models_pkg"
	"github.com/satori/go.uuid"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method ViewTemplate
 */
type ViewTemplateInput struct {
    Templateid      uuid.UUID       //The unique identifier for a template object
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ViewSharedShortcodes
 */
type ViewSharedShortcodesInput struct {
    Messagesid      string          //Message sid
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ListOutboundSharedShortcodes
 */
type ListOutboundSharedShortcodesInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //Messages sent from this number
    To              *string         //Messages sent to this number
    Datesent        *string         //Only list SMS messages sent in the specified date range
}

/*
 * Input structure for the method ListInboundSharedShortcodes
 */
type ListInboundSharedShortcodesInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //From Number to Inbound ShortCode
    Shortcode       *string         //Only list messages sent to this Short Code
    DateReceived    *string         //Only list messages sent with the specified date
}

/*
 * Input structure for the method SendSharedShortcode
 */
type SendSharedShortcodeInput struct {
    Shortcode             string          //The Short Code number that is the sender of this message
    To                    string          //A valid 10-digit number that should receive the message
    Templateid            uuid.UUID       //The unique identifier for the template used for the message
    ResponseType          string          //Response type format xml or json
    Data                  string          //format of your data, example: {companyname}:test,{otpcode}:1234
    Method                models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once the Short Code message is sent.
    MessageStatusCallback *string         //URL that can be requested to receive notification when Short Code message was sent.
}

/*
 * Input structure for the method ListTemplates
 */
type ListTemplatesInput struct {
    ResponseType    string          //Response type format xml or json
    Type            *string         //The type (category) of template Valid values: marketing, authorization
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //The count of objects to return per page.
    Shortcode       *string         //Only list templates of type
}

/*
 * Input structure for the method ViewKeyword
 */
type ViewKeywordInput struct {
    Keywordid       string          //The unique identifier of each keyword
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ListKeyword
 */
type ListKeywordInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    PageSize        *int64          //Number of individual resources listed in the response per page
    Keyword         *string         //Only list keywords of keyword
    Shortcode       *int64          //Only list keywords of shortcode
}

/*
 * Input structure for the method ViewAssignement
 */
type ViewAssignementInput struct {
    Shortcode       string          //List of valid Shortcode to your message360 account
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ListAssignment
 */
type ListAssignmentInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    PageSize        *int64          //Number of individual resources listed in the response per page
    Shortcode       *string         //Only list keywords of shortcode
}

/*
 * Input structure for the method UpdateAssignment
 */
type UpdateAssignmentInput struct {
    Shortcode         string          //List of valid shortcode to your message360 account
    ResponseType      string          //Response type format xml or json
    FriendlyName      *string         //User generated name of the shortcode
    CallbackUrl       *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished.
    CallbackMethod    models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required StatusCallBackUrl once call connects.
    FallbackUrl       *string         //URL used if any errors occur during execution of InboundXML or at initial request of the required Url provided with the POST.
    FallbackUrlMethod models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required FallbackUrl once call connects.
}

/*
 * Client structure as interface implementation
 */
type SHAREDSHORTCODE_IMPL struct { }

/**
 * View a Shared ShortCode Template
 * @param  ViewTemplateInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ViewTemplate (input *ViewTemplateInput) (string, error) {
    //validating required parameters
    if (input.Templateid == nil){
        return nil,errors.New("The property 'templateid' in the input object cannot be nil.")
}     //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/template/view.{ResponseType}"

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

        "templateid" : input.Templateid,

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
 * View a ShortCode Message
 * @param  ViewSharedShortcodesInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ViewSharedShortcodes (input *ViewSharedShortcodesInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/viewsms.{ResponseType}"

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

        "messagesid" : input.Messagesid,

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
 * List ShortCode Messages
 * @param  ListOutboundSharedShortcodesInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ListOutboundSharedShortcodes (input *ListOutboundSharedShortcodesInput) (string, error) {
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

        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "from" : input.From,
        "to" : input.To,
        "datesent" : input.Datesent,

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
 * List All Inbound ShortCode
 * @param  ListInboundSharedShortcodesInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ListInboundSharedShortcodes (input *ListInboundSharedShortcodesInput) (string, error) {
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

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "DateReceived" : input.DateReceived,
    })
    if err != nil {
        //error in query param handling
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

        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "from" : input.From,
        "Shortcode" : input.Shortcode,

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
 * Send an SMS from a message360 ShortCode
 * @param  SendSharedShortcodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) SendSharedShortcode (input *SendSharedShortcodeInput) (string, error) {
    //validating required parameters
    if (input.Templateid == nil){
        return nil,errors.New("The property 'templateid' in the input object cannot be nil.")
}     //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/sendsms.{ResponseType}"

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
        "templateid" : input.Templateid,
        "data" : input.Data,
        "Method" : models_pkg.HttpActionEnumToValue(input.Method),
        "MessageStatusCallback" : input.MessageStatusCallback,

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
 * List Shortcode Templates by Type
 * @param  ListTemplatesInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ListTemplates (input *ListTemplatesInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/template/lists.{ResponseType}"

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

        "type" : apihelper_pkg.ToString(*input.Type, "authorization"),
        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "Shortcode" : input.Shortcode,

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
 * View a set of properties for a single keyword.
 * @param  ViewKeywordInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ViewKeyword (input *ViewKeywordInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/keyword/view.{ResponseType}"

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

        "Keywordid" : input.Keywordid,

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
 * Retrieve a list of keywords associated with your message360 account.
 * @param  ListKeywordInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ListKeyword (input *ListKeywordInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/keyword/lists.{ResponseType}"

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
        "Keyword" : input.Keyword,
        "Shortcode" : input.Shortcode,

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
 * The response returned here contains all resource properties associated with the given Shortcode.
 * @param  ViewAssignementInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ViewAssignement (input *ViewAssignementInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/viewshortcode.{ResponseType}"

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
 * Retrieve a list of shortcode assignment associated with your message360 account.
 * @param  ListAssignmentInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) ListAssignment (input *ListAssignmentInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/listshortcode.{ResponseType}"

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

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "Shortcode" : input.Shortcode,
    })
    if err != nil {
        //error in query param handling
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
 * TODO: type endpoint description here
 * @param  UpdateAssignmentInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHAREDSHORTCODE_IMPL) UpdateAssignment (input *UpdateAssignmentInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/updateshortcode.{ResponseType}"

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
        "FriendlyName" : input.FriendlyName,
        "CallbackUrl" : input.CallbackUrl,
        "CallbackMethod" : models_pkg.HttpActionEnumToValue(input.CallbackMethod),
        "FallbackUrl" : input.FallbackUrl,
        "FallbackUrlMethod" : models_pkg.HttpActionEnumToValue(input.FallbackUrlMethod),

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

