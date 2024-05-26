package main

import "fmt"

type Receipt struct {
	name  string
	items map[string]float64
	price float64
}

// generationReceipt creates a new receipt with the given name
func generationReceipt(name string) Receipt {
	receipt := Receipt{
		name:  name,
		items: map[string]float64{"item1": 1.99, "item2": 5.99},
		price: 0,
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

	fs += fmt.Sprintf("%-25v ...$ %0.2f", "total:", total)
	return fs
}
