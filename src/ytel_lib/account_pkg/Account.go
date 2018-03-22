/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package account_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the ACCOUNT_IMPL
 */
type ACCOUNT interface {
    ViewAccount (*ViewAccountInput) (string, error)
}

/*
 * Factory for the ACCOUNT interaface returning ACCOUNT_IMPL
 */
func NewACCOUNT(config configuration_pkg.CONFIGURATION) *ACCOUNT_IMPL {
    client := new(ACCOUNT_IMPL)
    client.config = config
    return client
}