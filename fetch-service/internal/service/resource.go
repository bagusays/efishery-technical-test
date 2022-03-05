package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/currency_converter"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/efishery"
	"github.com/spf13/cast"
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
		resources, err = r.getResource(ctx)
		return err
	})

	errG.Go(func() error {
		usd, err = r.getUSD(ctx)
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

func (r Resource) getUSD(ctx context.Context) (float64, error) {
	key := "usd"
	usd, err := internal.GetCache(key)
	if err == nil {
		f, err := cast.ToFloat64E(usd)
		if err != nil {
			return 0, err
		}
		return f, nil
	}

	newUsd, err := r.currencyConverterClient.Convert(ctx, currency_converter.CurrencyUSD, currency_converter.CurrencyIDR)
	if err != nil {
		return 0, err
	}

	internal.SetCache(key, newUsd, time.Now().Add(24*time.Hour))
	return newUsd, nil
}

func (r Resource) getResource(ctx context.Context) ([]efishery.ResourceResponse, error) {
	key := "resource"
	data, err := internal.GetCache(key)
	if err == nil {
		if d, ok := data.([]efishery.ResourceResponse); ok {
			return d, nil
		}
	}

	resources, err := r.efisheryClient.FetchResource(ctx)
	if err != nil {
		return nil, err
	}

	internal.SetCache(key, resources, time.Now().Add(24*time.Hour))
	return resources, nil
}
