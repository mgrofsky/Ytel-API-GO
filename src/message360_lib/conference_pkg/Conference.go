/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package conference_pkg


/*
 * Interface for the CONFERENCE_IMPL
 */
type CONFERENCE interface {
    DeafMuteParticipant (*DeafMuteParticipantInput) (string, error)

    ViewParticipant (*ViewParticipantInput) (string, error)

    AddParticipant (*AddParticipantInput) (string, error)

    ViewConference (*ViewConferenceInput) (string, error)

    CreateConference (*CreateConferenceInput) (string, error)

    HangupParticipant (*HangupParticipantInput) (string, error)

    PlayConferenceAudio (*PlayConferenceAudioInput) (string, error)

    ListParticipant (*ListParticipantInput) (string, error)

    ListConference (*ListConferenceInput) (string, error)
}

/*
 * Factory for the CONFERENCE interaface returning CONFERENCE_IMPL
 */
func NewCONFERENCE() CONFERENCE {
    return &CONFERENCE_IMPL{}
}