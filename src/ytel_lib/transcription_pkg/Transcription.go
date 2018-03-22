/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package transcription_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the TRANSCRIPTION_IMPL
 */
type TRANSCRIPTION interface {
    ViewTranscription (*ViewTranscriptionInput) (string, error)

    RecordingTranscription (*RecordingTranscriptionInput) (string, error)

    AudioURLTranscription (*AudioURLTranscriptionInput) (string, error)

    ListTranscription (*ListTranscriptionInput) (string, error)
}

/*
 * Factory for the TRANSCRIPTION interaface returning TRANSCRIPTION_IMPL
 */
func NewTRANSCRIPTION(config configuration_pkg.CONFIGURATION) *TRANSCRIPTION_IMPL {
    client := new(TRANSCRIPTION_IMPL)
    client.config = config
    return client
}