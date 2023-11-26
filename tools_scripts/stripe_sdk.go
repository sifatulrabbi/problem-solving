package tools_scripts

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/card"
	"github.com/stripe/stripe-go/v75/paymentmethod"
)

const customer = "cus_OQbZ1rp3VG1DA6"

func prepareStripe() {
	stripeKey := os.Getenv("STRIPE_LIVE_API_KEY")
	if stripeKey == "" {
		log.Fatalln(errors.New("no stripe key found among the env vars"))
	}
	stripe.Key = stripeKey
}

func GetCustomerCards() {
	prepareStripe()
	params := &stripe.CardListParams{
		Customer: stripe.String(customer),
	}
	list := card.List(params)
	cards := []*stripe.Card{}
	for list.Next() {
		c := list.Card()
		if c != nil {
			fmt.Println("card found")
		}
		cards = append(cards, c)
	}
	fmt.Println(cards)
}

func GetCustomerPaymentMethods() {
	prepareStripe()
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customer),
		Type:     stripe.String("card"),
	}
	paymentMethods := []*stripe.PaymentMethod{}
	i := paymentmethod.List(params)
	for i.Next() {
		paymentMethods = append(paymentMethods, i.PaymentMethod())
	}
	fmt.Println(paymentMethods)
}
