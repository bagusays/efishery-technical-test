package service

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/currency_converter"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/efishery"
	"github.com/spf13/cast"
)

//go:generate mockgen -destination=mock/mock_resource.go -package=mock github.com/bagusays/efishery-technical-test/internal/service Resource
type Resource interface {
	FetchResource(ctx context.Context) ([]model.Resource, error)
	ResourceStatistics(ctx context.Context) (map[string][]model.ResourceStatistics, error)
}

type resource struct {
	currencyConverterClient currency_converter.Client
	efisheryClient          efishery.Client
}

func NewResource(currencyConverterClient currency_converter.Client, efisheryClient efishery.Client) Resource {
	return &resource{
		currencyConverterClient: currencyConverterClient,
		efisheryClient:          efisheryClient,
	}
}

func (r resource) ResourceStatistics(ctx context.Context) (map[string][]model.ResourceStatistics, error) {
	resources, err := r.FetchResource(ctx)
	if err != nil {
		return nil, err
	}

	byProvinceAndWeekly := make(map[string][]model.Resource)
	for _, d := range resources {
		key := fmt.Sprintf("%s|%d-%d", d.ProvinceArea, d.ParsedDate.Year(), d.ParsedDate.Month())
		if _, ok := byProvinceAndWeekly[key]; !ok {
			byProvinceAndWeekly[key] = []model.Resource{d}
		} else {
			byProvinceAndWeekly[key] = append(byProvinceAndWeekly[key], d)
		}
	}

	result := map[string][]model.ResourceStatistics{
		"byPrice": make([]model.ResourceStatistics, 0),
		"bySize":  make([]model.ResourceStatistics, 0),
	}
	for key, val := range byProvinceAndWeekly {
		date := strings.Split(key, "|")[1]
		byPrice, avgByPrice, medianByPrice := r.getStatisticByPrice(val)
		bySize, avgBySize, medianBySize := r.getStatisticBySize(val)
		result["byPrice"] = append(result["byPrice"], model.ResourceStatistics{
			ProvinceArea: byPrice[0].ProvinceArea,
			Date:         date,
			Statistics: model.Statistics{
				Min:     byPrice[0].Price,
				Max:     byPrice[len(byPrice)-1].Price,
				Median:  medianByPrice,
				Average: avgByPrice,
			},
		})
		result["bySize"] = append(result["bySize"], model.ResourceStatistics{
			ProvinceArea: bySize[0].ProvinceArea,
			Date:         date,
			Statistics: model.Statistics{
				Min:     bySize[0].Size,
				Max:     bySize[len(bySize)-1].Size,
				Median:  medianBySize,
				Average: avgBySize,
			},
		})

	}

	return result, nil
}

func (r resource) getStatisticByPrice(arr []model.Resource) ([]model.Resource, float64, float64) {
	avg := float64(arr[0].Price)
	for i := 1; i < len(arr); i++ {
		avg += arr[i].Price
		if arr[i-1].Price > arr[i].Price {
			for j := i; j > 0; j-- {
				if arr[j-1].Price > arr[j].Price {
					tmp := arr[j-1]
					arr[j-1] = arr[i]
					arr[i] = tmp
				}
			}
		}
	}

	median := len(arr) / 2
	if len(arr)%2 == 0 {
		median = int(math.Round(float64(len(arr)) / 2))
	}

	avg = avg / float64(len(arr))
	return arr, math.Round(avg*100) / 100, arr[median].Price
}

func (r resource) getStatisticBySize(arr []model.Resource) ([]model.Resource, float64, float64) {
	avg := float64(arr[0].Price)
	for i := 1; i < len(arr); i++ {
		avg += arr[i].Size
		if arr[i-1].Size > arr[i].Size {
			for j := i; j > 0; j-- {
				if arr[j-1].Size > arr[j].Size {
					tmp := arr[j-1]
					arr[j-1] = arr[i]
					arr[i] = tmp
				}
			}
		}
	}

	median := len(arr) / 2
	if len(arr)%2 == 0 {
		median = int(math.Round(float64(len(arr)) / 2))
	}

	avg = avg / float64(len(arr))
	return arr, math.Round(avg*100) / 100, arr[median].Size
}

func (r resource) FetchResource(ctx context.Context) ([]model.Resource, error) {
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
		// if it doesn't have price, skip
		if d.Price.String() == "" {
			continue
		}

		pricef64, err := d.Price.Float64()
		if err != nil {
			return nil, err
		}
		priceInUsd := fmt.Sprintf("%.2f", pricef64/usd)

		uuid := "unknownUUID"
		if d.UUID != "" {
			uuid = d.UUID
		}

		// if it doesn't have parsed date, skip
		if time.Time(d.ParsedDate).IsZero() {
			continue
		}

		size, err := d.Size.Float64()
		if err != nil {
			return nil, err
		}

		price, err := d.Price.Float64()
		if err != nil {
			return nil, err
		}

		finalResources[idx] = model.Resource{
			UUID:         uuid,
			Commodity:    d.Commodity,
			ProvinceArea: d.ProvinceArea,
			CityArea:     d.CityArea,
			Size:         size,
			Price:        price,
			PriceInUSD:   priceInUsd,
			ParsedDate:   time.Time(d.ParsedDate),
			Timestamp:    d.Timestamp,
		}
	}

	return finalResources, nil
}

func (r resource) getUSD(ctx context.Context) (float64, error) {
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

func (r resource) getResource(ctx context.Context) ([]efishery.ResourceResponse, error) {
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
