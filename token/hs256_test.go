package token_test

import (
	"testing"

	"github.com/esiteltd/zaincash-golang/token"
)

func TestHS256_Decode(t *testing.T) {
	d := token.NewHS256(
		"$2y$10$hBbAZo2GfSSvyqAyV2SaqOfYewgYpfR1O19gIh4SqyGWdmySZYPuS",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGF0dXMiOiJzdWNjZXNzIiwib3JkZXJpZCI6ImZvciBleGFtcGxlIiwiaWQiOiI2NzkwZDFkOTE0MTMxOGQzMmQ5MWM5Y2IiLCJvcGVyYXRpb25pZCI6IjEyNDMyNzU1NDY2ODI1MDAiLCJtc2lzZG4iOiI5NjQ3ODAyOTk5NTY5In0.yCbogBmVURng9AuDkkjH-p_CqN7YPV2pDl3resRhCck",
	)

	if err := d.Decode(); err != nil {
		t.Fatal(err)
	}

	if !d.Succeed() {
		t.Fatal(d.Message())
	}

	t.Logf("ID: %s OrderID: %s IssuedAt: %d ExpiresAt: %d", d.TransactionID(), d.OrderID(), d.IssuedAt(), d.ExpiresAt())
}
