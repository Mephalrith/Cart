### redbubbleCart

Takes the path to a cart.json file, and the path to a basePrices.json file. For each item in the cart, it compares the item to the relevant product in the base prices, and calculates the total cost.

Example:

```
./redbubbleCart /path/to/cart/cart.json /path/to/prices/prices.json

The cart total: 9363

Process finished with exit code 0
```

Automated tests include the creation of dummy items and base prices to test item methods, as well as benchmark tests for calculating final price and comparing item options to base price options.

Tests require having Golang installed. It can be downloaded from [https://golang.org/](https://golang.org/).

To run tests:
```
go test -v redbubbleCart
go test -bench .
```