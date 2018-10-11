/*
 * ytelapiv3_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package sharedshortcode_pkg

import "ytelapiv3_lib/configuration_pkg"

/*
 * Interface for the SHAREDSHORTCODE_IMPL
 */
type SHAREDSHORTCODE interface {
    CreateViewShortcode (string) (string, error)

    CreateViewKeyword (string) (string, error)

    CreateViewTemplate (uuid.UUID) (string, error)

    CreateListInboundSMS (*string, *int64, *int64, *string, *string) (string, error)

    CreateSendSMS (string, string, uuid.UUID, string, *string, *string) (string, error)

    CreateListTemplates (*string, *int64, *int64, *string) (string, error)

    CreateListKeywords (*int64, *int64, *string, *int64) (string, error)

    CreateListShortcodes (*string, *int64, *int64) (string, error)

    UpdateShortcode (string, *string, *string, *string, *string, *string) (string, error)
}

/*
 * Factory for the SHAREDSHORTCODE interaface returning SHAREDSHORTCODE_IMPL
 */
func NewSHAREDSHORTCODE(config configuration_pkg.CONFIGURATION) *SHAREDSHORTCODE_IMPL {
    client := new(SHAREDSHORTCODE_IMPL)
    client.config = config
    return client
}