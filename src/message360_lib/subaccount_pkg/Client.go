/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/12/2016
 */
package subaccount_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method CreateSubAccount
 */
type CreateSubAccountInput struct {
    Firstname       string          //TODO: Write general description for this field
    Lastname        string          //TODO: Write general description for this field
    Email           string          //TODO: Write general description for this field
    ResponseType    *string         //ResponseType Format either json or xml
}

/*
 * Input structure for the method CreateSuspendSubAccount
 */
type CreateSuspendSubAccountInput struct {
    Subaccountsid   string          //TODO: Write general description for this field
    Activate        models_pkg.ActivateStatus //TODO: Write general description for this field
    ResponseType    *string         //TODO: Write general description for this field
}

/*
 * Input structure for the method CreateDeleteMergeSubAccount
 */
type CreateDeleteMergeSubAccountInput struct {
    Subaccountsid   string          //TODO: Write general description for this field
    Mergenumber     models_pkg.MergeNumberStatus //TODO: Write general description for this field
    ResponseType    *string         //Response type format either json or xml
}

/*
 * Client structure as interface implementation
 */
type SUBACCOUNT_IMPL struct { }

/**
 * Create Sub account
 * @param  CreateSubAccountInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SUBACCOUNT_IMPL) CreateSubAccount (input *CreateSubAccountInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/user/createsubaccount.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : apihelper_pkg.ToString(*input.ResponseType, "json"),
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

        "firstname" : input.Firstname,
        "lastname" : input.Lastname,
        "email" : input.Email,

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
 * Suspend or unsuspend
 * @param  CreateSuspendSubAccountInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SUBACCOUNT_IMPL) CreateSuspendSubAccount (input *CreateSuspendSubAccountInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/user/subaccountactivation.{ResponseType}"

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

        "subaccountsid" : input.Subaccountsid,
        "activate" : models_pkg.ActivateStatusToValue(input.Activate),

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
 * Delete or Merge Sub account
 * @param  CreateDeleteMergeSubAccountInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SUBACCOUNT_IMPL) CreateDeleteMergeSubAccount (input *CreateDeleteMergeSubAccountInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/user/deletesubaccount.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : apihelper_pkg.ToString(*input.ResponseType, "json"),
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

        "subaccountsid" : input.Subaccountsid,
        "mergenumber" : models_pkg.MergeNumberStatusToValue(input.Mergenumber),

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

