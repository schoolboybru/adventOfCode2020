package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const target = 2020

func main() {
	numbers := readFile()

	part1(numbers, target)

	part2(numbers, target)

}

func part2 (numbers []int, target int) {
	m := make(map[int]int)

	for _, num := range numbers {
		m[num] = num
		for value := range m {
			if val, found := m[target - num - value]; found {fmt.Println(val * num * value)}
		}
	}
}

func part1 (numbers []int, target int) {
	m := make(map[int]int)
	for _, num := range numbers {
		var newValue = target - num
		if val, found := m[target - num]; found { fmt.Println(val * num)}
		m[num] = target - newValue
	}
}

func readFile() []int {
	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	var numbers []int

	for _,i := range sliceData {
		j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, j)
	}

	return numbers
}