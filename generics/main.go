package main

import "fmt"

func main() {
	division := GenericDivide(3.3,2)
	fmt.Println("Division: ", division)

	compareStrings := isEqual("lala", "lala")
	fmt.Println("strings: ", compareStrings)

	compareNumbers := isEqual(3, 2)
	fmt.Println("numbers: ", compareNumbers)

	// var number interface{}
	// number = 2
	// fmt.Printf("number type: %T", number)
}

func GenericDivide[N, D int | float64](num N, den D) float64{
	fmt.Printf("type of numerator: %T\n", num)
	return float64(num)/float64(den)
}

func isEqual[E comparable](first, second E) bool{
	fmt.Printf("type of comparation: %T\n", first)
	return first == second
}