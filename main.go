package main

import (
	"fmt"
)

func main() {
	receipt := generationReceipt("John Doe")
	fmt.Println(receipt.format())
}
