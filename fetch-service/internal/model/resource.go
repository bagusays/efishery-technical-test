package model

import (
	"time"
)

type Resource struct {
	UUID         string    `json:"uuid"`
	Commodity    string    `json:"komoditas"`
	ProvinceArea string    `json:"area_provinsi"`
	CityArea     string    `json:"area_kota"`
	Size         float64   `json:"size"`
	Price        float64   `json:"price"`
	PriceInUSD   string    `json:"priceInUsd"`
	ParsedDate   time.Time `json:"tgl_parsed"`
	Timestamp    string    `json:"timestamp"`
}

type ResourceStatistics struct {
	ProvinceArea string     `json:"province"`
	Date         string     `json:"date"`
	Statistics   Statistics `json:"statistics"`
}

type Statistics struct {
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
	Median  float64 `json:"median"`
	Average float64 `json:"average"`
}
