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

	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
}
