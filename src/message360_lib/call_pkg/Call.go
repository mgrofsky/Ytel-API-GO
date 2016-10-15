/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/15/2016
 */

package call_pkg

import "message360_lib/models_pkg"

/*
 * Interface for the CALL_IMPL
 */
type CALL interface {
    CreateViewCall (string, *string) (string, error)

    CreateMakeCall (string, string, string, string, string, models_pkg.HttpMethodEnum, *string, models_pkg.HttpMethodEnum, *string, models_pkg.HttpMethodEnum, *string, *bool, *int64, *string, *bool, *bool, *string, models_pkg.HttpMethodEnum, *bool, *string, models_pkg.IfMachineEnum, *string) (string, error)

    CreatePlayAudio (int64, models_pkg.DirectionEnum, bool, bool, *string, *string, *string) (string, error)

    CreateRecordCall (string, bool, models_pkg.DirectionEnum, *int64, *string, models_pkg.AudioFormatEnum, *string) (string, error)

    CreateVoiceEffect (string, models_pkg.AudioDirectionEnum, *float64, *float64, *float64, *float64, *float64, *string) (string, error)

    CreateSendDigit (string, string, models_pkg.DirectionEnum, *string) (string, error)

    CreateInterruptedCall (string, *string, models_pkg.HttpMethodEnum, models_pkg.InterruptedCallStatusEnum, *string) (string, error)

    CreateListCalls (*string, *string, *string, *string, *string, *string) (error)
}

/*
 * Factory for the CALL interaface returning CALL_IMPL
 */
func NewCALL() CALL {
    return &CALL_IMPL{}
}