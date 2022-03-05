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
