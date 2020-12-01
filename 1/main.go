package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open")
		
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var numList []int
	
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numList = append(numList, number)
	}
	
	file.Close()
	
	/**
	* The solution from this line
	**/
	
	type pair struct {
		value1 int
		value2 int
	}
	sumsRecord := make(map[int]pair)
	
	for i, firstOfPair := range numList {
		
		var numList2 = numList[i+1:]
		for _, secondOfPair := range numList2 {
			sumOfPairs := firstOfPair + secondOfPair
			if (sumOfPairs) == 2020 {
				fmt.Println("ans1 = ", firstOfPair*secondOfPair)
			}
			if (sumOfPairs) < 2020 {
				sumsRecord[sumOfPairs] = pair{
					firstOfPair, secondOfPair,
				}
			}
		}
	}
	
	for _, num := range numList {
		diff := 2020 - num
		perDiff, ok := sumsRecord[diff]
		if ok {
			fmt.Println("ans2 = ", perDiff.value1*perDiff.value2*num)
			return
		}
	}
}
