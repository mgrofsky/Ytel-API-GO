/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package usage_pkg


import(
	"ytel_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"ytel_lib/apihelper_pkg"
	"ytel_lib/configuration_pkg"
)

/*
 * Input structure for the method ListUsage
 */
type ListUsageInput struct {
    ResponseType       string          //Response type format xml or json
    ProductCode        models_pkg.ProductCodeEnum //Filter usage results by product type.
    StartDate          *string         //Filter usage objects by start date.
    EndDate            *string         //Filter usage objects by end date.
    IncludeSubAccounts *string         //Will include all subaccount usage data
}

/*
 * Client structure as interface implementation
 */
type USAGE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Retrieve usage metrics regarding your Ytel account. The results includes inbound/outbound voice calls and inbound/outbound SMS messages as well as carrier lookup requests.
 * @param  ListUsageInput     Structure with all inputs
 * @return	Returns the string response from the API call
 */
func (me *USAGE_IMPL) ListUsage (input *ListUsageInput) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.DEFAULT,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/usage/listusage.{ResponseType}"

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

        "ProductCode" : models_pkg.ProductCodeEnumToValue(input.ProductCode),
        "startDate" : apihelper_pkg.ToString(*input.StartDate, "2016-09-06"),
        "endDate" : apihelper_pkg.ToString(*input.EndDate, "2016-09-06"),
        "IncludeSubAccounts" : input.IncludeSubAccounts,

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

