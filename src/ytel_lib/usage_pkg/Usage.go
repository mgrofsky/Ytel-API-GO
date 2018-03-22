/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package usage_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the USAGE_IMPL
 */
type USAGE interface {
    ListUsage (*ListUsageInput) (string, error)
}

/*
 * Factory for the USAGE interaface returning USAGE_IMPL
 */
func NewUSAGE(config configuration_pkg.CONFIGURATION) *USAGE_IMPL {
    client := new(USAGE_IMPL)
    client.config = config
    return client
}