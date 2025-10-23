package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "hello world"
	// fmt.Println(name)

	fmt.Println("Index \t", "Char \t", "Byte Value")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%d \t %c \t %x \n", i, name[i], name[i])
	}

	fmt.Println("Index \t", "Rune \t", "Byte Value")
	for i, r := range name {
		fmt.Println(i, "\t", r, "\t", string(r))
	}

	str1 := "Hello"
	str2 := "world"
	addStrg := str1 + str2
	fmt.Println(addStrg)

	addStr1 := fmt.Sprintf("%s%s", str1, str2)
	fmt.Println(addStr1)

	var sb strings.Builder
	sb.WriteString(str1)
	sb.WriteString(str2)
	fmt.Println(sb.String())

}
