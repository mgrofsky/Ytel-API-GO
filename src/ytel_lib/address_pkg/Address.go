/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package address_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the ADDRESS_IMPL
 */
type ADDRESS interface {
    CreateAddress (*CreateAddressInput) (string, error)

    ViewAddress (*ViewAddressInput) (string, error)

    ListAddress (*ListAddressInput) (string, error)

    VerifyAddress (*VerifyAddressInput) (string, error)

    DeleteAddress (*DeleteAddressInput) (string, error)
}

/*
 * Factory for the ADDRESS interaface returning ADDRESS_IMPL
 */
func NewADDRESS(config configuration_pkg.CONFIGURATION) *ADDRESS_IMPL {
    client := new(ADDRESS_IMPL)
    client.config = config
    return client
}