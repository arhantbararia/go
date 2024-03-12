package main

import (
	"fmt"
)

func main() {

	fmt.Println("arrays")
	var intArr [3]int32 //fixed length //all element same type  //indexable //contiguous in memory
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2]) // 0 ,1  [0 123] [m:n] m-> (n-1)

	fmt.Println(&intArr[0]) // 0xc00000a0d0					//indexable //contiguous in memory
	fmt.Println(&intArr[1]) //0xc00000a0d4
	fmt.Println(&intArr[2]) //0xc00000a0d8

	var intArr1 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr1)

	intArr2 := [...]int32{1, 2, 3, 4}
	fmt.Println(intArr2)

	//Dynamic arrays or vectors are called slices in go   //all element same type  //indexable

	fmt.Println("\n\n Slices")
	var intSlice []int32 = []int32{2, 3, 4, 5}
	fmt.Println(intSlice)
	fmt.Printf("The length of intSlice is %v with capacity %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The new length is %v and capacity is %v ", len(intSlice), cap(intSlice)) // capicity is generally gets doubled

	//any index couldn't get larger than length

	//can also append slices (concatenation) using spread operator (...)
	var intSlice2 = []int32{6, 7, 8, 0}
	intSlice = append(intSlice, intSlice2...) // append() only takes 2 parameters
	fmt.Println(intSlice)

	//can also make slices with specified length and capacity using a make function
	length := 3
	capacity := 8
	var intSlice3 []int32 = make([]int32, length, capacity)
	fmt.Println(intSlice3)

	//hashmaps ,maps ,dicitonary
	fmt.Println("\n\nMaps")

	var MyMap map[string]uint16 = make(map[string]uint16)

	MyMap["arhant"] = 32
	MyMap["Dost"] = 21
	fmt.Println(MyMap["nice"]) //wrong key will show null value of type (here int16 therefore 0)

	//when returning data from a map it also returns a list of elements (value , found(bool) )
	age, found := MyMap["arhant"]

	delete(MyMap , "Dost")
	fmt.Println(age)
	if found {
		fmt.Println(MyMap)

	}else{
		fmt.Println("data not found!")
	}

}
