package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "expenses.json"

type Expense struct {
	Amount   float64
	Category string
}

func (e *Expense) Format() string {
	return fmt.Sprintln(e.Amount, e.Category)
}

func loadExpense() ([]Expense, error) {
	raw, err := os.ReadFile("expenses.json")
	if errors.Is(err, os.ErrNotExist) {
		return []Expense{}, nil
	}
	if err != nil {
		return nil, err
	}
	var list []Expense
	err = json.Unmarshal(raw, &list)
	if err != nil {
		return nil, err
	}
	return list, err
}

func saveExpenses(list []Expense) error {
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: add <amount> <category>")
		return
	}

	command := os.Args

	switch command[1] {
	case "add":
		if len(command) < 4 {
			fmt.Println("usage: add <amount> <category>")
			return
		}

		amount, err := strconv.ParseFloat(command[2], 64)

		if err != nil {
			fmt.Println("Invalid amount :", command[2])
			return
		}

		if amount < 0 {
			fmt.Println("Invalid must be positive")
			return
		}
		e := Expense{Amount: amount, Category: command[3]}

		list, err := loadExpense()
		if err != nil {
			fmt.Println("laod error:", err)
			return
		}

		list = append(list, e)
		err = saveExpenses(list)
		if err != nil {
			fmt.Println("save error:", err)
			return
		}
		fmt.Println("added", e.Format())
		fmt.Printf("added : %.2f %s", amount, command[3])

	case "transport":
		if len(command) < 4 {
			fmt.Println("usage: add <amount> <category>")
			return
		}

		amount, err := strconv.ParseFloat(command[2], 64)

		if err != nil {
			fmt.Println("Invalid amount :", command[2])
			return
		}

		if amount < 0 {
			fmt.Println("Invalid must be positive")
			return
		}

		e := Expense{Amount: amount, Category: command[3]}
		list, err := loadExpense()
		if err != nil {
			fmt.Println("load error:", err)
			return
		}
		list = append(list, e)

		err = saveExpenses(list)
		if err != nil {
			fmt.Println("save error:", err)
			return
		}
		fmt.Println("added:", e.Format())

		fmt.Printf("added : %.2f %s", amount, command[3])

	case "list":
		list, err := loadExpense()
		if err != nil {
			fmt.Println("load error", err)
			return
		}
		if len(list) == 0 {
			fmt.Println("no expenses yet")
			return
		}
		for i, e := range list {
			fmt.Printf("%d. %s\n", i+1, e.Format())
		}
	case "total":
		list, err := loadExpense()
		if err != nil {
			fmt.Println("load error:", err)
			return
		}
		var sum float64
		for _, e := range list {
			sum += e.Amount
		}
		fmt.Printf("total : %.2f\n", sum)
	default:
		fmt.Println("unknow command")
	}
}
