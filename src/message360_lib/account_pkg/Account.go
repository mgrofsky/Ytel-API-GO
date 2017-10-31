/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package account_pkg


/*
 * Interface for the ACCOUNT_IMPL
 */
type ACCOUNT interface {
    ViewAccount (*ViewAccountInput) (string, error)
}

/*
 * Factory for the ACCOUNT interaface returning ACCOUNT_IMPL
 */
func NewACCOUNT() ACCOUNT {
    return &ACCOUNT_IMPL{}
}