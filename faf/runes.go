package main

import (
	"fmt"
)

func main(){
	var myString = []rune("résumé")
	//runes are just aliases for the int32
	var indexed = myString[0]
	fmt.Printf("%v, %T\n ", indexed, indexed)

	for index, value := range myString {
		fmt.Println(index, value)
	}

}
