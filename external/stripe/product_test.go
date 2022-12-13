package stripe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/product"
)

func TestStripe_GetProduct(t *testing.T) {
	stripe.Key = testStripe.Key
	params := &stripe.ProductParams{
		Name: stripe.String("product1"),
	}
	p, _ := product.New(params)

	t.Run("ok", func(t *testing.T) {
		got, err := testStripe.GetProduct(p.ID)
		assert.NoError(t, err)
		assert.NotEmpty(t, got)

		assert.Equal(t, "product1", got.Name)

		t.Cleanup(func() {
			_, err = product.Del(p.ID, nil)
			assert.NoError(t, err)
		})
	})
}

func TestStripe_CreateProduct(t *testing.T) {
	t.Run("ok: title1", func(t *testing.T) {
		pro, pri, err := testStripe.CreteaProduct("title1")
		assert.NoError(t, err)
		assert.NotEmpty(t, pro)
		assert.NotEmpty(t, pri)

		assert.Equal(t, "title1", pro.Name)

		t.Cleanup(func() {
			// _, err = product.Del(pro.ID, nil)
			// assert.NoError(t, err)
		})
	})

}
