package main

import (
	"fmt"
	"os/exec"
)

func main() {
	prg := "fio"
	arg := "-v"
	// arg:= os.Args[1]
	// prg1 := "ls"
	// arg2 := "-ltr"

	cmd := exec.Command(prg, arg)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))

	// cmd2 := exec.Command(prg1, arg2)
	// stdout1, err1 := cmd2.Output()

	// if err1 != nil {
	// 	fmt.Println(err1.Error())
	// 	return
	// }
	// fmt.Print(string(stdout1))
}
