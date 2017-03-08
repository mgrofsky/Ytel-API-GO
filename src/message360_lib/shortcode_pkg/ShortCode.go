/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package shortcode_pkg

import "github.com/satori/go.uuid"

/*
 * Interface for the SHORTCODE_IMPL
 */
type SHORTCODE interface {
    CreateViewTemplate (*CreateViewTemplateInput) (string, error)

    CreateSendShortCode (*CreateSendShortCodeInput, map[string]interface{}) (string, error)

    CreateListInboundShortCode (*CreateListInboundShortCodeInput) (string, error)

    CreateListShortCode (*CreateListShortCodeInput) (string, error)

    CreateListTemplates (*CreateListTemplatesInput) (string, error)

    CreateViewShortCode (*CreateViewShortCodeInput) (string, error)
}

/*
 * Factory for the SHORTCODE interaface returning SHORTCODE_IMPL
 */
func NewSHORTCODE() SHORTCODE {
    return &SHORTCODE_IMPL{}
}