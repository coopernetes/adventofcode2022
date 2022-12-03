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
		debug := false
		mid := len(line) / 2
		left := line[:mid]
		right := line[mid:]
		if debug {
			fmt.Printf("left=%s right=%s\n", left, right)
		}

		var uniqueMatches = make(map[int]bool)
		for _, l := range left {
			for _, r := range right {
				if l == r {
					uniqueMatches[shift(uint8(l))] = true

				}
			}
		}
		for u, _ := range uniqueMatches {
			total += u
			if debug {
				fmt.Printf("u=%d total=%d\n\n", u, total)
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
