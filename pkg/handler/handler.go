package handler

import (
	"net/http"
	fs "stripe-example/external/firestore"
	st "stripe-example/external/stripe"
	"stripe-example/pkg/domain/output"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
)

type Handler struct {
	stService *st.Stripe
	fsService *fs.FireStore
}

func NewHandler(st *st.Stripe, fs *fs.FireStore) *Handler {
	return &Handler{
		stService: st,
		fsService: fs,
	}
}

func (h *Handler) Healthz(c echo.Context) error {
	stripe.Key = h.stService.Key

	id := "productId1"

	product_params := &stripe.ProductParams{
		ID:          &id,
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}
	starter_product, err := product.New(product_params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "intenal server error")
	}

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(1200),
	}
	starter_price, _ := price.New(price_params)

	pay := output.Payment{ProductID: starter_price.ID, PriceID: starter_price.ID}
	return c.JSON(http.StatusOK, pay)

}
