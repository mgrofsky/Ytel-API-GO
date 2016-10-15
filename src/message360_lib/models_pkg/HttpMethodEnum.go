/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/15/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for HttpMethodEnum enum
 */
type HttpMethodEnum int

/**
 * Value collection for HttpMethodEnum enum
 */
const (
    HttpMethod_GET HttpMethodEnum = 1 + iota
    HttpMethod_POST
)

func (r HttpMethodEnum) MarshalJSON() ([]byte, error) { 
    s := HttpMethodEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HttpMethodEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HttpMethodEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HttpMethodEnum to its string representation
 */
func HttpMethodEnumToValue(httpMethodEnum HttpMethodEnum) string {
    switch httpMethodEnum {
        case HttpMethod_GET:
    		return "GET"		
        case HttpMethod_POST:
    		return "POST"		
        default:
        	return "GET"
    }
}

/**
 * Converts HttpMethodEnum Array to its string Array representation
*/
func HttpMethodEnumArrayToValue(httpMethodEnum []HttpMethodEnum) []string {
    convArray := make([]string,len( httpMethodEnum))
    for i:=0; i<len(httpMethodEnum);i++ {
        convArray[i] = HttpMethodEnumToValue(httpMethodEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HttpMethodEnumFromValue(value string) HttpMethodEnum {
    switch value {
        case "GET":
            return HttpMethod_GET
        case "POST":
            return HttpMethod_POST
        default:
            return HttpMethod_GET
    }
}
