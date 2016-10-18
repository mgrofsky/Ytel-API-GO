/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/18/2016
 */

package transcription_pkg

import "message360_lib/models_pkg"

/*
 * Interface for the TRANSCRIPTION_IMPL
 */
type TRANSCRIPTION interface {
    CreateListTranscription (*int64, *int64, models_pkg.StatusEnum, *string, *string) (string, error)

    CreateRecordingTranscription (string, *string) (string, error)

    CreateViewTranscription (string, *string) (string, error)

    CreateAudioURLTranscription (string, *string) (string, error)
}

/*
 * Factory for the TRANSCRIPTION interaface returning TRANSCRIPTION_IMPL
 */
func NewTRANSCRIPTION() TRANSCRIPTION {
    return &TRANSCRIPTION_IMPL{}
}