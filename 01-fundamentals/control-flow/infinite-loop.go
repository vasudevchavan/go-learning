package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// ch := make(chan string)

	// go mylogger.ListenForLogs(ch)

	// fmt.Println("Capture User Input:")
	// fmt.Print("Enter text: ")

	// for i := 0; i < 1; i++ {
	// 	fmt.Println("->")
	// 	input, _ := reader.ReadString('\n')
	// 	ch <- input
	// 	time.Sleep(time.Second)
	// }

	tableMultiplier(3)

}

func tableMultiplier(a int) {
	for i := 1; i <= a; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}

	}
}
