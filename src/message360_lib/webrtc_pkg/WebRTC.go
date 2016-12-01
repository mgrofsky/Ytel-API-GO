/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 12/01/2016
 */

package webrtc_pkg


/*
 * Interface for the WEBRTC_IMPL
 */
type WEBRTC interface {
    CreateToken (*CreateTokenInput) (error)

    CreateCheckFunds (*CreateCheckFundsInput) (error)

    CreateAuthenticateNumber (*CreateAuthenticateNumberInput) (error)
}

/*
 * Factory for the WEBRTC interaface returning WEBRTC_IMPL
 */
func NewWEBRTC() WEBRTC {
    return &WEBRTC_IMPL{}
}