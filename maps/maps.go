package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70, // Must have trailing coma in multi line
	}

	// Number of items
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// Return 0 if value not found
	fmt.Println(stocks["TSLA"])

	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// Set a value
	stocks["TSLA"] = 822.12
	fmt.Println(stocks)

	// Delete item
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
