package main

import (
	"fmt"
	"os"
	"strconv"
)

type Expense struct {
	Amount   float64
	Category string
}

func (e Expense) Format() string {
	return fmt.Sprintf("%.2f %s", e.Amount, e.Category)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: add <amount> <category>")
		return
	}
	command := os.Args
	switch command[1] {
	case "add":
		if len(os.Args) < 4 { // 1. ไม่ครบ → บอกวิธีใช้ แล้วหยุด
			fmt.Println("usage: add <amount> <category>")
			return
		}
		amount, err := strconv.ParseFloat(command[2], 64)
		if err != nil {
			fmt.Println("Invalid must amount :", command[2])
			return
		}
		if amount <= 0 { // 3. ตัวเลขต้องมากกว่า 0
			fmt.Println("amount must be positive")
			return
		}

		e := Expense{Amount: 50, Category: "food"}
		fmt.Println("added:", e.Format())

	case "list":
		fmt.Println("lits : (empty)")
	default:
		fmt.Println("unknow command:", command[1])
	}
}
