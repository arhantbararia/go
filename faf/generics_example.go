package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ContactInfo struct {
	Name  string
	Email string
}

type PurchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}


func main(){
	var contactData := loadJSON[ContactInfo]("./contactInfo.json")
	var purchaseData := loadJSON[PurchaseInfo]("./purchaseInfo.json")
	

}


func loadJSON[T ContactInfo | PurchaseInfo](filepath string) ([]T) {
	data, _ := os.Open(filepath)
	defer data.Close()

	var loaded = []T{}

	decoder := json.NewDecoder(data)
	err := decoder.Decode(&loaded) //inplace storing that's why address is passed

	if err != nil {
		fmt.Println()
	}
	return loaded

}
