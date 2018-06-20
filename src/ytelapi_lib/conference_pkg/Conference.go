/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package conference_pkg

import "ytelapi_lib/configuration_pkg"

/*
 * Interface for the CONFERENCE_IMPL
 */
type CONFERENCE interface {
    CreateConferencesPlayAudio (*CreateConferencesPlayAudioInput) (string, error)

    CreateConferencesHangupParticipant (*CreateConferencesHangupParticipantInput) (string, error)

    CreateConferencesViewconference (string) (string, error)

    CreateConferencesListconference (*CreateConferencesListconferenceInput) (string, error)

    CreateConferencesListParticipant (*CreateConferencesListParticipantInput) (string, error)

    CreateConferencesViewParticipant (*CreateConferencesViewParticipantInput) (string, error)

    CreateConferencesAddParticipant (*CreateConferencesAddParticipantInput) (string, error)

    CreateConferencesCreateConference (*CreateConferencesCreateConferenceInput) (string, error)

    CreateConferencesDeafMuteParticipant (*CreateConferencesDeafMuteParticipantInput) (string, error)
}

/*
 * Factory for the CONFERENCE interaface returning CONFERENCE_IMPL
 */
func NewCONFERENCE(config configuration_pkg.CONFIGURATION) *CONFERENCE_IMPL {
    client := new(CONFERENCE_IMPL)
    client.config = config
    return client
}