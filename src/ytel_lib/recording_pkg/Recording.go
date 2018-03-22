/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package recording_pkg

import "ytel_lib/configuration_pkg"

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
func NewRECORDING(config configuration_pkg.CONFIGURATION) *RECORDING_IMPL {
    client := new(RECORDING_IMPL)
    client.config = config
    return client
}