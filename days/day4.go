package days

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var keys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

func ReadFileDay4() {
	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n\n")

	numberOfValidPassports, numberOfValidPassports2 := validateData(sliceData)

	fmt.Println(numberOfValidPassports)
	fmt.Println(numberOfValidPassports2)
}

func validateData(data []string) (int, int){
	var counter = 0
	var counter2 = 0
	for _,val := range data {
		var currentValue = organizeDataDay4(val)
		if isValidPassport(currentValue) {
			counter++
		}
		if isValldPassportPart2(currentValue) {
			counter2++
		}
	}

	return counter, counter2
}

func isValidPassport(passport map[string]string) bool{
	_, found := passport["cid"]
	if !found {
		if len(passport) == 7 {
			return true
		}
	}
	if len(passport) == 8 {
		return true
	}
	return false
}

func isValldPassportPart2(passport map[string]string) bool {
	var validPassportValues = 0
	for _,val := range keys {
		switch val {
		case "byr":
			var currentValue = passport[val]
			 i,err := strconv.Atoi(currentValue)
			 if err != nil {
			 	return false
			 }
			 if i >= 1920 && i <= 2002 {
			 	validPassportValues++
			 }
		case "iyr":
			var currentValue = passport[val]
			i,err := strconv.Atoi(currentValue)
			if err != nil {
				return false
			}
			if i >= 2010 && i <= 2020 {
				validPassportValues++
			}
		case "eyr":
			var currentValue = passport[val]
			i,err := strconv.Atoi(currentValue)
			if err != nil {
				return false
			}
			if i >= 2020 && i <= 2030 {
				validPassportValues++
			}
		case "hgt":
			var currentValue = passport[val]
			if matched,_ := regexp.MatchString("^(1[5-8][0-9]|19[0-3])(cm)$|^(59|6[0-9]|7[0-6])(in)$", currentValue); matched{validPassportValues++}
		case "hcl":
			var currentValue = passport[val]
			if matched,_ := regexp.MatchString("^#([0-9]|[a-f]){6}", currentValue); matched{validPassportValues++}
		case "ecl":
			var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			var currentValue = passport[val]

			for _,val := range eyeColors {
				if currentValue == val {
					validPassportValues++
					break
				}
			}
		case "pid":
			var currentValue = passport[val]
			if matched,_ := regexp.MatchString("^([0-9]{9}$)", currentValue); matched{validPassportValues++}
		case "cid":
			validPassportValues++
		}
	}

	if validPassportValues == 8 {
		return true
	}

	return false

}

func organizeDataDay4(passportField string) map[string]string{
	var regexRemoveWhiteSpace = regexp.MustCompile("[\\s]")

	var orderedValues = regexRemoveWhiteSpace.ReplaceAllString(passportField, "\n")
	values := strings.Split(orderedValues, "\n")

	m := make(map[string]string)

	for _,val := range values {
		currentValue := strings.Split(val, ":")
		m[currentValue[0]] = currentValue[1]
	}

	return m
}