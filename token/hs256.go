package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type HS256 struct {
	Key, Token string
	jwt.MapClaims
}

func NewHS256(key, token string) *HS256 {
	return &HS256{
		Key:   key,
		Token: token,
	}
}

func (h *HS256) Decode() error {
	t, err := jwt.ParseWithClaims(h.Token, &h.MapClaims, func(*jwt.Token) (interface{}, error) {
		return []byte(h.Key), nil
	})
	if err != nil {
		return fmt.Errorf("decode token: %w", err)
	}

	h.MapClaims = *t.Claims.(*jwt.MapClaims)
	return nil
}

func (h *HS256) Succeed() bool {
	v, found := h.MapClaims["status"]
	if !found {
		return false
	}

	status, ok := v.(string)
	if !ok {
		return false
	}

	return status == "success"
}

func (h *HS256) Message() string {
	v, found := h.MapClaims["msg"]
	if !found {
		return ""
	}

	msg, ok := v.(string)
	if !ok {
		return ""
	}

	return msg
}

func (h *HS256) TransactionID() string {
	v, found := h.MapClaims["id"]
	if !found {
		return ""
	}

	id, ok := v.(string)
	if !ok {
		return ""
	}

	return id
}

func (h *HS256) OrderID() string {
	v, found := h.MapClaims["orderid"]
	if !found {
		return ""
	}

	oid, ok := v.(string)
	if !ok {
		return ""
	}

	return oid
}

func (h *HS256) IssuedAt() int64 {
	v, found := h.MapClaims["iat"]
	if !found {
		return 0
	}

	iat, ok := v.(int64)
	if !ok {
		return 0
	}

	return iat
}

func (h *HS256) ExpiresAt() int64 {
	v, found := h.MapClaims["exp"]
	if !found {
		return 0
	}

	exp, ok := v.(int64)
	if !ok {
		return 0
	}

	return exp
}
