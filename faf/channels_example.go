package main

import (
	"fmt"
	"math/rand"
	"time"

)



var MAX_CHICKEN_PRICE float32 = 10
var MAX_TOFU_PRICE flaot32 = 10




func main(){
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com" , "costco.com" , "wholefoods.com"}

	for i := range websites {
		go checkChickenPrice(websites[i] , chickenChannel)
		go checkTofuPrice(websites[i]  , tofuChannel)

	}

	
}



func sendMessage(chiChan chan string , tofuChan chan string){
	switch{
	case website := <-chiChan:
		fmt.Printf("\n Text sent: Found deal on chicken at %v", website)
	case website := <-totofuChan:
		fmt.Printf("\n Email sent: Found deal on tofu at %v", website)

	}
}

func checkChickenPrice(website string , channel chan string){
	for{
		time.Sleep(time.Second + 1)
		var chickenPrice = rand.Float32()*20
		if(chickenPrice <= MAX_CHICKEN_PRICE){
			channel <- website
			break

		}
	}
}

func checkTofuPrice(website string , channel chan string){
	for{
		time.Sleep(time.Second + 1)
		var tofuPrice = rand.Float32() * 20
		if(tofuPrice <= MAX_TOFU_PRICE ){
			channel <- website
			break
		}
		
	}
}