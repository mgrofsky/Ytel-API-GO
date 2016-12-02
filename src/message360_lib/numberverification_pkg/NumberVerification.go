/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/02/2016
 */

package numberverification_pkg


/*
 * Interface for the NUMBERVERIFICATION_IMPL
 */
type NUMBERVERIFICATION interface {
    CreateVerifyNumber (*CreateVerifyNumberInput) (string, error)
}

/*
 * Factory for the NUMBERVERIFICATION interaface returning NUMBERVERIFICATION_IMPL
 */
func NewNUMBERVERIFICATION() NUMBERVERIFICATION {
    return &NUMBERVERIFICATION_IMPL{}
}