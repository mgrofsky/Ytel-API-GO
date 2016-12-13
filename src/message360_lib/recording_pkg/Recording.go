/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/12/2016
 */

package recording_pkg


/*
 * Interface for the RECORDING_IMPL
 */
type RECORDING interface {
    CreateListRecording (*CreateListRecordingInput) (string, error)

    CreateDeleteRecording (*CreateDeleteRecordingInput) (string, error)

    CreateViewRecording (*CreateViewRecordingInput) (string, error)
}

/*
 * Factory for the RECORDING interaface returning RECORDING_IMPL
 */
func NewRECORDING() RECORDING {
    return &RECORDING_IMPL{}
}