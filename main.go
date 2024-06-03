package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		p, err := strconv.ParseFloat(itemPrice, 64)
		if err != nil {
			fmt.Println("Invalid price")
			propmtChoices(receipt)
		}
		receipt.addItem(itemName, p)
		fmt.Println("Item added -", itemName, itemPrice)
		propmtChoices(receipt)

	case "t":
		tip, _ := getUserInput("Enter tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Invalid tip amount")
			propmtChoices(receipt)
		}
		receipt.updateTip(t)
		fmt.Println("tip added - ", tip)
		propmtChoices(receipt)

	case "q", "s":
		receipt.saveReceipt()
		fmt.Println("saving receipt\n", receipt.format())

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
