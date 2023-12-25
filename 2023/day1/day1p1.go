package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day1part1() {

	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	var sum int32 = 0
	for s.Scan() {
		var line string = s.Text()
		var j int = len(line) - 1
		var firstNum int8 = -1
		var secondNum int8 = -1
		for i := 0; i < len(line); i++ {

			if isInteger(line[i]) && firstNum == -1 {
				firstNum = int8(line[i])
			}

			if isInteger(line[j]) && secondNum == -1 {
				secondNum = int8(line[j])
			}

			//fmt.Println(string(line[i]), string(line[j]))
			j--
		}
		var comString string = string(firstNum) + string(secondNum)
		if combinedNum, err := strconv.Atoi(comString); err == nil {
			sum = sum + int32(combinedNum)
		}
	}
	fmt.Println(sum)
	f.Close()
}

func isInteger2(asciiVal byte) bool {

	if asciiVal > 47 && asciiVal < 58 {
		return true
	}
	return false
}
