/*
 * message360_lib
 *
 * This file was automatically generated for message360 by APIMATIC v2.0 ( https://apimatic.io ) on 10/18/2016
 */

package email_pkg


/*
 * Interface for the EMAIL_IMPL
 */
type EMAIL interface {
    CreateSendEmail (string, string, string, string, string, *string, *string, []byte, *string) (string, error)

    CreateDeleteUnsubscribes (string, *string) (string, error)

    CreateListUnsubscribes (*string, *string, *string) (string, error)

    AddUnsubscribes (string, *string) (string, error)

    CreateDeleteSpam (string, *string) (string, error)

    CreateDeleteBlock (string, *string) (string, error)

    CreateListInvalid (*string, *string, *string) (string, error)

    CreateDeleteBounces (string, *string) (string, error)

    CreateListBounces (*string, *string, *string) (string, error)

    CreateListSpam (string, *string, *string) (string, error)

    CreateListBlocks (*string, *string, *string) (string, error)
}

/*
 * Factory for the EMAIL interaface returning EMAIL_IMPL
 */
func NewEMAIL() EMAIL {
    return &EMAIL_IMPL{}
}