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
 * Type definition for ActivateStatusEnum enum
 */
type ActivateStatusEnum int

/**
 * Value collection for ActivateStatusEnum enum
 */
const (
    ActivateStatus_ACTIVATE ActivateStatusEnum = 1
    ActivateStatus_DEACTIVATE = 0
)

func (r ActivateStatusEnum) MarshalJSON() ([]byte, error) { 
    s := ActivateStatusEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ActivateStatusEnum) UnmarshalJSON(data []byte) error { 
    var s int 
    json.Unmarshal(data, &s)
    v :=  ActivateStatusEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ActivateStatusEnum to its int representation
 */
func ActivateStatusEnumToValue(activateStatusEnum ActivateStatusEnum) int {
    switch activateStatusEnum {
        case ActivateStatus_ACTIVATE:
    		return 1		
        case ActivateStatus_DEACTIVATE:
    		return 0		
        default:
        	return 1
    }
}

/**
 * Converts ActivateStatusEnum Array to its string Array representation
*/
func ActivateStatusEnumArrayToValue(activateStatusEnum []ActivateStatusEnum) []int {
    convArray := make([]int,len( activateStatusEnum))
    for i:=0; i<len(activateStatusEnum);i++ {
        convArray[i] = ActivateStatusEnumToValue(activateStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ActivateStatusEnumFromValue(value int) ActivateStatusEnum {
    switch value {
        case 1:
            return ActivateStatus_ACTIVATE
        case 0:
            return ActivateStatus_DEACTIVATE
        default:
            return ActivateStatus_ACTIVATE
    }
}
