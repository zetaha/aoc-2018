package main

import (
	"testing"
)

type FlagTests struct {
	F Flags
	S string
}

func flagChecker(teststring []byte, e Flags, t *testing.T) {
	flags := processString(teststring)
	if flags.Double != e.Double {
		t.Errorf("Error flag double for test string %v , got %v, expected %v", string(teststring), flags.Double, e.Double)
	}
	if flags.Triple != e.Triple {
		t.Errorf("Error flag triple for test string %v, got %v, expected %v", string(teststring), flags.Triple, e.Triple)
	}
}

func TestProcessString(t *testing.T) {
	var tests []FlagTests
	tests = append(tests, FlagTests{Flags{false, false}, "abcdef"})
	tests = append(tests, FlagTests{Flags{true, true}, "bababc"})
	tests = append(tests, FlagTests{Flags{true, false}, "abbcde"})
	tests = append(tests, FlagTests{Flags{false, true}, "abcccd"})

	for _, test := range tests {
		flagChecker([]byte(test.S), test.F, t)
	}
}

func BenchmarkFindEquivalentStrings(b *testing.B) {
	scanner, _ := getScanner("input.txt")
	var ids []string
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	for i := 0; i < b.N; i++ {
		findEquivalentStrings(ids, 2)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partTwo("input.txt")
	}
}
