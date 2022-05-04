package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 15 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()

	// "while" loop
	i := 0
	for {
		if i > 10 {
			break
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// iter an array
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
}
