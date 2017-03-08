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
 * Type definition for InterruptedCallStatus enum
 */
type InterruptedCallStatus int

/**
 * Value collection for InterruptedCallStatus enum
 */
const (
    InterruptedCallStatus_CANCELED InterruptedCallStatus = 1 + iota
    InterruptedCallStatus_COMPLETED
)

func (r InterruptedCallStatus) MarshalJSON() ([]byte, error) { 
    s := InterruptedCallStatusToValue(r)
    return json.Marshal(s) 
} 

func (r *InterruptedCallStatus) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  InterruptedCallStatusFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts InterruptedCallStatus to its string representation
 */
func InterruptedCallStatusToValue(interruptedCallStatus InterruptedCallStatus) string {
    switch interruptedCallStatus {
        case InterruptedCallStatus_CANCELED:
    		return "CANCELED"		
        case InterruptedCallStatus_COMPLETED:
    		return "COMPLETED"		
        default:
        	return "CANCELED"
    }
}

/**
 * Converts InterruptedCallStatus Array to its string Array representation
*/
func InterruptedCallStatusArrayToValue(interruptedCallStatus []InterruptedCallStatus) []string {
    convArray := make([]string,len( interruptedCallStatus))
    for i:=0; i<len(interruptedCallStatus);i++ {
        convArray[i] = InterruptedCallStatusToValue(interruptedCallStatus[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func InterruptedCallStatusFromValue(value string) InterruptedCallStatus {
    switch value {
        case "CANCELED":
            return InterruptedCallStatus_CANCELED
        case "COMPLETED":
            return InterruptedCallStatus_COMPLETED
        default:
            return InterruptedCallStatus_CANCELED
    }
}
