package main

import "fmt"

type cupOfCoffee struct {
	water, milk, beans int
}

type coffeeMachine struct {
	water, milk, beans int
}

func countCups(c cupOfCoffee, m coffeeMachine) (cups int) {
	for cups = 0; m.water >= c.water && m.milk >= c.milk && m.beans >= c.beans; cups++ {
		m.water -= c.water
		m.milk -= c.milk
		m.beans -= c.beans
	}
	return
}

func main() {
	oneCup := cupOfCoffee{200, 50, 15} // (water, milk, beans)

	var saeco coffeeMachine
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&saeco.water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&saeco.milk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&saeco.beans)

	var cups int
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)

	cupsAvail := countCups(oneCup, saeco)

	switch {
	case cupsAvail == cups:
		fmt.Println("Yes, I can make that amount of coffee")
	case cupsAvail > cups:
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", cupsAvail-cups)
	case cupsAvail < cups:
		fmt.Printf("No, I can make only %d cups of coffee\n", cupsAvail)
	}
}
