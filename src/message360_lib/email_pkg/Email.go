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
    CreateDeleteInvalid (*CreateDeleteInvalidInput) (string, error)

    CreateListBlocks (*CreateListBlocksInput) (string, error)

    CreateListSpam (*CreateListSpamInput) (string, error)

    CreateListBounces (*CreateListBouncesInput) (string, error)

    CreateDeleteBounces (*CreateDeleteBouncesInput) (string, error)

    CreateListInvalid (*CreateListInvalidInput) (string, error)

    CreateListUnsubscribes (*CreateListUnsubscribesInput) (string, error)

    CreateDeleteUnsubscribes (*CreateDeleteUnsubscribesInput) (string, error)

    AddUnsubscribes (*AddUnsubscribesInput) (string, error)

    CreateDeleteBlock (*CreateDeleteBlockInput) (string, error)

    CreateDeleteSpam (*CreateDeleteSpamInput) (string, error)

    CreateSendEmail (*CreateSendEmailInput) (string, error)
}

/*
 * Factory for the EMAIL interaface returning EMAIL_IMPL
 */
func NewEMAIL() EMAIL {
    return &EMAIL_IMPL{}
}