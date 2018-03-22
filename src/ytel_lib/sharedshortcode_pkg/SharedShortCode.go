/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package sharedshortcode_pkg

import "github.com/satori/go.uuid"
import "ytel_lib/configuration_pkg"

/*
 * Interface for the SHAREDSHORTCODE_IMPL
 */
type SHAREDSHORTCODE interface {
    ViewTemplate (*ViewTemplateInput) (string, error)

    ViewSharedShortcodes (*ViewSharedShortcodesInput) (string, error)

    ListOutboundSharedShortcodes (*ListOutboundSharedShortcodesInput) (string, error)

    ListInboundSharedShortcodes (*ListInboundSharedShortcodesInput) (string, error)

    SendSharedShortcode (*SendSharedShortcodeInput) (string, error)

    ListTemplates (*ListTemplatesInput) (string, error)

    ViewKeyword (*ViewKeywordInput) (string, error)

    ListKeyword (*ListKeywordInput) (string, error)

    ViewAssignement (*ViewAssignementInput) (string, error)

    ListAssignment (*ListAssignmentInput) (string, error)

    UpdateAssignment (*UpdateAssignmentInput) (string, error)
}

/*
 * Factory for the SHAREDSHORTCODE interaface returning SHAREDSHORTCODE_IMPL
 */
func NewSHAREDSHORTCODE(config configuration_pkg.CONFIGURATION) *SHAREDSHORTCODE_IMPL {
    client := new(SHAREDSHORTCODE_IMPL)
    client.config = config
    return client
}