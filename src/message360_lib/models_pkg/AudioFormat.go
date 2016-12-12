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
 * Type definition for AudioFormat enum
 */
type AudioFormat int

/**
 * Value collection for AudioFormat enum
 */
const (
    AudioFormat_MP3 AudioFormat = 1 + iota
    AudioFormat_WAV
)

func (r AudioFormat) MarshalJSON() ([]byte, error) { 
    s := AudioFormatToValue(r)
    return json.Marshal(s) 
} 

func (r *AudioFormat) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  AudioFormatFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts AudioFormat to its string representation
 */
func AudioFormatToValue(audioFormat AudioFormat) string {
    switch audioFormat {
        case AudioFormat_MP3:
    		return "mp3"		
        case AudioFormat_WAV:
    		return "wav"		
        default:
        	return "mp3"
    }
}

/**
 * Converts AudioFormat Array to its string Array representation
*/
func AudioFormatArrayToValue(audioFormat []AudioFormat) []string {
    convArray := make([]string,len( audioFormat))
    for i:=0; i<len(audioFormat);i++ {
        convArray[i] = AudioFormatToValue(audioFormat[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AudioFormatFromValue(value string) AudioFormat {
    switch value {
        case "mp3":
            return AudioFormat_MP3
        case "wav":
            return AudioFormat_WAV
        default:
            return AudioFormat_MP3
    }
}
