package main

import (
	"fmt"
)


func square(thing2 []int32 ) (result []int32) {
	fmt.Println(&thing2)
	for i := range thing2{
		
	}
}

func main() {
	var p *int32 = new(int32) //nil by defualt, new is return an address after allocation
	//when you initialize a pointer with a memory allocation (malloc in C) or new it zeroes out that location
	fmt.Println(p)
	*p = 19
	fmt.Println(*p)

	*p = 10
	k := &p
	fmt.Println(k)
	

	fmt.Println(*k)

	//when dealing with arrays and slices
	//slicenew = sliceold --> this is always copy by reference not value
	var myslice1 = []int32{1,2,3}
	slice2 := myslice1
	slice2[2] = 4
	fmt.Println(myslice1)
	fmt.Println(slice2)

	



}
