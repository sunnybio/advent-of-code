package utils

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
