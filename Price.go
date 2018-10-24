package main

type Price struct {
	ProductType string              `json:"product-type"`
	Options     map[string][]string `json:"options"`
	BasePrice   int                 `json:"base-price"`
}
