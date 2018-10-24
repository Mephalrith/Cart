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
	/* TODO:
	 * i'm having trouble figuring out how to
	 * differentiate between the options.... things
	 * are evaluating to true when they shouldn't be
	 */
	isEqual := false
	for key, val := range item.Options {
		if j, ok := price.Options[key]; ok {
			for _, v := range j {
				if v == val {
					isEqual = true
				}
			}
		}
	}
	return isEqual
}

func calculatePrice(item Item, price Price) int {
	// TODO: markup is evaluating to 0, need to google more
	markup := float64(price.BasePrice * (item.ArtistMarkup / 100))
	fmt.Println(markup)

	total := (price.BasePrice + int(Round(markup))) * item.Quantity
	fmt.Println(total)

	return total

	// one.json should be:
	//((4368 * 0.3) + 4368) * 2
	//(1310.4 + 4368) * 2
	//5678.4 * 2
	//11356.8
}

func main() {
	//reader := bufio.NewReader(os.Stdin)

	//fmt.Println("Enter Cart path.json:")
	//cartInput, _ := reader.ReadString('\n')
	//readInput(cartInput, &cart)
	readInput("one.json", &cart)

	//fmt.Println("Enter Base Prices path.json:")
	//priceInput, _ := reader.ReadString('\n')
	//readInput(priceInput, &basePrices)
	readInput("prices.json", &basePrices)

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
