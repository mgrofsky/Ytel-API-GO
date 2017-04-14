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
 * Type definition for AudioFormatEnum enum
 */
type AudioFormatEnum int

/**
 * Value collection for AudioFormatEnum enum
 */
const (
    AudioFormat_MP3 AudioFormatEnum = 1 + iota
    AudioFormat_WAV
)

func (r AudioFormatEnum) MarshalJSON() ([]byte, error) { 
    s := AudioFormatEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *AudioFormatEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  AudioFormatEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts AudioFormatEnum to its string representation
 */
func AudioFormatEnumToValue(audioFormatEnum AudioFormatEnum) string {
    switch audioFormatEnum {
        case AudioFormat_MP3:
    		return "mp3"		
        case AudioFormat_WAV:
    		return "wav"		
        default:
        	return "mp3"
    }
}

/**
 * Converts AudioFormatEnum Array to its string Array representation
*/
func AudioFormatEnumArrayToValue(audioFormatEnum []AudioFormatEnum) []string {
    convArray := make([]string,len( audioFormatEnum))
    for i:=0; i<len(audioFormatEnum);i++ {
        convArray[i] = AudioFormatEnumToValue(audioFormatEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AudioFormatEnumFromValue(value string) AudioFormatEnum {
    switch value {
        case "mp3":
            return AudioFormat_MP3
        case "wav":
            return AudioFormat_WAV
        default:
            return AudioFormat_MP3
    }
}
