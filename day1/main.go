package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, i := range array {
		result += i
	}
	return result
}

func readInput(file string) *bufio.Scanner {
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	return scanner
}

func main() {
	filePath := "./day_1_input.txt"
	input := readInput(filePath)

	var lineNumbers []int

	re := regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)
	for input.Scan() {
		var numbers []int

		line := input.Text()
		n := re.FindAllString(line, -1)
		if len(n) < 1 {
			log.Fatal("No numbers in current line")
		}

		for _, i := range n {
			switch i {
			case "one":
				i = "1"
			case "two":
				i = "2"
			case "three":
				i = "3"
			case "four":
				i = "4"
			case "five":
				i = "5"
			case "six":
				i = "6"
			case "seven":
				i = "7"
			case "eight":
				i = "8"
			case "nine":
				i = "9"
			}
			num, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, num)

		}
		resultNumber, err := strconv.Atoi(fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1]))
		if err != nil {
			log.Fatal(err)
		}
		lineNumbers = append(lineNumbers, resultNumber)
	}
	fmt.Printf("Day1 answer: %d\n", sum(lineNumbers))
}
