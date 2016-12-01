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
 * Type definition for ActivateStatus enum
 */
type ActivateStatus int

/**
 * Value collection for ActivateStatus enum
 */
const (
    ActivateStatus_ACTIVATE ActivateStatus = 1
    ActivateStatus_DEACTIVATE = 0
)

func (r ActivateStatus) MarshalJSON() ([]byte, error) { 
    s := ActivateStatusToValue(r)
    return json.Marshal(s) 
} 

func (r *ActivateStatus) UnmarshalJSON(data []byte) error { 
    var s int 
    json.Unmarshal(data, &s)
    v :=  ActivateStatusFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ActivateStatus to its int representation
 */
func ActivateStatusToValue(activateStatus ActivateStatus) int {
    switch activateStatus {
        case ActivateStatus_ACTIVATE:
    		return 1		
        case ActivateStatus_DEACTIVATE:
    		return 0		
        default:
        	return 1
    }
}

/**
 * Converts ActivateStatus Array to its string Array representation
*/
func ActivateStatusArrayToValue(activateStatus []ActivateStatus) []int {
    convArray := make([]int,len( activateStatus))
    for i:=0; i<len(activateStatus);i++ {
        convArray[i] = ActivateStatusToValue(activateStatus[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ActivateStatusFromValue(value int) ActivateStatus {
    switch value {
        case 1:
            return ActivateStatus_ACTIVATE
        case 0:
            return ActivateStatus_DEACTIVATE
        default:
            return ActivateStatus_ACTIVATE
    }
}
