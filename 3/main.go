package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	fmt.Println("Day 3")
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	in, err := os.ReadFile(path.Join(thisDir, "input"))
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(in[:]), "\n")
	var total int
	fmt.Printf("Line count: %d\n", len(lines))

	start := time.Now()
	for _, line := range lines {
		debug := true
		mid := len(line) / 2
		left := line[:mid]
		right := line[mid:]
		if debug {
			fmt.Printf("left=%s right=%s\n", left, right)
		}
		var leftExists = make(map[uint8]bool)
		var rightExists = make(map[uint8]bool)
		for i, _ := range make([]int, mid-1) {
			rightChar := right[i]
			leftChar := left[i]
			leftExists[leftChar] = true
			rightExists[rightChar] = true
		}

		for k, _ := range leftExists {
			if _, r := rightExists[k]; r {
				total += shift(k)
				if debug {
					fmt.Printf("k=%d (%s), shifted=%d, total=%d\n\n", k, string(k), shift(k), total)
				}
			}
		}
	}
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
	fmt.Printf("Total: %d\n", total)
}

func shift(r uint8) int {
	var shifted int
	if r >= 97 {
		shifted = int(r - 96)
	} else {
		shifted = int(r - 38)
	}
	return shifted
}
