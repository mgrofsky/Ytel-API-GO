/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/12/2016
 */

package phonenumber_pkg


/*
 * Interface for the PHONENUMBER_IMPL
 */
type PHONENUMBER interface {
    UpdatePhoneNumber (*UpdatePhoneNumberInput) (string, error)

    CreateBuyNumber (*CreateBuyNumberInput) (string, error)

    CreateReleaseNumber (*CreateReleaseNumberInput) (string, error)

    CreateViewNumberDetails (*CreateViewNumberDetailsInput) (string, error)

    CreateListNumber (*CreateListNumberInput) (string, error)

    CreateAvailablePhoneNumber (*CreateAvailablePhoneNumberInput) (string, error)
}

/*
 * Factory for the PHONENUMBER interaface returning PHONENUMBER_IMPL
 */
func NewPHONENUMBER() PHONENUMBER {
    return &PHONENUMBER_IMPL{}
}