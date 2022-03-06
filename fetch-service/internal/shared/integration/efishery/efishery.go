package efishery

import (
	"context"
	"net/http"
	"time"

	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration"
)

//go:generate mockgen -destination=mock/mock_efishery.go -package=mock github.com/bagusays/efishery-technical-test/internal/shared/integration/efishery Client
type Client interface {
	FetchResource(ctx context.Context) ([]ResourceResponse, error)
}

type client struct {
	httpClient *http.Client
}

func NewClient() Client {
	timeout := 10 * time.Second
	if config.GetConfig().Efishery.HttpTimeout != 0 {
		timeout = config.GetConfig().Efishery.HttpTimeout
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}
	return &client{
		httpClient: httpClient,
	}
}

type ResourceResponse struct {
	UUID         string `json:"uuid"`
	Commodity    string `json:"komoditas"`
	ProvinceArea string `json:"area_provinsi"`
	CityArea     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	ParsedDate   Time   `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

func (c client) FetchResource(ctx context.Context) ([]ResourceResponse, error) {
	url := "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"

	var result []ResourceResponse
	err := integration.PerformRequest(ctx, integration.Request{
		HttpClient: c.httpClient,
		Method:     http.MethodGet,
		URL:        url,
	}, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
