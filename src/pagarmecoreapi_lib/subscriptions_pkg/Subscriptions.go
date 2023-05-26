/*
 * pagarmecoreapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package subscriptions_pkg

import "time"
import "pagarmecoreapi_lib/configuration_pkg"
import "pagarmecoreapi_lib/models_pkg"

/*
 * Interface for the SUBSCRIPTIONS_IMPL
 */
type SUBSCRIPTIONS interface {
    UpdateSubscriptionCard (string, *models_pkg.UpdateSubscriptionCardRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    DeleteUsage (string, string, string, *string) (*models_pkg.GetUsageResponse, error)

    CreateDiscount (string, *models_pkg.CreateDiscountRequest, *string) (*models_pkg.GetDiscountResponse, error)

    CreateAnUsage (string, string, *string) (*models_pkg.GetUsageResponse, error)

    GetUsages (string, string, *int64, *int64, *string, *string, *time.Time, *time.Time) (*models_pkg.ListUsagesResponse, error)

    UpdateCurrentCycleStatus (string, *models_pkg.UpdateCurrentCycleStatusRequest, *string) (error)

    UpdateSubscriptionPaymentMethod (string, *models_pkg.UpdateSubscriptionPaymentMethodRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    DeleteDiscount (string, string, *string) (*models_pkg.GetDiscountResponse, error)

    GetSubscriptionItems (string, *int64, *int64, *string, *string, *string, *string, *string, *string) (*models_pkg.ListSubscriptionItemsResponse, error)

    CreateSubscriptionItem (string, *models_pkg.CreateSubscriptionItemRequest, *string) (*models_pkg.GetSubscriptionItemResponse, error)

    GetSubscriptionItem (string, string) (*models_pkg.GetSubscriptionItemResponse, error)

    UpdateSubscriptionItem (string, string, *models_pkg.UpdateSubscriptionItemRequest, *string) (*models_pkg.GetSubscriptionItemResponse, error)

    GetSubscriptions (*int64, *int64, *string, *string, *string, *string, *string, *string, *time.Time, *time.Time, *time.Time, *time.Time) (*models_pkg.ListSubscriptionsResponse, error)

    CreateSubscription (*models_pkg.CreateSubscriptionRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    CancelSubscription (string, *string, *models_pkg.CreateCancelSubscriptionRequest) (*models_pkg.GetSubscriptionResponse, error)

    GetSubscription (string) (*models_pkg.GetSubscriptionResponse, error)

    CreateIncrement (string, *models_pkg.CreateIncrementRequest, *string) (*models_pkg.GetIncrementResponse, error)

    GetDiscountById (string, string) (*models_pkg.GetDiscountResponse, error)

    UpdateSubscriptionAffiliationId (string, *models_pkg.UpdateSubscriptionAffiliationIdRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionMetadata (string, *models_pkg.UpdateMetadataRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    DeleteIncrement (string, string, *string) (*models_pkg.GetIncrementResponse, error)

    GetIncrementById (string, string) (*models_pkg.GetIncrementResponse, error)

    GetSubscriptionCycles (string, string, string) (*models_pkg.ListCyclesResponse, error)

    RenewSubscription (string, *string) (*models_pkg.GetPeriodResponse, error)

    GetDiscounts (string, int64, int64) (*models_pkg.ListDiscountsResponse, error)

    UpdateSubscriptionBillingDate (string, *models_pkg.UpdateSubscriptionBillingDateRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    DeleteSubscriptionItem (string, string, *string) (*models_pkg.GetSubscriptionItemResponse, error)

    GetIncrements (string, *int64, *int64) (*models_pkg.ListIncrementsResponse, error)

    UpdateSubscriptionDueDays (string, *models_pkg.UpdateSubscriptionDueDaysRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionStartAt (string, *models_pkg.UpdateSubscriptionStartAtRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateLatestPeriodEndAt (string, *models_pkg.UpdateCurrentCycleEndDateRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionMiniumPrice (string, *models_pkg.UpdateSubscriptionMinimumPriceRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    GetSubscriptionCycleById (string, string) (*models_pkg.GetPeriodResponse, error)

    GetUsageReport (string, string) (*models_pkg.GetUsageReportResponse, error)

    UpdateSplitSubscription (string, *models_pkg.UpdateSubscriptionSplitRequest) (*models_pkg.GetSubscriptionResponse, error)
}

/*
 * Factory for the SUBSCRIPTIONS interaface returning SUBSCRIPTIONS_IMPL
 */
func NewSUBSCRIPTIONS(config configuration_pkg.CONFIGURATION) *SUBSCRIPTIONS_IMPL {
    client := new(SUBSCRIPTIONS_IMPL)
    client.config = config
    return client
}