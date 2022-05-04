package main

import (
	"fmt"
	"math"
)

// const name string = "Panchito"

// const (
// 	TODO   = iota //0
// 	DOING         //1
// 	REVIEW        //2
// 	QA            //3
// 	DONE          //4
// )

var global string = "For all the program" //global variagle with value and type defined

var GlobalExported string = "Accesible for all packages"

// naming: CammelCase instead of snake_case (newVar or NewVar instead of new_var)

func main() {
	var m, n float64   //variables only type defined
	m, n = 1.223, 3.44 //define the value of m and n
	// m = "lalita"
	fmt.Println("m, n: ", m, n)

	// zero values
	var (
		boolean bool
		message string
		decimal float32
		integer int
	)

	fmt.Printf("Zero values: \nbool: %v\nstring: %v\nnumber: %v\ninteger: %v\n",
		boolean, message, decimal, integer)

	newVar := false

	fmt.Printf("newVar value: %v type: %T\n", newVar, newVar)

	//int division
	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T \n", x)

	div := x / k
	fmt.Println("div: ", div)

	//float division
	y := 4 / 2.3
	fmt.Println("y: ", y)

	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat: ", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}

func complexNumbers() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T", c1)
	fmt.Printf("Type of c2: %T", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3: ", c3)
	fmt.Printf("Type of c3: %T", c3)

	cZero := c3 - c3
	fmt.Println("cZero: ", cZero)
}

func maxValues() {
	fmt.Println("MAX int8 value: ", math.MaxInt8)
	fmt.Println("MAX int16 value: ", math.MaxInt16)
	fmt.Println("MAX int32 value: ", math.MaxInt32)
	fmt.Println("MAX int64 value: ", math.MaxInt64)
	fmt.Println("MAX int value: ", math.MaxInt)
	fmt.Println("MAX uint8 value: ", math.MaxUint8)
	fmt.Println("MAX uint16 value: ", math.MaxUint16)
	fmt.Println("MAX uint32 value: ", math.MaxUint32)
	// fmt.Println("MAX uint64 value: ", math.MaxUint64)
	// fmt.Println("MAX uint value: ", math.MaxUint)

	fmt.Println("MAX uint / int relation: ", math.MaxUint/math.MaxInt)

	fmt.Println("MAX float32 value: ", math.MaxFloat32)
	fmt.Println("MAX float64 value: ", math.MaxFloat64)
}
