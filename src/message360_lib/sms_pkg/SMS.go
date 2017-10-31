/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package sms_pkg


/*
 * Interface for the SMS_IMPL
 */
type SMS interface {
    SendSMS (*SendSMSInput) (string, error)

    ViewSMS (*ViewSMSInput) (string, error)

    ListSMS (*ListSMSInput) (string, error)

    ListInboundSMS (*ListInboundSMSInput) (string, error)
}

/*
 * Factory for the SMS interaface returning SMS_IMPL
 */
func NewSMS() SMS {
    return &SMS_IMPL{}
}