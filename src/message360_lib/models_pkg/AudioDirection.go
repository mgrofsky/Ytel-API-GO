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
 * Type definition for AudioDirection enum
 */
type AudioDirection int

/**
 * Value collection for AudioDirection enum
 */
const (
    AudioDirection_IN AudioDirection = 1 + iota
    AudioDirection_OUT
)

func (r AudioDirection) MarshalJSON() ([]byte, error) { 
    s := AudioDirectionToValue(r)
    return json.Marshal(s) 
} 

func (r *AudioDirection) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  AudioDirectionFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts AudioDirection to its string representation
 */
func AudioDirectionToValue(audioDirection AudioDirection) string {
    switch audioDirection {
        case AudioDirection_IN:
    		return "IN"		
        case AudioDirection_OUT:
    		return "OUT"		
        default:
        	return "IN"
    }
}

/**
 * Converts AudioDirection Array to its string Array representation
*/
func AudioDirectionArrayToValue(audioDirection []AudioDirection) []string {
    convArray := make([]string,len( audioDirection))
    for i:=0; i<len(audioDirection);i++ {
        convArray[i] = AudioDirectionToValue(audioDirection[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AudioDirectionFromValue(value string) AudioDirection {
    switch value {
        case "IN":
            return AudioDirection_IN
        case "OUT":
            return AudioDirection_OUT
        default:
            return AudioDirection_IN
    }
}
