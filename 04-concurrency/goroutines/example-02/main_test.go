package main

import "testing"

func Test_UpdateMessage(t *testing.T) {
	msg = "welcome"
	wg.Add(1)
	updateMessage("hey")
	wg.Wait()

	if msg != "hey" {
		t.Errorf("Missing required message")
	}
}
