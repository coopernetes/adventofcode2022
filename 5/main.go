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

type CrateStack struct {
	crates []rune
	position int
}

func (stk *CrateStack) peek() rune {
	var top rune
	if len(stk.crates) > 0 {
		top = stk.crates[len(stk.crates) - 1]
	}
	return top
}

func (stk *CrateStack) push(r rune) {
	newElements := make([]rune, cap(stk.crates) + 1)
	copy(newElements, stk.crates)
	newElements[len(newElements)-1] = r
	stk.crates = newElements
}

func (stk *CrateStack) pushMany(runes []rune) {
	newElements := append(stk.crates, runes...)
	stk.crates = newElements
}

func (stk *CrateStack) pop() rune {
	var item rune
	if (len(stk.crates) > 0) {
		item = stk.crates[len(stk.crates) - 1]
		newElements := make([]rune, cap(stk.crates) - 1)
		copy(newElements, stk.crates)
		stk.crates = newElements	
	}
	return item
}

func (stk *CrateStack) popOrdered(n int) []rune {
	items := make([]rune, n)
	copy(items, stk.crates[len(stk.crates)-n:len(stk.crates)])
	newElements := make([]rune, len(stk.crates) - n)
	copy(newElements, stk.crates)
	stk.crates = newElements
	return items
}

func (stk CrateStack) String() string {
	return fmt.Sprintf("%d:%s", stk.position, string(stk.crates))
}


func move(instruction string, stacks []*CrateStack) {
	parts := strings.SplitN(instruction, " ", 6)
	cratesToMove, _ := strconv.Atoi(parts[1])
	srcStackPos, _ := strconv.Atoi(parts[3])
	destStackPos, _ := strconv.Atoi(parts[5])
	log.Printf("Moving %d crates from pos %d to %d\n", cratesToMove, srcStackPos, destStackPos)
	for i := 0; i < cratesToMove; i++ {
		srcStack := stacks[srcStackPos-1]
		destStack := stacks[destStackPos-1]
		nextCrate := srcStack.peek()
		if nextCrate != 0 {
			crate := srcStack.pop()
			log.Printf("Crate %s\n", string(crate))
			destStack.push(crate)	
		}
	}
	log.Printf("CrateStack: %s\n", stacks)
}

func moveOrdered(instruction string, stacks []*CrateStack) {
	parts := strings.SplitN(instruction, " ", 6)
	cratesToMove, _ := strconv.Atoi(parts[1])
	srcStackPos, _ := strconv.Atoi(parts[3])
	destStackPos, _ := strconv.Atoi(parts[5])
	log.Printf("Moving %d crates from pos %d to %d\n", cratesToMove, srcStackPos, destStackPos)
	srcStack := stacks[srcStackPos-1]
	destStack := stacks[destStackPos-1]
	crates := srcStack.popOrdered(cratesToMove)
	log.Printf("Crates %s\n", string(crates))
	destStack.pushMany(crates)
	log.Printf("CrateStack: %s\n", stacks)
}


func main() {
	log.Println("Day 5")
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	crateInput, err := os.ReadFile(path.Join(thisDir, "crates.txt"))
	if err != nil {
		log.Fatal(err)
	}

	crateLines := strings.Split(string(crateInput[:]), "\n")

	instructionIn, err := os.ReadFile(path.Join(thisDir, "instructions.txt"))
	if err != nil {
		log.Fatal(err)
	}

	instructionLines := strings.Split(string(instructionIn[:]), "\n")

	start := time.Now()
	debug := true

	// initial array of CrateStacks using the position numbers
	// at the end of the file (the last line)
	lastline := crateLines[len(crateLines)-1]
	log.Println(lastline)
	var initialCrateStacks []*CrateStack
	for _, r := range lastline {
		log.Println(string(r))
		pos, err := strconv.Atoi(string(r));
		if err != nil {
			continue
		}
		cs := &CrateStack{
			crates: []rune{},
			position: pos,
		}
		initialCrateStacks = append(initialCrateStacks, cs)	
	}

	// each "crate entry" is separated in fixed length columns
	// so we grab each entry (even empty ones) every 4 characters
	crateEntryChunk := 4

	// crates are added to the stack in reverse (bottom of the file 
	// ascending up) order
	for i := len(crateLines)-2; i >= 0; i-- {
		line := crateLines[i]
		if debug {
			log.Printf("Line:\n\t%s\n", crateLines[i])
		}
		var entries []string
		for j := 0; j < len(line); j += crateEntryChunk {
			entries = append(entries, line[j:j+crateEntryChunk-1])
		}
		log.Printf("Entries (len: %d):\n\t%s\n", len(entries), entries)
		for i, e := range entries {
			if e != "   " {
				initialCrateStacks[i].push(rune(e[1]))
			}
		}
	}
	fmt.Printf("CrateStack start: %s\n", initialCrateStacks)
	log.Printf("CrateStack start: %s\n", initialCrateStacks)

	// part 1
	crates := make([]*CrateStack, len(initialCrateStacks))
	for i, c := range initialCrateStacks {
		v := *c
		crates[i] = &v
	}
	for _, instruction := range instructionLines {
		move(instruction, crates)
	}
	fmt.Printf("Part 1 CrateStack end: %+v\n", crates)
	log.Printf("Part 1 CrateStack end: %+v\n", crates)
	fmt.Print("Part 1 Result: ")
	for _, cs := range crates {
		if cs.peek() != 0 {
			fmt.Print(string(cs.peek()))
		}
	}
	fmt.Println()
	// part 2
	groupedCrates := make([]*CrateStack, len(initialCrateStacks))
	for i, c := range initialCrateStacks {
		v := *c
		groupedCrates[i] = &v
	}
	for _, instruction := range instructionLines[:] {
		moveOrdered(instruction, groupedCrates)
	}
	fmt.Printf("Part 2 CrateStack end: %s\n", groupedCrates)
	log.Printf("Part 2 CrateStack end: %s\n", groupedCrates)
	fmt.Print("Part 2 Result: ")
	for _, cs := range groupedCrates {
		if cs.peek() != 0 {
			fmt.Print(string(cs.peek()))
		}
	}

	done := time.Now()
	diff := done.Sub(start)
	log.Printf("\nExecution time: %d µSeconds\n", diff.Microseconds())
}
