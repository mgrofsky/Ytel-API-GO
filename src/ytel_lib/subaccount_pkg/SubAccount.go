/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package subaccount_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the SUBACCOUNT_IMPL
 */
type SUBACCOUNT interface {
    DeleteSubAccount (*DeleteSubAccountInput) (string, error)

    SuspendSubAccount (*SuspendSubAccountInput) (string, error)

    CreateSubAccount (*CreateSubAccountInput) (string, error)
}

/*
 * Factory for the SUBACCOUNT interaface returning SUBACCOUNT_IMPL
 */
func NewSUBACCOUNT(config configuration_pkg.CONFIGURATION) *SUBACCOUNT_IMPL {
    client := new(SUBACCOUNT_IMPL)
    client.config = config
    return client
}