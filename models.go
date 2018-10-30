package main

import (
	"math"
)

type Price struct {
	ProductType string              `json:"product-type"`
	Options     map[string][]string `json:"options"`
	BasePrice   int                 `json:"base-price"`
}

// Adds the method to get all options keys from a price as an array of strings.
func (price Price) getOptionsKeys() []string {
	var keys []string
	for key, option := range price.Options {
		for _, val := range option {
			keys = append(keys, key+":"+val)
		}
	}
	return keys
}

type Item struct {
	ProductType  string            `json:"product-type"`
	Options      map[string]string `json:"options"`
	ArtistMarkup int               `json:"artist-markup"`
	Quantity     int               `json:"quantity"`
}

// Adds the method to get all options keys, as well as the product type, as an array of strings.
func (item Item) getAllKeys() []string {
	var keys []string
	keys = append(keys, item.ProductType)
	for key, val := range item.Options {
		keys = append(keys, key+":"+val)
	}
	return keys
}

// Adds the method to calculate the price for an item, given it's base price.
func (item Item) calculatePrice(price int) int {
	markup := float64(price) * (float64(item.ArtistMarkup) / 100)
	return (price + int(Round(markup))) * item.Quantity
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
