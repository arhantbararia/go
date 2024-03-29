package main

import (
	"fmt"
)

func square1(thing2 *[5]int32) (result *[5]int32) {
	fmt.Printf("The memory location of the thing2 array is: %p  \n", thing2) //%v shows &[1 2 3 4 5] because it does not have the value
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
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
	var myslice1 = []int32{1, 2, 3}
	slice2 := myslice1
	slice2[2] = 4
	fmt.Println(myslice1)
	fmt.Println(slice2)

	var thing1 [5]int32 = [5]int32{1, 2, 3, 4, 5}
	var item1 *[5]int32 = &thing1

	fmt.Printf("The memory location of the thing1 array is: %p , %v \n", &thing1, &item1)
	//to print address we use & because in Printf when passing by value it does not retain original vars
	// address

	result := square1(&thing1)
	fmt.Println(result)
	fmt.Println(thing1) //not changed (pass by value)

}
