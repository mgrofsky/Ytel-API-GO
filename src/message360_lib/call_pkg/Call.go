/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package call_pkg


/*
 * Interface for the CALL_IMPL
 */
type CALL interface {
    MakeCall (*MakeCallInput) (string, error)

    PlayAudio (*PlayAudioInput) (string, error)

    RecordCall (*RecordCallInput) (string, error)

    VoiceEffect (*VoiceEffectInput) (string, error)

    SendDigit (*SendDigitInput) (string, error)

    InterruptedCall (*InterruptedCallInput) (string, error)

    GroupCall (*GroupCallInput) (string, error)

    ListCalls (*ListCallsInput) (string, error)

    SendRinglessVM (*SendRinglessVMInput) (string, error)

    ViewCall (*ViewCallInput) (string, error)
}

/*
 * Factory for the CALL interaface returning CALL_IMPL
 */
func NewCALL() CALL {
    return &CALL_IMPL{}
}