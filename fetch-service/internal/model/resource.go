package model

import "time"

type Resource struct {
	UUID         string    `json:"uuid"`
	Commodity    string    `json:"komoditas"`
	ProvinceArea string    `json:"area_provinsi"`
	CityArea     string    `json:"area_kota"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	PriceInUSD   string    `json:"priceInUsd"`
	ParsedDate   time.Time `json:"tgl_parsed"`
	Timestamp    string    `json:"timestamp"`
}

type ResourceStatistics struct {
	ProvinceArea string    `json:"province"`
	Date         time.Time `json:"date"`
	Min          string    `json:"min"`
	Max          string    `json:"max"`
	Median       string    `json:"median"`
	Average      string    `json:"average"`
}
