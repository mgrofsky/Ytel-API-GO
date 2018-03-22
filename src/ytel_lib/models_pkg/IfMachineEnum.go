/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for IfMachineEnum enum
 */
type IfMachineEnum int

/**
 * Value collection for IfMachineEnum enum
 */
const (
    IfMachine_CONTINUE IfMachineEnum = 1 + iota
    IfMachine_HANGUP
)

func (r IfMachineEnum) MarshalJSON() ([]byte, error) { 
    s := IfMachineEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *IfMachineEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  IfMachineEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts IfMachineEnum to its string representation
 */
func IfMachineEnumToValue(ifMachineEnum IfMachineEnum) string {
    switch ifMachineEnum {
        case IfMachine_CONTINUE:
    		return "CONTINUE"		
        case IfMachine_HANGUP:
    		return "HANGUP"		
        default:
        	return "CONTINUE"
    }
}

/**
 * Converts IfMachineEnum Array to its string Array representation
*/
func IfMachineEnumArrayToValue(ifMachineEnum []IfMachineEnum) []string {
    convArray := make([]string,len( ifMachineEnum))
    for i:=0; i<len(ifMachineEnum);i++ {
        convArray[i] = IfMachineEnumToValue(ifMachineEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func IfMachineEnumFromValue(value string) IfMachineEnum {
    switch value {
        case "CONTINUE":
            return IfMachine_CONTINUE
        case "HANGUP":
            return IfMachine_HANGUP
        default:
            return IfMachine_CONTINUE
    }
}
