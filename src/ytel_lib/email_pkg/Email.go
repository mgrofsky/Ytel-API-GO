/*
 * ytel_lib
 *
 * This file was automatically generated for ytel by APIMATIC v2.0 ( https://apimatic.io )
 */

package email_pkg

import "ytel_lib/configuration_pkg"

/*
 * Interface for the EMAIL_IMPL
 */
type EMAIL interface {
    DeleteSpam (*DeleteSpamInput) (string, error)

    DeleteBlock (*DeleteBlockInput) (string, error)

    AddUnsubscribes (*AddUnsubscribesInput) (string, error)

    SendEmail (*SendEmailInput) (string, error)

    DeleteUnsubscribes (*DeleteUnsubscribesInput) (string, error)

    ListUnsubscribes (*ListUnsubscribesInput) (string, error)

    ListInvalid (*ListInvalidInput) (string, error)

    DeleteBounces (*DeleteBouncesInput) (string, error)

    ListBounces (*ListBouncesInput) (string, error)

    ListSpam (*ListSpamInput) (string, error)

    ListBlocks (*ListBlocksInput) (string, error)

    DeleteInvalid (*DeleteInvalidInput) (string, error)
}

/*
 * Factory for the EMAIL interaface returning EMAIL_IMPL
 */
func NewEMAIL(config configuration_pkg.CONFIGURATION) *EMAIL_IMPL {
    client := new(EMAIL_IMPL)
    client.config = config
    return client
}