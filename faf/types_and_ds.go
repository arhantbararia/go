package main

import (
	"fmt"
	"unicode/utf8" //path is given instead of .
)

func main() {
	var intNum int16 = 32767 //max value

	intNum = intNum + 1 //go will allow it but will show ambiguos results
	fmt.Println(intNum) //-32768

	var floatNum float32 = 12345678.9
	var floatNum64 float64 = 12345678.9 // will store precise results

	println(floatNum)
	println(floatNum64)

	//cannot add 2 different types without explicit type casting: float64(int32value)

	//defaults to floor function i.e. int(3/2) = 1
	var myString string = "Hello" + " " + "World"

	fmt.Println(len(myString)) // this is not actual length but instead its bytes but mostly each character in utf-8 encoding is 1 byte it doesn't matter
	fmt.Println(len("γ"))      // gamma symbol output: 2

	//to get actual length
	fmt.Println(utf8.RuneCountInString("γ")) //output: 1

	var myboolean bool = false
	fmt.Println(myboolean)

	myVar := "text" // walrus operator
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// myVariable := foo() //will work but not prefer
	// var Variable string = foo() //this is nice

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.134
	print(pi)

}


// uint8  : 0 to 255 
// uint16 : 0 to 65535 
// uint32 : 0 to 4294967295 
// uint64 : 0 to 18446744073709551615 
// int8   : -128 to 127 
// int16  : -32768 to 32767 
// int32  : -2147483648 to 2147483647 
// int64  : -9223372036854775808 to 9223372036854775807

