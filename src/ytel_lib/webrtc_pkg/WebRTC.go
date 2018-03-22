/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package webrtc_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the WEBRTC_IMPL
 */
type WEBRTC interface {
    CreateToken (*CreateTokenInput) (string, error)

    CheckFunds (*CheckFundsInput) (string, error)
}

/*
 * Factory for the WEBRTC interaface returning WEBRTC_IMPL
 */
func NewWEBRTC(config configuration_pkg.CONFIGURATION) *WEBRTC_IMPL {
    client := new(WEBRTC_IMPL)
    client.config = config
    return client
}