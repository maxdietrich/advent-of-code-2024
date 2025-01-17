package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ListsFromInput struct {
	leftList  []string
	rightList []string
}

func main() {
	listsFromInput := readInputLists()
	leftList := convertToIntList(listsFromInput.leftList)
	rightList := convertToIntList(listsFromInput.rightList)
	appearancesInRightList := determineNumberOfAppearances(rightList)
	similarityScore := 0
	for _, leftEntry := range leftList {
		similarityScore += leftEntry * appearancesInRightList[leftEntry]
	}
	fmt.Println(similarityScore)
}

func determineNumberOfAppearances(list []int) map[int]int {
	appearanceMap := make(map[int]int)
	for _, entry := range list {
		appearanceMap[entry] = appearanceMap[entry] + 1
	}
	return appearanceMap
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
