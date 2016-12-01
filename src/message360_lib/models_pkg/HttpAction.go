/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/01/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for HttpAction enum
 */
type HttpAction int

/**
 * Value collection for HttpAction enum
 */
const (
    HttpAction_GET HttpAction = 1 + iota
    HttpAction_POST
)

func (r HttpAction) MarshalJSON() ([]byte, error) { 
    s := HttpActionToValue(r)
    return json.Marshal(s) 
} 

func (r *HttpAction) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HttpActionFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HttpAction to its string representation
 */
func HttpActionToValue(httpAction HttpAction) string {
    switch httpAction {
        case HttpAction_GET:
    		return "GET"		
        case HttpAction_POST:
    		return "POST"		
        default:
        	return "GET"
    }
}

/**
 * Converts HttpAction Array to its string Array representation
*/
func HttpActionArrayToValue(httpAction []HttpAction) []string {
    convArray := make([]string,len( httpAction))
    for i:=0; i<len(httpAction);i++ {
        convArray[i] = HttpActionToValue(httpAction[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HttpActionFromValue(value string) HttpAction {
    switch value {
        case "GET":
            return HttpAction_GET
        case "POST":
            return HttpAction_POST
        default:
            return HttpAction_GET
    }
}
