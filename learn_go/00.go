package main

import "fmt"

func main() {
	var number int
	var response = "Zero!"
	fmt.Scanf("%d", &number)
	switch {
	case number < 0:
		response = "Negative!"
	case number > 0:
		response = "Positive!"
	}
	fmt.Println(response)
}
