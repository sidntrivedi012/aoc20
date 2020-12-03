package main

import (
	"bufio"
	"os"
)

func parseInput(filePath string) []int {
	var inputData []string
	found := false
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := scanner.Text()
		check(err)
		inputData = append(inputData, input)
	}
	fmt.Println("Parsed input - \n", inputData)
	return inputData
}
