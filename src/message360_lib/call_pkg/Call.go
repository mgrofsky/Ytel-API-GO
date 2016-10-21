/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/21/2016
 */

package call_pkg

import "message360_lib/models_pkg"

/*
 * Interface for the CALL_IMPL
 */
type CALL interface {
    CreateViewCall (string, *string) (string, error)

    CreateMakeCall (string, string, string, string, string, models_pkg.HttpMethod, *string, models_pkg.HttpMethod, *string, models_pkg.HttpMethod, *string, *bool, *int64, *string, *bool, *bool, *string, models_pkg.HttpMethod, *bool, *string, models_pkg.IfMachine, *string) (string, error)

    CreatePlayAudio (int64, models_pkg.Direction, bool, bool, *string, *string, *string) (string, error)

    CreateRecordCall (string, bool, models_pkg.Direction, *int64, *string, models_pkg.AudioFormat, *string) (string, error)

    CreateVoiceEffect (string, models_pkg.AudioDirection, *float64, *float64, *float64, *float64, *float64, *string) (string, error)

    CreateSendDigit (string, string, models_pkg.Direction, *string) (string, error)

    CreateInterruptedCall (string, *string, models_pkg.HttpMethod, models_pkg.InterruptedCallStatus, *string) (string, error)

    CreateListCalls (*string, *string, *string, *string, *string, *string) (error)
}

/*
 * Factory for the CALL interaface returning CALL_IMPL
 */
func NewCALL() CALL {
    return &CALL_IMPL{}
}