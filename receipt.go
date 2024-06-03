package main

import (
	"fmt"
	"os"
)

type Receipt struct {
	name  string
	items map[string]float64
	price float64
	tip   float64
}

// generationReceipt creates a new receipt with the given name
func generationReceipt(name string) Receipt {
	receipt := Receipt{
		name:  name,
		items: map[string]float64{},
		price: 0,
		tip:   0,
	}
	return receipt
}

// format receipt as a string
func (receipt Receipt) format() string {
	fs := "Receipt details \n"
	var total float64 = 0
	for k, v := range receipt.items {
		fs += fmt.Sprintf("%-25v ...$ %v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintln("-------------------------------------")

	fs += fmt.Sprintf("%-25v ...$ %0.2f", "total:", total+receipt.tip)
	return fs
}

// update the tip of the receipt
func (receipt *Receipt) updateTip(tip float64) {
	receipt.tip = tip
}

// add an item to the receipt
func (receipt *Receipt) addItem(item string, price float64) {
	receipt.items[item] = price
}

// save receipt
func (receipt *Receipt) saveReceipt() {
	data := []byte(receipt.format())
	err := os.WriteFile("receipts/"+receipt.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Print("Receipt saved, Thank you!\n")
}
