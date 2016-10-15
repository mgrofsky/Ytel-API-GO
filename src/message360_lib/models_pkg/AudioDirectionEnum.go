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
 * Type definition for AudioDirectionEnum enum
 */
type AudioDirectionEnum int

/**
 * Value collection for AudioDirectionEnum enum
 */
const (
    AudioDirection_IN AudioDirectionEnum = 1 + iota
    AudioDirection_OUT
)

func (r AudioDirectionEnum) MarshalJSON() ([]byte, error) { 
    s := AudioDirectionEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *AudioDirectionEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  AudioDirectionEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts AudioDirectionEnum to its string representation
 */
func AudioDirectionEnumToValue(audioDirectionEnum AudioDirectionEnum) string {
    switch audioDirectionEnum {
        case AudioDirection_IN:
    		return "IN"		
        case AudioDirection_OUT:
    		return "OUT"		
        default:
        	return "IN"
    }
}

/**
 * Converts AudioDirectionEnum Array to its string Array representation
*/
func AudioDirectionEnumArrayToValue(audioDirectionEnum []AudioDirectionEnum) []string {
    convArray := make([]string,len( audioDirectionEnum))
    for i:=0; i<len(audioDirectionEnum);i++ {
        convArray[i] = AudioDirectionEnumToValue(audioDirectionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AudioDirectionEnumFromValue(value string) AudioDirectionEnum {
    switch value {
        case "IN":
            return AudioDirection_IN
        case "OUT":
            return AudioDirection_OUT
        default:
            return AudioDirection_IN
    }
}