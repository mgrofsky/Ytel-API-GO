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
 * Type definition for InterruptedCallStatusEnum enum
 */
type InterruptedCallStatusEnum int

/**
 * Value collection for InterruptedCallStatusEnum enum
 */
const (
    InterruptedCallStatus_CANCELED InterruptedCallStatusEnum = 1 + iota
    InterruptedCallStatus_COMPLETED
)

func (r InterruptedCallStatusEnum) MarshalJSON() ([]byte, error) { 
    s := InterruptedCallStatusEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *InterruptedCallStatusEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  InterruptedCallStatusEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts InterruptedCallStatusEnum to its string representation
 */
func InterruptedCallStatusEnumToValue(interruptedCallStatusEnum InterruptedCallStatusEnum) string {
    switch interruptedCallStatusEnum {
        case InterruptedCallStatus_CANCELED:
    		return "CANCELED"		
        case InterruptedCallStatus_COMPLETED:
    		return "COMPLETED"		
        default:
        	return "CANCELED"
    }
}

/**
 * Converts InterruptedCallStatusEnum Array to its string Array representation
*/
func InterruptedCallStatusEnumArrayToValue(interruptedCallStatusEnum []InterruptedCallStatusEnum) []string {
    convArray := make([]string,len( interruptedCallStatusEnum))
    for i:=0; i<len(interruptedCallStatusEnum);i++ {
        convArray[i] = InterruptedCallStatusEnumToValue(interruptedCallStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func InterruptedCallStatusEnumFromValue(value string) InterruptedCallStatusEnum {
    switch value {
        case "CANCELED":
            return InterruptedCallStatus_CANCELED
        case "COMPLETED":
            return InterruptedCallStatus_COMPLETED
        default:
            return InterruptedCallStatus_CANCELED
    }
}
