package days

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFileDay3() {
	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	fmt.Println(findTrees(sliceData))

	var slope1 = findTreesPart2(sliceData, 1, 1)
	var slope2 = findTreesPart2(sliceData, 3, 1)
	var slope3 = findTreesPart2(sliceData, 5, 1)
	var slope4 = findTreesPart2(sliceData, 7, 1)
	var slope5 = findTreesPart2(sliceData, 1, 2)

	fmt.Println(slope1 * slope2 * slope3 * slope4 * slope5)

}

func findTrees(data []string) int{
	var currentPosition = 0
	var counter = 0
	var size = len(data[0])

	for idx,val := range data {
		currentPosition = (3*idx) % size
		if string(val[currentPosition]) == "#" {
			counter++
		}
	}

	return counter
}

func findTreesPart2(data []string, right int, down int) int {
	var currentPosition = 0
	var counter = 0
	var size = len(data[0])

	for i := 0; i < len(data); i = i + down {
		currentValue := data[i]
		if string(currentValue[currentPosition]) == "#" {
			counter++
		}
		currentPosition = (currentPosition + right) % size;
	}

	return counter
}