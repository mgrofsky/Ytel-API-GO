/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/01/2016
 */
package conference_pkg


import(
	"message360_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"message360_lib"
	"message360_lib/apihelper_pkg"
)

/*
 * Input structure for the method CreateDeafMuteParticipant
 */
type CreateDeafMuteParticipantInput struct {
    ConferenceSid   string          //TODO: Write general description for this field
    ParticipantSid  string          //TODO: Write general description for this field
    Muted           *bool           //TODO: Write general description for this field
    Deaf            *bool           //TODO: Write general description for this field
    ResponseType    *string         //Response Type either json or xml
}

/*
 * Input structure for the method CreateListConference
 */
type CreateListConferenceInput struct {
    Page            *int64          //Which page of the overall response will be returned. Zero indexed
    PageSize        *int64          //Number of individual resources listed in the response per page
    FriendlyName    *string         //Only return conferences with the specified FriendlyName
    Status          models_pkg.InterruptedCallStatus //TODO: Write general description for this field
    DateCreated     *string         //TODO: Write general description for this field
    DateUpdated     *string         //TODO: Write general description for this field
    ResponseType    *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateViewConference
 */
type CreateViewConferenceInput struct {
    Conferencesid   string          //The unique identifier of each conference resource
    ResponseType    *string         //Response type format xml or json
}

/*
 * Input structure for the method AddParticipant
 */
type AddParticipantInput struct {
    Conferencesid     string          //Unique Conference Sid
    Participantnumber string          //Particiant Number
    Tocountrycode     int64           //TODO: Write general description for this field
    Muted             *bool           //TODO: Write general description for this field
    Deaf              *bool           //TODO: Write general description for this field
    ResponseType      *string         //Response type format xml or json
}

/*
 * Input structure for the method CreateListParticipant
 */
type CreateListParticipantInput struct {
    ConferenceSid   string          //unique conference sid
    Page            *int64          //page number
    Pagesize        *int64          //TODO: Write general description for this field
    Muted           *bool           //TODO: Write general description for this field
    Deaf            *bool           //TODO: Write general description for this field
    ResponseType    *string         //Response format, xml or json
}

/*
 * Input structure for the method CreateViewParticipant
 */
type CreateViewParticipantInput struct {
    ConferenceSid   string          //unique conference sid
    ParticipantSid  string          //TODO: Write general description for this field
    ResponseType    *string         //Response type format xml or json
}

/*
 * Client structure as interface implementation
 */
type CONFERENCE_IMPL struct { }

/**
 * Deaf Mute Participant
 * @param  CreateDeafMuteParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateDeafMuteParticipant (input *CreateDeafMuteParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/deafMuteParticipant.{ResponseType}"

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
 * List Conference
 * @param  CreateListConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateListConference (input *CreateListConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listconference.{ResponseType}"

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
        "PageSize" : input.PageSize,
        "FriendlyName" : input.FriendlyName,
        "Status" : models_pkg.InterruptedCallStatusToValue(input.Status),
        "DateCreated" : input.DateCreated,
        "DateUpdated" : input.DateUpdated,

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
 * @param  CreateViewConferenceInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateViewConference (input *CreateViewConferenceInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewconference.{ResponseType}"

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

        "conferencesid" : input.Conferencesid,
        "participantnumber" : input.Participantnumber,
        "tocountrycode" : input.Tocountrycode,
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
 * List Participant
 * @param  CreateListParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateListParticipant (input *CreateListParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/listparticipant.{ResponseType}"

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

        "ConferenceSid" : input.ConferenceSid,
        "Page" : input.Page,
        "Pagesize" : input.Pagesize,
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
 * @param  CreateViewParticipantInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *CONFERENCE_IMPL) CreateViewParticipant (input *CreateViewParticipantInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := message360_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/conferences/viewparticipant.{ResponseType}"

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

