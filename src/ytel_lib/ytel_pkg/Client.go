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
 * Client structure as interface implementation
 */
type YTEL_IMPL struct {
     webrtc webrtc_pkg.WEBRTC
     sharedshortcode sharedshortcode_pkg.SHAREDSHORTCODE
     conference conference_pkg.CONFERENCE
     phone number phone number_pkg.PHONE NUMBER
     transcription transcription_pkg.TRANSCRIPTION
     recording recording_pkg.RECORDING
     email email_pkg.EMAIL
     sms sms_pkg.SMS
     call call_pkg.CALL
     carrier carrier_pkg.CARRIER
     address address_pkg.ADDRESS
     sub account sub account_pkg.SUB ACCOUNT
     account account_pkg.ACCOUNT
     usage usage_pkg.USAGE
     shortcode shortcode_pkg.SHORTCODE
     postcard postcard_pkg.POSTCARD
     letter letter_pkg.LETTER
     areamail areamail_pkg.AREAMAIL
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *YTEL_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to WebRTC controller
     * @return Returns the WebRTC() instance
*/
func (me * YTEL_IMPL) WebRTC() webrtc_pkg.WEBRTC {
    if(me.webrtc) == nil {
        me.webrtc = webrtc_pkg.NewWEBRTC(me.config)
    }
    return me.webrtc
}
/**
     * Access to SharedShortCode controller
     * @return Returns the SharedShortCode() instance
*/
func (me * YTEL_IMPL) SharedShortCode() sharedshortcode_pkg.SHAREDSHORTCODE {
    if(me.sharedshortcode) == nil {
        me.sharedshortcode = sharedshortcode_pkg.NewSHAREDSHORTCODE(me.config)
    }
    return me.sharedshortcode
}
/**
     * Access to Conference controller
     * @return Returns the Conference() instance
*/
func (me * YTEL_IMPL) Conference() conference_pkg.CONFERENCE {
    if(me.conference) == nil {
        me.conference = conference_pkg.NewCONFERENCE(me.config)
    }
    return me.conference
}
/**
     * Access to PhoneNumber controller
     * @return Returns the PhoneNumber() instance
*/
func (me * YTEL_IMPL) PhoneNumber() phone number_pkg.PHONE NUMBER {
    if(me.phone number) == nil {
        me.phone number = phone number_pkg.NewPHONE NUMBER(me.config)
    }
    return me.phone number
}
/**
     * Access to Transcription controller
     * @return Returns the Transcription() instance
*/
func (me * YTEL_IMPL) Transcription() transcription_pkg.TRANSCRIPTION {
    if(me.transcription) == nil {
        me.transcription = transcription_pkg.NewTRANSCRIPTION(me.config)
    }
    return me.transcription
}
/**
     * Access to Recording controller
     * @return Returns the Recording() instance
*/
func (me * YTEL_IMPL) Recording() recording_pkg.RECORDING {
    if(me.recording) == nil {
        me.recording = recording_pkg.NewRECORDING(me.config)
    }
    return me.recording
}
/**
     * Access to Email controller
     * @return Returns the Email() instance
*/
func (me * YTEL_IMPL) Email() email_pkg.EMAIL {
    if(me.email) == nil {
        me.email = email_pkg.NewEMAIL(me.config)
    }
    return me.email
}
/**
     * Access to SMS controller
     * @return Returns the SMS() instance
*/
func (me * YTEL_IMPL) SMS() sms_pkg.SMS {
    if(me.sms) == nil {
        me.sms = sms_pkg.NewSMS(me.config)
    }
    return me.sms
}
/**
     * Access to Call controller
     * @return Returns the Call() instance
*/
func (me * YTEL_IMPL) Call() call_pkg.CALL {
    if(me.call) == nil {
        me.call = call_pkg.NewCALL(me.config)
    }
    return me.call
}
/**
     * Access to Carrier controller
     * @return Returns the Carrier() instance
*/
func (me * YTEL_IMPL) Carrier() carrier_pkg.CARRIER {
    if(me.carrier) == nil {
        me.carrier = carrier_pkg.NewCARRIER(me.config)
    }
    return me.carrier
}
/**
     * Access to Address controller
     * @return Returns the Address() instance
*/
func (me * YTEL_IMPL) Address() address_pkg.ADDRESS {
    if(me.address) == nil {
        me.address = address_pkg.NewADDRESS(me.config)
    }
    return me.address
}
/**
     * Access to SubAccount controller
     * @return Returns the SubAccount() instance
*/
func (me * YTEL_IMPL) SubAccount() sub account_pkg.SUB ACCOUNT {
    if(me.sub account) == nil {
        me.sub account = sub account_pkg.NewSUB ACCOUNT(me.config)
    }
    return me.sub account
}
/**
     * Access to Account controller
     * @return Returns the Account() instance
*/
func (me * YTEL_IMPL) Account() account_pkg.ACCOUNT {
    if(me.account) == nil {
        me.account = account_pkg.NewACCOUNT(me.config)
    }
    return me.account
}
/**
     * Access to Usage controller
     * @return Returns the Usage() instance
*/
func (me * YTEL_IMPL) Usage() usage_pkg.USAGE {
    if(me.usage) == nil {
        me.usage = usage_pkg.NewUSAGE(me.config)
    }
    return me.usage
}
/**
     * Access to ShortCode controller
     * @return Returns the ShortCode() instance
*/
func (me * YTEL_IMPL) ShortCode() shortcode_pkg.SHORTCODE {
    if(me.shortcode) == nil {
        me.shortcode = shortcode_pkg.NewSHORTCODE(me.config)
    }
    return me.shortcode
}
/**
     * Access to PostCard controller
     * @return Returns the PostCard() instance
*/
func (me * YTEL_IMPL) PostCard() postcard_pkg.POSTCARD {
    if(me.postcard) == nil {
        me.postcard = postcard_pkg.NewPOSTCARD(me.config)
    }
    return me.postcard
}
/**
     * Access to Letter controller
     * @return Returns the Letter() instance
*/
func (me * YTEL_IMPL) Letter() letter_pkg.LETTER {
    if(me.letter) == nil {
        me.letter = letter_pkg.NewLETTER(me.config)
    }
    return me.letter
}
/**
     * Access to AreaMail controller
     * @return Returns the AreaMail() instance
*/
func (me * YTEL_IMPL) AreaMail() areamail_pkg.AREAMAIL {
    if(me.areamail) == nil {
        me.areamail = areamail_pkg.NewAREAMAIL(me.config)
    }
    return me.areamail
}
