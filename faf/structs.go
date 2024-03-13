package main

import (
	"fmt"

)

type gasengine struct{
	kmpl uint8
	litres uint8
	owner_name owner
}

//we can define a struct method
func (e gasengine) kmsleft() uint8 {
	return (e.kmpl * e.litres)
}

func canMakeit(e gasengine , kms uint8 )(bool) {
	if kms <= e.kmsleft(){
		fmt.Println("You will make it")
		return true
	}else{
		fmt.Println("You won't make it")
		return false
	}
}


type owner struct{
	name string
}

func main(){
	var myEngine gasengine = gasengine{25, 15 , owner{"arhant"}}
	myEngine.litres = 20
	fmt.Println(myEngine.kmpl , myEngine.litres)
	fmt.Println(myEngine.owner_name.name)

	//can also create temporary structs inline definition
	var myEngine2 = struct{
		kmpl int32
		litres int32
	}{20 , 29}
	//the above definition method is not reusable, only one instance is created this can be used to declare constant objects
	fmt.Println(myEngine2.kmpl)

	//calling struct method
	fmt.Println(myEngine.kmsleft())
		
}