/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/18/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for DirectionEnum enum
 */
type DirectionEnum int

/**
 * Value collection for DirectionEnum enum
 */
const (
    Direction_IN DirectionEnum = 1 + iota
    Direction_OUT
    Direction_BOTH
)

func (r DirectionEnum) MarshalJSON() ([]byte, error) { 
    s := DirectionEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *DirectionEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  DirectionEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts DirectionEnum to its string representation
 */
func DirectionEnumToValue(directionEnum DirectionEnum) string {
    switch directionEnum {
        case Direction_IN:
    		return "IN"		
        case Direction_OUT:
    		return "OUT"		
        case Direction_BOTH:
    		return "BOTH"		
        default:
        	return "IN"
    }
}

/**
 * Converts DirectionEnum Array to its string Array representation
*/
func DirectionEnumArrayToValue(directionEnum []DirectionEnum) []string {
    convArray := make([]string,len( directionEnum))
    for i:=0; i<len(directionEnum);i++ {
        convArray[i] = DirectionEnumToValue(directionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DirectionEnumFromValue(value string) DirectionEnum {
    switch value {
        case "IN":
            return Direction_IN
        case "OUT":
            return Direction_OUT
        case "BOTH":
            return Direction_BOTH
        default:
            return Direction_IN
    }
}
