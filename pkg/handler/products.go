package handler

import (
	"log"
	"net/http"
	"stripe-example/pkg/domain/collection"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/stripe/stripe-go/v74"
)

func (h *Handler) CreateSubscription(c echo.Context) error {
	sub := &collection.Subscription{
		ID:    "subscriptionId1",
		Title: "味噌ラーメンわくわく定額プラン",
		Plans: []*collection.Plan{
			{
				ID:    "planId1",
				Title: "毎日ラーメン1杯無料プラン",
				Price: 3000,
				Benefits: []*collection.Benefit{
					{
						ID: uuid.New().String(),
					},
				},
			},
			{
				ID:    "palnId2",
				Title: "トッピング毎回1品無料",
				Price: 350,
				Benefits: []*collection.Benefit{
					{
						ID: uuid.New().String(),
					},
				},
			},
		},
	}

	for _, plan := range sub.Plans {
		// Subscriptionの商品及び価格の詳細はこちら: https://stripe.com/docs/billing/prices-guide

		// Productの作成 https://stripe.com/docs/api/products/create
		productParams := &stripe.ProductParams{
			Name:                stripe.String(plan.Title),
			StatementDescriptor: stripe.String("Chompy"), // 明細書に記載する文字列. 5 ~ 22文字でアルファベットと数字のみなので注意. https://stripe.com/docs/statement-descriptors
		}
		productParams.AddMetadata("subscription_id", sub.ID)
		productParams.AddMetadata("plan_id", plan.ID)

		product, _ := h.stService.StripeClient.Products.New(productParams)

		// Priceの作成 https://stripe.com/docs/api/prices/create
		priceParams := &stripe.PriceParams{
			Currency: stripe.String(string(stripe.CurrencyJPY)), // 通貨の設定, JPYを設定する
			Product:  stripe.String(product.ID),                 // 上記で作成したProductのIDを設定する
			Recurring: &stripe.PriceRecurringParams{ // サブスク期間の設定
				Interval:      stripe.String("day"), // 日毎
				IntervalCount: stripe.Int64(30),     // 30日
			},
			UnitAmount: stripe.Int64(3000), // 料金, 3000円
		}
		priceParams.AddMetadata("subscription_id", sub.ID)
		priceParams.AddMetadata("plan_id", plan.ID)
		price, _ := h.stService.StripeClient.Prices.New(priceParams)

		plan.StripeProductID = product.ID
		plan.StripePriceID = price.ID
	}

	err := h.fsService.CreateSubscription(sub)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, sub)
}
