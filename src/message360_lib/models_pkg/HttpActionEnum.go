/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for HttpActionEnum enum
 */
type HttpActionEnum int

/**
 * Value collection for HttpActionEnum enum
 */
const (
    HttpAction_GET HttpActionEnum = 1 + iota
    HttpAction_POST
)

func (r HttpActionEnum) MarshalJSON() ([]byte, error) { 
    s := HttpActionEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HttpActionEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HttpActionEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HttpActionEnum to its string representation
 */
func HttpActionEnumToValue(httpActionEnum HttpActionEnum) string {
    switch httpActionEnum {
        case HttpAction_GET:
    		return "GET"		
        case HttpAction_POST:
    		return "POST"		
        default:
        	return "GET"
    }
}

/**
 * Converts HttpActionEnum Array to its string Array representation
*/
func HttpActionEnumArrayToValue(httpActionEnum []HttpActionEnum) []string {
    convArray := make([]string,len( httpActionEnum))
    for i:=0; i<len(httpActionEnum);i++ {
        convArray[i] = HttpActionEnumToValue(httpActionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HttpActionEnumFromValue(value string) HttpActionEnum {
    switch value {
        case "GET":
            return HttpAction_GET
        case "POST":
            return HttpAction_POST
        default:
            return HttpAction_GET
    }
}
