package main


import (
	"fmt"
)

type GasEngine struct{
	gallons float32
	mpg float32
}

type ElectricEngine struct{
	kwh float32
	mpkwh float32

}


type car [ET GasEngine | ElectricEngine] struct{
	carMake string
	model string
	engineType ET
}

func main(){


	var gasCar = car[GasEngine]{
		carMake: "honda",
		model: "city"
		engineType: GasEngine{
			gallons: 12.4,
			mpg: 10.2,
		},
	}

	var elecCar = car[ElectricEngine]{
		carMake: "tesla"
		model: "model 3",
		engineType: ElectricEngine{
			kwh: 123,
			mpkwh: 123.5,
		},
	}


}