/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/08/2016
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for SendEmailAs enum
 */
type SendEmailAs int

/**
 * Value collection for SendEmailAs enum
 */
const (
    SendEmailAs_TEXT SendEmailAs = 1 + iota
    SendEmailAs_HTML
)

func (r SendEmailAs) MarshalJSON() ([]byte, error) { 
    s := SendEmailAsToValue(r)
    return json.Marshal(s) 
} 

func (r *SendEmailAs) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  SendEmailAsFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts SendEmailAs to its string representation
 */
func SendEmailAsToValue(sendEmailAs SendEmailAs) string {
    switch sendEmailAs {
        case SendEmailAs_TEXT:
    		return "TEXT"		
        case SendEmailAs_HTML:
    		return "HTML"		
        default:
        	return "TEXT"
    }
}

/**
 * Converts SendEmailAs Array to its string Array representation
*/
func SendEmailAsArrayToValue(sendEmailAs []SendEmailAs) []string {
    convArray := make([]string,len( sendEmailAs))
    for i:=0; i<len(sendEmailAs);i++ {
        convArray[i] = SendEmailAsToValue(sendEmailAs[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SendEmailAsFromValue(value string) SendEmailAs {
    switch value {
        case "TEXT":
            return SendEmailAs_TEXT
        case "HTML":
            return SendEmailAs_HTML
        default:
            return SendEmailAs_TEXT
    }
}
