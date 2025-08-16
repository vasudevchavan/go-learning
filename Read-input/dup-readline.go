package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	nscan := bufio.NewScanner(os.Stdin)
	nscan.Scan()
	born_year, err := strconv.ParseInt(nscan.Text(), 10, 64)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("you age is %d", 2023-born_year)

}
