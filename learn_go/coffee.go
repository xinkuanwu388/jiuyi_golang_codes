package main

import "fmt"

func main() {

	var cups int
	fmt.Println("Write how many cups of coffee you will need:") // Writing a request message to the stdout
	_, err := fmt.Scan(&cups)
	if err != nil {
		panic(err)
	}

	//输出结果
	fmt.Printf("For %d cups of coffee you will need:\n", cups)
	fmt.Printf("%d ml of water\n%d ml of milk\n%d g of coffee beans\n", cups*200, cups*50, cups*15)

}
