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
 * Type definition for ProductCodeEnum enum
 */
type ProductCodeEnum int

/**
 * Value collection for ProductCodeEnum enum
 */
const (
    ProductCode_ALL ProductCodeEnum = 0
    ProductCode_OUTBOUND_CALL = 1
    ProductCode_INBOUND_CALL = 2
    ProductCode_OUTBOUND_SMS = 3
    ProductCode_INBOUND_SMS = 4
    ProductCode_TRANSCRIPTION = 5
    ProductCode_EMAIL = 6
    ProductCode_CONFERENCE = 7
    ProductCode_CARRIER = 8
    ProductCode_NUMBER_PURCHASED = 9
    ProductCode_DIRECT_MAIL_AREAMAIL = 10
    ProductCode_DIRECT_MAIL_POSTCARD = 11
    ProductCode_DIRECT_MAIL_LETTERS = 12
    ProductCode_DIRECT_MAIL_VERIFIED_ADDRESS = 13
)

func (r ProductCodeEnum) MarshalJSON() ([]byte, error) { 
    s := ProductCodeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ProductCodeEnum) UnmarshalJSON(data []byte) error { 
    var s int 
    json.Unmarshal(data, &s)
    v :=  ProductCodeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ProductCodeEnum to its int representation
 */
func ProductCodeEnumToValue(productCodeEnum ProductCodeEnum) int {
    switch productCodeEnum {
        case ProductCode_ALL:
    		return 0		
        case ProductCode_OUTBOUND_CALL:
    		return 1		
        case ProductCode_INBOUND_CALL:
    		return 2		
        case ProductCode_OUTBOUND_SMS:
    		return 3		
        case ProductCode_INBOUND_SMS:
    		return 4		
        case ProductCode_TRANSCRIPTION:
    		return 5		
        case ProductCode_EMAIL:
    		return 6		
        case ProductCode_CONFERENCE:
    		return 7		
        case ProductCode_CARRIER:
    		return 8		
        case ProductCode_NUMBER_PURCHASED:
    		return 9		
        case ProductCode_DIRECT_MAIL_AREAMAIL:
    		return 10		
        case ProductCode_DIRECT_MAIL_POSTCARD:
    		return 11		
        case ProductCode_DIRECT_MAIL_LETTERS:
    		return 12		
        case ProductCode_DIRECT_MAIL_VERIFIED_ADDRESS:
    		return 13		
        default:
        	return 0
    }
}

/**
 * Converts ProductCodeEnum Array to its string Array representation
*/
func ProductCodeEnumArrayToValue(productCodeEnum []ProductCodeEnum) []int {
    convArray := make([]int,len( productCodeEnum))
    for i:=0; i<len(productCodeEnum);i++ {
        convArray[i] = ProductCodeEnumToValue(productCodeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProductCodeEnumFromValue(value int) ProductCodeEnum {
    switch value {
        case 0:
            return ProductCode_ALL
        case 1:
            return ProductCode_OUTBOUND_CALL
        case 2:
            return ProductCode_INBOUND_CALL
        case 3:
            return ProductCode_OUTBOUND_SMS
        case 4:
            return ProductCode_INBOUND_SMS
        case 5:
            return ProductCode_TRANSCRIPTION
        case 6:
            return ProductCode_EMAIL
        case 7:
            return ProductCode_CONFERENCE
        case 8:
            return ProductCode_CARRIER
        case 9:
            return ProductCode_NUMBER_PURCHASED
        case 10:
            return ProductCode_DIRECT_MAIL_AREAMAIL
        case 11:
            return ProductCode_DIRECT_MAIL_POSTCARD
        case 12:
            return ProductCode_DIRECT_MAIL_LETTERS
        case 13:
            return ProductCode_DIRECT_MAIL_VERIFIED_ADDRESS
        default:
            return ProductCode_ALL
    }
}
