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
 * Type definition for Status enum
 */
type Status int

/**
 * Value collection for Status enum
 */
const (
    Status_INPROGRESS Status = 1 + iota
    Status_SUCCESS
    Status_FAILURE
)

func (r Status) MarshalJSON() ([]byte, error) { 
    s := StatusToValue(r)
    return json.Marshal(s) 
} 

func (r *Status) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  StatusFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status to its string representation
 */
func StatusToValue(status Status) string {
    switch status {
        case Status_INPROGRESS:
    		return "INPROGRESS"		
        case Status_SUCCESS:
    		return "Success"		
        case Status_FAILURE:
    		return "Failure"		
        default:
        	return "INPROGRESS"
    }
}

/**
 * Converts Status Array to its string Array representation
*/
func StatusArrayToValue(status []Status) []string {
    convArray := make([]string,len( status))
    for i:=0; i<len(status);i++ {
        convArray[i] = StatusToValue(status[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusFromValue(value string) Status {
    switch value {
        case "INPROGRESS":
            return Status_INPROGRESS
        case "Success":
            return Status_SUCCESS
        case "Failure":
            return Status_FAILURE
        default:
            return Status_INPROGRESS
    }
}
