/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package call_pkg


import(
	"ytel_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"ytel_lib/apihelper_pkg"
	"ytel_lib/configuration_pkg"
)

/*
 * Input structure for the method MakeCall
 */
type MakeCallInput struct {
    From                  string          //A valid Ytel Voice enabled number (E.164 format) that will be initiating the phone call.
    To                    string          //To number
    Url                   string          //URL requested once the call connects
    ResponseType          string          //Response type format xml or json
    Method                models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once call connects.
    StatusCallBackUrl     *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished.
    StatusCallBackMethod  models_pkg.HttpActionEnum //Specifies the HTTP methodlinkclass used to request StatusCallbackUrl.
    FallBackUrl           *string         //URL requested if the initial Url parameter fails or encounters an error
    FallBackMethod        models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required FallbackUrl once call connects.
    HeartBeatUrl          *string         //URL that can be requested every 60 seconds during the call to notify of elapsed tim
    HeartBeatMethod       models_pkg.HttpActionEnum //Specifies the HTTP method used to request HeartbeatUrl.
    Timeout               *int64          //Time (in seconds) Ytel should wait while the call is ringing before canceling the call
    PlayDtmf              *string         //DTMF Digits to play to the call once it connects. 0-9, #, or *
    HideCallerId          *bool           //Specifies if the caller id will be hidden
    Record                *bool           //Specifies if the call should be recorded
    RecordCallBackUrl     *string         //Recording parameters will be sent here upon completion
    RecordCallBackMethod  models_pkg.HttpActionEnum //Method used to request the RecordCallback URL.
    Transcribe            *bool           //Specifies if the call recording should be transcribed
    TranscribeCallBackUrl *string         //Transcription parameters will be sent here upon completion
    IfMachine             models_pkg.IfMachineEnum //How Ytel should handle the receiving numbers voicemail machine
    IfMachineUrl          *string         //URL requested when IfMachine=continue
    IfMachineMethod       models_pkg.HttpActionEnum //Method used to request the IfMachineUrl.
    Feedback              *bool           //Specify if survey should be enable or not
    SurveyId              *string         //The unique identifier for the survey.
}

/*
 * Input structure for the method PlayAudio
 */
type PlayAudioInput struct {
    CallSid         string          //The unique identifier of each call resource
    AudioUrl        string          //URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav
    SayText         string          //Valid alphanumeric string that should be played to the In-progress call.
    ResponseType    string          //Response type format xml or json
    Length          *int64          //Time limit in seconds for audio play back
    Direction       models_pkg.DirectionEnum //The leg of the call audio will be played to
    Mix             *bool           //If false, all other audio will be muted
}

/*
 * Input structure for the method RecordCall
 */
type RecordCallInput struct {
    CallSid         string          //The unique identifier of each call resource
    Record          bool            //Set true to initiate recording or false to terminate recording
    ResponseType    string          //Response format, xml or json
    Direction       models_pkg.DirectionEnum //The leg of the call to record
    TimeLimit       *int64          //Time in seconds the recording duration should not exceed
    CallBackUrl     *string         //URL consulted after the recording completes
    Fileformat      models_pkg.AudioFormatEnum //Format of the recording file. Can be .mp3 or .wav
}

/*
 * Input structure for the method VoiceEffect
 */
type VoiceEffectInput struct {
    CallSid         string          //The unique identifier for the in-progress voice call.
    ResponseType    string          //Response type format xml or json
    AudioDirection  models_pkg.AudioDirectionEnum //The direction the audio effect should be placed on. If IN, the effects will occur on the incoming audio stream. If OUT, the effects will occur on the outgoing audio stream.
    PitchSemiTones  *float64        //Set the pitch in semitone (half-step) intervals. Value between -14 and 14
    PitchOctaves    *float64        //Set the pitch in octave intervals.. Value between -1 and 1
    Pitch           *float64        //Set the pitch (lowness/highness) of the audio. The higher the value, the higher the pitch. Value greater than 0
    Rate            *float64        //Set the rate for audio. The lower the value, the lower the rate. value greater than 0
    Tempo           *float64        //Set the tempo (speed) of the audio. A higher value denotes a faster tempo. Value greater than 0
}

/*
 * Input structure for the method SendDigit
 */
type SendDigitInput struct {
    CallSid           string          //The unique identifier of each call resource
    PlayDtmf          string          //DTMF digits to play to the call. 0-9, #, *, W, or w
    ResponseType      string          //Response type format xml or json
    PlayDtmfDirection models_pkg.DirectionEnum //The leg of the call DTMF digits should be sent to
}

/*
 * Input structure for the method InterruptedCall
 */
type InterruptedCallInput struct {
    CallSid         string          //The unique identifier for voice call that is in progress.
    ResponseType    string          //Response type format xml or json
    Url             *string         //URL the in-progress call will be redirected to
    Method          models_pkg.HttpActionEnum //The method used to request the above Url parameter
    Status          models_pkg.InterruptedCallStatusEnum //Status to set the in-progress call to
}

/*
 * Input structure for the method ListCalls
 */
type ListCallsInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    PageSize        *int64          //Number of individual resources listed in the response per page
    To              *string         //Filter calls that were sent to this 10-digit number (E.164 format).
    From            *string         //Filter calls that were sent from this 10-digit number (E.164 format).
    DateCreated     *string         //Return calls that are from a specified date.
}

/*
 * Input structure for the method SendRinglessVM
 */
type SendRinglessVMInput struct {
    From                string          //A valid Ytel Voice enabled number (E.164 format) that will be initiating the phone call.
    RVMCallerId         string          //A required secondary Caller ID for RVM to work.
    To                  string          //A valid number (E.164 format) that will receive the phone call.
    VoiceMailURL        string          //The URL requested once the RVM connects. A set of default parameters will be sent here.
    ResponseType        string          //Response type format xml or json
    Method              models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once call connects.
    StatusCallBackUrl   *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished.
    StatsCallBackMethod models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required StatusCallBackUrl once call connects.
}

/*
 * Input structure for the method ViewCall
 */
type ViewCallInput struct {
    Callsid         string          //The unique identifier for the voice call.
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method GroupCall
 */
type GroupCallInput struct {
    From                  string          //This number to display on Caller ID as calling
    To                    string          //Please enter multiple E164 number. You can add max 10 numbers. Add numbers separated with comma. e.g : 1111111111,2222222222
    Url                   string          //URL requested once the call connects
    ResponseType          string          //TODO: Write general description for this field
    GroupConfirmKey       string          //Define the DTMF that the called party should send to bridge the call. Allowed Values : 0-9, #, *
    GroupConfirmFile      models_pkg.AudioFormatEnum //Specify the audio file you want to play when the called party picks up the call
    Method                models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once call connects.
    StatusCallBackUrl     *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished.
    StatusCallBackMethod  models_pkg.HttpActionEnum //Specifies the HTTP methodlinkclass used to request StatusCallbackUrl.
    FallBackUrl           *string         //URL requested if the initial Url parameter fails or encounters an error
    FallBackMethod        models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required FallbackUrl once call connects.
    HeartBeatUrl          *string         //URL that can be requested every 60 seconds during the call to notify of elapsed time and pass other general information.
    HeartBeatMethod       models_pkg.HttpActionEnum //Specifies the HTTP method used to request HeartbeatUrl.
    Timeout               *int64          //Time (in seconds) we should wait while the call is ringing before canceling the call
    PlayDtmf              *string         //DTMF Digits to play to the call once it connects. 0-9, #, or *
    HideCallerId          *string         //Specifies if the caller id will be hidden
    Record                *bool           //Specifies if the call should be recorded
    RecordCallBackUrl     *string         //Recording parameters will be sent here upon completion
    RecordCallBackMethod  models_pkg.HttpActionEnum //Method used to request the RecordCallback URL.
    Transcribe            *bool           //Specifies if the call recording should be transcribed
    TranscribeCallBackUrl *string         //Transcription parameters will be sent here upon completion
}

/*
 * Client structure as interface implementation
 */
type CALL_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * You can experiment with initiating a call through Ytel and view the request response generated when doing so and get the response in json
 * @param  MakeCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) MakeCall (input *MakeCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/makecall.{ResponseType}"

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

        "From" : input.From,
        "To" : input.To,
        "Url" : input.Url,
        "Method" : models_pkg.HttpActionEnumToValue(input.Method),
        "StatusCallBackUrl" : input.StatusCallBackUrl,
        "StatusCallBackMethod" : models_pkg.HttpActionEnumToValue(input.StatusCallBackMethod),
        "FallBackUrl" : input.FallBackUrl,
        "FallBackMethod" : models_pkg.HttpActionEnumToValue(input.FallBackMethod),
        "HeartBeatUrl" : input.HeartBeatUrl,
        "HeartBeatMethod" : models_pkg.HttpActionEnumToValue(input.HeartBeatMethod),
        "Timeout" : input.Timeout,
        "PlayDtmf" : input.PlayDtmf,
        "HideCallerId" : input.HideCallerId,
        "Record" : input.Record,
        "RecordCallBackUrl" : input.RecordCallBackUrl,
        "RecordCallBackMethod" : models_pkg.HttpActionEnumToValue(input.RecordCallBackMethod),
        "Transcribe" : input.Transcribe,
        "TranscribeCallBackUrl" : input.TranscribeCallBackUrl,
        "IfMachine" : models_pkg.IfMachineEnumToValue(input.IfMachine),
        "IfMachineUrl" : input.IfMachineUrl,
        "IfMachineMethod" : models_pkg.HttpActionEnumToValue(input.IfMachineMethod),
        "Feedback" : input.Feedback,
        "SurveyId" : input.SurveyId,

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
 * @param  PlayAudioInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) PlayAudio (input *PlayAudioInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/playaudios.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "AudioUrl" : input.AudioUrl,
        "SayText" : input.SayText,
        "Length" : input.Length,
        "Direction" : models_pkg.DirectionEnumToValue(input.Direction),
        "Mix" : input.Mix,

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
 * @param  RecordCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) RecordCall (input *RecordCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/recordcalls.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "Record" : input.Record,
        "Direction" : models_pkg.DirectionEnumToValue(input.Direction),
        "TimeLimit" : input.TimeLimit,
        "CallBackUrl" : input.CallBackUrl,
        "Fileformat" : models_pkg.AudioFormatEnumToValue(input.Fileformat),

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
 * @param  VoiceEffectInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) VoiceEffect (input *VoiceEffectInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/voiceeffect.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "AudioDirection" : models_pkg.AudioDirectionEnumToValue(input.AudioDirection),
        "PitchSemiTones" : input.PitchSemiTones,
        "PitchOctaves" : input.PitchOctaves,
        "Pitch" : input.Pitch,
        "Rate" : input.Rate,
        "Tempo" : input.Tempo,

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
 * @param  SendDigitInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) SendDigit (input *SendDigitInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/senddigits.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "PlayDtmf" : input.PlayDtmf,
        "PlayDtmfDirection" : models_pkg.DirectionEnumToValue(input.PlayDtmfDirection),

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
 * @param  InterruptedCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) InterruptedCall (input *InterruptedCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/interruptcalls.{ResponseType}"

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

        "CallSid" : input.CallSid,
        "Url" : input.Url,
        "Method" : models_pkg.HttpActionEnumToValue(input.Method),
        "Status" : models_pkg.InterruptedCallStatusEnumToValue(input.Status),

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
 * @param  ListCallsInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) ListCalls (input *ListCallsInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/listcalls.{ResponseType}"

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
        "To" : input.To,
        "From" : input.From,
        "DateCreated" : input.DateCreated,

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
 * @param  SendRinglessVMInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) SendRinglessVM (input *SendRinglessVMInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/makervm.{ResponseType}"

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

        "From" : input.From,
        "RVMCallerId" : input.RVMCallerId,
        "To" : input.To,
        "VoiceMailURL" : input.VoiceMailURL,
        "Method" : models_pkg.HttpActionEnumToValue(input.Method),
        "StatusCallBackUrl" : input.StatusCallBackUrl,
        "StatsCallBackMethod" : models_pkg.HttpActionEnumToValue(input.StatsCallBackMethod),

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
 * @param  ViewCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) ViewCall (input *ViewCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/viewcalls.{ResponseType}"

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

        "callsid" : input.Callsid,

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
 * @param    string        callSid          parameter: Required
 * @param    string        responseType     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) ViewCallDetail (
            callSid string,
            responseType string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/viewcalldetail.{ResponseType}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "ResponseType" : responseType,
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
 * Group Call
 * @param  GroupCallInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CALL_IMPL) GroupCall (input *GroupCallInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/calls/groupcall.{ResponseType}"

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

        "From" : input.From,
        "To" : input.To,
        "Url" : input.Url,
        "GroupConfirmKey" : input.GroupConfirmKey,
        "GroupConfirmFile" : models_pkg.AudioFormatEnumToValue(input.GroupConfirmFile),
        "Method" : models_pkg.HttpActionEnumToValue(input.Method),
        "StatusCallBackUrl" : input.StatusCallBackUrl,
        "StatusCallBackMethod" : models_pkg.HttpActionEnumToValue(input.StatusCallBackMethod),
        "FallBackUrl" : input.FallBackUrl,
        "FallBackMethod" : models_pkg.HttpActionEnumToValue(input.FallBackMethod),
        "HeartBeatUrl" : input.HeartBeatUrl,
        "HeartBeatMethod" : models_pkg.HttpActionEnumToValue(input.HeartBeatMethod),
        "Timeout" : input.Timeout,
        "PlayDtmf" : input.PlayDtmf,
        "HideCallerId" : input.HideCallerId,
        "Record" : input.Record,
        "RecordCallBackUrl" : input.RecordCallBackUrl,
        "RecordCallBackMethod" : models_pkg.HttpActionEnumToValue(input.RecordCallBackMethod),
        "Transcribe" : input.Transcribe,
        "TranscribeCallBackUrl" : input.TranscribeCallBackUrl,

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

