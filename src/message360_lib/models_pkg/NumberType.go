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
 * Type definition for NumberType enum
 */
type NumberType int

/**
 * Value collection for NumberType enum
 */
const (
    NumberType_ALL NumberType = 1 + iota
    NumberType_VOICE
    NumberType_SMS
)

func (r NumberType) MarshalJSON() ([]byte, error) { 
    s := NumberTypeToValue(r)
    return json.Marshal(s) 
} 

func (r *NumberType) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  NumberTypeFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts NumberType to its string representation
 */
func NumberTypeToValue(numberType NumberType) string {
    switch numberType {
        case NumberType_ALL:
    		return "ALL"		
        case NumberType_VOICE:
    		return "Voice"		
        case NumberType_SMS:
    		return "SMS"		
        default:
        	return "ALL"
    }
}

/**
 * Converts NumberType Array to its string Array representation
*/
func NumberTypeArrayToValue(numberType []NumberType) []string {
    convArray := make([]string,len( numberType))
    for i:=0; i<len(numberType);i++ {
        convArray[i] = NumberTypeToValue(numberType[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NumberTypeFromValue(value string) NumberType {
    switch value {
        case "ALL":
            return NumberType_ALL
        case "Voice":
            return NumberType_VOICE
        case "SMS":
            return NumberType_SMS
        default:
            return NumberType_ALL
    }
}
