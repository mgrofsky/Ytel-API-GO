/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io )
 */

package email_pkg


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
func NewEMAIL() EMAIL {
    return &EMAIL_IMPL{}
}