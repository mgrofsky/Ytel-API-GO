/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/08/2016
 */

package subaccount_pkg


/*
 * Interface for the SUBACCOUNT_IMPL
 */
type SUBACCOUNT interface {
    CreateSubAccount (*CreateSubAccountInput) (string, error)

    CreateSuspendSubAccount (*CreateSuspendSubAccountInput) (string, error)

    CreateDeleteMergeSubAccount (*CreateDeleteMergeSubAccountInput) (string, error)
}

/*
 * Factory for the SUBACCOUNT interaface returning SUBACCOUNT_IMPL
 */
func NewSUBACCOUNT() SUBACCOUNT {
    return &SUBACCOUNT_IMPL{}
}