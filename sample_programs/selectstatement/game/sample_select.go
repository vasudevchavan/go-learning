package game

import (
	"fmt"
	"time"
)

var chan1 = make(chan string)
var chan2 = make(chan string)

func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "Task 1 completed"
}

func task2() {
	time.Sleep(1 * time.Second)
	chan2 <- "Task 2 completed"
}

func main() {

	go task1()
	go task2()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}

}
