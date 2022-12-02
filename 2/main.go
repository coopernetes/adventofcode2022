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

	//choices := map[string]int{"X": 1, "Y": 2, "Z": 3, "A": 1, "B": 2, "C": 3}
	//winners := map[string]string{"X": "C", "Y": "A", "Z": "B"}
	expected := map[string]map[string]int{"X": {"A": 3, "B": 1, "C": 2}, "Y": {"A": 1, "B": 2, "C": 3}, "Z": {"A": 2, "B": 3, "C": 1}}
	// A,X(1) rock -> beats C,Z(3) scissors
	// B,Y(2) paper -> beats A,X(1) rock
	// C,Z(3) scissors -> beats B,Y(2) paper
	// lose = 0, draw = 3, win  = 6
	// X = lose, Y = draw, Z = win
	var sum int
	lines := strings.Split(string(in[:]), "\n")
	fmt.Printf("Line count: %d\n", len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		theirs := parts[0]
		outcome := parts[1]
		if outcome == "Y" {
			sum += 3
		}
		if outcome == "Z" {
			sum += 6
		}
		sum += expected[outcome][theirs]
		//mine := parts[1]
		//sum += choices[mine]
		//if choices[mine] == choices[theirs] {
		//	sum += 3
		//	continue
		//}
		//if strings.Compare(theirs, winners[mine]) == 0 {
		//	sum += 6
		//}
	}
	log.Printf("%d", sum)
}
