/*
 * ytelapiv3_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for Type36Enum enum
 */
type Type36Enum int

/**
 * Value collection for Type36Enum enum
 */
const (
    Type36_TEXT Type36Enum = 1 + iota
    Type36_HTML
)

func (r Type36Enum) MarshalJSON() ([]byte, error) { 
    s := Type36EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type36Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type36EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type36Enum to its string representation
 */
func Type36EnumToValue(type36Enum Type36Enum) string {
    switch type36Enum {
        case Type36_TEXT:
    		return "TEXT"		
        case Type36_HTML:
    		return "HTML"		
        default:
        	return "TEXT"
    }
}

/**
 * Converts Type36Enum Array to its string Array representation
*/
func Type36EnumArrayToValue(type36Enum []Type36Enum) []string {
    convArray := make([]string,len( type36Enum))
    for i:=0; i<len(type36Enum);i++ {
        convArray[i] = Type36EnumToValue(type36Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type36EnumFromValue(value string) Type36Enum {
    switch value {
        case "TEXT":
            return Type36_TEXT
        case "HTML":
            return Type36_HTML
        default:
            return Type36_TEXT
    }
}