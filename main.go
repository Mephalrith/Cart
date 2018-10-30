package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var cart []Item
var basePrices []Price
var pricesMap map[string]interface{}

func init() {
	pricesMap = make(map[string]interface{})
}

// Takes in a string leading to the path for a .json file, and the interface we are unmarshaling the json to.
// Reads the file found at the path, or panics if the path was bad, then performs the unmarshal operation.
func readInput(input string, jsonType interface{}) {
	input = strings.TrimSpace(input)

	val, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(val, &jsonType)
}

// Takes in a price of type struct.
// Instantiate the map for each product type, if it isn't already. Get all options available to the given price.
// Create a series of nested maps containing all valid options for each possible item, and finally set the base price
// to relevant keys.
func addToPricesMap(price Price) {
	// The first time a specific product type comes up, we instantiate its nested map.
	if _, ok := pricesMap[price.ProductType]; !ok {
		// The item doesn't have any options, so we just set the base price and stop.
		if len(price.Options) == 0 {
			pricesMap[price.ProductType] = price.BasePrice
			return
		}
		pricesMap[price.ProductType] = make(map[string]interface{})
	}

	// We know that pricesMap[price.ProductType] is valid because we just instantiated it, so we panic if it isn't.
	currentMap, ok := pricesMap[price.ProductType].(map[string]interface{})
	if !ok {
		panic("Invalid Map")
	}

	validKeys := price.getOptionsKeys()
	maxDepth := len(price.Options) - 1

	// We loop through maps by their depth for each valid key. When we reach the end, set the base price.
	for depth := 0; depth <= maxDepth; depth++ {
		nextMap := make(map[string]interface{})

		for _, key := range validKeys {
			if val, ok := currentMap[key]; ok {
				fmt.Println(key, val)
			}
			if depth == maxDepth {
				currentMap[key] = price.BasePrice
			} else {
				currentMap[key] = nextMap
			}
		}
		currentMap = nextMap
	}
}

// Takes in an item of type struct.
// Checks if each key from the given item exists in the base prices map.
// Returns the base price of the item as an int, or 0 if it invalid.
func getBasePrice(item Item) int {
	nextMap := pricesMap
	keys := item.getAllKeys()

	for _, key := range keys {
		val := nextMap[key]

		if priceToReturn, ok := val.(int); ok {
			return priceToReturn
		}

		if nestedMap, ok := val.(map[string]interface{}); ok {
			nextMap = nestedMap
		}
	}
	fmt.Println("Weirdness Ensued")
	return 0
}

func main() {
	//// Read in two command line arguments. First argument for cart path, second for prices path.
	//// If two command line arguments are not provided, panics with a helpful message.
	//if len(os.Args) != 2 {
	//	panic("Paths to two json files required.\nExample: ./redbubbleCart path/to/cart.json path/to/prices.json")
	//}
	//cartArg := os.Args[1]
	//priceArg := os.Args[2]
	//
	//// Take paths and convert given json files to custom structs.
	//readInput(cartArg, &cart)
	//readInput(priceArg, &basePrices)

	readInput("json/cart2.json", &cart)
	readInput("json/prices.json", &basePrices)

	// Create a cache of base prices.
	for _, price := range basePrices {
		addToPricesMap(price)
	}

	fmt.Println(pricesMap)

	cartTotal := 0
	// Do the calculation for each item in the cart.
	for _, item := range cart {
		cartTotal += item.calculatePrice(getBasePrice(item))
	}
	fmt.Printf("\nThe cart total: %v\n", cartTotal)
}
