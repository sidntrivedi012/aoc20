package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/* 1. Read the file for input
2. Extract the input in an array
3. Iterate over the array, do the string operations and check each password by applying the required logic
4. Print the output. */

func main() {
	var inputData []int
	found := false
	file, err := os.Open("./inputs/one.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		check(err)
		inputData = append(inputData, input)
	}
	fmt.Println(inputData)
	var product int

	for i, value := range inputData {
		for j := i + 1; j < len(inputData); j++ {
			if value+inputData[j] == 2020 {
				product = value * inputData[j]
				found = true
				break
			}

		}
		if found == true {
			fmt.Println(product)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
