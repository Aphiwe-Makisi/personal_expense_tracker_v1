# 🧾 Expense Tracker (Go CLI)

A simple command-line application written in Go that helps you track your daily expenses.
You can add expenses, view them by category, see summaries, and clear all data — all from the terminal.

### ✨ Features

- Add Expense → Record a new expense with name, amount, and category.

- View All Expenses → List all expenses with total spending.

- View by Category → Group expenses by category with subtotals.

- View Summary → Get overall statistics including:

  - Total expenses

  - Average per expense

  - Most expensive item

  - Spend per category

- Clear All Data → Permanently delete all recorded expenses.

- Exit → Leave the app with a friendly farewell message.

### 📋 Menu Options
### MENU
=======================
1. Add Expense
2. View All Expenses
3. View by Category
4. Show Summary
5. Clear All Data
6. Exit

### 🚀 How to Run

Make sure you have Go installed (go version).

Clone this repository or copy the source code into a file called main.go.

Open your terminal and run:

`go run main.go`

### 🖥️ Usage Example

=== Hello 👋! Welcome to Expense Tracker. ===

### MENU
=======================
1. Add Expense
2. View All Expenses
3. View by Category
4. Show Summary
5. Clear All Data
6. Exit

========================

Choose from the menu: 1

Enter expense name: Coffee

Enter expense amount: 20

Enter expense category: Food

✅ Successfully added Coffee.

### ⚠️ Notes

Data is stored in memory only. Once you exit the program, all expenses are lost.

Clearing all data is permanent — there’s no undo.

Works on Linux, macOS, and Windows (terminal clearing is OS-aware).

🛠️ Possible Improvements

- Save expenses to a file (CSV/JSON) for persistence.

- Import/export expense history.

- Support recurring expenses.

- Add sorting/filtering options.

- Build a TUI (Text-based UI) for better navigation.

✍️ Built with Go for learning and fun!