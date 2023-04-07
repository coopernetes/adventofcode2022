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
	log.Println("Day n")
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	in, err := os.ReadFile(path.Join(thisDir, "input.txt"))
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
	log.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}
