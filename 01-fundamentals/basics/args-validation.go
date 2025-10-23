package main

import (
	"fmt"
	"os"
)

func main() {
	arg_len := len(os.Args[1:])
	fmt.Println("Argument legth: ", arg_len)
}
