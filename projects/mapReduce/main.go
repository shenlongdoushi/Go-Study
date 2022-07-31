package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Read CSV File to []string
// Create a mapTask Chan to build Go Object based on []string (go routine)
// Create a reduceTask Chan to filter Go Object
// pipe and print the result

type Player struct {
	ID   int
	Name string
	Age  int
}

var (
	wg sync.WaitGroup
)

func main() {
	filename := "./testData/soccer.csv"
	lines := readCSV(filename)

	mapQ := make(chan []Player)
	reduceQ := make(chan []Player)

	wg.Add(len(lines))

	go reduceTask(mapQ, reduceQ)

	for _, v := range lines {
		go func(line []string) {
			defer wg.Done()
			mapQ <- mapTask(line)
		}(v)
	}

	wg.Wait()
	close(mapQ)
	fmt.Println(<-reduceQ)
	close(reduceQ)
}

func mapTask(line []string) []Player {
	output := []Player{}
	id, _ := strconv.Atoi(line[0])
	age, _ := strconv.Atoi(line[4])
	output = append(output, Player{
		ID:   id,
		Name: line[1],
		Age:  age,
	})
	return output
}

func reduceTask(mapQ chan []Player, reduceQ chan []Player) {
	output := []Player{}
	for item := range mapQ {
		for _, v := range item {
			if v.Age >= 35 {
				output = append(output, v)
			}
		}
	}
	reduceQ <- output
}

func readCSV(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}
