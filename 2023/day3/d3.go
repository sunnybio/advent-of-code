package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(fileName string) bufio.Scanner {

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)

	}

	fileScanner := bufio.NewScanner(f)
	return *fileScanner
}
func main() {
	s := readFile("input.txt")
	for s.Scan() {
		fmt.Println(s.Text())
		var line string = s.Text()

		for i := 0; i < len(line); i++ {
			fmt.Println(line[i])
		}
	}
}
