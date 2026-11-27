package main

import (
	"fmt"
	"os"
	"strconv"
)

type Expense struct {
	amount   float64
	category string
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: add <amount> <category>")
	}

	command := os.Args

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

	switch command[1] {
	case "add":

		fmt.Printf("added : %.2f %s", amount, command[3])

	case "transport":

		fmt.Printf("added : %.2f %s", amount, command[3])

	case "list":
		for i := range os.Args {
			fmt.Println(i, ".", os.Args[2], os.Args[3])
		}
	case "total":
		fmt.Printf("total : %.2f", amount)
	}
}
