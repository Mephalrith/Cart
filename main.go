package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var cart []Item
var basePrices []Price
var pricesMap map[string]map[interface{}]int

func init() {
	pricesMap = make(map[string]map[interface{}]int)
}

func makeBasePricesMap() {
	fmt.Println("Making Base Price Map")
	var options interface{}

	for _, item := range basePrices {
		fmt.Println(item)
		if _, ok := pricesMap[item.ProductType]; !ok {
			pricesMap[item.ProductType] = make(map[interface{}]int)
		}
		var optionsList []string
		for _, vals := range item.Options {
			for _, val := range vals {
				optionsList = append(optionsList, val)
			}
		}
		for _, vals := range optionsList {
			options = append(options, vals)
		}
		pricesMap[item.ProductType][options] = item.BasePrice
		fmt.Println(pricesMap[item.ProductType])
		fmt.Println()
	}

	fmt.Println("Done Making Base Price Map")
}

func readInput(input string, jsonType interface{}) {
	input = strings.TrimSpace(input)

	val, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
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

func Round(x float64) float64 {
	t := math.Trunc(x)

	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}

	return t
}

func calculatePrice(item Item, price Price) int {
	markup := float64(price.BasePrice) * (float64(item.ArtistMarkup) / 100)
	return (price.BasePrice + int(Round(markup))) * item.Quantity
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//
	//fmt.Println("Enter Cart path.json:")
	//cartInput, err := reader.ReadString('\n')
	//if err != nil {
	//	panic(err)
	//}
	//readInput(cartInput, &cart)
	readInput("json/cart2.json", &cart)

	//fmt.Println("Enter Base Prices path.json:")
	//priceInput, err := reader.ReadString('\n')
	//if err != nil {
	//	panic(err)
	//}
	//readInput(priceInput, &basePrices)
	readInput("json/prices.json", &basePrices)

	makeBasePricesMap()
	fmt.Printf("Prices Map: %v\n", pricesMap)

	cartTotal := 0

	for _, item := range cart {
		fmt.Printf("Item: %v\n", item)
		if price, ok := pricesMap[item.ProductType]; ok {
			fmt.Printf("Price: %v\n", price)
			//if checkOptions(item, price) {
			//cartTotal += calculatePrice(item, price)
			//}
		}
	}

	fmt.Printf("\nThe cart total: %v\n", cartTotal)
}
