package main

import "fmt"

func main() {

	// //if else statement
	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// //if without an else
	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }

	// //A statement can precede conditionals
	// if num := 8342; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	// //Command arguments
	// arguments := os.Args
	// if len(arguments) < 2 {
	// 	fmt.Println("usage: switch number.")
	// 	os.Exit(1)
	// }

	// //switch
	// asString := arguments[1]
	// switch asString {
	// case "5":
	// 	fmt.Println("Five!")
	// case "0":
	// 	fmt.Println("Zero!")
	// default:
	// 	fmt.Println("Do not care!")
	// }

	// //convert to int, switch statement
	// number, err := strconv.Atoi(arguments[1])
	// if err != nil {
	// 	fmt.Println("This value is not an integer:", number)
	// 	os.Exit(1)
	// }

	// switch {
	// case number < 0:
	// 	fmt.Println("Less than zero!")
	// case number > 0:
	// 	fmt.Println("Bigger than zero!")
	// default:
	// 	fmt.Println("Zero!")
	// }

	// //regex with switch
	// var negative = regexp.MustCompile(`-`)
	// var floatingPoint = regexp.MustCompile(`\d?\.\d`)

	// switch {
	// case negative.MatchString(asString):
	// 	fmt.Println("Negative number")
	// case floatingPoint.MatchString(asString):
	// 	fmt.Println("Floating point!")
	// 	fallthrough
	// default:
	// 	fmt.Println("Something else!")
	// }

	// //guess type of an empty interface
	// var aType interface{}
	// aType = "hola"
	// switch aType.(type) {
	// case nil:
	// 	fmt.Println("It is nil interface!")
	// case bool:
	// 	fmt.Println("I'm a bool")
	// case int:
	// 	fmt.Println("I'm an int")
	// case string:
	// 	fmt.Println("I'm a string")
	// default:
	// 	fmt.Printf("Don't know type %T\n", aType)
	// }

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

	// "do while" loop
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
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
