package main

import (
	"fmt"
	"errors"
	"time"
)

func main(){
	//printValue := "Arhant"

	//way to check errors in a function\
	result , remainder , err := intDivision(4 ,6 )	
	// if(err != nil ) {
		// fmt.Println(err )
		// time.Sleep(3 * time.Second)
		// fmt.Println("Should you be even doing this?")
		
	// }else if remainder == 0 {
	// 	fmt.Printf("The result of the division is %v" , result)
	// }else {
	// 	fmt.Printf("The result of the division is %v and remainder is %v" , result , remainder)
	// }

	switch{
	case err != nil:
		fmt.Println(err )
		time.Sleep(3 * time.Second)
		fmt.Println("Should you be even doing this?")
	
	case remainder == 0:
		fmt.Printf("Result: %v" , remainder )
	
	default:
		fmt.Printf("Result: %v , Remainder %v" , result , remainder )
	}

	switch remainder{
	case 0:
		//some command
	case 1:
		//some command
	default:
		//some command
	}

}

func printMe(printValue string){
	fmt.Println(printValue)
}


func intDivision(numerator int , denominator int ) (int , int , error  ) { 
	
	//create a function with proper error handling

	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by zero")
		return 0, 0 ,err

	}

	var result int = numerator / denominator
	var remainder int = numerator  & denominator

	return result , remainder , nil
}

