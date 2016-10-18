/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/18/2016
 */
package recording_pkg


import(
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type RECORDING_IMPL struct { }

/**
 * Delete Recording Record
 * @param    string         recordingSid     parameter: Required
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateDeleteRecording (
            recordingSid string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/deleterecording.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : apihelper_pkg.ToString(*responseType, "json"),
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

        "RecordingSid" : recordingSid,

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
 * View a specific Recording
 * @param    string         recordingSid     parameter: Required
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateViewRecording (
            recordingSid string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/viewrecording.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : apihelper_pkg.ToString(*responseType, "json"),
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

        "RecordingSid" : recordingSid,

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
 * List out Recordings
 * @param    *int64         page             parameter: Optional
 * @param    *int64         pageSize         parameter: Optional
 * @param    *string        dateCreated      parameter: Optional
 * @param    *string        callSid          parameter: Optional
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateListRecording (
            page *int64,
            pageSize *int64,
            dateCreated *string,
            callSid *string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/listrecording.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : apihelper_pkg.ToString(*responseType, "json"),
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

        "Page" : page,
        "PageSize" : pageSize,
        "DateCreated" : dateCreated,
        "CallSid" : callSid,

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

