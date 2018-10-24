package main

type Item struct {
	ProductType  string            `json:"product-type"`
	Options      map[string]string `json:"options"`
	ArtistMarkup int               `json:"artist-markup"`
	Quantity     int               `json:"quantity"`
}