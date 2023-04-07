package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func isUnique(chunk string) bool {
	seen := make([]rune, 0)
	for _, r := range chunk {
		for _, s := range seen {
			if r == s {
				return false
			}
		}
		seen = append(seen, r)
	}
	return true
}


func main() {
	log.Println("Day 6")
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	in, err := os.ReadFile(path.Join(thisDir, "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	data := string(in[:])
	start := time.Now()

	chunkSize := 4
	for i := 0; i < len(data); i++ {
		log.Printf("Checking chunk %s at idx %d", data[i:i+chunkSize], i+chunkSize)
		if i+chunkSize <= len(data) && isUnique(data[i:i+chunkSize]) {
			fmt.Printf("Found unique chunk at idx %d: %s\n", i+chunkSize, data[i:i+chunkSize])
			break
		}
	}

	chunkSize = 14
	for i := 0; i < len(data); i++ {
		log.Printf("Checking chunk %s at idx %d", data[i:i+chunkSize], i+chunkSize)
		if i+chunkSize <= len(data) && isUnique(data[i:i+chunkSize]) {
			fmt.Printf("Found unique chunk at idx %d: %s\n", i+chunkSize, data[i:i+chunkSize])
			break
		}
	}
	
	done := time.Now()
	diff := done.Sub(start)
	log.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}
