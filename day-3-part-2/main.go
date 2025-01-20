package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	instructions := readRawInstructionStringsFromFile()
	factorsRegex := regexp.MustCompile(`\d+`)
	sumOfResults := 0
	instructionsAreEnabled := true
	for _, instruction := range instructions {
		if instruction == "do()" {
			instructionsAreEnabled = true
			continue
		}
		if instruction == "don't()" {
			instructionsAreEnabled = false
			continue
		}
		if instructionsAreEnabled {
			sumOfResults += evaluateInstruction(instruction, factorsRegex)
		}
	}
	fmt.Println(sumOfResults)
}

func evaluateInstruction(instruction string, factorsRegex *regexp.Regexp) int {
	factors := factorsRegex.FindAllString(instruction, 2)
	if len(factors) != 2 {
		log.Fatalf("Did not find exactly 2 factors in instruction %s", instruction)
	}
	return parseFactor(factors[0]) * parseFactor(factors[1])
}

func parseFactor(s string) int {
	factor, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Ran into error parsing factor: %v", err)
	}
	return factor
}

func readRawInstructionStringsFromFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	defer file.Close()

	instructionStrings := []string{}

	instructionRegex := regexp.MustCompile(`(mul\(\d+,\d+\))|(do(n't)?\(\))`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allInstructionsInLine := instructionRegex.FindAllString(scanner.Text(), -1)
		instructionStrings = append(instructionStrings, allInstructionsInLine...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return instructionStrings
}
