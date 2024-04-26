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

func main() {
	contactData := loadJSON[ContactInfo]("contactInfo.json")
	purchaseData := loadJSON[PurchaseInfo]("purchaseInfo.json")

	fmt.Println(contactData)
	fmt.Println(purchaseData)

}

func loadJSON[T ContactInfo | PurchaseInfo](filepath string) []T {
	data, err := os.Open(filepath)
	if err != nil {

		fmt.Println("some error occurred during reading file")
		fmt.Println(err)
	}
	defer data.Close()

	var loaded = []T{}

	decoder := json.NewDecoder(data)
	err = decoder.Decode(&loaded) //inplace storing that's why address is passed

	if err != nil {

		fmt.Println("some error occurred during Decoding")
		fmt.Println(err)
	}
	return loaded

}
