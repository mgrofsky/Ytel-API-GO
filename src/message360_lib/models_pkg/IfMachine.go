/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/08/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for IfMachine enum
 */
type IfMachine int

/**
 * Value collection for IfMachine enum
 */
const (
    IfMachine_CONTINUE IfMachine = 1 + iota
    IfMachine_HANGUP
)

func (r IfMachine) MarshalJSON() ([]byte, error) { 
    s := IfMachineToValue(r)
    return json.Marshal(s) 
} 

func (r *IfMachine) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  IfMachineFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts IfMachine to its string representation
 */
func IfMachineToValue(ifMachine IfMachine) string {
    switch ifMachine {
        case IfMachine_CONTINUE:
    		return "CONTINUE"		
        case IfMachine_HANGUP:
    		return "HANGUP"		
        default:
        	return "CONTINUE"
    }
}

/**
 * Converts IfMachine Array to its string Array representation
*/
func IfMachineArrayToValue(ifMachine []IfMachine) []string {
    convArray := make([]string,len( ifMachine))
    for i:=0; i<len(ifMachine);i++ {
        convArray[i] = IfMachineToValue(ifMachine[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func IfMachineFromValue(value string) IfMachine {
    switch value {
        case "CONTINUE":
            return IfMachine_CONTINUE
        case "HANGUP":
            return IfMachine_HANGUP
        default:
            return IfMachine_CONTINUE
    }
}
