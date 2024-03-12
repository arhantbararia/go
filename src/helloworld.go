package main

import (
	"fmt"
	"math"
	"strconv"
)


func square_root(n float64 ) (x float64  , y string ) {
	if n < 0{
		x = 0
		y = strconv.FormatFloat(math.Sqrt( -n ) , 'f' , -1 , 64) +  "i"
		return
	}

	x = math.Sqrt(n)
	y = "0"
	
	
	return

}


func main(){
	fmt.Println(square_root(-34))
	return
}