package main

import (
	"testing"
)

var validItem = Item{"hoodie", map[string]string{"size": "small", "colour": "dark", "print-location": "front"}, 20, 1}
var invalidItem = Item{"hoodie", map[string]string{"size": "xs", "colour": "grey", "print-location": "back"}, 20, 1}
var newItem = Item{"pin", map[string]string{"size": "medium", "character": "link", "style": "shiny"}, 30, 5}

var validPrice = Price{"hoodie", map[string][]string{"colour": {"white", "dark"}, "size": {"small", "medium"}}, 3800}
var newPrice = Price{"pin", map[string][]string{"size": {"small", "medium", "large"}, "character": {"link", "zelda", "ganondorf", "triforce"}, "style": {"shiny", "matte", "glitter"}}, 1000}

func Test_options(t *testing.T) {
	// validItem has options present in validPrice, so it should return true.
	t.Run("valid hoodie", testCheckOptions(validItem, validPrice, true))

	// invalidItem has no options present in validPrice, so it should return false.
	t.Run("invalid hoodie", testCheckOptions(invalidItem, validPrice, false))

	// newItem has options present in newPrice, so it should return true.
	// This test is useful for seeing if the method handles options that are not present in my test prices.json file.
	t.Run("new pin", testCheckOptions(newItem, newPrice, true))
}

func Test_Calculations(t *testing.T) {
	// validItem has an artist markup of 20, and quantity of 1. Using validPrice for the base price, the expected total is 4560.
	t.Run("valid hoodie", testCalculate(validItem, validPrice, 4560))

	// newItem has an artist markup of 30, and a quantity of 5. Using newPrice for the base price, the expected total is 6500.
	t.Run("new pin", testCalculate(newItem, newPrice, 6500))
}

func testCheckOptions(item Item, price Price, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		found := item.checkOptions(price)
		if found != expected {
			t.Errorf("checkOptions(price) error: expected %v but got %v", expected, found)
		}
	}
}

func testCalculate(item Item, price Price, expected int) func(*testing.T) {
	return func(t *testing.T) {
		calculated := item.calculatePrice(price.BasePrice)
		if calculated != expected {
			t.Errorf("calculatePrice(price) error: expected %v but got %v", expected, calculated)
		}
	}
}

// 53.0 ns/op
func Benchmark_checkOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCheckOptions(validItem, validPrice, true)
	}
}

// 7.24 ns/op
func Benchmark_calculatePrice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validItem.calculatePrice(validPrice.BasePrice)
	}
}
