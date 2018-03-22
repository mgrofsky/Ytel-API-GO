/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package phonenumber_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the PHONENUMBER_IMPL
 */
type PHONENUMBER interface {
    AvailablePhoneNumber (*AvailablePhoneNumberInput) (string, error)

    MassReleaseNumber (*MassReleaseNumberInput) (string, error)

    ViewNumberDetails (*ViewNumberDetailsInput) (string, error)

    ReleaseNumber (*ReleaseNumberInput) (string, error)

    BuyNumber (*BuyNumberInput) (string, error)

    UpdatePhoneNumber (*UpdatePhoneNumberInput) (string, error)

    TransferNumber (*TransferNumberInput) (string, error)

    ListNumber (*ListNumberInput) (string, error)

    MassUpdateNumber (*MassUpdateNumberInput) (string, error)

    GetDIDScoreNumber (*GetDIDScoreNumberInput) (string, error)

    BulkBuyNumber (*BulkBuyNumberInput) (string, error)
}

/*
 * Factory for the PHONENUMBER interaface returning PHONENUMBER_IMPL
 */
func NewPHONENUMBER(config configuration_pkg.CONFIGURATION) *PHONENUMBER_IMPL {
    client := new(PHONENUMBER_IMPL)
    client.config = config
    return client
}