package main

import (
	"fmt"
)

func square(array1 [5]int32) int32 {
	var sum int32 = 0

	for index, value := range array1 {
		sum += value
		fmt.Println("at index ", index)

	}
	return sum
}

func main() {
	fmt.Println("loops")

	//this is about loops
	//iterate over mapps
	var myMap map[string]uint8 = map[string]uint8{"Arhant": 23, "Sarah": 45}
	myMap["jatin"] = 22

	for key, value := range myMap {
		fmt.Printf("Name: %v , Age: %v \n", key, value)
	}

	fmt.Println("Iterating over loops \n\n\n")
	//iterating over arrays and slices
	mySlice := []int32{2, 3, 4, 5, 6}

	for index, value := range mySlice {
		fmt.Printf("Index: %v , Value: %v\n", index, value)

	}

	//normal loops

	for i := 0; i < 10; i++ {
		fmt.Printf("%v \n", i)
	}

	//but we can also do
	//while --in-go-> for

	var j int32 = 0
	for j < 10 {
		fmt.Printf("%v \n", j)
		j = j + 1
	}

	// or we can
	var k int32 = 0
	for {
		if k > 10 {
			break
		}
		k++
		fmt.Printf("%v \n", k)
	}

	var array2 [5]int32 = [5]int32{1, 2, 3, 4, 5}

	result := square(array2)
	fmt.Println(result)
}
