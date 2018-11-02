package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var cart []Item
var basePrices []Price

// Takes in a string leading to the path for a .json file, and the interface we are unmarshaling the json to.
// Reads the file found at the path, or panics if the path was bad, then performs the unmarshal operation.
func parseInput(input string, jsonType interface{}) {
	input = strings.TrimSpace(input)

	val, err := ioutil.ReadFile(input)
	if err != nil {
		panic("Invalid path: " + input)
	}

	json.Unmarshal(val, &jsonType)
}

// Do the calculation for each item in the cart if the item is listed in base prices.
// Returns the total price of the cart.
func totalCartPrice() int {
	cartTotal := 0
	for _, item := range cart {
		for _, price := range basePrices {

			if price.ProductType == item.ProductType {
				if item.checkOptions(price) {
					cartTotal += item.calculatePrice(price.BasePrice)
				}
			}
		}
	}
	return cartTotal
}

func main() {
	// Read in two command line arguments. First argument for cart path, second for prices path.
	// If two command line arguments are not provided, panics with a helpful message.
	if len(os.Args) != 3 {
		panic("Paths to two json files required.\nRequired format: ./redbubbleCart path/to/cart.json path/to/prices.json")
	}
	cartArg := os.Args[1]
	priceArg := os.Args[2]

	// Take paths and convert given json files to custom structs.
	parseInput(cartArg, &cart)
	parseInput(priceArg, &basePrices)

	fmt.Printf("\nThe cart total: %v\n", totalCartPrice())
}
