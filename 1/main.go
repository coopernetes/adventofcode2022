package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	in, err := os.ReadFile(path.Join(thisDir, "input"))
	if err != nil {
		log.Fatal(err)
	}
	data := []int{0}
	parts := strings.Split(string(in[:]), "\n")
	fmt.Printf("Parts count: %d\n", len(parts))
	var i int
	j := 0
	for i = 0; i < len(parts); i++ {
		line := parts[i]
		if line == "" {
			j++
			data = append(data, 0)
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data[j] += num
	}
	sort.Ints(data)
	topData := []int{0, 0, 0}
	copy(topData, data[len(data)-3:])
	log.Printf("%d", topData)
	var sum int
	for _, d := range topData {
		sum += d
	}
	log.Printf("%d", sum)
}
