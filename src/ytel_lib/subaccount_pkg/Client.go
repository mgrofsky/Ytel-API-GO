/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package subaccount_pkg


import(
	"ytel_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"ytel_lib/apihelper_pkg"
	"ytel_lib/configuration_pkg"
)

/*
 * Input structure for the method DeleteSubAccount
 */
type DeleteSubAccountInput struct {
    SubAccountSID   string          //The SubaccountSid to be deleted
    MergeNumber     models_pkg.MergeNumberStatusEnum //0 to delete or 1 to merge numbers to parent account.
    ResponseType    string          //Response type format xml or json
}

/*
 * Input structure for the method SuspendSubAccount
 */
type SuspendSubAccountInput struct {
    SubAccountSID   string          //The SubaccountSid to be activated or suspended
    Activate        models_pkg.ActivateStatusEnum //0 to suspend or 1 to activate
    ResponseType    string          //TODO: Write general description for this field
}

/*
 * Input structure for the method CreateSubAccount
 */
type CreateSubAccountInput struct {
    FirstName       string          //Sub account user first name
    LastName        string          //sub account user last name
    Email           string          //Sub account email address
    FriendlyName    string          //Descriptive name of the sub account
    Password        string          //The password of the sub account.  Please make sure to make as password that is at least 8 characters long, contain a symbol, uppercase and a number.
    ResponseType    string          //Response type format xml or json
}

/*
 * Client structure as interface implementation
 */
type SUBACCOUNT_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Delete sub account or merge numbers into parent
 * @param  DeleteSubAccountInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SUBACCOUNT_IMPL) DeleteSubAccount (input *DeleteSubAccountInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/user/deletesubaccount.{ResponseType}"

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

        "SubAccountSID" : input.SubAccountSID,
        "MergeNumber" : models_pkg.MergeNumberStatusEnumToValue(input.MergeNumber),

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
 * Suspend or unsuspend
 * @param  SuspendSubAccountInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SUBACCOUNT_IMPL) SuspendSubAccount (input *SuspendSubAccountInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/user/subaccountactivation.{ResponseType}"

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

        "SubAccountSID" : input.SubAccountSID,
        "Activate" : models_pkg.ActivateStatusEnumToValue(input.Activate),

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
 * Create a sub user account under the parent account
 * @param  CreateSubAccountInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *SUBACCOUNT_IMPL) CreateSubAccount (input *CreateSubAccountInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/user/createsubaccount.{ResponseType}"

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

        "FirstName" : input.FirstName,
        "LastName" : input.LastName,
        "Email" : input.Email,
        "FriendlyName" : input.FriendlyName,
        "Password" : input.Password,

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
