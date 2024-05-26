package main

import (
	"fmt"
)

func main() {
	receipt := generationReceipt("John Doe")

	receipt.addItem("Stationary", 5.09)
	receipt.updateTip(0.50)

	fmt.Println(receipt.format())
}
