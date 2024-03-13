package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	fmt.Println(len(myString))
	var indexed = myString[0]
	fmt.Printf("%v, %T\n ", indexed, indexed)

	for index, value := range myString {
		fmt.Println(index, value)
	}
//output
	// 	114, uint8
	//  0 114
	// 1 233
	// 3 115
	// 4 117
	// 5 109
	// 6 233
	//since char é is 2 bytes long. it skips 2nd index. and when accessing these bytes it gets converted to uint8
	// In loop1:

    // - myString[0] corresponds to the first byte of the UTF-8 representation of the character 'résumé', which is 114 in decimal.
    // - When iterating over the string using a range loop, it correctly splits the string into individual Unicode code points,
	//  printing the index and the corresponding code point.
	// - The range function is doing the heavy lifting of breaking down 2byte character to its code point

	var indexed2 = myString[2]
	fmt.Printf("%v, %T\n ", indexed2, indexed2)

	intArr := [...]int32{0, 1, 2, 3, 4, 5, 6, 7}
	for i, v := range intArr {
		fmt.Println(i, byte(myString[v]))
	}
// output
	// 0 114
	// 1 195
	// 2 169
	// 3 115
	// 4 117
	// 5 109
	// 6 195
	// 7 169
	// In loop2:

    //- The intArr contains numbers 0 through 7.
    //- The indices 1 and 6 correspond to the values 1 and 6 respectively in the intArr.
    //- In Go, when you index a string with an integer, it interprets it as the index of the UTF-8 byte sequence of the string. 
	// So, when accessing myString[v] with v being 1 and 6, it's attempting to access bytes at those indices. 
	// These bytes represent parts of the UTF-8 encoding of the characters 'é' and 'é',
	//  which result in different values than the Unicode code points.
	

	//strings are immutable in go.
	// var string2 = "arihant"
	//string2[2] = 'a' // this will not work 
	
	//but this will work
	var strSlice = []string{"a" , "r" , "h" , "a" , "n" , "t"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i]

	}
	fmt.Println(catStr)

	//because this loop will create an entire new string every time, :/ inefficient
	//instead we can use stringbuilder from strings package

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])

	}

	var fast_cat = strBuilder.String()
	fmt.Println(fast_cat)

}



