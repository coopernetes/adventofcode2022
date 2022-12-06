package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Day 4")
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	in, err := os.ReadFile(path.Join(thisDir, "input"))
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(in[:]), "\n")
	start := time.Now()
	debug := os.Getenv("DEBUG") != ""
	containsFully := 0
	overlapCount := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		left := strings.Split(pairs[0], "-")
		right := strings.Split(pairs[1], "-")
		l1, _ := strconv.Atoi(left[0])
		l2, _ := strconv.Atoi(left[1])
		r1, _ := strconv.Atoi(right[0])
		r2, _ := strconv.Atoi(right[1])
		leftRange := makeRange(l1, l2)
		rightRange := makeRange(r1, r2)
		if debug {
			fmt.Printf("left=%d-%d right=%d-%d\n", l1, l2, r1, r2)
		}
		if contains(leftRange, rightRange) {
			containsFully += 1
		}
		if overlaps(leftRange, rightRange) {
			if debug {
				fmt.Printf("\toverlaps (%d)\n", overlapCount)
			}
			overlapCount += 1
		}
	}

	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Ranges contained: %d\n", containsFully)
	fmt.Printf("Ranges overlapping: %d\n", overlapCount) // 779 = too low, 847 = too low, 910 = too high
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func contains(left []int, right []int) bool {
	if len(left) == 1 {
		return right[0] <= left[0] && left[0] <= right[len(right)-1]
	}
	if len(right) == 1 {
		return left[0] <= right[0] && right[0] <= left[len(left)-1]
	}
	return left[0] >= right[0] && left[len(left)-1] <= right[len(right)-1] ||
		right[0] >= left[0] && right[len(right)-1] <= left[len(left)-1]
}

func overlaps(left []int, right []int) bool {
	for _, i := range left {
		for _, j := range right {
			if i == j {
				return true
			}
		}
	}
	return false
	// I can't believe I wrote the code below (which didn't work btw Q_Q)
	//if len(left) == 1 {
	//	return right[0] <= left[0] && left[0] <= right[len(right)-1]
	//}
	//if len(right) == 1 {
	//	return left[0] <= right[0] && right[0] <= left[len(right)-1]
	//}
	//if left[0] < right[0] && left[len(left)-1] <= right[len(right)-1] {
	//	return true
	//}
	//if right[0] < left[0] && right[len(right)-1] <= left[len(left)-1] {
	//	return true
	//}
	//if left[0] < right[0] {
	//	return left[len(left)-1] >= right[0]
	//}
	//if right[0] < left[0] {
	//	return right[len(right)-1] >= left[0]
	//}
	//return false
}

func makeRange(lower int, upper int) []int {
	r := make([]int, upper-lower+1)
	var i = 0
	for i = 0; i < (upper - lower + 1); i++ {
		r[i] = lower + i
	}
	return r
}
