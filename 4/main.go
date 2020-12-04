package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("failed to open")
		
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	var passports []map[string]string
	
	currentPassport := make(map[string]string)
	
	for scanner.Scan() {
		line := scanner.Text()
		feats := strings.Split(line, " ")
		for _, feat := range feats {
			featMap := strings.Split(feat, ":")
			if len(featMap) == 2 {
				currentPassport[featMap[0]] = featMap[1]
			}
		}
		if line == "" {
			// New Pass
			passports = append(passports, currentPassport)
			currentPassport = make(map[string]string)
		}
	}
	// Append the last one
	passports = append(passports, currentPassport)
	_ = file.Close()
	
	ans2 := 0
	for _, pass := range passports {
		
		// byr (Birth Year)
		// four digits; at least 1920 and at most 2002.
		byrValid := false
		byr, hasByr := pass["byr"]
		if hasByr && len(byr) == 4 {
			year, err := strconv.Atoi(byr)
			if err == nil && (year >= 1920 && year <= 2002) {
				byrValid = true
			}
		}
		if !byrValid {
			continue
		}
		
		// iyr (Issue Year)
		// four digits; at least 2010 and at most 2020.
		iyrYearValid := false
		iyr, hasIyr := pass["iyr"]
		if hasIyr && len(iyr) == 4 {
			year, err := strconv.Atoi(iyr)
			if err == nil && (year >= 2010 && year <= 2020) {
				iyrYearValid = true
			}
		}
		if !iyrYearValid {
			continue
		}
		
		// eyr (Expiration Year)
		// four digits; at least 2020 and at most 2030.
		expYearValid := false
		eyr, hasEyr := pass["eyr"]
		if hasEyr && len(eyr) == 4 {
			year, err := strconv.Atoi(eyr)
			if err == nil && (year >= 2020 && year <= 2030) {
				expYearValid = true
			}
		}
		if !expYearValid {
			continue
		}
		
		//(Height) - a number followed by either cm or in:
		//If cm, the number must be at least 150 and at most 193.
		//If in, the number must be at least 59 and at most 76.
		heightValid := false
		hgt, hasHgt := pass["hgt"]
		if hasHgt &&
			(strings.HasSuffix(hgt, "cm") ||
				strings.HasSuffix(hgt, "in")) {
			if strings.HasSuffix(hgt, "cm") {
				hgt = strings.TrimSuffix(hgt, "cm")
				height, err := strconv.Atoi(hgt)
				if err != nil || (height >= 150 && height <= 193) {
					heightValid = true
				}
				
			} else if strings.HasSuffix(hgt, "in") {
				hgt = strings.TrimSuffix(hgt, "in")
				height, err := strconv.Atoi(hgt)
				if err != nil || (height >= 59 && height <= 76) {
					heightValid = true
				}
			}
		}
		if !heightValid {
			continue
		}
		
		//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		hclValid := false
		hcl, hasHcl := pass["hcl"]
		if hasHcl && (len(hcl) == 7) {
			//^#[a-f0-9_-]*$
			macthed, err := regexp.MatchString("^#[a-f0-9_-]*$", hcl)
			if err == nil && macthed {
				hclValid = true
			}
		}
		if !hclValid {
			continue
		}
		
		//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		eclValid := false
		ecl, hasEcl := pass["ecl"]
		if hasEcl {
			if (ecl == "amb") ||
				(ecl == "blu") ||
				(ecl == "brn") ||
				(ecl == "gry") ||
				(ecl == "grn") ||
				(ecl == "hzl") ||
				(ecl == "oth") {
				eclValid = true
			}
		}
		if !eclValid {
			continue
		}
		
		//pid (Passport ID) - a nine-digit number, including leading zeroes.
		pidValid := false
		pid, hasPid := pass["pid"]
		if hasPid && len(pid) == 9 {
			_, err := strconv.Atoi(pid)
			if err == nil {
				pidValid = true
			}
		}
		if !pidValid {
			continue
		}
		ans2++
	}
	fmt.Println(ans2)
	
}
