/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package phonenumber_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method CreateBuyNumber
 */
type CreateBuyNumberInput struct {
    PhoneNumber     string          //Phone number to be purchase
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method CreateReleaseNumber
 */
type CreateReleaseNumberInput struct {
    PhoneNumber     string          //Phone number to be relase
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method CreateViewNumberDetails
 */
type CreateViewNumberDetailsInput struct {
    PhoneNumber     string          //Get Phone number Detail
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method CreateListNumber
 */
type CreateListNumberInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    PageSize        *int64          //Number of individual resources listed in the response per page
    NumberType      models_pkg.NumberTypeEnum //TODO: Write general description for this field
    FriendlyName    *string         //TODO: Write general description for this field
}

/*
 * Input structure for the method CreateAvailablePhoneNumber
 */
type CreateAvailablePhoneNumberInput struct {
    NumberType      models_pkg.NumberTypeEnum //Number type either SMS,Voice or all
    AreaCode        string          //Phone Number Area Code
    ResponseType    string          //Response type format xml or json
    PageSize        *int64          //Page Size
}

/*
 * Input structure for the method UpdatePhoneNumber
 */
type UpdatePhoneNumberInput struct {
    PhoneNumber          string          //TODO: Write general description for this field
    ResponseType         string          //Response type format xml or json
    FriendlyName         *string         //TODO: Write general description for this field
    VoiceUrl             *string         //URL requested once the call connects
    VoiceMethod          models_pkg.HttpActionEnum //TODO: Write general description for this field
    VoiceFallbackUrl     *string         //URL requested if the voice URL is not available
    VoiceFallbackMethod  models_pkg.HttpActionEnum //TODO: Write general description for this field
    HangupCallback       *string         //TODO: Write general description for this field
    HangupCallbackMethod models_pkg.HttpActionEnum //TODO: Write general description for this field
    HeartbeatUrl         *string         //URL requested once the call connects
    HeartbeatMethod      models_pkg.HttpActionEnum //URL that can be requested every 60 seconds during the call to notify of elapsed time
    SmsUrl               *string         //URL requested when an SMS is received
    SmsMethod            models_pkg.HttpActionEnum //TODO: Write general description for this field
    SmsFallbackUrl       *string         //URL requested once the call connects
    SmsFallbackMethod    models_pkg.HttpActionEnum //URL requested if the sms URL is not available
}

/*
 * Client structure as interface implementation
 */
type PHONENUMBER_IMPL struct { }

/**
 * Buy Phone Number 
 * @param  CreateBuyNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateBuyNumber (input *CreateBuyNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/buynumber.{ResponseType}"

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

        "PhoneNumber" : input.PhoneNumber,

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
 * @param  CreateReleaseNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateReleaseNumber (input *CreateReleaseNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/releasenumber.{ResponseType}"

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

        "PhoneNumber" : input.PhoneNumber,

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
 * @param  CreateViewNumberDetailsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateViewNumberDetails (input *CreateViewNumberDetailsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/viewnumber.{ResponseType}"

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

        "PhoneNumber" : input.PhoneNumber,

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
 * @param  CreateListNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateListNumber (input *CreateListNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/listnumber.{ResponseType}"

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

        "Page" : input.Page,
        "PageSize" : apihelper_pkg.ToString(*input.PageSize, "10"),
        "NumberType" : models_pkg.NumberTypeEnumToValue(input.NumberType),
        "FriendlyName" : input.FriendlyName,

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
 * Available Phone Number
 * @param  CreateAvailablePhoneNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) CreateAvailablePhoneNumber (input *CreateAvailablePhoneNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/availablenumber.{ResponseType}"

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

        "NumberType" : models_pkg.NumberTypeEnumToValue(input.NumberType),
        "AreaCode" : input.AreaCode,
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
 * Update Phone Number Details
 * @param  UpdatePhoneNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) UpdatePhoneNumber (input *UpdatePhoneNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/updatenumber.{ResponseType}"

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

        "PhoneNumber" : input.PhoneNumber,
        "FriendlyName" : input.FriendlyName,
        "VoiceUrl" : input.VoiceUrl,
        "VoiceMethod" : models_pkg.HttpActionEnumToValue(input.VoiceMethod),
        "VoiceFallbackUrl" : input.VoiceFallbackUrl,
        "VoiceFallbackMethod" : models_pkg.HttpActionEnumToValue(input.VoiceFallbackMethod),
        "HangupCallback" : input.HangupCallback,
        "HangupCallbackMethod" : models_pkg.HttpActionEnumToValue(input.HangupCallbackMethod),
        "HeartbeatUrl" : input.HeartbeatUrl,
        "HeartbeatMethod" : models_pkg.HttpActionEnumToValue(input.HeartbeatMethod),
        "SmsUrl" : input.SmsUrl,
        "SmsMethod" : models_pkg.HttpActionEnumToValue(input.SmsMethod),
        "SmsFallbackUrl" : input.SmsFallbackUrl,
        "SmsFallbackMethod" : models_pkg.HttpActionEnumToValue(input.SmsFallbackMethod),

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

