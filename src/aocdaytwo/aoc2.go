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

func getDoubleTripleFlags(chars []byte) Flags {
	flags := Flags{false, false}

	// counts the occurrence of chars
	charmap := make(map[byte]int)

	// stores the doubles
	doubles := make(map[byte]int)

	//stores the triples
	triples := make(map[byte]int)

	for _, char := range chars {
		charmap[char]++
		value := charmap[char]
		if value == 2 {
			doubles[char] = 1
		}
		if value == 3 {
			delete(doubles, char)
			triples[char] = 1
		}
		if value > 3 {
			delete(triples, char)
		}
	}

	//
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
		fmt.Println("error retriving the scanner", err)
		return 0
	}

	fc := FlagsCounter{0, 0}
	for scanner.Scan() {
		flags := getDoubleTripleFlags(scanner.Bytes())
		fc.updateCounters(flags)
	}
	checksum := fc.Double * fc.Triple
	fmt.Printf("part one solution: checksum: %d", checksum)
	return checksum
}

// findEquivalentStrings finds equivalent strings in time n*k
// n = len of keys
// k = len of keys[i] i=0,1,..,n
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
	filename := "input.txt"
	partOne(filename)
	partTwo(filename)
}
