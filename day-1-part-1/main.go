package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ListsFromInput struct {
	leftList  []string
	rightList []string
}

func main() {
	listsFromInput := readInputLists()
	sortedNumbersInLeftList := toSortedIntList(listsFromInput.leftList)
	sortedNumbersInRightList := toSortedIntList(listsFromInput.rightList)
	if len(sortedNumbersInLeftList) != len(sortedNumbersInRightList) {
		log.Fatalf("Input lists have different lengths: %d and %d", len(sortedNumbersInLeftList), len(sortedNumbersInRightList))
	}
	sumOfDifferences := 0
	for i := 0; i < len(sortedNumbersInLeftList); i++ {
		difference := int(math.Abs(float64(sortedNumbersInLeftList[i] - sortedNumbersInRightList[i])))
		sumOfDifferences += difference
	}
	fmt.Println(sumOfDifferences)
}

func toSortedIntList(stringList []string) []int {
	intList := convertToIntList(stringList)
	slices.Sort(intList)
	return intList
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

func readInputLists() ListsFromInput {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	leftList := []string{}
	rightList := []string{}

	for scanner.Scan() {
		lineTokens := strings.Fields(scanner.Text())
		if len(lineTokens) != 2 {
			log.Fatalf("invalid line format: %s", scanner.Text())
		}
		leftList = append(leftList, lineTokens[0])
		rightList = append(rightList, lineTokens[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return ListsFromInput{
		leftList:  leftList,
		rightList: rightList,
	}
}
