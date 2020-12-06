package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Read input
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("failed to open")
		
	}
	
	type location struct {
		row    int
		column int
	}
	
	var passes []location
	seats := make(map[int]map[int]string)
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	lastRow := 127
	lastColumn := 7
	
	passLocation := func(indexer string) location {
		newFirstRow := 0
		newLastRow := lastRow
		
		newFirstColumn := 0
		newLastColumn := lastColumn
		
		cursor := strings.Split(indexer, "")
		for _, c := range cursor {
			
			if c == "F" || c == "B" {
				numberOfSeatsLeft := newLastRow - newFirstRow
				splitter := numberOfSeatsLeft / 2
				if c == "F" {
					newLastRow = (newLastRow - splitter) - 1
				} else if c == "B" {
					newFirstRow = (newFirstRow + splitter) + 1
				}
			}
			
			if c == "R" || c == "L" {
				numberOfSeatsLeft := newLastColumn - newFirstColumn
				splitter := numberOfSeatsLeft / 2
				// Should have found row by now
				// To be safe check
				if c == "L" {
					newLastColumn = (newLastColumn - splitter) - 1
				} else if c == "R" {
					newFirstColumn = (newFirstColumn + splitter) + 1
				}
			}
		}
		
		row, initRow := seats[newFirstRow]
		if !initRow {
			row = make(map[int]string)
		}
		row[newFirstColumn] = "X"
		seats[newFirstRow] = row
		
		result := location{
			newFirstRow, newFirstColumn,
		}
		return result
	}
	
	highestPassId := 0
	
	for scanner.Scan() {
		place := passLocation(scanner.Text())
		passes = append(passes, place)
		passId := (place.row * 8) + place.column
		if highestPassId < passId {
			highestPassId = passId
		}
		
	}
	_ = file.Close()
	fmt.Println(highestPassId)
	
	var rowsIds []int
	for k := range seats {
		rowsIds = append(rowsIds, k)
	}
	
	sort.Ints(rowsIds)
	
	for index, i := range rowsIds {
		row := seats[i]
		if index == 0 || index == len(rowsIds)-1 {
			continue
		}
		for j := 0; j <= lastColumn; j++ {
			_, takenSeat := row[j]
			if !takenSeat {
				passId := (i * 8) + j
				fmt.Println("pass id =", passId)
			}
		}
	}
	
}
