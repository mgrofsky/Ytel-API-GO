/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/21/2016
 */

package sms_pkg

import "message360_lib/models_pkg"

/*
 * Interface for the SMS_IMPL
 */
type SMS interface {
    CreateSendSMS (int64, string, int64, string, string, models_pkg.HttpMethod, *string, *string) (string, error)

    CreateViewSMS (string, *string) (string, error)

    CreateListSMS (*int64, *int64, *string, *string, *string, *string) (string, error)

    CreateListInboundSMS (*int64, *string, *string, *string, *string) (string, error)
}

/*
 * Factory for the SMS interaface returning SMS_IMPL
 */
func NewSMS() SMS {
    return &SMS_IMPL{}
}