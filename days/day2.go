package days

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFileDay2() {
	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	validatePasswords(sliceData)

}

func validatePasswords(data []string) {
	var counter1 = 0
	var counter2 = 0

	for _,item := range data {
		key, input, nums := organizeData(item)

		if isValidPassword1(key, input, nums) {
			counter1++
		}

		if isValidPassword2(key, input, nums) {
			counter2++
		}
	}

	fmt.Println(counter1)
	fmt.Println(counter2)

}

func organizeData(data string) (string, string, []int){

	items := strings.Split(data, " ")
	var nums []int

	numbers := regexp.MustCompile(`[\d]+`)
	num := numbers.FindAllString(items[0], -1)

	for _, i := range num {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		nums = append(nums, j)
	}

	key := strings.Split(items[1], ":")

	return key[0], items[2], nums

}

func isValidPassword1(key string, input string, validRange []int) bool {
	var counter = 0
	var low = validRange[0]
	var high = validRange[1]

	for _,i := range input {
		if string(i) == key {
			counter++
		}
	}

	if counter >= low && counter <= high {
		return true
	}

	return false

}

func isValidPassword2(key string, input string, positions []int) bool {
	var position1 = positions[0] - 1
	var position2 = positions[1] - 1

	return (string(input[position1]) == key && string(input[position2]) != key) || (string(input[position2]) == key && string(input[position1]) != key)

}