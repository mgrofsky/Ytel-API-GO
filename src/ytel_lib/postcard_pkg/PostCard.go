/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package postcard_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the POSTCARD_IMPL
 */
type POSTCARD interface {
    ViewPostcard (*ViewPostcardInput) (string, error)

    CreatePostcard (*CreatePostcardInput) (string, error)

    ListPostcards (*ListPostcardsInput) (string, error)

    DeletePostcard (*DeletePostcardInput) (string, error)
}

/*
 * Factory for the POSTCARD interaface returning POSTCARD_IMPL
 */
func NewPOSTCARD(config configuration_pkg.CONFIGURATION) *POSTCARD_IMPL {
    client := new(POSTCARD_IMPL)
    client.config = config
    return client
}