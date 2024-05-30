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

	// fmt.Println("Enter the consumer name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getUserInput("Enter the consumer name: ", reader)

	receiptInprogress := generationReceipt(name)

	fmt.Println("generating receipt for", receiptInprogress.name)
	return receiptInprogress
}

func propmtChoices(receipt Receipt) {
	reader := bufio.NewReader(os.Stdin)

	userChoice, _ := getUserInput("Select from bwlow, \n a to Add item \n t to add tip \n q/s to save receipt \n", reader)

	switch userChoice {
	case "a":
		itemName, _ := getUserInput("Enter item name: ", reader)
		itemPrice, _ := getUserInput("Enter item price: ", reader)

		fmt.Print(itemName, itemPrice)

	case "t":
		tip, _ := getUserInput("Enter tip amount: ", reader)
		fmt.Print(tip)

	case "q", "s":
		fmt.Println("saving receipt")

	default:
		fmt.Println("Invalid choice")
		propmtChoices(receipt)
	}
}

func main() {

	receipt := creatingReceipt()

	propmtChoices(receipt)

	// fmt.Println(receipt.format())

	// receipt := generationReceipt("John Doe")
	// receipt.addItem("Stationary", 5.09)
	// receipt.updateTip(0.50)
	// fmt.Println(receipt.format())
}
