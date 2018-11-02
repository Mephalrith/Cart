package main

import "math"

type Price struct {
	ProductType string              `json:"product-type"`
	Options     map[string][]string `json:"options"`
	BasePrice   int                 `json:"base-price"`
}

type Item struct {
	ProductType  string            `json:"product-type"`
	Options      map[string]string `json:"options"`
	ArtistMarkup int               `json:"artist-markup"`
	Quantity     int               `json:"quantity"`
}

// Adds the method to compare the options of the item with the options from a given base price.
// Using a counter lets us ignore options present on an item that are not present in a base price.
func (item Item) checkOptions(price Price) bool {
	counter := 0

	for key, vals := range price.Options {
		if val, ok := item.Options[key]; ok {
			for _, v := range vals {
				if v == val {
					counter++
				}
			}
		}
	}

	if counter == len(price.Options) {
		return true
	}
	return false
}

// Adds the method to calculate the price for an item, given its base price.
func (item Item) calculatePrice(price int) int {
	markup := float64(price) * (float64(item.ArtistMarkup) / 100)
	return (price + int(math.Round(markup))) * item.Quantity
}
