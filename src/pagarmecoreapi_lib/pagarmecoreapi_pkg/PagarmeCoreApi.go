/*
 * pagarmecoreapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package PagarmeCoreApiClient

import(
	"pagarmecoreapi_lib/configuration_pkg"
	"pagarmecoreapi_lib/plans_pkg"
	"pagarmecoreapi_lib/subscriptions_pkg"
	"pagarmecoreapi_lib/invoices_pkg"
	"pagarmecoreapi_lib/orders_pkg"
	"pagarmecoreapi_lib/customers_pkg"
	"pagarmecoreapi_lib/recipients_pkg"
	"pagarmecoreapi_lib/charges_pkg"
	"pagarmecoreapi_lib/transfers_pkg"
	"pagarmecoreapi_lib/tokens_pkg"
	"pagarmecoreapi_lib/sellers_pkg"
	"pagarmecoreapi_lib/transactions_pkg"
)


/*
 * Interface for the PAGARMECOREAPI_IMPL
 */
type PAGARMECOREAPI interface {
    Plans()                 plans_pkg.PLANS
    Subscriptions()         subscriptions_pkg.SUBSCRIPTIONS
    Invoices()              invoices_pkg.INVOICES
    Orders()                orders_pkg.ORDERS
    Customers()             customers_pkg.CUSTOMERS
    Recipients()            recipients_pkg.RECIPIENTS
    Charges()               charges_pkg.CHARGES
    Transfers()             transfers_pkg.TRANSFERS
    Tokens()                tokens_pkg.TOKENS
    Sellers()               sellers_pkg.SELLERS
    Transactions()          transactions_pkg.TRANSACTIONS
    Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the PAGARMECOREAPI interface returning PAGARMECOREAPI_IMPL
 */
func NewPAGARMECOREAPI(basicAuthUserName string, basicAuthPassword string) PAGARMECOREAPI {
    pagarmeCoreApiClient := new(PAGARMECOREAPI_IMPL)
    pagarmeCoreApiClient.config = configuration_pkg.NewCONFIGURATION()

    pagarmeCoreApiClient.config.SetBasicAuthUserName(basicAuthUserName)
    pagarmeCoreApiClient.config.SetBasicAuthPassword(basicAuthPassword)

    return pagarmeCoreApiClient
}
