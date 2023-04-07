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

func newCrateStack(initialCrates []rune, position int) *CrateStack {
	return &CrateStack{
		crates: initialCrates,
		position: position,
	}
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

func (stk CrateStack) String() string {
	return fmt.Sprintf("%d:%s", stk.position, string(stk.crates))
}


func move(instruction string, stacks []*CrateStack) {
	// move 2 from 8 to 2
	log.Printf("Instruction: %s\n", instruction)
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
	var crateStacks []*CrateStack
	for _, r := range lastline {
		log.Println(string(r))
		pos, err := strconv.Atoi(string(r));
		if err != nil {
			log.Print(err)
			continue
		}
		cs := newCrateStack([]rune{}, pos)
		crateStacks = append(crateStacks, cs)	
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
				crateStacks[i].push(rune(e[1]))
			}
		}
	}
	fmt.Printf("CrateStack start: %s\n", crateStacks)
	log.Printf("CrateStack start: %s\n", crateStacks)

	for _, instruction := range instructionLines {
		move(instruction, crateStacks)
	}
	fmt.Printf("CrateStack end: %+v\n", crateStacks)
	log.Printf("CrateStack end: %+v\n", crateStacks)
	fmt.Print("Result: ")
	for _, cs := range crateStacks {
		fmt.Print(string(cs.peek()))
	}
	done := time.Now()
	diff := done.Sub(start)
	log.Printf("\nExecution time: %d µSeconds\n", diff.Microseconds())
}
