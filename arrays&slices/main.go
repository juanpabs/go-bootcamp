package main

import (
	"fmt"
)

func main() {

	//Arrays
	anArray := [4]int{1, 2, 4, -4}
	var threeD [2][2]int
	threeD = [2][2]int{{1, 2}, {3, 4}}
	any := [...]int{2, 3, 5, 6}

	fmt.Println("The length of", anArray, " is", len(anArray)) //len() to know the length of an array or slice
	fmt.Println("The length of ", threeD, "is", len(threeD))
	fmt.Println("The length of ", any, "is", len(any))

	//Slices
	aSliceLiteral := []int{1, 2, 3, 4, 5} //no length defined
	// aSliceLiteral = nil
	fmt.Println("The length of ", aSliceLiteral, " is", len(aSliceLiteral))
	fmt.Println("The capacity of ", aSliceLiteral, " is", cap(aSliceLiteral))
	integerSlice := make([]int, 20, 50) //* initial values
	fmt.Println("The length of ", integerSlice, " is", len(integerSlice))
	fmt.Println("The capacity of ", integerSlice, " is", cap(integerSlice))

	//append values
	integerSlice = append(integerSlice, 1)
	fmt.Println("integer Slice with value appended ", integerSlice)

	//re-slicing
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	reSlice[0] = -100
	reSlice[1] = 123
	fmt.Println(s1) //* values
	fmt.Println(reSlice)

	//Capacity varies
	aSlice := []int{-1, 0, 1}
	fmt.Println("aSlice ", aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -10)
	fmt.Println("aSlice ", aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -20)
	aSlice = append(aSlice, -30)
	aSlice = append(aSlice, -40)
	fmt.Println("aSlice ", aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice)) //* capacity

	//Copy slices
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6: ", a6)
	fmt.Println("a4: ", a4)

	copy(a6, a4)
	fmt.Println("a6: ", a6)
	fmt.Println("a4: ", a4)
	fmt.Println()

	a6 = []int{-10, 1, 2, 3, 4, 5}
	a4 = []int{-1, -2, -3, -4}
	copy(a4, a6)
	fmt.Println("a6: ", a6)
	fmt.Println("a4: ", a4)
	fmt.Println()

	//copy on an array
	array4 := [4]int{1, 2, 3, 4}
	s3 := []int{5, 5, 5}
	fmt.Println("array4: ", array4)
	fmt.Println("s3: ", s3)
	elementsCopied := copy(s3, array4[:])
	fmt.Println("array4: ", array4)
	fmt.Println("s3: ", s3)
	fmt.Println(elementsCopied)
	s3[0] = 100
	fmt.Println("array4: ", array4)
	fmt.Println("s3: ", s3)
	array4[0] = 200

	//question reference array
	array := [3]int{1, 2, 3}
	refArray := array[:]

	fmt.Println(array)
	fmt.Println(refArray)
	array[1] = 100
	fmt.Println(refArray)

	// ... to append
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 6}
	ref := a[:]
	fmt.Println("slice: ", s)
	fmt.Println("Existing array: ", ref)
	s = append(s, ref...)
	fmt.Println("New slice: ", s)
}
