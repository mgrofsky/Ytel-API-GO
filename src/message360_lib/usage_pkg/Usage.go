/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/21/2016
 */

package usage_pkg


/*
 * Interface for the USAGE_IMPL
 */
type USAGE interface {
    CreateListUsage (string, string, string, *string) (string, error)
}

/*
 * Factory for the USAGE interaface returning USAGE_IMPL
 */
func NewUSAGE() USAGE {
    return &USAGE_IMPL{}
}