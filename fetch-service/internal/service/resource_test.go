package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/currency_converter"
	mockCurrencyConverter "github.com/bagusays/efishery-technical-test/internal/shared/integration/currency_converter/mock"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/efishery"
	mockEfishery "github.com/bagusays/efishery-technical-test/internal/shared/integration/efishery/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestResource_ResourceStatistics(t *testing.T) {
	getTime := func(tm string) time.Time {
		t, _ := time.Parse("2006-01-02", tm)
		return t
	}

	testCases := []struct {
		name      string
		mockCache func()
		isErr     bool
		want      map[string][]model.ResourceStatistics
	}{
		{
			name: "should be succeed",
			mockCache: func() {
				internal.SetCache("usd", 14000, time.Now().Add(1*time.Hour))
				internal.SetCache("resource", []efishery.ResourceResponse{{
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         "3",
					Price:        "2",
					ParsedDate:   efishery.Time(getTime("2020-01-01")),
					Timestamp:    "timestamp",
				}, {
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         "7",
					Price:        "15",
					ParsedDate:   efishery.Time(getTime("2020-01-02")),
					Timestamp:    "timestamp",
				}, {
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         "5",
					Price:        "10",
					ParsedDate:   efishery.Time(getTime("2020-01-02")),
					Timestamp:    "timestamp",
				}}, time.Now().Add(1*time.Hour))
			},
			want: map[string][]model.ResourceStatistics{
				"byPrice": {{
					ProvinceArea: "provinceArea",
					Date:         "2020-1",
					Statistics: model.Statistics{
						Min:     2,
						Max:     15,
						Median:  10,
						Average: (float64(2) + float64(10) + float64(15)) / 3,
					},
				}},
				"bySize": {{
					ProvinceArea: "provinceArea",
					Date:         "2020-1",
					Statistics: model.Statistics{
						Min:     3,
						Max:     7,
						Median:  5,
						Average: 5,
					},
				}},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() {
				ctrl.Finish()
				internal.FlushCache()
			}()

			svc := NewResource(nil, nil)
			tt.mockCache()
			resources, err := svc.ResourceStatistics(context.Background())
			if tt.isErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.EqualValues(t, tt.want, resources)

			cache, err := internal.GetCache("resource")
			assert.NoError(t, err)
			assert.NotNil(t, cache)
		})
	}
}

func TestResource_FetchResource(t *testing.T) {
	tm := time.Now()

	testCases := []struct {
		name                  string
		mockEfishery          func(m *mockEfishery.MockClient)
		mockCurrencyConverter func(m *mockCurrencyConverter.MockClient)
		mockCache             func()
		isErr                 bool
		want                  interface{}
	}{
		{
			name:      "should be succeed",
			mockCache: func() {},
			mockEfishery: func(m *mockEfishery.MockClient) {
				m.EXPECT().FetchResource(gomock.Any()).Return([]efishery.ResourceResponse{{
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         json.Number("1"),
					Price:        "20000",
					ParsedDate:   efishery.Time(tm),
					Timestamp:    "timestamp",
				}, {
					UUID:         "",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         json.Number("1"),
					Price:        "20000",
					ParsedDate:   efishery.Time(tm),
					Timestamp:    "timestamp",
				}}, nil)
			},
			mockCurrencyConverter: func(m *mockCurrencyConverter.MockClient) {
				m.EXPECT().Convert(gomock.Any(), currency_converter.CurrencyUSD, currency_converter.CurrencyIDR).Return(float64(14000), nil)
			},
			want: []model.Resource{
				{
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         1,
					Price:        20000,
					ParsedDate:   tm,
					Timestamp:    "timestamp",
					PriceInUSD:   fmt.Sprintf("%.2f", float64(20000)/float64(14000)),
				}, {
					UUID:         "unknownUUID",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         1,
					Price:        20000,
					ParsedDate:   tm,
					Timestamp:    "timestamp",
					PriceInUSD:   fmt.Sprintf("%.2f", float64(20000)/float64(14000)),
				},
			},
		},
		{
			name: "should be succeed (from cache)",
			mockCache: func() {
				internal.SetCache("usd", 14000, tm.Add(1*time.Hour))
				internal.SetCache("resource", []efishery.ResourceResponse{{
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         json.Number("1"),
					Price:        "20000",
					ParsedDate:   efishery.Time(tm),
					Timestamp:    "timestamp",
				}}, tm.Add(1*time.Hour))
			},
			mockEfishery:          func(m *mockEfishery.MockClient) {},
			mockCurrencyConverter: func(m *mockCurrencyConverter.MockClient) {},
			want: []model.Resource{
				{
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         1,
					Price:        20000,
					ParsedDate:   tm,
					Timestamp:    "timestamp",
					PriceInUSD:   fmt.Sprintf("%.2f", float64(20000)/float64(14000)),
				},
			},
		},
		{
			name:      "should be failed when get resource from efishery return error",
			mockCache: func() {},
			mockEfishery: func(m *mockEfishery.MockClient) {
				m.EXPECT().FetchResource(gomock.Any()).Return(nil, fmt.Errorf("unknownErr"))
			},
			mockCurrencyConverter: func(m *mockCurrencyConverter.MockClient) {
				m.EXPECT().Convert(gomock.Any(), currency_converter.CurrencyUSD, currency_converter.CurrencyIDR).Return(float64(14000), nil)
			},
			isErr: true,
		},
		{
			name:      "should be failed when get usd from currency converter return error",
			mockCache: func() {},
			mockEfishery: func(m *mockEfishery.MockClient) {
				m.EXPECT().FetchResource(gomock.Any()).Return([]efishery.ResourceResponse{{
					UUID:         "uuid",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         json.Number("1"),
					Price:        "20000",
					ParsedDate:   efishery.Time(tm),
					Timestamp:    "timestamp",
				}}, nil)
			},
			mockCurrencyConverter: func(m *mockCurrencyConverter.MockClient) {
				m.EXPECT().Convert(gomock.Any(), currency_converter.CurrencyUSD, currency_converter.CurrencyIDR).Return(float64(0), fmt.Errorf("unknownErr"))
			},
			isErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() {
				ctrl.Finish()
				internal.FlushCache()
			}()

			efisheryClientMock := mockEfishery.NewMockClient(ctrl)
			tt.mockEfishery(efisheryClientMock)

			currencyConverterMock := mockCurrencyConverter.NewMockClient(ctrl)
			tt.mockCurrencyConverter(currencyConverterMock)

			svc := NewResource(currencyConverterMock, efisheryClientMock)
			tt.mockCache()
			resources, err := svc.FetchResource(context.Background())
			if tt.isErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.EqualValues(t, tt.want, resources)

			cache, err := internal.GetCache("resource")
			assert.NoError(t, err)
			assert.NotNil(t, cache)
		})
	}
}
