/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/15/2016
 */

package recording_pkg


/*
 * Interface for the RECORDING_IMPL
 */
type RECORDING interface {
    CreateDeleteRecording (string, *string) (string, error)

    CreateViewRecording (string, *string) (string, error)

    CreateListRecording (*int64, *int64, *string, *string, *string) (string, error)
}

/*
 * Factory for the RECORDING interaface returning RECORDING_IMPL
 */
func NewRECORDING() RECORDING {
    return &RECORDING_IMPL{}
}