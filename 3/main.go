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
	var total int32
	fmt.Printf("Line count: %d\n", len(lines))

	start := time.Now()
	debug := true
	for _, line := range lines {
		mid := len(line) / 2
		left := line[:mid]
		right := line[mid:]
		if debug {
			fmt.Printf("left=%s right=%s\n", left, right)
		}

		var uniqueMatches = make(map[int32]bool)
		for _, l := range left {
			for _, r := range right {
				if l == r {
					uniqueMatches[shift(l)] = true

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
	var badgeSum int32
	var i int
	for i = 0; i < len(lines); i += 3 {
		var g1Chars = make(map[int32]bool)
		var g2Chars = make(map[int32]bool)
		var g3Chars = make(map[int32]bool)

		for _, c := range lines[i] {
			g1Chars[c] = true
		}
		for _, c := range lines[i+1] {
			g2Chars[c] = true
		}
		for _, c := range lines[i+2] {
			g3Chars[c] = true
		}
		if debug {
			fmt.Printf("%v\n%v\n%v\n", g1Chars, g2Chars, g3Chars)
		}
		for k, _ := range g1Chars {
			_, inG2 := g2Chars[k]
			_, inG3 := g3Chars[k]
			if inG2 && inG3 {
				badgeSum += shift(k)
			}
		}
	}
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Badge Sum: %d\n", badgeSum)
}

func shift(r int32) int32 {
	var shifted int32
	if r >= 97 {
		shifted = r - 96
	} else {
		shifted = r - 38
	}
	return shifted
}
