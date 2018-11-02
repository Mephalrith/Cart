package main

import (
	"encoding/json"
	"flag"
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
	// Uses flags to specify json file paths. Requires both cart path and prices path. If either
	// is not present, will print help text and exit.
	cartPath := flag.String("cart", "", "path to cart.json file (required)")
	pricesPath := flag.String("prices", "", "path to base prices.json file (required)")
	help := flag.Bool("help", false, "list of commands")

	flag.Parse()

	if *help == false {
		if *cartPath == "" || *pricesPath == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		// Take paths and convert given json files to custom structs.
		parseInput(*cartPath, &cart)
		parseInput(*pricesPath, &basePrices)
	} else {
		flag.PrintDefaults()
	}

	fmt.Printf("\nThe cart total: %v\n", totalCartPrice())
}
