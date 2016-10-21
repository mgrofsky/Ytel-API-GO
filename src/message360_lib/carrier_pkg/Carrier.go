/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/21/2016
 */

package carrier_pkg


/*
 * Interface for the CARRIER_IMPL
 */
type CARRIER interface {
    CreateCarrierLookup (string, *string) (string, error)

    CreateCarrierLookupList (*string, *string, *string) (string, error)
}

/*
 * Factory for the CARRIER interaface returning CARRIER_IMPL
 */
func NewCARRIER() CARRIER {
    return &CARRIER_IMPL{}
}