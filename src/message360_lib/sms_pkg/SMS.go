/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/08/2016
 */

package sms_pkg


/*
 * Interface for the SMS_IMPL
 */
type SMS interface {
    CreateViewSMS (*CreateViewSMSInput) (string, error)

    CreateListInboundSMS (*CreateListInboundSMSInput) (string, error)

    CreateListSMS (*CreateListSMSInput) (string, error)

    CreateSendSMS (*CreateSendSMSInput) (string, error)
}

/*
 * Factory for the SMS interaface returning SMS_IMPL
 */
func NewSMS() SMS {
    return &SMS_IMPL{}
}