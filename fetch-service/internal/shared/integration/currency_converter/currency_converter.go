package currency_converter

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration"
)

//go:generate mockgen -destination=mock/mock_currency_converter.go -package=mock github.com/bagusays/efishery-technical-test/internal/shared/integration/currency_converter Client
type Client interface {
	Convert(ctx context.Context, from, to Currency) (float64, error)
}

type client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient() Client {
	timeout := 10 * time.Second
	if config.GetConfig().CurrencyConverter.HttpTimeout != 0 {
		timeout = config.GetConfig().CurrencyConverter.HttpTimeout
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}

	return &client{
		httpClient: httpClient,
		apiKey:     config.GetConfig().CurrencyConverter.APIKey,
	}
}

type ResourceResponse struct {
	UUID         string `json:"uuid"`
	Commodity    string `json:"komoditas"`
	ProvinceArea string `json:"area_provinsi"`
	CityArea     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	ParsedDate   string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyIDR Currency = "IDR"
)

func (c client) Convert(ctx context.Context, from, to Currency) (float64, error) {
	key := fmt.Sprintf("%s_%s", from, to)
	url := c.buildURL(fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra", key))

	var result map[string]float64
	err := integration.PerformRequest(ctx, integration.Request{
		HttpClient: c.httpClient,
		Method:     http.MethodGet,
		URL:        url,
	}, &result)
	if err != nil {
		return 0, err
	}

	return result[key], nil
}

// for now we only use string as URL instead of use url.URL
func (c client) buildURL(s string) string {
	return fmt.Sprintf("%s&apiKey=%s", s, c.apiKey)
}
