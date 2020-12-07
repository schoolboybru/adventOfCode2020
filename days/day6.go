package days

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFileDay6() {

	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n\n")

	countNumberOfYesAnswers(sliceData)
}

func countNumberOfYesAnswers(answers []string) {

	var counter = 0
	var counterPart2 = 0

	for _,val := range answers {
		counter += countYesInEntry(val)
		counterPart2 += countYesInEntryPart2(val)

	}

	fmt.Println(counter)
	fmt.Println(counterPart2)
}

func countYesInEntry(entry string) int {

	m := make(map[string]int)

	for _,val := range entry {
		var currentValue = string(val)
		if currentValue != "\n" {
			m[currentValue] += 1
		}
	}

	return len(m)
}

func countYesInEntryPart2(entry string) int {

	var organizedEntry = strings.Split(entry, "\n")

	m := make(map[string]int)

	var size = len(organizedEntry)
	var counter = 0

	for _,val := range entry {
		var currentValue = string(val)
		if currentValue != "\n" {
			m[currentValue] = 0
		}
	}

	for k := range m {
		for _,val := range organizedEntry {
			if strings.Contains(val, k) {
				m[k]++
			}
		}
	}

	for _,val := range m {
		if val == size {
			counter++
		}
	}

	return counter
}
