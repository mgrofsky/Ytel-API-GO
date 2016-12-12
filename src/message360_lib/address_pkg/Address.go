/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/12/2016
 */

package address_pkg


/*
 * Interface for the ADDRESS_IMPL
 */
type ADDRESS interface {
    CreateAddress (*CreateAddressInput) (string, error)

    CreateDeleteAddress (*CreateDeleteAddressInput) (string, error)

    CreateVerifyAddress (*CreateVerifyAddressInput) (string, error)

    CreateListAddress (*CreateListAddressInput) (string, error)

    CreateViewAddress (*CreateViewAddressInput) (string, error)
}

/*
 * Factory for the ADDRESS interaface returning ADDRESS_IMPL
 */
func NewADDRESS() ADDRESS {
    return &ADDRESS_IMPL{}
}