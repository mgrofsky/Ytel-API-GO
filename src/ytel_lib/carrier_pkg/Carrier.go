/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package carrier_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the CARRIER_IMPL
 */
type CARRIER interface {
    CarrierLookupList (*CarrierLookupListInput) (string, error)

    CarrierLookup (*CarrierLookupInput) (string, error)
}

/*
 * Factory for the CARRIER interaface returning CARRIER_IMPL
 */
func NewCARRIER(config configuration_pkg.CONFIGURATION) *CARRIER_IMPL {
    client := new(CARRIER_IMPL)
    client.config = config
    return client
}