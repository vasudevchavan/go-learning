package main

import (
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(path)
}
