/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package call_pkg

import "ytel_lib/configuration_pkg"

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

    ListCalls (*ListCallsInput) (string, error)

    SendRinglessVM (*SendRinglessVMInput) (string, error)

    ViewCall (*ViewCallInput) (string, error)

    ViewCallDetail (string, string) (string, error)

    GroupCall (*GroupCallInput) (string, error)
}

/*
 * Factory for the CALL interaface returning CALL_IMPL
 */
func NewCALL(config configuration_pkg.CONFIGURATION) *CALL_IMPL {
    client := new(CALL_IMPL)
    client.config = config
    return client
}