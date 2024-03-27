package main

import (
	//"sync"
	"io"
	"os"
	"strings"
	"testing"
)

const newMessage = "New message"

func TestUpdateMassage(t *testing.T) {
	wg.Add(1)
	go updateMessage(newMessage)
	wg.Wait()

	if msg != newMessage {
		t.Errorf("updateMessage(): Expected 'New message', got %s", msg)
	}
}

func TestPrintMassage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = newMessage
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, newMessage) {
		t.Errorf("printMessage(): Expected 'New message', got %s", output)
	}
}

func TestMain(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("main(): Expected Hello, universe!, got %s", output)
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("main(): Expected Hello, cosmos!, got %s", output)
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("main(): Expected Hello, world!, got %s", output)
	}
}
