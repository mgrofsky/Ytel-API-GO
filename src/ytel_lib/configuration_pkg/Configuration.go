/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package configuration_pkg



type CONFIGURATION interface {
        BasicAuthUserName()  string
        BasicAuthPassword()  string
        SetBasicAuthUserName(basicAuthUserName   string)
        SetBasicAuthPassword(basicAuthPassword   string)
}   
/*
 * Factory for the CONFIGURATION interaface returning CONFIGURATION_IMPL
 */
func NewCONFIGURATION() CONFIGURATION{
    configuration := new(CONFIGURATION_IMPL)
    return configuration
}