/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package letter_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the LETTER_IMPL
 */
type LETTER interface {
    ViewLetter (*ViewLetterInput) (string, error)

    CreateLetter (*CreateLetterInput) (string, error)

    ListLetters (*ListLettersInput) (string, error)

    DeleteLetter (*DeleteLetterInput) (string, error)
}

/*
 * Factory for the LETTER interaface returning LETTER_IMPL
 */
func NewLETTER(config configuration_pkg.CONFIGURATION) *LETTER_IMPL {
    client := new(LETTER_IMPL)
    client.config = config
    return client
}