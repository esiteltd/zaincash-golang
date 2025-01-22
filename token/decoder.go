package token

type Decoder interface {
	Decode() error

	Succeed() bool
	Message() string

	TransactionID() string
	OrderID() string

	IssuedAt() int64
	ExpiresAt() int64
}
