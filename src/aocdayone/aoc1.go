package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Frequency struct {
	counter int
}

func (freq *Frequency) updateFreq(f int) {
	freq.counter = freq.counter + f
}

func processLine(line []byte) (int, error) {
	if len(line) < 1 {
		return 0, errors.New("empty line")
	}
	if string(line[0]) == "+" {
		freq, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			fmt.Println("error parsing number in line", err)
			return 0, err
		}
		return freq, nil
	}
	if string(line[0]) == "-" {
		freq, err := strconv.Atoi(string(line))
		if err != nil {
			fmt.Println("error parsing number in line", err)
			return 0, err
		}
		return freq, nil
	}

	return 0, errors.New("unhandled case?")
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

func partOne(filename string) int {

	scanner, err := getScanner(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Println("solving aoc day part 1")

	var freq Frequency
	freq.counter = 0

	for scanner.Scan() {
		f, err := processLine(scanner.Bytes())
		if err != nil {
			fmt.Println(err)
			break
		}
		freq.updateFreq(f)
	}
	fmt.Printf("Final frequency: %d \n", freq.counter)
	return freq.counter
}

func PartTwo(filename string) int {
	fmt.Println("solving aoc day part 2")
	var freq Frequency
	freq.counter = 0
	freqmap := make(map[int]int)
	for {
		scanner, err := getScanner(filename)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		freqmap[0] = 1
		for scanner.Scan() {
			f, err := processLine(scanner.Bytes())
			if err != nil {
				fmt.Println(err)
				break
			}
			freq.updateFreq(f)
			freqmap[freq.counter]++
			if freqmap[freq.counter] > 1 {
				fmt.Printf("part2 solution: %d \n", freq.counter)
				return freq.counter
			}
		}
	}
}

func main() {
	filename := "input.txt"
	partOne(filename)
	PartTwo(filename)

}
