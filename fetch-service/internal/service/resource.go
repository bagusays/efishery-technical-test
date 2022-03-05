package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/currency_converter"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/efishery"
)

type Resource struct {
	currencyConverterClient *currency_converter.Client
	efisheryClient          *efishery.Client
}

func NewResource(currencyConverterClient *currency_converter.Client, efisheryClient *efishery.Client) *Resource {
	return &Resource{
		currencyConverterClient: currencyConverterClient,
		efisheryClient:          efisheryClient,
	}
}

func (r Resource) FetchResource(ctx context.Context) ([]model.Resource, error) {
	var (
		resources []efishery.ResourceResponse
		usd       float64
		err       error
	)

	errG, _ := errgroup.WithContext(context.Background())
	errG.Go(func() error {
		resources, err = r.efisheryClient.FetchResource(ctx)
		return err
	})

	errG.Go(func() error {
		usd, err = r.currencyConverterClient.Convert(ctx, currency_converter.CurrencyUSD, currency_converter.CurrencyIDR)
		return err
	})

	if err := errG.Wait(); err != nil {
		return nil, err
	}

	finalResources := make([]model.Resource, len(resources))
	for idx, d := range resources {
		priceInUsd := ""
		if d.Price != "" {
			pricef64, err := strconv.ParseFloat(d.Price, 64)
			if err != nil {
				return nil, err
			}
			priceInUsd = fmt.Sprintf("%.2f", pricef64/usd)
		}

		finalResources[idx] = model.Resource{
			UUID:         d.UUID,
			Commodity:    d.Commodity,
			ProvinceArea: d.ProvinceArea,
			CityArea:     d.CityArea,
			Size:         d.Size,
			Price:        d.Price,
			PriceInUSD:   priceInUsd,
			ParsedDate:   time.Time(d.ParsedDate),
			Timestamp:    d.Timestamp,
		}
	}

	return finalResources, nil
}
