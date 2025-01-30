package main

import (
	"fmt"
	"strconv" // Needed for int to string conversion
)

// Coffee struct to store coffee details
type Coffee struct {
	name  string
	price int
}

func main() {
	var sum int = 0
	coffeeList := map[int]Coffee{
		1: {name: "Hot Coffee", price: 200},
		2: {name: "Cool Coffee", price: 350},
		3: {name: "Black Coffee", price: 400},
	}

	// Map to store purchased coffee
	parcess := make(map[int]Coffee)

	// Welcome message
	fmt.Println("#------------------- Self Service Coffee Shop ------------------#")

	// User input loop
	for {
		var n int
		fmt.Println("\nEnter the value 0 to view items")
		fmt.Println("Enter the value -1 to exit")

		fmt.Scan(&n)
		if n == -1 {
			break
		} else if n == 0 {
			// Display coffee menu
			fmt.Println("\nAvailable Coffees:")
			for quantity, value := range coffeeList {
				fmt.Println(quantity, value.name, "=", value.price)
			}

			// User selects coffee
			var userInput int
			fmt.Println("#------- Buy Coffee --------#")
			fmt.Scan(&userInput)

			// Slice to store coffee names
			var arr []string

			// Check if selected coffee exists
			if coffee, exists := coffeeList[userInput]; exists {
				parcess[userInput] = coffee
				arr = append(arr, strconv.Itoa(userInput)) // Convert int to string before appending
				sum += coffee.price // Add price to total sum
			} else {
				fmt.Println("Invalid selection! Please choose a valid coffee.")
			}
		}
	}

	// Final receipt
	fmt.Println("\nItem Name        Price")
	fmt.Println("--------------------------------------")
	var totalAmount int
	for _, coffee := range parcess {
		fmt.Println(coffee.name, "      ", coffee.price)
		totalAmount += coffee.price
	}
	fmt.Println("--------------------------------------")
	fmt.Println("Total Amount:       ", totalAmount)
}
