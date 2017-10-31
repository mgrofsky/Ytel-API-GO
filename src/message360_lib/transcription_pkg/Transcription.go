/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package transcription_pkg


/*
 * Interface for the TRANSCRIPTION_IMPL
 */
type TRANSCRIPTION interface {
    ListTranscription (*ListTranscriptionInput) (string, error)

    ViewTranscription (*ViewTranscriptionInput) (string, error)

    RecordingTranscription (*RecordingTranscriptionInput) (string, error)

    AudioURLTranscription (*AudioURLTranscriptionInput) (string, error)
}

/*
 * Factory for the TRANSCRIPTION interaface returning TRANSCRIPTION_IMPL
 */
func NewTRANSCRIPTION() TRANSCRIPTION {
    return &TRANSCRIPTION_IMPL{}
}