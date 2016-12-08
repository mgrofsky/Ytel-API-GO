/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/08/2016
 */

package transcription_pkg


/*
 * Interface for the TRANSCRIPTION_IMPL
 */
type TRANSCRIPTION interface {
    CreateAudioURLTranscription (*CreateAudioURLTranscriptionInput) (string, error)

    CreateRecordingTranscription (*CreateRecordingTranscriptionInput) (string, error)

    CreateViewTranscription (*CreateViewTranscriptionInput) (string, error)

    CreateListTranscription (*CreateListTranscriptionInput) (string, error)
}

/*
 * Factory for the TRANSCRIPTION interaface returning TRANSCRIPTION_IMPL
 */
func NewTRANSCRIPTION() TRANSCRIPTION {
    return &TRANSCRIPTION_IMPL{}
}