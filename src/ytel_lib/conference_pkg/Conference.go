/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package conference_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the CONFERENCE_IMPL
 */
type CONFERENCE interface {
    DeafMuteParticipant (*DeafMuteParticipantInput) (string, error)

    ViewParticipant (*ViewParticipantInput) (string, error)

    ViewConference (*ViewConferenceInput) (string, error)

    AddParticipant (*AddParticipantInput) (string, error)

    CreateConference (*CreateConferenceInput) (string, error)

    HangupParticipant (*HangupParticipantInput) (string, error)

    PlayConferenceAudio (*PlayConferenceAudioInput) (string, error)

    ListParticipant (*ListParticipantInput) (string, error)

    ListConference (*ListConferenceInput) (string, error)
}

/*
 * Factory for the CONFERENCE interaface returning CONFERENCE_IMPL
 */
func NewCONFERENCE(config configuration_pkg.CONFIGURATION) *CONFERENCE_IMPL {
    client := new(CONFERENCE_IMPL)
    client.config = config
    return client
}