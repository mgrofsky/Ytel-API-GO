/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package webrtc_pkg


/*
 * Interface for the WEBRTC_IMPL
 */
type WEBRTC interface {
    CreateCheckFunds (*CreateCheckFundsInput) (string, error)

    CreateToken (*CreateTokenInput) (string, error)
}

/*
 * Factory for the WEBRTC interaface returning WEBRTC_IMPL
 */
func NewWEBRTC() WEBRTC {
    return &WEBRTC_IMPL{}
}