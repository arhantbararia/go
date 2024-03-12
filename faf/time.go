package main

import (
	"fmt"
	"time"
)


func main(){
	var n int =  1000000
	var testSlice1 = []int{}
	var testSlice2 = make([]int , 0 , n)

	fmt.Printf("Total time without Preallocation: %v \n" , timeloop(testSlice1, n))
	fmt.Printf("Total time with Preallocation: %v \n" , timeloop(testSlice2 , n))
	

}


func timeloop(slice []int , n int) (time.Duration){
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice  , 1 )

	}

	return time.Since(t0)
}