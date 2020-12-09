package days

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadFileDay8() {
	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	organizeDataDay8(sliceData)
	Day8part2(sliceData)
}

func organizeDataDay8(instructions []string) {

	usedValues := make(map[int]bool)
	var accumulator = 0

	var idx = 0
	for idx < len(instructions) {
		var _, nextStep, addToAcc = computeInstruction(instructions[idx])
		if usedValues[idx] == true {
			break
		}
		usedValues[idx] = true
		idx += nextStep
		accumulator += addToAcc
	}

	fmt.Println(accumulator)

}

func Day8part2(instructions []string) {


	for idx,val := range instructions {
		var currInstruction, _, _ = computeInstruction(val)
		if currInstruction != "acc" {
			noLoop, accumulator := execute(currInstruction, idx, instructions)
			if noLoop {
				fmt.Println(accumulator)
				break
			}
		}
	}

}

func execute(currInstruction string, currIdx int, instructions []string) (bool, int) {
	usedValues := make(map[int]bool)
	var accumulator = 0

	var newInstruction = instructions[currIdx]

	if currInstruction == "nop" {
		newInstruction = strings.Replace(newInstruction, "nop", "jmp", 1)
	} else if currInstruction == "jmp" {
		newInstruction = strings.Replace(newInstruction, "jmp", "nop", 1)
	}

	var idx = 0

	for idx <= len(instructions) {
		if idx >= len(instructions) {
			return true, accumulator
		}
		var currIns = instructions[idx]
		if idx == currIdx {
			currIns = newInstruction
		}
		_, nextStep, addToAcc := computeInstruction(currIns)
		if usedValues[idx] == true {
			return false, accumulator
		}
		usedValues[idx] = true
		idx += nextStep
		accumulator += addToAcc
	}

	return true, accumulator

}

func computeInstruction(instruction string) (string, int, int) {
	var organizedInstruction = strings.Split(instruction, " ")

	var currInstruction = organizedInstruction[0]
	var nextStep = 0
	var addToAcc = 0
	currValue, err := strconv.Atoi(organizedInstruction[1])

	if err != nil {
		fmt.Println(err)
	}

	switch currInstruction {
	case "acc":
		addToAcc = currValue
		nextStep = 1

	case "jmp":
		addToAcc = 0
		nextStep = currValue

	case "nop":
		addToAcc = 0
		nextStep = 1

	}

	return currInstruction, nextStep, addToAcc

}