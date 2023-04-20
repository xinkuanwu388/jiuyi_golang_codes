package main

import "fmt"

func main() {

	var water, milk, grams, cups, available int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee grams the coffee machine has:")
	fmt.Scan(&grams)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)

	//if water/200 < milk/50 {
	//	canMake = water / 200
	//} else {
	//	canMake = milk / 50
	//}
	//if canMake > grams/15 {
	//	canMake = grams / 15
	//}

	availableIngredients := []int{water / 200, milk / 50, grams / 15}
	available = min(availableIngredients)

	switch {
	case available > cups:
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)", available-cups)
	case available < cups:
		fmt.Printf("No, I can make only %d cups of coffee", available)
	default:
		fmt.Printf("Yes, I can make that amount of coffee")
	}
}
func min(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < result {
			result = arr[i]
		}
	}
	return result
}
