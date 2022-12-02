package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
)

type Handler struct {
}

type payment struct {
	ProductID string `json:"product_id"`
	PriceID   string `json:"price_id"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) StripeHealthz(c echo.Context) error {
	stripe.Key = os.Getenv("SECRET_KEY")

	product_params := &stripe.ProductParams{
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}
	starter_product, _ := product.New(product_params)

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(1200),
	}
	starter_price, _ := price.New(price_params)

	pay := payment{ProductID: starter_price.ID, PriceID: starter_price.ID}
	return c.JSON(http.StatusOK, pay)

}
