package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var isGreetingDisplayed bool = true
var expenses = []map[string]interface{}{}

func main() {
	// Entry point
	var menuOption int

	if isGreetingDisplayed {
		fmt.Println("=== Hello 👋! Welcome to Expense Tracker. ===")
	}
	for {
		displayMenu()

		// program flow
		fmt.Print("Choose from the menu: ")
		fmt.Scanln(&menuOption)

		switch menuOption {
		case 1:
			clearTerminal()
			addExpense()
		case 2:
			viewAllExpenses()
		case 3:
			viewByCategory()
		case 6:
			exitApp()
			return
		}
	}
}

func displayMenu() {
	fmt.Println(`
MENU
=======================
1. Add Expense
2. View All Expenses
3. View by Category
4. Show Summary
5. Clear All Data
6. Exit
`)
}

func addExpense() {
	for {
		var amount float64
		var category, name, confirmation string

		fmt.Print("Enter expense name: ")
		if _, err := fmt.Scanln(&name); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Print("Enter expense amount: ")
		if _, err := fmt.Scanln(&amount); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if amount <= 0 {
			clearTerminal()
			fmt.Println("Amount must be greater than 0")
			continue
		}

		fmt.Print("Enter expense category: ")
		if _, err := fmt.Scanln(&category); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		expenses = append(expenses, map[string]interface{}{
			"category": strings.TrimSpace(category),
			"amount":   amount,
			"name":     strings.TrimSpace(name),
		})

		clearTerminal()
		fmt.Printf("✅ Successfully added %v.\n", strings.TrimSpace(name))

		for {
			fmt.Print("Do you want to add another one? (y/n): ")
			fmt.Scanln(&confirmation)
			var c = strings.TrimSpace(strings.ToLower(confirmation))

			if c == "y" {
				clearTerminal()
				break
			} else if c == "n" {
				clearTerminal()
				return
			} else {
				fmt.Println("Please enter 'y' for yes or 'n' for no.")
			}
		}
	}
}

func viewAllExpenses() {
	clearTerminal()

	if len(expenses) == 0 {
		fmt.Println("=== You have no expenses added. ===\n\n")
		return
	}

	var total float32

	fmt.Println("============= All Expenses =============\n")
	for i, e := range expenses {
		if amount, ok := e["amount"].(float32); ok {
			total += amount
		}
		fmt.Printf("%v. %v -- R%.2f [%v]\n", i+1, e["name"], e["amount"], e["category"])
	}

	fmt.Println("\n=========================================")
	fmt.Printf("Total: %v\n", total)
	fmt.Println("=========================================")
}

func viewByCategory() {
	clearTerminal()

	c := make(map[string][]map[string]interface{})

	for _, e := range expenses {
		ec := e["category"].(string)
		c[ec] = append(c[ec], e)
	}

	fmt.Println("=============== View by Category =================\n")
	if len(c) == 0 {
		fmt.Println("You have no expenses added.\n")
	}

	for cat, items := range c {
		fmt.Printf("--- %s ---\n", cat)
		ct := 0.0
		for _, e := range items {
			ct += e["amount"].(float64)
			fmt.Printf("• %v - R%.2f\n", e["name"], e["amount"])
		}
		fmt.Println("==================================================")
		fmt.Printf("Subtotal: R%.2f", ct)
		fmt.Println("\n==================================================")
	}
}

func viewSummary() {
	clearTerminal()
	fmt.Println("=== Expense Summary ===")
}

func exitApp() {
	var confirmation string
	farewellMsg := `
	=================================
	🧾  Expense Tracker
	=================================
	Your session has ended.
	Stay mindful of your spending 💸
	=================================
	`
	fmt.Print("You are about to leave the Expense Tracker. Are you sure you want to leave? (y/n): ")
	if _, err := fmt.Scanln(&confirmation); err != nil {
		fmt.Println("Invalid menu option: %v", err)
		return
	}

	for {
		if strings.ToLower(strings.TrimSpace(confirmation)) == "y" {
			clearTerminal()
			fmt.Println(farewellMsg)
			return
		} else if strings.ToLower(strings.TrimSpace(confirmation)) == "n" {
			isGreetingDisplayed = false
			clearTerminal()
			main()
			break
		}
	}
}

func clearTerminal() {
	// clears the terminal
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
