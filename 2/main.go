package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	
	type rule struct {
		minOcu       int
		maxOcu       int
		matchingChar string
		password     string
	}
	
	var rules []rule
	
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		matchingChar := strings.TrimRight(s[1], ":")
		ocu := strings.Split(s[0], "-")
		minOcu, _ := strconv.Atoi(ocu[0])
		maxOcu, _ := strconv.Atoi(ocu[1])
		r := rule{
			minOcu, maxOcu, matchingChar, s[2],
		}
		
		rules = append(rules, r)
	}
	
	file.Close()
	
	var inValidPassword int
	var validPassword int
	var validPasswordPos int
	for _, r := range rules {
		count := strings.Count(r.password, r.matchingChar)
		if count < r.minOcu || count > r.maxOcu {
			validPassword++
		} else {
			inValidPassword++
		}
		
		validationChar1 := r.password[r.minOcu-1]
		validationChar2 := r.password[r.maxOcu-1]
		
		if validationChar1 == r.matchingChar[0] || validationChar2 == r.matchingChar[0] {
			validPasswordPos++
		}
		
		if validationChar1 == r.matchingChar[0] && validationChar2 == r.matchingChar[0] {
			validPasswordPos--
		}
	}
	fmt.Println(inValidPassword)
	fmt.Println(validPasswordPos)
}
