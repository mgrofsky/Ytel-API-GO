/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package call_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method CreateViewCall
 */
type CreateViewCallInput struct {
    Callsid         string          //Call Sid id for particular Call
    ResponseType    *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateGroupCall
 */
type CreateGroupCallInput struct {
    FromCountryCode       string          //TODO: Write general description for this field
    From                  string          //TODO: Write general description for this field
    ToCountryCode         string          //TODO: Write general description for this field
    To                    string          //TODO: Write general description for this field
    Url                   string          //TODO: Write general description for this field
    Method                models_pkg.HttpAction //TODO: Write general description for this field
    StatusCallBackUrl     *string         //TODO: Write general description for this field
    StatusCallBackMethod  models_pkg.HttpAction //TODO: Write general description for this field
    FallBackUrl           *string         //TODO: Write general description for this field
    FallBackMethod        models_pkg.HttpAction //TODO: Write general description for this field
    HeartBeatUrl          *string         //TODO: Write general description for this field
    HeartBeatMethod       models_pkg.HttpAction //TODO: Write general description for this field
    Timeout               *int64          //TODO: Write general description for this field
    PlayDtmf              *string         //TODO: Write general description for this field
    HideCallerId          *string         //TODO: Write general description for this field
    Record                *bool           //TODO: Write general description for this field
    RecordCallBackUrl     *string         //TODO: Write general description for this field
    RecordCallBackMethod  models_pkg.HttpAction //TODO: Write general description for this field
    Transcribe            *bool           //TODO: Write general description for this field
    TranscribeCallBackUrl *string         //TODO: Write general description for this field
    ResponseType          *string         //TODO: Write general description for this field
}

/*
 * Input structure for the method CreateVoiceEffect
 */
type CreateVoiceEffectInput struct {
    CallSid         string          //TODO: Write general description for this field
    AudioDirection  models_pkg.AudioDirection //TODO: Write general description for this field
    PitchSemiTones  *float64        //value between -14 and 14
    PitchOctaves    *float64        //value between -1 and 1
    Pitch           *float64        //value greater than 0
    Rate            *float64        //value greater than 0
    Tempo           *float64        //value greater than 0
    ResponseType    *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateRecordCall
 */
type CreateRecordCallInput struct {
    CallSid         string          //The unique identifier of each call resource
    Record          bool            //Set true to initiate recording or false to terminate recording
    Direction       models_pkg.Direction //The leg of the call to record
    TimeLimit       *int64          //Time in seconds the recording duration should not exceed
    CallBackUrl     *string         //URL consulted after the recording completes
    Fileformat      models_pkg.AudioFormat //Format of the recording file. Can be .mp3 or .wav
    ResponseType    *string         //Response format, xml or json
}

/*
 * Input structure for the method CreatePlayAudio
 */
type CreatePlayAudioInput struct {
    CallSid         string          //The unique identifier of each call resource
    AudioUrl        string          //URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav
    Length          *int64          //Time limit in seconds for audio play back
    Direction       models_pkg.Direction //The leg of the call audio will be played to
    Loop            *bool           //Repeat audio playback indefinitely
    Mix             *bool           //If false, all other audio will be muted
    ResponseType    *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateInterruptedCall
 */
type CreateInterruptedCallInput struct {
    CallSid         string          //Call SId
    Url             *string         //URL the in-progress call will be redirected to
    Method          models_pkg.HttpAction //The method used to request the above Url parameter
    Status          models_pkg.InterruptedCallStatus //Status to set the in-progress call to
    ResponseType    *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateSendDigit
 */
type CreateSendDigitInput struct {
    CallSid           string          //The unique identifier of each call resource
    PlayDtmf          string          //DTMF digits to play to the call. 0-9, #, *, W, or w
    PlayDtmfDirection models_pkg.Direction //The leg of the call DTMF digits should be sent to
    ResponseType      *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateMakeCall
 */
type CreateMakeCallInput struct {
    FromCountryCode       string          //from country code
    From                  string          //This number to display on Caller ID as calling
    ToCountryCode         string          //To cuntry code number
    To                    string          //To number
    Url                   string          //URL requested once the call connects
    Method                models_pkg.HttpAction //Specifies the HTTP method used to request the required URL once call connects.
    StatusCallBackUrl     *string         //specifies the HTTP methodlinkclass used to request StatusCallbackUrl.
    StatusCallBackMethod  models_pkg.HttpAction //Specifies the HTTP methodlinkclass used to request StatusCallbackUrl.
    FallBackUrl           *string         //URL requested if the initial Url parameter fails or encounters an error
    FallBackMethod        models_pkg.HttpAction //Specifies the HTTP method used to request the required FallbackUrl once call connects.
    HeartBeatUrl          *string         //URL that can be requested every 60 seconds during the call to notify of elapsed tim
    HeartBeatMethod       *bool           //Specifies the HTTP method used to request HeartbeatUrl.
    Timeout               *int64          //Time (in seconds) Message360 should wait while the call is ringing before canceling the call
    PlayDtmf              *string         //DTMF Digits to play to the call once it connects. 0-9, #, or *
    HideCallerId          *bool           //Specifies if the caller id will be hidden
    Record                *bool           //Specifies if the call should be recorded
    RecordCallBackUrl     *string         //Recording parameters will be sent here upon completion
    RecordCallBackMethod  models_pkg.HttpAction //Method used to request the RecordCallback URL.
    Transcribe            *bool           //Specifies if the call recording should be transcribed
    TranscribeCallBackUrl *string         //Transcription parameters will be sent here upon completion
    IfMachine             models_pkg.IfMachine //How Message360 should handle the receiving numbers voicemail machine
    ResponseType          *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateListCalls
 */
type CreateListCallsInput struct {
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    PageSize        *int64          //Number of individual resources listed in the response per page
    To              *string         //Only list calls to this number
    From            *string         //Only list calls from this number
    DateCreated     *string         //Only list calls starting within the specified date range
    ResponseType    *string         //Response type format xml or json
}

/*
 * Client structure as interface implementation
 */
type CALL_IMPL struct { }

/**
 * View Call Response
 * @param  CreateViewCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateViewCall (input *CreateViewCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/viewcalls.{ResponseType}"

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

        "callsid" : input.Callsid,

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
 * Group Call
 * @param  CreateGroupCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateGroupCall (input *CreateGroupCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/groupcall.{ResponseType}"

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

        "FromCountryCode" : input.FromCountryCode,
        "From" : input.From,
        "ToCountryCode" : input.ToCountryCode,
        "To" : input.To,
        "Url" : input.Url,
        "Method" : models_pkg.HttpActionToValue(input.Method),
        "StatusCallBackUrl" : input.StatusCallBackUrl,
        "StatusCallBackMethod" : models_pkg.HttpActionToValue(input.StatusCallBackMethod),
        "FallBackUrl" : input.FallBackUrl,
        "FallBackMethod" : models_pkg.HttpActionToValue(input.FallBackMethod),
        "HeartBeatUrl" : input.HeartBeatUrl,
        "HeartBeatMethod" : models_pkg.HttpActionToValue(input.HeartBeatMethod),
        "Timeout" : input.Timeout,
        "PlayDtmf" : input.PlayDtmf,
        "HideCallerId" : input.HideCallerId,
        "Record" : input.Record,
        "RecordCallBackUrl" : input.RecordCallBackUrl,
        "RecordCallBackMethod" : models_pkg.HttpActionToValue(input.RecordCallBackMethod),
        "Transcribe" : input.Transcribe,
        "TranscribeCallBackUrl" : input.TranscribeCallBackUrl,

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
 * @param  CreateVoiceEffectInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateVoiceEffect (input *CreateVoiceEffectInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/voiceeffect.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "AudioDirection" : models_pkg.AudioDirectionToValue(input.AudioDirection),
        "PitchSemiTones" : input.PitchSemiTones,
        "PitchOctaves" : input.PitchOctaves,
        "Pitch" : input.Pitch,
        "Rate" : input.Rate,
        "Tempo" : input.Tempo,

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
 * @param  CreateRecordCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateRecordCall (input *CreateRecordCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/recordcalls.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "Record" : input.Record,
        "Direction" : models_pkg.DirectionToValue(input.Direction),
        "TimeLimit" : input.TimeLimit,
        "CallBackUrl" : input.CallBackUrl,
        "Fileformat" : models_pkg.AudioFormatToValue(input.Fileformat),

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
 * @param  CreatePlayAudioInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreatePlayAudio (input *CreatePlayAudioInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/playaudios.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "AudioUrl" : input.AudioUrl,
        "Length" : input.Length,
        "Direction" : models_pkg.DirectionToValue(input.Direction),
        "Loop" : input.Loop,
        "Mix" : input.Mix,

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
 * @param  CreateInterruptedCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateInterruptedCall (input *CreateInterruptedCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/interruptcalls.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "Url" : input.Url,
        "Method" : models_pkg.HttpActionToValue(input.Method),
        "Status" : models_pkg.InterruptedCallStatusToValue(input.Status),

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
 * @param  CreateSendDigitInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateSendDigit (input *CreateSendDigitInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/senddigits.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "PlayDtmf" : input.PlayDtmf,
        "PlayDtmfDirection" : models_pkg.DirectionToValue(input.PlayDtmfDirection),

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
 * @param  CreateMakeCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateMakeCall (input *CreateMakeCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/makecall.{ResponseType}"

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

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "Method" : models_pkg.HttpActionToValue(input.Method),
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

        "FromCountryCode" : input.FromCountryCode,
        "From" : input.From,
        "ToCountryCode" : input.ToCountryCode,
        "To" : input.To,
        "Url" : input.Url,
        "StatusCallBackUrl" : input.StatusCallBackUrl,
        "StatusCallBackMethod" : models_pkg.HttpActionToValue(input.StatusCallBackMethod),
        "FallBackUrl" : input.FallBackUrl,
        "FallBackMethod" : models_pkg.HttpActionToValue(input.FallBackMethod),
        "HeartBeatUrl" : input.HeartBeatUrl,
        "HeartBeatMethod" : input.HeartBeatMethod,
        "Timeout" : input.Timeout,
        "PlayDtmf" : input.PlayDtmf,
        "HideCallerId" : input.HideCallerId,
        "Record" : input.Record,
        "RecordCallBackUrl" : input.RecordCallBackUrl,
        "RecordCallBackMethod" : models_pkg.HttpActionToValue(input.RecordCallBackMethod),
        "Transcribe" : input.Transcribe,
        "TranscribeCallBackUrl" : input.TranscribeCallBackUrl,
        "IfMachine" : models_pkg.IfMachineToValue(input.IfMachine),

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
 * @param  CreateListCallsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) CreateListCalls (input *CreateListCallsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/listcalls.{ResponseType}"

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

        "Page" : input.Page,
        "PageSize" : apihelper_pkg.ToString(*input.PageSize, "10"),
        "To" : input.To,
        "From" : input.From,
        "DateCreated" : input.DateCreated,

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

