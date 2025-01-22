package zaincash_golang

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Transaction struct {
	Amount            int64  `json:"amount,omitempty"`
	ServiceType       string `json:"serviceType,omitempty"`
	WalletPhoneNumber int64  `json:"msisdn,omitempty"`
	OrderID           string `json:"orderId,omitempty"`
	RedirectionURL    string `json:"redirectUrl,omitempty"`

	jwt.RegisteredClaims
}

func (t *Transaction) Sign(key string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, *t).SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("signing token: %v", err)
	}
	return token, nil
}
