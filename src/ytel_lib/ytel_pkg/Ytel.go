/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package YtelClient

import(
	"ytel_lib/configuration_pkg"
	"ytel_lib/webrtc_pkg"
	"ytel_lib/sharedshortcode_pkg"
	"ytel_lib/conference_pkg"
	"ytel_lib/phone number_pkg"
	"ytel_lib/transcription_pkg"
	"ytel_lib/recording_pkg"
	"ytel_lib/email_pkg"
	"ytel_lib/sms_pkg"
	"ytel_lib/call_pkg"
	"ytel_lib/carrier_pkg"
	"ytel_lib/address_pkg"
	"ytel_lib/sub account_pkg"
	"ytel_lib/account_pkg"
	"ytel_lib/usage_pkg"
	"ytel_lib/shortcode_pkg"
	"ytel_lib/postcard_pkg"
	"ytel_lib/letter_pkg"
	"ytel_lib/areamail_pkg"
)


/*
 * Interface for the YTEL_IMPL
 */
type YTEL interface {
     WebRTC()                webrtc_pkg.WEBRTC
     SharedShortCode()       sharedshortcode_pkg.SHAREDSHORTCODE
     Conference()            conference_pkg.CONFERENCE
     PhoneNumber()          phone number_pkg.PHONE NUMBER
     Transcription()         transcription_pkg.TRANSCRIPTION
     Recording()             recording_pkg.RECORDING
     Email()                 email_pkg.EMAIL
     SMS()                   sms_pkg.SMS
     Call()                  call_pkg.CALL
     Carrier()               carrier_pkg.CARRIER
     Address()               address_pkg.ADDRESS
     SubAccount()           sub account_pkg.SUB ACCOUNT
     Account()               account_pkg.ACCOUNT
     Usage()                 usage_pkg.USAGE
     ShortCode()             shortcode_pkg.SHORTCODE
     PostCard()              postcard_pkg.POSTCARD
     Letter()                letter_pkg.LETTER
     AreaMail()              areamail_pkg.AREAMAIL
     Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the YTEL interaface returning YTEL_IMPL
 */
func NewYTEL(basicAuthUserName string, basicAuthPassword string) YTEL {
    ytelClient := new(YTEL_IMPL)
    ytelClient.config = configuration_pkg.NewCONFIGURATION()
    ytelClient.config.SetBasicAuthUserName(basicAuthUserName)
    ytelClient.config.SetBasicAuthPassword(basicAuthPassword)
    return ytelClient
}
