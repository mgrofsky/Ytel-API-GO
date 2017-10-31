/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package shortcode_pkg

import "time"

/*
 * Interface for the SHORTCODE_IMPL
 */
type SHORTCODE interface {
    SendDedicatedShortcode (*SendDedicatedShortcodeInput) (string, error)

    ViewShortcode (*ViewShortcodeInput) (string, error)

    ListShortcode (*ListShortcodeInput) (string, error)

    ListInboundShortcode (*ListInboundShortcodeInput) (string, error)
}

/*
 * Factory for the SHORTCODE interaface returning SHORTCODE_IMPL
 */
func NewSHORTCODE() SHORTCODE {
    return &SHORTCODE_IMPL{}
}