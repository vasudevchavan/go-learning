package main

import (
	"fmt"
	"os"
	"log"
)




func main() {
	fileName := "testing.txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	// Once all operations are done this function will close the file.
  defer file.Close()
  fmt.Println("file created")
  
  // Write bytes to file
  wdata := []byte("hello \n")
  _,err = file.Write(wdata)
  	if err != nil {
		log.Printf("Error %s", err)
	}

  // Writing a slice of string
	inputData := []string{"we","are","in","india"}

  for _,v:= range inputData{
    v=v+"\n"
    _,err = file.WriteString(v)
  	if err != nil {
		log.Printf("Error %s", err)
	}
  }
}
