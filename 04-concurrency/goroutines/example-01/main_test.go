package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateGreeting("testing")
	wg.Wait()
	if msg != "testing" {
		t.Errorf("Expected output missing")
	}
}

func Test_PrintMessageTest(t *testing.T) {
	stdOut := os.Stdout  // Store Current standard output
	r, w, _ := os.Pipe() // Create a Pipe to read and write
	os.Stdout = w        // Write ouput to the Pipe we have created
	msg = "Welcome"
	printGreeting() // start a go routine to print output.
	wg.Wait()       // wait for go routine to completed
	_ = w.Close()   // close the pipe

	result, _ := io.ReadAll(r) // read everything from the read end of the pipe
	output := string(result)   // convert byte slice to string
	os.Stdout = stdOut         // restore the standard output

	if !strings.Contains(output, "Welcome") {
		t.Errorf("Expected Output missing")
	}

}

func Test_MainFunction(t *testing.T) {
	stdOut := os.Stdout  // Store Current standard output
	r, w, _ := os.Pipe() // Create a Pipe to read and write
	os.Stdout = w        // Write ouput to the Pipe we have created
	main()
	_ = w.Close()
	result, _ := io.ReadAll(r) // read everything from the read end of the pipe
	output := string(result)   // convert byte slice to string
	os.Stdout = stdOut         // restore the standard output

	if !strings.Contains(output, "Hello go") {
		t.Errorf("Expected Output missing")
	}
	if !strings.Contains(output, "Hello golang") {
		t.Errorf("Expected Output missing")
	}
	if !strings.Contains(output, "Hello goroutine") {
		t.Errorf("Expected Output missing")
	}

}
