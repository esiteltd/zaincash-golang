# ZainCash Client

This library is a simple client for [ZainCash](https://zaincash.iq/). It's currently under development, but key features have been implemented.

```bash
go get -u github.com/esiteltd/zaincash-golang
```

## Usage

The following example creates a transaction and ZainCash will redirect the customer to the localhost.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/esiteltd/zaincash"
)

func main() {
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

	id, err := p.Send_CreateTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(id)
}
```
