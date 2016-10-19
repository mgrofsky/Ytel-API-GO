/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/18/2016
 */

package conference_pkg

import "message360_lib/models_pkg"

/*
 * Interface for the CONFERENCE_IMPL
 */
type CONFERENCE interface {
    CreateViewParticipant (string, string, *string) (string, error)

    CreateListParticipant (string, *int64, *int64, *bool, *bool, *string) (string, error)

    AddParticipant (string, string, int64, *bool, *bool, *string) (string, error)

    CreateViewConference (string, *string) (string, error)

    CreateListConference (*int64, *int64, *string, models_pkg.InterruptedCallStatusEnum, *string, *string, *string) (string, error)
}

/*
 * Factory for the CONFERENCE interaface returning CONFERENCE_IMPL
 */
func NewCONFERENCE() CONFERENCE {
    return &CONFERENCE_IMPL{}
}