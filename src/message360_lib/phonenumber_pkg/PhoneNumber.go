/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package phonenumber_pkg


/*
 * Interface for the PHONENUMBER_IMPL
 */
type PHONENUMBER interface {
    AvailablePhoneNumber (*AvailablePhoneNumberInput) (string, error)

    ListNumber (*ListNumberInput) (string, error)

    ViewNumberDetails (*ViewNumberDetailsInput) (string, error)

    ReleaseNumber (*ReleaseNumberInput) (string, error)

    BuyNumber (*BuyNumberInput) (string, error)

    UpdatePhoneNumber (*UpdatePhoneNumberInput) (string, error)
}

/*
 * Factory for the PHONENUMBER interaface returning PHONENUMBER_IMPL
 */
func NewPHONENUMBER() PHONENUMBER {
    return &PHONENUMBER_IMPL{}
}