/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for SendEmailAsEnum enum
 */
type SendEmailAsEnum int

/**
 * Value collection for SendEmailAsEnum enum
 */
const (
    SendEmailAs_TEXT SendEmailAsEnum = 1 + iota
    SendEmailAs_HTML
)

func (r SendEmailAsEnum) MarshalJSON() ([]byte, error) { 
    s := SendEmailAsEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *SendEmailAsEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  SendEmailAsEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts SendEmailAsEnum to its string representation
 */
func SendEmailAsEnumToValue(sendEmailAsEnum SendEmailAsEnum) string {
    switch sendEmailAsEnum {
        case SendEmailAs_TEXT:
    		return "TEXT"		
        case SendEmailAs_HTML:
    		return "HTML"		
        default:
        	return "TEXT"
    }
}

/**
 * Converts SendEmailAsEnum Array to its string Array representation
*/
func SendEmailAsEnumArrayToValue(sendEmailAsEnum []SendEmailAsEnum) []string {
    convArray := make([]string,len( sendEmailAsEnum))
    for i:=0; i<len(sendEmailAsEnum);i++ {
        convArray[i] = SendEmailAsEnumToValue(sendEmailAsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SendEmailAsEnumFromValue(value string) SendEmailAsEnum {
    switch value {
        case "TEXT":
            return SendEmailAs_TEXT
        case "HTML":
            return SendEmailAs_HTML
        default:
            return SendEmailAs_TEXT
    }
}
