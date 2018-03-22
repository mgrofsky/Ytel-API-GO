/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package areamail_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the AREAMAIL_IMPL
 */
type AREAMAIL interface {
    CreateAreaMail (*CreateAreaMailInput) (string, error)

    ViewAreaMail (*ViewAreaMailInput) (string, error)

    ListAreaMail (*ListAreaMailInput) (string, error)

    DeleteAreaMail (*DeleteAreaMailInput) (string, error)
}

/*
 * Factory for the AREAMAIL interaface returning AREAMAIL_IMPL
 */
func NewAREAMAIL(config configuration_pkg.CONFIGURATION) *AREAMAIL_IMPL {
    client := new(AREAMAIL_IMPL)
    client.config = config
    return client
}