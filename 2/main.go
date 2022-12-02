package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	fmt.Println("Day 2")
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	in, err := os.ReadFile(path.Join(thisDir, "input"))
	if err != nil {
		log.Fatal(err)
	}

	choices := map[string]int{"X": 1, "Y": 2, "Z": 3, "A": 1, "B": 2, "C": 3}
	winners := map[string]string{"X": "C", "Y": "A", "Z": "B"}
	// A,X rock -> beats C,Z scissors
	// B,Y paper -> beats A,X rock
	// C,Z scissors -> beats B,Y paper
	// lose = 0, draw = 3, win  = 6
	var sum int
	lines := strings.Split(string(in[:]), "\n")
	fmt.Printf("Line count: %d\n", len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		theirs := parts[0]
		mine := parts[1]
		sum += choices[mine]
		if choices[mine] == choices[theirs] {
			sum += 3
			continue
		}
		if strings.Compare(theirs, winners[mine]) == 0 {
			sum += 6
		}
	}
	log.Printf("%d", sum)
}
