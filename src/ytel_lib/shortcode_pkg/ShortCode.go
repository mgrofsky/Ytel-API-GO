/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package shortcode_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the SHORTCODE_IMPL
 */
type SHORTCODE interface {
    SendDedicatedShortcode (*SendDedicatedShortcodeInput) (string, error)

    ViewShortcode (*ViewShortcodeInput) (string, error)

    ListShortcode (*ListShortcodeInput) (string, error)

    ListInboundShortcode (*ListInboundShortcodeInput) (string, error)

    ViewDedicatedShortcodeAssignment (*ViewDedicatedShortcodeAssignmentInput) (string, error)

    UpdateDedicatedShortCodeAssignment (*UpdateDedicatedShortCodeAssignmentInput) (string, error)

    ListShortCodeAssignment (*ListShortCodeAssignmentInput) (string, error)
}

/*
 * Factory for the SHORTCODE interaface returning SHORTCODE_IMPL
 */
func NewSHORTCODE(config configuration_pkg.CONFIGURATION) *SHORTCODE_IMPL {
    client := new(SHORTCODE_IMPL)
    client.config = config
    return client
}