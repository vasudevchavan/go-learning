package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_MainFunction(t *testing.T) {
	stdOut := os.Stdout  // Store Current standard output
	r, w, _ := os.Pipe() // Create a Pipe to read and write
	os.Stdout = w        // Write ouput to the Pipe we have created
	main()
	_ = w.Close() // close the pipe

	result, _ := io.ReadAll(r) // read everything from the read end of the pipe
	output := string(result)   // convert byte slice to string
	os.Stdout = stdOut         // restore the standard output

	if !strings.Contains(output, "130") {
		t.Errorf("Expected Output missing")
	}

}
