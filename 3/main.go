package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open")
		
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	var slopes [][]string
	
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		slopes = append(slopes, line)
	}
	
	_ = file.Close()
	
	//Right 1, down 1
	santasPostion1 := 0
	treeCount1 := 0
	
	//Right 3, down 1.
	santasPostion2 := 0
	treeCount2 := 0
	
	// Right 5, down 1.
	santasPostion3 := 0
	treeCount3 := 0
	
	// Right 5, down 1.
	santasPostion4 := 0
	treeCount4 := 0
	
	//Right 7, down 1.
	santasPostion5 := 0
	treeCount5 := 0
	
	sizeOfEverySlope := len(slopes[0])
	
	getSantasNextPos := func(santasPostion int, nextStepSize int) int {
		predNextPos := santasPostion + nextStepSize
		
		if predNextPos < sizeOfEverySlope {
			return predNextPos
		}
		
		return predNextPos - sizeOfEverySlope
	}
	
	isTree := func(slope []string, santasPostion int, right int, treeCount int) (int, int) {
		place := slope[santasPostion]
		santasPostion = getSantasNextPos(santasPostion, right)
		if place == "#" {
			treeCount++
		}
		return treeCount, santasPostion
	}
	
	down := 0
	for _, slope := range slopes {
		treeCount1, santasPostion1 = isTree(slope, santasPostion1, 1, treeCount1)
		treeCount2, santasPostion2 = isTree(slope, santasPostion2, 3, treeCount2)
		treeCount3, santasPostion3 = isTree(slope, santasPostion3, 5, treeCount3)
		treeCount4, santasPostion4 = isTree(slope, santasPostion4, 7, treeCount4)
		if down%2 == 0 {
			treeCount5, santasPostion5 = isTree(slope, santasPostion5, 1, treeCount5)
		}
		
		down++
	}
	
	fmt.Println(treeCount1 * treeCount2 * treeCount3 * treeCount4 * treeCount5)
}
