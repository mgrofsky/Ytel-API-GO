/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package sms_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method SendSMS
 */
type SendSMSInput struct {
    From                  string          //SMS enabled Message360 number to send this message from
    To                    string          //Number to send the SMS to
    Body                  string          //Text Message To Send
    ResponseType          string          //Response type format xml or json
    Method                models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once SMS sent.
    Messagestatuscallback *string         //URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished.
    Smartsms              *bool           //Check's 'to' number can receive sms or not using Carrier API, if wireless = true then text sms is sent, else wireless = false then call is recieved to end user with audible message.
}

/*
 * Input structure for the method ViewSMS
 */
type ViewSMSInput struct {
    Messagesid      string          //Message sid
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ListSMS
 */
type ListSMSInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //Messages sent from this number
    To              *string         //Messages sent to this number
    Datesent        *string         //Only list SMS messages sent in the specified date range
}

/*
 * Input structure for the method ListInboundSMS
 */
type ListInboundSMSInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    Pagesize        *int64          //Number of individual resources listed in the response per page
    From            *string         //From Number to Inbound SMS
    To              *string         //To Number to get Inbound SMS
    DateSent        *string         //Filter sms message objects by this date.
}

/*
 * Client structure as interface implementation
 */
type SMS_IMPL struct { }

/**
 * Send an SMS from a message360 number
 * @param  SendSMSInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) SendSMS (input *SendSMSInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/sendsms.{ResponseType}"

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

        "from" : input.From,
        "to" : input.To,
        "body" : input.Body,
        "method" : models_pkg.HttpActionEnumToValue(input.Method),
        "messagestatuscallback" : input.Messagestatuscallback,
        "smartsms" : apihelper_pkg.ToString(*input.Smartsms, false),

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
 * View a Particular SMS
 * @param  ViewSMSInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) ViewSMS (input *ViewSMSInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/viewsms.{ResponseType}"

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
 * List All SMS
 * @param  ListSMSInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) ListSMS (input *ListSMSInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/listsms.{ResponseType}"

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
 * List All Inbound SMS
 * @param  ListInboundSMSInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) ListInboundSMS (input *ListInboundSMSInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/getInboundsms.{ResponseType}"

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
        "DateSent" : input.DateSent,

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

