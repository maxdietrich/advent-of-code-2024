package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func main() {
	reports := readReports()
	sumOfSafeLevels := 0
	for _, report := range reports {
		if areLevelsSafe(report.levels) {
			sumOfSafeLevels++
		}
	}
	fmt.Println(sumOfSafeLevels)
}

func areLevelsSafe(levels []int) bool {
	startIsIncreasing := false

	for i := 0; i < len(levels); i++ {
		if i == 1 && levels[i] > levels[i-1] {
			startIsIncreasing = true
		}
		if i > 0 {
			if (startIsIncreasing && levels[i] < levels[i-1]) || (!startIsIncreasing && levels[i] > levels[i-1]) {
				return false
			}
			differenceToLast := math.Abs(float64(levels[i] - levels[i-1]))
			if differenceToLast < 1 || differenceToLast > 3 {
				return false
			}
		}
	}
	return true
}

func convertToIntList(stringList []string) []int {
	intList := make([]int, 0, len(stringList))
	for _, listEntry := range stringList {
		intValue, err := strconv.Atoi(listEntry)
		if err != nil {
			log.Fatalf("could not convert value %s to int: %v", listEntry, err)
		}
		intList = append(intList, intValue)
	}
	return intList
}

func readReports() []Report {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	defer file.Close()

	reports := []Report{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineTokens := strings.Fields(scanner.Text())
		levels := convertToIntList(lineTokens)
		reports = append(reports, Report{levels})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return reports
}
