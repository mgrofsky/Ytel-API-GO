/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package conference_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method DeafMuteParticipant
 */
type DeafMuteParticipantInput struct {
    ConferenceSid   string          //ID of the active conference
    ParticipantSid  string          //ID of an active participant
    ResponseType    string          //Response Type either json or xml
    Muted           *bool           //Mute a participant
    Deaf            *bool           //Make it so a participant cant hear
}

/*
 * Input structure for the method ViewParticipant
 */
type ViewParticipantInput struct {
    ConferenceSid   string          //unique conference sid
    ParticipantSid  string          //TODO: Write general description for this field
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method AddParticipant
 */
type AddParticipantInput struct {
    Conferencesid     string          //Unique Conference Sid
    Participantnumber string          //Particiant Number
    ResponseType      string          //Response type format xml or json
    Muted             *bool           //add muted
    Deaf              *bool           //add without volume
}

/*
 * Input structure for the method ViewConference
 */
type ViewConferenceInput struct {
    Conferencesid   string          //The unique identifier of each conference resource
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method CreateConference
 */
type CreateConferenceInput struct {
    From                 string          //This number to display on Caller ID as calling
    To                   string          //To number
    Url                  string          //URL requested once the call connects
    Method               models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once call connects.
    RecordCallbackUrl    string          //Recording parameters will be sent here upon completion.
    ResponseType         string          //Response type format xml or json
    StatusCallBackUrl    *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the conference is finished.
    StatusCallBackMethod models_pkg.HttpActionEnum //Specifies the HTTP methodlinkclass used to request StatusCallbackUrl.
    FallBackUrl          *string         //URL requested if the initial Url parameter fails or encounters an error
    FallBackMethod       models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required FallbackUrl once call connects.
    Record               *bool           //Specifies if the conference should be recorded.
    RecordCallbackMethod models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once conference connects.
    SchdeuleTime         *string         //Schedule conference in future. Schedule time must be greater than current time
    Timeout              *int64          //The number of seconds the call stays on the line while waiting for an answer. The max time limit is 999 and the default limit is 60 seconds but lower times can be set.
}

/*
 * Input structure for the method HangupParticipant
 */
type HangupParticipantInput struct {
    ConferenceSid   string          //The unique identifier for a conference object.
    ParticipantSid  string          //The unique identifier for a participant object.
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method PlayConferenceAudio
 */
type PlayConferenceAudioInput struct {
    ConferenceSid   string          //The unique identifier for a conference object.
    ParticipantSid  string          //The unique identifier for a participant object.
    AudioUrl        models_pkg.AudioFormatEnum //The URL for the audio file that is to be played during the conference. Multiple audio files can be chained by using a semicolon.
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ListParticipant
 */
type ListParticipantInput struct {
    ConferenceSid   string          //unique conference sid
    ResponseType    string          //Response format, xml or json
    Page            *int64          //page number
    Pagesize        *int64          //Amount of records to return per page
    Muted           *bool           //Participants that are muted
    Deaf            *bool           //Participants cant hear
}

/*
 * Input structure for the method ListConference
 */
type ListConferenceInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    PageSize        *int64          //Number of individual resources listed in the response per page
    FriendlyName    *string         //Only return conferences with the specified FriendlyName
    DateCreated     *string         //Conference created date
}

/*
 * Client structure as interface implementation
 */
type CONFERENCE_IMPL struct { }

/**
 * Deaf Mute Participant
 * @param  DeafMuteParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) DeafMuteParticipant (input *DeafMuteParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/deafMuteParticipant.{ResponseType}"

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

        "conferenceSid" : input.ConferenceSid,
        "ParticipantSid" : input.ParticipantSid,
        "Muted" : input.Muted,
        "Deaf" : input.Deaf,

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
 * View Participant
 * @param  ViewParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ViewParticipant (input *ViewParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewparticipant.{ResponseType}"

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

        "ConferenceSid" : input.ConferenceSid,
        "ParticipantSid" : input.ParticipantSid,

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
 * Add Participant in conference 
 * @param  AddParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) AddParticipant (input *AddParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/addParticipant.{ResponseType}"

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

        "conferencesid" : input.Conferencesid,
        "participantnumber" : input.Participantnumber,
        "muted" : input.Muted,
        "deaf" : input.Deaf,

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
 * View Conference
 * @param  ViewConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ViewConference (input *ViewConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewconference.{ResponseType}"

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

        "conferencesid" : input.Conferencesid,

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
 * Here you can experiment with initiating a conference call through message360 and view the request response generated when doing so.
 * @param  CreateConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateConference (input *CreateConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/createConference.{ResponseType}"

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

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "Url" : input.Url,
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

        "From" : input.From,
        "To" : input.To,
        "Method" : models_pkg.HttpActionEnumToValue(input.Method),
        "RecordCallbackUrl" : input.RecordCallbackUrl,
        "StatusCallBackUrl" : input.StatusCallBackUrl,
        "StatusCallBackMethod" : models_pkg.HttpActionEnumToValue(input.StatusCallBackMethod),
        "FallBackUrl" : input.FallBackUrl,
        "FallBackMethod" : models_pkg.HttpActionEnumToValue(input.FallBackMethod),
        "Record" : input.Record,
        "RecordCallbackMethod" : models_pkg.HttpActionEnumToValue(input.RecordCallbackMethod),
        "SchdeuleTime" : input.SchdeuleTime,
        "Timeout" : input.Timeout,

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
 * Remove a participant from a conference.
 * @param  HangupParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) HangupParticipant (input *HangupParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/hangupparticipant.{ResponseType}"

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

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "ParticipantSid" : input.ParticipantSid,
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

        "ConferenceSid" : input.ConferenceSid,

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
 * Play an audio file during a conference.
 * @param  PlayConferenceAudioInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) PlayConferenceAudio (input *PlayConferenceAudioInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/playaudio.{ResponseType}"

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

        "ConferenceSid" : input.ConferenceSid,
        "ParticipantSid" : input.ParticipantSid,
        "AudioUrl" : models_pkg.AudioFormatEnumToValue(input.AudioUrl),

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
 * List Participant
 * @param  ListParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ListParticipant (input *ListParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listparticipant.{ResponseType}"

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

        "ConferenceSid" : input.ConferenceSid,
        "Page" : apihelper_pkg.ToString(*input.Page, "1"),
        "Pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "Muted" : input.Muted,
        "Deaf" : input.Deaf,

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
 * List Conference
 * @param  ListConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ListConference (input *ListConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listconference.{ResponseType}"

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

        "Page" : apihelper_pkg.ToString(*input.Page, "1"),
        "PageSize" : apihelper_pkg.ToString(*input.PageSize, "10"),
        "FriendlyName" : input.FriendlyName,
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

