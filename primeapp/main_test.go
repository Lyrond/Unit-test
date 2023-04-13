package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func TestCheckNumbers(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		done     bool
	}{
		{"7", "7 is a prime number!", false},
		{"-1", "Negative numbers are not prime, by definition!", false},
		{"0", "0 is not prime, by definition!", false},
		{"1", "1 is not prime, by definition!", false},
		{"q", "", true},
	}

	for _, tc := range testCases {
		scanner := bufio.NewScanner(strings.NewReader(tc.input + "\n"))
		result, done := checkNumbers(scanner)

		if result != tc.expected {
			t.Errorf("For input '%s', expected '%s', but got '%s'", tc.input, tc.expected, result)
		}

		if done != tc.done {
			t.Errorf("For input '%s', expected done=%t, but got done=%t", tc.input, tc.done, done)
		}
	}
}

func Test_intro(t *testing.T) {
	// set up a pipe to capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// call the intro function
	intro()

	// restore the original stdout
	w.Close()
	os.Stdout = old

	// read the captured output from the pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// check that the output is correct
	expectedOutput := "Is it Prime?\n" +
		"------------\n" +
		"Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n" +
		"-> "

	if buf.String() != expectedOutput {
		t.Errorf("Unexpected output from intro(). Got: %s, Expected: %s", buf.String(), expectedOutput)
	}
}

func Test_readUserInput(t *testing.T) {
	// prepare input to the function
	input := "5\nq\n"
	reader := strings.NewReader(input)

	// replace os.Stdout with a pipe to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// create a channel to indicate when the function is done
	doneChan := make(chan bool)

	// call the function
	go readUserInput(reader, doneChan)

	// wait for the function to finish
	<-doneChan

	// restore os.Stdout and read the captured output
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old

	// expected output
	expected := "Please enter a whole number!\n-> 7 is a prime number!\n-> Goodbye.\n"

	// compare the expected output with the actual output
	if buf.String() != expected {
		t.Errorf("readUserInput failed, expected %s but got %s", expected, buf.String())
	}
}
