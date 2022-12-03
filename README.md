# golang-stripe
stripe API

## Article
1. FireStroe
 - https://ema-hiro.hatenablog.com/entry/2020/09/28/022730
 - https://qiita.com/subaru44k/items/a88e638333b8d5cc29f2
 -

2. Stripe
 - https://stripe.com/docs/api
 - https://note.com/ogiogi93/n/na35ad141101a
 - https://github.com/ogiogi93/stripe-subscription-samples


3. Echo
 - https://echo.labstack.com/
 - https://github.com/labstack/echo

## API
```
$ curl -X GET http://www.localhost:8080
```
```json
{
	"ID": "97b2b954-2131-4e6c-82b1-e173e3f62706",
	"Title": "味噌ラーメンわくわく定額プラン",
	"Plans": [
		{
			"ID": "2de57c99-0360-4fbd-8397-9a56489d8998",
			"Title": "毎日ラーメン1杯無料プラン",
			"StripeProductID": "prod_MugIuC0Ppqqe92",
			"StripePriceID": "price_1MAqw8D41fJNajsLUBuU1xV0",
			"Price": 3000,
			"Benefits": [
				{
					"ID": "5770c18f-b9f3-431e-9510-574c9670529b",
					"Title": ""
				}
			]
		},
		{
			"ID": "ba9cdae8-84ed-4b37-bb94-11b3a94126b7",
			"Title": "トッピング毎回1品無料",
			"StripeProductID": "prod_MugI3kHqbbUaWP",
			"StripePriceID": "price_1MAqw8D41fJNajsL0p41Ie1a",
			"Price": 350,
			"Benefits": [
				{
					"ID": "ef96ef24-3802-46dd-9314-3287210454eb",
					"Title": ""
				}
			]
		}
	]
}
```
