/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package configuration_pkg

import(
	"ytel_lib/apihelper_pkg"

/* Setting up enums for Environments and Servers 
*/
type Environments int

type Servers int

const (
     PRODUCTION Environments = 1 + iota
     //Pre-Production environment used to test your code in our beta systems.  There is a good chance nothing will work here.  Don't use it unless instructed by our staff.
     PREPRODUCTION
)
const (
 	 DEFAULT Servers = 1 + iota
)

type CONFIGURATION_IMPL struct {
/** The username to use with basic authentication */
    /** Replace the value of basicauthusername with SetBasicAuthUserName function */
    basicauthusername string
/** The password to use with basic authentication */
    /** Replace the value of basicauthpassword with SetBasicAuthPassword function */
    basicauthpassword string
}
 
  
/*
 * Getter function returning basicauthusername
 */
func (me *CONFIGURATION_IMPL) BasicAuthUserName() string{
     return me.basicauthusername
}

/*
 * Getter function returning basicauthpassword
 */
func (me *CONFIGURATION_IMPL) BasicAuthPassword() string{
     return me.basicauthpassword
}

/*
 * Setter function setting up basicauthusername
 */
func (me *CONFIGURATION_IMPL) SetBasicAuthUserName(basicAuthUserName string) {
      me.basicauthusername = basicAuthUserName
}

/*
 * Setter function setting up basicauthpassword
 */
func (me *CONFIGURATION_IMPL) SetBasicAuthPassword(basicAuthPassword string) {
      me.basicauthpassword = basicAuthPassword
}

// Setting up Default Environment
var Environment = PRODUCTION

//A map of environments and their corresponding servers/baseurls
 var EnvironmentsMap = map[Environments](map[Servers]string){

    PRODUCTION : map[Servers]string{
        DEFAULT:"https://api.message360.com/api/v3",
    },

    PREPRODUCTION : map[Servers]string{
        DEFAULT:"https://api-preprod.message360.com/api/v2",
    },
}
 
// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
     kvpMap := map[string]interface{}{
    }
    return kvpMap;
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
    url := EnvironmentsMap[Environment][server]
    appendedURL, _ := apihelper_pkg.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
    return appendedURL

}
