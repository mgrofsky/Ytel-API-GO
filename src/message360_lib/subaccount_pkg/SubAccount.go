/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package subaccount_pkg


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
func NewSUBACCOUNT() SUBACCOUNT {
    return &SUBACCOUNT_IMPL{}
}