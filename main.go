package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var isGreetingDisplayed bool = true

func main() {
	// Entry point
	var menuOption int

	if isGreetingDisplayed {
		fmt.Println("=== Hello 👋! Welcome to Expense Tracker. ===")
	}
	displayMenu()

	// program flow
	fmt.Print("Choose from the menu: ")
	fmt.Scanln(&menuOption)

	switch menuOption {
	case 6:
		exitApp()
		break
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
