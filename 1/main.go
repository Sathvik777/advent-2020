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
	var text []string
	
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	
	file.Close()
	
	/**
	* The solution from this line
	**/
	
	recorder := make(map[int]int)
	
	for _, eachLn := range text {
		i, _ := strconv.Atoi(eachLn)
		diff := 2020 - i
		perDiff, ok := recorder[i]
		if ok {
			fmt.Println(perDiff * i)
			return
		}
		recorder[diff] = i
	}
}
