package main

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
