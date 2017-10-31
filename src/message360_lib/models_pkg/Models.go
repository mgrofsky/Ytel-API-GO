/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package models_pkg


/*
 * Structure for the custom type MessageModel
 */
type MessageModel struct {
    ApiVersion      string          `json:"ApiVersion" form:"ApiVersion"` //TODO: Write general description for this field
    MessageSid      string          `json:"MessageSid" form:"MessageSid"` //TODO: Write general description for this field
    From            string          `json:"From" form:"From"` //TODO: Write general description for this field
    To              string          `json:"To" form:"To"` //TODO: Write general description for this field
    MessagePrice    string          `json:"MessagePrice" form:"MessagePrice"` //TODO: Write general description for this field
    Body            string          `json:"Body" form:"Body"` //TODO: Write general description for this field
    DateSent        string          `json:"DateSent" form:"DateSent"` //TODO: Write general description for this field
    Status          string          `json:"Status" form:"Status"` //TODO: Write general description for this field
    TemplateId      string          `json:"TemplateId" form:"TemplateId"` //TODO: Write general description for this field
    TemplateData    TemplateDataModel `json:"TemplateData" form:"TemplateData"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TemplateDataModel
 */
type TemplateDataModel struct {
    Companyname     string          `json:"companyname" form:"companyname"` //TODO: Write general description for this field
    Otpcode         string          `json:"otpcode" form:"otpcode"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ErrorsModel
 */
type ErrorsModel struct {
    Error           []*ErrorModel   `json:"Error" form:"Error"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ErrorModel
 */
type ErrorModel struct {
    Code            string          `json:"Code" form:"Code"` //TODO: Write general description for this field
    Message         string          `json:"Message" form:"Message"` //TODO: Write general description for this field
    MoreInfo        []string        `json:"MoreInfo" form:"MoreInfo"` //TODO: Write general description for this field
}
