/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/01/2016
 */

package usage_pkg


/*
 * Interface for the USAGE_IMPL
 */
type USAGE interface {
    CreateListUsage (*CreateListUsageInput) (string, error)
}

/*
 * Factory for the USAGE interaface returning USAGE_IMPL
 */
func NewUSAGE() USAGE {
    return &USAGE_IMPL{}
}