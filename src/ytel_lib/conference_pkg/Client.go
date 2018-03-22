/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package conference_pkg


import(
	"ytel_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"ytel_lib/apihelper_pkg"
	"ytel_lib/configuration_pkg"
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
    ConferenceSid   string          //The unique identifier for a conference object.
    ParticipantSid  string          //The unique identifier for a participant object.
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method ViewConference
 */
type ViewConferenceInput struct {
    ConferenceSid   string          //The unique identifier of each conference resource
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method AddParticipant
 */
type AddParticipantInput struct {
    ConferenceSid     string          //The unique identifier for a conference object.
    ParticipantNumber string          //The phone number of the participant to be added.
    ResponseType      string          //Response type format xml or json
    Muted             *bool           //Specifies if participant should be muted.
    Deaf              *bool           //Specifies if the participant should hear audio in the conference.
}

/*
 * Input structure for the method CreateConference
 */
type CreateConferenceInput struct {
    From                 string          //A valid 10-digit number (E.164 format) that will be initiating the conference call.
    To                   string          //A valid 10-digit number (E.164 format) that is to receive the conference call.
    Url                  string          //URL requested once the conference connects
    ResponseType         string          //Response type format xml or json
    Method               models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once call connects.
    StatusCallBackUrl    *string         //URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the conference is finished.
    StatusCallBackMethod models_pkg.HttpActionEnum //Specifies the HTTP methodlinkclass used to request StatusCallbackUrl.
    FallbackUrl          *string         //URL requested if the initial Url parameter fails or encounters an error
    FallbackMethod       models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required FallbackUrl once call connects.
    Record               *bool           //Specifies if the conference should be recorded.
    RecordCallBackUrl    *string         //Recording parameters will be sent here upon completion.
    RecordCallBackMethod models_pkg.HttpActionEnum //Specifies the HTTP method used to request the required URL once conference connects.
    ScheduleTime         *string         //Schedule conference in future. Schedule time must be greater than current time
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
    ConferenceSid   string          //The unique identifier for a conference.
    ResponseType    string          //Response format, xml or json
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //The count of objects to return per page.
    Muted           *bool           //Specifies if participant should be muted.
    Deaf            *bool           //Specifies if the participant should hear audio in the conference.
}

/*
 * Input structure for the method ListConference
 */
type ListConferenceInput struct {
    ResponseType    string          //Response type format xml or json
    Page            *int64          //The page count to retrieve from the total results in the collection. Page indexing starts at 1.
    Pagesize        *int64          //Number of individual resources listed in the response per page
    FriendlyName    *string         //Only return conferences with the specified FriendlyName
    DateCreated     *string         //Conference created date
}

/*
 * Client structure as interface implementation
 */
type CONFERENCE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Deaf Mute Participant
 * @param  DeafMuteParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) DeafMuteParticipant (input *DeafMuteParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "conferenceSid" : input.ConferenceSid,
        "ParticipantSid" : input.ParticipantSid,
        "Muted" : input.Muted,
        "Deaf" : input.Deaf,

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
 * Retrieve information about a participant by its ParticipantSid.
 * @param  ViewParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ViewParticipant (input *ViewParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewParticipant.{ResponseType}"

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

        "ConferenceSid" : input.ConferenceSid,
        "ParticipantSid" : input.ParticipantSid,

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
 * Retrieve information about a conference by its ConferenceSid.
 * @param  ViewConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ViewConference (input *ViewConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "ConferenceSid" : input.ConferenceSid,

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
 * Add Participant in conference 
 * @param  AddParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) AddParticipant (input *AddParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "ConferenceSid" : input.ConferenceSid,
        "ParticipantNumber" : input.ParticipantNumber,
        "Muted" : input.Muted,
        "Deaf" : input.Deaf,

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
 * Here you can experiment with initiating a conference call through Ytel and view the request response generated when doing so.
 * @param  CreateConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateConference (input *CreateConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "From" : input.From,
        "To" : input.To,
        "Method" : models_pkg.HttpActionEnumToValue(input.Method),
        "StatusCallBackUrl" : input.StatusCallBackUrl,
        "StatusCallBackMethod" : models_pkg.HttpActionEnumToValue(input.StatusCallBackMethod),
        "FallbackUrl" : input.FallbackUrl,
        "FallbackMethod" : models_pkg.HttpActionEnumToValue(input.FallbackMethod),
        "Record" : input.Record,
        "RecordCallBackUrl" : input.RecordCallBackUrl,
        "RecordCallBackMethod" : models_pkg.HttpActionEnumToValue(input.RecordCallBackMethod),
        "ScheduleTime" : input.ScheduleTime,
        "Timeout" : input.Timeout,

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
 * Remove a participant from a conference.
 * @param  HangupParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) HangupParticipant (input *HangupParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/hangupParticipant.{ResponseType}"

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "ConferenceSid" : input.ConferenceSid,

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
 * Play an audio file during a conference.
 * @param  PlayConferenceAudioInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) PlayConferenceAudio (input *PlayConferenceAudioInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/playAudio.{ResponseType}"

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

        "ConferenceSid" : input.ConferenceSid,
        "ParticipantSid" : input.ParticipantSid,
        "AudioUrl" : models_pkg.AudioFormatEnumToValue(input.AudioUrl),

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
 * Retrieve a list of participants for an in-progress conference.
 * @param  ListParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ListParticipant (input *ListParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listParticipant.{ResponseType}"

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

        "ConferenceSid" : input.ConferenceSid,
        "Page" : apihelper_pkg.ToString(*input.Page, "1"),
        "Pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "Muted" : input.Muted,
        "Deaf" : input.Deaf,

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
 * Retrieve a list of conference objects.
 * @param  ListConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) ListConference (input *ListConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

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
        "user-agent" : "ytel-api",
    }

    //form parameters
    parameters := map[string]interface{} {

        "page" : apihelper_pkg.ToString(*input.Page, "1"),
        "pagesize" : apihelper_pkg.ToString(*input.Pagesize, "10"),
        "FriendlyName" : input.FriendlyName,
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

