package examples_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	zaincash "github.com/esiteltd/zaincash-golang"
)

func TestCreateTransaction(t *testing.T) {
	p := &zaincash.Provider{
		Host:           zaincash.DefaultHost,
		Language:       zaincash.DefaultLanguage,
		MerchantID:     "5ffacf6612b5777c6d44266f",
		MerchantSecret: "$2y$10$hBbAZo2GfSSvyqAyV2SaqOfYewgYpfR1O19gIh4SqyGWdmySZYPuS",
		HTTPClient:     http.DefaultClient,
	}

	tx := zaincash.Transaction{
		Amount:            1000,
		ServiceType:       "A book",
		WalletPhoneNumber: 9647835077893,
		OrderID:           "for example",
		RedirectionURL:    "http://localhost",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now())},
	}

	id, err := p.Send_CreateTransaction(t.Context(), tx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)
}

func TestCheckTransaction(t *testing.T) {
	p := &zaincash.Provider{
		Host:           zaincash.DefaultHost,
		Language:       zaincash.DefaultLanguage,
		MerchantID:     "5ffacf6612b5777c6d44266f",
		MerchantSecret: "$2y$10$hBbAZo2GfSSvyqAyV2SaqOfYewgYpfR1O19gIh4SqyGWdmySZYPuS",
		HTTPClient:     http.DefaultClient,
	}

	id, err := p.Send_CreateTransaction(t.Context(), zaincash.Transaction{
		Amount:            1000,
		ServiceType:       "A book",
		WalletPhoneNumber: 9647835077893,
		OrderID:           "for example",
		RedirectionURL:    "http://localhost",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now())},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)

	status, err := p.CheckTransaction(t.Context(), zaincash.Transaction{
		ID:                id,
		WalletPhoneNumber: 9647835077893,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(status)
}
