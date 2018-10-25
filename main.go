package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

var cart []Item
var basePrices []Price

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func readInput(input string, jsonType interface{}) {
	input = strings.TrimSpace(input)

	val, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(val, &jsonType)
}

func checkOptions(item Item, price Price) bool {
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

func calculatePrice(item Item, price Price) int {
	markup := float64(price.BasePrice) * (float64(item.ArtistMarkup) / 100)
	return (price.BasePrice + int(Round(markup))) * item.Quantity
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Cart path.json:")
	cartInput, _ := reader.ReadString('\n')
	readInput(cartInput, &cart)
	//readInput("one.json", &cart)

	fmt.Println("Enter Base Prices path.json:")
	priceInput, _ := reader.ReadString('\n')
	readInput(priceInput, &basePrices)
	//readInput("prices.json", &basePrices)

	cartTotal := 0

	for _, item := range cart {
		for _, price := range basePrices {

			if price.ProductType == item.ProductType {
				if checkOptions(item, price) {
					cartTotal += calculatePrice(item, price)
				}
			}
		}
	}

	fmt.Printf("\nThe cart total: %v\n", cartTotal)
}
