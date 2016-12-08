package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type info struct {
	groupName string
	remainder int
	number    int
}

var digits = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var tens = map[int]string{
	1:  "ten",
	2:  "twenty",
	20: "twenty",
	3:  "thirty",
	4:  "fourty",
	5:  "fifty",
	6:  "sixty",
	7:  "seventy",
	8:  "eighty",
	9:  "ninety",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

func main() {
	remaining, _ := strconv.Atoi(os.Args[1])
	outStr := ""

	for remaining > 0 {
		currentInfo := getInfo(remaining)

		current := currentInfo.number

		if current >= 100 {
			h := current / 100
			current = current % 100
			outStr += fmt.Sprintf("%v hundred ", digits[h])
		}

		if current >= 10 {
			if current <= 20 {
				outStr += fmt.Sprintf("%v ", tens[current])
			} else {
				t := current / 10
				d := current % 10
				outStr += fmt.Sprintf("%v ", tens[t])
				outStr += fmt.Sprintf("%v ", digits[d])
			}
		}

		if currentInfo.groupName != "" {
			outStr += currentInfo.groupName
		}

		remaining = currentInfo.remainder
	}

	fmt.Println(outStr)
}

func floorDiv(num, dem int) int {
	return int(math.Floor(float64(num / dem)))
}

func getInfo(num int) *info {
	if num >= 1e9 {
		return &info{
			"billion ",
			num % 1e9,
			num / 1e9,
		}
	} else if num >= 1e6 {
		return &info{
			"million ",
			num % 1e6,
			num / 1e6,
		}
	} else if num >= 1e3 {
		return &info{
			"thousand ",
			num % 1e3,
			num / 1e3,
		}
	}
	return &info{
		"",
		0,
		num,
	}
}
