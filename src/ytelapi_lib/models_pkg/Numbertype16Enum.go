/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for Numbertype16Enum enum
 */
type Numbertype16Enum int

/**
 * Value collection for Numbertype16Enum enum
 */
const (
    Numbertype16_ALL Numbertype16Enum = 1 + iota
    Numbertype16_VOICE
    Numbertype16_SMS
)

func (r Numbertype16Enum) MarshalJSON() ([]byte, error) { 
    s := Numbertype16EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Numbertype16Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Numbertype16EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Numbertype16Enum to its string representation
 */
func Numbertype16EnumToValue(numbertype16Enum Numbertype16Enum) string {
    switch numbertype16Enum {
        case Numbertype16_ALL:
    		return "ALL"		
        case Numbertype16_VOICE:
    		return "Voice"		
        case Numbertype16_SMS:
    		return "SMS"		
        default:
        	return "ALL"
    }
}

/**
 * Converts Numbertype16Enum Array to its string Array representation
*/
func Numbertype16EnumArrayToValue(numbertype16Enum []Numbertype16Enum) []string {
    convArray := make([]string,len( numbertype16Enum))
    for i:=0; i<len(numbertype16Enum);i++ {
        convArray[i] = Numbertype16EnumToValue(numbertype16Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Numbertype16EnumFromValue(value string) Numbertype16Enum {
    switch value {
        case "ALL":
            return Numbertype16_ALL
        case "Voice":
            return Numbertype16_VOICE
        case "SMS":
            return Numbertype16_SMS
        default:
            return Numbertype16_ALL
    }
}
