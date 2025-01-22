package zaincash_golang

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	TestHost = "test.zaincash.iq"
	LiveHost = "api.zaincash.iq"

	DefaultHost = TestHost
)

const (
	CreateTransactionEndpoint = "/transaction/init"
	GetTransactionEndpoint    = "/transaction/get"
)

const (
	English = "en"
	Arabic  = "ar"

	DefaultLanguage = English
)

type Provider struct {
	Host     string
	Language string

	MerchantID     string
	MerchantSecret string

	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

func (p *Provider) Send_CreateTransaction(ctx context.Context, tx Transaction) (id string, err error) {
	t, err := tx.Sign(p.MerchantSecret)
	if err != nil {
		return "", fmt.Errorf("encode token: %w", err)
	}

	data := url.Values{}
	data.Set("token", url.QueryEscape(t))
	data.Set("merchantId", p.MerchantID)
	data.Set("lang", p.Language)

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		fmt.Sprintf("https://%s%s", p.Host, CreateTransactionEndpoint),
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	res, err := p.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("unmarshal response: %w", err)
	}

	// Unfortunately, the status codes are not correct.

	v, found := response["id"]
	if !found {
		return "", fmt.Errorf("no id found in response: %s", body)
	}

	id, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("id is not string: %T", v)
	}

	return id, nil
}
