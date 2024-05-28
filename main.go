package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func creatingReceipt() Receipt {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the consumer name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	receiptInprogress := generationReceipt(name)

	fmt.Println("generating receipt for", receiptInprogress.name)
	return receiptInprogress
}

func main() {

	receipt := creatingReceipt()

	fmt.Println(receipt.format())

	// receipt := generationReceipt("John Doe")
	// receipt.addItem("Stationary", 5.09)
	// receipt.updateTip(0.50)
	// fmt.Println(receipt.format())
}
