/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package phonenumber_pkg


import(
	"ytel_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"ytel_lib/apihelper_pkg"
	"ytel_lib/configuration_pkg"
)

/*
 * Input structure for the method AvailablePhoneNumber
 */
type AvailablePhoneNumberInput struct {
    Numbertype      models_pkg.NumberTypeEnum //Number type either SMS,Voice or all
    Areacode        string          //Specifies the area code for the returned list of available numbers. Only available for North American numbers.
    ResponseType    string          //Response type format xml or json
    Pagesize        *int64          //The count of objects to return.
}

/*
 * Input structure for the method MassReleaseNumber
 */
type MassReleaseNumberInput struct {
    PhoneNumber     string          //A valid Ytel comma separated numbers (E.164 format).
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ViewNumberDetails
 */
type ViewNumberDetailsInput struct {
    PhoneNumber     string          //A valid Ytel 10-digit phone number (E.164 format).
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ReleaseNumber
 */
type ReleaseNumberInput struct {
    PhoneNumber     string          //A valid 10-digit Ytel number (E.164 format).
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method BuyNumber
 */
type BuyNumberInput struct {
    PhoneNumber     string          //A valid 10-digit Ytel number (E.164 format).
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method UpdatePhoneNumber
 */
type UpdatePhoneNumberInput struct {
    PhoneNumber          string          //A valid Ytel number (E.164 format).
    VoiceUrl             string          //URL requested once the call connects
    ResponseType         string          //Response type format xml or json
    FriendlyName         *string         //Phone number friendly name description
    VoiceMethod          models_pkg.HttpActionEnum //Post or Get
    VoiceFallbackUrl     *string         //URL requested if the voice URL is not available
    VoiceFallbackMethod  models_pkg.HttpActionEnum //Post or Get
    HangupCallback       *string         //callback url after a hangup occurs
    HangupCallbackMethod models_pkg.HttpActionEnum //Post or Get
    HeartbeatUrl         *string         //URL requested once the call connects
    HeartbeatMethod      models_pkg.HttpActionEnum //URL that can be requested every 60 seconds during the call to notify of elapsed time
    SmsUrl               *string         //URL requested when an SMS is received
    SmsMethod            models_pkg.HttpActionEnum //Post or Get
    SmsFallbackUrl       *string         //URL used if any errors occur during execution of InboundXML from an SMS or at initial request of the SmsUrl.
    SmsFallbackMethod    models_pkg.HttpActionEnum //The HTTP method Ytel will use when URL requested if the SmsUrl is not available.
}

/*
 * Input structure for the method TransferNumber
 */
type TransferNumberInput struct {
    Phonenumber     string          //A valid 10-digit Ytel number (E.164 format).
    Fromaccountsid  string          //A specific Accountsid from where Number is getting transfer.
    Toaccountsid    string          //A specific Accountsid to which Number is getting transfer.
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ListNumber
 */
type ListNumberInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Page indexing starts at 1.
    PageSize        *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    NumberType      models_pkg.NumberTypeEnum //The capability supported by the number.Number type either SMS,Voice or all
    FriendlyName    *string         //A human-readable label added to the number object.
}

/*
 * Input structure for the method MassUpdateNumber
 */
type MassUpdateNumberInput struct {
    PhoneNumber          string          //A valid comma(,) separated Ytel numbers. (E.164 format).
    VoiceUrl             string          //The URL returning InboundXML incoming calls should execute when connected.
    ResponseType         string          //Response type format xml or json
    FriendlyName         *string         //A human-readable value for labeling the number.
    VoiceMethod          models_pkg.HttpActionEnum //Specifies the HTTP method used to request the VoiceUrl once incoming call connects.
    VoiceFallbackUrl     *string         //URL used if any errors occur during execution of InboundXML on a call or at initial request of the voice url
    VoiceFallbackMethod  models_pkg.HttpActionEnum //Specifies the HTTP method used to request the VoiceFallbackUrl once incoming call connects.
    HangupCallback       *string         //URL that can be requested to receive notification when and how incoming call has ended.
    HangupCallbackMethod models_pkg.HttpActionEnum //The HTTP method Ytel will use when requesting the HangupCallback URL.
    HeartbeatUrl         *string         //URL that can be used to monitor the phone number.
    HeartbeatMethod      models_pkg.HttpActionEnum //The HTTP method Ytel will use when requesting the HeartbeatUrl.
    SmsUrl               *string         //URL requested when an SMS is received.
    SmsMethod            models_pkg.HttpActionEnum //The HTTP method Ytel will use when requesting the SmsUrl.
    SmsFallbackUrl       *string         //URL used if any errors occur during execution of InboundXML from an SMS or at initial request of the SmsUrl.
    SmsFallbackMethod    models_pkg.HttpActionEnum //The HTTP method Ytel will use when URL requested if the SmsUrl is not available.
}

/*
 * Input structure for the method GetDIDScoreNumber
 */
type GetDIDScoreNumberInput struct {
    Phonenumber     string          //Specifies the multiple phone numbers for check updated spamscore .
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method BulkBuyNumber
 */
type BulkBuyNumberInput struct {
    NumberType      models_pkg.NumberTypeEnum //The capability the number supports.
    AreaCode        string          //Specifies the area code for the returned list of available numbers. Only available for North American numbers.
    Quantity        string          //A positive integer that tells how many number you want to buy at a time.
    ResponseType    string          //Response type format xml or json
    Leftover        *string         //If desired quantity is unavailable purchase what is available .
}

/*
 * Client structure as interface implementation
 */
type PHONENUMBER_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Retrieve a list of available phone numbers that can be purchased and used for your Ytel account.
 * @param  AvailablePhoneNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) AvailablePhoneNumber (input *AvailablePhoneNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "numbertype" : models_pkg.NumberTypeEnumToValue(input.Numbertype),
        "areacode" : input.Areacode,
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),

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
 * Remove a purchased Ytel number from your account.
 * @param  MassReleaseNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) MassReleaseNumber (input *MassReleaseNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/massreleasenumber.{ResponseType}"

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

        "PhoneNumber" : input.PhoneNumber,

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
 * Retrieve the details for a phone number by its number.
 * @param  ViewNumberDetailsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) ViewNumberDetails (input *ViewNumberDetailsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "PhoneNumber" : input.PhoneNumber,

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
 * Remove a purchased Ytel number from your account.
 * @param  ReleaseNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) ReleaseNumber (input *ReleaseNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "PhoneNumber" : input.PhoneNumber,

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
 * Purchase a phone number to be used with your Ytel account
 * @param  BuyNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) BuyNumber (input *BuyNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "PhoneNumber" : input.PhoneNumber,

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
 * Update properties for a Ytel number that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.
 * @param  UpdatePhoneNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) UpdatePhoneNumber (input *UpdatePhoneNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "PhoneNumber" : input.PhoneNumber,
        "VoiceUrl" : input.VoiceUrl,
        "FriendlyName" : input.FriendlyName,
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
 * Transfer phone number that has been purchased for from one account to another account.
 * @param  TransferNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) TransferNumber (input *TransferNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/transferphonenumbers.{ResponseType}"

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

        "phonenumber" : input.Phonenumber,
        "fromaccountsid" : input.Fromaccountsid,
        "toaccountsid" : input.Toaccountsid,

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
 * Retrieve a list of purchased phones numbers associated with your Ytel account.
 * @param  ListNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) ListNumber (input *ListNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Page" : apihelper_pkg.ToString(*input.Page, "1"),
        "PageSize" : apihelper_pkg.ToString(*input.PageSize, "10"),
        "NumberType" : models_pkg.NumberTypeEnumToValue(input.NumberType),
        "FriendlyName" : input.FriendlyName,

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
 * Update properties for a Ytel numbers that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.
 * @param  MassUpdateNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) MassUpdateNumber (input *MassUpdateNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/massupdatenumber.{ResponseType}"

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

        "PhoneNumber" : input.PhoneNumber,
        "VoiceUrl" : input.VoiceUrl,
        "FriendlyName" : input.FriendlyName,
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
 * Get DID Score Number
 * @param  GetDIDScoreNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) GetDIDScoreNumber (input *GetDIDScoreNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/getdidscore.{ResponseType}"

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

        "Phonenumber" : input.Phonenumber,

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
 * Purchase a selected number of DID's from specific area codes to be used with your Ytel account.
 * @param  BulkBuyNumberInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *PHONENUMBER_IMPL) BulkBuyNumber (input *BulkBuyNumberInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/incomingphone/bulkbuy.{ResponseType}"

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

        "NumberType" : models_pkg.NumberTypeEnumToValue(input.NumberType),
        "AreaCode" : input.AreaCode,
        "Quantity" : input.Quantity,
        "Leftover" : input.Leftover,

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

