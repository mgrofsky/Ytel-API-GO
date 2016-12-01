/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/01/2016
 */

package carrier_pkg


/*
 * Interface for the CARRIER_IMPL
 */
type CARRIER interface {
    CreateCarrierLookup (*CreateCarrierLookupInput) (string, error)

    CreateCarrierLookupList (*CreateCarrierLookupListInput) (string, error)
}

/*
 * Factory for the CARRIER interaface returning CARRIER_IMPL
 */
func NewCARRIER() CARRIER {
    return &CARRIER_IMPL{}
}