package main

import (
	"fmt"
	"strings"
	"time"

)


func direct_add(strSlice []string  ) time.Duration {
	var catStr = ""
	var t0 = time.Now()
	

	for _, v := range strSlice{
		catStr += v
	}
	time_taken := time.Since(t0)
	return time_taken

}

func stringBuild(strSlice []string) time.Duration {
	
	var strBuilder strings.Builder
	var t0 = time.Now()
	for _, v  := range strSlice {
		strBuilder.WriteString(v)

	}
	var fast_cat = strBuilder.String()
	// fmt.Println(fast_cat)
	if fast_cat != "" {
		fmt.Println("ok")
	}

	var time_taken = time.Since(t0)
	return time_taken
	

	
}

func main(){
	var t0 = time.Now()
	millionAs := make([]string, 100000)
    for i := range millionAs {
        millionAs[i] = "A"
    }
	fmt.Println(time.Since(t0))


	fmt.Printf("Total time with direct add %v \n" , direct_add(millionAs))
	fmt.Printf("Total time with stringbuilder: %v \n" , stringBuild(millionAs))
	
}