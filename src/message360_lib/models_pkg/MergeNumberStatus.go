/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/02/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for MergeNumberStatus enum
 */
type MergeNumberStatus int

/**
 * Value collection for MergeNumberStatus enum
 */
const (
    MergeNumberStatus_DELETE MergeNumberStatus = 0
    MergeNumberStatus_MERGE = 1
)

func (r MergeNumberStatus) MarshalJSON() ([]byte, error) { 
    s := MergeNumberStatusToValue(r)
    return json.Marshal(s) 
} 

func (r *MergeNumberStatus) UnmarshalJSON(data []byte) error { 
    var s int 
    json.Unmarshal(data, &s)
    v :=  MergeNumberStatusFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts MergeNumberStatus to its int representation
 */
func MergeNumberStatusToValue(mergeNumberStatus MergeNumberStatus) int {
    switch mergeNumberStatus {
        case MergeNumberStatus_DELETE:
    		return 0		
        case MergeNumberStatus_MERGE:
    		return 1		
        default:
        	return 0
    }
}

/**
 * Converts MergeNumberStatus Array to its string Array representation
*/
func MergeNumberStatusArrayToValue(mergeNumberStatus []MergeNumberStatus) []int {
    convArray := make([]int,len( mergeNumberStatus))
    for i:=0; i<len(mergeNumberStatus);i++ {
        convArray[i] = MergeNumberStatusToValue(mergeNumberStatus[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func MergeNumberStatusFromValue(value int) MergeNumberStatus {
    switch value {
        case 0:
            return MergeNumberStatus_DELETE
        case 1:
            return MergeNumberStatus_MERGE
        default:
            return MergeNumberStatus_DELETE
    }
}
