package main

type Receipt struct {
	name  string
	items map[string]float64
	price float64
}

// generationReceipt creates a new receipt with the given name
func generationReceipt(name string) Receipt {
	receipt := Receipt{
		name:  name,
		items: make(map[string]float64),
		price: 0,
	}
	return receipt
}
