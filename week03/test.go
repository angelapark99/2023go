package main

import (
		"fmt"
		"math")

func main(){
	var price int = 100
	fmt.Println("price is", price, "dollars")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate //100.0 * 0.08
	fmt.Println("tax is", tax, "dollars")

	var total float64 = float64(price) + tax //100.0 + (100.0 * 0.08)
	fmt.Println("total cost is", total, "dollars")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("within budget", total <= float64(availableFunds))
}

