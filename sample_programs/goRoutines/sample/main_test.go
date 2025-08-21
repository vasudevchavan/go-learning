package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomeThing(t *testing.T) {
	stdOut := os.Stdout              // Store Current standard output
	r, w, _ := os.Pipe()             // Create a Pipe to read and write
	os.Stdout = w                    // Write ouput to the Pipe we have created
	var wg sync.WaitGroup            // Initialize the wait group
	wg.Add(1)                        // Add wait
	go printGreeting("Welcome", &wg) // start a go routine to print output.
	wg.Wait()                        // wait for go routine to completed
	_ = w.Close()                    // close the pipe

	result, _ := io.ReadAll(r) // read everything from the read end of the pipe
	output := string(result)   // convert byte slice to string
	os.Stdout = stdOut         // restore the standard output

	if !strings.Contains(output, "Welcome") {
		t.Errorf("Expected Output missing")
	}

}
