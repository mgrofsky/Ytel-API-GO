/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package shortcode_pkg


import(
	"errors"
	"github.com/satori/go.uuid"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method CreateViewTemplate
 */
type CreateViewTemplateInput struct {
    Templateid      uuid.UUID       //The unique identifier for a template object
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method CreateSendShortCode
 */
type CreateSendShortCodeInput struct {
    Shortcode             string          //The Short Code number that is the sender of this message
    Tocountrycode         string          //The country code for sending this message
    To                    string          //A valid 10-digit number that should receive the message+
    Templateid            uuid.UUID       //The unique identifier for the template used for the message
    ResponseType          string          //Response type format xml or json
    Data                  string          //format of your data, example: {companyname}:test,{otpcode}:1234
    Method                *string         //Specifies the HTTP method used to request the required URL once the Short Code message is sent.
    MessageStatusCallback *string         //URL that can be requested to receive notification when Short Code message was sent.
}

/*
 * Input structure for the method CreateListInboundShortCode
 */
type CreateListInboundShortCodeInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //From Number to Inbound ShortCode
    Shortcode       *string         //Only list messages sent to this Short Code
    DateReceived    *string         //Only list messages sent with the specified date
}

/*
 * Input structure for the method CreateListShortCode
 */
type CreateListShortCodeInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //Messages sent from this number
    To              *string         //Messages sent to this number
    Datesent        *string         //Only list SMS messages sent in the specified date range
}

/*
 * Input structure for the method CreateListTemplates
 */
type CreateListTemplatesInput struct {
    ResponseType    string          //Response type format xml or json
    Type            *string         //The type (category) of template Valid values: marketing, authorization
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //The count of objects to return per page.
}

/*
 * Input structure for the method CreateViewShortCode
 */
type CreateViewShortCodeInput struct {
    Messagesid      string          //Message sid
    ResponseType    string          //Response type format xml or json
}

/*
 * Client structure as interface implementation
 */
type SHORTCODE_IMPL struct { }

/**
 * View a Shared ShortCode Template
 * @param  CreateViewTemplateInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) CreateViewTemplate (input *CreateViewTemplateInput) (string, error) {
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
 * Send an SMS from a message360 ShortCode
 * @param  CreateSendShortCodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) CreateSendShortCode (input *CreateSendShortCodeInput) (string, error) {
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
        "tocountrycode" : input.Tocountrycode,
        "to" : input.To,
        "templateid" : input.Templateid,
        "data" : input.Data,
        "Method" : apihelper_pkg.ToString(*input.Method, "GET"),
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
 * List All Inbound ShortCode
 * @param  CreateListInboundShortCodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) CreateListInboundShortCode (input *CreateListInboundShortCodeInput) (string, error) {
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

        "page" : input.Page,
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
 * List ShortCode Messages
 * @param  CreateListShortCodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) CreateListShortCode (input *CreateListShortCodeInput) (string, error) {
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

        "page" : input.Page,
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
 * List Shortcode Templates by Type
 * @param  CreateListTemplatesInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) CreateListTemplates (input *CreateListTemplatesInput) (string, error) {
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
        "page" : input.Page,
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),

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
 * @param  CreateViewShortCodeInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SHORTCODE_IMPL) CreateViewShortCode (input *CreateViewShortCodeInput) (string, error) {
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

