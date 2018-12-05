package main

import (
	"testing"
)

func TestProcessLine(t *testing.T) {
	testline := []byte("+1")

	change, err := processLine(testline)
	expectedChange := 1
	if err != nil {
		t.Errorf("Processline returned an unexpected error, got %v, want %v", err, nil)
	}
	if change != expectedChange {
		t.Errorf("Error processing line, got %v, want %v", change, expectedChange)
	}

	testline = []byte("-1231")
	change, _ = processLine(testline)
	expectedChange = -1231
	if change != -1231 {
		t.Errorf("Error processing line, got %v, want %v", change, expectedChange)
	}

}

func TestPartTwo(t *testing.T) {
	sol := PartTwo("testinputparttwo.txt")
	if sol != 2 {
		t.Errorf("Error, wrong solution, got %v, want %v", sol, 2)
	}
}
