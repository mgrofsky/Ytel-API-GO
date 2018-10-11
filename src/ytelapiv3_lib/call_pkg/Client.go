/*
 * ytelapiv3_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package call_pkg


import(
	"fmt"
	"ytelapiv3_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"ytelapiv3_lib/apihelper_pkg"
	"ytelapiv3_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type CALL_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Retrieve a single voice call’s information by its CallSid.
 * @param    string        callSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateViewCall1 (
            callSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/viewcalldetail.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "callSid" : callSid,

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
 * Retrieve a single voice call’s information by its CallSid.
 * @param    string        callsid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateViewCall (
            callsid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/viewcalls.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "callsid" : callsid,

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
 * Play Dtmf and send the Digit
 * @param    string                                  callSid               parameter: Required
 * @param    string                                  playDtmf              parameter: Required
 * @param    models_pkg.PlayDtmfDirectionEnum        playDtmfDirection     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreatePlayDTMF (
            callSid string,
            playDtmf string,
            playDtmfDirection models_pkg.PlayDtmfDirectionEnum) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/senddigits.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "CallSid" : callSid,
        "PlayDtmf" : playDtmf,
        "PlayDtmfDirection" : models_pkg.PlayDtmfDirectionEnumToValue(playDtmfDirection),

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
 * You can experiment with initiating a call through Ytel and view the request response generated when doing so and get the response in json
 * @param    string                          from                      parameter: Required
 * @param    string                          to                        parameter: Required
 * @param    string                          url                       parameter: Required
 * @param    *string                         method                    parameter: Optional
 * @param    *string                         statusCallBackUrl         parameter: Optional
 * @param    *string                         statusCallBackMethod      parameter: Optional
 * @param    *string                         fallBackUrl               parameter: Optional
 * @param    *string                         fallBackMethod            parameter: Optional
 * @param    *string                         heartBeatUrl              parameter: Optional
 * @param    *string                         heartBeatMethod           parameter: Optional
 * @param    *int64                          timeout                   parameter: Optional
 * @param    *string                         playDtmf                  parameter: Optional
 * @param    *bool                           hideCallerId              parameter: Optional
 * @param    *bool                           record                    parameter: Optional
 * @param    *string                         recordCallBackUrl         parameter: Optional
 * @param    *string                         recordCallBackMethod      parameter: Optional
 * @param    *bool                           transcribe                parameter: Optional
 * @param    *string                         transcribeCallBackUrl     parameter: Optional
 * @param    models_pkg.IfMachineEnum        ifMachine                 parameter: Optional
 * @param    *string                         ifMachineUrl              parameter: Optional
 * @param    *string                         ifMachineMethod           parameter: Optional
 * @param    *bool                           feedback                  parameter: Optional
 * @param    *string                         surveyId                  parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateMakeCall (
            from string,
            to string,
            url string,
            method *string,
            statusCallBackUrl *string,
            statusCallBackMethod *string,
            fallBackUrl *string,
            fallBackMethod *string,
            heartBeatUrl *string,
            heartBeatMethod *string,
            timeout *int64,
            playDtmf *string,
            hideCallerId *bool,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod *string,
            transcribe *bool,
            transcribeCallBackUrl *string,
            ifMachine models_pkg.IfMachineEnum,
            ifMachineUrl *string,
            ifMachineMethod *string,
            feedback *bool,
            surveyId *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/makecall.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "From" : from,
        "To" : to,
        "Url" : url,
        "Method" : method,
        "StatusCallBackUrl" : statusCallBackUrl,
        "StatusCallBackMethod" : statusCallBackMethod,
        "FallBackUrl" : fallBackUrl,
        "FallBackMethod" : fallBackMethod,
        "HeartBeatUrl" : heartBeatUrl,
        "HeartBeatMethod" : heartBeatMethod,
        "Timeout" : timeout,
        "PlayDtmf" : playDtmf,
        "HideCallerId" : hideCallerId,
        "Record" : record,
        "RecordCallBackUrl" : recordCallBackUrl,
        "RecordCallBackMethod" : recordCallBackMethod,
        "Transcribe" : transcribe,
        "TranscribeCallBackUrl" : transcribeCallBackUrl,
        "IfMachine" : models_pkg.IfMachineEnumToValue(ifMachine),
        "IfMachineUrl" : ifMachineUrl,
        "IfMachineMethod" : ifMachineMethod,
        "Feedback" : feedback,
        "SurveyId" : surveyId,

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
 * Play Audio from a url
 * @param    string                          callSid       parameter: Required
 * @param    string                          audioUrl      parameter: Required
 * @param    string                          sayText       parameter: Required
 * @param    *int64                          length        parameter: Optional
 * @param    models_pkg.DirectionEnum        direction     parameter: Optional
 * @param    *bool                           mix           parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreatePlayAudio (
            callSid string,
            audioUrl string,
            sayText string,
            length *int64,
            direction models_pkg.DirectionEnum,
            mix *bool) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/playaudios.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "CallSid" : callSid,
        "AudioUrl" : audioUrl,
        "SayText" : sayText,
        "Length" : length,
        "Direction" : models_pkg.DirectionEnumToValue(direction),
        "Mix" : mix,

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
 * Start or stop recording of an in-progress voice call.
 * @param    string                           callSid         parameter: Required
 * @param    bool                             record          parameter: Required
 * @param    models_pkg.Direction4Enum        direction       parameter: Optional
 * @param    *int64                           timeLimit       parameter: Optional
 * @param    *string                          callBackUrl     parameter: Optional
 * @param    models_pkg.FileformatEnum        fileformat      parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateRecordCall (
            callSid string,
            record bool,
            direction models_pkg.Direction4Enum,
            timeLimit *int64,
            callBackUrl *string,
            fileformat models_pkg.FileformatEnum) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/recordcalls.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "CallSid" : callSid,
        "Record" : record,
        "Direction" : models_pkg.Direction4EnumToValue(direction),
        "TimeLimit" : timeLimit,
        "CallBackUrl" : callBackUrl,
        "Fileformat" : models_pkg.FileformatEnumToValue(fileformat),

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
 * Add audio voice effects to the an in-progress voice call.
 * @param    string                               callSid            parameter: Required
 * @param    models_pkg.AudioDirectionEnum        audioDirection     parameter: Optional
 * @param    *float64                             pitchSemiTones     parameter: Optional
 * @param    *float64                             pitchOctaves       parameter: Optional
 * @param    *float64                             pitch              parameter: Optional
 * @param    *float64                             rate               parameter: Optional
 * @param    *float64                             tempo              parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateVoiceEffect (
            callSid string,
            audioDirection models_pkg.AudioDirectionEnum,
            pitchSemiTones *float64,
            pitchOctaves *float64,
            pitch *float64,
            rate *float64,
            tempo *float64) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/voiceeffect.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "CallSid" : callSid,
        "AudioDirection" : models_pkg.AudioDirectionEnumToValue(audioDirection),
        "PitchSemiTones" : pitchSemiTones,
        "PitchOctaves" : pitchOctaves,
        "Pitch" : pitch,
        "Rate" : rate,
        "Tempo" : tempo,

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
 * Interrupt the Call by Call Sid
 * @param    string                       callSid     parameter: Required
 * @param    *string                      url         parameter: Optional
 * @param    *string                      method      parameter: Optional
 * @param    models_pkg.StatusEnum        status      parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateInterruptCall (
            callSid string,
            url *string,
            method *string,
            status models_pkg.StatusEnum) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/interruptcalls.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "CallSid" : callSid,
        "Url" : url,
        "Method" : method,
        "Status" : models_pkg.StatusEnumToValue(status),

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
 * A list of calls associated with your Ytel account
 * @param    *int64         page            parameter: Optional
 * @param    *int64         pageSize        parameter: Optional
 * @param    *string        to              parameter: Optional
 * @param    *string        from            parameter: Optional
 * @param    *string        dateCreated     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateListCalls (
            page *int64,
            pageSize *int64,
            to *string,
            from *string,
            dateCreated *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/listcalls.json"

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
        "user-agent" : "APIMATIC 2.0",
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
 * Initiate an outbound Ringless Voicemail through Ytel.
 * @param    string         from                    parameter: Required
 * @param    string         rVMCallerId             parameter: Required
 * @param    string         to                      parameter: Required
 * @param    string         voiceMailURL            parameter: Required
 * @param    *string        method                  parameter: Optional
 * @param    *string        statusCallBackUrl       parameter: Optional
 * @param    *string        statsCallBackMethod     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateSendRVM (
            from string,
            rVMCallerId string,
            to string,
            voiceMailURL string,
            method *string,
            statusCallBackUrl *string,
            statsCallBackMethod *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/makervm.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "From" : from,
        "RVMCallerId" : rVMCallerId,
        "To" : to,
        "VoiceMailURL" : voiceMailURL,
        "Method" : method,
        "StatusCallBackUrl" : statusCallBackUrl,
        "StatsCallBackMethod" : statsCallBackMethod,

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
 * Group Call
 * @param    string                                 from                      parameter: Required
 * @param    string                                 to                        parameter: Required
 * @param    string                                 url                       parameter: Required
 * @param    string                                 groupConfirmKey           parameter: Required
 * @param    models_pkg.GroupConfirmFileEnum        groupConfirmFile          parameter: Required
 * @param    *string                                method                    parameter: Optional
 * @param    *string                                statusCallBackUrl         parameter: Optional
 * @param    *string                                statusCallBackMethod      parameter: Optional
 * @param    *string                                fallBackUrl               parameter: Optional
 * @param    *string                                fallBackMethod            parameter: Optional
 * @param    *string                                heartBeatUrl              parameter: Optional
 * @param    *string                                heartBeatMethod           parameter: Optional
 * @param    *int64                                 timeout                   parameter: Optional
 * @param    *string                                playDtmf                  parameter: Optional
 * @param    *string                                hideCallerId              parameter: Optional
 * @param    *bool                                  record                    parameter: Optional
 * @param    *string                                recordCallBackUrl         parameter: Optional
 * @param    *string                                recordCallBackMethod      parameter: Optional
 * @param    *bool                                  transcribe                parameter: Optional
 * @param    *string                                transcribeCallBackUrl     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateGroupCall (
            from string,
            to string,
            url string,
            groupConfirmKey string,
            groupConfirmFile models_pkg.GroupConfirmFileEnum,
            method *string,
            statusCallBackUrl *string,
            statusCallBackMethod *string,
            fallBackUrl *string,
            fallBackMethod *string,
            heartBeatUrl *string,
            heartBeatMethod *string,
            timeout *int64,
            playDtmf *string,
            hideCallerId *string,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod *string,
            transcribe *bool,
            transcribeCallBackUrl *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/groupcall.json"

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
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "From" : from,
        "To" : to,
        "Url" : url,
        "GroupConfirmKey" : groupConfirmKey,
        "GroupConfirmFile" : models_pkg.GroupConfirmFileEnumToValue(groupConfirmFile),
        "Method" : method,
        "StatusCallBackUrl" : statusCallBackUrl,
        "StatusCallBackMethod" : statusCallBackMethod,
        "FallBackUrl" : fallBackUrl,
        "FallBackMethod" : fallBackMethod,
        "HeartBeatUrl" : heartBeatUrl,
        "HeartBeatMethod" : heartBeatMethod,
        "Timeout" : timeout,
        "PlayDtmf" : playDtmf,
        "HideCallerId" : hideCallerId,
        "Record" : record,
        "RecordCallBackUrl" : recordCallBackUrl,
        "RecordCallBackMethod" : recordCallBackMethod,
        "Transcribe" : transcribe,
        "TranscribeCallBackUrl" : transcribeCallBackUrl,

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
