/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/12/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for Direction enum
 */
type Direction int

/**
 * Value collection for Direction enum
 */
const (
    Direction_IN Direction = 1 + iota
    Direction_OUT
    Direction_BOTH
)

func (r Direction) MarshalJSON() ([]byte, error) { 
    s := DirectionToValue(r)
    return json.Marshal(s) 
} 

func (r *Direction) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  DirectionFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Direction to its string representation
 */
func DirectionToValue(direction Direction) string {
    switch direction {
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
 * Converts Direction Array to its string Array representation
*/
func DirectionArrayToValue(direction []Direction) []string {
    convArray := make([]string,len( direction))
    for i:=0; i<len(direction);i++ {
        convArray[i] = DirectionToValue(direction[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DirectionFromValue(value string) Direction {
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
