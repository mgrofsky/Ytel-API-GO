/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package recording_pkg


/*
 * Interface for the RECORDING_IMPL
 */
type RECORDING interface {
    ViewRecording (*ViewRecordingInput) (string, error)

    DeleteRecording (*DeleteRecordingInput) (string, error)

    ListRecording (*ListRecordingInput) (string, error)
}

/*
 * Factory for the RECORDING interaface returning RECORDING_IMPL
 */
func NewRECORDING() RECORDING {
    return &RECORDING_IMPL{}
}