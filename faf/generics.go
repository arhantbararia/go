package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3}
	var float32slice = []float32{1, 2.1, 3}
	var emptyIntSlice = []int{}
	var emptyfloat32slice = []float32{}

	fmt.Println(SumintSlice(intSlice))
	fmt.Println(Sumfloat32Slice(float32slice))
	//1. as we can see here it is not very good to have different functions for different types for the same function

	//2. Hence Generics are used
	//the sumSlice() function acts as a generic template for handling different types
	fmt.Println("By Generics")
	fmt.Println(sumSlice[float32](float32slice))
	fmt.Println(sumSlice[int](intSlice))

	//3. 'any' type in generics
	fmt.Println(isEmpty[int](emptyIntSlice))         // OR isEmpty(emptyIntSlice)
	fmt.Println(isEmpty[float32](emptyfloat32slice)) // OR isEmpty(emptyfloat32slice)

}

// go does not support function overloading
func SumintSlice(intSlice []int) int {
	var sum int
	for _, v := range intSlice {
		sum += v
	}
	return sum
}



func Sumfloat32Slice(float32slice []float32) float32 {
	var sum float32
	for _, v := range float32slice {
		sum += v
	}
	return sum
}

// 2. using generics
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// 3. Using 'any'
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
