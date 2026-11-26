package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	amount, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("invalid amount: ", err)
		return
	}
	fmt.Print("amount : ", amount, " Err : ", err)
}
