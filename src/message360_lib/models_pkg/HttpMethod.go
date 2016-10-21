/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/21/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for HttpMethod enum
 */
type HttpMethod int

/**
 * Value collection for HttpMethod enum
 */
const (
    HttpMethod_GET HttpMethod = 1 + iota
    HttpMethod_POST
)

func (r HttpMethod) MarshalJSON() ([]byte, error) { 
    s := HttpMethodToValue(r)
    return json.Marshal(s) 
} 

func (r *HttpMethod) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HttpMethodFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HttpMethod to its string representation
 */
func HttpMethodToValue(httpMethod HttpMethod) string {
    switch httpMethod {
        case HttpMethod_GET:
    		return "GET"		
        case HttpMethod_POST:
    		return "POST"		
        default:
        	return "GET"
    }
}

/**
 * Converts HttpMethod Array to its string Array representation
*/
func HttpMethodArrayToValue(httpMethod []HttpMethod) []string {
    convArray := make([]string,len( httpMethod))
    for i:=0; i<len(httpMethod);i++ {
        convArray[i] = HttpMethodToValue(httpMethod[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HttpMethodFromValue(value string) HttpMethod {
    switch value {
        case "GET":
            return HttpMethod_GET
        case "POST":
            return HttpMethod_POST
        default:
            return HttpMethod_GET
    }
}
