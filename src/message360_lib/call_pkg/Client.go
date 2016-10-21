/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/21/2016
 */
package call_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type CALL_IMPL struct { }

/**
 * View Call Response
 * @param    string         callsid          parameter: Required
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateViewCall (
            callsid string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/viewcalls.{ResponseType}"

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

        "callsid" : callsid,

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
 * You can experiment with initiating a call through Message360 and view the request response generated when doing so and get the response in json
 * @param    string                       fromCountryCode           parameter: Required
 * @param    string                       from                      parameter: Required
 * @param    string                       toCountryCode             parameter: Required
 * @param    string                       to                        parameter: Required
 * @param    string                       url                       parameter: Required
 * @param    models_pkg.HttpMethod        method                    parameter: Optional
 * @param    *string                      statusCallBackUrl         parameter: Optional
 * @param    models_pkg.HttpMethod        statusCallBackMethod      parameter: Optional
 * @param    *string                      fallBackUrl               parameter: Optional
 * @param    models_pkg.HttpMethod        fallBackMethod            parameter: Optional
 * @param    *string                      heartBeatUrl              parameter: Optional
 * @param    *bool                        heartBeatMethod           parameter: Optional
 * @param    *int64                       timeout                   parameter: Optional
 * @param    *string                      playDtmf                  parameter: Optional
 * @param    *bool                        hideCallerId              parameter: Optional
 * @param    *bool                        record                    parameter: Optional
 * @param    *string                      recordCallBackUrl         parameter: Optional
 * @param    models_pkg.HttpMethod        recordCallBackMethod      parameter: Optional
 * @param    *bool                        transcribe                parameter: Optional
 * @param    *string                      transcribeCallBackUrl     parameter: Optional
 * @param    models_pkg.IfMachine         ifMachine                 parameter: Optional
 * @param    *string                      responseType              parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateMakeCall (
            fromCountryCode string,
            from string,
            toCountryCode string,
            to string,
            url string,
            method models_pkg.HttpMethod,
            statusCallBackUrl *string,
            statusCallBackMethod models_pkg.HttpMethod,
            fallBackUrl *string,
            fallBackMethod models_pkg.HttpMethod,
            heartBeatUrl *string,
            heartBeatMethod *bool,
            timeout *int64,
            playDtmf *string,
            hideCallerId *bool,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod models_pkg.HttpMethod,
            transcribe *bool,
            transcribeCallBackUrl *string,
            ifMachine models_pkg.IfMachine,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/makecall.{ResponseType}"

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

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "Method" : models_pkg.HttpMethodToValue(method),
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

        "FromCountryCode" : fromCountryCode,
        "From" : from,
        "ToCountryCode" : toCountryCode,
        "To" : to,
        "Url" : url,
        "StatusCallBackUrl" : statusCallBackUrl,
        "StatusCallBackMethod" : models_pkg.HttpMethodToValue(statusCallBackMethod),
        "FallBackUrl" : fallBackUrl,
        "FallBackMethod" : models_pkg.HttpMethodToValue(fallBackMethod),
        "HeartBeatUrl" : heartBeatUrl,
        "HeartBeatMethod" : heartBeatMethod,
        "Timeout" : timeout,
        "PlayDtmf" : playDtmf,
        "HideCallerId" : hideCallerId,
        "Record" : record,
        "RecordCallBackUrl" : recordCallBackUrl,
        "RecordCallBackMethod" : models_pkg.HttpMethodToValue(recordCallBackMethod),
        "Transcribe" : transcribe,
        "TranscribeCallBackUrl" : transcribeCallBackUrl,
        "IfMachine" : models_pkg.IfMachineToValue(ifMachine),

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
 * Play Dtmf and send the Digit
 * @param    int64                       length           parameter: Required
 * @param    models_pkg.Direction        direction        parameter: Required
 * @param    bool                        loop             parameter: Required
 * @param    bool                        mix              parameter: Required
 * @param    *string                     callSid          parameter: Optional
 * @param    *string                     audioUrl         parameter: Optional
 * @param    *string                     responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreatePlayAudio (
            length int64,
            direction models_pkg.Direction,
            loop bool,
            mix bool,
            callSid *string,
            audioUrl *string,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/playaudios.{ResponseType}"

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

        "Length" : length,
        "Direction" : models_pkg.DirectionToValue(direction),
        "Loop" : loop,
        "Mix" : mix,
        "CallSid" : callSid,
        "AudioUrl" : audioUrl,

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
 * Record a Call
 * @param    string                        callSid          parameter: Required
 * @param    bool                          record           parameter: Required
 * @param    models_pkg.Direction          direction        parameter: Optional
 * @param    *int64                        timeLimit        parameter: Optional
 * @param    *string                       callBackUrl      parameter: Optional
 * @param    models_pkg.AudioFormat        fileformat       parameter: Optional
 * @param    *string                       responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateRecordCall (
            callSid string,
            record bool,
            direction models_pkg.Direction,
            timeLimit *int64,
            callBackUrl *string,
            fileformat models_pkg.AudioFormat,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/recordcalls.{ResponseType}"

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

        "CallSid" : callSid,
        "Record" : record,
        "Direction" : models_pkg.DirectionToValue(direction),
        "TimeLimit" : timeLimit,
        "CallBackUrl" : callBackUrl,
        "Fileformat" : models_pkg.AudioFormatToValue(fileformat),

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
 * Voice Effect
 * @param    string                           callSid            parameter: Required
 * @param    models_pkg.AudioDirection        audioDirection     parameter: Optional
 * @param    *float64                         pitchSemiTones     parameter: Optional
 * @param    *float64                         pitchOctaves       parameter: Optional
 * @param    *float64                         pitch              parameter: Optional
 * @param    *float64                         rate               parameter: Optional
 * @param    *float64                         tempo              parameter: Optional
 * @param    *string                          responseType       parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateVoiceEffect (
            callSid string,
            audioDirection models_pkg.AudioDirection,
            pitchSemiTones *float64,
            pitchOctaves *float64,
            pitch *float64,
            rate *float64,
            tempo *float64,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/voiceeffect.{ResponseType}"

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

        "CallSid" : callSid,
        "AudioDirection" : models_pkg.AudioDirectionToValue(audioDirection),
        "PitchSemiTones" : pitchSemiTones,
        "PitchOctaves" : pitchOctaves,
        "Pitch" : pitch,
        "Rate" : rate,
        "Tempo" : tempo,

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
 * Play Dtmf and send the Digit
 * @param    string                      callSid               parameter: Required
 * @param    string                      playDtmf              parameter: Required
 * @param    models_pkg.Direction        playDtmfDirection     parameter: Optional
 * @param    *string                     responseType          parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateSendDigit (
            callSid string,
            playDtmf string,
            playDtmfDirection models_pkg.Direction,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/senddigits.{ResponseType}"

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

        "CallSid" : callSid,
        "PlayDtmf" : playDtmf,
        "PlayDtmfDirection" : models_pkg.DirectionToValue(playDtmfDirection),

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
 * Interrupt the Call by Call Sid
 * @param    string                                  callSid          parameter: Required
 * @param    *string                                 url              parameter: Optional
 * @param    models_pkg.HttpMethod                   method           parameter: Optional
 * @param    models_pkg.InterruptedCallStatus        status           parameter: Optional
 * @param    *string                                 responseType     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateInterruptedCall (
            callSid string,
            url *string,
            method models_pkg.HttpMethod,
            status models_pkg.InterruptedCallStatus,
            responseType *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/interruptcalls.{ResponseType}"

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

        "CallSid" : callSid,
        "Url" : url,
        "Method" : models_pkg.HttpMethodToValue(method),
        "Status" : models_pkg.InterruptedCallStatusToValue(status),

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
 * A list of calls associated with your Message360 account
 * @param    *string        page             parameter: Optional
 * @param    *string        pageSize         parameter: Optional
 * @param    *string        to               parameter: Optional
 * @param    *string        from             parameter: Optional
 * @param    *string        dateCreated      parameter: Optional
 * @param    *string        responseType     parameter: Optional
 * @return	Returns the  response from the API call
 */
func (me *CALL_IMPL) CreateListCalls (
            page *string,
            pageSize *string,
            to *string,
            from *string,
            dateCreated *string,
            responseType *string) (error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/listcalls.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : apihelper_pkg.ToString(*responseType, "json"),
    })
    if err != nil {
        //error in template param handling
        return err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "message360-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Page" : page,
        "PageSize" : pageSize,
        "To" : to,
        "From" : from,
        "DateCreated" : dateCreated,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, message360_lib.Config.BasicAuthUserName, message360_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil
}

