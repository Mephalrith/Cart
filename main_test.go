package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	readInput("json/cart2.json", &cart)
	readInput("json/prices.json", &basePrices)

	// Create a cache of base prices.
	for _, price := range basePrices {
		addToPricesMap(price)
	}

	cartTotal := 0
	// Do the calculation for each item in the cart.
	for _, item := range cart {
		cartTotal += item.calculatePrice(getBasePrice(item))
	}

	fmt.Printf("\nThe cart total: %v\n", cartTotal)
	os.Exit(m.Run())
	// Output:
	//
	// The cart total: 9363
}

func BenchmarkAddToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, price := range basePrices {
			addToPricesMap(price)
		}
	}
}

func BenchmarkCalculatePrice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cartTotal := 0
		for _, item := range cart {
			cartTotal += item.calculatePrice(getBasePrice(item))
		}
	}
}
