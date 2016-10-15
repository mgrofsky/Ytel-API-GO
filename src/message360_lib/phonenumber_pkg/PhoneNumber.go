/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/15/2016
 */

package phonenumber_pkg

import "message360_lib/models_pkg"

/*
 * Interface for the PHONENUMBER_IMPL
 */
type PHONENUMBER interface {
    CreateAvailablePhoneNumber (string, string, *int64, *string) (string, error)

    CreateListNumber (*int64, *int64, *string, *string, *string) (string, error)

    CreateReleaseNumber (string, *string) (string, error)

    CreateBuyNumber (string, *string) (string, error)

    CreateViewNumberDetails (string, *string) (string, error)

    UpdatePhoneNumber (string, *string, *string, models_pkg.HttpMethodEnum, *string, models_pkg.HttpMethodEnum, *string, models_pkg.HttpMethodEnum, *string, models_pkg.HttpMethodEnum, *string, models_pkg.HttpMethodEnum, *string, models_pkg.HttpMethodEnum, *string) (string, error)
}

/*
 * Factory for the PHONENUMBER interaface returning PHONENUMBER_IMPL
 */
func NewPHONENUMBER() PHONENUMBER {
    return &PHONENUMBER_IMPL{}
}