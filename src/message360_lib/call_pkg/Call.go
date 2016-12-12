/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/12/2016
 */

package call_pkg


/*
 * Interface for the CALL_IMPL
 */
type CALL interface {
    CreateViewCall (*CreateViewCallInput) (string, error)

    CreateGroupCall (*CreateGroupCallInput) (string, error)

    CreateVoiceEffect (*CreateVoiceEffectInput) (string, error)

    CreateRecordCall (*CreateRecordCallInput) (string, error)

    CreatePlayAudio (*CreatePlayAudioInput) (string, error)

    CreateListCalls (*CreateListCallsInput) (string, error)

    CreateInterruptedCall (*CreateInterruptedCallInput) (string, error)

    CreateSendDigit (*CreateSendDigitInput) (string, error)

    CreateMakeCall (*CreateMakeCallInput) (string, error)
}

/*
 * Factory for the CALL interaface returning CALL_IMPL
 */
func NewCALL() CALL {
    return &CALL_IMPL{}
}