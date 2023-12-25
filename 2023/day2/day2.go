package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reading file for input
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(f)

	colorBook := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	fmt.Println(colorBook)
	var sum int = 0
	for s.Scan() {
		var gameData []string = strings.Split(s.Text(), ":")

		var gameNumber string = gameData[0]
		var gameIterationList []string = strings.Split(gameData[1], ";")

		var flag bool = true
		for _, gameIteration := range gameIterationList {
			for _, diceData := range strings.Split(gameIteration, ",") {

				var finalSplit []string = strings.Split(strings.Trim(diceData, " "), " ")
				fmt.Println(finalSplit)
				number, err := strconv.Atoi(finalSplit[0])
				if err != nil {
					panic(err)
				}
				var color string = finalSplit[1]

				if number > colorBook[color] {
					flag = false
				}
				fmt.Println(diceData)
			}

		}
		if flag {
			num, err := strconv.Atoi(strings.Split(gameNumber, " ")[1])
			if err != nil {
				panic(err)
			}

			sum = sum + num
		}
		fmt.Println(gameNumber)
	}
	fmt.Println(sum)

}
