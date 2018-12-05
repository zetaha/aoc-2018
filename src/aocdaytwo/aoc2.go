package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Flags struct {
	Double bool
	Triple bool
}

type FlagsCounter struct {
	Double int
	Triple int
}

type Seq struct {
	S1 string
	S2 string
}

func (fc *FlagsCounter) updateCounters(flags Flags) {
	if flags.Double {
		fc.Double++
	}
	if flags.Triple {
		fc.Triple++
	}
}
func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening the file", err)
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	return scanner, nil
}

func processString(chars []byte) Flags {
	flags := Flags{false, false}
	charmap := make(map[byte]int)
	doubles := make(map[byte]int)
	triples := make(map[byte]int)
	for _, char := range chars {
		charmap[char]++
		if charmap[char] == 2 {
			doubles[char] = 1
		}
		if charmap[char] == 3 {
			delete(doubles, char)
			triples[char] = 1
		}
		if charmap[char] > 3 {
			delete(triples, char)
		}
	}
	if len(doubles) > 0 {
		flags.Double = true
	}
	if len(triples) > 0 {
		flags.Triple = true
	}

	return flags
}

func partOne(filename string) int {
	fmt.Println("solving part one!")
	scanner, err := getScanner(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fc := FlagsCounter{0, 0}
	for scanner.Scan() {
		flags := processString(scanner.Bytes())
		fc.updateCounters(flags)
	}
	checksum := fc.Double * fc.Triple
	fmt.Printf("part one solution: checksum: %d", checksum)
	return checksum
}

func compareStrings(s1, s2 string) bool {
	this := []byte(s1)
	isEqual := false
	other := []byte(s2)
	var length int
	if len(this) < len(other) {
		length = len(this)
	} else {
		length = len(other)
	}
	diff := 0
	for i := 0; i < length; i++ {
		if this[i] != other[i] {
			diff++
		}
		if diff > 1 {
			isEqual = true
			break
		}
	}
	return isEqual
}

func assembleSequence(line string, skipindex int) string {
	seq := ""
	for j, char := range line {
		if j != skipindex {
			seq = seq + string(char)
		}
	}
	return seq
}

func findEquivalentStrings(keys []string, skipindex int) (Seq, error) {
	charmap := make(map[Seq]int)
	for k := range keys {
		seq := Seq{keys[k][:skipindex], keys[k][skipindex+1:]}
		charmap[seq]++
		if charmap[seq] >= 2 {
			return seq, nil
		}
	}
	return Seq{}, errors.New("could not find solution")
}

func partTwo(filename string) int {
	scanner, err := getScanner(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	var ids []string
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	for skip := 0; skip < len(ids[0]); skip++ {
		findEquivalentStrings(ids, skip)
	}

	return 0
}

func main() {
	partTwo("input.txt")
}
