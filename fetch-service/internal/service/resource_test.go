package service

import (
	"context"
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
					Size:         "size",
					Price:        "20000",
					ParsedDate:   efishery.Time(tm),
					Timestamp:    "timestamp",
				}, {
					UUID:         "",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         "size",
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
					Size:         "size",
					Price:        "20000",
					ParsedDate:   tm,
					Timestamp:    "timestamp",
					PriceInUSD:   fmt.Sprintf("%.2f", float64(20000)/float64(14000)),
				}, {
					UUID:         "unknownUUID",
					Commodity:    "commodity",
					ProvinceArea: "provinceArea",
					CityArea:     "cityArea",
					Size:         "size",
					Price:        "20000",
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
					Size:         "size",
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
					Size:         "size",
					Price:        "20000",
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
					Size:         "size",
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
