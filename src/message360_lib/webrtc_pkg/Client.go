/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package webrtc_pkg


import(
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method CheckFunds
 */
type CheckFundsInput struct {
    AccountSid      string          //Your message360 Account SID
    AuthToken       string          //Your message360 Token
}

/*
 * Input structure for the method CreateToken
 */
type CreateTokenInput struct {
    AccountSid      string          //Your message360 Account SID
    AuthToken       string          //Your message360 Token
    Username        string          //WebRTC username authentication
    Password        string          //WebRTC password authentication
}

/*
 * Client structure as interface implementation
 */
type WEBRTC_IMPL struct { }

/**
 * TODO: type endpoint description here
 * @param  CheckFundsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *WEBRTC_IMPL) CheckFunds (input *CheckFundsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/webrtc/checkFunds.json"

    //variable to hold errors
    var err error = nil
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

        "account_sid" : input.AccountSid,
        "auth_token" : input.AuthToken,

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
 * message360 webrtc
 * @param  CreateTokenInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *WEBRTC_IMPL) CreateToken (input *CreateTokenInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/webrtc/agentLogin.json"

    //variable to hold errors
    var err error = nil
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

        "account_sid" : input.AccountSid,
        "auth_token" : input.AuthToken,
        "username" : input.Username,
        "password" : input.Password,

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

