/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/08/2016
 */

package conference_pkg


/*
 * Interface for the CONFERENCE_IMPL
 */
type CONFERENCE interface {
    CreateDeafMuteParticipant (*CreateDeafMuteParticipantInput) (string, error)

    CreateListConference (*CreateListConferenceInput) (string, error)

    CreateViewConference (*CreateViewConferenceInput) (string, error)

    AddParticipant (*AddParticipantInput) (string, error)

    CreateListParticipant (*CreateListParticipantInput) (string, error)

    CreateViewParticipant (*CreateViewParticipantInput) (string, error)
}

/*
 * Factory for the CONFERENCE interaface returning CONFERENCE_IMPL
 */
func NewCONFERENCE() CONFERENCE {
    return &CONFERENCE_IMPL{}
}