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
 * Type definition for MergeNumberStatusEnum enum
 */
type MergeNumberStatusEnum int

/**
 * Value collection for MergeNumberStatusEnum enum
 */
const (
    MergeNumberStatus_DELETE MergeNumberStatusEnum = 0
    MergeNumberStatus_MERGE = 1
)

func (r MergeNumberStatusEnum) MarshalJSON() ([]byte, error) { 
    s := MergeNumberStatusEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *MergeNumberStatusEnum) UnmarshalJSON(data []byte) error { 
    var s int 
    json.Unmarshal(data, &s)
    v :=  MergeNumberStatusEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts MergeNumberStatusEnum to its int representation
 */
func MergeNumberStatusEnumToValue(mergeNumberStatusEnum MergeNumberStatusEnum) int {
    switch mergeNumberStatusEnum {
        case MergeNumberStatus_DELETE:
    		return 0		
        case MergeNumberStatus_MERGE:
    		return 1		
        default:
        	return 0
    }
}

/**
 * Converts MergeNumberStatusEnum Array to its string Array representation
*/
func MergeNumberStatusEnumArrayToValue(mergeNumberStatusEnum []MergeNumberStatusEnum) []int {
    convArray := make([]int,len( mergeNumberStatusEnum))
    for i:=0; i<len(mergeNumberStatusEnum);i++ {
        convArray[i] = MergeNumberStatusEnumToValue(mergeNumberStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func MergeNumberStatusEnumFromValue(value int) MergeNumberStatusEnum {
    switch value {
        case 0:
            return MergeNumberStatus_DELETE
        case 1:
            return MergeNumberStatus_MERGE
        default:
            return MergeNumberStatus_DELETE
    }
}
