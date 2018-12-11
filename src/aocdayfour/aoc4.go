package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening the file", err)
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	return scanner, nil
}

type Event struct {
	DT    time.Time
	event string
}

const (
	WakeUp     = "wakes up"
	Asleep     = "falls asleep"
	Guard      = "Guard"
	NonProcess = "WrongEvent"
)

type Events []Event

func (es *Events) processLine(s string) {
	words := strings.Split(s, " ")
	layout := "2006-01-02 15:04"
	d := words[0][1:]
	t := words[1][:len(words[1])-1]
	event := strings.Join(words[2:], " ")
	dt, _ := time.Parse(layout, d+" "+t)
	//	fmt.Println(dt, d, t, event)

	*es = append(*es, Event{dt, event})
}

func (es Events) Len() int {
	return len(es)
}

func (es Events) Swap(i, j int) {
	(es)[j], (es)[i] = (es)[i], (es)[j]
}

func (es Events) Less(i, j int) bool {

	if (es)[j].DT.After((es)[i].DT) {
		return true
	}
	return false
}

func getID(s string) string {
	return strings.Split(s, " ")[1]
}

func getEventType(e Event) (string, int) {
	t := e.DT.Minute()
	if strings.HasPrefix(e.event, Guard) {
		return Guard, t
	}
	if strings.HasPrefix(e.event, Asleep) {
		return Asleep, t
	}
	if strings.HasPrefix(e.event, WakeUp) {
		return WakeUp, t
	}

	return NonProcess, 0
}

func findMaxMinuteAmount(m map[int]int) (int, int) {
	MAX := 0
	MAXkey := 0
	for k, v := range m {
		if v > MAX {
			MAX = v
			MAXkey = k
		}
	}
	return MAXkey, MAX
}

func findSum(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func generateGuardMinutesMap(es Events) map[string]map[int]int {
	var guardID string
	AsleepMap := make(map[string]map[int]int)
	for i := range es {
		kind, minute := getEventType(es[i])
		if kind == Guard {
			guardID = getID(es[i].event)
		}
		// init map
		if AsleepMap[guardID] == nil {
			AsleepMap[guardID] = make(map[int]int)
		}
		if kind == Asleep && i < len(es) {
			_, nextminute := getEventType(es[i+1])
			i++
			for jj := minute; jj < nextminute; jj++ {
				AsleepMap[guardID][jj]++
			}

		}
	}
	return AsleepMap
}

func main() {
	var es Events
	scanner, _ := getScanner("input.txt")

	for scanner.Scan() {
		line := scanner.Text()
		es.processLine(line)
	}

	// I preferred sort.Sort at the end over sorting during the scan
	// just for getting used to implementing an interface.
	sort.Sort(es)

	AsleepMap := generateGuardMinutesMap(es)

	PartOne(AsleepMap)
	PartTwo(AsleepMap)

}

func PartOne(AsleepMap map[string]map[int]int) {
	var Max int
	var FinalGuard string
	Max = 0
	for guard := range AsleepMap {
		sum := findSum(AsleepMap[guard])
		if sum > Max {
			Max = sum
			FinalGuard = guard
		}
	}

	Max = 0
	FinalMinute := 0
	for k, v := range AsleepMap[FinalGuard] {
		if v > Max {
			Max = v
			FinalMinute = k
		}
	}
	fmt.Printf("part one, Guard id %v, Minute %v \n", FinalGuard, FinalMinute)
}

func PartTwo(AsleepMap map[string]map[int]int) {
	var MaxMin int
	var FinalGuard string
	MaxMin = 0
	var MaxValue int
	MaxValue = 0
	for guard := range AsleepMap {
		minute, amount := findMaxMinuteAmount(AsleepMap[guard])
		if amount > MaxValue {
			MaxMin = minute
			FinalGuard = guard
			MaxValue = amount
		}
	}
	fmt.Printf("part two, Guard id %v, Minute %v \n", FinalGuard, MaxMin)

}
