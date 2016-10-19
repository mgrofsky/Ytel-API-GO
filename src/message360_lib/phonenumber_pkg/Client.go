/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/18/2016
 */
package phonenumber_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type PHONENUMBER_IMPL struct { }

/**
 * Available Phone Number
 * @param    string         numberType       parameter: Required
 * @param    string         areaCode         parameter: Required
 * @param    *int64         pageSize         parameter: Optional
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateAvailablePhoneNumber (
            numberType string,
            areaCode string,
            pageSize *int64,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/availablenumber.{ResponseType}"

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

        "NumberType" : numberType,
        "AreaCode" : areaCode,
        "PageSize" : pageSize,

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
 * List Account's Phone number details
 * @param    *int64         page             parameter: Optional
 * @param    *int64         pageSize         parameter: Optional
 * @param    *string        numberType       parameter: Optional
 * @param    *string        friendlyName     parameter: Optional
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateListNumber (
            page *int64,
            pageSize *int64,
            numberType *string,
            friendlyName *string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/listnumber.{ResponseType}"

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
        "NumberType" : numberType,
        "FriendlyName" : friendlyName,

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
 * Release number from account
 * @param    string         phoneNumber      parameter: Required
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateReleaseNumber (
            phoneNumber string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/releasenumber.{ResponseType}"

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

        "PhoneNumber" : phoneNumber,

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
 * Buy Phone Number 
 * @param    string         phoneNumber      parameter: Required
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateBuyNumber (
            phoneNumber string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/buynumber.{ResponseType}"

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

        "PhoneNumber" : phoneNumber,

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
 * Get Phone Number Details
 * @param    string         phoneNumber      parameter: Required
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateViewNumberDetails (
            phoneNumber string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/viewnumber.{ResponseType}"

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

        "PhoneNumber" : phoneNumber,

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
 * Update Phone Number Details
 * @param    string                           phoneNumber              parameter: Required
 * @param    *string                          friendlyName             parameter: Optional
 * @param    *string                          voiceUrl                 parameter: Optional
 * @param    models_pkg.HttpMethodEnum        voiceMethod              parameter: Optional
 * @param    *string                          voiceFallbackUrl         parameter: Optional
 * @param    models_pkg.HttpMethodEnum        voiceFallbackMethod      parameter: Optional
 * @param    *string                          hangupCallback           parameter: Optional
 * @param    models_pkg.HttpMethodEnum        hangupCallbackMethod     parameter: Optional
 * @param    *string                          heartbeatUrl             parameter: Optional
 * @param    models_pkg.HttpMethodEnum        heartbeatMethod          parameter: Optional
 * @param    *string                          smsUrl                   parameter: Optional
 * @param    models_pkg.HttpMethodEnum        smsMethod                parameter: Optional
 * @param    *string                          smsFallbackUrl           parameter: Optional
 * @param    models_pkg.HttpMethodEnum        smsFallbackMethod        parameter: Optional
 * @param    *string                          responseType             parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) UpdatePhoneNumber (
            phoneNumber string,
            friendlyName *string,
            voiceUrl *string,
            voiceMethod models_pkg.HttpMethodEnum,
            voiceFallbackUrl *string,
            voiceFallbackMethod models_pkg.HttpMethodEnum,
            hangupCallback *string,
            hangupCallbackMethod models_pkg.HttpMethodEnum,
            heartbeatUrl *string,
            heartbeatMethod models_pkg.HttpMethodEnum,
            smsUrl *string,
            smsMethod models_pkg.HttpMethodEnum,
            smsFallbackUrl *string,
            smsFallbackMethod models_pkg.HttpMethodEnum,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/updatenumber.{ResponseType}"

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

        "PhoneNumber" : phoneNumber,
        "FriendlyName" : friendlyName,
        "VoiceUrl" : voiceUrl,
        "VoiceMethod" : models_pkg.HttpMethodEnumToValue(voiceMethod),
        "VoiceFallbackUrl" : voiceFallbackUrl,
        "VoiceFallbackMethod" : models_pkg.HttpMethodEnumToValue(voiceFallbackMethod),
        "HangupCallback" : hangupCallback,
        "HangupCallbackMethod" : models_pkg.HttpMethodEnumToValue(hangupCallbackMethod),
        "HeartbeatUrl" : heartbeatUrl,
        "HeartbeatMethod" : models_pkg.HttpMethodEnumToValue(heartbeatMethod),
        "SmsUrl" : smsUrl,
        "SmsMethod" : models_pkg.HttpMethodEnumToValue(smsMethod),
        "SmsFallbackUrl" : smsFallbackUrl,
        "SmsFallbackMethod" : models_pkg.HttpMethodEnumToValue(smsFallbackMethod),

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

