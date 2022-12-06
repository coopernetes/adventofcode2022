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
	debug := true
	for _, line := range lines {
		if debug {
			fmt.Printf("%s\n", line)
		}

	}

	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}
