package main

import (
	"fmt"
)



var False = false
var True = true


type gasengine struct{
	kmpl uint16
	litres uint16
	owner_name  owner //instead of owner_name = owner.name, d_o_p = owner.date_of_purchase
}




func (e gasengine) kmsleft() uint8 {
	return uint8(e.kmpl * e.litres)
}

type elecEngine struct {
	kmpkwh uint16
	kwh uint16
}


func (e elecEngine) kmsleft() uint8 {
	return uint8(e.kmpkwh * e.kwh)

}


// let's define an interface, or AKA FunctionAPI\
type engine interface{
	kmsleft() uint8 //defined a functonal endpoint
}


func canMakeit (e engine , kms uint8 ) (bool) {
	if kms <= e.kmsleft(){ //using the interface defined endpoint
		fmt.Println("You will make it")
		return true
	}else{
		fmt.Println("You won't make it")
		return false
	}
}




type owner struct{
	name string
	date_of_purchase string
}


func main(){
	var gEngine gasengine = gasengine{25 , 26 , owner{"test_name" , "02-08-00"}}
	gEngine.litres = 20
	fmt.Println(gEngine)
	fmt.Println(gEngine.owner_name.date_of_purchase)

	//can also create temporary struct inline definition
	var myEngine2 = struct{
		kmpl int32
		litres int32

	}{20, 23}
	//the above definition method is not reusable, only one instance is created this can be used to declare constant objects
	fmt.Println(myEngine2.kmpl)

	canMakeit(gEngine , 20)
	
}
