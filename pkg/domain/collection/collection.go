package collection

import (
	"time"

	"github.com/stripe/stripe-go"
)

type Plan struct {
	ID              string     `firestore:"id"`
	Title           string     `firestore:"title"`
	StripeProductID string     `firestore:"stripe_product_id"`
	StripePriceID   string     `firestore:"stripe_price_id"`
	Price           int32      `firestore:"price"`
	Benefits        []*Benefit `firestore:"benefits"`
}

// Benefit サブスクリプション適用のためのデータを定義(割引額等)。今回は触れない
type Benefit struct {
	ID    string `firestore:"id"`
	Title string `firestore:"title"`
	// DiscountValue int32 `firestore:"discount_value"`
}

// Subscription サブスクリプションは複数のプランを保持できる
type Subscription struct {
	ID    string  `firestore:"-"`
	Title string  `firestore:"title"`
	Plans []*Plan `firestore:"plans"`
}

// UserSubscription ユーザー毎のサブスクリプションプランの状態を定義
type UserSubscription struct {
	ID                    string                    `firestore:"-"`
	CustomerID            string                    `firestore:"customer_id"`
	SubscriptionID        string                    `firestore:"subscription_id"`
	PlanID                string                    `firestore:"plan_id"`
	NextPlanID            string                    `firestore:"next_plan_id"`
	Status                stripe.SubscriptionStatus `firestore:"status"`
	LatestPaymentIntentID string                    `firestore:"latest_payment_intent_id"`
	StartedAt             time.Time                 `firestore:"started_at"`

	StripeSubscriptionID     string `firestore:"stripe_subscription_id"`
	StripeSubscriptionItemID string `firestore:"stripe_subscription_item_id"`

	CurrentPeriodStart time.Time `firestore:"current_period_start"`
	CurrentPeriodEnd   time.Time `firestore:"current_period_end"`
}
