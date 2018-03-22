/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package sms_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the SMS_IMPL
 */
type SMS interface {
    SendSMS (*SendSMSInput) (string, error)

    ViewSMS (*ViewSMSInput) (string, error)

    ListSMS (*ListSMSInput) (string, error)

    ListInboundSMS (*ListInboundSMSInput) (string, error)

    ViewDetailSMS (*ViewDetailSMSInput) (string, error)
}

/*
 * Factory for the SMS interaface returning SMS_IMPL
 */
func NewSMS(config configuration_pkg.CONFIGURATION) *SMS_IMPL {
    client := new(SMS_IMPL)
    client.config = config
    return client
}